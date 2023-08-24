package iotl

import (
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func Render() {
	fmt.Println(uuid.New().String())
	fmt.Println(yaml.ScalarNode)

}

func a() {}
