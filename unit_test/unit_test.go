package unit_test

import (
	"errors"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// TestAdd tests that the Add function returns the correct result.
func TestAdd(t *testing.T) {
	ans := Add(2, 3)
	if ans != 5 {
		t.Errorf("Expected 5, got %d", ans)
	}
}

// TestDivideTableDriven tests the Divide function using a table-driven approach.
// It checks for correct division results and proper error handling for division by zero.
func TestDivideTableDriven(t *testing.T) {
	tests := []struct {
		name       string
		a, b, want int
		wantErr    error
	}{
		{"Divide 2 by 1", 2, 1, 2, nil},
		{"Divide 4 by 2", 4, 2, 2, nil},
		{"Divide 0 by 2", 0, 2, 0, nil},
		{"Divide 1 by 0", 1, 0, 0, errors.New("division by zero")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Divide(test.a, test.b)

			// cek apakah terjadi error
			if err != nil && test.wantErr == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err == nil && test.wantErr != nil {
				t.Errorf("Expected error %v, got nil", test.wantErr)
			}
			if err != nil && test.wantErr != nil && err.Error() != test.wantErr.Error() {
				t.Errorf("Expected error %v, got %v", test.wantErr, err)
			}

			// cek hasil jika tidak ada error
			if got != test.want {
				t.Errorf("Expected %d, got %d", test.want, got)
			}
		})
	}
}
