package lib

type HttpMethod int

const (
	HTTP_GET HttpMethod  = iota
	HTTP_POST
	HTTP_DELETE
	HTTP_PUT
)
