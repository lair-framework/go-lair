package lair

// Issue contains all the needed information for a single
// issue in a project.
type Issue struct {
	ID             string           `json:"_id" bson:"_id"`
	ProjectID      string           `json:"projectId" bson:"projectId"`
	Title          string           `json:"title" bson:"title"`
	CVSS           float64          `json:"cvss" bson:"cvss"`
	Rating         string           `json:"rating" bson:"rating"`
	IsConfirmed    bool             `json:"isConfirmed" bson:"isConfirmed"`
	Description    string           `json:"description" bson:"description"`
	Evidence       string           `json:"evidence" bson:"evidence"`
	Solution       string           `json:"solution" bson:"solution"`
	Hosts          []IssueHost      `json:"hosts" bson:"hosts"`
	PluginIDs      []PluginID       `json:"pluginIds" bson:"pluginIds"`
	CVEs           []string         `json:"cves" bson:"cves"`
	References     []IssueReference `json:"references" bson:"references"`
	IdentifiedBy   []IdentifiedBy   `json:"identified_by" bson:"identifiedBy"`
	IsFlagged      bool             `json:"isFlagged" bson:"isFlagged"`
	Status         string           `json:"status" bson:"status"`
	LastModifiedBy string           `json:"lastModifiedBy "bson:"lastModifiedBy"`
	Notes          []Note           `json:"notes" bson:"notes"`
	Files          []File           `json:"files" bson:"filed"`
}

// IssueHost is a single host record for a issue.
type IssueHost struct {
	IPv4     string `json:"ipv4" bson:"ipv4"`
	Port     int    `json:"port" bson:"port"`
	Protocol string `json:"protocol" bson:"protocol"`
}

// PluginID is used to store a id from an automated tool.
type PluginID struct {
	Tool string `json:"tool" bson:"tool"`
	ID   string `json:"id" bson:"id"`
}

// IssueReference is a 3rd party containing issue information
type IssueReference struct {
	Link string `json:"link" bson:"link"`
	Name string `json:"name" bson:"name"`
}

// IdentifiedBy is used to show what identfied the issue.
type IdentifiedBy struct {
	Tool string `json:"tool"`
}
