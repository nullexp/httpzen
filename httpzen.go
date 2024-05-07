package httpzen

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"time"
)

type (
	// The Api interface defines the methods required to run an HTTP server and handle incoming requests.
	Api interface {
		Run(ip string, port uint, mode string) error
		AppendModule(mod Module)
		AppendMiddleware(string, Action) // TODO: add middleware base on order. study it!
		GetRoute(url, method string) *RequestDefinition
		SetCors(cors []string)
		// TODO: add crash report handler with mode( dev or etc)

		// For testing
		TestHandle(*httptest.ResponseRecorder, *http.Request) error

		// OpenAPI
		SetExternalDocs(ExternalDocs)
		SetInfo(Info)
		SetContact(Contact)
		SetServers([]Server)
		EnableOpenApi(route string) error
		SetErrors([]string)
	}

	// The Module interface defines the methods required to handle a group of related request handlers.
	Module interface {
		GetRequestHandlers() []*RequestDefinition
		GetBaseURL() string
		GetTag() Tag
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
		Get(key string) (interface{}, bool)
		MustGet(key string) interface{}
		GetDto() (interface{}, bool)
		MustGetDto() interface{}
		Negotiate(stausCode int, err error, Dto interface{})
		RangeFile(status int, err error, file SeekerFile)
		WriteFile(status int, err error, file File)
		ReturnMultipartMixed(status int, err error, out ...Multipart)

		SetFile(key string, f FileHeader)
		SetFiles(key string, f []FileHeader)
		GetDTO() (interface{}, bool)
		MustGetDTO() interface{}
		GetFile(partName string) (FileHeader, bool)
		MustGetFile(partName string) FileHeader
		GetFiles(partName string) ([]FileHeader, bool)
		MustGetFiles(partName string) []FileHeader
		GetPagination() (Pagination, bool)
		GetCursorPagination() (CursorPagination, bool)
		GetCaller() (Caller, bool)
		MustGetCaller() Caller
		GetSort() []Sort
		GetQuery() []Query
		GetDefaultQuery() (string, bool)
		IsAndQuery() bool
	}

	// The Verifier interface defines the method required to validate input data.
	Verifier interface {
		Verify() error
	}

	// An Action is a function that handles an HTTP request.
	Action func(req Request)

	// A RequestDefinition contains information about an HTTP route, including the URL pattern,
	// expected data types, and HTTP method.
	RequestDefinition struct {
		Route   string
		Dto     Verifier
		Handler Action
		Method  HTTPMethod

		// For open api
		Summary             string
		Description         string
		OperationID         string // Tools use this UNIQUE id to idenitfy the action example: GetUserByID,GetUsers, ....
		Deprecated          bool   // Set true if this definition is deprecated
		ResponseDefinitions []ResponseDefinition
	}

	// HTTPMethod is a string representing an HTTP method (GET, POST, DELETE, etc.).
	HTTPMethod string

	Info struct {
		Version     string
		Description string
		Title       string
	}

	Contact struct {
		Name  string
		Email string
		URL   string
	}

	Server struct {
		URL         string
		Description string
	}

	Tag struct {
		Name         string
		Description  string
		ExternalDocs *ExternalDocs
	}

	// The name is plural by open api definition
	ExternalDocs struct {
		Description string
		URL         string
	}

	NameGetter interface {
		GetName() string
	}

	DescriptionGetter interface {
		GetDescription() string
	}

	MultipartDefinition interface {
		IsOptional() bool
		GetPartName() string
		IsSingle() bool
		GetObject() interface{}
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

	FileMultiaprt interface {
		File
		Multipart
	}

	FileHeader interface {
		GetHeader() textproto.MIMEHeader
		GetSize() int64
		GetFilename() string
		OpenFile() (File, error)
	}

	File interface {
		io.Reader
		io.Closer
		GetFilename() string
		GetMimeType() string
		GetLastModifiedDate() time.Time
	}

	SeekerFile interface {
		io.Seeker
		File
	}
	ResponseDefinition struct {
		Description string // For example: success operation
		Status      int
		Dto         any // Might be basic dto,
	}
)
