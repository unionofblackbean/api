package jws

type Header map[string]interface{}

func NewHeader() Header {
	return make(Header)
}

func (h Header) Get(key string) interface{} {
	return h[key]
}

func (h Header) Set(key string, value interface{}) {
	h[key] = value
}

func (h Header) Remove(key string) {
	delete(h, key)
}

func (h Header) Exists(key string) bool {
	_, found := h[key]
	return found
}
