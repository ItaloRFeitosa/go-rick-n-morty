package alert

type Alert interface {
	String() string
}

type Data struct {
	Title   string
	Message string
}
