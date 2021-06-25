package entity

type Music struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Singer      string `json:"singer"`
	Duration    string `json:"duration"`
	Album       string `json:"album"`
	ReleaseYear string `json:"release_year"`
}

type MusicResult struct {
	List []Music `json:"list"`
	Pagination
}

type MusicFilter struct {
	ID        string
	DataTable FilterDataTable
}
