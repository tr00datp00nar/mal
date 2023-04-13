package mal

import (
	"context"
	"flag"

	"github.com/charmbracelet/log"
	"github.com/nstratos/go-myanimelist/mal"
)

func runWatching() error {
	config, err := LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	var (
		clientID     = flag.String("client-id", config.ClientId, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", config.ClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		state        = flag.String("state", "", "token to protect against CSRF attacks")
	)
	flag.Parse()

	ctx := context.Background()

	tokenClient, err := authenticate(ctx, *clientID, *clientSecret, *state)
	if err != nil {
		return err
	}

	c := client{
		Client: mal.NewClient(tokenClient),
	}

	return c.showcaseWatching(ctx)
}

func runPlanToWatch() error {
	config, err := LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	var (
		clientID     = flag.String("client-id", config.ClientId, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", config.ClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		state        = flag.String("state", "", "token to protect against CSRF attacks")
	)
	flag.Parse()

	ctx := context.Background()

	tokenClient, err := authenticate(ctx, *clientID, *clientSecret, *state)
	if err != nil {
		return err
	}

	c := client{
		Client: mal.NewClient(tokenClient),
	}

	return c.showcasePlanToWatch(ctx)
}

func runCompleted() error {
	config, err := LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	var (
		clientID     = flag.String("client-id", config.ClientId, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", config.ClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		state        = flag.String("state", "", "token to protect against CSRF attacks")
	)
	flag.Parse()

	ctx := context.Background()

	tokenClient, err := authenticate(ctx, *clientID, *clientSecret, *state)
	if err != nil {
		return err
	}

	c := client{
		Client: mal.NewClient(tokenClient),
	}

	return c.showcaseCompleted(ctx)
}

func runMangaList() error {
	config, err := LoadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	var (
		clientID     = flag.String("client-id", config.ClientId, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", config.ClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		state        = flag.String("state", "", "token to protect against CSRF attacks")
	)
	flag.Parse()

	ctx := context.Background()

	tokenClient, err := authenticate(ctx, *clientID, *clientSecret, *state)
	if err != nil {
		return err
	}

	c := client{
		Client: mal.NewClient(tokenClient),
	}

	return c.showcaseMangaList(ctx)
}
