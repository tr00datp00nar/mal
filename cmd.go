package mal

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/log"
	"github.com/ktr0731/go-fuzzyfinder"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/to"

	"github.com/tr00datp00nar/animerecommender"
	"github.com/tr00datp00nar/fn"
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
		animerecommender.Cmd,
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

	Call: func(x *Z.Cmd, args ...string) error {
		clientPublic()
		clientOauth2()

		anime := GetUserAnimeListWatching()

		idx, err := fuzzyfinder.FindMulti(
			anime,
			func(i int) string {
				title := anime[i].Anime.AlternativeTitles.English
				if title == "" {
					title = anime[i].Anime.Title
				}
				return title
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := anime[i].Anime.Title
				engTitle := anime[i].Anime.AlternativeTitles.English
				if engTitle == "" {
					title = origTitle
				} else {
					title = engTitle
				}

				var broadcast string
				if anime[i].Anime.EndDate.Month == 0 {
					if len(anime[i].Anime.Broadcast.DayOfTheWeek) == 0 {
						broadcast = "Airs on: Unknown"
					} else {
						broadcast = "Airs on: " + strings.Title(to.String(anime[i].Anime.Broadcast.DayOfTheWeek))
					}
				} else {
					broadcast = "Finished Airing on: " + fmt.Sprintf("%d/%d/%d",
						anime[i].Anime.EndDate.Month,
						anime[i].Anime.EndDate.Day,
						anime[i].Anime.EndDate.Year)
				}

				syn, _ := to.Wrapped(anime[i].Anime.Synopsis, (w/2)-5)
				episodeCount := "👀 Watched: " + to.String(anime[i].Anime.MyListStatus.NumEpisodesWatched) + "/" + to.String(anime[i].Anime.NumEpisodes)
				comments := anime[i].Anime.MyListStatus.Comments

				return fmt.Sprintf("📺  %s\n%s\n%s\n%s\n\nSynopsis:\n%v\n",
					title,
					broadcast,
					comments,
					episodeCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := to.String(anime[indx].Anime.ID)
			url := "https://myanimelist.net/anime/" + sel
			clipboard.WriteAll(url)
			fmt.Println("Url copied to system clipboard...")
			fmt.Println("Opening myanimelist.net in the default browser...")
			fn.OpenURL(url)
		}
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

	Call: func(x *Z.Cmd, args ...string) error {
		clientPublic()
		clientOauth2()

		anime := GetUserAnimeListPlanToWatch()

		idx, err := fuzzyfinder.FindMulti(
			anime,
			func(i int) string {
				title := anime[i].Anime.AlternativeTitles.English
				if title == "" {
					title = anime[i].Anime.Title
				}
				return title
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := anime[i].Anime.Title
				engTitle := anime[i].Anime.AlternativeTitles.English
				if engTitle == "" {
					title = origTitle
				} else {
					title = engTitle
				}

				var broadcast string
				if anime[i].Anime.StartDate.Month == 0 {
					broadcast = "Upcoming"
				} else if anime[i].Anime.EndDate.Month == 0 {
					if len(anime[i].Anime.Broadcast.DayOfTheWeek) == 0 {
						broadcast = "Airs on: Unknown"
					} else {
						broadcast = "Airs on: " + strings.Title(to.String(anime[i].Anime.Broadcast.DayOfTheWeek))
					}
				} else {
					broadcast = "Finished Airing on: " + fmt.Sprintf("%d/%d/%d",
						anime[i].Anime.EndDate.Month,
						anime[i].Anime.EndDate.Day,
						anime[i].Anime.EndDate.Year)
				}

				syn, _ := to.Wrapped(anime[i].Anime.Synopsis, (w/2)-5)
				episodeCount := "👀 Watched: " + to.String(anime[i].Anime.MyListStatus.NumEpisodesWatched) + "/" + to.String(anime[i].Anime.NumEpisodes)
				comments := anime[i].Anime.MyListStatus.Comments

				return fmt.Sprintf("📺  %s\n%s\n%s\n%s\n\nSynopsis:\n%v\n",
					title,
					broadcast,
					comments,
					episodeCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := to.String(anime[indx].Anime.ID)
			url := "https://myanimelist.net/anime/" + sel
			clipboard.WriteAll(url)
			fmt.Println("Url copied to system clipboard...")
			fmt.Println("Opening myanimelist.net in the default browser...")
			fn.OpenURL(url)
		}
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

	Call: func(x *Z.Cmd, args ...string) error {
		clientPublic()
		clientOauth2()

		anime := GetUserAnimeListCompleted()

		idx, err := fuzzyfinder.FindMulti(
			anime,
			func(i int) string {
				title := anime[i].Anime.AlternativeTitles.English
				if title == "" {
					title = anime[i].Anime.Title
				}
				return title
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := anime[i].Anime.Title
				engTitle := anime[i].Anime.AlternativeTitles.English
				if engTitle == "" {
					title = origTitle
				} else {
					title = engTitle
				}

				var broadcast string
				if anime[i].Anime.EndDate.Month == 0 {
					if len(anime[i].Anime.Broadcast.DayOfTheWeek) == 0 {
						broadcast = "Airs on: Unknown"
					} else {
						broadcast = "Airs on: " + strings.Title(to.String(anime[i].Anime.Broadcast.DayOfTheWeek))
					}
				} else {
					broadcast = "Finished Airing on: " + fmt.Sprintf("%d/%d/%d",
						anime[i].Anime.EndDate.Month,
						anime[i].Anime.EndDate.Day,
						anime[i].Anime.EndDate.Year)
				}

				syn, _ := to.Wrapped(anime[i].Anime.Synopsis, (w/2)-5)
				episodeCount := "👀 Watched: " + to.String(anime[i].Anime.MyListStatus.NumEpisodesWatched) + "/" + to.String(anime[i].Anime.NumEpisodes)
				comments := anime[i].Anime.MyListStatus.Comments

				return fmt.Sprintf("📺  %s\n%s\n%s\n%s\n\nSynopsis:\n%v\n",
					title,
					broadcast,
					comments,
					episodeCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := to.String(anime[indx].Anime.ID)
			url := "https://myanimelist.net/anime/" + sel
			clipboard.WriteAll(url)
			fmt.Println("Url copied to system clipboard...")
			fmt.Println("Opening myanimelist.net in the default browser...")
			fn.OpenURL(url)
		}
		return nil
	},
}

// // ─────────────────────────── Manga ───────────────────────────────

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
		clientPublic()
		clientOauth2()

		manga := GetUserMangaList()

		idx, err := fuzzyfinder.FindMulti(
			manga,
			func(i int) string {
				return to.String(manga[i].Manga.Title)
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := manga[i].Manga.Title
				engTitle := manga[i].Manga.AlternativeTitles.English
				if engTitle == "" {
					title = "📘 " + origTitle
				} else {
					title = "📘 " + engTitle
				}

				var readCount string
				if manga[i].Manga.NumChapters == 0 {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + "?"
				} else {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + to.String(manga[i].Manga.NumChapters)
				}

				var status string
				s := manga[i].Manga.Status
				if s == "currently_publishing" {
					status = "🔖 Ongoing"
				}
				if s == "finished" {
					status = "🔖 Complete"
				}

				syn, _ := to.Wrapped(manga[i].Manga.Synopsis, (w/2)-5)
				return fmt.Sprintf("%s\n%s\n%s\n\n%s\n",
					title,
					status,
					readCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := manga[indx].Manga.Title
			clipboard.WriteAll(sel)
			fmt.Println(sel)
		}

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
		clientPublic()
		clientOauth2()

		manga := GetUserMangaListPlanToRead()

		idx, err := fuzzyfinder.FindMulti(
			manga,
			func(i int) string {
				return to.String(manga[i].Manga.Title)
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := manga[i].Manga.Title
				engTitle := manga[i].Manga.AlternativeTitles.English
				if engTitle == "" {
					title = "📘 " + origTitle
				} else {
					title = "📘 " + engTitle
				}

				var readCount string
				if manga[i].Manga.NumChapters == 0 {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + "?"
				} else {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + to.String(manga[i].Manga.NumChapters)
				}

				var status string
				s := manga[i].Manga.Status
				if s == "currently_publishing" {
					status = "🔖 Ongoing"
				}
				if s == "finished" {
					status = "🔖 Complete"
				}

				syn, _ := to.Wrapped(manga[i].Manga.Synopsis, (w/2)-5)
				return fmt.Sprintf("%s\n%s\n%s\n\n%s\n",
					title,
					status,
					readCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := manga[indx].Manga.Title
			clipboard.WriteAll(sel)
			fmt.Println(sel)
		}

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
		clientPublic()
		clientOauth2()

		manga := GetUserMangaListDone()

		idx, err := fuzzyfinder.FindMulti(
			manga,
			func(i int) string {
				return to.String(manga[i].Manga.Title)
			},
			fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}

				var title string
				origTitle := manga[i].Manga.Title
				engTitle := manga[i].Manga.AlternativeTitles.English
				if engTitle == "" {
					title = "📘 " + origTitle
				} else {
					title = "📘 " + engTitle
				}

				var readCount string
				if manga[i].Manga.NumChapters == 0 {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + "?"
				} else {
					readCount = "🔎 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead) + "/" + to.String(manga[i].Manga.NumChapters)
				}

				var status string
				s := manga[i].Manga.Status
				if s == "currently_publishing" {
					status = "🔖 Ongoing"
				}
				if s == "finished" {
					status = "🔖 Complete"
				}

				syn, _ := to.Wrapped(manga[i].Manga.Synopsis, (w/2)-5)
				return fmt.Sprintf("%s\n%s\n%s\n\n%s\n",
					title,
					status,
					readCount,
					syn)
			}))
		if err != nil {
			log.Fatal(err)
		}
		for _, indx := range idx {
			sel := manga[indx].Manga.Title
			clipboard.WriteAll(sel)
			fmt.Println(sel)
		}

		return nil
	},
}
