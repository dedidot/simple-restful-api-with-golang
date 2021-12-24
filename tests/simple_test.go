package tests

import (
	"fmt"
	"restful-api/simple"
	"testing"
)

func TestServiceProvider(t *testing.T) {
	simpleServie, err := simple.InitializedService(false)
	fmt.Println(simpleServie.Action)
	fmt.Println(err)
}