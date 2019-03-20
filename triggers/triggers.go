package triggers

type Favorite struct {
	Author string `json:"author,omitempty" firestore:"author"`
	BookID string `json:"bookID,omitempty" firestore:"bookID"`
	Site   string `json:"site,omitempty" firestore:"site"`
	Title  string `json:"title,omitempty" firestore:"title"`
}
