package cfg

import "time"

//  Config - конфигурация из файла config.yaml
type Config struct {
	CBR CfgCBR `yaml:"cbr"`
}

// CfgCBR - конфигурация интеграции с апи Центробанка
type CfgCBR struct {
	Host    string        `yaml:"host"`
	Method  string        `yaml:"method"`
	Days    int           `yaml:"days"`
	Timeout time.Duration `yaml:"timeout"`
}
