package utils

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReplacePlaceholders(template string, values map[string]string) string {
	for key, value := range values {
		template = strings.ReplaceAll(template, key, value)
	}
	return template
}
