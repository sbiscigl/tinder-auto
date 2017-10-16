package conf

import (
	"flag"
	"fmt"
)

const (
	/*FbID const for type of facebook id number*/
	FbID = "fb_id"
	/*FbToken const for type of facebook token*/
	FbToken = "fb_token"
)

/*Config type for passing around*/
type Config struct {
	configMap map[string]string
}

/*New returns a new instance of config map*/
func New() *Config {
	fbID := flag.String(FbID,
		"foo",
		"facebook id number")
	fbToken := flag.String(FbToken,
		"bar",
		"facebook id number")
	flag.Parse()
	flagMap := make(map[string]string, 0)
	flagMap[FbID] = *fbID
	flagMap[FbToken] = *fbToken
	return &Config{
		configMap: flagMap,
	}
}

/*GetFbID gets the fb id from the map*/
func (c *Config) GetFbID() string {
	val, ok := c.configMap[FbID]
	if ok != true {
		panic(fmt.Sprintf("face book id not passed in, pass in with -fb_id"))
	}
	return val
}

/*GetFbToken gets the fb token from the map*/
func (c *Config) GetFbToken() string {
	val, ok := c.configMap[FbToken]
	if ok != true {
		panic(fmt.Sprintf("face book id not passed in, pass in with -fb_token"))
	}
	return val
}
