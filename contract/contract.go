package contract

type Link interface {
	Title() string
	Url() string
	Raw() string
	OldTitleUrl() string
	IsRepoUrl() bool
	ApiEndpoint() string
}

type Readme interface {
	Links() []Link
	Raw() string
	GetStars() map[string]int
}

type Puller interface {
	Pull(url string) PullResult
}

type PullResult interface {
	Stars() int
	Ok() bool
}
