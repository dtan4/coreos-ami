package coreos

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRetrieveAMIFeed(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	body, err := ioutil.ReadFile(filepath.Join("..", "testdata", "aws-beta.json"))
	if err != nil {
		t.Fatalf("Cannot open testdata: %s", err)
	}

	httpmock.RegisterResponder("GET", "https://coreos.com/dist/aws/aws-beta.json",
		httpmock.NewBytesResponder(200, body))

	amiFeed, err := RetrieveAMIFeed("beta")
	if err != nil {
		t.Errorf("Error should not be raised: %s", err)
	}

	expectedReleaseInfo := map[string]string{
		"version":      "1298.4.0",
		"release_date": "2017-02-23 04:40:00 +0000",
		"platform":     "aws",
	}

	if !reflect.DeepEqual(amiFeed.ReleaseInfo, expectedReleaseInfo) {
		t.Errorf("ReleaseInfo does not match. expected: %#v, got: %#v", expectedReleaseInfo, amiFeed.ReleaseInfo)
	}

	expectedAMIs := map[string]map[string]string{
		"eu-central-1": map[string]string{
			"hvm": "ami-273df748",
			"pv":  "ami-4233f92d",
		},
		"ap-northeast-1": map[string]string{
			"hvm": "ami-91f3bbf6",
			"pv":  "ami-45f3bb22",
		},
		"us-gov-west-1": map[string]string{
			"hvm": "ami-0447fd65",
			"pv":  "ami-ac46fccd",
		},
		"ap-northeast-2": map[string]string{
			"hvm": "ami-4c16c622",
			"pv":  "ami-1614c478",
		},
		"ca-central-1": map[string]string{
			"hvm": "ami-a17ec3c5",
			"pv":  "ami-1278c576",
		},
		"ap-south-1": map[string]string{
			"hvm": "ami-4a334325",
			"pv":  "ami-963343f9",
		},
		"sa-east-1": map[string]string{
			"hvm": "ami-9c1375f0",
			"pv":  "ami-14177178",
		},
		"ap-southeast-2": map[string]string{
			"hvm": "ami-76adad15",
			"pv":  "ami-4faeae2c",
		},
		"ap-southeast-1": map[string]string{
			"hvm": "ami-8dcb7aee",
			"pv":  "ami-b6cc7dd5",
		},
		"us-east-1": map[string]string{
			"hvm": "ami-eba777fd",
			"pv":  "ami-84bf6f92",
		},
		"us-east-2": map[string]string{
			"hvm": "ami-faa2879f",
			"pv":  "ami-2aad884f",
		},
		"us-west-2": map[string]string{
			"hvm": "ami-ec4ccc8c",
			"pv":  "ami-ef4ccc8f",
		},
		"us-west-1": map[string]string{
			"hvm": "ami-3e1d435e",
			"pv":  "ami-551b4535",
		},
		"eu-west-1": map[string]string{
			"hvm": "ami-fcf1dc9a",
			"pv":  "ami-3cf0dd5a",
		},
		"eu-west-2": map[string]string{
			"hvm": "ami-742b3e10",
			"pv":  "ami-062b3e62",
		},
	}

	if !reflect.DeepEqual(amiFeed.AMIs, expectedAMIs) {
		t.Errorf("AMIs does not match. expected: %#v, got: %#v", expectedAMIs, amiFeed.AMIs)
	}
}
