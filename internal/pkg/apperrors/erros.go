package apperrors

import "errors"

var (
	ErrDurationTooShort     = errors.New("duration too short")
	ErrQuantityNotAvailable = errors.New("requested quantity not available")

	ErrDbFetching = errors.New("error while fetching DB data")
	ErrDbParsing  = errors.New("error while parsing db data")
	ErrDbExce     = errors.New("error while inserting data into DB")

	ErrPathValueParsing = errors.New("error while paring path value")
)
