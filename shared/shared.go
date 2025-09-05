package shared

type Art struct {
	PrimaryImage string `json:"primaryImage"`
	Title string `json:"title"`
	IsHighlight bool `json:"isHighlight"`
	ObjectID int `json:"objectID"`
	Department string `json:"department"`
	Medium string `json:"medium"` // refers to materials that were used
	Country string `json:"country"`
	City string `json:"city"`
	ObjectDate string `json:"objectDate"`
	Period string `json:"period"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ArtistDisplayBio string `json:"artistDisplayBio"`
	ArtistWikidata_URL string `json:"artistWikidata_URL"`
	ObjectURL string `json:"objectURL"` // URL to objects page on metmuseum.org
	ObjectWikidata_URL string `json:"objectWikidata_URL"`
}

type Gallery struct {
	Total int `json:"total"`
	ObjectIDs []int `json:"objectIDs"`
}
