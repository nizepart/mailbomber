package app

import (
	"os"
	"strconv"
)

type Value struct {
	value string
}

func (v *Value) String() string {
	return v.value
}

func (v *Value) Int() int {
	value, _ := strconv.Atoi(v.value)
	return value
}

func GetValue(key string, fallback string) *Value {
	if value, ok := os.LookupEnv(key); ok {
		return &Value{value: value}
	}
	return &Value{value: fallback}
}
