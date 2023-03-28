package mal

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nstratos/go-myanimelist/mal"
)

func main(args string) {
	if err := run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Authorization Documentation:
//
// https://myanimelist.net/apiconfig/references/authorization

func run(args string) error {
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

	l := len(args)
	switch l {
	case 0:
		methods := []func(context.Context){c.userAnimeListWatching}
		for _, m := range methods {
			m(ctx)
		}
		if c.err != nil {
			return c.err
		}
		return nil
	case 1:
		if strings.Contains(args, "plan") {
			methods := []func(context.Context){c.userAnimeListPlanToWatch}
			for _, m := range methods {
				m(ctx)
			}
			if c.err != nil {
				return c.err
			}
			return nil
		}
		if strings.Contains(args, "done") {
			methods := []func(context.Context){c.userAnimeListCompleted}
			for _, m := range methods {
				m(ctx)
			}
			if c.err != nil {
				return c.err
			}
			return nil
		}
	}

	// if args == "" {
	// 	methods := []func(context.Context){c.userAnimeListWatching}
	// 	for _, m := range methods {
	// 		m(ctx)
	// 	}
	// 	if c.err != nil {
	// 		return c.err
	// 	}
	// 	return nil
	// }
	//
	// if args == "done" {
	// 	methods := []func(context.Context){c.userAnimeListCompleted}
	// 	for _, m := range methods {
	// 		m(ctx)
	// 	}
	// 	if c.err != nil {
	// 		return c.err
	// 	}
	// 	return nil
	// }
	//
	// if args == "planned" {
	// 	methods := []func(context.Context){c.userAnimeListPlanToWatch}
	// 	for _, m := range methods {
	// 		m(ctx)
	// 	}
	// 	if c.err != nil {
	// 		return c.err
	// 	}
	// 	return nil
	// }

	return c.showcase(ctx)
}
