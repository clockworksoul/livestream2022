package main

import "errors"

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func Delete(key string) error {
	delete(store, key)

	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Put(key string, value string) error {
	store[key] = value

	return nil
}
