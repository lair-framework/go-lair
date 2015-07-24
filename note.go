package lair

// Note is used to hold information data about
// as project, host, or port.
type Note struct {
	Title          string `json:"title" bson:"title"`
	Content        string `json:"content" bson:"content"`
	LastModifiedBy string `json:"lastModifiedBy" bson:"lastModifiedBy"`
}
