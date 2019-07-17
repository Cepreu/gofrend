package cloud

import (
	"encoding/json"
)

// GetConfig returns the config currently in
func GetConfig() (map[string]string, error) {
	data, err := download(GcpConfigFileName)
	if err != nil {
		return nil, err
	}
	var config map[string]string
	err = json.Unmarshal(data, &config)
	return config, err
}

// UpdateConfig - add members of newConfig to config in cloud
func UpdateConfig(newConfig map[string]string) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}
	for k, v := range newConfig {
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
