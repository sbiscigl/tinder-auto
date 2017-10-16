package main

import (
	"github.com/sbiscigl/tinder-auto/conf"
	"github.com/sbiscigl/tinder-auto/tinder"
)

func main() {
	conf := conf.New()
	tinderAccess := tinder.New(conf)
	tinderAccess.LikeUsers()
}
