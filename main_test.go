package configuration

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type MyConfig struct {
	SQLConfig
}

func loadConfig() MyConfig {
	c := MyConfig{}
	c.SQLConfig = LoadSQLConfig()
	return c
}

func Test_Config(t *testing.T) {
	Convey("Test config loading", t, func(c C) {
		dbName := os.Getenv("DB_NAME")
		log.Println("DB_NAME =", dbName)
		config := loadConfig()
		So(config.SQLConfig.Name, ShouldEqual, dbName)
		So(config.Name, ShouldEqual, dbName)
		So(dbName, ShouldEqual, "alpaca")
	})
}
