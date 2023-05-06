package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type WebConfig struct {
	CMD string `yaml:"CMD"`
	URL string `yaml:"URL"`
}

type StoreConfig struct {
	TPC struct {
		Status bool     `yaml:"Status"`
		CMD    []string `yaml:"CMD"`
	} `yaml:"2PC"`
	RAFT struct {
		Status bool     `yaml:"Status"`
		CMD    []string `yaml:"CMD"`
	} `yaml:"RAFT"`
}

type BalancerConfig struct {
	Status bool     `yaml:"Status"`
	CMD    string   `yaml:"CMD"`
	URL    string   `yaml:"URL"`
	Web    []string `yaml:"Web"`
}

type StaticConfig struct {
	Course   string `yaml:"Course"`
	Student  string `yaml:"Student"`
	RootPath string `yaml:"RootPath"`
	Build    string `yaml:"Build"`
	Clean    string `yaml:"Clean"`
	Debug    bool   `yaml:"Debug"`
	HttpDebug    bool   `yaml:"HttpDebug"`
}

type t struct {
	WebConfig      WebConfig      `yaml:"Web"`
	StoreConfig    StoreConfig    `yaml:"Store"`
	BalancerConfig BalancerConfig `yaml:"Balancer"`
	StaticConfig   StaticConfig   `yaml:"Static"`
}

var config t

func Init() {

	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Panicf("[Init] 打开配置文件失败: %v", err.Error())
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Panicf("[Init] 格式化配置文件失败: %v", err.Error())
	}

}

func Check() {
	if config.BalancerConfig.Status {
		if config.BalancerConfig.CMD == "" {
			log.Panicf("[Check] 配置文件错误: %v", "config.BalancerConfig.CMD")
		}
		if len(config.BalancerConfig.Web) < 2 {
			log.Panicf("[Check] 配置文件错误: %v", "config.BalancerConfig.Web")
		}
		if config.BalancerConfig.URL == "" {
			log.Panicf("[Check] 配置文件错误: %v", "config.BalancerConfig.URL")
		}
	}

	if config.StoreConfig.TPC.Status {
		if config.StoreConfig.RAFT.Status {
			log.Panicf("[Check] 配置文件错误: %v", "config.StoreConfig 2PC和RAFT只能开启一个")
		}
		if len(config.StoreConfig.TPC.CMD) < 4 {
			log.Panicf("[Check] 配置文件错误: %v", "config.StoreConfig.TPC.CMD")
		}
	}
	if config.StoreConfig.RAFT.Status {
		if config.StoreConfig.TPC.Status {
			log.Panicf("[Check] 配置文件错误: %v", "config.StoreConfig 2PC和RAFT只能开启一个")
		}
		if len(config.StoreConfig.RAFT.CMD) < 4 {
			log.Panicf("[Check] 配置文件错误: %v", "config.StoreConfig.RAFT.CMD")
		}
	}
	if config.StaticConfig.Course == "" {
		log.Panicf("[Check] 配置文件错误: %v", "config.StaticConfig.Course")
	}
	if config.StaticConfig.Student == "" {
		log.Panicf("[Check] 配置文件错误: %v", "config.StaticConfig.Student")
	}
}
func GetWeb() WebConfig {
	return config.WebConfig
}

func GetStore() StoreConfig {
	return config.StoreConfig
}

func GetBalancer() BalancerConfig {
	return config.BalancerConfig
}

func GetStatic() StaticConfig {
	return config.StaticConfig
}
