package test

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	data := `{"access_token":"a9c5bb08571250e43f71c29f6c93a27b4579c4587fa26d8efcf5762499ae3b056d29239fad7d58783ab07","expires_in":86400,"user_id":497677293,"email":"deutschman1999@mail.ru"}`
	fmt.Println(gjson.Get(data, "#.user_id"))
}
