package fastid

import (
	"log"
	"testing"

	"gotest.tools/assert"
)

func TestUnique(t *testing.T) {
	u1 := Unique(1)
	u11 := Unique(1)
	assert.Assert(t, u1 != u11, "u1=%v u11= %v", u1, u11)
	u2 := Unique(2.3)
	u3 := Unique(2.3)
	assert.Assert(t, u2 != u3, "u2=%v u3= %v", u2, u3)
	u4 := Unique("")
	u5 := Unique("")
	assert.Assert(t, u4 != u5, "u4=%v u5= %v", u4, u5)
	u6 := Unique(nil)
	u7 := Unique(nil)
	assert.Assert(t, u6 != u7, "u6=%v u7= %v", u6, u7)
	u8 := Unique("", 0)
	u9 := Unique("", 0)
	assert.Assert(t, u8 != u9, "u8=%v u9= %v", u8, u9)
}

func TestCryptoUnique(t *testing.T) {
	u1 := CUnique()
	u11 := CUnique()
	log.Printf("u1=%s, u11=%s", u1, u11)
	assert.Assert(t, u1 != u11, "u1=%v u11= %v", u1, u11)
	u2 := CUnique()
	u3 := CUnique()
	assert.Assert(t, u2 != u3, "u2=%v u3= %v", u2, u3)
	u4 := CUnique()
	u5 := CUnique()
	assert.Assert(t, u4 != u5, "u4=%v u5= %v", u4, u5)
	u6 := CUnique()
	u7 := CUnique()
	assert.Assert(t, u6 != u7, "u6=%v u7= %v", u6, u7)
	u8 := CUnique()
	u9 := CUnique()
	assert.Assert(t, u8 != u9, "u8=%v u9= %v", u8, u9)
}
