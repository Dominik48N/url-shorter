package main

import "testing"

func TestIsValidURL(t *testing.T) {
	validURL := "AbcDefG"
	if !isValidURL(validURL) {
		t.Errorf("Expected %s to be a valid URL", validURL)
	}

	invalidURL := "Abc DefG"
	if isValidURL(invalidURL) {
		t.Errorf("Expected %s to be an invalid URL", invalidURL)
	}

	invalidLengthURL := "Ab"
	if isValidURL(invalidLengthURL) {
		t.Errorf("Expected %s to be an invalid URL due to length", invalidLengthURL)
	}

	invalidCharsURL := "AbcDefG123"
	if isValidURL(invalidCharsURL) {
		t.Errorf("Expected %s to be an invalid URL due to invalid characters", invalidCharsURL)
	}
}
