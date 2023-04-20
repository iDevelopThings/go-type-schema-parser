package main

type TypeDef struct {
	TypeName string
	Tag      string
	Fields   []Field
}

type Field struct {
	Name     string
	Type     string
	Nullable bool
}

type FieldType string

const (
	ID       FieldType = "ID"
	String   FieldType = "string"
	Int      FieldType = "int"
	Float    FieldType = "float"
	Bool     FieldType = "bool"
	Date     FieldType = "date"
	DateTime FieldType = "datetime"
	ArrayInt FieldType = "array<int>"
)
