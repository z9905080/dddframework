package env_variable

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
)

var Setting Config

func init() {
	viper.SetConfigFile(getEnvPath())
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.Unmarshal(&Setting)
}

func getEnvPath() string {
	// .env 讀取
	gotenv.Load()
	env := os.Getenv("ENV")
	return "environment/" + env + ".json"
}

// LogLevel => 0.Debug 1.Info 2.Warn 3.Error 4.Fatal
// LogMode => # 0.Stdout 1.File
type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Debug    bool `json:"debug"`
	LogLevel int  `json:"log_level"`
	LogMode  int  `json:"log_mode"`
}
