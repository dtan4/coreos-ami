package coreos

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AMIFeed represents a structure of CoreOS AMI JSON feed
type AMIFeed struct {
	ReleaseInfo map[string]string
	AMIs        map[string]map[string]string
}

func RetrieveAMIFeed(channel string) (*AMIFeed, error) {
	url := feedURL(channel)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feed map[string]map[string]string

	if err := json.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return nil, err
	}

	releaseInfo := feed["release_info"]
	delete(feed, "release_info")

	amiInfo := &AMIFeed{
		ReleaseInfo: releaseInfo,
		AMIs:        feed,
	}

	return amiInfo, nil
}

func feedURL(channel string) string {
	return fmt.Sprintf("https://coreos.com/dist/aws/aws-%s.json", channel)
}
