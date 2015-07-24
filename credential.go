package lair

// Credential is some representation of a password All or none of the fields may be populated.
// Host and Service may be used to created loose relationships to host by ip and/or port.
// Username can be used to create a loose relationship to a person by principal name.
type Credential struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Format   string `json:"format" bson:"format"`
	Hash     string `json:"hash" bson:"hash"`
	Host     string `json:"host" bson:"host"`
	Service  string `json:"service" bson:"service"`
}
