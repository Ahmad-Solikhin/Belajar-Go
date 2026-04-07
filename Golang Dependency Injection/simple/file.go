package simple

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.CLose()
	}
}

func (f *File) CLose() {
	fmt.Println("Close file", f.Name)
}
