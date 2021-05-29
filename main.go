package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// Version contains the binary version. This is added at build time.
var Version = "uncommitted"

var showVersion = false

var configFileName string
var forkAndExit bool

func init() {
	flag.BoolVar(&showVersion, "version", showVersion, `Displays version information, then exits`)
	flag.StringVar(&configFileName, "config", "", "JSON or YAML configuration file to use. Required.")
	flag.BoolVar(&forkAndExit, "fork", false, "(INTERNAL) whether to execute commands synchronously, then exit.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\trofi -modi 'pick:rofixec -config CONFIG.yaml|CONFIG.json' -show pick\n\n")
		fmt.Fprintf(os.Stderr, "See also: https://github.com/mfontani/rofixec\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nThis is rofixec %s\n", Version)
	}
	flag.Parse()
	if showVersion {
		fmt.Printf("%s\n", Version)
		os.Exit(0)
	}
	if configFileName == "" {
		panic("Need a config file name. See -config path/to/file.json or -config path/to/file.yaml")
	}
}

func main() {
	items := unmarshalConfig(configFileName)
	pickedOption := flag.Arg(0)
	// Run the requested command, with args
	if pickedOption != "" {
		for _, v := range items {
			if pickedOption == v.Name {
				if len(v.Commands) > 0 {
					// Commands must be run sequentially, in a separate process so as to not hold rofi open
					if !forkAndExit {
						newArgs := make([]string, 0)
						newArgs = append(newArgs, "-fork")
						for i, v := range os.Args {
							if i == 0 {
								continue
							}
							newArgs = append(newArgs, v)
						}
						cmd := exec.Command(os.Args[0], newArgs...)
						err := cmd.Start()
						if err != nil {
							panic(err)
						}
						os.Exit(0)
					}
				}
				if v.Exec != "" {
					cmd := exec.Command(v.Exec, v.Args...)
					err := cmd.Start()
					if err != nil {
						panic(err)
					}
				}
				if len(v.Commands) > 0 {
					for _, e := range v.Commands {
						if e.Exec != "" {
							cmd := exec.Command(e.Exec, e.Args...)
							err := cmd.Run() // Run sequentially!
							if err != nil {
								panic(err)
							}
						}
					}
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
