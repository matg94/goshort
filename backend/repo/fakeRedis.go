package repo

type FakeRedis struct {
	store map[string]string
}

func (r *FakeRedis) GET(key string) string {
	return r.store[key]
}

func (r *FakeRedis) SET(key string, value string) {
	r.store[key] = value
}
