package models

import "github.com/nikaydo/grpc-contract/gen/video"

type Video struct {
	Title string `json:"title"`
	Vid   []byte `json:"Vid"`
	Token string `json:"Token"`
}

type User struct {
	Id           int    `json:"id,omitempty"`
	Login        string `json:"login"`
	Pass         string `json:"pass"`
	RefreshToken string `json:"refresh,omitempty"`
}

type Tokens struct {
	Token []string `json:"tokens"`
}

type VideoData struct {
	Token string `json:"token"`
	Uuid  string `json:"uuid,omitempty"`
	Name  string `json:"name,omitempty"`
}

type VideoList struct {
	Video []*video.SavedVideo `json:"title,omitempty"`
}
