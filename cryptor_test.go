package goutil

import (
	"bytes"
	"testing"
)

type CipherTest struct {
	name  string
	input []byte
	key   int
	want  []byte
}

func TestCaeserEncode(t *testing.T) {
	tests := []CipherTest{
		{
			name:  "Caeser Encode low key",
			input: []byte("Hello World"),
			key:   1,
			want:  []byte("Ifmmp Xpsme"),
		},
		{
			name:  "Caeser Encode high key",
			input: []byte("Hello World!"),
			key:   300,
			want:  []byte("Vszzc Kcfzr!"),
		},
		{
			name:  "Caeser Encode negativ key",
			input: []byte("Hello World"),
			key:   -1,
			want:  []byte("Gdkkn Vnqkc"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CaeserEncode(test.input, test.key)
			if !bytes.Equal(got, test.want) {
				t.Errorf("CaeserEncode(%s, %d) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}

func TestCaeserDecode(t *testing.T) {
	tests := []CipherTest{
		{
			name:  "Caeser Decode low key",
			input: []byte("Ifmmp Xpsme"),
			key:   1,
			want:  []byte("Hello World"),
		},
		{
			name:  "Caeser Decode high key",
			input: []byte("Vszzc Kcfzr!"),
			key:   14,
			want:  []byte("Hello World!"),
		},
		{
			name:  "Caeser Decode negativ key",
			input: []byte("Gdkkn Vnqkc"),
			key:   -1,
			want:  []byte("Hello World"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CaeserDecode(test.input, test.key)
			if !bytes.Equal(got, test.want) {
				t.Errorf("CaeserDecode(%s, %d) = %s, want %s", test.input, test.key, got, test.want)
			}
		})
	}
}

func TestRot13Encode(t *testing.T) {
	tests := []CipherTest{
		{
			name:  "Rot13",
			input: []byte("Hello World"),
			want:  []byte("Uryyb Jbeyq"),
		},
		{
			name:  "Rot13",
			input: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			want:  []byte("NOPQRSTUVWXYZABCDEFGHIJKLM"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Rot13Encode(test.input)
			if !bytes.Equal(got, test.want) {
				t.Errorf("Rot13(%s) = %s, want %s", test.input, got, test.want)
			}
		})
	}
}
