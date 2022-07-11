package goutil

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

// Ceaser cipher
// The key is the shift amount.
// The key must be between 0 and 255.
func Ceaser(b []byte, key int) []byte {
	for i, v := range b {
		b[i] = v + byte(key)
	}
	return b
}

// Ceaser brute force
func CeaserBrute(b []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = Ceaser(b, i)
	}
	return b
}

// Vigenere cipher
// The key is the string to use as the key.
// The key must be between 0 and 255.
func Vigenere(b []byte, key []byte) []byte {
	for i, v := range b {
		b[i] = v + key[i%len(key)]
	}
	return b
}

// Vinegere brute force
func VigenereBrute(b []byte, key []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = Vigenere(b, key)
	}
	return b
}

// Xor cipher
// The key is the byte to use as the key.
func Xor(b []byte, key byte) []byte {
	for i, v := range b {
		b[i] = v ^ key
	}
	return b
}

// Xor brute force
func XorBrute(b []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = Xor(b, byte(i))
	}
	return b
}

// Rot13 cipher
func Rot13(b []byte) []byte {
	for i, v := range b {
		if v >= 'a' && v <= 'z' {
			b[i] = v + 13
			if b[i] > 'z' {
				b[i] = b[i] - 26
			}
		} else if v >= 'A' && v <= 'Z' {
			b[i] = v + 13
			if b[i] > 'Z' {
				b[i] = b[i] - 26
			}
		}
	}
	return b
}

// Rot13 brute force
func Rot13Brute(b []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = Rot13(b)
	}
	return b
}

// Rot47 cipher
func Rot47(b []byte) []byte {
	for i, v := range b {
		if v >= '!' && v <= '~' {
			b[i] = v + 47
			if b[i] > '~' {
				b[i] = b[i] - 94
			}
		}
	}
	return b
}

// Rot47 brute force
func Rot47Brute(b []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = Rot47(b)
	}
	return b
}

// SecureRandom is a cryptographically secure random number generator.
// The number generated is between 0 and max.
func SecureRandom(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

// Generate a random token of the given length.
// The token is a base64 encoded string.
// The token is URL safe.
// The token is base64 padded.
func GenerateToken(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
=======
	"regexp"
	"strings"
)

// SecureRandom is a cryptographically secure random number generator.
// The number generated is between 0 and max.
func SecureRandom(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

// Generate a random token of the given length.
// The token will have following characteristics:
// URL-safe,
// Base64 encoded,
// padded.
func GenerateToken(length int) string {
	if length <= 0 {
		length = 16
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Caeser encoder:
// This function encodes the input string using the key.
// Only letters of the basic alphabet will be edited.
// The input is the byte slice and the key.
// The output is the encoded byte slice.
func CaeserEncode(input string, key int) string {
	var output []rune = make([]rune, len(input))
	offset := byte(key % 26)

	for i, v := range input {
		if v == ' ' {
			output[i] = ' '
			continue
		} else if !regexp.MustCompile(`[a-zA-Z]`).MatchString(string(v)) {
			output[i] = v
			continue
		}

		output[i] = rune(byte(v) + offset)
		if v >= 'A' && v <= 'Z' {
			if output[i] > 'Z' {
				output[i] = output[i] - 26
			}
		} else if v >= 'a' && v <= 'z' {
			if output[i] > 'z' {
				output[i] = output[i] - 26
			}
		}
	}
	return string(output)
}

// Caeser decoder:
// This function encodes the input string using the key.
// Only letters of the basic alphabet will be edited.
// The input is the byte slice and the key.
// The output is the encoded byte slice.
func CaeserDecode(input string, key int) string {
	var output []rune = make([]rune, len(input))
	offset := byte(key % 26)

	for i, v := range input {
		if v == ' ' {
			output[i] = ' '
			continue
		} else if !regexp.MustCompile(`[a-zA-Z]`).MatchString(string(v)) {
			output[i] = v
			continue
		}

		output[i] = rune(byte(v) - offset)
		if v >= 'A' && v <= 'Z' {
			if output[i] < 'A' {
				output[i] = output[i] + 26
			}
		} else if v >= 'a' && v <= 'z' {
			if output[i] < 'a' {
				output[i] = output[i] + 26
			}
		}
	}
	return string(output)
}

// Vigenere encoder:
// This function encodes the input string using the key.
// Only letters of the basic alphabet will be edited.
// The input is the byte slice and the key.
// The output is the encoded byte slice.
func VigenereEncode(input string, key string) string {
	var output []rune = make([]rune, len(input))
	var offset byte

	for i, v := range input {
		key = strings.ToLower(key)
		offset = key[i%len(key)] - 'a'

		caeser := CaeserEncode(string(v), int(offset))
		output[i] = rune(caeser[0])
	}
	return string(output)
}

// Vigenere decoder:
// This function decodes the input string using the key.
// Only letters of the basic alphabet will be edited.
// The input is the ciphered byte slice and the key.
// The output is the deciphered byte slice.
func VigenereDecode(input string, key string) string {
	var output []rune = make([]rune, len(input))
	var offset byte

	for i, v := range input {
		key = strings.ToLower(key)
		offset = key[i%len(key)] - 'a'

		caeser := CaeserDecode(string(v), int(offset))
		output[i] = rune(caeser[0])
	}
	return string(output)
}

// Rot13 cipher encoder:
func Rot13Encode(input string) string {
	return CaeserEncode(input, 13)
}

func Rot13Decode(input string) string {
	return CaeserDecode(input, 13)
}

// Rot47 cipher encoder
func Rot47Encode(input string) string {
	output := make([]rune, len(input))
	offset := 47

	for i, v := range input {
		asciiCode := int(v)
		if 33 <= asciiCode && asciiCode <= 126 {
			output[i] = v + rune(offset)
			if output[i] > 126 {
				output[i] = output[i] - 94
			}
		} else {
			output[i] = v
		}
	}
	return string(output)
}

// Rot47 cipher decoder
func Rot47Decode(input string) string {
	output := make([]rune, len(input))
	offset := 47

	for i, v := range input {
		asciiCode := int(v)
		if 33 <= asciiCode && asciiCode <= 126 {
			output[i] = v - rune(offset)
			if output[i] < 33 {
				output[i] = output[i] + 94
			}
		} else {
			output[i] = v
		}
	}
	return string(output)
}
