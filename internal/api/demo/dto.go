package demo

type MovieRequestDto struct {
	Name string `json:"name"`
}

type MovieDto struct {
	Id          int    `json:"id"`
	Title       string `json:"original_title""`
	ReleaseDate string `json:"release_date"`
}

type MovieResp struct {
	Page    int        `json:"page"`
	Results []MovieDto `json:"results"`
}
