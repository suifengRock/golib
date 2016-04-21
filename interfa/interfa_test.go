package interfa

import (
	"testing"
)

func TestString(t *testing.T) {
	var b interface{}
	b = "false"
	b, err := String(b)
	if nil != err {
		t.Error(err.Error())
	}

}

func TestBool(t *testing.T) {
	var b interface{}
	b = false
	b, err := Bool(b)
	if nil != err {
		t.Error(err.Error())
	}
	if b != false {
		t.Error("error")
	}
}

func TestBytes(t *testing.T) {
	var b interface{}
	b = "test_btye"
	b, err := Bytes(b)
	if nil != err {
		t.Error(err.Error())
	}
}

func TestInt(t *testing.T) {
	var data interface{}
	data = 1024
	i, err := Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

	data = float32(1024)
	i, err = Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

	data = float64(1024)
	i, err = Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

	data = int16(1024)
	i, err = Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

	data = int32(1024)
	i, err = Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

	data = int64(1024)
	i, err = Int(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != int(1024) {
		t.Error("error")
	}

}

func TestFloat(t *testing.T) {
	var data interface{}
	data = 1024
	i, err := Float64(data)
	if nil != err {
		t.Error(err.Error())
	}
	if i != float64(1024) {
		t.Error("error")
	}
}
