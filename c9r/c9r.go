package c9r

type Favorite struct {
	Author string `json:"author,omitempty" firestore:"author"`
	Site   string `json:"site,omitempty" firestore:"site"`
	Title  string `json:"title,omitempty" firestore:"title"`
}
