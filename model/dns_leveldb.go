//go:build leveldb
// +build leveldb

package model

import (
	"dns/g"
	"fmt"
)

// 获取DNS列表
func DnsList() []*Dns {
	result, found := g.Mycache.Get(g.Cache_Name)
	if found {
		return result.([]*Dns)
	} else {
		mymap := make(map[string]string)
		iter := LevelDB.SelectAll()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			fmt.Printf("k: %s\n", k)
			fmt.Printf("v: %s\n", v)
			mymap[string(k)] = string(v)
		}
		iter.Release()
		err := iter.Error()
		if err != nil {
			fmt.Printf("查询LevelDB错误: %v\n", err)
			return nil
		}
		result := make([]*Dns, 0, len(mymap))
		for k, v := range mymap {
			result = append(result, LevelDBkey2Host(k, v))
		}
		return result
	}
}

// 获取DNS信息
func DnsGet(key string) (*Dns, error) {
	value, err := LevelDB.Get(key)
	if err != nil {
		return nil, err
	}
	return LevelDBkey2Host(key, string(value)), nil
}

// 添加DNS信息
func DnsAdd(key, value string) (bool, error) {
	err := LevelDB.Put(key, value)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 编辑DNS信息
func DnsEdit(key, value string) error {
	LevelDB.Delete(key)
	return LevelDB.Put(key, value)
}

// 删除DNS信息
func DnsDel(key string) error {
	return LevelDB.Delete(key)
}

// 监听数据库
func WatchDBUpdate() {
	// TODO:待实现
	fmt.Println("不支持的功能")
}
