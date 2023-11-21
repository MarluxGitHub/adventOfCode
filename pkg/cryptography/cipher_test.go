package cryptography_test

import (
	"MarluxGitHub/adventOfCode/pkg/cryptography"
	"testing"
)

func TestCipher(t *testing.T) {
	tests := []struct {
		name  string
		input string
		shift int
		want  string
	}{
		{
			name:  "Test 1: Cipher with shift 1",
			input: "abc",
			shift: 1,
			want:  "bcd",
		},
		{
			name:  "Test 2: Cipher with shift 0",
			input: "abc",
			shift: 0,
			want:  "abc",
		},
		{
			name:  "Test 3: Cipher with negative shift",
			input: "abc",
			shift: -1,
			want:  "zab",
		},
		{
			name:  "Test 4: Cipher with shift greater than 26",
			input: "abc",
			shift: 27,
			want:  "bcd",
		},
		{
			name:  "Test 5: Cipher with hyphen",
			input: "a-b-c",
			shift: 1,
			want:  "b c d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cryptography.Cipher(tt.input, tt.shift); got != tt.want {
				t.Errorf("Cipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
