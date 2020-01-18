package model

// FieldType Information about the type of a field
type FieldType struct {

	// Name The name of the field type
	Name string

	// VariableLength Whether the field is variable length
	VariableLength bool

	// FixedLength If fixed length field, otherwise -1
	FixedLength int
}
