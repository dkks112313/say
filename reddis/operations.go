package reddis

func CreateDataStorage() *Data {
	return &Data{
		make(map[string]string),
	}
}

func AddIntoStorage(d *Data, key, value string) {
	d.set(key, value)
}

func GetFromStorage(d *Data, key string) string {
	return d.get(key)
}
