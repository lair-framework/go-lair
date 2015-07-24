package lair

type WebDirectory struct {
	ID             string `json:"_id" bson:"_id"`
	ProjectID      string `json:"projectId" bson:"projectId"`
	HostID         string `json:"hostId" bson:"hostId"`
	Path           string `json:"path" bson:"path"`
	Port           int    `json:"port" bson:"port"`
	ResponseCode   string `json:"responseCode" bson:"responseCode"`
	LastModifiedBy string `json:"lastModifiedBy" bson:"lastModifiedBy"`
	IsFlagged      bool   `json:"isFlagged" bson:"isFlagged"`
}
