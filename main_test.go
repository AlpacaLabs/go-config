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
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")

		log.Println("DB_NAME =", dbName)

		config := loadConfig()

		So(config.SQLConfig.Name, ShouldEqual, dbName)
		So(config.Name, ShouldEqual, dbName)

		So(config.SQLConfig.User, ShouldEqual, dbUser)
		So(config.User, ShouldEqual, dbUser)

		So(config.SQLConfig.Pass, ShouldEqual, dbPass)
		So(config.Pass, ShouldEqual, dbPass)

		So(config.SQLConfig.Host, ShouldEqual, dbHost)
		So(config.Host, ShouldEqual, dbHost)

		So(dbName, ShouldEqual, "alpaca")
		So(dbUser, ShouldEqual, "alpaca")
		So(dbPass, ShouldEqual, "alpaca")
		So(dbHost, ShouldEqual, "alpaca")
	})
}
