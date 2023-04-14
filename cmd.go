package mal

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `mal`,
	Usage:       `COMMAND`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mal),
	Description: help.D(_mal),

	Commands: []*Z.Cmd{
		animeCmd,
		mangaCmd,
		help.Cmd,
	},
}

// ─────────────────────────── Anime ───────────────────────────────

var animeCmd = &Z.Cmd{
	Name:        `anime`,
	Usage:       `COMMAND`,
	Version:     `v0.0.3`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_anime),
	Description: help.D(_anime),

	Commands: []*Z.Cmd{
		animeDoneCmd,
		animeListCmd,
		animePlanCmd,
		help.Cmd,
	},
}

var animeListCmd = &Z.Cmd{
	Name:        `list`,
	Usage:       `[help]`,
	Version:     `v0.0.3`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_animeList),
	Description: help.D(_animeList),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		runWatching()
		return nil
	},
}

var animePlanCmd = &Z.Cmd{
	Name:        `plan`,
	Usage:       `[help]`,
	Version:     `v0.0.3`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_animePlan),
	Description: help.D(_animePlan),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		runPlanToWatch()
		return nil
	},
}

var animeDoneCmd = &Z.Cmd{
	Name:        `done`,
	Usage:       `[help]`,
	Version:     `v0.0.3`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_animeDone),
	Description: help.D(_animeDone),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		runCompleted()
		return nil
	},
}

// ─────────────────────────── Manga ───────────────────────────────

var mangaCmd = &Z.Cmd{
	Name:        `manga`,
	Usage:       `COMMAND`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_manga),
	Description: help.D(_manga),

	Commands: []*Z.Cmd{
		mangaListCmd,
		mangaPlanCmd,
		mangaDoneCmd,
		help.Cmd,
	},
}

var mangaListCmd = &Z.Cmd{
	Name:        `list`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mangaList),
	Description: help.D(_mangaList),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		runMangaList()
		return nil
	},
}

var mangaPlanCmd = &Z.Cmd{
	Name:        `plan`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mangaPlan),
	Description: help.D(_mangaPlan),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		runMangaListPlan()
		return nil
	},
}

var mangaDoneCmd = &Z.Cmd{
	Name:        `done`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mangaDone),
	Description: help.D(_mangaDone),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		runMangaListDone()
		return nil
	},
}
