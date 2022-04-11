package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"library/utils/yamlparse"
	"time"
)

var TokenExpireDuration time.Duration
var Secret string
var Issuer string

type Conf struct {
	Secret  string `yaml:"Secret"`
	Issuer  string `yaml:"Issuer"`
	Expires int    `yaml:"Expires"`
}

type Claims struct {
	ID   int      `json:"user_id"`
	Role RoleType `json:"user_role"`
	jwt.StandardClaims
}

func InitAuth() {
	confMap, err := readConfig("auth")
	if err != nil {
		panic(err)
	}

	var conf Conf
	conf, _ = confMap["Library"]
	Secret = conf.Secret
	Issuer = conf.Issuer
	TokenExpireDuration = time.Hour * time.Duration(conf.Expires)
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

func GenToken(id int, role RoleType) (string, error) {
	c := Claims{id, role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    Issuer,                                     // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	key, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return "Bearer " + key, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
