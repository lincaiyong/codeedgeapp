package config

import "context"

type ChatFn func(ctx context.Context, model, q string, f func(string)) (string, error)

type Config struct {
	AppId       string
	AppSecret   string
	DataUrl     map[string]string
	DataFields  map[string][]string
	SamplesRepo string
	ChatFn      ChatFn
}
