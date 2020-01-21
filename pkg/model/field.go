package model

var (
	// StringType represents a string
	StringType FieldType = FieldType{
		Name:           "string",
		VariableLength: true,
		FixedLength:    -1,
	}
	// BooleanType represents a boolean
	BooleanType FieldType = FieldType{
		Name:           "bool",
		VariableLength: false,
		FixedLength:    1,
	}
	// Int64Type represents a int64
	Int64Type FieldType = FieldType{
		Name:           "int64",
		VariableLength: false,
		FixedLength:    8,
	}
	// Float64Type represents a float64
	Float64Type FieldType = FieldType{
		Name:           "float64",
		VariableLength: false,
		FixedLength:    8,
	}
	// BinaryType represents a byte[]
	BinaryType FieldType = FieldType{
		Name:           "binary",
		VariableLength: true,
		FixedLength:    -1,
	}
)

// Field The field, defining a name/value associated with an event
type Field struct {
	name      string
	fieldType FieldType
	value     []byte
}

// NewField returns a new field with name, field type and value
func NewField(name string, fieldType FieldType, value []byte) *Field {
	return &Field{
		name:      name,
		fieldType: fieldType,
		value:     value,
	}
}

// Name the field name
func (field *Field) Name() string {
	return field.name
}

// Type the field type
func (field *Field) Type() FieldType {
	return field.fieldType
}

// Value the field value
func (field *Field) Value() []byte {
	return field.value
}
