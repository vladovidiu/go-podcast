package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vladovidiu/go-podcast/feeds"
	"github.com/vladovidiu/go-podcast/graph/generated"
	"github.com/vladovidiu/go-podcast/graph/model"
	"github.com/vladovidiu/go-podcast/itunes"
	"github.com/vladovidiu/go-podcast/utils"
)

func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	as := itunes.NewAPIServices()

	res, err := as.Search(term)
	if err != nil {
		return nil, err
	}

	var podcasts []*model.Podcast

	for _, result := range res.Results {
		podcast := &model.Podcast{
			Artist:        result.ArtistName,
			PodcastName:   result.TrackName,
			FeedURL:       result.FeedURL,
			Thumbnail:     result.ArtworkURL100,
			EpisodesCount: result.TrackCount,
			Genres:        result.Genres,
		}

		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feedItems []*model.FeedItem

	for _, item := range res.Channel.Item {
		feedItem := &model.FeedItem{
			PubDate:     item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       utils.CheckNullString(item.Image.Href),
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
