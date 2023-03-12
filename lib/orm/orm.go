package orm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/prometheus"
)

type MySQL struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password,omitempty"`
	Database         string `json:"database"`
	PrometheusEnable bool   `json:"prometheusEnabled"`
}

func NewMySQL(db *MySQL) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username,
		db.Password, db.Host, db.Port, db.Database)

	sqlDB, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	orm, err := gorm.Open(mysql.New(
		mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
	})

	if db.PrometheusEnable {
		if err := orm.Use(prometheus.New(prometheus.Config{
			DBName: "mysql",
			MetricsCollector: []prometheus.MetricsCollector{
				&MySQLMetricsCollector{},
			},
		})); err != nil {
			return nil, err
		}
	}

	return orm, err
}
