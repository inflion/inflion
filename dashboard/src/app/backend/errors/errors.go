package errors

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
)

// NewInvalid return a statusError
// which is an error intended for consumption by a REST API server; it can also be
// reconstructed by clients from a REST response. Public to allow easy type switches.
func NewInvalid(reason string) *StatusError {
	return &StatusError{
		ErrStatus: Status{
			Status:  StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  StatusReasonInvalid,
			Message: reason,
		},
	}
}

// NewUnauthorized returns an error indicating the client is not authorized to perform the requested
// action.
func NewUnauthorized(reason string) *errors.StatusError {
	return errors.NewUnauthorized(reason)
}

// NewBadRequest creates an error that indicates that the request is invalid and can not be processed.
func NewBadRequest(reason string) *errors.StatusError {
	return errors.NewBadRequest(reason)
}

// NewTokenExpired return a statusError
// which is an error intended for consumption by a REST API server; it can also be
// reconstructed by clients from a REST response. Public to allow easy type switches.
func NewTokenExpired(reason string) *StatusError {
	return &StatusError{
		ErrStatus: Status{
			Status:  StatusFailure,
			Code:    http.StatusUnauthorized,
			Reason:  StatusReasonExpired,
			Message: reason,
		},
	}
}

// IsTokenExpired determines if the err is an error which errStatus' message is MsgTokenExpiredError
func IsTokenExpired(err error) bool {
	statusErr, ok := err.(*errors.StatusError)
	if !ok {
		return false
	}

	return statusErr.ErrStatus.Message == MsgTokenExpiredError
}

// IsAlreadyExists determines if the err is an error which indicates that a specified resource already exists.
func IsAlreadyExists(err error) bool {
	return errors.IsAlreadyExists(err)
}

// IsUnauthorized determines if err is an error which indicates that the request is unauthorized and
// requires authentication by the user.
func IsUnauthorized(err error) bool {
	return errors.IsUnauthorized(err)
}
