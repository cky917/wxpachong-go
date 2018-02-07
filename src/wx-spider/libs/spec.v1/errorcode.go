package spec

import (
	"log"
)

func HandleError(err error) (errcode int, message string) {
	log.Println(err.Error())
	return 599, err.Error()
}
