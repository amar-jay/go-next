package cmd

import (
	"log"
	"os"
	"os/user"
	"strconv"

	models "github.com/amar-jay/go-api-boilerplate/database/domain/session"
	"github.com/amar-jay/go-api-boilerplate/pkg/config"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
// router = gin.Default()
// db   *gorm.DB
// err  error
// conf config.Config
)

func Initialize() (conf config.Config, db *gorm.DB, port string, err error) {
	// load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env files. $PORT must be set")
	}

	conf = config.GetConfig()
	db, err = gorm.Open(
		conf.Postgres.GetConnectionInfo(),
		conf.Postgres.Config(),
	)

	if err != nil {
		panic(err)
	}

	//port = strconv.Itoa(config.Port);
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = strconv.Itoa(conf.Port)
	}

	return
}

func Main() {
}

func Migrate(db *gorm.DB) (err error) {
	// all gorm database schema migrations go here
	// Migrate the schema
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in Migrate", r)
		}
	}()
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}
	// err = db.AutoMigrate(&password_reset.PasswordReset{})
	// if err != nil {
	// 	panic(err)
	// }
	err = db.AutoMigrate(&models.Session{})
	if err != nil {
		panic(err)
	}

	log.Println("Database migrated successfully")
	return
}
