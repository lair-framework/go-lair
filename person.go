package lair

// Person is used to store information related to a single person
// for a project. Most fields map to AD. LastloggedIn is used to create
// a loose relationship to a Host by ip.
type Person struct {
	ID                string            `json:"_id" bson:"_id"`
	ProjectID         string            `json:"projectId" bson:"projectId"`
	PrincipalName     string            `json:"principalName" bson:"principalName"`
	SAMAccountName    string            `json:"samAccountName" bson:"samAccountName"`
	DistinguishedName string            `json:"distinguishedName" bson:"distinguishedName"`
	FirstName         string            `json:"firstName" bson:"firstName"`
	MiddleName        string            `json:"middleName" bson:"middleName"`
	LastName          string            `json:"lastName" bson:"lastName"`
	DisplayName       string            `json:"displayName" bson:"displayName"`
	Department        string            `json:"department" bson:"department"`
	Description       string            `json:"description" bson:"description"`
	Address           string            `json:"address" bson:"address"`
	Emails            []string          `json:"emails" bson:"emails"`
	Phones            []string          `json:"phones" bson:"phones"`
	References        []PersonReference `json:"references" bson:"references"`
	Groups            []string          `json:"groups" bson:"groups"`
	LastLogon         string            `json:"lastLogon" bson:"lastLogon"`
	LastLogoff        string            `json:"lastLogoff" bson:"lastLogoff"`
	LoggedIn          []string          `json:"loggedIn" bson:"loggedIn"`
}

// PersonReference is used to link a person to a 3rd party site.
type PersonReference struct {
	Description string `json:"description" bson:"description"`
	Username    string `json:"username" bson:"username"`
	Link        string `json:"link" bson:"link"`
}
