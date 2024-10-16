package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	err := godotenv.Load(".env." + env)

	if err != nil {
		fmt.Println("加载环境变量失败")
	}
}
