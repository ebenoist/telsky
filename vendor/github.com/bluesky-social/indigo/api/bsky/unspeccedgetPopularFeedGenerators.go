// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.unspecced.getPopularFeedGenerators

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// UnspeccedGetPopularFeedGenerators_Output is the output of a app.bsky.unspecced.getPopularFeedGenerators call.
type UnspeccedGetPopularFeedGenerators_Output struct {
	Cursor *string                   `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Feeds  []*FeedDefs_GeneratorView `json:"feeds" cborgen:"feeds"`
}

// UnspeccedGetPopularFeedGenerators calls the XRPC method "app.bsky.unspecced.getPopularFeedGenerators".
func UnspeccedGetPopularFeedGenerators(ctx context.Context, c *xrpc.Client, cursor string, limit int64, query string) (*UnspeccedGetPopularFeedGenerators_Output, error) {
	var out UnspeccedGetPopularFeedGenerators_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
		"query":  query,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.unspecced.getPopularFeedGenerators", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}