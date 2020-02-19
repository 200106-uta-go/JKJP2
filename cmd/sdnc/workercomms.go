package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Send hash to each worker's overlay ip address.
func sendToWorkers(hash string, workerAddrs []string) *http.Response {
	var workerCount int64 = int64(len(workerAddrs))  // How many slave ips have we registered?
	var pages int64 = dictionaryLength / workerCount // How big will the assigned tasks will be
	myResponse := make(chan *http.Response, 0)

	// Iterate over worker addresses and submit a hash, startindex, and length of entries.
	for i := 0; i < len(workerAddrs); i++ {
		startIndex := int64(i) * pages

		// This go routine submits values to the PODS not the ec2s.
		go func(index int) {

			resp, er := tryGet(workerAddrs[index], hash, startIndex, pages, 60)
			if er == nil {
				myResponse <- resp
			}
			log.Fatalf("Timeout: failed to connect - %v", er)

		}(i)
	}
	return <-myResponse
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, index int64, length int64, t int) (*http.Response, error) {
	var er error
	var colliderPort string = "8080"

	for i := 0; i < t; i++ {

		// Submit request to colliders. Post request will be used because GET query and
		// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
		// myURL.PasreForm()
		colliderURL := fmt.Sprintf("%s:%s/?hash=%s&index=%s&length=%s", addr, colliderPort, hash, string(index), string(length))
		// contentType := "application/x-www-form-urlencoded"
		// content := fmt.Sprintf("hash=%s&index=%s&length=%s", hash, index, length)

		req, er := http.NewRequest("GET", colliderURL, nil)
		client := http.Client{}
		resp, er := client.Do(req)
		// Normal operation
		if er == nil {
			return resp, er
		}

		// Print error to standard console and wait to try again.
		fmt.Println(er)
		time.Sleep(time.Second)

	}

	// We did not connect to collider in time. Don't return a resp back to sendToWorkers.
	// Return error from POST request.
	return nil, er

}
