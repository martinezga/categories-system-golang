package shared

import (
	"errors"
	"fmt"

	"github.com/teris-io/shortid"
)

func GenerateUniqueString(identifier string, finalLength int) (string, error) {
	minLength := 16
	if finalLength < minLength {
		return "", errors.New("final length is too short")
	}
	// Configure the shortid generator
	sid, err := shortid.New(1, shortid.DefaultABC, uint64(finalLength*5))
	if err != nil {
		return "", err
	}
	// Generate unique strings
	string1, err := sid.Generate()
	if err != nil {
		return "", err
	}
	string2, err := sid.Generate()
	if err != nil {
		return "", err
	}
	// Create the final string
	result := fmt.Sprintf("%s_%s%s", identifier, string1, string2)

	return result[:finalLength], nil
}
