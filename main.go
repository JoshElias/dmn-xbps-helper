package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func main() {
	appPath, err := getAppPath()
	if err != nil {
		return
	}
	repoUrls, err := getXbpsRepos()
	if err != nil {
		return
	}

	for _, url := range repoUrls {
		fmt.Println(url)
	}

	appDir := filepath.Join(appPath, "void-repos")
	err = cleanDir(appDir)
	if err != nil {
		return
	}
}

func getAppPath() (path string, err error) {
	viper.AutomaticEnv()
	path = viper.GetString("HOME")
	if path == "" {
		err = errors.New("HOME is not defined")
		return
	}
	fmt.Println("App path:", path)
	return
}

func getXbpsRepos() (repoUrls []string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting cwd:", err)
		return
	}
	fmt.Println("CWD:", cwd)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cwd)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to read config")
	}

	repoUrls = viper.GetStringSlice("repoUrls")
	if repoUrls == nil {
		fmt.Print("Please provide repo urls")
	}
	fmt.Println("repoUrls:", repoUrls)
	return
}

func cleanDir(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println("App dir does not exist")
		return
	}
	err = os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error deleting app dir:", err)
	}
	return
}
