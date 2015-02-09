package lair

// Host is a single Lair host.
type Host struct {
	Id             string   `json"_id" bson:"_id"`
	ProjectId      string   `json:"project_id" bson:"project_id"`
	LongAddr       uint32   `json:"long_addr" bson:"long_addr"`
	StringAddr     string   `json:"string_addr" bson:"string_addr"`
	MacAddr        string   `json:"mac_addr" bson:"mac_addr"`
	Hostnames      []string `json:"hostnames" bson:"hostnames"`
	OS             []OS     `json:"os" bson:"os"`
	Notes          []Note   `json:"notes" bson:"notes"`
	Alive          bool     `json:"alive" bson:"alive"`
	Status         string   `json:"status" bson:"status"`
	LastModifiedBy string   `json:"last_modified_by" bson:"last_modified_by"`
	Flag           bool     `json:"flag" bson:"flag"`
}

// OS fingerprint for a host.
type OS struct {
	Tool        string `json:"tool" bson:"tool"`
	Weight      int    `json:"weight" bson:"weight"`
	Fingerprint string `json:"fingerprint" bson:"fingerprint"`
}
