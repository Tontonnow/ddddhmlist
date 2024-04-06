package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func GenerateRandomString(length int) string {
	characters := "0123456789abcdef"
	return GenerateRandomStringWithCharset(length, characters)
}
func GenerateRandomNumber(length int) string {
	characters := "0123456789"
	return GenerateRandomStringWithCharset(length, characters)
}
func GenerateRandomStringWithCharset(length int, charset string) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomMac() string {
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", GenerateRandomString(2), GenerateRandomString(2), GenerateRandomString(2), GenerateRandomString(2), GenerateRandomString(2), GenerateRandomString(2))
}

func GenerateRandomAndroidId() string {
	return GenerateRandomString(16)
}

func GenerateRandomImei() string {
	return GenerateRandomNumber(15)
}

func GenerateRandomUdid() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", GenerateRandomString(8), GenerateRandomString(4), GenerateRandomString(4), GenerateRandomString(4), GenerateRandomString(12))
}
func GenerateUUID() string {
	return uuid.New().String()
}
