package httpzen

import "github.com/gin-gonic/gin"

type ginApp struct {
	ginDomainHandlers []Module
	preHandlers       map[string][]Action
	cors              []string
	gin               *gin.Engine

	// for openapi
	ApiInfo      *Info
	Servers      []Server
	Contact      *Contact
	ExternalDocs *ExternalDocs
	openApiRoute string
	jsonOpenApi  string
}
