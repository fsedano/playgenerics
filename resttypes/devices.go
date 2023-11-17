package resttypes

type Device struct {
	Data []struct {
		Avatar    string `json:"avatar"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		ID        int    `json:"id"`
		LastName  string `json:"last_name"`
	} `json:"data"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Support struct {
		Text string `json:"text"`
		URL  string `json:"url"`
	} `json:"support"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
