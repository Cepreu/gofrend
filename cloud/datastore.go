package cloud

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Cepreu/gofrend/ivr"
	"google.golang.org/api/option"
)

func DownloadScript(scriptHash string) (*ivr.StorageScript, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, GcpProjectID, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return nil, err
	}
	key := datastore.NameKey("IvrScript", scriptHash, nil)
	script := new(ivr.StorageScript)
	err = client.Get(ctx, key, script)
	if err != nil {
		return nil, err
	}
	return script, client.Close()
}

func UploadScript(script *ivr.StorageScript, scriptHash string) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, GcpProjectID, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return err
	}
	key := datastore.NameKey("IvrScript", scriptHash, nil)
	_, err = client.Put(ctx, key, script)
	if err != nil {
		return err
	}
	return client.Close()
}

func DeleteScript(scriptHash string) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, GcpProjectID, option.WithCredentialsFile(GcpCredentialsFileName))
	if err != nil {
		return err
	}
	key := datastore.NameKey("IvrScript", scriptHash, nil)
	err = client.Delete(ctx, key)
	if err != nil {
		return err
	}
	return client.Close()
}
