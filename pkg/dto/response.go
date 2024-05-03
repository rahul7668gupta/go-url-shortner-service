package dto

type Reponse struct {
	ShortenedUrl string `json:"shortenedUrl"`
}

type Metrics struct {
	DomainName string `json:"domainName"`
	Count      int64  `json:"count"`
}
