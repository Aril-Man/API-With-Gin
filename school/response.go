package school

type SchoolResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   string `json:"class"`
	Major   string `json:"major"`
}
