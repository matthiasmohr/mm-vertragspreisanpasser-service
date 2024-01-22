package usecase

const (
	ErrTypeDatabase = "database"
	ErrTypeDomain   = "domain"
)

// Error is a wrapper for possible usecase errors.
type Error struct {
	Message string
	Code    int
	Type    string `json:"-"`
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// Domain error codes.
const (
	ErrCodeDomainInternal = iota + 5001000
)

// Domain errors.
var (
	ErrDomainInternal = &Error{
		Message: "internal error",
		Code:    ErrCodeDomainInternal,
		Type:    ErrTypeDomain,
	}
)

// Database error codes.
const (
	ErrCodeDatabaseInternal = iota + 5002000
	ErrCodeDatabaseCustomerNotFound
)

// Database errors.
var (
	ErrDatabaseInternal = &Error{
		Message: "internal database error",
		Code:    ErrCodeDatabaseInternal,
		Type:    ErrTypeDatabase,
	}

	ErrDatabaseCustomerNotFound = &Error{
		Message: "case not found",
		Code:    ErrCodeDatabaseCustomerNotFound,
		Type:    ErrTypeDatabase,
	}
)
