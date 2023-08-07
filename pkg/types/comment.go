package types

import "time"

type Comment struct {
	Timestamp time.Time
	Author    ID
	Markdown  string
}
