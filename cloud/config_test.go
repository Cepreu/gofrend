package cloud

import "testing"

func TestInitialize(t *testing.T) {
	checkNil(DeleteConfig(), t)
}

func TestConfig(t *testing.T) {
	config := map[string]interface{}{"Color": "Red"}
	err := UpdateConfig(config)
	checkNil(err, t)
	config, err = GetConfig()
	checkNil(err, t)
	if config["Color"] != "Red" {
		panic("Color not red")
	}
	err = DeleteConfig()
	checkNil(err, t)
	config, err = GetConfig()
	checkNil(err, t)
	_, ok := config["Color"]
	if ok {
		panic("Color in map")
	}
}
