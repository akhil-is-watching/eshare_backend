// package main

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Generate random bytes for shared secret
// 	sharedSecret := make([]byte, 32)
// 	if _, err := io.ReadFull(rand.Reader, sharedSecret); err != nil {
// 		fmt.Printf("Error generating shared secret: %v\n", err)
// 		return
// 	}

// 	// Convert the secret key to base64 and print it
// 	encodedKey := base64.StdEncoding.EncodeToString(sharedSecret)
// 	fmt.Printf("Secret Key (save this for decryption): %s\n", encodedKey)

// 	// Example usage
// 	err := encryptPDFFile("normal.pdf", "encrypted.pdf", sharedSecret)
// 	if err != nil {
// 		fmt.Printf("Error encrypting file: %v\n", err)
// 		return
// 	}
// 	fmt.Println("File encrypted successfully!")

// 	err = decryptPDFFileWithEncodedKey("encrypted.pdf", "decrypted.pdf", encodedKey)
// 	if err != nil {
// 		fmt.Printf("Error decrypting file: %v\n", err)
// 		return
// 	}
// 	fmt.Println("File decrypted successfully!")
// }

// func encryptPDFFile(inputPath, outputPath string, sharedSecret []byte) error {
// 	// Read input file
// 	plaintext, err := os.ReadFile(inputPath)
// 	if err != nil {
// 		return fmt.Errorf("failed to read input file: %v", err)
// 	}

// 	// Create AES cipher using the shared secret
// 	block, err := aes.NewCipher(sharedSecret)
// 	if err != nil {
// 		return fmt.Errorf("failed to create cipher: %v", err)
// 	}

// 	// Create GCM mode
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return fmt.Errorf("failed to create GCM: %v", err)
// 	}

// 	// Generate nonce
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
// 		return fmt.Errorf("failed to generate nonce: %v", err)
// 	}

// 	// Encrypt the file
// 	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

// 	// Write to output file
// 	err = os.WriteFile(outputPath, ciphertext, 0644)
// 	if err != nil {
// 		return fmt.Errorf("failed to write output file: %v", err)
// 	}

// 	return nil
// }

// func decryptPDFFile(inputPath, outputPath string, sharedSecret []byte) error {
// 	// Read encrypted file
// 	ciphertext, err := os.ReadFile(inputPath)
// 	if err != nil {
// 		return fmt.Errorf("failed to read encrypted file: %v", err)
// 	}

// 	// Create AES cipher using the shared secret
// 	block, err := aes.NewCipher(sharedSecret)
// 	if err != nil {
// 		return fmt.Errorf("failed to create cipher: %v", err)
// 	}

// 	// Create GCM mode
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return fmt.Errorf("failed to create GCM: %v", err)
// 	}

// 	// Extract nonce from ciphertext
// 	nonceSize := gcm.NonceSize()
// 	if len(ciphertext) < nonceSize {
// 		return fmt.Errorf("ciphertext too short")
// 	}
// 	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

// 	// Decrypt the file
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to decrypt file: %v", err)
// 	}

// 	// Write decrypted data to output file
// 	err = os.WriteFile(outputPath, plaintext, 0644)
// 	if err != nil {
// 		return fmt.Errorf("failed to write decrypted file: %v", err)
// 	}

// 	return nil
// }

// func decryptPDFFileWithEncodedKey(inputPath, outputPath, encodedKey string) error {
// 	// Decode the base64 key
// 	sharedSecret, err := base64.StdEncoding.DecodeString(encodedKey)
// 	if err != nil {
// 		return fmt.Errorf("failed to decode key: %v", err)
// 	}

// 	return decryptPDFFile(inputPath, outputPath, sharedSecret)
// }
