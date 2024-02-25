package config

import (
	"flag"
	"fmt"
	"log"
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
	gottenEnv := flag.String("env", "development", "Environment name e.g. test, development, staging, production")

	validateEnv(*gottenEnv)

	env.name = *gottenEnv
}

func (env *Environment) loadEnvs() {
	if err := godotenv.Load(fmt.Sprintf(".env.%s", env.name)); err != nil {
		log.Fatalf("Error loading %s file", fmt.Sprintf(".env.%s", env.name))
	}

	if err := godotenv.Load(".env"); err != nil {
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
