package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

//func gerSecret() (string, error) {
//	var T struct {
//		App struct {
//			Ip     string `json:"ip"`
//			Port   string `json:"port"`
//			Secret string `json:"secret"`
//		} `json:"app"`
//	}
//	str, _ := config.GetConfig("DEFAULT_GROUP", conts.ServiceName)
//	json.Unmarshal([]byte(str), &T)
//	return T.App.Secret, nil
//}

func GetJwtToken(iat, seconds, payload int64) (string, error) {
	//secret, err := gerSecret()
	//if err != nil {
	//	return "", err
	//}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte("12345"))
}
func SetJwtToken(tokens string) (int64, error) {
	//secret, err := gerSecret()
	//if err != nil {
	//	return 0, err
	//}
	token, _ := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return []byte("12345"), nil
	})
	client := token.Claims.(jwt.MapClaims)
	if _, ok := client["payload"].(int64); !ok {
		return 0, errors.New("token invalid")
	}
	return client["payload"].(int64), nil
}
