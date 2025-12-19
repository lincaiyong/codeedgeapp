package handler

var config Config

type Config struct {
	AppId       string
	AppSecret   string
	DataUrl     map[string]string
	DataFields  map[string][]string
	SamplesRepo string
}

func Init(conf Config) {
	config = conf
}
