package tinder

import (
	"encoding/json"
	"fmt"
)

/*FacebookAuth auth type for facebook*/
type FacebookAuth struct {
	FbID    string `json:"facebook_id"`
	FbToken string `json:"facebook_token"`
}

/*NewFBAuth for new instance of facebook auth instance*/
func NewFBAuth(id string, token string) *FacebookAuth {
	return &FacebookAuth{
		FbID:    id,
		FbToken: token,
	}
}

/*ToJSON casts a faceboook auth instance into json*/
func (fa *FacebookAuth) ToJSON() []byte {
	byts, err := json.Marshal(fa)
	if err != nil {
		panic(fmt.Sprintf("could not marshall facebook auth"))
	}
	return byts
}
