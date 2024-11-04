package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IConfig interface{}

func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Config holds the configuration settings.
type Config struct {
	Port       string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
}

// NewConfig creates a new Config instance.
func NewConfig(filename string) *Config {
	fmt.Println("FILENAME: ", filename)
	if err := loadEnv(filename); err != nil {
		fmt.Println("Error loading file")
		return nil
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error converting DB_PORT to integer")
		return nil
	}

	return &Config{
		Port:       os.Getenv("PORT"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     dbPort,
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
