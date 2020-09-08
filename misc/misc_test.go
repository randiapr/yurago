package misc

import (
	"reflect"
	"testing"
)

func TestPopArrayStr(t *testing.T) {
	sample := []string{"aa", "ab", "ac", "ad"}
	compare := []string{"aa", "ab", "ad"}

	out := PopArrayStr(sample, "ac")
	if !reflect.DeepEqual(out, compare) {
		t.Errorf("output %+v and compare %+v", out, sample)
	}

}
