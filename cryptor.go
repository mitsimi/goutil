package goutil

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"regexp"
)

// Caeser cipher encode
// The Caeser cipher is a substitution cipher.
// The key is the shift amount.
// The input is the byte slice to cipher.
// The output is the ciphered byte slice.
// The output is the same length as the input.
// The output is in the same order as the input.
//
// Only letters of the basic alphabet will be ciphered.
func CaeserEncode(input []byte, key int) []byte {
	var output []byte = make([]byte, len(input))
	copy(output, input)
	offset := byte(key % 26)

	for i, v := range input {
		if regexp.MustCompile(`[a-zA-Z]`).MatchString(string(v)) {
			output[i] = v + offset
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
	}
	return output
}

// Caeser cipher decode
// The Caeser cipher is a substitution cipher.
// The key is the shift amount.
// The input is the ciphered byte slice.
// The output is the deciphered byte slice.
// The output is the same length as the input.
// The output is in the same order as the input.
func CaeserDecode(input []byte, key int) []byte {
	var output []byte = make([]byte, len(input))
	copy(output, input)
	offset := byte(key % 26)

	for i, v := range input {
		if regexp.MustCompile(`[a-zA-Z]`).MatchString(string(v)) {
			output[i] = v - offset
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
	}
	return output
}

// Caeser brute force
func CaeserBrute(b []byte, min, max int) (s string, key int) {
	for key := min; key <= max; key++ {
		s += fmt.Sprintf("%s\n", string(CaeserEncode(b, key)))
	}
	return
}

// Vigenere cipher
// The key is the string to use as the key.
// The input is the byte slice to cipher.
// The output is the ciphered byte slice.
// The output is the same length as the input.
// The output is in the same order as the input.
//
// Only letters of the basic alphabet will be ciphered.
func VigenereEncode(b []byte, key []byte) []byte {
	for i, v := range b {
		offset := key[i%len(key)]
		b[i] = CaeserEncode([]byte{v}, int(offset))[0]
	}
	return b
}

// Vigenere cipher
// The key is the string to use as the key.
// The input is the ciphered byte slice.
// The output is the deciphered byte slice.
// The output is the same length as the input.
// The output is in the same order as the input.
//
// Only letters of the basic alphabet will be deciphered.
func VigenereDecode(b []byte, key []byte) []byte {
	for i, v := range b {
		b[i] = v - key[i%len(key)]
	}
	return b
}

// Vinegere brute force
func VigenereBrute(b []byte, key []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		b = VigenereDecode(b, key)
	}
	return b
}

// Xor cipher
// The key is the byte to use as the key.
func XorEncode(b []byte, key byte) []byte {
	for i, v := range b {
		b[i] = v ^ key
	}
	return b
}

// Xor brute force
func XorBrute(b []byte, min, max int) []byte {
	for i := min; i <= max; i++ {
		//		b = Xor(b, byte(i))
	}
	return b
}

// Rot13 cipher
func Rot13Encode(b []byte) []byte {
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
		//		b = Rot13(b)
	}
	return b
}

// Rot47 cipher
func Rot47Encode(b []byte) []byte {
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
		//		b = Rot47(b)
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
}
