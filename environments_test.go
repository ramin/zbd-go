package zebedee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {
	tests := []struct {
		environment Environment
		expected    string
	}{
		{Production, "https://api.zebedee.io/v0"},
		{Sandbox, "https://sandbox-api.zebedee.io/v0"},
	}

	for _, test := range tests {
		actual := endpoints[test.environment]
		if actual != test.expected {
			assert.Equal(t, actual, test.expected)
		}
	}
}

func TestEnvironmentString(t *testing.T) {
	tests := []struct {
		environment Environment
		expected    string
	}{
		{Production, "production"},
		{Sandbox, "sandbox"},
		{Public, "public"},
		{Environment(69429), "unknown"},
	}

	for _, test := range tests {
		actual := test.environment.String()
		assert.Equal(t, test.expected, actual)
	}
}
