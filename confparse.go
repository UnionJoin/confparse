package confparse

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory", path)
	}
	return nil
}

func ParseConfig(config interface{}) error {
	value := reflect.ValueOf(config).Elem()
	typ := reflect.TypeOf(config).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			err := ParseConfig(field.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		tag := fieldType.Tag.Get("required")
		if tag == "true" {
			if !field.IsValid() {
				return fmt.Errorf("missing config field %s", fieldType.Name)
			}
			if field.Kind() == reflect.Bool {
				return nil
			} else if field.IsZero() {
				return fmt.Errorf("config field %s requires a value", fieldType.Name)
			}
		}
	}

	return nil
}

func LoadConfig(configPath string, config interface{}) error {
	err := ValidateConfigPath(configPath)
	if err != nil {
		return err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(content), config)
	if err != nil {
		return err
	}

	err = ParseConfig(config)
	if err != nil {
		return err
	}

	return nil
}
