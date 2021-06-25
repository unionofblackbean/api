package jws

type JWS struct {
	Header  Header
	Payload Payload
}

func New() *JWS {
	return &JWS{
		Header:  NewHeader(),
		Payload: NewPayload(),
	}
}
