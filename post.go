package qbittorrentapi

import (
	"bytes"
)

func (q *QbittorrentClient) Post(endpoint, data string) (*Response, error) {
	postBody := bytes.NewBufferString(data)
	resp, err := q.Do("POST", endpoint, postBody)
	return resp, err
}
