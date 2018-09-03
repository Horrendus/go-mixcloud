package mixcloud

import "time"

type ListOptions struct {
	Limit int		`url:"limit"`
	Offset int		`url:"offset,omitempty"`
	Since time.Time	`url:"since,omitempty,unix"`
	Until time.Time `url:"until,omitempty,unix"`
}
