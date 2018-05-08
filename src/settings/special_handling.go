package settings

import (
	"fmt"
	"models"

	runtime "github.com/go-openapi/runtime"

	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
)

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
					return err
				}
				return readPreferenceDefinition(all, original)
			}
		} else if reflect.TypeOf(data) == reflect.TypeOf(pSample) {
			fmt.Println("USING SPECIAL CONSUMER")

			original, ok := data.(*models.Profile)
			if ok {
				all, err := ioutil.ReadAll(reader)
				if err != nil {
					return err
				}
				return readProfile(all, original)
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

func readPreferenceDefinition(bytes []byte, data *models.PreferenceDefinition) error {

	// Unmarshall "most" of this object.
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}

	// Now unmarshal the schema and layout fields manually
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(bytes, &objMap)
	if err != nil {
		return err
	}
	data.JsonData = objMap

	// Parse schema object
	if schMsg, ok := objMap["schema"]; ok {
		schema := string(*schMsg)
		data.Schema = schema
	}

	// Layout schema object
	if layMsg, ok := objMap["layout"]; ok {
		layout := string(*layMsg)
		data.Layout = layout
	}

	fmt.Printf("ID: %s\n", data.ID)
	fmt.Printf("Cat: %s\n", data.Category)
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("schema: %s\n", data.Schema)
	fmt.Printf("layout: %s\n", data.Layout)
	return nil
}

func readProfile(bytes []byte, data *models.Profile) error {

	// Unmarshall "most" of this object.
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}

	// Now unmarshal the schema and layout fields manually
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(bytes, &objMap)
	if err != nil {
		return err
	}
	data.JsonData = objMap

	// var rawMessagesForPreferences []*json.RawMessage
	// err = json.Unmarshal(*objMap["preferences"], &rawMessagesForPreferences)
	// if err != nil {
	//     return err
	// }
	// var m map[string]string
	// var prefs = make([]*models.Preference, len(rawMessagesForPreferences) )
	// for i, rawMessage := range rawMessagesForPreferences {
	// 	p := &models.Preference{}
	// 	err = json.Unmarshal(*rawMessage, &p)
	// 	p.JsonData = rawMessage
	//     if err != nil {
	//         return err
	//     }
	//     fmt.Println(m)
	// }

	// data.JsonData = objMap

	// Parse schema object
	// if valueMsg, ok := objMap["value"]; ok {
	// 	value := string(*valueMsg)
	// 	data.Value = value
	// }
	return nil
}
