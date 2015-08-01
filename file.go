package lair

// File is used to share links to files for a project.
type File struct {
	FileName string `json:"fileName" bson:"fileName"`
	URL      string `json:"url" bson:"url"`
}
