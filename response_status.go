package gohttp

import "net/http"

// Http response status 2XX ---------------

// OK returns true if status equals StatusOK
func (r *response) OK() bool { return r.Status() == http.StatusOK }

// Created returns true if status equals StatusCreated
func (r *response) Created() bool { return r.Status() == http.StatusCreated }

// Accepted returns true if status equals StatusAccepted
func (r *response) Accepted() bool { return r.Status() == http.StatusAccepted }

// NonAuthoritativeInfo returns true if status equals StatusNonAuthoritativeInfo
func (r *response) NonAuthoritativeInfo() bool { return r.Status() == http.StatusNonAuthoritativeInfo }

// NoContent returns true if status equals StatusNoContent
func (r *response) NoContent() bool { return r.Status() == http.StatusNoContent }

// ResetContent returns true if status equals StatusResetContent
func (r *response) ResetContent() bool { return r.Status() == http.StatusResetContent }

// PartialContent returns true if status equals StatusPartialContent
func (r *response) PartialContent() bool { return r.Status() == http.StatusPartialContent }

// MultiStatus returns true if status equals StatusMultiStatus
func (r *response) MultiStatus() bool { return r.Status() == http.StatusMultiStatus }

// AlreadyReported returns true if status equals StatusAlreadyReported
func (r *response) AlreadyReported() bool { return r.Status() == http.StatusAlreadyReported }

// IMUsed returns true if status equals StatusIMUsed
func (r *response) IMUsed() bool { return r.Status() == http.StatusIMUsed }

// Http response status 4XX ---------------

// BadRequest will return true if status equals StatusBadRequest
func (r *response) BadRequest() bool { return r.Status() == http.StatusBadRequest }

// Unauthorized will return true if status equals StatusUnauthorized
func (r *response) Unauthorized() bool { return r.Status() == http.StatusUnauthorized }

// PaymentRequired will return true if status equals StatusPaymentRequired
func (r *response) PaymentRequired() bool { return r.Status() == http.StatusPaymentRequired }

// Forbidden will return true if status equals StatusForbidden
func (r *response) Forbidden() bool { return r.Status() == http.StatusForbidden }

// NotFound will return true if status equals StatusNotFound
func (r *response) NotFound() bool { return r.Status() == http.StatusNotFound }

// MethodNotAllowed will return true if status equals StatusMethodNotAllowed
func (r *response) MethodNotAllowed() bool { return r.Status() == http.StatusMethodNotAllowed }

// NotAcceptable will return true if status equals StatusNotAcceptable
func (r *response) NotAcceptable() bool { return r.Status() == http.StatusNotAcceptable }

// ProxyAuthRequired will return true if status equals StatusProxyAuthRequired
func (r *response) ProxyAuthRequired() bool { return r.Status() == http.StatusProxyAuthRequired }

// RequestTimeout will return true if status equals StatusRequestTimeout
func (r *response) RequestTimeout() bool { return r.Status() == http.StatusRequestTimeout }

// Conflict will return true if status equals StatusConflict
func (r *response) Conflict() bool { return r.Status() == http.StatusConflict }

// Gone will return true if status equals StatusGone
func (r *response) Gone() bool { return r.Status() == http.StatusGone }

// LengthRequired will return true if status equals StatusLengthRequired
func (r *response) LengthRequired() bool { return r.Status() == http.StatusLengthRequired }

// PreconditionFailed will return true if status equals StatusPreconditionFailed
func (r *response) PreconditionFailed() bool { return r.Status() == http.StatusPreconditionFailed }

// RequestEntityTooLarge will return true if status equals StatusRequestEntityTooLarge
func (r *response) RequestEntityTooLarge() bool {
	return r.Status() == http.StatusRequestEntityTooLarge
}

// RequestURITooLong will return true if status equals StatusRequestURITooLong
func (r *response) RequestURITooLong() bool { return r.Status() == http.StatusRequestURITooLong }

// UnsupportedMediaType will return true if status equals StatusUnsupportedMediaType
func (r *response) UnsupportedMediaType() bool { return r.Status() == http.StatusUnsupportedMediaType }

