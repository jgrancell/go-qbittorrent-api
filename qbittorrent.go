package qbittorrentapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	loginCookieName string = "SID"
	userAgent       string = "goQbittorrentClient/1.0"
)

type QbittorrentClient struct {
	FQDN       string
	AuthConfig *AuthConfig
	HttpConfig HttpConfig
	Version    string
}

type AuthConfig struct {
	Method   string
	Username string
	Password string
	SID      string
}

type HttpConfig struct {
	Timeout time.Duration
}

func (q *QbittorrentClient) Do(method, endpoint string, data io.Reader) (*Response, error) {
	client := &http.Client{Timeout: q.HttpConfig.Timeout}
	url := fmt.Sprintf("%s/%s", q.FQDN, strings.TrimPrefix(endpoint, "/"))
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Setting default headers
	req.Header.Add("Referer", q.FQDN)
	req.Header.Add("User-Agent", userAgent)

	if q.AuthConfig.SID != "" && q.AuthConfig.Method == "cookie" {
		req.AddCookie(&http.Cookie{Name: loginCookieName, Value: q.AuthConfig.SID})
	}

	if q.AuthConfig.Method == "basic" {
		req.SetBasicAuth(q.AuthConfig.Username, q.AuthConfig.Password)
	}

	if data != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	response := &Response{
		StatusCode: resp.StatusCode,
		Cookies:    resp.Cookies(),
		Body:       bodyBytes,
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return response, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	response.Body = bodyBytes
	return response, err
}
