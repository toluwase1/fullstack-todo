package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo/models"
)

type GormDB struct {
	DB *gorm.DB
}

func GetDB() *GormDB {
	gormDB := &GormDB{}
	gormDB.Init()
	return gormDB
}

func (g *GormDB) Init() {
	g.DB = getPostgresDB()
	if err := migrate(g.DB); err != nil {
		log.Fatalf("unable to run migrations: %v", err)
	}
}

func getPostgresDB() *gorm.DB {
	log.Printf("Connecting to postgres")
	postgresDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		"localhost", "postgres", "toluwase", "todolist", 5432) //, sslmode)
	//postgresDSN := "postgres://postgres:toluwase@localhost:5432/grovepaytest?sslmode=disable"
	postgresDB, err := gorm.Open(postgres.Open(postgresDSN))
	if err != nil {
		log.Fatal(err)
	}
	return postgresDB
}
func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		return fmt.Errorf("migrations error: %v", err)
	}
	return nil
}
