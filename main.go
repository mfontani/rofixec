package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "", "JSON configuration file to use. Required.")
	flag.Parse()
	if configFileName == "" {
		panic("Need a config file name. See -config path/to/file.json")
	}
}

func main() {
	items := unmarshalConfig(configFileName)
	pickedOption := flag.Arg(0)
	// Run the requested command, with args
	if pickedOption != "" {
		for _, v := range items {
			if pickedOption == v.Name {
				cmd := exec.Command(v.Exec, v.Args...)
				err := cmd.Start()
				if err != nil {
					panic(err)
				}
				os.Exit(0)
			}
		}
		panic("No such option found in " + configFileName + ": " + pickedOption)
	}
	// Display the list
	for _, v := range items {
		fmt.Println(v.Name)
	}
	os.Exit(0)
}
