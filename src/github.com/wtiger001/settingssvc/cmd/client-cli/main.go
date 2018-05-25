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
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/andrew-d/go-termutil"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/wtiger001/settingssvc/models"
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

// Config ...
type Config struct {
	Baseurl    string
	PrettyJSON bool
}

var debug = false
var pipedata string
var cfg *Config

func main() {
	debug = contains(os.Args, "debug")
	if termutil.Isatty(os.Stdin.Fd()) {
		pipedata = ""
	} else {
		pipedata = getDataFromPipe()
	}

	if len(os.Args) < 2 {
		fatalf("Not enough Arguments")
	}

	op := os.Args[1]

	if op == "configure" {
		handleConfigureCommand()
	}

	// read config
	cfg = readConfig()
	extractConfig()

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

func getDataFromPipe() string {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output)
}

func handleAdd() {
	if len(os.Args) < 3 {
		fatalf("Not enough arguments")
	}

	r := &request{
		modelType: os.Args[2],
		method:    "post",
	}
	completeBody(r)
	issueRequest(r)
}

func completeBody(r *request) {
	fields := mapArgs()

	if pipedata != "" {
		r.body = pipedata
	} else if val, ok := fields["data"]; ok {
		r.body = val
	} else if val, ok := fields["f"]; ok {
		bytes, err := ioutil.ReadFile(val)
		fatal(err)
		r.body = string(bytes)
	} else {
		object, err := makeObject(r.modelType, fields)
		fatal(err)
		r.body = object
	}
}

func handleDelete() {
	if len(os.Args) < 4 {
		fatalf("Not enough arguments")
	}

	r := &request{
		modelType: os.Args[2],
		id:        os.Args[3],
		method:    "delete",
	}

	issueRequest(r)
}

func handleUpdate() {
	if len(os.Args) < 4 {
		fatalf("Not enough arguments")
	}

	r := &request{
		modelType: os.Args[2],
		id:        os.Args[3],
		method:    "put",
	}
	completeBody(r)
	issueRequest(r)
}

func handleGet() {
	id := ""
	if len(os.Args) >= 4 && !strings.Contains(os.Args[3], "=") {
		id = os.Args[3]
	}
	r := &request{
		modelType: os.Args[2],
		id:        id,
		method:    "get",
	}

	fields := mapArgs()
	if val, ok := fields["version"]; ok {
		r.version = val
	} else {
		if contains(os.Args, "versions") && r.modelType == "profile" {
			r.modelType = "profile/versions"
		}
	}

	issueRequest(r)
}

type request struct {
	modelType string
	id        string
	method    string
	body      string
	version   string
}

func (r *request) url() string {
	val := getBaseURL() + "/" + r.modelType
	if r.id != "" {
		val += "/" + r.id
	}
	return val
}

func issueRequest(r *request) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	debugf(strings.ToUpper(r.method))
	url := r.url()
	debugf(url)

	req, err := http.NewRequest(r.method, url, nil)
	if err != nil {
		fatal(err)
	}
	if r.body != "" {
		debugf("body: %s", r.body)
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(bytes.NewReader([]byte(r.body)))
	}

	res, err := netClient.Do(req)
	fatal(err)
	debugf("Status Code: %v", res.StatusCode)

	if res.StatusCode >= 400 && res.StatusCode < 500 {
		fatalf(res.Status)
	}
	if res.StatusCode >= 500 {
		fatalf(res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fatal(err)
	}
	if len(bytes) == 0 && r.body != "" {
		fmtJSON(r.body)
	} else if len(bytes) > 0 {
		fmtJSON(string(bytes))
	}

}

func extractConfig() {
	fields := mapArgs()

	if base, ok := fields["baseurl"]; ok {
		cfg.Baseurl = base
	}
	if pretty, ok := fields["pretty"]; ok {
		result, err := strconv.ParseBool(pretty)
		fatal(err)
		cfg.PrettyJSON = result
	}
}

func readConfig() *Config {
	myconfig := &Config{
		Baseurl:    "",
		PrettyJSON: true,
	}

	// Read from .prefs file in home dir
	dir, err := homedir.Dir()
	if err != nil {
		return myconfig
	}
	configFile := filepath.Join(dir, ".prefs_command")
	debugf("Reading configuration from %s", configFile)

	if _, err := os.Stat(configFile); err != nil {
		// fatalf("No Configuration")
		return myconfig
	}

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		// fatal(err)
		return myconfig
	}

	err = json.Unmarshal(bytes, &myconfig)
	if err != nil {
		// fatal(err)
		return myconfig
	}
	return myconfig
}

