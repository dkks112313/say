package reddis

type Data struct {
	data map[string]string
}

func (d *Data) set(key, value string) {
	d.data[key] = value
}

func (d *Data) get(key string) string {
	return d.data[key]
}
