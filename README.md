# :deciduous_tree: Mal

This is mal, a tool to interact with MyAnimeList.net. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

This branch is still very much a work in progress. PRs welcome.

If you just want to try it out, grab the release binary with curl and put into your PATH:

```bash
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-linux-amd64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-darwin-amd64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-darwin-arm64 -o ~/.local/bin/mal
curl -L https://github.com/tr00datp00nar/mal/releases/latest/download/mal-windows-amd64 -o ~/.local/bin/mal
```
To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see my personal Bonzai stateful command tree: https://github.com/tr00datp00nar/c

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z

## To-do

- [ ] Create some kind of nice looking UI to display to returned API data.
- [ ] Better logging
- [ ] Better documentation
