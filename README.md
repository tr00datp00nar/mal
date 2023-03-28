# :deciduous_tree: Mal

This is mal, a tool to interact with MyAnimeList.net. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

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

You will need to create a file `credentials.go` in the same directory as the source code with the following content:

```bash
package mal

const (
	defaultClientID     = "YOUR-CLIENT-ID"
	defaultClientSecret = "YOUR-CLIENT-SECRET"
)

```

## CAUTION
This branch is incomplete. Currently, it can only be used by cloning this repository to a location on your system and using `go install` to compile the binary with your credential information.

## Resources
To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z

## To-do

- [ ] Better way to manage credentials
- [ ] Create some kind of nice looking UI to display to returned API data.
- [ ] Better logging
- [ ] Better documentation
