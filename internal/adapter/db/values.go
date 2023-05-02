package db

import "time"

const (
	defaultPagination       = 10
	defaultURI              = "localhost:27017"
	defaultDbName           = "images"
	defaultColName          = "info"
	defaultCheckingDuration = time.Minute * 5
)
