package main

type FileReq struct {
	Raw     string `json:"raw"`
	Content string `json:"content"`
}

type DecodeResult struct {
	Text  string `json:"text"`
	Words []Word `json:"words"`
}

type Word struct {
	Text     string  `json:"text"`
	Score    float32 `json:"score"`
	Position Pos     `json:"position"`
}

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
