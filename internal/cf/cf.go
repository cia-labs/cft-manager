package cf

import (
	"io"
	"io/ioutil"
	"net/http"
)

type CF struct {
	apiKey string
	client *http.Client
}

func NewCF(apiKey, email string) *CF {
	return &CF{
		apiKey: apiKey,
		// client object with Authorization header set to Bearer apiKey
		client: &http.Client{},
	}
}

func (c *CF) apiCall(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	return c.client.Do(req)
}

func (c *CF) GetTunnelList() (string, error) {
	// implementation to get tunnel list
	apiResonse, apiErr := c.apiCall("GET", "", nil)
	if apiErr != nil {
		return "", apiErr
	}
	defer apiResonse.Body.Close()
	body, err := ioutil.ReadAll(apiResonse.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// func (c *CF) GetZoneID(zoneName string) (string, error) {
// 	// implementation to get zone ID
// }

// func (c *CF) CreateDNSRecord(zoneID, recordType, name, content string, ttl int) error {
// 	// implementation to create DNS record
// }

// func (c *CF) UpdateDNSRecord(zoneID, recordID, recordType, name, content string, ttl int) error {
// 	// implementation to update DNS record
// }

// func (c *CF) DeleteDNSRecord(zoneID, recordID string) error {
// 	// implementation to delete DNS record
// }
