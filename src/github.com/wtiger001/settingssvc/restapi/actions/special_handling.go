package actions

import (
	"fmt"
	"github.com/wtiger001/settingssvc/models"
	"os"

	runtime "github.com/go-openapi/runtime"

	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	stor "github.com/wtiger001/settingssvc/store"
	"strings"
)

var store stor.ItemStore

// GetStore profides the current store
func GetStore() stor.ItemStore {
	return store
}

// SetupStore Initializes the choosen store
func SetupStore() {
	storeType := strings.ToUpper(os.Getenv("SETTINGS_STORE"))
	switch storeType {
	case "FILE":
		fmt.Println("Using File Store")
		store = stor.NewFileStore("c:/dev/fs")
	case "SQL":
		fmt.Println("Using SQL Store")
		store = stor.NewSQLStore()
	default:
		fmt.Println("No store selected: Using SQL Store as default")
		store = stor.NewSQLStore()
	}
	if err := store.InitStore(); err != nil {
		panic(err)
	}

}

//var prefDefStore = make(map[string]map[string]*json.RawMessage)

// ProducePreferenceDef handles the special logic for
// encoding these guys. We store the parsed data as a map
func ProducePreferenceDef(writer io.Writer, data interface{}) error {
	var sample *models.PreferenceDefinition
	if reflect.TypeOf(data) == reflect.TypeOf(sample) {
		fmt.Printf("I AM SO SPECIAL\n")
	}
	fmt.Printf("PROD: %s", data)
	fmt.Printf("PROD2: %s\n", reflect.TypeOf(data))

	enc := json.NewEncoder(writer)
	return enc.Encode(data)
}

// JSONConsumer creates a new JSON consumer
func JSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		var sample *models.PreferenceDefinition
		var pSample *models.Profile

		if reflect.TypeOf(data) == reflect.TypeOf(sample) {
			fmt.Println("USING SPECIAL CONSUMER")

			original, ok := data.(*models.PreferenceDefinition)
			if ok {
				all, err := ioutil.ReadAll(reader)
				if err != nil {
					fmt.Printf("\tError: %s\n", err.Error())
					return err
				}
				return stor.ParsePreferenceDefinition(all, original)
			}
		} else if reflect.TypeOf(data) == reflect.TypeOf(pSample) {
			fmt.Println("USING SPECIAL CONSUMER")

			original, ok := data.(*models.Profile)
			if ok {
				all, err := ioutil.ReadAll(reader)
				if err != nil {
					fmt.Printf("\tError: %s\n", err.Error())
					return err
				}
				return stor.ParseProfile(all, original)
			}
		}

		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		return dec.Decode(data)
	})
}

// JSONProducer creates a new JSON producer
func JSONProducer() runtime.Producer {
	return runtime.ProducerFunc(func(writer io.Writer, data interface{}) error {
		var sample *models.PreferenceDefinition
		var pSample *models.Profile
		if reflect.TypeOf(data) == reflect.TypeOf(sample) {
			original, ok := data.(*models.PreferenceDefinition)
			if ok {
				fmt.Println("USING SPECIAL PRODUCER")
				enc := json.NewEncoder(writer)
				return enc.Encode(original.JsonData)
			}
		} else if reflect.TypeOf(data) == reflect.TypeOf(pSample) {
			original, ok := data.(*models.Profile)
			if ok {
				fmt.Println("USING SPECIAL PRODUCER")
				enc := json.NewEncoder(writer)
				return enc.Encode(original.JsonData)
			}
		}

		enc := json.NewEncoder(writer)
		return enc.Encode(data)
	})
}
