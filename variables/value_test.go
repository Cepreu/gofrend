package variables

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	//	v1 := "012345"
	iv := IntegerValue{defaultValueImpl{false}, 1234}
	strIv, err := iv.convertToString()
	if err != nil {
		t.Error(err)
		return
	}
	if strIv != "1234" {
		t.Error("Expected 1234", "got ", strIv)
	}
}
func TestCompare(t *testing.T) {
	iv1 := IntegerValue{defaultValueImpl{false}, 1234}
	compArray := [...][3]int{{0, 1234, 0}, {1, 1234, 0}, {0, 2000, -1}, {0, 1000, 1}}
	for _, turple := range compArray {
		ivI := IntegerValue{defaultValueImpl{1 == turple[0]}, int64(turple[1])}
		res, _ := iv1.compareTo(&ivI)
		if res != turple[2] {
			t.Error("Expected ", turple[2], ", got ", res)
		}
	}
}
func TestConvertion(t *testing.T) {
	iv1 := IntegerValue{defaultValueImpl{false}, 1234}
	_, err := iv1.toDate()
	fmt.Println("OK: ", err)
	if err == nil {
		t.Error("Expected err, received OK")
	}
}
func TestTime(t *testing.T) {
	type timeTest struct {
		s   string
		v   int32
		err bool
	}
	timeArray := [...]timeTest{{"6:34pm", 18*60 + 34, true}, {"12:34", 12*60 + 34, true}, {"12:34PM", 12*60 + 34, true}, {"12:34 AM", 34, true}, {"00:34", 34, true}, {"00:34am", 0, false}, {"13:34 AM", 0, false}}
	for _, turple := range timeArray {
		tm, e := vuStringToMinutes(turple.s)
		if e == nil && turple.err {
			if tm != turple.v {
				t.Error(turple.s, "Expected ", turple.v, ", got ", tm)
			}
		} else if e == nil {
			t.Error(turple.s, "Expected error, got OK")
		}
	}
}
