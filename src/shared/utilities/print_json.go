package utilities

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func PrintJson(data any) {
	str, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		str := `{
				"str": "foo",
				"num": 100,
				"bool": false,
				"null": null,
				"array": ["foo", "bar", "baz"],
				"obj": { "a": 1, "b": 2 }
			}`
	*/
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 2

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
