package peep

import (
	"reflect"
	"testing"
	"time"
)

func TestParsePsStat(t *testing.T) {
	input1 := "USER                      STARTED COMMAND\n"
	out1, _ := parsePsStat(input1)
	if out1 != nil {
		t.Errorf("out1 should be nil but: %#v", out1)
	}

	input2 := `USER                      STARTED COMMAND
root     Thu Feb 18 11:12:38 2016 /sbin/init  hoge
`
	expect := psStat{
		User:    "root",
		Command: "/sbin/init  hoge",
		StartAt: time.Date(2016, time.February, 18, 11, 12, 38, 0, time.Local),
	}
	out2, _ := parsePsStat(input2)
	if !reflect.DeepEqual(*out2, expect) {
		t.Errorf("something went wrong:\n   out: %+v\nexpect: %+v", *out2, expect)
	}
}
