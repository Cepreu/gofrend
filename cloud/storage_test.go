package cloud

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Cepreu/gofrend/utils"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var fname1 = "test_files/is_large_test.five9ivr"

// func TestBucket(t *testing.T) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx, option.WithCredentialsFile("F9-Test-Agent-0925974a682a.json"))
// 	check(err)
// 	bkt := client.Bucket("f9-test-agent.appspot.com")
// 	_, err = bkt.Attrs(ctx)
// 	check(err)
// 	// fmt.Println(attrs)

// 	data, err := ioutil.ReadFile(fname1)
// 	hashbytes := sha256.Sum256(data)
// 	hash := base64.URLEncoding.EncodeToString(hashbytes[:])
// 	object := bkt.Object(fmt.Sprintf("five9ivr-files/%s", hash))

// 	_, err = object.NewReader(ctx)
// 	if err != nil { // Check if object exists in storage bucket
// 		fmt.Println("Writing")
// 		writer := object.NewWriter(ctx)
// 		_, err = writer.Write(data)
// 		check(err)
// 		err = writer.Close()
// 		check(err)
// 	} else {
// 		fmt.Println("Not writing")
// 	}
// }

func TestUploadXML(t *testing.T) {
	data, err := ioutil.ReadFile(fname1)
	check(err)
	err = UploadXML(data)
	check(err)
}

func TestDownloadXML(t *testing.T) {
	expectedData, err := ioutil.ReadFile(fname1)
	check(err)
	hash := utils.ScriptHash(expectedData)
	data, err := DownloadXML(hash)
	check(err)
	if !bytes.Equal(data, expectedData) {
		t.Fatalf("Download data does not match file contents")
	}
}
