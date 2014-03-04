package experiment

import (
	"testing"
)

func TestSayWithNumber(t *testing.T) {
	if Say("number") != "number123" {
		t.Fail()
	}
}
