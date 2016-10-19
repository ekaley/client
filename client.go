package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"

	"net/http"

	"github.com/RackHD/ipam/controllers/helpers"
	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/resources/factory"
	"github.com/hashicorp/go-cleanhttp"
)

// Client is used to configure the creation of a client
type Client struct {
	Address string
	Scheme  string
}

// NewClient returns a new client
func NewClient(address string) *Client {
	// bootstrap the config
	c := &Client{
		Address: address,
		Scheme:  "http",
	}
	return c
}

// SendResource is used to send the passed resource type
func (c *Client) SendResource(method, path string, in interfaces.Resource) (string, error) {

	body, err := encodeBody(in)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(method, c.Scheme+"://"+c.Address+path, body)
	if err != nil {
		return "", err
	}
	req.Header.Set(
		"Content-Type",
		mime.FormatMediaType(
			fmt.Sprintf("%s+%s", in.Type(), "json"),
			map[string]string{"version": in.Version()},
		),
	)

	client := cleanhttp.DefaultClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return "", errors.New(resp.Status)
	}

	return resp.Header.Get("Location"), nil
}

// ReceiveResource is used to receive the passed reasource type
func (c *Client) ReceiveResource(method, path, rType, rVersion string) (interfaces.Resource, error) {

	req, err := http.NewRequest(method, c.Scheme+"://"+c.Address+path, nil)

	req.Header.Set(
		"Content-Type",
		mime.FormatMediaType(
			fmt.Sprintf("%s+%s", rType, "json"),
			map[string]string{"version": rVersion},
		),
	)

	client := cleanhttp.DefaultClient()
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	mediaType, err := helpers.NewMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	resource, err := factory.Require(mediaType.Type, mediaType.Version)
	if err != nil {
		return nil, err
	}

	err = decodeBody(r, &resource)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

// SendReceiveResource is used to send a resource type and then
// upon success, fetch and recieve that resource type
func (c *Client) SendReceiveResource(methodSend, methodReceive, path string, in interfaces.Resource) (interfaces.Resource, error) {

	location, err := c.SendResource(methodSend, path, in)
	if err != nil {
		return nil, err
	}
	out, err := c.ReceiveResource(methodReceive, location, "", "")
	return out, err
}

// decodeBody is used to JSON decode a body
func decodeBody(resp *http.Response, out interface{}) error {
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// encodeBody is used to encode a request body
func encodeBody(obj interface{}) (io.Reader, error) {
	if obj == nil {
		return nil, nil
	}
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}
