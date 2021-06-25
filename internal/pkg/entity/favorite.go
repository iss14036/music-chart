package entity

type Favorite struct {
	UserID  int `json:"user_id"`
	MusicID int `json:"music_id"`
}

type FavoriteResult struct {
	List []Music `json:"list"`
	Pagination
}

type FavoriteFilter struct {
	UserID string
	DataTable FilterDataTable
}