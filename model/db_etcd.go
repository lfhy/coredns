package model

import (
	"dns/config"
	"dns/util"
	"encoding/json"

	client "go.etcd.io/etcd/client/v3"

	//"dns/controller"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/patrickmn/go-cache"
)

var (
	kapi    client.KV
	EtcdDao = &etcddao{}
)

// 校验是否可以连接
func OninitCheck() {
	c, err := client.New(client.Config{
		Endpoints: config.Etcd_url,
	})
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	kapi = client.NewKV(c)
	_, err = kapi.Get(context.Background(), config.DBKeyPath, nil)
	//fmt.Println(rep.Node.Value)
	//_, err = kapi.Set(context.Background(),ippath, "0",&client.SetOptions{PrevExist:client.PrevExist,PrevValue:"0",Dir:false})
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

type etcddao struct {
}

func (this *etcddao) DnsList() []*Dns {
	result, found := config.Mycache.Get(config.Cache_Name)
	if found {
		return result.([]*Dns)
	} else {
		mymap, err := etcdList()
		if err != nil {
			return nil
		}
		result := make([]*Dns, 0, len(mymap))
		for k, v := range mymap {
			result = append(result, Etcdkey2Host(k, v))
		}
		return result
	}
}
func etcdALL() []*Dns {
	mymap, err := etcdList()
	if err != nil {
		return nil
	}
	result := make([]*Dns, 0, len(mymap))
	for k, v := range mymap {
		result = append(result, Etcdkey2Host(k, v))
	}
	config.Mycache.Set(config.Cache_Name, result, cache.DefaultExpiration)
	return result
}
func (this *etcddao) DnsAdd(key, value string) (bool, error) {
	return etcdAdd(key, value)
}
func (this *etcddao) DnsDel(key string) error {
	return etcdDel(key)
}
func (this *etcddao) DnsEdit(key, value string) error {
	return etcdEdit(key, value)
}

func (this *etcddao) DnsGet(key string) (*Dns, error) {
	resp, err := etcdGet(key)
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		return Etcdkey2Host(key, string(kv.Value)), nil
	}
	return nil, fmt.Errorf("Not Found")
}

// 获取KV Map
func etcdGetmap(resp *client.GetResponse, mymap map[string]string) {
	for _, kv := range resp.Kvs {
		mymap[string(kv.Key)] = string(kv.Value)
	}
}

func etcdList() (map[string]string, error) {
	node, err := etcdGet(config.DBKeyPath)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string, 10000)
	etcdGetmap(node, result)
	return result, nil
}

// 获取Key的返回
func etcdGet(key string) (*client.GetResponse, error) {
	return kapi.Get(context.Background(), key, nil)
}

func etcdAdd(key, value string) (bool, error) {
	keylist := strings.Split(key, ".")
	util.Reverse(keylist)
	prekey := strings.Join(keylist, "/")
	if !strings.HasPrefix(prekey, "/") {
		prekey = "/" + prekey
	}
	key = config.DBKeyPath + prekey
	_, err := kapi.Put(context.Background(), key, value, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
func etcdDel(key string) error {
	fmt.Println(key)
	_, err := kapi.Delete(context.Background(), key, nil)
	if err != nil {
		fmt.Println(key, err)
		return err
	}
	return nil
}
func etcdEdit(key, value string) error {
	_, err := kapi.Put(context.Background(), key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func WatchEtcd() {
	// TODO: 下次调
	fmt.Println("不支持的功能")
	// watcher := kapi.Watcher(config.DBKeyPath, &client.WatcherOptions{Recursive: true})
	// fmt.Println(122222)
	// for {
	// 	select {
	// 	case <-config.Exit:
	// 		break
	// 	default:

	// 	}
	// 	res, err := watcher.Next(context.Background())
	// 	if err != nil {
	// 		continue
	// 	}
	// 	if res.Action == "expire" {
	// 		continue
	// 	} else if res.Action == "set" || res.Action == "update" || res.Action == "create" || res.Action == "delete" {
	// 		fmt.Println(res.Action)
	// 		result := etcdALL()
	// 		if result != nil {
	// 			NewMessage <- result
	// 		}
	// 	}

	// }
}

func Etcdkey2Host(key, value string) *Dns {
	temp := reg.ReplaceAllString(key, "")
	temp = strings.Replace(temp, config.DBKeyPath, "", 1)
	list := strings.Split(temp, "/")
	util.Reverse(list)
	aaa := A{}
	json.Unmarshal([]byte(value), &aaa)
	return &Dns{Origin: strings.Join(list, "."), NameServer: aaa.Host, TTL: aaa.TTL, Key: key, Value: value}
}
