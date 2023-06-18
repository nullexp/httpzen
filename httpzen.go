package httpzen

type (
	// The Api interface defines the methods required to run an HTTP server and handle incoming requests.
	Api interface {
		Run(ip string, port uint, mode string) error
		AppendModule(mod Module)
		AppendMiddleware(string, Action) // TODO: add middleware base on order. study it!
		GetRoute(url, method string) *RequestDefinition
		SetCors(cors []string)
		// TODO: add crash report handler with mode( dev or etc)
	}

	// The Module interface defines the methods required to handle a group of related request handlers.
	Module interface {
		GetRequestHandlers() []*RequestDefinition
		GetBaseURL() string
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
	}

	// HTTPMethod is a string representing an HTTP method (GET, POST, DELETE, etc.).
	HTTPMethod string
)
