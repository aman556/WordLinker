package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func LinkWords(dirPath string) error {
	var fileConfig FilesConfig
	err := fileConfig.AddConfig()
	if err != nil {
		return err
	}

	var files Files
	err = files.FilesList(dirPath, fileConfig)
	if err != nil {
		return err
	}

	for _, f := range files.List {
		if err = linkingWords(f); err != nil {
			return err
		}
	}
	return nil
}

func (f *FilesConfig) AddConfig() error {
	// Read the file
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &f)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func (f *Files) FilesList(dirPath string, fileConfig FilesConfig) error {
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			for _, exclude := range fileConfig.Exclude {
				if strings.HasPrefix(path, exclude) {
					return nil
				}
			}
			extension := filepath.Ext(path)
			for _, ext := range fileConfig.Include {
				if ext == extension {
					f.List = append(f.List, path)
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func linkingWords(filePath string) error {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
