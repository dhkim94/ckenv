package ckenv

import (
	"testing"
)

func TestSample(t *testing.T) {
	if !Init("./sample.properties") {
		t.Fatal("fail initialize ckenv")
	}

	str1 := GetValue("daemon.name")
	if str1 == "" {
		t.Fatal("fail read property value")
	}

	// 없는 값은 nil 을 리턴 해야 한다.
	str2 := GetValue("no.value")

	if str2 != "" {
		t.Fatal("fail not found valud")
	}
}