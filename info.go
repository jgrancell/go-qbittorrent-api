package qbittorrentapi

import (
	"encoding/json"

	"github.com/jgrancell/qbittorrent-api/schema"
)

func (q *QbittorrentClient) Info() ([]*schema.TorrentInfo, error) {
	// Checking to see if we need to perform a cookie login.
	err := q.DoLoginCheck()
	if err != nil {
		return nil, err
	}

	// Looking up the correct Qbittorrent endpoint for our version.
	endpoint, err := q.GetEndpoint("info")
	if err != nil {
		return nil, err
	}

	// Getting data from the Qbittorrent endpoint
	resp, err := q.Get(endpoint)
	if err != nil {
		return nil, err
	}

	// Parsing the received API information
	var info []*schema.TorrentInfo
	err = json.Unmarshal(resp.Body, &info)
	return info, err
}
