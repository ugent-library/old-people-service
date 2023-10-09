package models

type Identifier struct {
	PropertyID string `json:"property_id"`
	Value      string `json:"value"`
}

func NewIdentifier(propertyID string, value string) Identifier {
	return Identifier{
		PropertyID: propertyID,
		Value:      value,
	}
}
