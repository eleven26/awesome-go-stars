package contract

type Link interface {
	Title() string
	Url() string
	Raw() string
	OldMarkdownLink() string
	NewMarkdownLink(star int) string
	IsRepoUrl() bool
	ApiEndpoint() string
}

type Puller interface {
	Pull(url string) PullResult
}

type PullResult interface {
	Stars() int
	Ok() bool
}

type Handler interface {
	Handle() error
}
