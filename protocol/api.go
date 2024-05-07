package protocol

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/textproto"

	"github.com/nullexp/httpzen/misc"
	fileProtocol "github.com/nullexp/httpzen/protocol/file"
	"github.com/nullexp/httpzen/protocol/model"
	"github.com/nullexp/httpzen/protocol/model/openapi"
)

type (
	// The Api interface defines the methods required to run an HTTP server and handle incoming requests.
	Api interface {
		Run(ip string, port uint, mode string) error
		AppendDuplexModule(mod DuplexModule)
		AppendModule(mod Module)
		AppendPreHandlers(string, Action)
		GetRoute(url, method string) *RequestDefinition
		AppendAuthorizer(baseURL string, authorizer Authorizer)
		AppendAuthenticator(baseURL string, authorizer Authenticator)
		SetCors(cors []string)
		SetLogHandler(LogHandler)
		SetLogPolicy(model.LogPolicy)
		TestHandle(*httptest.ResponseRecorder, *http.Request) error
		// OpenAPI
		SetExternalDocs(openapi.ExternalDocs)
		SetInfo(openapi.Info)
		SetContact(openapi.Contact)
		SetServers([]openapi.Server)
		EnableOpenApi(route string) error
		SetErrors([]string)
	}

	LogHandler interface {
		Handle(model.HttpLog)
	}

	// The Module interface defines the methods required to handle a group of related request handlers.
	Module interface {
		GetRequestHandlers() []*RequestDefinition
		GetBaseURL() string
		GetTag() openapi.Tag
	}

	// The Request interface defines the methods available to an HTTP request handler.
	Request interface {
		// GetJson and other stuff
		SetServerError(msg string)
		SetForbidden()
		SetUnauthorized(msg string, code string)
		SetBadRequest(msg string, code string)
		SetNotFound(msg string, code string)
		ReturnStatus(int, error)
		Set(key string, value interface{})
		SetFile(key string, f FileHeader)
		SetFiles(key string, f []FileHeader)
		Get(key string) (interface{}, bool)
		MustGet(key string) interface{}
		GetDTO() (interface{}, bool)
		MustGetDTO() interface{}
		GetFile(partName string) (FileHeader, bool)
		MustGetFile(partName string) FileHeader
		GetFiles(partName string) ([]FileHeader, bool)
		MustGetFiles(partName string) []FileHeader
		Negotiate(stausCode int, err error, dto interface{})
		RangeFile(status int, err error, file fileProtocol.SeekerFile)
		WriteFile(status int, err error, file fileProtocol.File)
		ReturnMultipartMixed(status int, err error, out ...Multipart)
		GetPagination() (misc.Pagination, bool)
		GetCursorPagination() (misc.CursorPagination, bool)
		GetCaller() (misc.Caller, bool)
		MustGetCaller() misc.Caller
		GetSort() []misc.Sort
		GetQuery() []misc.Query
		GetDefaultQuery() (string, bool)
		IsAndQuery() bool
	}

	// The Verifier interface defines the method required to validate input data.
	Verifier interface {
		Verify() error
	}

	// An Action is a function that handles an HTTP request.
	Action func(req Request)

	Authorizer func(identity string, permission string) (bool, error)

	MultipartDefinition interface {
		IsOptional() bool
		GetPartName() string
		IsSingle() bool
		GetObject() interface{}
		GetSupportedTypes() []string
	}

	MultipartFileDefinition interface {
		MultipartDefinition
		Verify(FileHeader) error
	}

	MultipartValueDefinition interface {
		MultipartDefinition
		Verify() error
	}

	Multipart interface {
		io.ReadCloser
		GetPartName() string
		GetMimeType() string
	}

	FileMultipart interface {
		fileProtocol.File
		Multipart
	}

	FileHeader interface {
		GetHeader() textproto.MIMEHeader
		GetSize() int64
		GetFilename() string
		OpenFile() (fileProtocol.File, error)
	}
)

const (
	KeyRole  = "Role"
	KeyQuery = "Query"
	KeyAuth  = "Auth"
	KeyDTO   = "DTO"
	KeyFile  = "File"
	MaxLimit = "MaxLimit"
)
