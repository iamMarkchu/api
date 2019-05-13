package helpers

import (
	"testing"
)

func TestCheckEmail(t *testing.T) {
	email1 := "18627032049@163.com"
	matched, _ := CheckEmail(email1)
	if matched != true {
		t.Errorf("CheckEmail(%s) = (%t); expected (%t)", email1, matched, true)
	}

	email2 := "1862$7032049@163.com"
	matched, _ = CheckEmail(email2)
	if matched != false {
		t.Errorf("CheckEmail(%s) = (%t); expected (%t)", email2, matched, false)
	}
}
