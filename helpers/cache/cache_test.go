package cache

import "testing"

func TestGetCacheInstance(t *testing.T) {
	bm := GetCacheInstance()
	err = bm.Incr("article")
	if err != nil {
		t.Error("出错了!")
	}
}
