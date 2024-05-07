package protocol

type RequestDefinition struct {
	Route          string
	Parameters     []RequestParameter
	Dto            Verifier
	DtoArray       Verifier
	FileParts      []MultipartFileDefinition
	ValueParts     []MultipartValueDefinition
	Handler        Action
	MaxLimit       int
	Method         HTTPMethod
	AnyPermissions []string
	FreeRoute      bool // Free route require neither authentication nor authorization

	// Specific for Swagger
	Summary             string
	Description         string
	OperationId         string // Tools use this UNIQUE id to idenitfy the action example: GetUserById,GetUsers, ....
	Deprecated          bool   // Set true if this definition is deprecated
	ResponseDefinitions []ResponseDefinition
}

type HTTPMethod string
