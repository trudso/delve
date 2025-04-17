package comm

import (
	"encoding/json"
	"os"
)

func SaveAsJson(path, filename string, data map[string]any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic("error serializing to json: " + err.Error())
	}

	writeFile(path, filename, bytes)
}

func writeFile(path, filename string, data []byte) {
	root, err := os.OpenRoot(path)
	if err != nil {
		panic(err)
	}
	defer root.Close()

	file, err := root.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(data)
}
