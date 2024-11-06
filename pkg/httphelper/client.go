package httphelper

import (
	"context"
	"encoding/json"
	"net/http"
	neturl "net/url"
)

func Get[Response any](ctx context.Context, client http.Client, url string, headers map[string]string, params map[string]string) (*Response, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	values := u.Query()
	for key, value := range params {
		values.Add(key, value)
	}
	u.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var o Response
	if err := json.NewDecoder(res.Body).Decode(&o); err != nil {
		return nil, err
	}

	return &o, nil
}
