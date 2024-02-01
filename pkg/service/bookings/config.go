package bookings

// Config configures the required information for accessing bookings endpoints.
type Config struct {
	Host        string
	Credentials Credentials
}

// Credentials object holds data allowing the service to be authenticated by bookings adapter.
type Credentials struct {
	User     string
	Password string
}
