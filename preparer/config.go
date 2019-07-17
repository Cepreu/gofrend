package preparer

import (
	"encoding/json"
	"io/ioutil"
)

const configFileName string = "config.json"

func getConfig(filename string) (map[string]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config map[string]string
	err = json.Unmarshal(data, &config)
	return config, err
}
