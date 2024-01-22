package postgresql

// Config configures the connection to the PostgreSQL database.
type Config struct {
	Host         string
	Port         int `default:"5432"`
	DatabaseName string
	Username     string
	Password     string
	LogMode      bool   `default:"true"`
	Level        string `default:"debug"`
}
