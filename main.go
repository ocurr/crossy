package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Loc struct {
	Link []map[string]string `yaml:",omitempty,flow"`
	Copy []map[string]string `yaml:",omitempty,flow"`
}

type T struct {
	Config map[string]Loc `yaml:",flow"`
}

func main() {

	var t T

	configName := flag.String("config", "config.yaml", "the name of the config file")

	flag.Parse()

	buf, err := ioutil.ReadFile(*configName)
	if err != nil {
		fmt.Println(err)
		return
	}

	yaml.Unmarshal(buf, &t)

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n\nWorking Directory: %s\n\n", path)

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	home := usr.HomeDir

	for i := 0; i < len(t.Config["home"].Link); i++ {
		for k, v := range t.Config["home"].Link[i] {
			if v[:2] == "~/" {
				v = filepath.Join(home, v[2:])
			}
			err = os.Symlink(filepath.Join(path, k), v)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	/*
		for i := 0; i < len(t.Config["home"].Copy); i++ {
			for k, v := range t.Config["home"].Copy[i] {
					err = os.Symlink(k, v)
					if err != nil {
						fmt.Println(err)
					}
			}
		}
	*/
}
