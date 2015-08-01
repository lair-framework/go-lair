package lair

// Service contains all the needed information for a single service.
type Service struct {
	ID             string `json:"_id" bson:"_id"`
	ProjectID      string `json:"projectId" bson:"projectId"`
	HostID         string `json:"hostId" bson:"hostId"`
	Port           int    `json:"port" bson:"port"`
	Protocol       string `json:"protocol" bson:"protocol"`
	Service        string `json:"service" bson:"service"`
	Product        string `json:"product" bson:"product"`
	Status         string `json:"status" bson:"status"`
	IsFlagged      bool   `json:"isFlagged" bson:"isFlagged"`
	LastModifiedBy string `json:"lastModifiedBy" bson:"lastModifiedBy"`
	Notes          []Note `json:"notes" bson:"notes"`
	Files          []File `json:"files" bson:"filed"`
}
