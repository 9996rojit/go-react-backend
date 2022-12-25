package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BodyParser(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		fmt.Println("This is a message body that is send by client to parse", body)
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
