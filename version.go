package lair

// Version is the major version of the lair document schema.
type Version struct {
	Value string `json:"version" bson:"version"`
}
