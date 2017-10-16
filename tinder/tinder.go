package tinder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sbiscigl/tinder-auto/conf"
)

var headerMap = map[string]string{
	"user-agent":   "Tinder/4.0.9 (iPhone; iOS 8.0.2; Scale/2.00)",
	"content-type": "application/json",
}

/*Tinder a type for interfacing with tinder api*/
type Tinder struct {
	auth string
}

/*New a new instane of a tinder object*/
func New(conf *conf.Config) *Tinder {
	client := &http.Client{}

	fa := NewFBAuth(conf.GetFbID(), conf.GetFbToken())
	req, err := http.NewRequest("POST", "https://api.gotinder.com/auth",
		bytes.NewReader(fa.ToJSON()))
	if err != nil {
		panic("could not create request")
	}
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic("could not execute response")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("could not get api token")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("could notread body")
	}

	ta := fromJSON(body)
	return &Tinder{
		auth: ta.AuthToken,
	}
}

/*GetRecs gets a list of reccomendations*/
func (t *Tinder) getRecs() *Reccomendations {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.gotinder.com/recs", nil)
	if err != nil {
		panic("could not create request")
	}
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	req.Header.Add("X-Auth-Token", t.auth)
	resp, err := client.Do(req)
	if err != nil {
		panic("could not execute response")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		panic(fmt.Sprintf("could not get reccomendations with staus code %d",
			resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("could not read body for recomendations")
	}
	return FromJSON(body)
}

/*LikeUsers likes users until i cant anymore*/
func (t *Tinder) LikeUsers() {
	client := &http.Client{}
	recs := t.getRecs().Users
	noneLeft := false
	for noneLeft == false {
		if len(recs) <= 0 {
			log.Println("no reccomendations left")
			noneLeft = true
		} else {
			for _, usr := range recs {
				reqString := "https://api.gotinder.com/like/" + usr.ID
				req, err := http.NewRequest("GET", reqString, nil)
				if err != nil {
					panic("could not create request")
				}
				for k, v := range headerMap {
					req.Header.Add(k, v)
				}
				req.Header.Add("X-Auth-Token", t.auth)
				time.Sleep(time.Second)
				resp, err := client.Do(req)
				if err != nil {
					panic("could not execute response")
				}
				defer resp.Body.Close()
				if resp.StatusCode != 200 {
					fmt.Println()
					panic("could not like")
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(fmt.Sprintf("could not read body"))
				}
				lr := LikeFromJSON(body)
				log.Printf("Just liked %s who is %d miles away ", usr.Name, usr.Distance)
				if !lr.Match {
					log.Printf("did not match. ")
				} else {
					log.Printf("was a match! ")
				}
				log.Printf("%d likes left.", lr.Likes)
			}
			recs = t.getRecs().Users
		}
	}
}
