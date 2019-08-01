package main

import (
	"fmt"

	"github.com/Cepreu/gofrend/cloud"
	"github.com/Cepreu/gofrend/utils"
)

func main() {
	accessToken := utils.GenerateAccessToken()
	fmt.Println(accessToken)
	config := map[string]string{cloud.GcpConfigAccessTokenKeyString: accessToken}
	err := cloud.UpdateConfig(config)
	if err != nil {
		fmt.Println(err)
	}
}
