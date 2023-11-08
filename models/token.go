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

func (t *Token) Dup() *Token {
	t2 := NewToken(t.PropertyID, t.Value)
	return &t2
}

type ByToken []Token

func (tokens ByToken) Len() int {
	return len(tokens)
}
func (tokens ByToken) Swap(i, j int) {
	tokens[i], tokens[j] = tokens[j], tokens[i]
}
func (tokens ByToken) Less(i, j int) bool {
	if tokens[i].PropertyID != tokens[j].PropertyID {
		return tokens[i].PropertyID < tokens[j].PropertyID
	}
	return tokens[i].Value != tokens[j].Value
}
