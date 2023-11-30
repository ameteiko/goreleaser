package iotl

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ameteiko/errors"
	"gopkg.in/yaml.v3"
)

func Render() {
	fmt.Println(uuid.New().String())
	fmt.Println(yaml.ScalarNode)

	_ = errors.New("test")
}

func a() {}
