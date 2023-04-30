package mal

import (
	"github.com/rl404/nagato"
)

func GetUserAnimeListWatching() (response []nagato.UserAnime) {
	d, _, err := nagatoClient.GetUserAnimeList(nagato.GetUserAnimeListParam{
		Username: "@me",
		Status:   nagato.UserAnimeStatusWatching,
		NSFW:     true,
		Sort:     nagato.UserAnimeSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
		nagato.AnimeFieldUserStatus(nagato.UserAnimeTags),
		nagato.AnimeFieldUserStatus(nagato.UserAnimeComments),
	)
	if err != nil {
		panic(err)
	}
	return d
}

func GetUserAnimeListPlanToWatch() (response []nagato.UserAnime) {
	d, _, err := nagatoClient.GetUserAnimeList(nagato.GetUserAnimeListParam{
		Username: "@me",
		Status:   nagato.UserAnimeStatusPlanToWatch,
		NSFW:     true,
		Sort:     nagato.UserAnimeSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
		nagato.AnimeFieldUserStatus(nagato.UserAnimeTags),
		nagato.AnimeFieldUserStatus(nagato.UserAnimeComments),
	)
	if err != nil {
		panic(err)
	}
	return d
}

func GetUserAnimeListCompleted() (response []nagato.UserAnime) {
	d, _, err := nagatoClient.GetUserAnimeList(nagato.GetUserAnimeListParam{
		Username: "@me",
		Status:   nagato.UserAnimeStatusCompleted,
		NSFW:     true,
		Sort:     nagato.UserAnimeSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
		nagato.AnimeFieldUserStatus(nagato.UserAnimeTags),
		nagato.AnimeFieldUserStatus(nagato.UserAnimeComments),
	)
	if err != nil {
		panic(err)
	}
	return d
}

func GetUserMangaList() (response []nagato.UserManga) {
	d, _, err := nagatoClient.GetUserMangaList(nagato.GetUserMangaListParam{
		Username: "@me",
		Status:   nagato.UserMangaStatusReading,
		NSFW:     true,
		Sort:     nagato.UserMangaSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
		nagato.MangaFieldUserStatus(nagato.UserMangaTags),
	)
	if err != nil {
		panic(err)
	}
	return d
}

func GetUserMangaListPlanToRead() (response []nagato.UserManga) {
	d, _, err := nagatoClient.GetUserMangaList(nagato.GetUserMangaListParam{
		Username: "@me",
		Status:   nagato.UserMangaStatusPlanToRead,
		NSFW:     true,
		Sort:     nagato.UserMangaSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
		nagato.MangaFieldUserStatus(nagato.UserMangaTags),
	)
	if err != nil {
		panic(err)
	}
	return d
}

func GetUserMangaListDone() (response []nagato.UserManga) {
	d, _, err := nagatoClient.GetUserMangaList(nagato.GetUserMangaListParam{
		Username: "@me",
		Status:   nagato.UserMangaStatusCompleted,
		NSFW:     true,
		Sort:     nagato.UserMangaSortUpdatedAt,
		Offset:   0,
		// Limit:    -1,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
		nagato.MangaFieldUserStatus(nagato.UserMangaTags),
	)
	if err != nil {
		panic(err)
	}
	return d
}
