package httpzen

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

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

func NewGinZen() Api {
	return ginApp{}
}

func (ginApp) Run(ip string, port uint, mode string) error {
	panic("not implemented")
}
func (ginApp) AppendModule(mod Module) {
	panic("not implemented")
}
func (ginApp) AppendMiddleware(string, Action) {
	panic("not implemented")
}
func (ginApp) GetRoute(url, method string) *RequestDefinition {
	panic("not implemented")
}
func (ginApp) SetCors(cors []string) {
	panic("not implemented")
}
func (ginApp) TestHandle(*httptest.ResponseRecorder, *http.Request) error {
	panic("not implemented")
}
func (ginApp) SetExternalDocs(ExternalDocs) {
	panic("not implemented")
}
func (ginApp) SetInfo(Info) {
	panic("not implemented")
}
func (ginApp) SetContact(Contact) {
	panic("not implemented")
}
func (ginApp) SetServers([]Server) {
	panic("not implemented")
}
func (ginApp) EnableOpenApi(route string) error {
	panic("not implemented")
}
func (ginApp) SetErrors([]string) {
	panic("not implemented")
}
