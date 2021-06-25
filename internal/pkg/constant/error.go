package constant

import "errors"

var (
	ErrorDataTableNotValid = errors.New("DataTable request is not valid")
	ErrorMusicNotFound     = errors.New("Music not found")
	ErrorDataNotFound      = errors.New("Data not found")
)
