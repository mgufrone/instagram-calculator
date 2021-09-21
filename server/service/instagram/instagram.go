package instagram

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mgufrone/insta-calculator/server/helpers"
	"log"
	"net/http"
	"regexp"
)

type Service struct {
	agent *http.Client
}

func New(agent *http.Client) *Service {
	return &Service{agent}
}

func (i *Service) RetrieveProfile(username string) (eng *Engagement, err error) {
	var (
		req  *http.Request
		resp *http.Response
		html *goquery.Document
		res  *Response
	)
	err = helpers.TryOrError(func() (err error) {
		uriString := fmt.Sprintf("https://www.instagram.com/%s/", username)
		log.Println("loading profile", uriString)
		req, err = http.NewRequestWithContext(context.Background(), "GET", uriString , nil)
		return
	}, func() (err error) {
		resp, err = i.agent.Do(req)
		return
	}, func() (err error) {
		defer resp.Body.Close()
		log.Println("response", resp.Status, resp.StatusCode)
		if resp.StatusCode > 299 {
			return errors.New(fmt.Sprintf("request failed: %s", resp.Status))
		}
		html, err = goquery.NewDocumentFromReader(resp.Body)
		return
	}, func() error {
		script := html.Find("script:contains(edge_owner_to_timeline_media)")
		extract := regexp.MustCompile(`window\._sharedData = (.*);(window)?`)
		str := extract.FindStringSubmatch(script.Text())
		return json.Unmarshal([]byte(str[1]), &res)
	}, func() error {
		eng = &Engagement{
			Followers: int64(res.EntryData.ProfilePage[0].Graphql.User.EdgeFollowedBy.Count),
			Likes: []int64{},
			Comments: []int64{},
		}
		if res.EntryData.ProfilePage[0].Graphql.User.IsPrivate {
			return nil
		}
		totalPosts := len(res.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges)
		likes := make([]int64, totalPosts)
		comments := make([]int64, totalPosts)
		for idx, p := range res.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges {
			likes[idx] = int64(p.Node.EdgeLikedBy.Count)
			comments[idx] = int64(p.Node.EdgeMediaToComment.Count)
		}
		eng.Likes = likes
		eng.Comments = comments
		return nil
	})
	return
}
