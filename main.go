package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting cwd:", err)
		return
	}

	appPath, err := getAppPath()
	if err != nil {
		return
	}
	repoUrls, err := getXbpsRepos(cwd)
	if err != nil {
		return
	}

	appDir := filepath.Join(appPath, "void-repos")
	err = cleanDir(appDir)
	if err != nil {
		return
	}

	fmt.Println("Trying to create app dir:", appDir)
	err = os.MkdirAll(appDir, os.ModePerm)
	if err != nil {
		fmt.Println("error creating app directory")
		return
	}
	fmt.Println("Created app dir:", appDir)
	for _, url := range repoUrls {
		if err := cloneRepo(appDir, url); err != nil {
			return
		}
	}
	// if err = cloneRepo(appDir, repoUrls[0]); err != nil {
	// 	return
	// }
	fmt.Println("Successfully cloned repos")
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

func getXbpsRepos(cwd string) (repoUrls []string, err error) {
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
		err = nil
		return
	}
	err = os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error deleting app dir:", err)
	}
	fmt.Println("Successfully cleaned app dir")
	return
}

func cloneRepo(appDir string, repoUrl string) (err error) {
	parsedUrl, err := url.Parse(repoUrl)
	if err != nil {
		return
	}
	repoWithExt := path.Base(parsedUrl.Path)
	repoName := strings.TrimSuffix(repoWithExt, ".git")

	_, err = git.PlainClone(appDir+"/"+repoName, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error cloning repo:", repoUrl)
		return
	}
	fmt.Println("Successfully cloned repo", repoUrl)
	return
}
