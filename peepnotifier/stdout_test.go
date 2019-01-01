package peepnotifier

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestStdout_run(t *testing.T) {
	st := &stdout{}
	re := &result{
		User:    "Songmu",
		Command: `perl -E 'say "Hello"'`,
		Host:    "localhost",
		StartAt: time.Now().UTC(),
		EndAt:   time.Now().UTC(),
	}
	var (
		buf  = &bytes.Buffer{}
		orig = outStream
	)
	outStream = buf
	defer func() { outStream = orig }()

	st.run(re, nil)
	out := &result{}
	err := json.NewDecoder(buf).Decode(out)
	if err != nil {
		t.Errorf("err should be nil but: %s", err)
	}

	if !reflect.DeepEqual(re, out) {
		t.Errorf("something went wrong")
	}
}