func handleConfigureCommand() {
	configureCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	baseURL := configureCmd.String("baseurl", "http://127.0.0.1:4201", "Base Url where preferences service is located")
	pretty := configureCmd.Bool("prettyJSON", true, "Format the JSON")

	configureCmd.Parse(os.Args[2:])
	if configureCmd.Parsed() {
		myconfig := new(Config)
		myconfig.Baseurl = *baseURL
		myconfig.PrettyJSON = *pretty

		debugf("Configuration: BaseURL:  %s", myconfig.Baseurl)
		debugf("Configuration: PrettyJSON:  %v", myconfig.PrettyJSON)

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
	if err != nil {
		fmt.Printf("Errors: %s\n", err.Error())
		os.Exit(1)
	}
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
				fields[strings.ToLower(parts[0])] = parts[1]
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

// Make a url for the command
func (cmd *command) url() (string, error) {
	return "", nil
}

func makeObject(modelType string, fields map[string]string) (string, error) {
	var item interface{}
	var err error
	switch modelType {
	case "category":
		item, err = makeCategory(fields)
	case "type":
		item, err = makeType(fields)
	case "definition":
		item, err = makeDefinition(fields)
	case "owner":
		item, err = makeOwner(fields)
	case "profile":
	}

	if err != nil {
		fatal(err)
	}

	bytes, err := json.Marshal(item)
	if err != nil {
		fatal(err)
	}

	return string(bytes), nil
}

func makeDefinition(fields map[string]string) (*models.PreferenceDefinition, error) {
	item := &models.PreferenceDefinition{}

	if val, ok := fields["id"]; ok {
		item.ID = val
	} else {
		item.ID = makeUUID()
	}
	if val, ok := fields["name"]; ok {
		item.Name = val
	}
	if val, ok := fields["order"]; ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		item.Order = int64(v)
	}
	if val, ok := fields["category"]; ok {
		item.Category = val
	}

	// if val, ok := fields["layout"]; ok {
	// 	item.Layout = val
	// }

	if val, ok := fields["schema"]; ok {
		item.Schema = val
	}

	return item, nil
}

func makeCategory(fields map[string]string) (*models.Category, error) {
	item := &models.Category{}

	if val, ok := fields["id"]; ok {
		item.ID = val
	} else {
		item.ID = makeUUID()
	}
	if val, ok := fields["name"]; ok {
		item.Name = val
	}
	if val, ok := fields["order"]; ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		item.Order = int64(v)
	}

	return item, nil
}

func makeOwner(fields map[string]string) (*models.PreferenceOwner, error) {
	item := &models.PreferenceOwner{}

	if val, ok := fields["id"]; ok {
		item.ID = val
	} else {
		item.ID = makeUUID()
	}
	if val, ok := fields["active"]; ok {
		item.Active = val
	}
	if val, ok := fields["type"]; ok {
		item.Type = val
	}

	return item, nil
}

func makeType(fields map[string]string) (*models.OwnerType, error) {
	item := &models.OwnerType{}

	// item.Definitions = make([]string, 0, 0)
	if val, ok := fields["id"]; ok {
		item.ID = val
	} else {
		item.ID = makeUUID()
	}
	if val, ok := fields["name"]; ok {
		item.Name = val
	}
	if val, ok := fields["description"]; ok {
		item.Description = val
	}
	return item, nil
}

func makeUUID() string {
	uid, err := uuid.NewRandom()
	fatal(err)
	debugf(uid.String())
	return uid.String()
}

func fmtJSON(data string) {
	if cfg.PrettyJSON {
		var out bytes.Buffer
		err := json.Indent(&out, ([]byte)(data), "", "  ")
		fatal(err)
		fmt.Print(out.String())
	} else {
		fmt.Print(data)
	}
}

func getBaseURL() string {
	// First check the args
	baseurl := cfg.Baseurl
	if baseurl == "" {
		fatalf("No Baseurl. Please provide a base url argument or run the configure command")
	}
	if strings.HasSuffix(baseurl, "/") {
		baseurl = baseurl[:len(baseurl)-1]
	}

	return baseurl
}
