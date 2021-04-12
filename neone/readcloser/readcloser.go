package readcloser
import "io"

type ReadCloser interface {
	io.Reader
	Close() error
}
