package multipart

import (
	"time"

	protocol "github.com/nullexp/httpzen/protocol"
	file "github.com/nullexp/httpzen/protocol/file"
)

type filePart struct {
	file     file.File
	partName string
}

func (f *filePart) Close() error {
	return f.file.Close()
}

func (f *filePart) Read(p []byte) (n int, err error) {
	return f.file.Read(p)
}

func (f *filePart) GetPartName() string {
	return f.partName
}

func (f *filePart) GetContentType() string {
	return f.file.GetMimeType()
}

func (f *filePart) GetFilename() string {
	return f.file.GetFilename()
}

func (f *filePart) GetLastModifiedDate() time.Time {
	return f.file.GetLastModifiedDate()
}

func (f *filePart) GetMimeType() string {
	return f.file.GetMimeType()
}

func NewFilePart(f file.File, partName string) protocol.FileMultipart {
	return &filePart{file: f, partName: partName}
}
