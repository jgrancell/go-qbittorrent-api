package qbittorrentapi

import (
	"fmt"
)

func (q *QbittorrentClient) DoLoginCheck() error {
	if q.AuthConfig.Method == "cookie" {
		if q.AuthConfig.SID == "" {
			return q.Login()
		}
	}
	return nil
}

func (q *QbittorrentClient) Login() error {
	// Looking up the correct Qbittorrent endpoint for our version.
	endpoint, err := q.GetEndpoint("login")
	if err != nil {
		return err
	}

	credentials := fmt.Sprintf("username=%s&password=%s", q.AuthConfig.Username, q.AuthConfig.Password)
	resp, err := q.Post(endpoint, credentials)
	if err != nil {
		return err
	}

	for _, cookie := range resp.Cookies {
		if cookie.Name == loginCookieName {
			q.AuthConfig.SID = cookie.Value
			return nil
		}
	}
	return fmt.Errorf("correct cookie with name %s not found", loginCookieName)
}
