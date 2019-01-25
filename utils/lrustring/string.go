package lrustring

import "github.com/satori/go.uuid"

func Uuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
