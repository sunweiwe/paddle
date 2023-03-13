package cmd

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/sirupsen/logrus"
	"github.com/sunweiwe/paddle/core/config"
	"github.com/sunweiwe/paddle/core/http/api/v1/user"
	"github.com/sunweiwe/paddle/lib/orm"
	authUser "github.com/sunweiwe/paddle/pkg/auth/user"
	"github.com/sunweiwe/paddle/pkg/parameter"
	"github.com/sunweiwe/paddle/pkg/parameter/manager"
	ormhook "github.com/sunweiwe/paddle/pkg/util/ormHook"
	"gorm.io/gorm"

	userController "github.com/sunweiwe/paddle/core/controller/user"
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

type RegisterI interface {
	RegisterRoutes(engine *gin.Engine)
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
	SetupLog(flags)

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

	redisClient := redis.NewClient(&redis.Options{
		Network:  coreConfig.RedisConfig.Protocol,
		Addr:     coreConfig.RedisConfig.Address,
		Password: coreConfig.RedisConfig.Password,
		DB:       int(coreConfig.RedisConfig.DB),
	})

	store, err := redisstore.NewRedisStore(context.Background(), redisClient)
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: int(coreConfig.SessionConfig.MaxAge),
	})

	gob.Register(&authUser.DefaultUser{})
	r := gin.New()
	middlewares := []gin.HandlerFunc{}
	r.Use(middlewares...)
	gin.ForceConsoleColor()

	registerRouter(r, mysqlDB, store)

	log.Printf("Server started")
	log.Print(r.Run(fmt.Sprintf(":%d", coreConfig.ServerConfig.Port)))
}

func registerRouter(r *gin.Engine, db *gorm.DB, store *redisstore.RedisStore) {
	m := manager.CreateManager(db)
	parameter := &parameter.Parameter{
		Manager: m,
	}

	var (
		userCtl = userController.NewController(parameter)

		userAPI = user.NewAPI(userCtl, store)
	)

	v1 := []RegisterI{
		userAPI,
	}

	for _, register := range v1 {
		register.RegisterRoutes(r)
	}
}

func SetupLog(flags *Flags) {
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
