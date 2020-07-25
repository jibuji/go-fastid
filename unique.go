package fastid

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var global int64 = 0

//Unique Not Strict Unique, but good enough
func Unique(salt ...interface{}) string {
	timePart := strconv.FormatInt(time.Now().UnixNano(), 36)[2:]
	s := fmt.Sprintf("%v", salt)
	s = digest(s)[:10]
	randomPart := strconv.FormatInt(rand.Int63n(time.Now().UnixNano()), 36)
	gPart := strconv.FormatInt(global, 36)
	global++
	return fmt.Sprintf("%s%s%s%s", timePart, s, gPart, randomPart)
}

func digest(str string) string {
	bs := []byte(str)
	m := md5.Sum(bs)
	hexDigest := hex.EncodeToString(m[:])
	return hexDigest
}
