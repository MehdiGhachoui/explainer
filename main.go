package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
)

type Exec struct {
	Description string `json:"description"`
	Command     string `json:"command"`
	Detail      string `json:"detail"`
}
type Content struct {
	Family string `json:"family"`
	Exec   []Exec `json:"exec"`
}
type Info struct {
	Key     string    `json:"key"`
	Content []Content `json:"content"`
}
type FileData struct {
	Info []Info `json:"info"`
}

func main() {
	// defining the flags
	key_flag := flag.String("key", "", "")
	family_flag := flag.String("family", "", "")
	detail_flag := flag.String("detail", "", "")
	flag.Parse()

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
	info_idx := slices.IndexFunc(data.Info, func(i Info) bool {
		return i.Key == *key_flag
	})
	info := data.Info[info_idx]

	content_idx := slices.IndexFunc(info.Content, func(c Content) bool {
		return c.Family == *family_flag
	})
	content := info.Content[content_idx]
	fmt.Println("Family:", content.Family)

	for _, exec := range content.Exec {
		if exec.Detail == *detail_flag {
			fmt.Println("====================== ======================")
			fmt.Println("Command:", exec.Command)
			fmt.Println("Description:", exec.Description)
		}
	}
}
