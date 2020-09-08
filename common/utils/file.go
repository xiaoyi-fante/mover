package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

//根据文件路径获取所有的markdown文件
func GetAllFiles(dir string) ([]string, error) {
	dirPath, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	sep := string(os.PathSeparator)

	for _, fi := range dirPath {
		if fi.IsDir() { //如果还是一个目录，则递归去遍历
			subFiles, err := GetAllFiles(dir + sep + fi.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			//过滤指定格式的文件
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dir+sep+fi.Name())
			}
		}
	}
	return files, nil
}
