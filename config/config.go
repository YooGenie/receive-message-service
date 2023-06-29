package config

import (
	"os"

	"github.com/jinzhu/configor"
)

var Config = struct {
	ServiceName string
	Environment string
	AwsSqs      struct {
		QueueUrl string
	}
}{}

func InitConfig(cfg string) {
	configor.Load(&Config, cfg)
}

func ConfigureEnvironment(path string) {
	configor.Load(&Config, path+"/config.json")

	Config.AwsSqs.QueueUrl = os.Getenv("AWS_SQS_QUEUE_URL")
}
