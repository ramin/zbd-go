//go:build e2e

package testing

import (
	"os"

	zebedee "github.com/zebedeeio/zbd-go"
)

func NewClient() *zebedee.Client {
	apiKey := os.Getenv("ZEBEDEE_API_KEY")
	return zebedee.New(apiKey).Sandbox()
}
