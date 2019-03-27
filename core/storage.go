package core

var storageDataPool = make(map[string]string)

func StoragePut(key, value string) {
	storageDataPool[key] = value
}

func StorageGet(key string) string {
	return storageDataPool[key]
}

func StrageDelete(key string) {
	delete(storageDataPool, key)
}
