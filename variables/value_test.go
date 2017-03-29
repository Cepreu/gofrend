package variables

import (
	"testing"
)

func TestInteger(t *testing.T) {
	//	v1 := "012345"
	iv := IntegerValue{false, 1234}
	strIv, err := iv.convertToString()
	if err != nil {
		t.Error(err)
		return
	}
	if strIv != "1234" {
		t.Error("Expected 1234", "got ", strIv)
	}
}
func TestCompare(t  *testing.T) {
	iv1 := IntegerValue{false, 1234}
	compArray := [...][3]int {{0, 1234, 0}, {1, 1234, 0}, {0, 2000, -1}, {0, 1000, 1}}
	for _, turple := range compArray {
		ivI := IntegerValue{1==turple[0], int64(turple[1])}
		res, _ := iv1.compareTo(&ivI)
		if res != turple[2] {
			t.Error("Expected ", turple[2],  ", got ", res)
		}
	}
}
