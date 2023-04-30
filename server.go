package mal

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/rl404/nagato"
	"github.com/rl404/nagato/mal"
	"github.com/spf13/viper"
)

var (
	userHomeDir, _ = os.UserHomeDir()
	cacheDir       = filepath.Join(userHomeDir, ".config/mal")
	cacheName      = filepath.Join(cacheDir, "/mal-token-cache.txt")

	nagatoClient *nagato.Client
	config       MalConfig

	clientID     = config.ClientId
	clientSecret = config.ClientSecret
	state        = ""
)

type MalConfig struct {
	ClientId     string `yaml:"mal.client_id"`
	ClientSecret string `yaml:"mal.client_secret"`
}

func LoadConfigFile() (config MalConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Clean("$HOME/.config/c"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}
	err = viper.Unmarshal(&config)
	return
}

func clientPublic() {
	nagatoClient = nagato.New(clientID)
}

func clientOauth2() {
	nagatoClient = nagato.New(clientID)

	// Generating token should be done outside your project flow
	// and should be done one time only.
	if err := generateToken(context.Background(), clientID, clientSecret, state); err != nil {
		panic(err)
	}

	cachedToken, err := loadCachedToken()
	if err != nil {
		panic(err)
	}

	malWithOauth2, err := mal.NewWithOauth2(mal.Oauth2Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		State:        state,
		AccessToken:  cachedToken.AccessToken,
		RefreshToken: cachedToken.RefreshToken,
		Expiry:       cachedToken.Expiry,
	})
	if err != nil {
		panic(err)
	}

	nagatoClient.SetMalClient(malWithOauth2)
}

func toJson(name string, d interface{}) {
	jsonData, _ := json.MarshalIndent(d, "", " ")

	f, _ := os.Create(name + ".json")
	defer f.Close()

	f.Write(jsonData)
}
