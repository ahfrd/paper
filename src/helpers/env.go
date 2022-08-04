package helpers

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	// log "micro-template/src/helpers"
)

// Env is a struct for env helpers
type Env struct{}

const projectDirName = "micro-emoney" // change to relevant project name

// StartingCheck is a function for checking a env ready to running
func (o Env) StartingCheck() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// logModule.Initialize{
	// 	App:    os.Getenv("APP"),
	// 	Status: os.Getenv("STATUS"),
	// 	Port:   os.Getenv("MYPORT"),
	// }.InitializeLogging()

}

// TestStartingCheck is a
func (o Env) TestStartingCheck() {
	loadEnvTesting()
	logrus.SetOutput(ioutil.Discard)
	// logModule.Initialize{
	// 	App:    os.Getenv("APP"),
	// 	Status: os.Getenv("STATUS"),
	// 	Port:   os.Getenv("MYPORT"),
	// }.InitializeLogging()

}

func loadEnvTesting() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/config/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
