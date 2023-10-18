package loadbalancer

import "time"

type RoundRobinEntity struct {
	Url           string    `json:"url"`
	LastRequestAt time.Time `json:"lastRequestAt"`
}

type RRCacheEntity struct {
	Hosts []RoundRobinEntity `json:"hosts"`
}