// RequestedRangeNotSatisfiable will return true if status equals StatusRequestedRangeNotSatisfiable
func (r *response) RequestedRangeNotSatisfiable() bool {
	return r.Status() == http.StatusRequestedRangeNotSatisfiable
}

// ExpectationFailed will return true if status equals StatusExpectationFailed
func (r *response) ExpectationFailed() bool { return r.Status() == http.StatusExpectationFailed }

// Teapot will return true if status equals StatusTeapot
func (r *response) Teapot() bool { return r.Status() == http.StatusTeapot }

// MisdirectedRequest will return true if status equals StatusMisdirectedRequest
func (r *response) MisdirectedRequest() bool { return r.Status() == http.StatusMisdirectedRequest }

// UnprocessableEntity will return true if status equals StatusUnprocessableEntity
func (r *response) UnprocessableEntity() bool { return r.Status() == http.StatusUnprocessableEntity }

// Locked will return true if status equals StatusLocked
func (r *response) Locked() bool { return r.Status() == http.StatusLocked }

// FailedDependency will return true if status equals StatusFailedDependency
func (r *response) FailedDependency() bool { return r.Status() == http.StatusFailedDependency }

// TooEarly will return true if status equals StatusTooEarly
func (r *response) TooEarly() bool { return r.Status() == http.StatusTooEarly }

// UpgradeRequired will return true if status equals StatusUpgradeRequired
func (r *response) UpgradeRequired() bool { return r.Status() == http.StatusUpgradeRequired }

// PreconditionRequired will return true if status equals StatusPreconditionRequired
func (r *response) PreconditionRequired() bool { return r.Status() == http.StatusPreconditionRequired }

// TooManyRequests will return true if status equals StatusTooManyRequests
func (r *response) TooManyRequests() bool { return r.Status() == http.StatusTooManyRequests }

// RequestHeaderFieldsTooLarge will return true if status equals StatusRequestHeaderFieldsTooLarge
func (r *response) RequestHeaderFieldsTooLarge() bool {
	return r.Status() == http.StatusRequestHeaderFieldsTooLarge
}

// UnavailableForLegalReasons will return true if status equals StatusUnavailableForLegalReasons
func (r *response) UnavailableForLegalReasons() bool {
	return r.Status() == http.StatusUnavailableForLegalReasons
}

// Http response status 5XX  ---------------

// InternalServerError returns status equals StatusInternalServerError
func (r *response) InternalServerError() bool { return r.Status() == http.StatusInternalServerError }

// NotImplemented returns status equals StatusNotImplemented
func (r *response) NotImplemented() bool { return r.Status() == http.StatusNotImplemented }

// BadGateway returns status equals StatusBadGateway
func (r *response) BadGateway() bool { return r.Status() == http.StatusBadGateway }

// ServiceUnavailable returns status equals StatusServiceUnavailable
func (r *response) ServiceUnavailable() bool { return r.Status() == http.StatusServiceUnavailable }

// GatewayTimeout returns status equals StatusGatewayTimeout
func (r *response) GatewayTimeout() bool { return r.Status() == http.StatusGatewayTimeout }

// HTTPVersionNotSupported returns status equals StatusHTTPVersionNotSupported
func (r *response) HTTPVersionNotSupported() bool {
	return r.Status() == http.StatusHTTPVersionNotSupported
}

// VariantAlsoNegotiates returns status equals StatusVariantAlsoNegotiates
func (r *response) VariantAlsoNegotiates() bool {
	return r.Status() == http.StatusVariantAlsoNegotiates
}

// InsufficientStorage returns status equals StatusInsufficientStorage
func (r *response) InsufficientStorage() bool { return r.Status() == http.StatusInsufficientStorage }

// LoopDetected returns status equals StatusLoopDetected
func (r *response) LoopDetected() bool { return r.Status() == http.StatusLoopDetected }

// NotExtended returns status equals StatusNotExtended
func (r *response) NotExtended() bool { return r.Status() == http.StatusNotExtended }

// NetworkAuthenticationRequired returns status equals StatusNetworkAuthenticationRequired
func (r *response) NetworkAuthenticationRequired() bool {
	return r.Status() == http.StatusNetworkAuthenticationRequired
}
