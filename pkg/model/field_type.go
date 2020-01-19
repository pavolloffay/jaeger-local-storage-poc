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

// ConvBoolToBytes convert the boolean value to a byte array
func ConvBoolToBytes(value bool) []byte {
	b := make([]byte, 1)
	if value {
		b[0] = 1
	} else {
		b[0] = 0
	}
	return b
}
