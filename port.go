package lair

// Port contains all the needed information for a single port.
type Port struct {
	Id             string       `json:"_id" bson:"_id"`
	ProjectId      string       `json:"project_id" bson:"project_id"`
	HostId         string       `json:"host_id" bson:"host_id"`
	Port           int          `json:"port" bson:"port"`
	Protocol       string       `json:"protocol" bson:"protocol"`
	Service        string       `json:"service" bson:"service"`
	Product        string       `json:"product" bson:"product"`
	Alive          bool         `json:"alive" bson:"alive"`
	Status         string       `json:"status" bson:"status"`
	Credentials    []Credential `json:"credentials" bson:"credentials"`
	LastModifiedBy string       `json:"last_modified_by" bson:"last_modified_by"`
	Notes          []Note       `json:"notes" bson:"notes"`
	Flag           bool         `json:"flag" bson:"flag"`
}

// Credential contains a username, password, and hash for a single port.
type Credential struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Hash     string `json:"hash" bson:"hash"`
}
