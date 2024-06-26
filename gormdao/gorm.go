package gormdao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormDao struct {
	Db *gorm.DB

	Config *Config
}

func (this *GormDao) WithConfig(config *Config) *GormDao {
	this.Config = config
	return this
}

// now only support mysql TODO:support other database by config
func (this *GormDao) Open() *GormDao {
	var err error
	if !this.Config.HasDsn() {
		panic("you must set gorm dsn first")
	}
	this.Db, err = gorm.Open(mysql.Open(this.Config.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}

	if d, err := this.Db.DB(); err != nil {
		panic(err.Error())
	} else {
		d.SetMaxIdleConns(this.Config.MaxOpenConns)
		d.SetMaxOpenConns(this.Config.MaxOpenConns)
		d.SetConnMaxIdleTime(this.Config.ConnMaxIdleTime)
	}
	return this
}

func NewGormDao() (dao *GormDao) {
	dao = &GormDao{}
	dao.Config = DefaultConfig()
	return
}
