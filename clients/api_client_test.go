package clients

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSetDefaultApiUrl(t *testing.T) {
	setUpViper()
	client := NewApiClient()
	assert.Equal(t, client.URL, "http://api.devportal.name" )
}

func TestOverrideDefaultApiUrl(t *testing.T) {
	os.Setenv("API_SERVICE_BASE_URL", "http://localhost:5000")
	setUpViper()
	client := NewApiClient()

	assert.Equal(t, client.URL, "http://localhost:5000")
}

func setUpViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}