package config

type Config struct {
	AppId       string
	AppSecret   string
	DataUrl     map[string]string
	DataFields  map[string][]string
	SamplesRepo string
}
