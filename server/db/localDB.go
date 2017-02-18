package db

type LocalDB map[string]interface{}

func (l LocalDB) Put(key string, value interface{}) error {
	l[key] = value
	return nil
}

func (l LocalDB) Get(key string) (interface{}, error) {
	return l[key], nil
}
