package model

// DuplexError is used to inform clients from their action errors.
type DuplexError struct {
	// Message is details of the error
	Message string `json:"message"`
	// Code is based on <<Response Rules>> , Refer to authn and authz docs.
	Code string `json:"code"`
}
