package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
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
	key_flag := flag.String("k", "", "")
	family_flag := flag.String("f", "", "")
	detail_flag := flag.String("d", "", "")
	list_flag := flag.Bool("l", false, "")
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
	fmt.Print("\n\n")

	if *list_flag {

		for _, cnt := range data.Info {
			fmt.Println(" >:", cnt.Key)
			fmt.Println("========================")
		}
	} else {
		info_idx := slices.IndexFunc(data.Info, func(i Info) bool {
			return strings.ToLower(i.Key) == strings.ToLower(*key_flag)
		})
		info := data.Info[info_idx]

		if *key_flag != "" && *family_flag == "" && *detail_flag == "" {
			for _, cnt := range info.Content {
				fmt.Println(" >:", cnt.Family)
				fmt.Println("========================")
			}
		} else if *family_flag != "" && *key_flag != "" && *detail_flag == "" {
			content_idx := slices.IndexFunc(info.Content, func(c Content) bool {
				return strings.ToLower(c.Family) == strings.ToLower(*family_flag)
			})
			content := info.Content[content_idx]

			for _, exec := range content.Exec {
				fmt.Println("====================== ======================")
				fmt.Println("-", exec.Description)

				fmt.Println("\n>", exec.Command)
			}
		}
	}
}
