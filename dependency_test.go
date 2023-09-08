package apiutils

import (
	"errors"
	"testing"
)

func Test_ValidateDependencies(t *testing.T) {
	tests := []struct {
		name     string
		input    Dependencies
		expected error
	}{
		{
			name:     "Test_Nil_Dependency_Map",
			input:    nil,
			expected: errors.New("dependencies map is nil"),
		},
		{
			name:     "Test_Empty_Dependencies",
			input:    Dependencies{},
			expected: nil,
		},
		{
			name: "Test_Valid_Dependencies",
			input: Dependencies{
				"db":  "some_db_connection",
				"log": "some_logger",
			},
			expected: nil,
		},
		{
			name: "Test_Nil_Dependency",
			input: Dependencies{
				"db": nil,
			},
			expected: errors.New("db is nil"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDependencies(tt.input); err != nil {
				if tt.expected == nil {
					t.Errorf("got error %v, want nil", err)
				} else if err.Error() != tt.expected.Error() {
					t.Errorf("got error %v, want %v", err, tt.expected)
				}
			} else if tt.expected != nil {
				t.Errorf("got nil error, want %v", tt.expected)
			}
		})
	}
}
