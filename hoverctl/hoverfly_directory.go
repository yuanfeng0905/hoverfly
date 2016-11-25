package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"regexp"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
)

type HoverflyDirectory struct {
	Path string
}

func NewHoverflyDirectory(config Config) (HoverflyDirectory, error) {
	var hoverflyDirectory HoverflyDirectory

	if len(config.GetFilepath()) == 0 {
		log.Debug("Missing a config file")
		log.Debug("Creating a new  a config file")
		hoverflyDirectory = HoverflyDirectory{Path: createHoverflyDirectory(getHomeDirectory())}

		err := config.WriteToFile(hoverflyDirectory)
		if err != nil {
			log.Debug(err.Error())
			return HoverflyDirectory{}, errors.New("Could not write new config to disk")
		}

	} else {
		hoverflyDirectory = HoverflyDirectory{
			Path: path.Dir(config.GetFilepath()),
		}
	}

	return hoverflyDirectory, nil
}

func (h *HoverflyDirectory) GetAllPids() ([]int, error) {
	pids := []int{}

	files, err := ioutil.ReadDir(h.Path)
	if err != nil {
		log.Debug(err.Error())
		return nil, errors.New("Could not find directory")
	}
	pattern, err := regexp.Compile(`hoverfly.(\d{1,5}).(\d{1,5}).pid`)
	if err != nil {
		log.Debug(err.Error())
		return nil, errors.New("Could not compile regex expression")
	}

	for _, fileInfo := range files {
		// This should return something like This
		// ["hoverfly", "8888", "8500"]
		results := pattern.FindSubmatch([]byte(fileInfo.Name()))

		if len(results) == 3 {
			pid, _ := h.GetPid(string(results[1]), string(results[2]))
			pids = append(pids, pid)
		}
	}

	return pids, nil
}

func (h *HoverflyDirectory) GetPid(adminPort, proxyPort string) (int, error) {
	hoverflyPidFile := h.buildPidFilePath(adminPort, proxyPort)
	if fileIsPresent(hoverflyPidFile) {
		pidFileData, err := ioutil.ReadFile(hoverflyPidFile)
		if err != nil {
			return 0, err
		}

		pid, err := strconv.Atoi(string(pidFileData))
		if err != nil {
			return 0, err
		}

		return pid, nil
	}

	return 0, nil
}

func (h *HoverflyDirectory) WritePid(adminPort, proxyPort string, pid int) error {
	pidFilePath := h.buildPidFilePath(adminPort, proxyPort)
	if fileIsPresent(pidFilePath) {
		return errors.New("Hoverfly pid already exists")
	}
	return ioutil.WriteFile(pidFilePath, []byte(strconv.Itoa(pid)), 0644)
}

func (h *HoverflyDirectory) DeletePid(adminPort, proxyPort string) error {
	return os.Remove(h.buildPidFilePath(adminPort, proxyPort))
}

func (h *HoverflyDirectory) buildPidFilePath(adminPort, proxyPort string) string {
	pidName := fmt.Sprintf("hoverfly.%v.%v.pid", adminPort, proxyPort)
	return filepath.Join(h.Path, pidName)
}

func getHomeDirectory() string {
	homeDirectory, err := homedir.Dir()
	if err != nil {
		log.Debug(err.Error())
		log.Fatal("Unable to get home directory")
	}

	return homeDirectory
}

func createHoverflyDirectory(homeDirectory string) string {
	hoverflyDirectory := filepath.Join(homeDirectory, "/.hoverfly")

	if !fileIsPresent(hoverflyDirectory) {
		err := os.Mkdir(hoverflyDirectory, 0777)
		if err != nil {
			log.Debug(err.Error())
			log.Fatal("Could not create a .hoverfly directory")
		}

		return hoverflyDirectory
	}

	return hoverflyDirectory
}

func fileIsPresent(fileURI string) bool {
	if _, err := os.Stat(fileURI); err != nil {
		return os.IsExist(err)
	}

	return true
}
