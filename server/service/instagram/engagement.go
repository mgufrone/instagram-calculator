package instagram

import "log"

type Engagement struct {
	Likes     []int64
	Comments  []int64
	Followers int64
}

func (e *Engagement) sumEngagement() (res int64, totalPosts int64) {
	totalPosts = int64(len(e.Likes))
	log.Println("recorded", e.Comments, e.Likes)
	for i, v := range e.Likes {
		total := v + e.Comments[i]
		log.Println("engagement rate", v, e.Comments[i], total, e.Followers, float64(total/e.Followers))
		res += total
	}
	return
}

func (e *Engagement) Rate() float64 {
	if len(e.Likes) == 0 {
		return 0
	}
	total, posts := e.sumEngagement()
	rate := float64(total) / float64(posts * e.Followers)
	return rate
}
