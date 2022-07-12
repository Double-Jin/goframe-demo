package utils

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	r "math/rand"
	"time"
)

// 字符类
var Charset = new(charset)

type charset struct{}

func (util *charset) GetMapKeysByString(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func (util *charset) Md5ToString(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}


func (util *charset) RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	}
	var bytes = make([]byte, n)
	var randBy bool
	r.Seed(time.Now().UnixNano())
	if num, err := rand.Read(bytes); num != n || err != nil {
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}


func (util *charset) GetStack(err error) []string {
	stackList := gstr.Split(gerror.Stack(err), "\n")
	for i := 0; i < len(stackList); i++ {
		stackList[i] = gstr.Replace(stackList[i], "\t", "--> ")
	}

	return stackList
}

func (util *charset) IsExists(elems interface{}, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}

	return false
}
