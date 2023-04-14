package mal

import _ "embed"

//go:embed text/en/mal.md
var _mal string

//go:embed text/en/anime.md
var _anime string

//go:embed text/en/manga.md
var _manga string

//go:embed text/en/animelist.md
var _animeList string

//go:embed text/en/animeplan.md
var _animePlan string

//go:embed text/en/animedone.md
var _animeDone string

//go:embed text/en/mangalist.md
var _mangaList string

//go:embed text/en/mangaplan.md
var _mangaPlan string

//go:embed text/en/mangadone.md
var _mangaDone string
