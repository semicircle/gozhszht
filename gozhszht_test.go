package gozhszht

import (
	"testing"
)

var (
	sample_sp = "今天天气真不错"
	sample_tr = "今天天氣真不錯"
)

func Test1(t *testing.T) {
	tr := ToTradition(sample_sp)
	if tr != sample_tr {
		t.Error("ToTradition", sample_tr, tr)
	}

	sp := ToSimple(tr)
	if sp != sample_sp {
		t.Error("ToSimple", sample_sp, sp)
	}
}
