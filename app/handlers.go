package app

type RespVisits struct {
	RespVisits []RespVisit `json:"visits"`
}
type RespVisit struct {
	Mark       int    `json:"mark"`
	Visited_at int    `json:"visited_at"`
	Place      string `json:"place"`
}

func (v RespVisits) Len() int {
	return len(v.RespVisits)
}
func (v RespVisits) Swap(i, j int) {
	v.RespVisits[i], v.RespVisits[j] = v.RespVisits[j], v.RespVisits[i]
}
func (v RespVisits) Less(i, j int) bool {
	return v.RespVisits[i].Visited_at < v.RespVisits[j].Visited_at
}

type VisIDs struct {
	IDs []int `json:"ids,omitempty"`
}
