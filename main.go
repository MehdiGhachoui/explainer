package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Content struct {
	Family      string `json:"family"`
	Description string `json:"description"`
	Command     string `json:"command"`
}
type Info struct {
	Key     string    `json:"key"`
	Content []Content `json:"content"`
}
type FileData struct {
	Info []Info `json:"info"`
}

func main() {
	// reading from file
	var data FileData
	jsonFile, err := os.Open("./explainer.json")
	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	//io is better than os - since os read the entire file to memory (bad for big files) -
	jsonByte, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonByte, &data)
	for i := 0; i < len(data.Info); i++ {
		fmt.Println(data.Info[i].Key)
		for j := 0; j < len(data.Info[i].Content); j++ {
			fmt.Println(data.Info[i].Content[j].Family)
			fmt.Println(data.Info[i].Content[j].Description)
			fmt.Println(data.Info[i].Content[j].Command)
		}
	}
}
