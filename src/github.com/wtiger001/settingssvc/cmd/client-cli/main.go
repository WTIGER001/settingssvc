package main

/*
	Example:

-- Adds a new type specified by the type.json

$ ./prefs add type -f=type.json
$ ./prefs get type abcd
$ ./prefs add type id=abc

Commands:
- add <type|category|definition|owner|profile> <data=<JSON>|f=pathtofile|fields>
- delete <type|category|definition|owner|profile> id=<IDTODLETE>
- update <type|category|definition|owner|profile> <data=<JSON>|f=pathtofile|fields>
- patch <type|category|definition|owner|profile> <data=<JSON>|f=pathtofile|fields>
- get  <type|category|definition|owner|profile|profileversions> id=<IDTODLETE> -v=<VERSION>

*/
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
)

var modelTypes = []string{
	"type", "category", "definition", "owner", "profile", "profileversions",
}

var opTypes = []string{
	"add", "delete", "update", "patch", "get", "configure",
}

type command struct {
	operation string
	modelType string
	data      string
	file      string
	fields    map[string]string
	id        string
	version   string
}

type Config struct {
	Baseurl string
}

const debug = true

var cfg *Config

func main() {
	if len(os.Args) < 2 {
		fatalf("Not enough Arguments")
	}

	op := os.Args[1]

	if op == "configure" {
		handleConfigureCommand()
	}

	// read config
	cfg = readConfig()

	switch op {
	case "add":
		handleAdd()
	case "delete":
		handleDelete()
	case "update":
		handleUpdate()
	case "get":
		handleGet()
	default:
		fmt.Printf("Invalid Command: %s, expected one of %v", op, opTypes)
	}

}

func handleAdd() {

}

func handleDelete() {

}

func handleUpdate() {

}

func handleGet() {
	modelType := os.Args[2]
	// id := ""
	// if len(os.Args) >= 4 {
	// 	id = os.Args[3]
	// }
	r := &request{
		modelType: modelType,
		method:    "get",
	}
	issueRequest(r)
}

type request struct {
	modelType string
	method    string
}

func (r *request) url() string {
	val := cfg.Baseurl + "/" + r.modelType

	return val
}

func issueRequest(r *request) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	url := r.url()
	switch r.method {
	case "get":
		res, err := netClient.Get(url)
		if err != nil {
			fatal(err)
		}
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fatal(err)
		}
		fmt.Println(string(bytes))
	case "delete":
		req, err := http.NewRequest("DELETE", url, nil)

		if err != nil {
			fatal(err)
		}
		res, err := netClient.Do(req)
		if err != nil {
			fatal(err)
		}
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fatal(err)
		}
		fmt.Println(string(bytes))
	}

	// response, _ := netClient.Get(cmd.url())

}

func readConfig() *Config {
	// Read from .prefs file in home dir
	dir, err := homedir.Dir()
	if err != nil {
		fatal(err)
	}
	configFile := filepath.Join(dir, ".prefs_command")
	debugf("Reading configuration from %s", configFile)

	if _, err := os.Stat(configFile); err != nil {
		fatalf("No Configuration")
	}

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		fatal(err)
	}
	var myconfig *Config
	err = json.Unmarshal(bytes, &myconfig)
	if err != nil {
		fatal(err)
	}
	return myconfig
}

func handleConfigureCommand() {
	configureCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	baseURL := configureCmd.String("baseurl", "http://127.0.0.1:4201", "Base Url where preferences service is located")

	configureCmd.Parse(os.Args[2:])
	if configureCmd.Parsed() {
		myconfig := new(Config)
		myconfig.Baseurl = *baseURL

		debugf("Configuration: BaseURL:  %s", myconfig.Baseurl)

		dir, err := homedir.Dir()
		if err != nil {
			fatal(err)
		}
		configFile := filepath.Join(dir, ".prefs_command")
		debugf("Writing configuration to %s", configFile)
		bytes, err := json.Marshal(myconfig)
		debugf("%s", string(bytes))

		if err != nil {
			fatal(err)
		}
		ioutil.WriteFile(configFile, bytes, 0644)
		os.Exit(0)
	}

	fmt.Printf("Invalid Configure Command\n")
	configureCmd.PrintDefaults()
	os.Exit(1)
}

func fatal(err error) {
	fmt.Printf("Errors: %s\n", err.Error())
	os.Exit(1)
}

func debugf(message string, items ...interface{}) {
	if debug {
		fmt.Printf(message+"\n", items...)
	}
}

func fatalf(message string, items ...interface{}) {
	fmt.Printf(message+"\n", items...)
	os.Exit(1)
}

func mapArgs() map[string]string {
	fields := make(map[string]string)
	for i := 3; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "=") {
			parts := strings.Split(os.Args[i], "=")
			if len(parts) == 2 {
				fields[parts[0]] = parts[1]
			}
		}
	}
	return fields
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func printUsage(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func makeRequest(cmd *command) error {
	// var netClient = &http.Client{
	// 	Timeout: time.Second * 10,
	// }

	// response, _ := netClient.Get(cmd.url())
	return nil
}

// Make a url for the command
func (cmd *command) url() (string, error) {
	return "", nil
}
