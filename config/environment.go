package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var env Environment

func InitEnvironment() {
	env = Environment{}
	env.loadEnvironment()
	env.loadEnvs()
}

func GetName() string {
	return env.name
}

type Environment struct {
	name string
}

func (env *Environment) loadEnvironment() {
	gottenEnv := os.Getenv("ENV")

	if gottenEnv == "" {
		gottenEnv = "development"
	} else {
		validateEnv(gottenEnv)
	}

	env.name = gottenEnv
}

func (env *Environment) loadEnvs() {
	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	localDir := strings.Split(currentDir, "go_webserver")[0]

	if err := godotenv.Load(fmt.Sprintf("%s/go_webserver/.env.%s", localDir, env.name)); err != nil {
		log.Fatalf("Error loading %s file", fmt.Sprintf(".env.%s", env.name))
	}

	if err := godotenv.Load(fmt.Sprintf("%s/go_webserver/.env", localDir)); err != nil {
		log.Println("Error loading .env file")
	}
}

func validateEnv(envToValidate string) bool {
	supportedEnvs := []string{"test", "development", "staging", "production"}

	for _, env := range supportedEnvs {
		if envToValidate == env {
			return true
		}
	}

	panic(fmt.Sprintf("Not supported env:%s\n", envToValidate))
}
