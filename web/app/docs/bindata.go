// Code generated for package docs by go-bindata DO NOT EDIT. (@generated)
// sources:
// docs/swagger.json
// docs/swagger.yaml
package docs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _swaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x6e\xdc\x44\x14\xbe\xef\x53\x8c\x0c\x52\x6f\x2a\x7b\xb3\x14\x09\x45\x42\x02\x4a\xa8\x56\xd0\xb4\x6a\x4a\x6f\x68\xb5\x9a\xb5\xcf\x7a\xa7\xb2\x67\xcc\xcc\x78\xd3\x28\x8a\x84\x10\x3f\xa9\x50\x09\x48\x29\x11\x25\x94\x72\x11\x1a\x55\x90\xf4\xa6\x28\xa4\x29\x79\x99\xb5\x37\x7d\x0b\x34\xb3\xde\xec\x9f\xbd\xf5\x26\x8e\x82\xa0\x77\xbb\x33\x67\xbe\x39\xdf\xf9\x99\x6f\xc6\x8b\x67\x10\x42\xc8\x10\x76\x03\x7c\x10\xc6\x34\xfa\x44\x0f\xe8\xc1\x86\x94\x81\x30\xf4\xff\x9b\xe7\x12\xbb\x79\xec\xba\xc0\x8d\x69\x64\x94\xcd\x92\x91\x8c\x12\x5a\x67\xc6\x34\x5a\xec\x2d\x75\x40\xd8\x9c\x04\x92\x30\xaa\x6c\xdf\x45\x73\x20\x11\xab\x23\x05\x89\x70\x40\x50\x9d\x71\x54\xc3\xd4\x0d\x7d\x62\x26\x30\x7a\xa1\x24\xd2\x03\xb5\xa4\xe6\xfa\x6f\xbc\xa5\x4d\x05\xf0\x26\xf0\x01\x23\xe0\xbe\xb8\x5c\x9f\x03\xde\x24\xb6\xb6\x9e\x65\xfd\xf3\x36\xa3\x12\xdb\x72\xc0\x25\x3d\x41\xb1\xaf\xcd\xaf\x71\xe2\x97\xa7\xfa\x96\xe8\xd9\x90\x7b\x6a\x52\xd3\x9e\xb6\x2c\x97\xc8\x46\x58\x33\x6d\xe6\x5b\xe9\xf6\xe0\x63\xa2\x57\x90\x77\xa4\x36\x30\x7d\x30\x0e\x4d\x96\xfa\x1c\xf2\x88\x0d\x54\x40\xb6\x43\x97\x2a\xd7\x26\xf5\xc6\xd2\x11\xb2\x6a\x1e\xab\x59\x0e\x34\xad\x8f\x2a\x17\x66\x66\xe7\x66\xd2\x1d\x68\x02\x17\x49\x2e\x4a\x66\xc9\x9c\xea\x58\x25\x16\x46\x83\x09\x15\x2c\x03\x07\xc4\xd4\xa8\x6a\x9b\x6e\x72\x6b\x58\xc0\x15\x2c\x1b\xca\xc0\xea\x0e\x06\x58\x36\xc4\x60\xca\x95\x3f\xa6\x6c\x5a\xcd\x29\xcb\xc6\x1e\x50\x07\x73\x6b\x31\x14\xc0\xab\xc4\x59\x1a\x65\xee\xc2\x68\x7e\xd2\x4a\x27\x7e\xf8\x57\x7c\x77\x2b\x5a\xdf\x6c\xaf\x7f\xd3\xbe\xff\x45\xfb\xde\xe3\xe8\xce\xa3\xf6\xea\x2f\xf1\xf2\x77\xc4\x16\xf1\xc3\xbd\x68\x6f\x25\x5e\xdb\x88\xbe\xfd\xea\x06\x8d\x7e\xfb\x3c\x7e\xb0\x1e\x3f\x5d\x39\x78\xb4\x1c\xfd\xb8\x79\xb0\xb5\xff\x62\x6d\x2b\x5e\xfb\xb3\xb5\x77\xff\x60\x7f\x35\xfa\xe9\x41\x7b\x77\x3b\xfe\xe1\xeb\x78\xfd\xf7\xf8\xde\x93\xf8\xee\xd6\x0d\x1a\xad\x6f\xb6\x9e\xef\xb7\x57\x37\x3b\x10\x38\x08\xa2\xed\xbd\xe8\xcb\x8d\xfe\x45\x9d\xa9\xce\x8a\xa1\x24\x75\x62\xc1\x99\x13\xda\x43\xcd\x33\x60\x21\xe1\xb6\x3c\x0c\x4a\x0a\x84\x36\xc2\x41\xe0\x11\x1b\x2b\xde\xd6\x2d\xc1\xa8\x31\x62\x76\x33\x65\x73\x11\xfa\x3e\xe6\x0b\x2a\x52\x2e\x50\xe0\x58\x02\xc2\xa8\xbb\x17\xaa\x73\xe6\xa3\x79\x2c\xed\x06\xa1\x2e\xb2\x99\xe7\x81\xad\x23\x9b\x02\xc5\x02\xb5\x9c\x30\x5a\x71\x14\x5c\x77\xd5\x85\x04\xeb\xfa\x70\xfd\x27\x85\xc0\xb1\x0f\x12\x78\x36\xfd\xd1\x1c\xf7\x02\xb3\x10\xe8\xf2\x17\x92\x13\xea\x66\x44\x06\xa5\x54\x45\x52\x57\xe7\x50\xb4\xb2\xdd\x7a\xb6\xa1\xd2\xb8\xf3\xb8\xb5\xbb\xdb\xda\x79\xf6\xe2\xd7\xa7\xaa\x4e\x9e\x7f\x1f\xfd\xfc\x59\x6b\xe7\x8f\xf8\xce\xdf\xd1\xf2\x93\x71\xd0\xdd\x1e\x4c\x30\xc7\x99\x12\xbd\xb9\x2a\xfe\x71\x56\x1c\x3e\x0d\x09\x07\x15\x45\xc9\x43\x48\x35\x5c\xca\x95\x5d\x0e\x22\x60\x54\x80\x48\x6d\x15\x6d\x52\x2e\x95\x32\x27\x51\x4a\xe4\xf2\xd4\x22\x3a\x94\x04\x3c\x16\x1b\xa5\xe4\x30\xd3\x78\x94\x31\x1a\x3c\xa4\x06\x50\xcf\x97\xce\x4f\x44\x6b\x96\x49\xf4\x01\x0b\xe9\xd8\xf4\xe5\xa5\xf4\x3a\x87\xba\x02\x7d\xcd\x72\xa0\x4e\x28\x51\x9b\x08\x8b\x83\x30\x67\x38\x67\xbc\x30\x8e\xe5\xf2\x44\x1c\x3f\xa6\x01\x67\x36\x08\x81\x6b\x1e\xa0\x19\x2a\x89\x5c\x38\x59\xba\xd7\xb1\x47\x1c\x7d\x26\x14\x4a\xfc\xcd\xd2\x64\xc4\xdf\xc3\x0e\xba\x88\x25\xcc\xe3\x13\xe6\x7b\x34\x96\x23\xa3\x83\x23\x4b\xa9\x8a\x7c\x28\x97\x65\xab\x49\x60\xbe\x4a\x02\x4b\x84\xb5\x5b\x60\x4b\x6b\x31\xf9\x71\x3c\xe1\xbc\x08\x12\x71\xf0\xb0\x04\x07\xcd\x75\x00\x05\xc2\xd4\x41\xb2\x01\x84\x77\xa6\x14\xf1\xa3\x49\x5a\x01\x6a\x25\x51\xc2\xb3\xe7\x0b\xf2\x71\x90\x43\x9a\x54\xc0\x2a\x57\x4e\x4e\x8e\x08\x95\xe0\xc2\xd8\xd3\x71\xf8\x82\x6b\xdb\x2c\xa4\x12\x55\xde\xcf\xa3\x34\xbd\xfc\xfe\xc7\xc5\xe6\xf2\x87\x27\xd7\xaf\xaa\x08\x48\x60\x26\xb1\xbc\x84\x83\xab\x20\x5e\xa9\xcf\x04\x1c\xff\xbf\xea\x33\x59\x11\x57\xa8\x04\x4e\xb1\x87\xe6\xf4\x43\x14\x75\x9c\xf9\x17\x26\x3a\xbf\x0e\x0d\xbc\xfe\xfa\x36\x1f\x7c\xd1\xf5\x3c\x19\x11\xa1\xee\x49\xc9\x74\xf3\x0d\x3f\x5d\x03\xae\x0e\x6c\x49\x32\x8e\x12\xc3\x57\x55\xe4\x8e\x3e\x87\x47\xe0\xc7\xbe\x0b\x0c\xb8\x8d\xfd\xa0\xf3\xb5\xc0\x65\x12\x61\x04\xca\x59\x44\x68\xf2\xc9\x00\x8d\x46\x30\xa5\x26\x0c\x21\xb1\x0c\xc7\x1c\x7a\x13\x3b\x03\xe9\xd9\xcb\x73\x2b\x48\x2b\xfb\x82\x83\x5f\x27\xe0\x39\xd5\x44\x8a\x8a\xa2\xdc\xa7\x68\xb9\x62\x5e\x7c\x05\x88\x06\x0b\x3d\x07\xd5\xd4\xf3\xb7\xab\xdf\x47\x4a\x41\xa2\x2b\xe0\xa4\xb8\x77\xcc\xd0\x13\x27\x37\xe1\x7c\x51\xc4\xc1\xcb\x01\xb3\x83\x91\x7a\x07\xe8\xdc\xc4\x8a\xf5\x53\xb0\x90\xdb\x39\x92\x3d\x91\xab\x12\xf3\xac\xab\x70\x4e\xd0\xfc\xc5\x40\x99\x73\x2a\xc5\x30\x51\x40\x88\x3f\xb6\xa3\x86\x54\xed\x6d\x74\xd6\xc3\x84\x9a\xc9\x33\x84\xf8\xae\x45\x59\x95\xd8\x8c\x56\x93\x66\x36\x03\xea\x9e\xcd\xea\xbc\xa3\x94\xc1\xc8\x67\xe2\x54\xc4\x24\x9a\xa7\xd4\x00\x13\x9d\x8b\xb9\x11\xab\x76\xd1\x1d\xd5\x3b\x6f\x8b\xe5\xdf\x05\x4e\x16\x1f\xd5\xe5\xfc\xad\xd5\xbb\xbf\x17\xdd\x60\xea\x04\xcf\x21\xeb\x98\xf3\xcc\xcf\x0a\x06\x91\xe0\x67\x63\xa0\x97\xbe\x4e\xb4\x8a\xe4\xbd\xad\xa5\x15\x0f\x73\x4e\x9d\x83\x3e\xfc\x8e\x7b\xe3\x1c\x97\xf8\xab\x29\x14\x8f\x99\x7b\x07\xcb\xec\x0b\xf8\x78\xec\x7c\x41\xe9\x2b\xdb\x13\xb9\xf0\x14\x7a\x73\x9d\x5c\x02\x3b\x0f\x84\x33\x4b\xff\x04\x00\x00\xff\xff\x9f\x8a\xf1\x4d\x1f\x1c\x00\x00")

func swaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_swaggerJson,
		"swagger.json",
	)
}

func swaggerJson() (*asset, error) {
	bytes, err := swaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swagger.json", size: 7199, mode: os.FileMode(438), modTime: time.Unix(1576780938, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\x5d\x6f\x1b\x45\x14\x7d\xf7\xaf\xb8\x32\x48\x7e\xa1\xbb\x8e\x29\x12\x5a\x09\x89\x52\x42\x65\x41\x3f\x44\x4a\x5f\xad\xf1\xee\xf5\x7a\xaa\xdd\x99\x65\x66\xd6\x69\x54\x90\x10\xe2\x23\x15\x2a\x01\x29\x25\xa2\x84\x52\x1e\x42\xa3\x0a\x92\xbe\x14\x85\x34\x25\x7f\xc6\xbb\x4e\xff\x05\x9a\xdd\xf5\x7a\xd7\x71\xd6\x69\xd4\x37\x1e\x77\x3c\xf7\xdc\x73\xee\x9c\x7b\x67\xdc\x25\x12\xaf\x11\xd5\xb7\xc0\xac\x39\xd8\xa3\x8c\x2a\xca\x99\xb4\x6a\x00\x02\xa5\xb1\x28\x04\x17\xfa\x03\x20\x10\x3c\x40\xa1\x28\xca\xf4\x1b\xc0\x47\x29\x89\x8b\xe3\x4f\x00\xbc\x45\xfc\xc0\x43\x0b\x1a\x2e\x57\x40\x00\x75\x34\x50\x06\x12\xc5\x00\x05\x34\xf2\x9d\x6a\x25\x40\x0b\xa4\x12\x94\xb9\xd9\xa2\x54\x44\x85\x72\x06\x58\x82\x72\x72\x64\xba\xc0\xbb\x37\xd1\x56\x19\xeb\x1b\xc4\xa3\x0e\xd1\x42\x2a\xf9\xf7\x28\x7a\x4e\x87\x11\x7f\x96\x04\x19\x26\x88\x1d\xea\x54\x91\x3e\xb9\x04\xb2\xcf\x43\xcf\x81\x2e\x02\x01\xca\x14\xba\x78\x7a\x0d\x03\x8a\xcb\x34\x30\xd0\x19\x23\x1f\xe7\x4e\x1d\xab\x92\x17\x09\xa6\x7f\x2f\x93\x10\xe8\x25\x05\xaa\x44\x91\x3c\x14\x36\x56\x03\x29\x22\x5c\x54\x55\x7b\x66\x8b\x63\xdc\x79\x09\x71\xe5\x9c\xd4\x2f\xd5\xdc\x41\x69\x0b\x1a\x24\x72\xe0\x1d\x68\x78\x84\x32\xa3\xeb\xfa\x86\x1a\x98\xd4\x77\x4d\xc6\x3b\xd4\xe6\xac\x93\x1d\xa9\x11\x30\xb7\xd2\x89\x94\xf5\xf8\x34\x81\x9c\xfe\xa9\xaa\x5b\xf6\xd4\x8c\x14\x7a\x43\xc7\x9e\x53\xfd\xdc\x81\xd5\xc9\xc6\xfb\x92\xdf\x5e\xd6\x63\x59\xf0\xe5\xb1\xa2\xe3\x87\xa1\x5d\x58\x68\x4b\xaa\xd0\x2f\x7c\x02\xbc\x2e\xb0\x67\x41\xe3\x35\xb3\x30\x3e\xcc\x82\x83\xa7\x4b\x4d\x84\x20\x2b\xe3\x32\x70\xe7\xcc\xd8\x3a\xf6\x64\xec\x79\x6a\x3f\x1e\xe7\x3d\x2e\xd8\x21\x8a\x4c\x48\x54\x51\x98\xc0\x4d\x13\x29\xdb\x65\x7a\x48\x9c\x62\xfa\xcd\x69\xa3\x3e\x97\xca\x02\x12\x50\xed\xf3\x37\xdf\x36\x6c\xee\xd7\xc6\xb6\xb5\x39\x53\xc4\xce\x5a\x12\x7d\x42\x3d\x0b\xe8\xbb\x4a\x50\xbf\xb5\x60\xf8\x58\xcb\x0d\x0a\xd7\x93\xb5\x64\x21\x14\x9e\x05\x7d\xa5\x02\x69\x99\xa6\x4b\x55\x3f\xec\x6a\x50\x33\xdf\x52\x6a\xb2\x0b\xb0\x84\x0a\x78\x2f\x89\xd0\x34\xa0\xc7\x05\x74\x09\x73\x43\x9f\x1a\x35\x00\x8f\xda\xc8\x64\x26\x39\x4d\x76\xb9\x7d\xfd\x14\x99\xcc\x44\x8f\xd9\xf5\x78\xd7\x74\x70\x60\x7e\xd4\xbe\xb8\x78\x65\x69\xb1\x06\xa0\x50\xf8\xf2\x6a\x6f\x09\xc5\x80\xda\x68\x41\xfd\x0a\xaf\xeb\x65\xaa\xf4\xb8\x4d\xc2\x12\x26\xe9\x5d\xa3\xcf\x1c\x85\x4c\xd8\x36\x8d\xa6\xb1\x50\x0b\x88\xea\x27\x05\x36\xb3\xd1\x30\x58\x30\x6d\xe2\x21\x73\x88\x30\x6f\x87\x12\x45\x87\x3a\x9f\xa7\x8c\x0b\x13\xad\xa4\xfb\xb3\x73\xf9\x01\xc5\x0f\xff\x89\xef\xee\x44\x9b\xdb\xa3\xcd\xef\x46\xf7\xbf\x1a\xdd\x7b\x1c\xdd\x79\x34\x5a\xff\x2d\x5e\xfd\x81\xda\x32\x7e\x78\x10\x1d\xac\xc5\x1b\x5b\xd1\xf7\xdf\xe4\x21\xd1\x1f\x5f\xc6\x0f\x36\xe3\xa7\x6b\x47\x8f\x56\xa3\x9f\xb7\x8f\x76\x0e\x5f\x6c\xec\xc4\x1b\x7f\x0f\x0f\xee\x1f\x1d\xae\x47\xbf\x3c\x18\xed\xef\xc6\x3f\x7d\x1b\x6f\xfe\x19\xdf\x7b\x12\xdf\xdd\x99\x04\x6e\x6e\x0f\x9f\x1f\x8e\xd6\xb7\x53\x40\x12\x04\xd1\xee\x41\xf4\xf5\x56\x31\x36\xfd\xa9\x14\xa8\x8d\x9d\x0c\xf9\xb6\x63\xc1\x32\x51\x76\x9f\x32\xf7\x62\x26\xf9\xc6\x42\xb6\x2b\x20\x82\xf8\xa8\x50\xe4\xee\x3b\x57\x16\x9d\x95\xe6\x0d\x88\xd6\x76\x87\xcf\xb6\x34\x99\xbd\xc7\xc3\xfd\xfd\xe1\xde\xb3\x17\xbf\x3f\xd5\xda\x9f\xff\x18\xfd\xfa\xc5\x70\xef\xaf\xf8\xce\xbf\xd1\xea\x93\x49\x3b\x33\x0b\x74\xd5\xf3\x85\xd4\x08\x19\x5e\xbe\x2a\xf0\xd3\x90\x0a\x74\x2c\x50\x22\xc4\xaa\x26\x09\x04\x77\x42\x1b\x0b\x44\x15\xde\x52\xf9\x29\xe6\xab\x24\x08\x3c\x6a\x27\xd2\xcd\x9b\x92\xb3\xfc\xce\x93\x01\x67\xb2\x38\x71\xea\xad\x66\xb3\x5e\x9c\x38\x25\xe5\xb3\xd0\x93\x66\xb5\xfb\xe8\x93\x62\xd8\x4c\xba\x00\xf5\xf3\xcd\xf3\x27\xc3\x5f\xe1\x0a\x3e\xe0\x21\x73\xe6\x40\xcf\x1c\x41\xf9\xeb\xac\x51\xc8\xd6\x6a\x9d\x9c\xed\x13\x16\x08\x6e\xeb\x51\xd4\xf5\x10\x16\x99\xa2\x6a\xe5\xac\x89\xa7\x1e\x58\x05\x0a\x6f\x35\x2b\x28\xbc\x47\x1c\xb8\x44\x14\x2e\x93\x33\x67\x2e\xe5\x93\xa1\xef\x13\xb1\x62\x81\x8b\x4c\x5b\x5d\xbf\xb5\xc6\xc7\x05\x3d\xc1\xfd\xdc\xf5\x60\x73\xcf\x43\x5b\xe3\x14\x27\x40\x2b\x99\xe4\x1d\x1a\x98\xd9\x28\x37\x6f\x4f\x6e\xdd\x39\xb3\xe0\x12\xaa\xf4\x15\x85\x0e\x2c\xa5\x41\x12\x08\x73\x40\xf5\x91\x8a\xfc\x81\x25\x67\x35\xa3\xce\xda\xbe\x76\xda\x06\xbc\x60\xdb\x3c\x64\x0a\xda\xef\xcf\x6b\xad\x19\x8f\xd6\xca\xee\x2a\x3f\x24\x8e\xb7\xd7\xab\x6a\xa4\xab\x1f\x9e\xe5\xbc\x67\x5d\xda\x8d\xff\x75\x6f\x55\x94\xb8\xcd\x14\x0a\x46\x3c\x58\x4a\xff\x6b\x2d\x96\xfe\x35\xbd\xa2\x2e\x53\x63\x87\x4d\xec\xad\x1f\xc3\xb5\x04\x5c\xdb\xe1\x5c\x7a\xb3\xd7\xe4\x32\x71\x5d\x14\x16\xd4\x5b\x46\xb3\x5e\xfb\x2f\x00\x00\xff\xff\x88\xdd\x5f\x5d\x66\x0e\x00\x00")

func swaggerYamlBytes() ([]byte, error) {
	return bindataRead(
		_swaggerYaml,
		"swagger.yaml",
	)
}

func swaggerYaml() (*asset, error) {
	bytes, err := swaggerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swagger.yaml", size: 3686, mode: os.FileMode(438), modTime: time.Unix(1576780938, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"swagger.json": swaggerJson,
	"swagger.yaml": swaggerYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"swagger.json": &bintree{swaggerJson, map[string]*bintree{}},
	"swagger.yaml": &bintree{swaggerYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
