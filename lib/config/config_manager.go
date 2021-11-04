package config

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Psone struct {
	//	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Configs struct {
	Psones map[string]Psone `yaml:"psones"`
}

func getConfig() *Configs {
	conf := Configs{}

	data, err := os.ReadFile(".psone.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Printf("cannot unmarshal data: %v", err)
	}
	return &conf
}

func ListPS1() {
	conf := getConfig()

	fmt.Println("psones:")
	for ps1, _ := range conf.Psones {
        fmt.Printf("  - name: %s:\n    value: \"%s\"\n", ps1, conf.Psones[ps1].Value)
	}
}

func SetPS1(NewPS1 string, permanent bool) {
    conf := getConfig()
    homeEnv := os.Getenv("HOME")

    // Check if passed PS1 exists in config file.
    if _, err := conf.Psones[NewPS1]; ! err {
        fmt.Printf("PS1 \"%s\" does not exists.\nCheck available PS1 with: psone get\n", NewPS1)
        panic(err)
    }

    data, err := os.ReadFile(homeEnv + "/.bashrc")
    if err != nil {
        panic(err)
    }

    // Find PS1 in .bashrc and replace with new one.
    match := regexp.MustCompile(`(?m)(^export PS1\=)\"(.*)\"$`)
    res := match.ReplaceAllString(string(data), "${1}" + "\"" + conf.Psones[NewPS1].Value + "\"")

    // If argument --write has been passed, override ~/.bashrc with newest PS1.
    if permanent {
        err := os.WriteFile(homeEnv + "/.bashrc", []byte(res), 0644)
        if err != nil {
            fmt.Printf("error updating .bashrc.")
        }
    } else {
        fmt.Println(string(res))
    }
}
