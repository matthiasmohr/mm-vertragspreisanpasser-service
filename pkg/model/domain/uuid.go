package domain

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

// UUID represents an UUID according to RFC 4122.
type UUID string

// Value returns the value of the underlying type.
// It also makes UUID implement driver.Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}

	val, err := uuid.Parse(string(u))
	if err != nil {
		return nil, fmt.Errorf("uuid value: parse uuid: %w", err)
	}

	return val, nil
}

// Scan scans the source value from the database into the underlying type.
// It also makes UUID implement sql.Scanner interface.
func (u *UUID) Scan(value interface{}) error {
	val, err := uuid.Parse(value.(string))
	if err != nil {
		return fmt.Errorf("scan uuid: %w", err)
	}

	*u = UUID(val.String())

	return nil
}

// ParseUUID decodes s into a UUID or returns an error.
func ParseUUID(s string) (UUID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", fmt.Errorf("parse uuid: %w", err)
	}

	return UUID(s), nil
}

// String returns the string value of the UUID.
func (u UUID) String() string {
	return string(u)
}

// New generates UUID.
func (u *UUID) New() {
	*u = UUID(uuid.New().String())
}

// NewUUID creates a new UUID.
func NewUUID() (UUID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to genrate a new random uuid %w", err)
	}

	return UUID(uuid.String()), nil
}
