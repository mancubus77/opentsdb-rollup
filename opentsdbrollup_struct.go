package main

import (
	"sync"
)

// RoundRobin struct
type RoundRobin struct {
	sync.Mutex

	current int
	pool    []string
}

// Lookup struct for search lookup api response
type Lookup struct {
	Type    string        `json:"type"`
	Metric  string        `json:"metric"`
	Tags    []interface{} `json:"tags"`
	Limit   int           `json:"limit"`
	Time    float64       `json:"time"`
	Results []struct {
		Tsuid  string            `json:"tsuid"`
		Metric string            `json:"metric"`
		Tags   map[string]string `json:"tags"`
	} `json:"results"`
	StartIndex   int `json:"startIndex"`
	TotalResults int `json:"totalResults"`
}

// Rollup payload struct
type Rollup struct {
	//Aggregator string            `json:"aggregator"`
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	Metric    string  `json:"metric"`
	//Interval   string            `json:"interval"`
	Tags map[string]string `json:"tags"`
}

// SubQuery data to rollup
type SubQuery struct {
	Aggregator  string   `json:"aggregator"`
	Metric      string   `json:"metric,omitempty"`
	RollupTable string   `json:"rollupTable,omitempty"`
	Downsample  string   `json:"downsample"`
	Tsuids      []string `json:"tsuids,omitempty"`
}

// Query data to rollup
type Query struct {
	Start   int64      `json:"start"`
	End     int64      `json:"end"`
	Queries []SubQuery `json:"queries"`
}

// QueryRespItem acts as the implementation of Response in the /api/query scene.
type QueryRespItem []struct {
	Metric        string             `json:"metric"`
	AggregateTags []string           `json:"aggregateTags"`
	Tags          map[string]string  `json:"tags"`
	Dps           map[string]float64 `json:"dps"`
}

// RollupResponse struct for post output
type RollupResponse struct {
	Success int
	Failed  int
}

// tomlConfig file
type tomlConfig struct {
	Title    string
	Servers  servers
	Settings settings
	API      api
	Metric   metric
	Rest     rest
}

// Gettslist channel struct
type Gettslist struct {
	jobid string
	list  []string
}

// Getrollup channel struct
type Getrollup struct {
	jobid  string
	rollup []Rollup
}

// servers
type servers struct {
	WriteEndpoint []string
	ReadEndpoint  []string
}

// api
type api struct {
	SuggestMax    string
	LookupLimit   string
	LookupUseMeta string
}

// settings
type settings struct {
	HoursPast    int
	Batch        int
	ReadWorkers  int
	WriteWorkers int
	Interval     string
	LogLevel     int
}

// rest
type rest struct {
	Debug            bool
	Timeout          int
	RetryCount       int
	RetryWaitTime    int
	RetryMaxWaitTime int
}

// metric
type metric struct {
	List []string
}
