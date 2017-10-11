package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func getCWD() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return usr.HomeDir
}

func filterHomeDir(path, homeDir string) string {
	return strings.Replace(path, "~", homeDir, -1)
}

func addCWD(path, cwd string) (string, bool) {
	if strings.Contains(path, "~/") {
		return path, true
	}

	return cwd + "/" + path, false
}

func executeScript(script string) {
}

func main() {

	var t Config

	configName := flag.String("config", "config.yaml", "the name of the config file")
	pro := flag.String("profile", "default", "the profile, specified in the config, to be used for the procedure")
	force := flag.Bool("force", false, "should the action specified be forced")
	del := flag.Bool("delete", false, "deletes any generated symlinks or copies specified in the config file")

	flag.Parse()

	fmt.Println("Using profile: ", *pro)

	NewConfig(*configName, &t)

	if len(t.Profiles[*pro]) == 0 {
		fmt.Printf("Error: specified profile (%s) is not included in config file\n", *pro)
		if len(t.Profiles) > 0 {
			fmt.Printf("Please choose from:\n")
			for k, _ := range t.Profiles {
				fmt.Printf("\t%s\n", k)
			}
			fmt.Println("e.g. crossy -profile home")
		}
		return
	}

	pwd := getCWD()
	home := getHomeDir()

	fmt.Println(t.Profiles[*pro])
	fmt.Println(t.Profiles[*pro][0]["vimrc"].Before)

	for i := 0; i < len(t.Profiles[*pro]); i++ {
		for k, v := range t.Profiles[*pro][i] {
			v.Link = filterHomeDir(v.Link, home)
			k, needHome := addCWD(k, pwd)
			if needHome {
				filterHomeDir(k, home)
			}

			executeScript(v.Before)

			if *force || *del {
				fmt.Println("Removing:", v.Link)
				err := os.Remove(v.Link)
				if err != nil {
					fmt.Println(err)
				}
			}
			if !(*del) {
				fmt.Println("Symlinking", v.Link, "to", k)
				err := os.Symlink(k, v.Link)
				if err != nil {
					fmt.Println(err)
				}
				executeScript(v.After)
			}
		}
	}
}
