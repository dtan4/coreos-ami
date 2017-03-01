package coreos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"text/tabwriter"
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

func (a *AMIFeed) Tabularize() string {
	buf := new(bytes.Buffer)

	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', 0)

	for region, amis := range a.AMIs {
		for amiType, amiID := range amis {
			fmt.Fprintln(w, strings.Join([]string{region, amiType, amiID}, "\t"))
		}
	}

	w.Flush()

	lines := strings.Split(buf.String(), "\n")
	sort.Strings(lines[0 : len(lines)-1])

	return strings.Join(lines, "\n")
}
