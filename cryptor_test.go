package goutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	var randNum map[int64]int = make(map[int64]int)
	var duplicate int = 0
	max := 1 << 32
	fmt.Println("max:", max)

	for i := 0; i < 1000; i++ {
		t.Run("Token", func(t *testing.T) {
			gen := SecureRandom(int64(max))
			if _, ok := randNum[gen]; ok {
				t.Errorf("Number %v already exists, after: %v", gen, i)
				duplicate++
			} else {
				randNum[gen] = i
			}
		})
		if duplicate >= 3 {
			break
		}
	}
}
func TestToken(t *testing.T) {
	var token map[string]int = make(map[string]int)
	for i := 0; i < 1000; i++ {
		t.Run("Token", func(t *testing.T) {
			gen := GenerateToken(32)
			if _, ok := token[gen]; ok {
				t.Errorf("Token %s already exists, after: %v", gen, i)
			} else {
				token[gen] = i
			}

		})
	}
}

func TestCaeserEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   int
		want  string
	}{
		{
			name:  "Caeser Encode low key",
			input: "Hello World",
			key:   1,
			want:  "Ifmmp Xpsme",
		},
		{
			name:  "Caeser Encode high key",
			input: "Hello World!",
			key:   300,
			want:  "Vszzc Kcfzr!",
		},
		{
			name:  "Caeser Encode negativ key",
			input: "Hello World",
			key:   -1,
			want:  "Gdkkn Vnqkc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CaeserEncode(test.input, test.key)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("CaeserEncode(%s, %v) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}

func TestCaeserDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   int
		want  string
	}{
		{
			name:  "Caeser Decode low key",
			input: "Ifmmp Xpsme",
			key:   1,
			want:  "Hello World",
		},
		{
			name:  "Caeser Decode high key",
			input: "Vszzc Kcfzr!",
			key:   14,
			want:  "Hello World!",
		},
		{
			name:  "Caeser Decode negativ key",
			input: "Gdkkn Vnqkc",
			key:   -1,
			want:  "Hello World",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CaeserDecode(test.input, test.key)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("CaeserDecode(%s, %d) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}

func TestVigenereEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   string
		want  string
	}{
		{
			name:  "Vigenere Encode key same length as input",
			input: "Hello World",
			key:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:  "Hfnos Cvzun",
		},
		{
			name:  "Vigenere Encode key shorter than input",
			input: "Hello World!",
			key:   "key",
			want:  "Rijvs Gspvh!",
		},
		{
			name:  "Vigenere Encode key upper case",
			input: "Hello World",
			key:   "Oopsie",
			want:  "Vsadw Kcgdl",
		},
		{
			name:  "Vigenere Encode short key",
			input: "Hello World",
			key:   "A",
			want:  "Hello World",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := VigenereEncode(test.input, test.key)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("CaeserEncode(%s, %s) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}

func TestVigenereDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		key   string
		want  string
	}{
		{
			name:  "Caeser Decode low key",
			input: "Hfnos Cvzun",
			key:   "ABCDEFGHIJK",
			want:  "Hello World",
		},
		{
			name:  "Caeser Decode high key",
			input: "Rijvs Gspvh!",
			key:   "key",
			want:  "Hello World!",
		},
		{
			name:  "Caeser Decode negativ key",
			input: "Vsadw Kcgdl",
			key:   "Oopsie",
			want:  "Hello World",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := VigenereDecode(test.input, test.key)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("CaeserDecode(%s, %s) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}
func TestRot13Encode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Rot13 Simple",
			input: "Hello World",
			want:  "Uryyb Jbeyq",
		},
		{
			name:  "Rot13 Long",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:  "NOPQRSTUVWXYZABCDEFGHIJKLM",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Rot13Encode(test.input)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("Rot13(%s) = %s, want %s", test.input, got, test.want)
			}
		})
	}
}

func TestRot13Decode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Rot13 Simple",
			input: "Uryyb Jbeyq",
			want:  "Hello World",
		},
		{
			name:  "Rot13 Long",
			input: "NOPQRSTUVWXYZABCDEFGHIJKLM",
			want:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Rot13Decode(test.input)
			if strings.Compare(got, test.want) != 0 {
				t.Errorf("Rot13(%s) = %s, want %s", test.input, got, test.want)
			}
		})
	}
}

func TestRot47Encode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Rot47 Simple",
			input: "Hello World",
			want:  "w6==@ (@C=5",
		},
		{
			name:  "Rot47 All",
			input: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
			want:  "pqrstuvwxyz{|}~!\"#$%&'()*+23456789:;<=>?@ABCDEFGHIJK_`abcdefghPQRSTUVWXYZ[\\]^ijklmno,-./01LMNO",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Rot47Encode(test.input)
			if got != test.want {
				t.Errorf("Rot47(%s) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestRot47Decode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Rot47 Simple",
			input: "w6==@ (@C=5",
			want:  "Hello World",
		},
		{
			name:  "Rot47 All",
			input: "pqrstuvwxyz{|}~!\"#$%&'()*+23456789:;<=>?@ABCDEFGHIJK_`abcdefghPQRSTUVWXYZ[\\]^ijklmno,-./01LMNO",
			want:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Rot47Decode(test.input)
			if got != test.want {
				t.Errorf("Rot47(%s) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
