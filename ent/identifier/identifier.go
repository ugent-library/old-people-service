// Code generated by ent, DO NOT EDIT.

package identifier

const (
	// Label holds the string label denoting the identifier type in the database.
	Label = "identifier"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the identifier in the database.
	Table = "identifiers"
)

// Columns holds all SQL columns for identifier fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
