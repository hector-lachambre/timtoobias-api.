package main

import (
	"log"
	"reflect"
	"testing"
)

func TestReadConfigWrongFile(t *testing.T) {

	config, err := ReadConfig("test/wrong_config.ini")

	log.Println(reflect.TypeOf(err))

	log.Println(reflect.TypeOf(&ConfigurationError{}))

	if reflect.TypeOf(err) != reflect.TypeOf(&ConfigurationError{}) {
		t.Errorf("Wrong error type, ConfigurationError wanted, got %v", reflect.TypeOf(err))
	}

	if config != nil {
		t.Errorf("Configuration must be nil if error occur, got %v", config)
	}
}
