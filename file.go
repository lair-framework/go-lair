package lair

// File is used to share links to files for a project.
type File struct {
	ID        string `json:"_id" bson:"_id"`
	HostID    string `json:"hostId" bson:"hostId"`
	ServiceID string `json:"serviceId" bson:"serviceId"`
	IssueID   string `json:"issueId" bson:"issueId"`
	FileName  string `json:"fileName" bson:"fileName"`
	URL       string `json:"url" bson:"url"`
}
