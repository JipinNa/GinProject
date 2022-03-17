package server

type Method string

const (
	POST  Method = `POST`
	PATCH Method = `PATCH`
	GET   Method = `GET`
)

const (
	ListenAddr = `127.0.0.1:5544`
)
