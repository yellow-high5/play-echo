package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	db, err := sql.Open("sqlite3", "db/test.db")
	if err != nil {
		panic("failed to connect database")
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		fmt.Print(err)
		panic("failed to database migration")
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func GetInstance() *gorm.DB {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Warn,            // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
