package config

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	Etcd_url   = make([]string, 0)
	DBKeyPath  = "/coredns"
	Exit       = make(chan struct{}, 1)
	Mycache    = cache.New(60*time.Minute, 60*time.Minute)
	Cache_Name = "dns"
)
