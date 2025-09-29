// http_task_test.go
package main

import "testing"

func TestMaskURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "regular URL",
			input:    "http://hehefouls.netHAHAHA",
			expected: "http://*******************",
		},
		{
			name:     "URL with Cyrillic",
			input:    "http://Привет.net",
			expected: "http://**********",
		},
		{
			name:     "mixed characters URL",
			input:    "http://googleпривет.net",
			expected: "http://****************",
		},
		{
			name:     "short URL",
			input:    "http://",
			expected: "http://",
		},
		{
			name:     "very short URL",
			input:    "http:",
			expected: "http:",
		},
		{
			name:     "empty URL",
			input:    "",
			expected: "",
		},
		{
			name:     "exactly 7 characters",
			input:    "http://",
			expected: "http://",
		},
		{
			name:     "8 characters URL",
			input:    "http://a",
			expected: "http://*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskURL(tt.input)
			if result != tt.expected {
				t.Errorf("MaskURL(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaskURL_EdgeCases(t *testing.T) {
	// Тест с символами, занимающими несколько байт
	testCases := []struct {
		input    string
		expected string
	}{
		{"http://🚀example.com", "http://************"},
		{"http://café.com", "http://********"},
		{"http://世界.com", "http://******"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := MaskURL(tc.input)
			if result != tc.expected {
				t.Errorf("MaskURL(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}
