package blog

import "time"

type BlogPost struct {
	ID				int
	Title			string
	Content		string
	CreatedAt time.Time
}
