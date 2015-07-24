package lair

type Credential struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Format   string `json:"format" bson:"format"`
	Hash     string `json:"hash" bson:"hash"`
	Host     string `json:"host" bson:"host"`
	Service  string `json:"service" bson:"service"`
}
