package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(secret []byte, expiration time.Duration, username string) (string, error) {
	now := time.Now().Unix()
	// 获取当前时间
	exp := now + int64(expiration.Seconds())
	// 计算过期时间
	data := fmt.Sprintf("%d:%s", exp, username)
	// 将数据转换为字符串
	s := hmac.New(sha256.New, secret)
	// 使用HMAC-SHA256算法生成签名
	s.Write([]byte(data))
	signature := s.Sum(nil)
	// 生成Token
	token := data + ":" + base64.URLEncoding.EncodeToString(signature)
	return token, nil
}

func TokenExpired(secret []byte, token string) (bool, string, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 3 {
		return false, "", fmt.Errorf("invalid token format")
	}
	// 解析Token
	signature, err := base64.URLEncoding.DecodeString(parts[2])
	if err != nil {
		return false, "", err
	}
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(parts[0] + ":" + parts[1]))
	expected := h.Sum(nil)
	if !hmac.Equal(signature, expected) {
		return false, "", fmt.Errorf("invalid signature")
	}
	username := parts[1]
	if err != nil {
		return false, "", err
	}
	// 验证签名
	exp, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return false, "", err
	}
	// 获取时间
	now := time.Now().Unix()
	if exp > now {
		return true, username, nil
	}
	// 验证是否已过期
	return false, username, nil
}
