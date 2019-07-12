package cloud

import (
	"fmt"

	"github.com/Cepreu/gofrend/utils"
)

// UploadXML uploads static xml file to cloud
func UploadXML(data []byte) error {
	return upload(fmt.Sprintf("five9ivr-files/%s", utils.HashToString(data)), data)
}

// DownloadXML downloads static xml file from cloud
func DownloadXML(hash string) ([]byte, error) {
	return download(fmt.Sprintf("five9ivr-files/%s", hash))
}
