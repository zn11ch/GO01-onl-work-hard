package internal

type Config struct {
	CheckLimit int `json:"CheckLimit"`
	Workers    struct {
		FileRead   int `json:"FileRead"`
		CountWords int `json:"CountWords"`
		Check      int `json:"Check"`
	} `json:"workers"`
}

func NewConfig() *Config {
	c := new(Config)
	c.CheckLimit = 10
	c.Workers.FileRead = 1
	c.Workers.CountWords = 1
	c.Workers.Check = 1
	return c
}
