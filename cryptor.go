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
}
