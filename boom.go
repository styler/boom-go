package boom

import (
	"encoding/json"
	"net/http"
)

// APIError is an error
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message,omitempty"`
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	error := APIError{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
		Message:    message,
	}

	writeJSON(w, statusCode, error)
}

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	bs, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bs)
}

// BadRequest responds with a 400 Bad Request error.
func BadRequest(w http.ResponseWriter, message string) {
	writeError(w, http.StatusBadRequest, message)
}

// Unathorized responds with a 401 Unauthorized error.
func Unathorized(w http.ResponseWriter, message string) {
	writeError(w, http.StatusUnauthorized, message)
}

// Forbidden responds with a 403 Forbidden error.
func Forbidden(w http.ResponseWriter, message string) {
	writeError(w, http.StatusForbidden, message)
}

// NotFound responds with a 404 Not Found error.
func NotFound(w http.ResponseWriter, message string) {
	writeError(w, http.StatusNotFound, message)
}

// MethodNotAllowed responds with a 405 Method Not Allowed error.
func MethodNotAllowed(w http.ResponseWriter, message string) {
	writeError(w, http.StatusMethodNotAllowed, message)
}

// NotAcceptable responds with a 406 Not Acceptable error.
func NotAcceptable(w http.ResponseWriter, message string) {
	writeError(w, http.StatusNotAcceptable, message)
}

// ClientTimeout responds with a 408 Request Time-out error.
func ClientTimeout(w http.ResponseWriter, message string) {
	writeError(w, http.StatusRequestTimeout, message)
}

// Conflict responds with a 409 Conflict error.
func Conflict(w http.ResponseWriter, message string) {
	writeError(w, http.StatusConflict, message)
}

// ResourceGone responds with a 410 Gone error.
func ResourceGone(w http.ResponseWriter, message string) {
	writeError(w, http.StatusGone, message)
}

// EntityTooLarge responds with a 413 Request Entity Too Large error.
func EntityTooLarge(w http.ResponseWriter, message string) {
	writeError(w, http.StatusRequestEntityTooLarge, message)
}

// URITooLong responds with a 414 Request-URI Too Large error
func URITooLong(w http.ResponseWriter, message string) {
	writeError(w, http.StatusRequestURITooLong, message)
}

// Teapot responds with a 418 I'm a Teapot error.
func Teapot(w http.ResponseWriter, message string) {
	writeError(w, http.StatusTeapot, message)
}

// BadData responds with a 422 Unprocessable Entity error.
func BadData(w http.ResponseWriter, message string) {
	writeError(w, http.StatusUnprocessableEntity, message)
}

// TooManyRequests responds with a 429 Too Many Requests error.
func TooManyRequests(w http.ResponseWriter, message string) {
	writeError(w, http.StatusTooManyRequests, message)
}

// Illegal responds with a 451 Unavailable For Legal Reasons error.
func Illegal(w http.ResponseWriter, message string) {
	writeError(w, http.StatusUnavailableForLegalReasons, message)
}

// Internal responds with a 500 Internal Server Error error.
func Internal(w http.ResponseWriter, message string) {
	writeError(w, http.StatusInternalServerError, message)
}

// BadImplementation responds with a 500 Internal Server Error error.
func BadImplementation(w http.ResponseWriter, message string) {
	writeError(w, http.StatusInternalServerError, message)
}

// NotImplemented responds with a 501 Not Implemented error.
func NotImplemented(w http.ResponseWriter, message string) {
	writeError(w, http.StatusNotImplemented, message)
}

// BadGateway responds with a 502 Bad Gateway error.
func BadGateway(w http.ResponseWriter, message string) {
	writeError(w, http.StatusBadGateway, message)
}

// ServerUnavailable .eturns a 503 Service Unavailable error.
func ServerUnavailable(w http.ResponseWriter, message string) {
	writeError(w, http.StatusServiceUnavailable, message)
}

// GatewayTimeout responds with a 504 Gateway Time-out error.
func GatewayTimeout(w http.ResponseWriter, message string) {
	writeError(w, http.StatusGatewayTimeout, message)
}
