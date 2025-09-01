package Generators

import (
	"github.com/google/uuid"
)

func GenerateUuidSalt() string {
	return uuid.New().String()
}
