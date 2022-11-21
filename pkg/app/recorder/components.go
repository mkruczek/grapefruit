package recorder

type HttpClient interface {
	ListenAndServe() error
}
