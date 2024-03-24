package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	t.Log("Similar to fmt.Println() and concurrency safe")
	t.Fail()    // wil show a test case has failed in the results
	t.FailNow() // t.Fail() + safely exit without continuing
	t.Error("t.Log() + t.Fail()")
	t.Fatal("t.Log() + t.FailNow()")

}
