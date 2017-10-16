package tinder

import "encoding/json"

/*LikeResp response type of like action*/
type LikeResp struct {
  Match bool `json:"match"`
  Likes int `json:"ikes_remaining"`
}

/*LikeFromJSON gets the object from the json*/
func LikeFromJSON(bytes []byte) *LikeResp {
	var like LikeResp
	if err := json.Unmarshal(bytes, &like); err != nil {
		panic(err)
	}
	return &like
}
