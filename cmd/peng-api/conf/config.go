package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Mongo string `yaml:"mongo"`
}

var gConfig Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		filename, _ := filepath.Abs("config.yml")

		// 로컬 개발환경에서는 실행 위치에 config.yml 파일이 없어서 config 원본 경로 사용
		if _, err := os.Stat(filename); err != nil {
			filename = "/Users/yongho/Works/Projects/depthon-2019/cmd/peng-api/conf/config.yml"
		}

		fmt.Println("filename : " + filename)
		yamlFile, err := ioutil.ReadFile(filename)
		err = yaml.Unmarshal(yamlFile, &gConfig)
		if err != nil {
			panic(err)
		}

		fmt.Printf("load config. %v\n", gConfig)
	})
	return gConfig
}
