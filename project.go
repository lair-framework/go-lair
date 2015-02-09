package lair

// Project is the root for all information pertaining to a Lair project.
type Project struct {
	Id           string    `json:"_id" bson:"_id"`
	ProjectName  string    `json:"project_name" bson:"project_name"`
	Industry     string    `json:"industry" bson:"industry"`
	CreationDate string    `json:"creation_date" bson:"creation_date"`
	Description  string    `json:"description" bson:"description"`
	Owner        string    `json:"description" bson:"owner"`
	Contributors []string  `json"contributors" bson:"contributors"`
	Commands     []string  `json:"commands" bson:"commands"`
	Notes        []Note    `json:"notes" bson:"notes"`
	DroneLog     []string  `json:"drone_log" bson:"drone_log"`
	Messages     []Message `json:"messages" bson:"messages"`
	Files        []File    `json:"files" bson:"files"`
}

// Message is for sending messages to other users in the project.
type Message struct {
	User    string `json:"user" bson:"user"`
	Message string `json:"message" bson:"message"`
}

// File is used to share links to files for a project.
type File struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}
