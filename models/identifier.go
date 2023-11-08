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

func (id *Identifier) Dup() *Identifier {
	newId := NewIdentifier(id.PropertyID, id.Value)
	return &newId
}

type ByIdentifier []Identifier

func (ids ByIdentifier) Len() int {
	return len(ids)
}
func (ids ByIdentifier) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}
func (ids ByIdentifier) Less(i, j int) bool {
	if ids[i].PropertyID != ids[j].PropertyID {
		return ids[i].PropertyID < ids[j].PropertyID
	}
	return ids[i].Value < ids[j].Value
}
