package data

import (
	"go-leaf/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewLeafAllocRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	ll := gl.Warn
	if c.Database.Debug {
		ll = gl.Info
	}
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: gl.Default.LogMode(ll),
	})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// 配置数据库
	idb, err := db.DB()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	idb.SetMaxIdleConns(int(c.Database.MaxIdle))
	idb.SetMaxOpenConns(int(c.Database.MaxOpen))

	return &Data{db: db}, cleanup, nil
}
