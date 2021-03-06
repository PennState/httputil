package httpstatus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClass(t *testing.T) {
	tests := []struct {
		Name  string
		Code  int
		Class Class
	}{
		{"Invalid - too low", 99, ""},
		{"Invalid - too high", 600, ""},
		{"Switching Protocols", 101, "1xx"},
		{"Created", 201, "2xx"},
		{"Moved Permanently", 301, "3xx"},
		{"Unauthorized", 401, "4xx"},
		{"Not Implemented", 501, "5xx"},
	}
	for idx := range tests {
		test := tests[idx]
		t.Run(test.Name, func(t *testing.T) {
			cls := Family(test.Code)
			assert.Equal(t, test.Class, cls)
		})
	}
}

func ExampleClass() {
	codes := []int{101, 201, 301, 401, 501}
	for _, code := range codes {
		if Family(code) == Class(ClassSuccessful) {
			fmt.Println("Successful:", code)
			continue
		}
		fmt.Println("Unsuccessful:", code)
	}
	// Output:
	// Unsuccessful: 101
	// Successful: 201
	// Unsuccessful: 301
	// Unsuccessful: 401
	// Unsuccessful: 501
}
