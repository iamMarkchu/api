package jwt

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	GetToken("10086")
}

func TestCheckToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAwODYiLCJleHAiOjE1NTc5Mjg5NzksImlzcyI6InRlc3QiLCJuYmYiOjE1NTc5MjcxNzl9.nXqgUjpVmcMNDu8_45LC3-J0xQfqxP_7oNd1KdyfTAc"
	UserId, correct := CheckToken(token)
	if UserId != "10086" && !correct {
		t.Error("出现错误!")
	}
}
