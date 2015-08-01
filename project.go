package lair

// Project is the root for all information pertaining to a Lair project.
type Project struct {
	ID             string          `json:"_id" bson:"_id"`
	Name           string          `json:"name" bson:"name"`
	Industry       string          `json:"industry" bson:"industry"`
	CreatedAt      string          `json:"createdAt" bson:"createdAt"`
	Description    string          `json:"description" bson:"description"`
	Owner          string          `json:"owner" bson:"owner"`
	Contributors   []string        `json:"contributors" bson:"contributors"`
	Commands       []Command       `json:"commands" bson:"commands"`
	Notes          []Note          `json:"notes" bson:"notes"`
	DroneLog       []string        `json:"droneLog" bson:"droneLog"`
	Tool           string          `json:"tool"`
	Hosts          []Host          `json:"hosts"`
	Issues         []Issue         `json:"issues"`
	AuthInterfaces []AuthInterface `json:"authInterfaces"`
	Netblocks      []Netblock      `json:"netblocks"`
	People         []Person        `json:"people"`
	Credentials    []Credential    `json:"credentials"`
}

// Command is used to track tool commands that were run.
type Command struct {
	Tool    string `json:"tool" bson:"tool"`
	Command string `json:"command" bson:"command"`
}
