package phe

import "testing"

func requireNoErr(t testing.TB, err error) {
	if err != nil {
		t.Error(err)
	}
}

func requireErr(t testing.TB, err error) {
	if err == nil {
		t.Error("error expected")
	}
}

func requireEquals(t testing.TB, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("%v expected; %v provided", expected, actual)
	}
}

func requireTrue(t testing.TB, val bool) {
	if !val {
		t.Error("must be true")
	}
}
