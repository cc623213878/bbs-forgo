package db

import (
	"bbs-forgo/log"
	"bbs-forgo/utils/autoconfig"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	CasbinConn *casbin.Enforcer
	DB         *gorm.DB
)

func Conn() error {
	var dialector gorm.Dialector
	var err error
	dataConf := autoconfig.Config.Base.Database
	createDatabase(&dataConf)
	if dataConf.DBType == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dataConf.Username, dataConf.Password, dataConf.Host, dataConf.Port, dataConf.DBName)
		dialector = mysql.New(mysql.Config{DSN: dsn, DefaultStringSize: 256})
	} else if dataConf.DBType == "postgresql" {
		dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable TimeZone=Asia/Shanghai",
			dataConf.Host, dataConf.Port, dataConf.Username, dataConf.DBName, dataConf.Password)
		dialector = postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
	}
	DB, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: "t_",
	},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)                                     // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)                                    // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Duration(8*3600) * time.Second) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	//casbin 初始化
	adapter, _ := gormadapter.NewAdapterByDB(DB) // Your driver and data source.
	m, err := model.NewModelFromString(`
    [request_definition]
    r = sub, obj, act
    
    [policy_definition]
    p = sub, obj, act
    
    [role_definition]
    g = _, _
    g2 = _, _
    
    [policy_effect]
    e = some(where (p.eft == allow))
    
    [matchers]
    m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act
    `)
	if err != nil {
		log.GetLogger().Error(err.Error())
		return err
	}
	CasbinConn, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.GetLogger().Error(err.Error())
		return err
	}
	// Load the policy from DB.
	err = CasbinConn.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

func GetConn() *gorm.DB {
	return DB
}

func GetCasbinConn() *casbin.Enforcer {
	return CasbinConn
}

func createDatabase(dataConf *autoconfig.DatabaseData) {
	var dialector gorm.Dialector
	if dataConf.DBType == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/?charset=utf8mb4&parseTime=True&loc=Local",
			dataConf.Username, dataConf.Password, dataConf.Host, dataConf.Port)
		dialector = mysql.New(mysql.Config{DSN: dsn, DefaultStringSize: 256})
	} else if dataConf.DBType == "postgresql" {
		dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v sslmode=disable TimeZone=Asia/Shanghai",
			dataConf.Host, dataConf.Port, dataConf.Username, dataConf.Password)
		dialector = postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
	}
	dbcreate, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
	dbconn, err := dbcreate.DB()
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
	defer dbconn.Close()
	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;",
		dataConf.DBName,
	)
	err = dbcreate.Exec(createSQL).Error
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
}
