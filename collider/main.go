package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/200106-uta-go/JKJP2/collider/proc"
)

func main() {

	// Get service port value from environment variable.
	//port := os.Getenv("SERV_PORT")

	// Response function initiator.
	http.HandleFunc("/", proc.Resp)

	// Listen http via service port as definded in enviroment variable.
	fmt.Printf("Server is running on port %s\n", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}
