package models

type Identifier struct {
	PropertyID string `json:"property_id"`
	Value      string `json:"value"`
}

func NewIdentifier(typ string, val string) Identifier {
	return Identifier{
		PropertyID: typ,
		Value:      val,
	}
}
