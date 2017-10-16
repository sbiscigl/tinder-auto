package tinder

import "encoding/json"

/*TAuth type for getting the tinder auth response*/
type TAuth struct {
	AuthToken string `json:"token"`
}

func fromJSON(bytes []byte) *TAuth {
	var auth TAuth
	if err := json.Unmarshal(bytes, &auth); err != nil {
		panic(err)
	}
	return &auth
}
