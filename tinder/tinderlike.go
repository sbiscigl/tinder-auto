package tinder

import "encoding/json"

/*LikeResp response type of like action*/
type LikeResp struct {
	Likes            int   `json:"likes_remaining"`
	RateLimitedUntil int64 `json:"rate_limited_until"`
}

/*LikeFromJSON gets the object from the json*/
func LikeFromJSON(bytes []byte) *LikeResp {
	var like LikeResp
	if err := json.Unmarshal(bytes, &like); err != nil {
		panic(err)
	}
	return &like
}
