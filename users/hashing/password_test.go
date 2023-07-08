package hashing_test

import (
	"testing"

	"github.com/Dominik48N/url-shorter/users/hashing"
)

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		password string
		hashed   string
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"world", "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7"},
		{"123456", "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"},
	}

	for _, tc := range testCases {
		result := hashing.HashPassword(tc.password)
		if result != tc.hashed {
			t.Errorf("HashPassword(%s) = %s; want %s", tc.password, result, tc.hashed)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		password       string
		hashedPassword string
		expected       bool
	}{
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", true},
		{"world", "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7", true},
		{"123456", "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", true},
		{"hello", "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7", false},
	}

	for _, tc := range testCases {
		result := hashing.CheckPassword(tc.password, tc.hashedPassword)
		if result != tc.expected {
			t.Errorf("CheckPassword(%s, %s) = %t; want %t", tc.password, tc.hashedPassword, result, tc.expected)
		}
	}
}
