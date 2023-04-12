package es6

type indexedOrganization struct {
	Id          string              `json:"id"`
	Type        string              `json:"type"`
	OtherId     map[string][]string `json:"other_id,omitempty"`
	ParentId    string              `json:"parent_id,omitempty"`
	Tree        []map[string]string `json:"tree,omitempty"`
	NameDut     string              `json:"name_dut,omitempty"`
	NameEng     string              `json:"name_eng,omitempty"`
	DateCreated string              `json:"date_created"`
	DateUpdated string              `json:"date_updated"`
}
