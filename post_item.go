package main

import (
	"fmt"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
)

type PostItem struct {
	raw *bsky.FeedDefs_FeedViewPost
}

func (p PostItem) Title() string {
	createdAt, _ := time.Parse(time.RFC3339, p.raw.Post.Record.Val.(*bsky.FeedPost).CreatedAt)
	return fmt.Sprintf(
		"%s (@%s) - %s",
		*p.raw.Post.Author.DisplayName,
		p.raw.Post.Author.Handle,
		createdAt.Local().Format(time.Kitchen),
	)
}

func (p PostItem) Description() string {
	return p.raw.Post.Record.Val.(*bsky.FeedPost).Text
}

func (p PostItem) FilterValue() string { return p.Title() }

func buildPosts(raw *bsky.FeedGetTimeline_Output) []PostItem {
	posts := make([]PostItem, 0, len(raw.Feed))

	for _, r := range raw.Feed {
		posts = append(posts, PostItem{raw: r})
	}

	return posts;
}
