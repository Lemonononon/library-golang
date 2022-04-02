package yamlparse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func LoadYamlFile(fileName string, res interface{}) error {
	dbBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(dbBytes, res); err != nil {
		log.Fatal(fmt.Sprintf("yaml parse error, file name: %v", fileName))
		return err
	}
	fmt.Println(res)
	return err
}
