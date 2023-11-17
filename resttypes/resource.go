package resttypes

type Resource struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Year         int    `json:"year"`
		Color        string `json:"color"`
		PantoneValue string `json:"pantone_value"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}
