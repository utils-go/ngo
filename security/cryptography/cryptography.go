// Package cryptography provides .NET System.Security.Cryptography-like hashing and encryption utilities.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.security.cryptography?view=netframework-4.7.2
package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
)

// ---- Hash Functions ----

// MD5Hash computes the MD5 hash of the input data.
func MD5Hash(data []byte) []byte {
	h := md5.Sum(data)
	return h[:]
}

// MD5HashString computes the MD5 hash as a hex string.
func MD5HashString(data string) string {
	h := md5.Sum([]byte(data))
	return hex.EncodeToString(h[:])
}

// SHA1Hash computes the SHA1 hash of the input data.
func SHA1Hash(data []byte) []byte {
	h := sha1.Sum(data)
	return h[:]
}

// SHA1HashString computes the SHA1 hash as a hex string.
func SHA1HashString(data string) string {
	h := sha1.Sum([]byte(data))
	return hex.EncodeToString(h[:])
}

// SHA256Hash computes the SHA256 hash of the input data.
func SHA256Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}

// SHA256HashString computes the SHA256 hash as a hex string.
func SHA256HashString(data string) string {
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}

// SHA256HashStream computes SHA256 from a reader.
func SHA256HashStream(reader io.Reader) ([]byte, error) {
	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// ---- HMAC ----

// HMACSHA256 computes an HMAC-SHA256 hash.
func HMACSHA256(data, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

// HMACSHA256String computes an HMAC-SHA256 hash as a hex string.
func HMACSHA256String(data string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// HMACSHA1 computes an HMAC-SHA1 hash.
func HMACSHA1(data, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

// ---- AES ----

// AESEncrypt encrypts data using AES with the specified key.
// Uses GCM mode (authenticated encryption).
func AESEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Prepend nonce to ciphertext
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AESDecrypt decrypts data using AES with the specified key.
func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}

// AESEncryptString encrypts a string and returns a hex-encoded string.
func AESEncryptString(plaintext string, key string) (string, error) {
	plainBytes := []byte(plaintext)
	keyBytes := padKey([]byte(key), 32) // AES-256 by default

	ciphertext, err := AESEncrypt(plainBytes, keyBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

// AESDecryptString decrypts a hex-encoded ciphertext string.
func AESDecryptString(ciphertextHex string, key string) (string, error) {
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	keyBytes := padKey([]byte(key), 32)
	plaintext, err := AESDecrypt(ciphertext, keyBytes)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// ---- Utility ----

// GenerateRandomBytes generates cryptographically secure random bytes.
func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return bytes, err
}

// GenerateKey generates a random key of the specified length suitable for AES.
func GenerateKey(bits int) ([]byte, error) {
	if bits%8 != 0 {
		return nil, fmt.Errorf("key length must be a multiple of 8")
	}
	return GenerateRandomBytes(bits / 8)
}

// padKey pads or truncates a key to the specified length.
func padKey(key []byte, length int) []byte {
	if len(key) >= length {
		return key[:length]
	}
	padded := make([]byte, length)
	copy(padded, key)
	return padded
}

// NewHash creates a new hash.Hash from the specified algorithm name.
func NewHash(algorithm string) (hash.Hash, error) {
	switch algorithm {
	case "MD5":
		return md5.New(), nil
	case "SHA1":
		return sha1.New(), nil
	case "SHA256":
		return sha256.New(), nil
	default:
		return nil, fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
}
