package core

import (
	"github.com/Unknwon/goconfig"
	"log"
)

// type Conf struct {
// 	Web struct {
// 		HttpPort int `gcfg:"http-port"`
// 		Timezone int
// 	}
// 	Db struct {
// 		Driver string
// 		Source string
// 	}
// }

// var Config Conf
var Config *goconfig.ConfigFile

func init() {
	initConfig()
}

func initConfig() {
	/* gcfg
	if &Config != nil {
		err := gcfg.ReadFileInto(&Config, "conf.gcfg")
		if err != nil {
			log.Fatalf("Failed to parse gcfg data: %s", err)
			panic(err)
		}
	}
	*/
	if &Config != nil {
		var err error
		Config, err = goconfig.LoadConfigFile("config.ini")
		if err != nil {
			log.Fatalf("Failed to parse config data: %s", err)
			panic(err)
		}
	}
}
