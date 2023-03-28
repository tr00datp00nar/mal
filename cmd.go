package mal

import (
	"strconv"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `mal`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_mal),
	Description: help.D(_mal),

	Commands: []*Z.Cmd{
		doneCmd,
		listCmd,
		planCmd,
		help.Cmd,
	},
}

var listCmd = &Z.Cmd{
	Name:        `list`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_list),
	Description: help.D(_list),

	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		i, _ := strconv.Atoi("1")
		main(i)
		return nil
	},
}

var planCmd = &Z.Cmd{
	Name:        `plan`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_plan),
	Description: help.D(_plan),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		i, _ := strconv.Atoi("2")
		main(i)
		return nil
	},
}

var doneCmd = &Z.Cmd{
	Name:        `done`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_done),
	Description: help.D(_done),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		i, _ := strconv.Atoi("3")
		main(i)
		return nil
	},
}
