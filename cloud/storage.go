package cloud

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/Cepreu/gofrend/utils"
	"google.golang.org/api/option"
)

// UploadXML uploads static xml file to cloud
func UploadXML(data []byte) error {
	object, err := getObjectHandleFromData(data)
	if err != nil {
		return err
	}

	ctx := context.Background()

	writer := object.NewWriter(ctx) // Currently writes regardless of whether or not file already exists. There may or may not be reason to rethink.
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	return nil
}

// DownloadXML downloads static xml file from cloud
func DownloadXML(hash string) ([]byte, error) {
	object, err := getObjectHandleFromHash(hash)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

func getObjectHandleFromData(data []byte) (*storage.ObjectHandle, error) {
	hash := utils.HashToString(data)
	return getObjectHandleFromHash(hash)
}

func getObjectHandleFromHash(hash string) (*storage.ObjectHandle, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return nil, err
	}
	object := client.Bucket(GcpStorageBucketName).Object(fmt.Sprintf("five9ivr-files/%s", hash))
	return object, nil
}