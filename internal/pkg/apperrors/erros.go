package apperrors

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrDbServer = errors.New("database serve is down")

	// Auth header errors
	ErrHeaderMissing = errors.New("authorization header missing")
	ErrInvalidToken  = errors.New("invalid or expired token")
	ErrIdMissmatch   = errors.New("user ID mismatch")

	// User errors
	ErrDuplicateUsername  = errors.New("username already exists")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrDuplicateUid       = errors.New("uid already exists")
	ErrInvalidData        = errors.New("invalid user data provided")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user with given ID not found")
	ErrInvalidUserID      = errors.New("invalid user ID: must be a positive integer")

	// Equipment errors
	ErrDurationTooShort     = errors.New("duration too short")
	ErrQuantityNotAvailable = errors.New("requested quantity not available")
	ErrFailedToCreate       = errors.New("failed to create equipment")
	ErrInvalidQuantity      = errors.New("invalid quantity: must be greater than zero")
	ErrEquipmentNotFound    = errors.New("equipment not found")
	ErrNotAnOwner           = errors.New("not an owner of the equipment, operation not allowed")

	// DB errors
	ErrDbFetching = errors.New("error while fetching DB data")
	ErrDbParsing  = errors.New("error while parsing db data")
	ErrDbExce     = errors.New("error while inserting data into DB")
	ErrDbScan     = errors.New("error while scannig data from db to model")
	ErrNoData     = errors.New("no data found")
	ErrDbDelete   = errors.New("error while deleting data from DB")

	ErrInvalidReqBody = errors.New("invalid request body")
	ErrPathParam      = errors.New("invalid path ID, must be a number")
	ErrAtoi           = errors.New("error while converting id param form string into int")
)
