package httpzen

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/nullexp/httpzen/protocol"
	"github.com/nullexp/httpzen/protocol/model"
	openapi "github.com/nullexp/httpzen/protocol/model/openapi"
)

type ginApp struct {
	ginDomainHandlers []protocol.Module
	preHandlers       map[string][]protocol.Action
	cors              []string
	gin               *gin.Engine

	// for openapi
	ApiInfo      *openapi.Info
	Servers      []openapi.Server
	Contact      *openapi.Contact
	ExternalDocs *openapi.ExternalDocs
	openApiRoute string
	jsonOpenApi  string
}

func NewGinZen() protocol.Api {
	return ginApp{}
}

func (ginApp) Run(ip string, port uint, mode string) error {
	panic("not implemented")
}

func (ginApp) AppendDuplexModule(mod protocol.DuplexModule) {
	panic("not implemented")
}

func (ginApp) AppendModule(mod protocol.Module) {
	panic("not implemented")
}

func (ginApp) AppendPreHandlers(string, protocol.Action) {
	panic("not implemented")
}

func (ginApp) GetRoute(url, method string) *protocol.RequestDefinition {
	panic("not implemented")
}

func (ginApp) AppendAuthorizer(baseURL string, authorizer protocol.Authorizer) {
	panic("not implemented")
}

func (ginApp) AppendAuthenticator(baseURL string, authorizer protocol.Authenticator) {
	panic("not implemented")
}

func (ginApp) SetCors(cors []string) {
	panic("not implemented")
}

func (ginApp) SetLogHandler(protocol.LogHandler) {
	panic("not implemented")
}

func (ginApp) SetLogPolicy(model.LogPolicy) {
	panic("not implemented")
}

func (ginApp) TestHandle(*httptest.ResponseRecorder, *http.Request) error {
	panic("not implemented")
}

// OpenAPI
func (ginApp) SetExternalDocs(openapi.ExternalDocs) {
	panic("not implemented")
}

func (ginApp) SetInfo(openapi.Info) {
	panic("not implemented")
}

func (ginApp) SetContact(openapi.Contact) {
	panic("not implemented")
}

func (ginApp) SetServers([]openapi.Server) {
	panic("not implemented")
}

func (ginApp) EnableOpenApi(route string) error {
	panic("not implemented")
}

func (ginApp) SetErrors([]string) {
	panic("not implemented")
}
