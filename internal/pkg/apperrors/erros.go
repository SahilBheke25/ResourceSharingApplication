package apperrors

import "errors"

var (
	ErrDurationTooShort     = errors.New("duration too short")
	ErrQuantityNotAvailable = errors.New("requested quantity not available")

	ErrDbFetching = errors.New("error while fetching DB data")
	ErrDbParsing  = errors.New("error while parsing db data")
	ErrDbExce     = errors.New("error while inserting data into DB")
	ErrDbScan     = errors.New("error while scannig data from db to model")

	ErrPathValueParsing = errors.New("error while paring path value")
	ErrAtoi             = errors.New("error while converting equipment id param form string into int")
)
