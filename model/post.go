package model

import (
	"time"
)

type Post struct {
	Id string
	Title string
	Date time.Time
	HtmlContent string
}