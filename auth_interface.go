package lair

type AuthInterface struct {
	ID            string `json:"_id" bson:"_id"`
	ProjectID     string `json:"projectId" bson:"projectId"`
	IsMultifactor bool   `json:"isMultifactor" bson:"isMultifactor"`
	Kind          string `json:"kind" bson:"kind"`
	URL           string `json:"url" bson:"url"`
	Description   string `json:"description" bson:"description"`
}
