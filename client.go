package main

import (
	"context"
	"log"

	"github.com/bluesky-social/indigo/util/cliutil"
	"github.com/bluesky-social/indigo/xrpc"

	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/bluesky-social/indigo/api/bsky"
)

func newClient() *xrpc.Client {
	return &xrpc.Client{
		Client: cliutil.NewHttpClient(),
		Host: "https://bsky.social",
		Auth: getAuth(),
	}
}

func getAuth() (*xrpc.AuthInfo) {
	auth, err := cliutil.ReadAuth(filepath.Join(xdg.ConfigHome, "telsky", "client.auth"))
	if err == nil {
		return auth
	}

	log.Panicf("could not read config - %s")
	return nil
}

func getAuthorFeed() ([]PostItem, error) {
	client := newClient()
	resp, err := bsky.FeedGetTimeline(context.TODO(), client, "reverse-chronological", "", 50)

	return buildPosts(resp), err
}
