package html

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

//go:embed *
var HtmlFile embed.FS

func GetHtmlFS() fs.FS {
	sfs, err := fs.Sub(HtmlFile, "views")
	if err != nil {
		panic(err)
	}
	return sfs
}

// 加载HTTP HTML FS
func GetHttpHtmlFS() http.FileSystem {
	return http.FS(GetHtmlFS())
}

func WriteHTMLTemplate() {
	htmlfs := GetHtmlFS()
	filenames := []string{"dns.html", "dns2.html", "edit.html", "footer.html", "header.html", "index.html", "login.html"}
	for _, v := range filenames {
		fmt.Printf("加载文件: %v\n", v)
		f, err := htmlfs.Open(v)
		if err != nil {
			fmt.Printf("加载HTML文件错误: %v\n", err)
			continue
		}
		data, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("读取HTML文件错误: %v\n", err)
			continue
		}
		os.MkdirAll(os.TempDir()+"/views", 0755)
		fd, _ := os.OpenFile(os.TempDir()+"/views/"+v, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		_, err = fd.Write(data)
		fd.Close()
		if err != nil {
			fmt.Printf("写入HTML文件错误: %v\n", err)
			continue
		}

	}

}
