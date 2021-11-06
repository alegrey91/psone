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
	homeEnv := os.Getenv("HOME")

	data, err := os.ReadFile(homeEnv + "/.psone.yaml")
	if err != nil {
		fmt.Printf("error reading ~/.psone.yaml file.\n")
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Printf("cannot unmarshal data: %v\n", err)
		os.Exit(1)
	}
	return &conf
}

// ListPS1 show the list of your saved PS1.
func ListPS1() {
	conf := getConfig()

	fmt.Println("psones:")
	for ps1, _ := range conf.Psones {
		fmt.Printf("  name: %s:\n    value: \"%s\"\n", ps1, conf.Psones[ps1].Value)
	}
}

// SetPS1 write to file one of the PS1 in the list.
func SetPS1(NewPS1 string, debug bool) {
	conf := getConfig()
	homeEnv := os.Getenv("HOME")

	// Check if passed PS1 exists in config file.
	if _, err := conf.Psones[NewPS1]; !err {
		fmt.Printf("PS1 \"%s\" does not exists.\nCheck available PS1 with: psone get\n", NewPS1)
		os.Exit(1)
	}

	data, err := os.ReadFile(homeEnv + "/.bashrc")
	if err != nil {
		fmt.Printf("error reading ~/.bashrc file.\n")
		os.Exit(1)
	}

	// Find PS1 in .bashrc and replace with new one.
	match := regexp.MustCompile(`(?m)(^export PS1\=)\"(.*)\"$`)
	res := match.ReplaceAllString(string(data), "${1}"+"\""+conf.Psones[NewPS1].Value+"\"")

	// If argument --debug is passed, print the result output (whitout write to file).
	if !debug {
		err := os.WriteFile(homeEnv+"/.bashrc", []byte(res), 0644)
		if err != nil {
			fmt.Printf("error updating .bashrc file.\n")
			os.Exit(1)
		}
	} else {
		fmt.Println(string(res))
	}
}

// AddPS1 add a new PS1 to your list.
func AddPS1(PS1Name string, PS1Value string) {
	conf := getConfig()
	homeEnv := os.Getenv("HOME")

	if _, ok := conf.Psones[PS1Name]; ok {
		fmt.Printf("PS1 %s already exists.\n", PS1Name)
		os.Exit(1)
	}

	conf.Psones[PS1Name] = Psone{
		Value: PS1Value,
	}

	data, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Printf("cannot marshall data %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(homeEnv+"/.psone.yaml", []byte(data), 0644)
	if err != nil {
		fmt.Printf("error updating .psone.yaml file.\n")
		os.Exit(1)
	}
}

// RemovePS1 remove a PS1 from your list.
func RemovePS1(PS1Name string) {
	conf := getConfig()
	homeEnv := os.Getenv("HOME")

	if _, ok := conf.Psones[PS1Name]; !ok {
		fmt.Printf("PS1 %s does not exists.\n", PS1Name)
		os.Exit(1)
	}
	delete(conf.Psones, PS1Name)

	data, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Printf("cannot marshall data %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(homeEnv+"/.psone.yaml", []byte(data), 0644)
	if err != nil {
		fmt.Printf("error updating .psone.yaml file.\n")
		os.Exit(1)
	}
}
