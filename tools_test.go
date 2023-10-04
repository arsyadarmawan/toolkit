package toolkit_test

import (
	"github.com/arsyadarmawan/toolkit"
	"testing"
)

func TestName(t *testing.T) {
	tools := toolkit.Tools{}
	randString := tools.RandomString(10)
	if len(randString) != 10 {
		t.Error("wrong length number")
	}
}
