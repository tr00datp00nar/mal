package mal

import (
	"context"
	"fmt"
	"unicode"

	"github.com/atotto/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"

	"github.com/charmbracelet/log"
	"github.com/nstratos/go-myanimelist/mal"
	"github.com/rwxrob/to"
)

// client has methods showcasing the usage of the different MyAnimeList API
// methods. It stores the first error it encounters so error checking only needs
// to be done once.
//
// This pattern is used for convenience and should not be used in concurrent
// code without guarding the error.
type client struct {
	*mal.Client
	err error
}

func (c *client) showcaseWatching(ctx context.Context) error {
	methods := []func(context.Context){
		c.userAnimeListWatching,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) showcasePlanToWatch(ctx context.Context) error {
	methods := []func(context.Context){
		c.userAnimeListPlanToWatch,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) showcaseCompleted(ctx context.Context) error {
	methods := []func(context.Context){
		c.userAnimeListCompleted,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) showcaseMangaList(ctx context.Context) error {
	methods := []func(context.Context){
		c.userMangaList,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) showcaseMangaListPlan(ctx context.Context) error {
	methods := []func(context.Context){
		c.userMangaListPlan,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) showcaseMangaListDone(ctx context.Context) error {
	methods := []func(context.Context){
		c.userMangaListDone,
	}
	for _, m := range methods {
		m(ctx)
	}
	if c.err != nil {
		return c.err
	}
	return nil
}

func (c *client) userAnimeListWatching(ctx context.Context) {
	if c.err != nil {
		return
	}
	anime, _, err := c.User.AnimeList(ctx, "@me",
		mal.Fields{"my_list_status{comments}", "num_episodes", "synopsis"},
		mal.AnimeStatusWatching,
		mal.SortAnimeListByListUpdatedAt,
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}

	idx, err := fuzzyfinder.FindMulti(
		anime,
		func(i int) string {
			return to.String(anime[i].Anime.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {

			if i == -1 {
				return ""
			}

			syn, _ := to.Wrapped(anime[i].Anime.Synopsis, 50)
			episodeCount := "👀 Watched: " + to.String(anime[i].Anime.MyListStatus.NumEpisodesWatched) + "/" + to.String(anime[i].Anime.NumEpisodes)

			return fmt.Sprintf("📺  %s\n%s\n%s\n\nSynopsis:\n%v\n",
				anime[i].Anime.Title,
				anime[i].Anime.MyListStatus.Comments,
				episodeCount,
				syn)
		}))
	if err != nil {
		log.Fatal(err)
	}
	for _, indx := range idx {
		sel := anime[indx].Anime.Title
		clipboard.WriteAll(sel)
		fmt.Println(sel)
	}
}

func (c *client) userAnimeListPlanToWatch(ctx context.Context) {
	anime, _, err := c.User.AnimeList(ctx, "@me",
		mal.Fields{"my_list_status{comments}", "list_status", "synopsis"},
		mal.AnimeStatusPlanToWatch,
		mal.SortAnimeListByListUpdatedAt,
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}

	idx, err := fuzzyfinder.FindMulti(
		anime,
		func(i int) string {
			return to.String(anime[i].Anime.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			syn, _ := to.Wrapped(anime[i].Anime.Synopsis, 72)
			return fmt.Sprintf("📺  %s\n%s\n\nSynopsis:\n%v\n",
				anime[i].Anime.Title,
				anime[i].Anime.MyListStatus.Comments,
				syn)
		}))
	if err != nil {
		log.Fatal(err)
	}
	for _, indx := range idx {
		sel := anime[indx].Anime.Title
		clipboard.WriteAll(sel)
		fmt.Println(sel)
	}
}

func (c *client) userAnimeListCompleted(ctx context.Context) {
	if c.err != nil {
		return
	}
	anime, _, err := c.User.AnimeList(ctx, "@me",
		mal.Fields{"my_list_status{comments}", "list_status", "synopsis"},
		mal.AnimeStatusCompleted,
		mal.SortAnimeListByListUpdatedAt,
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}
	idx, err := fuzzyfinder.FindMulti(
		anime,
		func(i int) string {
			return to.String(anime[i].Anime.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			syn, _ := to.Wrapped(anime[i].Anime.Synopsis, 72)
			return fmt.Sprintf("%s\n%s\n\n%s\n",
				anime[i].Anime.Title,
				anime[i].Anime.MyListStatus.Comments,
				syn)
		}))
	if err != nil {
		log.Fatal(err)
	}
	for _, indx := range idx {
		sel := anime[indx].Anime.Title
		clipboard.WriteAll(sel)
		fmt.Println(sel)
	}
}

func (c *client) userMangaList(ctx context.Context) {
	if c.err != nil {
		return
	}
	manga, _, err := c.User.MangaList(ctx, "@me",
		mal.MangaStatusReading,
		mal.SortMangaListByListUpdatedAt,
		mal.Fields{"list_status", "synopsis", "my_list_status{comments}"},
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}
	idx, err := fuzzyfinder.FindMulti(
		manga,
		func(i int) string {
			return to.String(manga[i].Manga.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			readCount := "📖 Chapters Read: " + to.String(manga[i].Manga.MyListStatus.NumChaptersRead)
			syn, _ := to.Wrapped(manga[i].Manga.Synopsis, 72)
			return fmt.Sprintf("%s\n%s\n%s\n\n%s\n",
				manga[i].Manga.Title,
				manga[i].Manga.MyListStatus.Comments,
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
}

func (c *client) userMangaListPlan(ctx context.Context) {
	if c.err != nil {
		return
	}
	manga, _, err := c.User.MangaList(ctx, "@me",
		mal.MangaStatusPlanToRead,
		mal.SortMangaListByListUpdatedAt,
		mal.Fields{"list_status", "synopsis", "my_list_status{comments}"},
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}
	idx, err := fuzzyfinder.FindMulti(
		manga,
		func(i int) string {
			return to.String(manga[i].Manga.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			syn, _ := to.Wrapped(manga[i].Manga.Synopsis, 72)
			return fmt.Sprintf("%s\n%s\n\n%s\n",
				manga[i].Manga.Title,
				manga[i].Manga.MyListStatus.Comments,
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
}

func (c *client) userMangaListDone(ctx context.Context) {
	if c.err != nil {
		return
	}
	manga, _, err := c.User.MangaList(ctx, "@me",
		mal.MangaStatusCompleted,
		mal.SortMangaListByListUpdatedAt,
		mal.Fields{"list_status", "synopsis", "my_list_status{comments}"},
		mal.Limit(100),
	)
	if err != nil {
		c.err = err
		return
	}
	idx, err := fuzzyfinder.FindMulti(
		manga,
		func(i int) string {
			return to.String(manga[i].Manga.Title)
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			syn, _ := to.Wrapped(manga[i].Manga.Synopsis, 72)
			return fmt.Sprintf("%s\n%s\n\n%s\n",
				manga[i].Manga.Title,
				manga[i].Manga.MyListStatus.Comments,
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
}

// ────────────────────────────────────────────────────────────

// func (c *client) userMyInfo(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	u, _, err := c.User.MyInfo(ctx)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("ID: %5d, Joined: %v, Username: %s\n", u.ID, u.JoinedAt.Format("Jan 2006"), u.Name)
// }
//
// func (c *client) animeList(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	anime, _, err := c.Anime.List(ctx, "hokuto no ken",
// 		mal.Fields{"rank", "popularity", "start_season"},
// 		mal.Limit(3),
// 		mal.Offset(0),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, a := range anime {
// 		fmt.Printf("ID: %5d, Rank: %5d, Popularity: %5d %s (%d)\n", a.ID, a.Rank, a.Popularity, a.Title, a.StartSeason.Year)
// 	}
// }
//
// func (c *client) mangaList(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	manga, _, err := c.Manga.List(ctx, "parasyte",
// 		mal.Fields{"num_volumes", "num_chapters", "alternative_titles"},
// 		mal.Limit(3),
// 		mal.Offset(0),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, m := range manga {
// 		fmt.Printf("ID: %5d, Volumes: %3d, Chapters: %3d %s (%s)\n", m.ID, m.NumVolumes, m.NumChapters, m.Title, m.AlternativeTitles.En)
// 	}
// }
//
// func (c *client) animeDetails(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	a, _, err := c.Anime.Details(ctx, 967,
// 		mal.Fields{
// 			"alternative_titles",
// 			"media_type",
// 			"num_episodes",
// 			"start_season",
// 			"source",
// 			"genres",
// 			"studios",
// 			"average_episode_duration",
// 		},
// 	)
//
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("%s\n", a.Title)
// 	fmt.Printf("ID: %d\n", a.ID)
// 	fmt.Printf("English: %s\n", a.AlternativeTitles.En)
// 	fmt.Printf("Type: %s\n", strings.ToUpper(a.MediaType))
// 	fmt.Printf("Episodes: %d\n", a.NumEpisodes)
// 	fmt.Printf("Premiered: %s %d\n", strings.Title(a.StartSeason.Season), a.StartSeason.Year)
// 	fmt.Print("Studios: ")
// 	delim := ""
// 	for _, s := range a.Studios {
// 		fmt.Printf("%s%s", delim, s.Name)
// 		delim = " "
// 	}
// 	fmt.Println()
// 	fmt.Printf("Source: %s\n", strings.Title(a.Source))
// 	fmt.Print("Genres: ")
// 	delim = ""
// 	for _, g := range a.Genres {
// 		fmt.Printf("%s%s", delim, g.Name)
// 		delim = " "
// 	}
// 	fmt.Println()
// 	fmt.Printf("Duration: %d min. per ep.\n", a.AverageEpisodeDuration/60)
// }
//
// func (c *client) mangaDetails(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	m, _, err := c.Manga.Details(ctx, 401,
// 		mal.Fields{
// 			"alternative_titles",
// 			"media_type",
// 			"num_volumes",
// 			"num_chapters",
// 			"authors{last_name, first_name}",
// 			"genres",
// 			"status",
// 		},
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("%s\n", m.Title)
// 	fmt.Printf("ID: %d\n", m.ID)
// 	fmt.Printf("English: %s\n", m.AlternativeTitles.En)
// 	fmt.Printf("Type: %s\n", strings.Title(m.MediaType))
// 	fmt.Printf("Volumes: %d\n", m.NumVolumes)
// 	fmt.Printf("Chapters: %d\n", m.NumChapters)
// 	fmt.Print("Studios: ")
// 	delim := ""
// 	for _, s := range m.Authors {
// 		fmt.Printf("%s%s, %s (%s)", delim, s.Person.LastName, s.Person.FirstName, s.Role)
// 		delim = " "
// 	}
// 	fmt.Println()
// 	fmt.Print("Genres: ")
// 	delim = ""
// 	for _, g := range m.Genres {
// 		fmt.Printf("%s%s", delim, g.Name)
// 		delim = " "
// 	}
// 	fmt.Println()
// 	fmt.Printf("Status: %s\n", strings.Title(m.Status))
// }
//
// func (c *client) animeListForLoop(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	offset := 0
// 	for {
// 		anime, resp, err := c.Anime.List(ctx, "kiseijuu",
// 			mal.Fields{"rank", "popularity", "start_season"},
// 			mal.Limit(100),
// 			mal.Offset(offset),
// 		)
// 		if err != nil {
// 			c.err = err
// 			return
// 		}
// 		for _, a := range anime {
// 			fmt.Printf("ID: %5d, Rank: %5d, Popularity: %5d %s (%d)\n", a.ID, a.Rank, a.Popularity, a.Title, a.StartSeason.Year)
// 		}
// 		fmt.Println("--------")
// 		fmt.Printf("Next offset: %d\n", resp.NextOffset)
// 		offset = resp.NextOffset
// 		if offset == 0 {
// 			break
// 		}
// 	}
// }
//
// func (c *client) updateMyAnimeListStatus(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	s, _, err := c.Anime.UpdateMyListStatus(ctx, 967,
// 		mal.AnimeStatusWatching,
// 		mal.NumEpisodesWatched(73),
// 		mal.Score(8),
// 		mal.Comments("You wa shock!"),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("Status: %q, Score: %d, Episodes Watched: %d, Comments: %s\n", s.Status, s.Score, s.NumEpisodesWatched, s.Comments)
// }
//
// func (c *client) updateMyMangaListStatus(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	s, _, err := c.Manga.UpdateMyListStatus(ctx, 401,
// 		mal.MangaStatusReading,
// 		mal.NumVolumesRead(1),
// 		mal.NumChaptersRead(5),
// 		mal.Comments("Migi"),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("Status: %q, Volumes Read: %d, Chapters Read: %d, Comments: %s\n", s.Status, s.NumVolumesRead, s.NumChaptersRead, s.Comments)
// }
//
// func (c *client) deleteMyAnimeListItem(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	_, err := c.Anime.DeleteMyListItem(ctx, 820)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// }
//
// func (c *client) deleteMyMangaListItem(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	_, err := c.Manga.DeleteMyListItem(ctx, 1)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// }
//
// func (c *client) animeRanking(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	rankings := []mal.AnimeRanking{
// 		mal.AnimeRankingAiring,
// 		mal.AnimeRankingAll,
// 		mal.AnimeRankingByPopularity,
// 	}
// 	for _, r := range rankings {
// 		fmt.Println("Ranking:", r)
// 		anime, _, err := c.Anime.Ranking(ctx, r,
// 			mal.Fields{"rank", "popularity"},
// 		)
// 		if err != nil {
// 			c.err = err
// 			return
// 		}
// 		for _, a := range anime {
// 			fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
// 		}
// 		fmt.Println("--------")
// 	}
// }
//
// func (c *client) mangaRanking(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	manga, _, err := c.Manga.Ranking(ctx,
// 		mal.MangaRankingByPopularity,
// 		mal.Fields{"rank", "popularity"},
// 		mal.Limit(6),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, m := range manga {
// 		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", m.Rank, m.Popularity, m.Title)
// 	}
// }
//
// func (c *client) animeSeasonal(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	anime, _, err := c.Anime.Seasonal(ctx, 2020, mal.AnimeSeasonFall,
// 		mal.Fields{"rank", "popularity"},
// 		mal.SortSeasonalByAnimeNumListUsers,
// 		mal.Limit(3),
// 		mal.Offset(0),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, a := range anime {
// 		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
// 	}
// }
//
// func (c *client) animeSuggested(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	anime, _, err := c.Anime.Suggested(ctx,
// 		mal.Limit(3),
// 		mal.Fields{"rank", "popularity"},
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, a := range anime {
// 		fmt.Printf("Rank: %5d, Popularity: %5d %s\n", a.Rank, a.Popularity, a.Title)
// 	}
// }
//
// func (c *client) forumBoards(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	forum, _, err := c.Forum.Boards(ctx)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, category := range forum.Categories {
// 		fmt.Printf("%s\n", category.Title)
// 		for _, b := range category.Boards {
// 			fmt.Printf("|-> %s\n", b.Title)
// 			for _, b := range b.Subboards {
// 				fmt.Printf("    |-> %s\n", b.Title)
// 			}
// 		}
// 		fmt.Println("---")
// 	}
// }
//
// func (c *client) forumTopics(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	topics, _, err := c.Forum.Topics(ctx,
// 		mal.Query("JoJo opening"),
// 		mal.SortTopicsRecent,
// 		mal.Limit(2),
// 	)
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	for _, t := range topics {
// 		fmt.Printf("ID: %5d, Title: %5q created by %q\n", t.ID, t.Title, t.CreatedBy.Name)
// 	}
// }
//
// func (c *client) forumTopicDetails(ctx context.Context) {
// 	if c.err != nil {
// 		return
// 	}
// 	topicDetails, _, err := c.Forum.TopicDetails(ctx, 1877721, mal.Limit(3), mal.Offset(0))
// 	if err != nil {
// 		c.err = err
// 		return
// 	}
// 	fmt.Printf("Topic title: %q\n", topicDetails.Title)
// 	if topicDetails.Poll != nil {
// 		fmt.Printf("Poll: %q\n", topicDetails.Poll.Question)
// 		for _, o := range topicDetails.Poll.Options {
// 			fmt.Printf("- %-25s %2d\n", o.Text, o.Votes)
// 		}
// 	}
// 	for _, p := range topicDetails.Posts {
// 		fmt.Printf("Post: %2d created by %q\n", p.Number, p.CreatedBy.Name)
// 	}
// }

func EllipticalTruncate(text string, maxLen int) string {
	lastSpaceIx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}

// func truncateString(s string, max int) string {
// 	if max > len(s) {
// 		return s
// 	}
// 	return s[:strings.LastIndexAny(s[:max], " .,:;-")]
// }
