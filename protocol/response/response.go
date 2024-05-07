package response

const (
	// InvalidAuthInfo explaining an invalid username and password.
	InvalidAuthInfo = "InvalidAuthInfo"
	// BannedeIntegrityFailed explaining a client is banned because of token integrity.
	BannedeIntegrityFailed = "BannedeIntegrityFailed"
	// SessionExpired explaining that client token is expired.
	SessionExpired = "SessionExpired"
	// UnknownFormat explaining a client send unknown format data
	UnknownFormat = "UnknownFormat"
	// ValidationError explaining a client send invalid data
	ValidationError = "ValidationError"
	// MissingAuth explaining a client did not send auth.
	MissingAuth = "MissingAuth"
	// Banned explaining a client is banned.
	Banned = "Banned"
	// ServerError indicate sth bad happen at server
	ServerError = "ServerError"
	// MissingDefaults indicate missing important header fields like send time and message identifier
	MissingDefaults = "MissingDefaults"
	// NotFound indicate missing data
	NotFound = "NotFound"
	// MalformMultipart indicate client send malformed multipart form data.
	MalformMultipart = "MalformMultipart"
	AccessDenied     = "AccessDenied"
)

func GetErrors() []string {
	return []string{
		InvalidAuthInfo,
		BannedeIntegrityFailed,
		SessionExpired,
		UnknownFormat,
		ValidationError,
		MissingAuth,
		Banned,
		ServerError,
		MissingDefaults,
		NotFound,
		MalformMultipart,
		AccessDenied,
	}
}
