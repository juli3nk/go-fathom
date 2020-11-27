package fathom

import (
	"net/http"
)

type Config struct {
	Cookies []*http.Cookie
}

type Data struct {
	ChartData       []ChartData         `json:"chart_data"`
	SiteStats       map[string]SiteStat `json:"site_stats"`
	PageStats       []PageStat          `json:"page_stats"`
	ReferrerStats   []StandardStat      `json:"referrer_stats"`
	DeviceTypeStats []StandardStat      `json:"device_type_stats"`
	BrowserStats    []StandardStat      `json:"browser_stats"`
	CountryStats    []StandardStat      `json:"country_stats"`
	GoalStats       []StandardStat      `json:"goal_stats"`
}

type ChartData struct {
	PageViews string `json:"pageviews"`
	Visits    string `json:"visits"`
	Date      string `json:"date"`
}

type SiteStat struct {
	PageViews   string  `json:"pageviews"`
	Visits      string  `json:"visits"`
	AvgDuration float64 `json:"avg_duration"`
	BounceRate  float64 `json:"bounce_rate"`
	Period      string  `json:"period"`
}

type PageStat struct {
	Url      string `json:"url"`
	PathName string `json:"pathname"`
	Uniques  string `json:"uniques"`
	Views    string `json:"views"`
}

type StandardStat struct {
	Label string `json:""label"`
	Total string `json:"total"`
}
