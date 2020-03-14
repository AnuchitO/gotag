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

func main() {
	user := User{
		Name:     "Nong",
		Password: "pass1234",
	}

	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if _, ok := field.Tag.Lookup(encrypt); ok {
			f := reflect.ValueOf(&user).Elem().Field(i)
			pwd := f.String()
			pwdEncoded := base64.StdEncoding.EncodeToString([]byte(pwd))
			f.SetString(pwdEncoded)
		}
	}
	fmt.Println(t)
	fmt.Println(user)
}
