package cityhashutil

// HashCollision describes a hash - collision pair
type HashCollision struct {
	InputHash string `json:"hash"`
	Collision string `json:"result"`
}
