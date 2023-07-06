package models

import "time"

var BeginningOfTime = time.Date(
	0, time.January, 1, 0, 0, 0, 0, time.UTC,
)
var EndOfTime = time.Date(
	9999, time.December, 31, 23, 59, 59, 999, time.UTC,
)
