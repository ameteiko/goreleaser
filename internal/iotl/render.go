package iotl

import (
	"fmt"

	"github.com/google/uuid"
)

func Render() {
	fmt.Println(uuid.New().String())

}
