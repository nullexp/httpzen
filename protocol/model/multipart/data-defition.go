package multipart

import http "github.com/nullexp/httpzen/protocol"

type DataDefinition struct {
	Name     string
	Object   http.Verifier
	Optional bool
	Single   bool
}

func (f *DataDefinition) GetPartName() string {
	return f.Name
}

func (f *DataDefinition) GetSupportedTypes() []string {
	return []string{}
}

func (f *DataDefinition) GetObject() interface{} {
	return f.Object
}

func (f *DataDefinition) IsSingle() bool {
	return f.Single
}

func (f *DataDefinition) IsOptional() bool {
	return f.Optional
}

func (f *DataDefinition) Verify() error {
	return f.Object.Verify()
}

const UnknownData = "unknown data"
