package config

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

var (
	DefaultPS1Name     string = "default"
	DefaultPS1Value    string = `[\u@\h \W] 🎮 $ `
	DefaultPS1FileName string = ".psone.yaml"
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

	data, err := os.ReadFile(homeEnv + "/" + DefaultPS1FileName)
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
		fmt.Printf("  %s:\n    value: \"%s\"\n", ps1, conf.Psones[ps1].Value)
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

	err = os.WriteFile(homeEnv+"/"+DefaultPS1FileName, []byte(data), 0644)
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

	err = os.WriteFile(homeEnv+"/"+DefaultPS1FileName, []byte(data), 0644)
	if err != nil {
		fmt.Printf("error updating .psone.yaml file.\n")
		os.Exit(1)
	}
}

// AddPS1 add a new PS1 to your list.
func GenerateFilePS1(Path string, Force bool) {
	conf := Configs{}
	conf.Psones = make(map[string]Psone)
	conf.Psones[DefaultPS1Name] = Psone{
		Value: DefaultPS1Value,
	}

    pathToConfigFile := ""
	homeEnv := os.Getenv("HOME")
	// Checking if --output has been provided.
	if Path != "" {
	    pathToConfigFile = Path + DefaultPS1FileName
	} else {
        pathToConfigFile = homeEnv + "/" + DefaultPS1FileName
	}

	data, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Printf("cannot marshall data %v\n", err)
		os.Exit(1)
	}

    // Checking if file already exists.
    if _, err := os.Stat(pathToConfigFile); err == nil {
        // If --force option is enabled, then override config file.
        if Force {
            err = os.WriteFile(pathToConfigFile, []byte(data), 0644)
            if err != nil {
                fmt.Printf("error generating file to %s.\n", pathToConfigFile)
                os.Exit(1)
            }
        } else {
            fmt.Printf("file %s already exists.\nuse option --force if you want to override it.\n", pathToConfigFile)
            os.Exit(1)
        }
    }
    // File doesn't exist, so we create it.
    err = os.WriteFile(pathToConfigFile, []byte(data), 0644)
    if err != nil {
        fmt.Printf("error generating file to %s.\n", pathToConfigFile)
        os.Exit(1)
    }
}
