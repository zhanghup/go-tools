package tools

import (
	"fmt"
	"io"
)

/*
	Filepath 获取Upload文件路径
	fieldId: UUID，例如：51e761c0-d4ff-478d-923a-14fb5b2bd0af
	ym: 例如202201
	contentType: image/jpeg
	fType: jpg
*/
func Filepath(fileId, ym, contentType, fType string) string {
	path := fmt.Sprintf("upload/%s/%s/%s.%s", ym, contentType, fileId, fType)
	data := []byte(path)
	for i := 0; len(data) < 150; i++ {
		data = append(data, 32)
	}
	return string(data)
}

/*
	FileUploadIO 文件上传
	reader: 文件流
	name: 文件名，例如1.jpg
	contentType: image/jpeg
	fType: jpg
*/
func FileUploadIO(reader io.Reader, name, contentType, salt string) (string, error) {
	//data, err := ioutil.ReadAll(reader)
	//if err != nil {
	//	return "", fmt.Errorf("[UploadIO] 读取文件失败【1】，error: %s", err.Error())
	//}
	//checksum, err := SHA256(append(data, []byte(salt)...))
	//if err != nil {
	//	return "", err
	//}
	return "", nil
}
