package main

import (
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	maximumAcceptableStatusCode = http.StatusBadRequest
)

var (
	statusNotHTML DeadLink = errors.New("not html")
	status400     DeadLink = errors.New("bad request")
	status404     DeadLink = errors.New("dead link")
	status500     DeadLink = errors.New("server error")
	status503     DeadLink = errors.New("service unavailable")
	status504     DeadLink = errors.New("gateway timeout")
	statusOther   DeadLink = errors.New("other error")
)

type DeadLink error

type Client struct {
	client *http.Client
}

// let timeout be provided from app call
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
	}
}

func (c *Client) Get(url string) ([]byte, DeadLink) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if deadLink := getDeadLink(resp.StatusCode); deadLink != nil {
		return nil, deadLink
	}

	body, _ := io.ReadAll(resp.Body)
	return body, nil
}

func getDeadLink(statusCode int) DeadLink {
	if statusCode < maximumAcceptableStatusCode {
		return nil
	}
	switch statusCode {
	case http.StatusBadRequest:
		return status400
	case http.StatusNotFound:
		return status404
	case http.StatusInternalServerError:
		return status500
	case http.StatusServiceUnavailable:
		return status503
	case http.StatusGatewayTimeout:
		return status504
	default:
		break
	}
	return statusOther
}
