package md5

import (
	md5Pkg "crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5(str string) string {
	hash := md5Pkg.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// 对文件进行MD5加密（用于校验文件完整性）
func Md5File(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建MD5哈希器
	hash := md5Pkg.New()
	// 将文件内容写入哈希器
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 计算并返回哈希值
	return hex.EncodeToString(hash.Sum(nil)), nil
}