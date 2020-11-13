package errors

// HandleHTTPError is used to handle HTTP Errors more accurately based on the localized consts
import (
	"k8s.io/apimachinery/pkg/api/errors"
	"log"
	"net/http"
)

// NonCriticalErrors is an array of error statuses, that are non-critical. That means, that this error can be
// silenced and displayed to the user as a warning on the frontend side.
var NonCriticalErrors = []int32{http.StatusForbidden, http.StatusUnauthorized}

func HandleHTTPError(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}
	if err.Error() == MsgTokenExpiredError || err.Error() == MsgLoginUnauthorizedError || err.Error() == MsgEncryptionKeyChanged {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

// HandleError handles single error, that occurred during API GET call. If it is not critical, then it will be
// returned as a part of error array. Otherwise, it will be returned as a second value. Usage of this functions
// allows to distinguish critical errors from non-critical ones. It is needed to handle them in a different way.
func HandleError(err error) ([]error, error) {
	nonCriticalErrors := make([]error, 0)
	return AppendError(err, nonCriticalErrors)
}

// AppendError handles single error, that occurred during API GET call. If it is not critical, then it will be
// returned as a part of error array. Otherwise, it will be returned as a second value. Usage of this functions
// allows to distinguish critical errors from non-critical ones. It is needed to handle them in a different way.
func AppendError(err error, nonCriticalErrors []error) ([]error, error) {
	if err != nil {
		if isErrorCritical(err) {
			return nonCriticalErrors, LocalizeError(err)
		}
		log.Printf("Non-critical error occurred during resource retrieval: %s", err)
		nonCriticalErrors = appendMissing(nonCriticalErrors, LocalizeError(err))
	}
	return nonCriticalErrors, nil
}

func isErrorCritical(err error) bool {
	status, ok := err.(*errors.StatusError)
	if !ok {
		// Assume, that error is critical if it cannot be mapped.
		return true
	}
	return !contains(NonCriticalErrors, status.ErrStatus.Code)
}

func appendMissing(slice []error, toAppend ...error) []error {
	m := make(map[string]bool, 0)

	for _, s := range slice {
		m[s.Error()] = true
	}

	for _, a := range toAppend {
		_, ok := m[a.Error()]
		if !ok {
			slice = append(slice, a)
			m[a.Error()] = true
		}
	}

	return slice
}

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
