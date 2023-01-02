package lsky

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 调用lsky图床的接口
func Upload(base64img string) string {
	url := "http://192.168.1.1:9080/api/v1/upload"
	method := "POST"
	filename := Base64ToJSFile(base64img)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filename)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return ""
	}
	defer file.Close()
	part1, _ := writer.CreateFormFile("file", filepath.Base(filename))
	io.Copy(part1, file)
	writer.Close()
	os.Remove(filename)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", "Bearer 1|Y7M1ROhk5hu5HZOwb4FqPfLzwJW5s955WCjjvRxt")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 从返回的json中获取图片的url
	re := regexp.MustCompile(`"url":\s*"([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		url := matches[1]
		return url
	}
	return ""
}

// 将base64编码转换为js的File类型
func Base64ToJSFile(base64img string) string {
	// 分别获取文件类型和文件数据
	typeinfo := strings.Split(base64img, ";")[0]
	base64img = strings.Split(base64img, ",")[1]
	data, err := base64.StdEncoding.DecodeString(base64img)
	if err != nil {
		fmt.Println("文件解码失败:", err)
	}
	// 生成一个随机文件名
	filename, _ := generateRandomFilename()
	// 解析文件类型
	// data:image/png;base64 => image/png
	typeinfo = strings.Split(typeinfo, ":")[1]
	hz := strings.Replace(typeinfo, "image/", "", 1)
	// 生成文件
	file, _ := os.Create(filename + "." + hz)
	defer file.Close()
	io.Copy(file, bytes.NewReader(data))
	// 写入文件
	return filename + "." + hz
}
func generateRandomFilename() (string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
