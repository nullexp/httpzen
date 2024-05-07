package multipart

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/nullexp/httpzen/misc"
	http "github.com/nullexp/httpzen/protocol"
)

type FileDefinition struct {
	Name           string
	MinSize        misc.ByteSize
	MaxSize        misc.ByteSize
	Optional       bool
	Single         bool
	SupportedTypes []string
}

func (f *FileDefinition) GetObject() interface{} {
	return &multipart.FileHeader{}
}

func (f FileDefinition) IsOptional() bool {
	return f.Optional
}

func (f FileDefinition) GetPartName() string {
	return f.Name
}

func (f FileDefinition) IsSingle() bool {
	return f.Single
}

func (f FileDefinition) GetSupportedTypes() []string {
	return f.SupportedTypes
}

const (
	TooSmall        = "less Than Minimum size. must be atleast "
	TooBig          = "exeeding Max size. must be lest than "
	NoName          = "filename should not be empty"
	CantRecognized  = "could not recognized the file type"
	InvalidFileType = "invalid file type"
)

func (f FileDefinition) Verify(header http.FileHeader) error {
	if f.MinSize >= 0 {
		if header.GetSize() < int64(f.MinSize) {
			return errors.New(TooSmall + fmt.Sprint(f.MinSize+1))
		}
	}

	if f.MaxSize > 0 {
		if header.GetSize() > int64(f.MaxSize) {
			return errors.New(TooBig + fmt.Sprint(f.MaxSize+1))
		}
	}

	if header.GetFilename() == "" {
		return errors.New(NoName)
	}

	if len(f.SupportedTypes) != 0 {
		freader, err := header.OpenFile()
		if err != nil {
			return errors.New(CantRecognized)
		}
		defer freader.Close()

		t, _ := filetype.MatchReader(freader)

		if t == types.Unknown {
			return errors.New(CantRecognized)
		}

		found := false
		for _, v := range f.SupportedTypes {
			if v == t.MIME.Value {
				found = true
				break
			}
		}

		if !found {
			return errors.New(InvalidFileType)
		}
	}

	return nil
}
