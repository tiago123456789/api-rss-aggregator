package service

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/model"
	"github.com/tiago123456789/api-rss-aggregator/internal/api_rss/repository"
	"github.com/tiago123456789/api-rss-aggregator/pkg/rss"
)

type RssFeedService struct {
	postRepository *repository.PostRepository
	feedRepository *repository.FeedRepository
}

func NewRssFeedService(
	postRepository *repository.PostRepository,
	feedRepository *repository.FeedRepository,
) *RssFeedService {
	return &RssFeedService{
		postRepository: postRepository,
		feedRepository: feedRepository,
	}
}

func (service *RssFeedService) ProcessToUpdateNewPostsRssFeedUrl() {
	fmt.Println("Start process to take posts based RSS feed url")
	feeds, err := service.feedRepository.GetFeedsToProcessRSS()
	if err != nil {
		fmt.Printf("%v", err)
	}

	var feedIds []int64
	for _, feed := range feeds {
		feedIds = append(feedIds, feed.ID)
	}

	postsByFeedIds, err := service.postRepository.GetPostByFeedIds(feedIds)
	if err != nil {
		fmt.Printf("%v", err)
	}

	var feedWithPostAlreadyCreated map[string]bool = make(map[string]bool)
	for _, item := range postsByFeedIds {
		key := fmt.Sprintf("%v#%v", item.FeedId, item.Title)
		feedWithPostAlreadyCreated[key] = true
	}

	for _, feed := range feeds {
		fmt.Printf("\nStart process to take posts the RSS feed with url %s of feed id %d", feed.Url, feed.ID)
		rssFeed := rss.Parse(feed.Url)
		var postsToSave []model.Post

		fmt.Printf("\nProcessing new posts of url %s of feed id %d", feed.Url, feed.ID)
		for _, v := range rssFeed.Channel.Item {
			key := fmt.Sprintf("%v#%v", feed.ID, v.Title)
			if !feedWithPostAlreadyCreated[key] {
				publishedAt, _ := dateparse.ParseLocal(v.PublishedAt)
				post := model.Post{
					Title:       v.Title,
					Link:        v.Link,
					PublishedAt: publishedAt,
					FeedId:      feed.ID,
				}
				postsToSave = append(postsToSave, post)
			}
		}

		if len(postsToSave) > 0 {
			fmt.Printf("\nSaving new posts of url %s of feed id %d", feed.Url, feed.ID)
			service.postRepository.InsertMany(postsToSave)
			service.feedRepository.SetLastFetchedAt(feed.ID)
			fmt.Printf("\nSaved new posts of url %s", feed.Url)
		} else {
			service.feedRepository.SetLastFetchedAt(feed.ID)
		}
	}
	fmt.Printf("\nFinished process to take posts the RSS feed")
}
