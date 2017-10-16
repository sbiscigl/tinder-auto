package tinder

/*User a type for liking someone*/
type User struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Distance int    `json:"distance_mi"`
}
