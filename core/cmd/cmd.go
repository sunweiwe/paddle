package cmd

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sunweiwe/paddle/core/config"
	"github.com/sunweiwe/paddle/lib/orm"
	ormhook "github.com/sunweiwe/paddle/pkg/util/ormHook"
)

// Flags defines agent CLI flags.

type Flags struct {
	ConfigFile          string
	RoleConfigFile      string
	ScopeRoleFile       string
	BuildJSONSchemaFile string
	BuildUISchemaFile   string
	Dev                 bool
	Environment         string
	LogLevel            string
}

func ParseFlags() *Flags {
	var flags Flags

	flag.StringVar(
		&flags.ConfigFile, "config", "", "configuration file path")

	flag.BoolVar(
		&flags.Dev, "dev", false, "if true, turn off the userMiddleware to skip login")

	flag.StringVar(
		&flags.Environment, "environment", "production", "environment string tag")

	flag.StringVar(
		&flags.LogLevel, "logLevel", "info", "the logLevel(panic/fatal/error/warn/info/debug/trace))")

	flag.Parse()
	return &flags
}

func Run(flags *Flags) {
	InitLog(flags)

	coreConfig, err := config.LoadConfig(flags.ConfigFile)
	if err != nil {
		panic(err)
	}

	_, err = json.MarshalIndent(coreConfig, "", " ")
	if err != nil {
		panic(err)
	}

	// init database
	mysqlDB, err := orm.NewMySQL(&orm.MySQL{
		Host:             coreConfig.DBConfig.Host,
		Port:             coreConfig.DBConfig.Port,
		Username:         coreConfig.DBConfig.Username,
		Password:         coreConfig.DBConfig.Password,
		Database:         coreConfig.DBConfig.Database,
		PrometheusEnable: coreConfig.DBConfig.PrometheusEnabled,
	})
	if err != nil {
		panic(err)
	}

	ormhook.RegisterCustomHooks(mysqlDB)

}

func InitLog(flags *Flags) {
	if flags.Environment == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(flags.LogLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
}
