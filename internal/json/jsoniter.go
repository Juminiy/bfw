package json

import (
	"bfw/internal/logger"
	"os"

	jsoniter "github.com/json-iterator/go"
)

// Marshal is exported by gin/json package.
// Unmarshal is exported by gin/json package.
// MarshalIndent is exported by gin/json package.
// NewDecoder is exported by gin/json package.
// NewEncoder is exported by gin/json package.
var (
	json            = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal         = json.Marshal
	MarshalToString = json.MarshalToString
	Unmarshal       = json.Unmarshal
	MarshalIndent   = json.MarshalIndent
	NewDecoder      = json.NewDecoder
	NewEncoder      = json.NewEncoder
)

func ReadJsonFileParseToMap(filePath string) (map[string]interface{}, error) {
	jsonFilePtr, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := jsonFilePtr.Close(); err != nil {
			logger.Errorf(err.Error())
		}
	}()
	var m map[string]interface{}
	jsonDecoder := NewDecoder(jsonFilePtr)
	if err = jsonDecoder.Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}

func ReadJsonFileParseToObject(filePath string, object interface{}) error {
	jsonFilePtr, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if err0 := jsonFilePtr.Close(); err != nil {
			logger.Errorf(err0.Error())
		}
	}()
	jsonDecoder := NewDecoder(jsonFilePtr)
	if err = jsonDecoder.Decode(object); err != nil {
		return err
	}
	return nil
}
