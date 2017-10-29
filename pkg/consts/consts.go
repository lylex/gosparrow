package consts

// Service related constants
const (
	ServiceName        = "Sparrow"
	ServiceDefaultPort = "8080"
)

// Environment related constants
const (
	EnvVariablePort = "PORT"
)

// Log levels
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelFatal = "fatal"
	LogLevelPanic = "panic"
)

// Error message related constants
const (
	ErrLogLevelParsePattern           = "parse logger level %s failed"
	ErrInternalServerErrorMsg         = "internal server error"
	ErrHandlerBadRequestBodyMsg       = "bad or unsupported request body"
	ErrHandlerUnsupportedMediaTypeMsg = "unsupported media type"
)

// Request header related constants
const (
	ReqHeaderCorrelationID        = "X-Correlation-Id"
	ReqHeaderContentType          = "Content-Type"
	ReqHeaderAuthorization        = "Authorization"
	ReqHeaderAccept               = "Accept"
	ReqHeaderValueApplicationJSON = "application/json"
)

// Response header related constants
const (
	RespHeaderContentType          = "Content-Type"
	RespHeaderContentLength        = "Content-Length"
	RespHeaderValueApplicationJSON = "application/json; charset=UTF-8"
)
