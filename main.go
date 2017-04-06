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
	force := flag.Bool("force", false, "should the action specified be forced")

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
			if k[:2] == "~/" {
				k = filepath.Join(home, k[2:])
			} else {
				k = filepath.Join(path, k)
			}
			if *force {
				fmt.Println("Removing:", v)
				err = os.Remove(v)
			}
			fmt.Println("Symlinking", v, "to", k)
			err = os.Symlink(k, v)
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
