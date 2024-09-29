package qbittorrentapi

import (
	"net/http"
)

type Response struct {
	StatusCode int
	Cookies    []*http.Cookie
	Body       []byte
}
