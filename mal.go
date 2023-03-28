package mal

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/nstratos/go-myanimelist/mal"
)

func main(arg int) {
	switch arg {
	case 1:
		if err := runWatching(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case 2:
		if err := runPlanToWatch(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case 3:
		if err := runCompleted(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}

func runWatching() error {
	var (
		clientID     = flag.String("client-id", defaultClientID, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", defaultClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		// state is a token to protect the user from CSRF attacks. In a web
		// application, you should provide a non-empty string and validate that
		// it matches the state query parameter on the redirect URL callback
		// after the MyAnimeList authentication. It can stay empty here.
		state = flag.String("state", "", "token to protect against CSRF attacks")
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
	// return c.showcasePlanToWatch(ctx)
	// return c.showcaseCompleted(ctx)
}

func runPlanToWatch() error {
	var (
		clientID     = flag.String("client-id", defaultClientID, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", defaultClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		// state is a token to protect the user from CSRF attacks. In a web
		// application, you should provide a non-empty string and validate that
		// it matches the state query parameter on the redirect URL callback
		// after the MyAnimeList authentication. It can stay empty here.
		state = flag.String("state", "", "token to protect against CSRF attacks")
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
	var (
		clientID     = flag.String("client-id", defaultClientID, "your registered MyAnimeList.net application client ID")
		clientSecret = flag.String("client-secret", defaultClientSecret, "your registered MyAnimeList.net application client secret; optional if you chose App Type 'other'")
		// state is a token to protect the user from CSRF attacks. In a web
		// application, you should provide a non-empty string and validate that
		// it matches the state query parameter on the redirect URL callback
		// after the MyAnimeList authentication. It can stay empty here.
		state = flag.String("state", "", "token to protect against CSRF attacks")
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
