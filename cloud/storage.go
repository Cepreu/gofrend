package cloud

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func download(filename string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return nil, err
	}
	object := client.Bucket(GcpStorageBucketName).Object(filename)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = reader.Close()
	if err != nil {
		return nil, err
	}
	return data, client.Close()
}

func upload(filename string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return err
	}
	object := client.Bucket(GcpStorageBucketName).Object(filename)
	writer := object.NewWriter(ctx)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	return client.Close()
}
