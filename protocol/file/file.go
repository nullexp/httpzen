package protocol

import (
	"io"
	"time"
)

type File interface {
	io.Reader
	io.Closer
	GetFilename() string
	GetMimeType() string
	GetLastModifiedDate() time.Time
}

type SeekerFile interface {
	io.Seeker
	File
}

type readableFile struct {
	readcloser io.ReadCloser
	fileName   string
	mimeType   string
	lmd        time.Time
}

func (r *readableFile) Read(p []byte) (n int, err error) {
	return r.readcloser.Read(p)
}

func (r *readableFile) Close() error {
	return r.readcloser.Close()
}

func (r *readableFile) GetFilename() string {
	return r.fileName
}

func (r *readableFile) GetMimeType() string {
	return r.mimeType
}

func (r *readableFile) GetLastModifiedDate() time.Time {
	return r.lmd
}

func NewFile(rc io.ReadCloser, fname, mimetype string, lmd time.Time) File {
	return &readableFile{readcloser: rc, fileName: fname, mimeType: mimetype, lmd: lmd}
}
