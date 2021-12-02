package main

import (
	"fmt"
	"gb/Golang2_Lesson7/task1/structurecopy"
	"log"
)

type Config struct {
	A int64
	B string
	C *Config
	D struct{ D1 int }
	E struct {
		E1 int
		E2 int
	}
}

type ConfigMap map[string]interface{}

func main() {
	config := Config{
		A: -8,
		B: "123",
	}

	mapConfig := ConfigMap{
		"A": interface{}(int64(5)),
		"B": "Hello",
		"C": &Config{},
		"D": struct{ D1 int }{2},
		"E": map[string]interface{}{"E1": 888, "E2": 999},
	}

	err := structurecopy.StructureSet(&config, mapConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результирующая структура: %+v\n", config)
}
