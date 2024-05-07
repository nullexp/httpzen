package model

// RequestError is used to inform clients from their action errors.
type RequestError struct {
	// Message is details of the error
	Message string `json:"message" description:"details of the error, for developer consumption only"`
	// Code is based on <<Response Rules>> , Refer to authn and authz docs.
	Code string `json:"code" description:"error code to evaluate"`
}
