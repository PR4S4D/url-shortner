package utils

import "github.com/google/uuid"

func GenerateUniqueId(length int) string {
	return uuid.New().String()[:length]
}
