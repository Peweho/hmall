package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 生成长度为size的随机字符串，每个字符为[0,10)的整数
func RandomNumeric(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		panic("{ size : " + strconv.Itoa(size) + " } must be more than 0 ")
	}
	value := ""
	for index := 0; index < size; index++ {
		value += strconv.Itoa(r.Intn(10))
	}

	return value
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}

// MD5加密
func Md5Password(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// 根据jwt获得当前用户的id
func GetUsr(ctx context.Context, jwtKey string) (int, error) {
	empIdJson := ctx.Value(jwtKey).(json.Number) //获得当前用户Id
	return strconv.Atoi(empIdJson.String())
}

// 构建缓存key
func CacheKey(prefix string, val string) string {
	return fmt.Sprintf("%s#%s", prefix, val)
}
