package qbittorrentapi

import "fmt"

var (
	ApiEndpoints = map[string]map[string]string{
		"v2": {
			"login": "/api/v2/auth/login",
			"info":  "/api/v2/torrents/info",
		},
	}
)

func (q *QbittorrentClient) GetEndpoint(name string) (string, error) {
	for version, endpoints := range ApiEndpoints {
		if version == q.Version {
			for endpoint, value := range endpoints {
				if endpoint == name {
					return value, nil
				}
			}
			return "", fmt.Errorf("the api endpoint %s is not supported in version %s", name, q.Version)
		}
	}
	return "", fmt.Errorf("the qbittorrent api version %s is not supported", q.Version)
}
