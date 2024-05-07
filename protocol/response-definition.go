package protocol

// Describe the output
type ResponseDefinition struct {
	Description string // For example: success operation
	Status      int
	Dto         any // Might be basic dto,
	IsFile      bool
}
