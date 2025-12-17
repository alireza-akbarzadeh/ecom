// Package env contains the environment variables for the ecom app.
package env

import "os"

func GetEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
