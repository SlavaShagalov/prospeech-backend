package files

import "context"

type File struct {
	Name string
	Data []byte
}

type Repository interface {
	Create(ctx context.Context, file *File) (location string, err error)
	Get(location string) (imgData []byte, err error)
	Update(ctx context.Context, location string, imgData []byte) (err error)
	Delete(location string) (err error)
}
