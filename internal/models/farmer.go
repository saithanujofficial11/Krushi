package models

type ContactInfo struct {
	PrimaryPhone string `json:"primaryPhone"`
	Address      string `json:"address"`
}

type Farmer struct {
		FarmerID    string      `json:"farmerId" bson:"farmer_id"`
	Name        string      `json:"name"`
	ContactInfo ContactInfo `json:"contactInfo"`
}
