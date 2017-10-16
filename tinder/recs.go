package tinder

import "encoding/json"

/*Reccomendations a structure for storing reccomnedations*/
type Reccomendations struct {
	Status int    `json:"status"`
	Users  []User `json:"results"`
}

/*FromJSON gets the object from the json*/
func FromJSON(bytes []byte) *Reccomendations {
	var recs Reccomendations
	if err := json.Unmarshal(bytes, &recs); err != nil {
		panic(err)
	}
	return &recs
}
