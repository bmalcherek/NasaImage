package models

// Model of A Picture Of the Day from NASA api

type APOD struct {
	Copyright       string
	Date            string
	Explanation     string
	Hdurl           string
	Media_type      string
	Service_version string
	Title           string
	Url             string
}
