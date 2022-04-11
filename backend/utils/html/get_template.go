package html

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Get_template(template_name string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	template_dir := "templates/" + template_name
	file, err := os.Open(template_dir)
	if err != nil {
		fmt.Println("template does not exist!")
	}
	file_text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error while reading file!")
	}
	return string(file_text)
}
