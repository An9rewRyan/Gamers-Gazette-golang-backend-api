package test

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	data := `{"response":[{"id":497677293,"first_name":"Michael","last_name":"Goldberg"}]}`
	fmt.Println(gjson.Get(data, "response.#.id"))
}
