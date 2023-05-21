//go:build leveldb
// +build leveldb

package model

import (
	"dns/g"
	"dns/util"
	"encoding/json"
	"strings"

	leveldb1 "github.com/jeffcail/leveldb"
)

var (
	LevelDB *leveldb1.LevelDB
)

// 校验是否可以连接
func OninitCheck() {
	var err error
	LevelDB, err = leveldb1.CreateLevelDB("./dns_data")
	if err != nil {
		panic(err)
	}
}

func LevelDBkey2Host(key, value string) *Dns {
	temp := reg.ReplaceAllString(key, "")
	temp = strings.Replace(temp, g.DBKeyPath, "", 1)
	list := strings.Split(temp, "/")
	util.Reverse(list)
	aaa := A{}
	json.Unmarshal([]byte(value), &aaa)
	return &Dns{Origin: strings.Join(list, "."), NameServer: aaa.Host, TTL: aaa.TTL, Key: key, Value: value}
}
