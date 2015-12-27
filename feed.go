package main

import (
	_ "encoding/json"
)

type AMIFeed struct {
	ReleaseInfo  *ReleaseInfo `json:"release_info"`
	SAEast1      *AMIIDs      `json:"sa-east-1"`
	APNorthEast1 *AMIIDs      `json:"ap-norheast-1"`
	APSouthEast1 *AMIIDs      `json:"ap-southeast-1"`
	APSouthEast2 *AMIIDs      `json:"ap-southeast-2"`
	EUCentral1   *AMIIDs      `json:"eu-central-1"`
	EUWest1      *AMIIDs      `json:"eu-west-1"`
	USEast1      *AMIIDs      `json:"us-east-1"`
	USGovWest1   *AMIIDs      `json:"us-gov-west-1"`
	USWest1      *AMIIDs      `json:"us-west-1"`
	USWest2      *AMIIDs      `json:"us-west-2"`
}

type AMIIDs struct {
	HVM string `json:"hvm"`
	PV  string `json:"pv"`
}

type ReleaseInfo struct {
	Version     string `json:"version"`
	ReleaseDate string `json:"release_date"`
	Platform    string `json:"aws"`
}

func FeedURLOf(channel string) string {
	return "https://coreos.com/dist/aws/aws-" + channel + ".json"
}
