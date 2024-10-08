package model

type Netflix struct {
	Movie   string `json: "movie"`
	Watched bool   `json:"watched"`
}
