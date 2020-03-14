package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

const encrypt = "encrypt"

type User struct {
	Name     string `json:"name"`
	Password string `encrypt:"base64"`
}

type Secret struct {
	Token string `encrypt:"base64"`
}

func encryptData(user interface{}) {
	t := reflect.TypeOf(user).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if _, ok := field.Tag.Lookup(encrypt); ok {
			f := reflect.ValueOf(user).Elem().FieldByName(field.Name)
			pwd := f.String()
			pwdEncoded := base64.StdEncoding.EncodeToString([]byte(pwd))
			f.SetString(pwdEncoded)
		}
	}
}

func main() {
	user := User{
		Name:     "Nong",
		Password: "pass1234",
	}
	fmt.Println(user)

	encryptData(&user)

	fmt.Println(user)

	secret := Secret{Token: "token1234"}
	fmt.Println(secret)

	encryptData(&secret)

	fmt.Println(secret)
}
