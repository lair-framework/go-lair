package lair

// Netblock is used to store Whois data.
type Netblock struct {
	ID             string `json:"_id" bson:"_id"`
	ProjectID      string `json:"projectId" bson:"projectId"`
	ASN            string `json:"asn" bson:"asn"`
	ASNCountryCode string `json:"asnCountryCode" bson:"asnCountryCode"`
	ASNCIDR        string `json:"asnCidr" bson:"asnCidr"`
	ASNDate        string `json:"asnDate" bson:"asnDate"`
	ASNRegistry    string `json:"asnRegistry" bson:"asnRegistry"`
	CIDR           string `json:"cidr" bson:"cidr"`
	AbuseEmails    string `json:"abuseEmails" bson:"abuseEmails"`
	MiscEmails     string `json:"miscEmails" bson:"miscEmails"`
	TechEmails     string `json:"techEmails" bson:"techEmails"`
	Name           string `json:"name" bson:"name"`
	City           string `json:"city" bson:"city"`
	Country        string `json:"country" bson:"country"`
	PostalCode     string `json:"postalCode" bson:"postalCode"`
	Created        string `json:"created" bson:"created"`
	Updated        string `json:"updated" bson:"updated"`
	Description    string `json:"description" bson:"description"`
	Handle         string `json:"handle" bson:"handle"`
}
