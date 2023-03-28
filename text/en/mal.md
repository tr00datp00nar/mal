Commands for interacting with MyAnimeList.net

Requires a configuration file to be located at `$HOME/.config/mal/mal.yaml` with the following lines:


    mal_client_id: "YOUR-CLIENT-ID"
    mal_client_secret: "YOUR-CLIENT-SECRET"

On first run, will launch your default browser and redirect to an oauth2 login screen for your myanimelist.net app. Token will be cached in `$HOME/.config/mal/mal-token-cache.txt`

Subsequent runs do not re-authenticate or re-cache the token.
