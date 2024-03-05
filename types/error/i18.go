package error

import (
	"context"
	"errors"
	"fmt"
)

const (
	I18nLanguageEn = "en"
	I18nLanguageCh = "ch"
)

var i i18n

type languageMapping map[StatusErrorKey]string

func init() {
	i.collections = make(map[string]languageMapping, 0)
}

type i18n struct {
	collections map[string]languageMapping
}

func GetI18nMessage(ctx context.Context, key StatusErrorKey) (string, error) {
	return GetI18nMessageF(ctx, key, nil)
}

func GetI18nMessageF(ctx context.Context, key StatusErrorKey, args interface{}) (string, error) {
	location := "en"
	// get collection
	collection, ok := i.collections[location]
	if !ok {
		return "", errors.New("location not registration")
	}

	// get message key
	msg, ok := collection[key]
	if !ok {
		return "", fmt.Errorf("message key not registration on location %s", location)
	}

	if args == nil {
		return msg, nil
	}

	return fmt.Sprintf(msg, args), nil
}

func InjectI18nCollection(location string, collections map[StatusErrorKey]string) {
	i.collections[location] = collections
}
