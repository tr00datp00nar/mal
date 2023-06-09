package mal

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// TODO: Move from using viper for configuration to the rwxrob/conf package instead.
//		c := conf.C(Id: `bar`, Dir: `testdata`, File: `config.yaml`)
//		clientId := c.Query(.mal.client_id)
//		clientId := c.Query(.mal.client_secret)

var userHomeDir, _ = os.UserHomeDir()
var cacheDir = filepath.Join(userHomeDir, ".config/mal")
var cacheName = filepath.Join(cacheDir, "/mal-token-cache.txt")

type MalConfig struct {
	ClientId     string `yaml:"mal_client_id"`
	ClientSecret string `yaml:"mal_client_secret"`
}

func LoadConfigFile() (config MalConfig, err error) {
	viper.AddConfigPath(filepath.Clean("$HOME/.config/mal"))
	viper.SetConfigName("mal")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}
	err = viper.Unmarshal(&config)
	return
}

func authenticate(ctx context.Context, clientID, clientSecret, state string) (*http.Client, error) {
	// Prepare the oauth2 configuration with your application ID, secret, the
	// MyAnimeList authentication and token URLs as specified in:
	//
	// https://myanimelist.net/apiconfig/references/authorization
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	oauth2Token, err := loadCachedToken()
	if err == nil {
		refreshedToken, err := conf.TokenSource(ctx, oauth2Token).Token()
		if err == nil && (oauth2Token != refreshedToken) {
			fmt.Println("Caching refreshed oauth2 token...")
			if err := cacheToken(*refreshedToken); err != nil {
				return nil, fmt.Errorf("caching refreshed oauth2 token: %s", err)
			}
			return conf.Client(ctx, refreshedToken), nil
		}
		return conf.Client(ctx, oauth2Token), nil
	}

	// Generate a code verifier, a high-entropy cryptographic random string. It
	// will be set as the code_challenge in the authentication URL. It should
	// have a minimum length of 43 characters and a maximum length of 128
	// characters.
	const codeVerifierLength = 128
	codeVerifier, err := generateCodeVerifier(codeVerifierLength)
	if err != nil {
		return nil, fmt.Errorf("generating code verifier: %v", err)
	}

	// Produce the authentication URL where the user needs to be redirected and
	// allow your application to access their MyAnimeList data.
	authURL := conf.AuthCodeURL(state,
		oauth2.SetAuthURLParam("code_challenge", codeVerifier),
	)
	err = openBrowser(authURL)
	if err != nil {
		fmt.Println("Could not open browser.")
	}

	fmt.Printf("Your browser should open: %v\n", authURL)
	fmt.Print("After authenticating, copy the code from the browser URL and paste it here: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := scanner.Text()
	if scanner.Err() != nil {
		return nil, fmt.Errorf("reading code from terminal: %v", err)
	}

	// Exchange the authentication code for a token. MyAnimeList currently only
	// supports the plain code_challenge_method so to verify the string, just
	// make sure it is the same as the one you entered in the code_challenge.
	token, err := conf.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
	if err != nil {
		return nil, fmt.Errorf("exchanging code for token: %v", err)
	}
	fmt.Println("Authentication was successful. Caching oauth2 token...")
	if err := cacheToken(*token); err != nil {
		return nil, fmt.Errorf("caching oauth2 token: %s", err)
	}

	return conf.Client(ctx, token), nil
}

func cacheToken(token oauth2.Token) error {
	b, err := json.MarshalIndent(token, "", "   ")
	if err != nil {
		return fmt.Errorf("marshaling token %s: %v", token, err)
	}
	err = os.WriteFile(cacheName, b, 0644)
	if err != nil {
		return fmt.Errorf("writing token %s to file %q: %v", token, cacheName, err)
	}
	return nil
}

func loadCachedToken() (*oauth2.Token, error) {
	b, err := os.ReadFile(cacheName)
	if err != nil {
		return nil, fmt.Errorf("reading oauth2 token from cache file %q: %v", cacheName, err)
	}
	token := new(oauth2.Token)
	if err := json.Unmarshal(b, token); err != nil {
		return nil, fmt.Errorf("unmarshaling oauth2 token: %v", err)
	}
	return token, nil
}

func generateCodeVerifier(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstvuwxyz" +
		"0123456789-._~"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes), nil
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("openBrowser: unsupported operating system: %v", runtime.GOOS)
	}
}
