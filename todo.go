package main

type Effect struct {
	Id     int    `json:"id"`
	Brand  string `json:"brand"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Stereo bool   `json:"stereo"`
}

type Effects []Effect
