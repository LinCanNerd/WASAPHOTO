package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validUsername(identifier string) bool {
	var trimmedId = strings.TrimSpace(identifier)
	return len(identifier) >= 3 && len(identifier) <= 16 && trimmedId != "" && !strings.ContainsAny(trimmedId, "?_")
}

var Root = filepath.Join("/tmp", "photos")

func (p *Photo) path() (path string) {
	return filepath.Join(Root, fmt.Sprintf("%d.%s", p.ID, p.Format))
}

func CreatePhotoFile(photo Photo, binaryImage []byte) (err error) {
	err = os.MkdirAll(Root, 0755)
	if err != nil {
		return
	}
	file, err := os.Create(photo.path())
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write(binaryImage)
	return
}

func DeletePhotoFile(photo Photo) (err error) {
	err = os.Remove(photo.path())
	return
}
