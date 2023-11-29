package lal

import (
	"errors"

	"github.com/ameteiko/goreleaser/internal/iotl"
)

func goes() {
	_ = errors.New("test")

	iotl.Feature()
}
