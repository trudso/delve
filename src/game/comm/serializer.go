package comm

import (
	"io"
	"encoding/json"
	"os"
)

func SaveJson(path, filename string, data map[string]any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic("error serializing to json: " + err.Error())
	}

	writeFile(path, filename, bytes)
}

func LoadJson(path, filename string) map[string]any {
	bytes := readFile(path, filename)	
	mapData := make(map[string]any)
	err := json.Unmarshal(bytes, &mapData)
	if err != nil {
		panic("error deserializing from json: " + err.Error())
	}

	return mapData
}

func readFile(path, filename string) []byte {
	root, err := os.OpenRoot(path)
	if err != nil {
		panic(err)
	}
	defer root.Close()

	file, err := root.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return bytes
}

func writeFile(path, filename string, data []byte) {
	root, err := os.OpenRoot(path)
	if err != nil {
		panic(err)
	}
	defer root.Close()

	file, err := root.Create( filename )
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(data)
}
