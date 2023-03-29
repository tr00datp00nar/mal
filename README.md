# :deciduous_tree: Mal

This is `mal`, a tool to interact with MyAnimeList.net. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

This branch is still very much a work in progress. PRs welcome.

## Installation

This branch can be composed into any Bonzai stateful command tree.

If you just want to try it out, grab the release binary with curl and put into your PATH:

```bash
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-linux-amd64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-darwin-amd64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-darwin-arm64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-windows-amd64 -o ~/.local/bin/mal
```

## Authentication
In order to create a client ID and secret for your application:

1. Navigate to https://myanimelist.net/apiconfig or go to your MyAnimeList profile, click Edit Profile and select the API tab on the far right.
2. Click Create ID and submit the form with your application details.

You will need to create a config file `$HOME/.config/mal/mal.yaml` with the following content:

```yaml
client_id: "YOUR-CLIENT-ID"
client_secret: "YOUR-CLIENT-SECRET"
```
This config location is currently hardcoded. PRs welcome to place configuration elsewhere.

## Resources
To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z

## To-do

- [ ] Allow user to define custom configuration and token cache locations.
- [ ] Allow user to change filters/sorting.
- [ ] Create some kind of nice looking UI to display to returned API data.
