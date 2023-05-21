package model

// 获取DNS列表
func DnsList() []*Dns {
	return EtcdDao.DnsList()
}

// 获取DNS信息
func DnsGet(key string) (*Dns, error) {
	return EtcdDao.DnsGet(key)
}

// 添加DNS信息
func DnsAdd(key, value string) (bool, error) {
	return EtcdDao.DnsAdd(key, value)
}

// 编辑DNS信息
func DnsEdit(key, value string) error {
	return EtcdDao.DnsEdit(key, value)
}

// 删除DNS信息
func DnsDel(key string) error {
	return EtcdDao.DnsDel(key)
}

// 监听数据库
func WatchDBUpdate() {
	WatchEtcd()
}
