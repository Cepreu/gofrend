package cloud

import (
	"encoding/json"
	"fmt"
)

// GetConfig returns the config currently in
func GetConfig() (map[string]interface{}, error) {
	data, err := download(GcpConfigFileName)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	return config, err
}

// UpdateConfig - add members of newConfig to config in cloud
func UpdateConfig(newConfig map[string]interface{}) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}
	for k, v := range newConfig {
		switch v.(type) {
		case int, string, float64:
			break
		default:
			return fmt.Errorf("Config value type must be int, string, or float64, instead got: %T", v)
		}
		config[k] = v
	}
	marshaled, err := json.Marshal(config)
	err = upload(GcpConfigFileName, marshaled)
	return err
}

// DeleteConfig - delete all configurations from cloud
func DeleteConfig() error {
	return upload(GcpConfigFileName, []byte("{}"))
}
