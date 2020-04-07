package util

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

type UniqueID interface {
	Generate() string
}

type UUIDUniqueID struct{}

func (UUIDUniqueID) Generate() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
