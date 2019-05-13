package main

import (
	"reflect"
	"testing"
)

func TestReadConfigFakeFile(t *testing.T) {

	config, err := readConfig("config-fake.ini")

	if reflect.TypeOf(err) != reflect.TypeOf(&ConfigurationError{}) {
		t.Errorf("Wrong error type, ConfigurationError wanted, got %v", reflect.TypeOf(err))
	}

	if config != nil {
		t.Errorf("Configuration must be nil if error occur, got %v", config)
	}
}

func TestReadConfigRightFile(t *testing.T) {

	config, err := readConfig("config.ini")

	if err != nil {
		t.Errorf("Wrong error type, ConfigurationError wanted, got %v", reflect.TypeOf(err))
	}

	if reflect.TypeOf(config) != reflect.TypeOf(&Config{}) {
		t.Errorf("Wrong error type, ConfigurationError wanted, got %v", reflect.TypeOf(err))
	}
}
