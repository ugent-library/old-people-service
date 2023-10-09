package models

type Token struct {
	PropertyID string `json:"property_id"`
	Value      string `json:"value"`
}

func NewToken(propertyID string, value string) Token {
	return Token{
		PropertyID: propertyID,
		Value:      value,
	}
}
