package config

import (
	"log"
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var configData []byte

//go:embed assets/beep.mp3
var soundData []byte

type Config struct {
	WinW   int        `json:"win_width"`
	WinH   int        `json:"win_height"`
	AlignX int        `json:"align_x"`
	AlignY int        `json:"align_y"`
	TextC  []float32  `json:"text_color"`
	BackC  []float32  `json:"back_color"`
	Sound  []byte     `json:"-"`
}

func Parse() (Config, error) {
	var cfg Config
	if err := json.Unmarshal(configData, &cfg); err != nil {
		log.Printf("Unmarshal config err: %v", err)
		return cfg, err
	}
	cfg.Sound = soundData
	return cfg, nil
}
