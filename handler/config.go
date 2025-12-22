package handler

import "context"

type ChatFn func(ctx context.Context, model, q string, f func(string)) (string, error)
type ObjectFn func(ctx context.Context, key string) ([]byte, error)

type Config struct {
	AppId      string
	AppSecret  string
	DataUrl    map[string]string
	DataFields map[string][]string
	SamplesUrl string
	ChatFn     ChatFn
	ObjectFn   ObjectFn
	ResetCache bool
}
