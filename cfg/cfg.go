package cfg

import "time"

// CfgCBR - конфигурация интеграции с апи Центробанка
type Config struct {
	CBR CfgCBR `yaml:"cbr"`
}

type CfgCBR struct {
	Host    string        `yaml:"host"`
	Method  string        `yaml:"method"`
	Days    int           `yaml:"days"`
	Timeout time.Duration `yaml:"timeout"`
}
