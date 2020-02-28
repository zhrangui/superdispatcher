package db

import (
	"superdispatcher/config"
	"superdispatcher/logger"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// SQL holds sql database connection string ...
type SQL struct {
	config *config.Config
	*logger.Logger
	*gorm.DB
}

type AtlasUsers struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primary_key; column:ID"`
	IP           string    `gorm:"column:IP"`
	ComputerName string    `gorm:"column:ComputerName"`
	Version      string    `gorm:"column:Version"`
}

func (AtlasUsers) TableName() string {
	return "AtlasUsers"
}

// NewSQL constructs SQL interface to interact with SQL database
func NewSQL(config *config.Config) (*SQL, error) {
	return &SQL{
		config: config,
	}, nil
}

// Open database connection
func (sql *SQL) Open() error {
	connString := sql.config.Constants.MSSQL.ConnectionString
	db, err := gorm.Open("mssql", connString)
	if err == nil {
		sql.DB = db
	}
	return err
}

// Close database connection
func (sql *SQL) Close() {
	sql.DB.Close()
}
