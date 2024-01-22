package server

import "time"

// Config configures the server.
type Config struct {
	Address      string        `default:":5000"`
	Debug        bool          `default:"true"`
	ReadTimeout  time.Duration `default:"1m"`
	WriteTimeout time.Duration `default:"1m"`
	Version      string
	CORS         CORSConfig
}

// CORSConfig configures CORS.
type CORSConfig struct {
	AllowCredentials bool `default:"true"`
	Headers          []string
	Methods          []string
	Origins          []string
}
