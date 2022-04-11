package crypto

import "library/utils/yamlparse"

var SaltA string
var SaltB string
var AesKey string
var AesIv string

type Conf struct {
	SaltA  string `yaml:"SaltA"`
	SaltB  string `yaml:"SaltB"`
	AesKey string `yaml:"AesKey"`
	AesIv  string `yaml:"AesIv"`
}

func readConfig(configName string) (map[string]Conf, error) {
	confMap := make(map[string]Conf)
	confFilePath := "./config/" + configName + ".yaml"

	err := yamlparse.LoadYamlFile(confFilePath, &confMap)
	if err != nil {
		return nil, err
	}
	return confMap, nil
}
func InitCrypto() {
	confMap, err := readConfig("crypto")
	if err != nil {
		panic(err)
	}

	var conf Conf
	conf, _ = confMap["Library"]

	SaltA = conf.SaltA
	SaltB = conf.SaltB
	AesKey = conf.AesKey
	AesIv = conf.AesIv
}

// Password2Secret convert the password(plaintext) to salted, use double SHA1 and saltA saltB, database store its output
func Password2Secret(password string) string {
	return SHA1(SaltB + SHA1(SaltA+password))
}
