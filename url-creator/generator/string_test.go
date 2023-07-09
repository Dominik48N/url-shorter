package generator_test

import (
	"testing"

	"github.com/Dominik48N/url-shorter/url-creator/generator"
)

func TestGenerateRandomString(t *testing.T) {
	minLength := 5
	maxLength := 10
	randomString := generator.GenerateRandomString(minLength, maxLength)

	if len(randomString) < minLength || len(randomString) > maxLength {
		t.Errorf("Generated string length is not within the allowed range")
	}
}
