package html_test

import (
	"dns/html"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	fs := html.GetHtmlFS()
	file, err := fs.Open("dns.html")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		t.Fail()
		os.Exit(1)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		t.Fail()
		os.Exit(1)
	}
	fmt.Printf("data: %v\n", string(data))
}
