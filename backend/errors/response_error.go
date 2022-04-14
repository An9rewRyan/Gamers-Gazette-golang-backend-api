package errors

import (
	"io"
	"log"
	"net/http"
)

func RespErr(w http.ResponseWriter, err error) {
	_, er := io.WriteString(w, err.Error())
	if er != nil {
		log.Println(err)
	}
}
