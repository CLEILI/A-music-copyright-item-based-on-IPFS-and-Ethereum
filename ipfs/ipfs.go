package ipfs

import (
	"fmt"
	"os"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

func Upload(path string) string {
	// 连接到本地IPFS节点，默认API地址是127.0.0.1:5001
	shell := ipfsapi.NewShell("localhost:5001")
	// 打开要上传的文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	// 上传文件到IPFS
	cid, err := shell.Add(file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cid
}
func Download(cid string, dpath string) string {
	shell := ipfsapi.NewShell("localhost:5001")
	err := shell.Get(cid, dpath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return dpath
}
