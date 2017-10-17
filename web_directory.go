package lair

// WebDirectory is a path from a web listener and is a child
// of a Host. Port is not a relationship to a Service.
type WebDirectory struct {
	HostID           string `json:"hostId" bson:"hostId"`
	ID               string `json:"_id" bson:"_id"`
	IsFlagged        bool   `json:"isFlagged" bson:"isFlagged"`
	LastModifiedBy   string `json:"lastModifiedBy" bson:"lastModifiedBy"`
	Path             string `json:"path" bson:"path"`
	Port             int    `json:"port" bson:"port"`
	ProjectID        string `json:"projectId" bson:"projectId"`
	RequestData      string `json:"requestData,omitempty" bson:"requestData,omitempty"`
	RequestHeaders   string `json:"requestHeaders,omitempty" bson:"requestHeaders,omitempty"`
	RequestMethod    string `json:"requestMethod,omitempty" bson:"requestMethod,omitempty"`
	RequestProtocol  string `json:"requestProtocol,omitempty" bson:"requestProtocol,omitempty"`
	ResponseBody     string `json:"responseBody,omitempty" bson:"responseBody,omitempty"`
	ResponseCode     string `json:"responseCode" bson:"responseCode"`
	ResponseLength   uint32 `json:"responseLength,omitempty" bson:"responseLength,omitempty"`
	ResponseLocation string `json:"responseLocation,omitempty" bson:"responseLocation,omitempty"`
}
