package main


type Downloader interface {
	Download() (err error)
}

type Archiver interface {
	Archive(names []string) (err error)
}