package jws

type Payload map[string]interface{}

func NewPayload() Payload {
	return make(Payload)
}

func (p Payload) Get(key string) interface{} {
	return p[key]
}

func (p Payload) Set(key string, value string) {
	p[key] = value
}

func (p Payload) Remove(key string) {
	delete(p, key)
}

func (p Payload) Exists(key string) bool {
	_, found := p[key]
	return found
}
