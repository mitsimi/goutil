package goutil

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
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
