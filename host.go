package lair

// Host is a single Lair host.
type Host struct {
	ID             string    `json:"_id" bson:"_id"`
	ProjectID      string    `json:"projectId" bson:"projectId"`
	LongIPv4Addr   uint64    `json:"longIpv4Addr" bson:"longIpv4Addr"`
	IPv4           string    `json:"ipv4" bson:"ipv4"`
	MAC            string    `json:"mac" bson:"mac"`
	Hostnames      []string  `json:"hostnames" bson:"hostnames"`
	OS             OS        `json:"os" bson:"os"`
	Notes          []Note    `json:"notes" bson:"notes"`
	StatusMessage  string    `json:"statusMessage" bson:"statusMessage"`
	Tags           []string  `json:"tags" bson:"tags`
	Status         string    `json:"status" bson:"status"`
	LastModifiedBy string    `json:"lastModifiedBy" bson:"lastModifiedBy"`
	IsFlagged      bool      `json:"isFlagged" bson:"isFlagged"`
	Services       []Service `json:"services"`
}

// OS fingerprint for a host.
type OS struct {
	Tool        string `json:"tool" bson:"tool"`
	Weight      int    `json:"weight" bson:"weight"`
	Fingerprint string `json:"fingerprint" bson:"fingerprint"`
}
