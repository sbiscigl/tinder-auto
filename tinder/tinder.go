package tinder

import (
	"bytes"
	"log"
	"time"

	"github.com/sbiscigl/tinder-auto/conf"
	"github.com/sbiscigl/tinder-auto/util"
)

var headerMap = map[string]string{
	"user-agent":   "Tinder/4.0.9 (iPhone; iOS 8.0.2; Scale/2.00)",
	"content-type": "application/json",
}

/*Tinder a type for interfacing with tinder api*/
type Tinder struct {
	auth       string
	httpHelper *util.HTTPHelper
}

/*New a new instane of a tinder object*/
func New(conf *conf.Config) *Tinder {
	httpHelper := util.New()

	fa := NewFBAuth(conf.GetFbID(), conf.GetFbToken())
	body := httpHelper.MakeReq("POST", "https://api.gotinder.com/auth",
		bytes.NewReader(fa.ToJSON()), []util.HTTPHeader{})

	ta := fromJSON(body)

	return &Tinder{
		auth:       ta.AuthToken,
		httpHelper: httpHelper,
	}
}

/*GetRecs gets a list of reccomendations*/
func (t *Tinder) getRecs() *Reccomendations {
	body := t.httpHelper.MakeReq("GET", "https://api.gotinder.com/recs", nil,
		[]util.HTTPHeader{
			util.HTTPHeader{
				Key:   "X-Auth-Token",
				Value: t.auth,
			},
		},
	)
	return FromJSON(body)
}

func (t *Tinder) waitForLikesBack(waitTill int64) {
	waitTillMillis := waitTill / 1000
	log.Printf("You got rate blocked! this will last till : %v\n"+
		"program will wait", time.Unix(waitTillMillis, 0))

	currentTime := time.Now().Unix()
	for currentTime < waitTillMillis {
		time.Sleep(time.Hour)
		currentTime = time.Now().Unix()
		log.Printf("waited an hour, still waiting, currently:\n%v "+
			"waiting for:\n%v", currentTime, waitTillMillis)
	}
	log.Println("AND WE'RE BACK")
}

/*LikeUsers likes users until i cant anymore*/
func (t *Tinder) LikeUsers() {
	recs := t.getRecs().Users
	noneLeft := false
	for noneLeft == false {
		if len(recs) <= 0 {
			log.Println("no reccomendations left")
			noneLeft = true
		} else {
			for _, usr := range recs {
				/*stupid fucking rate limiting*/
				time.Sleep(time.Second)
				reqString := "https://api.gotinder.com/like/" + usr.ID
				body := t.httpHelper.MakeReq("GET", reqString, nil,
					[]util.HTTPHeader{
						util.HTTPHeader{
							Key:   "X-Auth-Token",
							Value: t.auth,
						},
					},
				)
				lr := LikeFromJSON(body)
				log.Printf("Just liked %s who is %d miles away ", usr.Name, usr.Distance)
				log.Printf("%d likes left.", lr.Likes)
				if lr.Likes == 0 {
					t.waitForLikesBack(lr.RateLimitedUntil)
				}
			}
			recs = t.getRecs().Users
		}
	}
}
