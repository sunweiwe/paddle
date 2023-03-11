package cmd

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
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
