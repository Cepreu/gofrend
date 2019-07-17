package cloud

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Cepreu/gofrend/utils"
)

func checkNil(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func checkNotNil(err error, t *testing.T) {
	if err == nil {
		t.Fatalf("Expected error, got nil.")
	}
}

var fname1 = "test_files/comparison_test.five9ivr"

func TestUploadXML(t *testing.T) {
	data, err := ioutil.ReadFile(fname1)
	checkNil(err, t)
	err = UploadXML(data)
	checkNil(err, t)
}

func TestDownloadXML(t *testing.T) {
	TestUploadXML(t)

	expectedData, err := ioutil.ReadFile(fname1)
	checkNil(err, t)
	hash := utils.HashToString(expectedData)
	data, err := DownloadXML(hash)
	checkNil(err, t)
	if !bytes.Equal(data, expectedData) {
		t.Fatalf("Download data does not match file contents")
	}
}

func TestBadDownload(t *testing.T) {
	_, err := DownloadXML("fake_file_name")
	checkNotNil(err, t)
}
