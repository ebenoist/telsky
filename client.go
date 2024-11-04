package main

import (
	"log"
	"fmt"
	"context"

	"github.com/bluesky-social/indigo/util/cliutil"
	"github.com/bluesky-social/indigo/xrpc"

	"path/filepath"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
)

type Post struct {
	raw *bsky.FeedDefs_FeedViewPost
}

func getAuth() (*xrpc.AuthInfo) {
	auth, err := cliutil.ReadAuth(filepath.Join("/home/erik/.config/telsky/client.auth"))
	if err == nil {
		return auth
	}

	panic(fmt.Sprintf("PANIC %s", err))

	return nil
}

func newClient() *xrpc.Client {
	return &xrpc.Client{
		Client: cliutil.NewHttpClient(),
		Host: "https://bsky.social",
		Auth: getAuth(),
	}
}

func (p Post) Title() string {
	createdAt, _ := time.Parse(time.RFC3339, p.raw.Post.Record.Val.(*bsky.FeedPost).CreatedAt)
	return fmt.Sprintf(
		"%s (@%s) - %s",
		*p.raw.Post.Author.DisplayName,
		p.raw.Post.Author.Handle,
		createdAt.Local().Format(time.Kitchen),
	)
}

func (p Post) Description() string {
	return p.raw.Post.Record.Val.(*bsky.FeedPost).Text
}

func (p Post) FilterValue() string { return p.Title() }

func buildPosts(raw *bsky.FeedGetTimeline_Output) []Post {
	posts := make([]Post, 0, len(raw.Feed))

	for _, r := range raw.Feed {
		posts = append(posts, Post{raw: r})
	}

	return posts;
}

func getAuthorFeed(author string) ([]Post, error) {
	log.Printf("getting feed for %s", author)
	client := newClient()

	resp, err := bsky.FeedGetTimeline(context.TODO(), client, "reverse-chronological", "", 24)

	return buildPosts(resp), err
}
