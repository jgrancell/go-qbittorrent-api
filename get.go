package qbittorrentapi

func (q *QbittorrentClient) Get(endpoint string) (*Response, error) {
	resp, err := q.Do("GET", endpoint, nil)
	return resp, err

}
