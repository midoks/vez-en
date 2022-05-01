// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../templates/about.tmpl
// ../../../templates/base/footer.tmpl
// ../../../templates/base/header.tmpl
// ../../../templates/home.tmpl
// ../../../templates/page/content.tmpl
// ../../../templates/sidebar/sidebar.tmpl
// ../../../templates/sidebar/sidebar_new.tmpl
// ../../../templates/soso.html
package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _aboutTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4a\x2c\x4e\xd5\xcf\x48\x4d\x4c\x49\x2d\x52\x52\xd0\xab\xad\xe5\xe2\xe2\xc2\x90\x4e\xcb\xcf\x2f\x81\x4a\x03\x02\x00\x00\xff\xff\xc7\xd7\xde\x9b\x3b\x00\x00\x00"

func aboutTmplBytes() ([]byte, error) {
	return bindataRead(
		_aboutTmpl,
		"about.tmpl",
	)
}

func aboutTmpl() (*asset, error) {
	bytes, err := aboutTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "about.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseFooterTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xc1\x8e\x9c\x38\x10\xbd\xf3\x15\xb5\xec\x6a\xe9\xd6\xc8\x30\x33\xdd\xa7\x1d\x60\xa5\x9d\x3d\x24\x52\xa4\x48\x51\x72\xc8\xd1\x40\x01\xee\x36\x36\xb1\x0b\x68\xd4\xea\x7f\x8f\x0c\x0d\x9d\x99\x28\x87\x8c\x0f\xa8\x28\xfb\xbd\x57\xf5\x5c\xf6\x3c\xcf\x8b\x0b\xd1\x43\x2e\xb9\xb5\x89\x9f\x6b\x45\x5c\x28\x34\x7e\xea\x01\xc4\xa5\xd6\x84\x66\xd9\x2c\x58\x29\xf1\x04\xee\xc3\x06\xc3\x5b\x38\x74\x96\x44\x39\x32\x87\x42\x45\x2c\x43\x1a\x10\x15\x70\x29\x2a\xc5\x04\x61\x63\x59\x8e\xca\x51\xb4\x23\xdb\x41\x33\xb2\x3d\x64\xda\x14\x68\x18\xe9\x76\xd2\x00\x78\xa9\x2f\x59\x53\xb0\x3d\x5c\xb5\x7e\x66\xba\x82\x00\x62\x0e\xb5\xc1\x32\xf1\x23\x7f\x41\x37\x99\x13\x41\xf6\x08\x4d\xe6\x68\xee\x81\xf0\x44\xac\xe9\x08\x8b\x39\x2c\x30\xd7\x86\x93\xd0\x8a\x29\xad\x10\x64\xcd\x1e\x56\x46\x80\xd8\xf6\xd5\x42\x96\x09\x1f\x06\x51\x50\x9d\xf8\xbb\x7b\x1f\x6a\x14\x55\x4d\x89\xff\xb8\xf7\xd3\xb8\xb3\x08\x27\x29\xd4\xf1\x9f\xb9\x84\x3f\x33\xad\xc9\x92\xe1\xad\x1f\xa5\x71\x64\xfb\x6a\xad\x32\xe2\x6b\x68\x5b\xae\x16\xf6\x5b\x61\x7e\xfa\x77\xae\xdb\xf1\x09\xce\xe7\xaf\xc8\xcd\xe5\x12\x47\xee\xe0\xd5\x9b\xa8\x10\x7d\xea\xcd\x71\x27\x17\xb4\xe2\x3d\xac\x56\xbd\xbe\x05\x54\x05\x48\x61\x89\x75\xca\xd2\x28\xb1\xb8\x9a\x79\x73\x4e\x8a\xd5\x31\xcb\x76\x7e\x1a\x8b\xa6\x5a\x7a\x7d\xdc\xbf\xe8\x15\xac\xc9\x13\x3f\x2a\x79\x2f\x72\xad\xc2\x56\x55\x53\x87\x52\x2c\xf5\x75\x72\x9a\x94\x68\x1e\x95\xd4\x5b\x2a\x8e\x6d\x6e\x44\x4b\x57\xbc\x25\x4e\x22\x8f\x0e\x36\x3a\x7c\xeb\xd0\x8c\x6c\x17\xee\xc2\x87\xb0\x11\x2a\x3c\xd8\x7f\xfb\xe4\x7c\xfe\xaf\x13\xb2\x78\xd6\x4d\x23\xe8\x72\xf1\x9d\x89\x13\x3e\xfd\x25\xd1\x6a\xf9\xdb\x59\x6a\x51\xd5\xd2\xb5\x7a\x8b\xde\x4e\x46\x3c\x93\x78\xab\x8a\x4d\xff\xbf\x45\x47\x63\x8b\xf3\x64\x44\x07\xde\xf3\x39\xeb\xa7\x5e\x2d\x0f\x36\x14\x4a\xd0\xbb\xa5\x4a\xa1\xaa\x8f\xea\x83\xe6\xc5\x66\xfb\xe4\xfd\xb5\x09\x26\xad\x60\x1b\xae\xea\x9f\x5d\x62\x73\xbe\x6c\x9f\xbc\x9b\xce\x22\x94\x3a\x48\x78\x62\x16\xb9\xc9\xeb\x60\x1b\xda\x2e\x6b\x04\x6d\xca\x4e\xe5\xee\x69\x6c\x70\x7b\x9e\x6e\x17\xc3\xd6\x60\x8f\x8a\xfe\xc7\x92\x77\x92\x9c\x9a\xcb\xf7\xdc\xc0\x71\x80\x04\x66\x9e\xe3\x10\x6c\xc3\x9e\xcb\x65\x5b\x94\x9b\x3f\x8e\xc3\x95\xc3\xad\x1f\x8e\x95\x3a\xef\xec\x72\xd0\x2d\x83\xd4\x19\x05\x25\x97\x16\xe7\xec\x65\xd5\xe8\x8c\x84\x04\x82\xc8\xea\x28\x80\x3b\x40\x95\xeb\x02\xbf\x7c\x7a\xff\xac\x9b\x56\x2b\x54\xb4\x39\x0e\x5b\xb8\x83\x20\xac\xa9\x91\xc1\x0c\x1f\x84\x2a\xf4\x10\x4a\x9d\x4f\xef\x1c\x12\x47\x33\x6f\xbd\xd4\x7a\x6d\x4e\x94\xe9\x62\x74\x03\xec\xc8\xd2\xef\x01\x00\x00\xff\xff\xbf\x0f\x83\x1e\x1a\x05\x00\x00"

func baseFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_baseFooterTmpl,
		"base/footer.tmpl",
	)
}

func baseFooterTmpl() (*asset, error) {
	bytes, err := baseFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base/footer.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseHeaderTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4b\x8f\x1b\xc7\x11\xbe\xf3\x57\x94\x27\x17\x09\x50\x35\xfb\x35\x2f\x2d\xb9\x81\xb4\x52\x42\x27\x5a\x39\xb1\xf5\xb0\x61\x18\x41\x73\xa6\xc9\x19\xed\x3c\xe8\x99\xe6\xec\x2e\x57\x7b\x32\x82\xdc\x72\xf4\x39\x87\x1c\x72\x0c\x90\x20\x08\x92\x43\xfe\x8c\x6d\x20\xff\x22\xe8\x19\x92\x3b\xa4\x28\x69\xe3\x07\x04\x2d\xbb\x6a\xaa\xbf\xaa\xfa\xaa\xab\xa7\x7b\x46\x1f\x3c\xfa\xe8\xe4\xd9\x67\xbf\x79\x0c\x89\xc9\xb3\xe3\xc1\x68\xfd\xf3\x01\x22\x7c\xfb\xe7\xbf\xff\xf7\xab\xbf\x00\xa2\x55\x6b\x15\x1f\x0f\x46\xb9\x36\x0a\x12\x63\x16\xa8\xbf\x5c\xa6\xcd\xd8\x39\x29\x0b\xa3\x0b\x83\xcf\x2e\x17\xda\x81\xa8\x93\xc6\x8e\xd1\x17\x66\x68\xa1\x8e\x20\x4a\x54\x55\x6b\x33\x7e\xfe\xec\x17\x18\x38\x30\x3c\x84\xf2\x29\x3e\x7f\x80\x27\x65\xbe\x50\x26\x9d\x66\x7d\xa0\x0f\x1f\x8f\x75\x3c\xd7\xce\xbb\x7c\x3f\x51\xc5\x7c\xa9\xe6\xfd\x69\x51\xb1\x9d\x51\xa8\x5c\x8f\x9d\x26\xd5\xe7\x8b\xb2\x32\x3d\x9b\xf3\x34\x36\xc9\x38\xd6\x4d\x1a\x69\x6c\x85\x7b\x90\x16\xa9\x49\x55\x86\x75\xa4\x32\x3d\x66\x84\xde\x83\x5c\x5d\xa4\xf9\x32\xef\xab\x96\xb5\xae\x5a\x59\x4d\x33\x3d\x2e\xca\x2e\xab\x2c\x2d\xce\xa0\xd2\xd9\xd8\xa9\x93\xb2\x32\xd1\xd2\x40\x1a\x95\x85\x03\x49\xa5\x67\x63\x67\x38\x53\x8d\x95\xc9\xa2\x98\x3b\xc7\x83\x8e\xe2\xef\xbe\xfe\x2b\x3c\x2c\x4b\x53\x9b\x4a\x2d\xe0\xbb\x3f\xfd\xf3\xdb\xff\x7c\x05\x27\x9f\x7c\x02\xdf\x7d\xfd\x87\x6f\xfe\xfd\x8f\x8e\xfb\x1e\xb0\xb9\xcc\x74\x9d\x68\x6d\xb6\xa8\xb5\x51\x26\x8d\x86\x51\x5d\x0f\xa7\x1b\x20\x92\xa7\x05\x89\xea\xfa\xe7\xcd\xf8\xea\xea\xe1\x32\xcd\xe2\x93\x32\xcf\x53\x73\x7d\xed\x7c\x1f\x38\xac\xb4\x1d\xff\x70\x54\x63\xf9\xea\xe1\xb6\xf2\x0f\x87\x4d\xd2\x79\x92\xa5\xf3\xc4\x0c\x3b\x9b\x61\xac\x67\x6a\x99\xfd\x08\x01\xd7\x3a\x9b\xb5\x7f\xde\x0a\xb3\xae\xe3\xa8\x8e\xaa\x74\x61\xa0\xae\xa2\x9b\xd9\xaf\xea\xe1\xab\x2f\x97\xba\xba\x44\x46\x18\x23\xb4\x8d\xe7\xd5\x41\x9c\xd1\xb0\x03\x38\x6e\x2b\x3e\x18\x6c\xf0\x62\x65\x14\xaa\x18\xa3\x2c\xed\xd6\xb5\xc2\xc5\x72\x8a\xbe\x2f\x3c\x2f\x74\x7d\x8f\xbb\x2c\xa4\x42\x38\xa0\xea\xcb\x22\xea\xdc\xdb\x16\xa9\xef\x0f\x87\x0b\x35\xd7\x2a\xe6\x64\x5e\x96\xf3\x4c\xd7\x97\x45\x9c\x46\xca\xa4\x65\x41\xa2\x32\x5f\x3f\xb5\x31\xaa\xb8\x9e\x5e\x76\x46\xe4\x55\xdd\x8b\x65\xa7\x81\x3a\x03\xac\x53\xa3\xb1\xd1\x55\x3a\x5b\x83\xf5\x1a\x2a\x7e\x18\xfe\xaa\x7c\x1a\x3d\x7d\xf4\xdb\x8f\x5e\x7e\xf6\xbc\x7a\xb9\x7c\x14\x3c\x8e\x99\x78\x12\xfd\xf2\xd1\x99\x78\xf9\x61\x43\x17\xd4\x8d\x4e\x9f\x7f\x5a\x89\xa8\xb7\x0f\x74\xf0\x53\x95\xc6\xcb\x77\xa2\x47\x65\xac\xf1\xe5\xb9\x77\x96\xf1\x2f\x27\x8b\xae\xe7\x6e\x98\x3a\x1e\x34\xaa\x82\xdf\x25\xb9\x81\x71\xf7\xf3\xfa\x35\x7c\xfe\xc5\xd1\xe0\xce\x6c\x59\x44\x16\xea\xce\x5d\xb8\x1a\x00\x58\xb3\x24\x87\x31\xc4\x65\xb4\xcc\x75\x61\x48\x54\x69\x65\xf4\xe3\x4c\x5b\xe9\x8e\xd3\xe1\x39\x77\x8f\x06\x00\x49\x4e\xea\x2a\x82\x31\x6c\x49\x4d\x72\xd2\x86\xda\x72\x98\xe4\xb6\x9c\x91\xe7\xf9\xd4\xd3\x74\x16\xb1\xa9\xc7\x43\xa1\x3c\x3e\x0b\xf4\x54\xb9\x41\x20\x42\xaa\x23\xe7\x68\xed\xb6\xee\x7b\x9d\x6b\xb3\x76\x59\x3f\xbc\x7c\xa6\xe6\x4f\x55\xae\x6f\x9c\x7f\x4e\xbf\x38\x82\x01\x40\x4d\x16\xaa\xd2\x85\x79\x5a\xc6\x9a\xa4\x45\xad\x2b\xf3\x50\xcf\xca\x4a\xdf\x49\xf2\x7b\x50\xdf\x3d\x1a\x5c\xdf\xbd\x73\xf7\x68\x70\x53\xb3\xc1\xe0\xea\x2a\x9d\x01\x79\x50\x99\x34\xca\xf4\xf5\xf5\x60\x64\x52\x93\xe9\xe3\xab\xab\x8d\x8e\x3c\xb3\x8a\xd7\x13\xad\xe2\x76\x74\x7d\x0d\x08\x1f\x6b\x15\x8f\x86\x9d\xe9\x4e\x65\xce\xf4\xe5\x79\x59\xc5\x75\xaf\x14\xfb\x50\xf0\x1a\x7a\x60\xf7\x2c\xd4\xde\xf6\x1b\xeb\x2e\xbe\xdd\x92\xf6\x70\x26\x26\xcf\x0e\xc3\x5c\x5d\xe9\xac\xee\xe5\xf1\xff\x04\x7a\xeb\x48\x6e\x7c\x15\xf1\xf5\xb5\x5d\x56\x6f\xbc\x6d\x2e\x30\x2e\x6a\x5c\x54\x7a\xa6\x4d\x94\xa0\x9d\x5a\x95\x59\x0f\xa3\x2c\x76\xb7\x95\xbe\xf9\x76\x63\x19\xa6\xb9\x9a\x6b\x12\xa9\x28\xd1\x51\xa2\xec\x32\xba\xcd\xac\x5c\x2c\x83\xfd\x49\x83\xf7\x4f\x7b\xe7\x16\x70\x1b\xbf\x66\x11\x7d\xff\xc9\xdd\x44\x15\xd7\x64\x4e\xe2\x72\x39\xcd\x74\x94\xa5\xd1\x19\x29\xb4\xb9\xcd\xf4\xf3\xf3\xf3\xb5\xef\xdb\x3a\x9c\x2d\x8b\x38\x2d\xe6\x51\x52\xa6\x91\xae\x73\x5d\xd7\x6a\xae\xeb\x5d\x90\xc1\x60\x34\xec\x0e\x33\x83\xd1\xb4\x8c\x2f\xed\x6f\xdd\xcc\xe1\x22\xcf\x8a\xba\xdb\x3c\xef\x77\xbe\xcf\x05\x29\xab\xf9\x90\x53\x4a\x87\x75\x33\x77\xa0\x7d\x4f\x8c\x9d\x38\xad\x17\x99\xba\xbc\x0f\x45\x59\xe8\x23\xe7\x78\x00\x30\xaa\x2f\xf3\x69\x99\x41\x1a\x8f\x9d\xed\xeb\xcd\x01\x7b\xe8\x78\x58\x5e\x8c\x1d\x0a\x14\x18\x0b\x20\x94\xad\x39\xc0\x7a\x2d\x6f\x5f\xfd\xdb\x05\xdd\x3e\x5c\x28\x93\xc0\x2c\xcd\x32\xac\x96\xd6\xa3\x6e\x74\x51\xc6\xb1\x03\x51\x96\x2e\xf6\x75\xf1\xd8\x39\xe5\x92\xb8\x34\x04\x1a\xa1\x47\x7c\x21\x80\x22\x63\xc4\x67\x2e\xb8\x24\x08\x85\x15\x64\xc8\x81\x71\xc2\x03\x49\x38\x93\xe0\x11\x26\x91\x50\x4f\x02\x93\x84\x86\x1c\x39\xa1\x9e\x07\x9c\x12\xd7\xf7\x4f\x02\x12\x4a\x01\x22\x24\xc2\xb3\x10\xae\xf4\x41\x0a\x22\x03\x17\x28\x48\x49\x28\x93\x8d\x4b\x42\x9f\x47\xed\x33\xe2\xf2\x10\xba\x29\x92\x78\x32\x04\x46\x49\xe8\x32\x60\x8c\x30\x57\x00\x27\x94\x72\xf0\xda\xe9\x36\x00\xeb\x51\x0a\x1f\x76\x3c\x32\x4e\xfc\x50\x42\x10\x10\x46\x3d\x60\x3e\xf1\x7d\x0f\x42\x09\x36\x31\x06\xa1\x9c\x84\x82\xb8\xd1\x3a\x39\x68\x93\x93\xd8\x26\x07\x6d\x72\x0c\xbb\xe4\xd0\x66\x87\x36\x3b\x9b\x1c\x76\xc9\x75\xae\xb0\x73\xd5\xc6\x13\xa2\x47\x5c\x2a\xc1\x25\x22\xf4\x90\x51\xe2\x71\xd9\x86\x2d\x5b\xb2\x98\x2b\x1a\xec\x32\xc4\x36\x45\xb4\x39\x62\x40\x42\x21\xb1\xcd\x11\x77\x8c\xb1\xcd\x11\x6d\x8e\x12\x6d\x1c\xd8\xe5\x88\x7d\xc7\x27\x8c\xba\x2d\xf7\x36\x6c\xeb\x8d\x12\xde\x66\x63\x53\x03\x3a\x69\x6b\x18\xac\x4e\x03\x0a\xae\x4f\x02\x4f\x9c\x04\x14\x3c\x8f\x78\x9e\x00\x5f\x10\x29\x3c\xf0\x39\x78\x9c\xb8\x52\x80\xcf\x27\x52\x2a\x0e\x1c\x28\x50\x86\x1c\xf9\x0b\x7e\x23\x73\xe4\x09\x0b\x6c\x04\x51\x48\x68\xd0\x52\xe6\x12\x2a\x25\x48\x62\xd7\x41\x37\x66\x9c\x48\x5f\x02\xb5\x2c\x50\x8e\xb6\xae\x96\x03\x2a\x43\x0c\x09\x63\x6d\x1d\x83\xa0\x21\xdc\xf7\x4f\x7c\x97\x08\xe6\x83\xf4\x88\xb0\x65\xa2\xe0\x32\xc2\x19\x6c\x63\x5d\x9d\x7a\x94\xb8\x9c\x01\x0f\x88\x90\x13\x19\x92\x50\x06\x0d\x93\x96\xb1\x24\x20\x21\xb5\xc5\x0b\x02\xeb\xcd\xb2\x1d\x20\x27\xbe\xcf\xd7\x63\x9f\xf8\xdc\x07\xda\x32\x2b\x50\x10\xee\x49\xf4\x09\xa7\x3e\x86\xc4\x26\xd3\x8e\x57\xa7\x1d\x28\xc8\x90\xf0\x86\x79\x44\xba\xc1\xc4\xa3\x24\x64\x91\x4f\x98\xe7\x77\xc8\xa1\x67\x0b\x10\xf8\xde\x46\x08\x08\x0f\x18\x50\x74\x89\xa4\x1e\x0a\x12\x52\x81\x01\x61\x7e\xd0\x36\x05\x77\x3b\x61\x1d\xf0\xca\x1e\x7e\x6c\xf3\xb5\x3d\x3d\xec\x9a\x7a\xbf\xbf\x93\x32\xd7\xfb\xad\xed\x01\xf3\x9c\x7e\xf3\xda\xa6\x0c\x88\x70\x25\x30\xc2\xa4\xa7\x88\x6b\x6b\x6c\xff\x21\xf1\x69\x00\x34\x43\x0f\xbc\x07\x37\x6a\x60\xc4\x05\x9f\xb8\x8d\xdf\xb3\x05\x3b\x4c\x24\x71\x77\x75\x48\xdc\x06\x65\xc2\x1b\xb9\x6f\x3b\x61\xf2\x80\xa9\xbf\xe3\x9e\x49\x0f\x6d\x64\x4f\x98\x00\x97\xf8\xa1\x78\xc1\x77\x1c\x60\x3b\x2b\x41\xb6\xab\x23\x6e\xc3\x08\x0f\xc5\x93\x5e\x5a\xab\x53\x4e\x5c\x60\xf2\x85\x4f\x7c\xea\x67\x2e\x71\x6d\xbb\xc0\xfa\xff\x0b\x26\x27\x8c\x36\x28\x0f\x80\x8b\x37\xc0\xe5\x84\x13\x77\xe5\x0c\xdf\x45\x7d\x7b\x5f\xb8\x15\xf7\x14\xf8\xa6\x1b\x80\x41\xdb\x0f\x3b\x0a\xe0\x4d\x5f\x81\x1c\xf8\x64\x47\xb6\x0d\xb5\xca\x99\x0b\x3c\x41\xd9\x88\x44\xbe\x90\xab\x9c\x82\xdc\x48\x41\x4f\x12\x8a\x01\xdb\x94\x11\x59\x83\x7c\x95\xa3\x0b\xa2\x41\x31\xf1\xac\x75\x2b\x52\x2b\xb2\x86\xf7\x8d\x81\x25\x62\x95\xa3\x44\x69\x21\x27\xac\x11\xab\x9c\xb6\x82\xec\x04\x17\x45\xe7\x7c\xe2\xad\x72\x09\xb2\xc3\x7b\x11\xbc\x87\xa8\x85\x2e\x17\x99\xc6\x28\xad\xa2\x5b\x12\xc6\x18\x78\x4a\x80\xe8\xa2\x42\x0f\x28\x74\x92\x95\x3d\xa0\x6b\x87\xef\x7a\x39\x75\xbc\x07\x2a\x80\xa0\x43\xb1\xae\xe8\x83\x4e\xa4\x6d\xd2\xc1\x2a\x0f\xd0\x57\x3e\xf8\xeb\xd2\xbb\x44\x7a\x81\xdd\xc2\x85\x7f\x22\x08\x97\xdc\x8e\x39\xf7\x40\x92\x80\xba\xc0\x28\x04\xc0\x68\x2d\x89\xef\xfa\x60\x9f\xd8\x95\x65\xa7\x70\x22\xfc\x07\x5b\x1c\x6b\xf5\x1e\x4a\xe6\x55\x1a\xdf\x8e\x09\xe0\xc4\x7d\x60\xdb\x91\xad\x57\x68\xab\xb1\xa5\xda\xd3\xda\x77\x89\xdb\xec\x6b\xed\xba\xf6\x13\xdc\x57\xb7\x0f\x1a\x14\x5d\xcf\xf0\x37\xd6\xbf\x78\xa3\xef\xc5\x81\x56\x16\x87\x7a\x69\x95\x7b\x56\xbb\xef\x90\x1e\x0c\xda\x2e\xe9\x03\x51\x33\x71\x38\xec\x70\x13\x76\xce\x5a\x77\x3f\x6a\xdc\xa7\x5d\x94\x07\xc8\x0e\x0f\x90\x6d\x4d\x0f\xb2\xcd\xdc\x43\x74\xdb\x8c\x7e\xaa\xc0\xdf\x4e\x78\x68\x77\x83\x5d\x75\x27\x35\x7b\x7a\x5c\x4b\x87\x28\xff\x29\x43\xdf\xef\x92\x91\x3d\xdb\xda\x43\xb0\x3d\x14\xeb\x0a\xa2\x4c\xd5\xf5\xd8\x59\xa0\x80\xe9\x1c\x63\x55\x9d\x81\xd1\x17\x06\xcf\x93\xd4\x68\xc8\xa7\x28\x36\x1d\x13\xa7\xcd\xc6\xda\xde\x81\x54\x5a\xe8\x6a\xfd\x6c\xf7\x69\x8c\xb3\x4c\x5f\x80\xfd\x83\xe7\x95\x5a\x80\xca\xd2\x79\x81\xa9\xd1\x79\x8d\x91\x2e\x8c\xae\xe0\xd5\xb2\x36\xe9\xec\x12\xd7\xb7\xa9\xb7\xa9\xb3\x39\xd6\x46\x55\x66\xeb\x07\x60\xa4\x36\xe7\x7f\x67\xcf\xe1\x01\x37\xf9\x14\xb9\xfd\x93\xcd\x91\xf6\xf3\x6a\x87\xb1\x8e\xca\xaa\xbd\xe1\xa0\x3d\xda\xf7\x7c\xd8\xad\xa4\x99\x6f\xe0\xa7\x29\xe4\x1a\xb9\x03\xdd\x17\x44\x47\x52\x07\x12\x9d\xce\x13\x33\x76\x04\x77\xa0\x2a\xed\xbe\x98\xe6\x73\x07\x54\x95\x2a\xcc\xd4\xd4\xde\x59\xb6\x87\x7d\xe7\x78\xb4\xac\x35\x5c\xd8\xeb\xcc\xfd\x2e\xf6\x9f\xdd\xdc\x1a\x86\xc7\xeb\x92\x6c\x5d\x0f\xd5\xf1\xe0\x46\x5a\x66\x9b\x38\x0a\xd5\x40\x54\x66\xc8\x78\xfb\x93\xcd\x51\x2d\x4d\x69\x63\xdb\x0e\x6d\xba\x6f\xa1\x36\x9f\x62\x1e\x23\xdd\xcd\x32\x4b\x8f\x0f\xf0\x59\xa8\x06\xdb\xbb\xd7\xe2\x02\x79\x8f\x36\xe7\x78\x52\xe6\xda\xc6\x37\x1a\x66\xe9\xdb\x81\x2a\x55\xc4\xb7\x00\xfb\x58\x15\x71\x99\x1f\x84\x6b\xbf\xb1\xed\x60\xaa\x69\xb9\x34\xb7\x00\xfd\xf6\xf7\x7f\xfb\xe6\x5f\x7f\xdc\x82\xb6\xdf\xd7\x6e\x98\x5d\x66\x7d\x6a\x67\x65\x95\x6f\x10\x2f\xb0\xd6\xaa\x8a\x92\x83\x0c\x4f\x51\x6c\x57\x51\x47\xb7\xd8\xe5\x31\x2d\x16\x4b\x03\xe6\x72\xa1\xc7\x4e\x87\xe3\xdc\x00\x9f\x9d\x83\xf5\xb4\xf9\x84\xb0\x23\xb4\x1d\xe7\xc0\x22\x53\x91\x4e\xca\x2c\xd6\xd5\xd8\xf9\xa4\x05\x20\x84\x38\xd0\xa8\x6c\xa9\xdb\x0f\x27\xbf\xee\xbe\x76\x5c\x5f\xef\x2e\xb2\xce\xb6\xdf\x1e\x43\x0b\xbf\x6d\xcb\x61\x9c\x36\xeb\xfe\xdd\x0c\xbb\xfb\xb0\xae\x8e\xff\x17\x00\x00\xff\xff\xff\x37\x77\x88\x0c\x18\x00\x00"

func baseHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_baseHeaderTmpl,
		"base/header.tmpl",
	)
}

func baseHeaderTmpl() (*asset, error) {
	bytes, err := baseHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base/header.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xc1\x8a\xdb\x3c\x10\xc7\xcf\xf6\x53\x0c\x3a\x2e\x28\xce\xb2\xdf\x07\xa5\x75\x0c\x85\x1e\xb6\x87\x96\x40\xfb\x02\x63\x7b\x62\xab\x2b\x4b\x41\x1a\x27\x59\xb4\x7e\xf7\x22\x3b\x4e\x9c\xb4\x7b\x6a\x0e\xb1\xa4\x19\x7e\x1a\xfd\x67\xf8\x87\xc0\xd4\xed\x35\x32\x81\x28\xd1\x53\xd6\x12\xd6\xe4\x04\xac\x86\x21\x4d\xd3\xbc\x56\x07\xa8\x34\x7a\xbf\x11\x95\x35\x8c\xca\x90\x13\x45\x1a\x42\xf6\x00\x9e\xd1\x31\x3c\x64\x31\x73\x99\xe8\xec\x11\x1a\xf9\xbf\x28\x52\x80\x5b\x80\x96\x5d\x2d\x3f\x88\x02\x00\x52\x88\xbf\x64\x19\xaf\xe5\x4e\xd3\x09\xe2\x9f\xac\xac\xee\x3b\x03\xa8\x55\x63\xa4\x62\xea\xbc\xf4\xec\x88\xab\x76\x8a\xfb\xd6\x29\xf3\x22\xd7\x50\x36\xf2\xd8\x2a\x26\x01\x9e\x5f\x35\x6d\xc4\x51\xd5\xdc\x7e\x04\xec\xd9\x7e\x12\x45\x9a\x26\xf1\x9e\xe5\x35\x5a\x79\x96\x8d\xb3\xfd\x1e\xae\x4b\xb9\xd3\xbd\x6f\xa1\xb4\xae\x26\x27\x4b\xcb\x6c\x3b\xf0\x95\xb3\x5a\xa3\x23\x14\xc5\xc4\x49\x42\x70\x68\x1a\x82\xd5\x67\xc7\xaa\xd2\xe4\x87\x61\x8a\x00\xe4\x08\xad\xa3\xdd\x46\x64\x2d\x77\x3a\x0b\x61\xf5\xad\xf9\xfa\x65\x18\x56\x71\x2b\xfe\xbc\x7d\x7c\x16\xdc\xed\x25\x56\xac\xac\x81\xfd\xab\x7c\x02\xdd\x4a\x56\x4d\xcb\x02\xd0\x29\x94\x55\xef\x1c\x19\xde\x08\x76\x3d\xcd\x15\xdd\xbf\xee\x2c\xe2\x51\x3e\xae\xd7\x37\xf2\x55\x64\x98\x1c\xfc\xea\x3d\xab\xdd\xab\x8c\xdd\x24\xc3\xb2\x24\x3e\x12\x99\x25\x0e\x20\xf7\xec\xac\x69\x66\x66\x57\xca\x47\x51\x84\xb0\xfa\xa9\x58\x13\xbc\xc1\x0f\xdc\xd1\x30\xe4\xd9\x94\xb6\xac\x24\xab\xd5\xe1\x9d\xca\x62\xfb\x1f\xd7\x10\x61\xe0\x3b\xd4\x7a\x44\x3e\x73\xa7\xe1\x0d\x9e\x09\xeb\x91\x1e\xb1\x4b\x46\x9e\xe1\x65\x1d\x02\x99\x7a\x16\xfc\x9a\x96\x26\x49\x32\xef\x92\x9b\x89\xea\x4a\xf9\x24\x8a\x73\x6c\x0c\x1a\x3c\x4c\x62\x6a\x2c\x49\x6f\xc4\x16\x1b\x02\x83\x07\xd5\xe0\xa8\x3b\x9d\xb0\xdb\xeb\x51\xdd\x24\x89\xf3\xd9\xeb\x19\xb6\xc7\x46\x99\x29\xeb\x5e\xc4\x49\x5b\x31\xdd\x01\x90\x24\x21\xa8\x1d\xac\xb6\x8e\xb6\x76\x9c\x90\x64\x9a\x9e\x5c\xab\x05\x8d\xc6\xce\x88\x22\xc7\x9b\x43\xad\xcc\x8b\x98\x47\x69\xef\xe8\x10\x47\x69\x46\x89\x62\xeb\xe8\x10\x45\xc9\x33\xad\x8a\x0b\x79\x96\x66\x71\x10\x2b\xf8\x4e\x27\xfe\xd7\x12\x0c\x9d\x38\x96\x70\x61\x89\x22\x2e\xdf\xaf\xe1\xac\x5c\xd6\xeb\xa9\x21\x99\xc1\xc3\xe8\x05\xe7\x3e\xfc\xd5\x15\xfe\x1b\xed\x22\xf6\xf8\x6a\x48\x5e\xd5\x54\xa2\xcb\xce\xdf\xc9\x94\x2e\x9c\x19\x37\x7a\x11\x99\x7a\x72\xa2\xeb\xe9\x9d\xb1\xed\xac\xe5\xb3\xb1\xfd\x0e\x00\x00\xff\xff\x24\xa1\x3d\x40\xf5\x04\x00\x00"

func homeTmplBytes() ([]byte, error) {
	return bindataRead(
		_homeTmpl,
		"home.tmpl",
	)
}

func homeTmpl() (*asset, error) {
	bytes, err := homeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pageContentTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xcd\x6e\xc3\x20\x10\x84\xef\x3c\xc5\xca\xc7\x48\x84\x43\x52\xa9\x07\xd7\x52\x6f\x3d\x56\x6a\x5f\x60\x31\x9b\x18\x09\x7b\x2d\x58\xa5\x07\xe2\x77\xaf\x62\xe3\xc4\xfd\xb9\xc0\xc2\x8c\x3e\xc4\x4c\xce\x42\xfd\x18\x50\x08\x2a\x8b\x89\x4c\x47\xe8\x28\x56\xb0\x9f\x26\xa5\x94\xaa\x9d\xbf\x40\x1b\x30\xa5\x97\xaa\xe5\x41\xd0\x0f\x14\xab\x46\xe5\x6c\x76\x90\x04\xa3\xc0\xce\x4c\x93\x02\xd8\x3a\x23\x7f\xc1\x59\x3f\x55\x8d\x02\xf8\xa9\xb4\x1c\x74\xef\xf4\x73\x91\x00\xea\xee\xb0\x6a\xa3\xd5\x47\xe8\x6f\x8b\xe5\xe8\x28\x6a\xcb\x22\xdc\xdf\xad\x00\x39\xef\x5f\xa3\xf8\x36\xd0\xfe\xd3\x4b\x20\xb8\xc2\x07\x9e\x68\x7e\x7f\x86\x99\xee\xd0\xa8\xf5\x80\x8b\x75\xc5\xdb\xc0\x67\x3d\x72\x92\xff\x81\x6f\xd2\x07\xb8\xc2\x3b\xc6\x44\xb7\x79\x03\x2d\xa0\x42\xae\x8d\xf3\x97\x75\xfe\xfb\xb5\xe3\x1d\xbf\xcd\x36\x79\x47\x16\xa3\x29\xfb\x92\xef\x83\x35\xc7\x49\x83\x5b\xc2\x7c\xdc\xfe\x2a\xe7\xc4\x2c\xa5\x9c\xef\x00\x00\x00\xff\xff\x21\x1f\x67\x65\xb9\x01\x00\x00"

func pageContentTmplBytes() ([]byte, error) {
	return bindataRead(
		_pageContentTmpl,
		"page/content.tmpl",
	)
}

func pageContentTmpl() (*asset, error) {
	bytes, err := pageContentTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "page/content.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sidebarSidebarTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x4d\x53\xe3\x38\x10\xbd\xe7\x57\xf4\xea\x04\x54\x64\xb3\x24\x10\x02\x71\xf6\xc2\x79\x4f\x7b\xd8\xaa\xa9\x29\xaa\x6d\xb5\x6d\x81\x2c\xb9\xd4\x72\x5c\x9e\x90\xff\x3e\xe5\x2f\xc8\x30\xd4\xe4\x10\xd9\x7a\xaf\x5f\xbf\xfe\xf0\xe2\x78\x8c\xaf\x20\x33\xc8\x9c\x88\xda\xb1\x0e\xda\x59\xc9\x41\x67\xaf\x9d\x80\xab\xf8\x74\x5a\xec\x94\x3e\x00\x87\xce\x50\x22\x82\xab\x1f\xe0\xc6\x53\xf5\x28\xf6\x8b\x05\x00\xc0\x80\xce\xf1\x72\x0d\x55\x2a\x57\x90\x16\xd2\xe8\xa2\x0c\xe0\x5d\x63\x15\x29\xb1\x1f\xb8\x00\xbb\x7a\xe6\x56\xa9\xbc\x16\xfb\x27\x74\x90\xa1\x85\x94\xe0\x09\xdd\x12\x0e\xe4\xbb\xfe\xe9\x11\x2c\x56\x34\x43\x39\x56\xae\xe1\x09\x1d\x5f\xa2\xc5\xbf\x58\x91\x21\xe6\x25\x84\x92\x20\xa5\x42\x5b\xab\x6d\x01\x2e\x87\x92\xf0\x40\x16\xd0\x2a\x20\xf4\xa1\x5c\xbe\x2b\xf4\xd4\xca\x85\x92\x7c\xcf\x43\x63\x20\x94\xda\x16\x1c\x2d\xfe\x2b\xc9\x53\xee\x3c\x0d\x24\x4f\xa0\x19\xd0\xb4\xd8\x31\x58\x07\x8a\x58\x7b\x82\xe0\xc0\xa5\x4c\xfe\x40\xa0\x03\x43\xeb\xac\x22\x9f\x37\xc6\x0e\x3e\xfa\x7c\x9f\x63\xbf\x0e\xe4\x16\xbb\x21\x25\x13\x84\xd6\x41\x89\x07\x1a\xbc\x71\x5f\xb5\xf3\xba\xd0\x16\xd2\x26\x80\xd2\x79\x4e\x9e\x6c\x18\xfa\xf1\x91\xa3\x03\xf4\x7d\x7b\x8c\x21\x05\xff\x37\x68\x97\xc3\xff\x8f\x52\x0f\x8c\xf1\xa6\x17\x54\xce\xf9\x3e\x7b\x5f\xea\x68\x97\xa3\x5d\x5c\x8f\xf3\xd8\xc5\x4a\x1f\xe6\x39\x72\xe6\x75\x1d\x00\xb9\xb3\x19\xb0\xcf\x12\x51\x86\x50\xf3\x43\x1c\xd7\x58\x10\xaa\x9b\xa8\x70\xae\x30\xc4\x9d\x55\x3a\xc3\x7e\x4f\xa2\xcc\x55\x13\x1a\xbf\x70\x8c\x8a\xd3\x6e\x24\x45\x2f\xfc\x4f\x66\x34\xd9\x90\x64\x28\xeb\x26\x95\x9b\xcd\xea\xee\x6e\x7b\xbb\xb9\xbb\xb9\xfd\x7b\x7b\xbd\x5a\x89\x71\x23\x32\xef\x98\xc7\x82\x13\x81\xd6\xd9\xae\x1f\x94\xd8\xef\xe2\xd1\xcf\x64\xf4\x2f\x29\x81\xb5\xa2\x67\x43\x79\x00\x29\xa7\x6b\x6d\x79\xde\xa7\xb3\xe4\x93\x74\xff\x9b\xd6\x56\x69\xae\x0d\x76\x0f\xa9\x71\xd9\xeb\x19\xac\x30\xa0\x44\x25\x27\xab\xe2\xcf\x5e\xcf\x23\xd8\xb8\x90\x88\xfb\xf5\xf6\x7e\xb3\xbe\xdf\x6e\xd7\x5f\x50\x72\xe7\x2b\x0c\x89\xc0\x26\xb8\xcf\x70\xde\x18\x23\x5b\xad\x42\x29\x3d\x71\xed\x2c\xeb\x43\xff\x79\xf9\x86\xfa\xda\xb5\xe5\xfd\xf9\x54\xf6\x1f\xe1\x17\x67\x85\x42\x02\xad\xb6\xca\xb5\xd1\xf9\xe5\xdb\x1b\x7c\xfb\x7e\x19\xd5\x0d\x97\x17\xc7\xd3\xe5\xe3\x34\xea\x59\xe9\xd7\x86\xa6\xe8\x81\x03\xfa\x8f\xa6\x1e\x8f\x81\xaa\xda\x60\x20\x10\x13\x23\x9e\xce\x67\x4b\xad\x80\xe8\x74\xfa\x5d\x83\xac\x9a\x15\x16\xef\x8b\x35\x1e\x3f\x03\x00\x00\xff\xff\xa5\x54\x7a\xad\x63\x04\x00\x00"

func sidebarSidebarTmplBytes() ([]byte, error) {
	return bindataRead(
		_sidebarSidebarTmpl,
		"sidebar/sidebar.tmpl",
	)
}

func sidebarSidebarTmpl() (*asset, error) {
	bytes, err := sidebarSidebarTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sidebar/sidebar.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sidebarSidebar_newTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcb\xce\xd3\x30\x10\x85\xf7\x79\x8a\x51\x58\x4f\xd2\x0a\xb1\x4b\xbb\x62\x83\x04\x6c\xe0\x05\x1c\x7b\x12\x5b\x9d\xd8\x91\x3d\xbd\xa0\xe0\x77\x47\x71\x53\xd1\xf6\xdf\xcc\x45\x23\x9f\x73\xfc\x55\x55\xd5\x19\x77\x01\xcd\x2a\xa5\x43\x6d\x70\x60\xba\xc1\x5a\x50\x07\x3e\x4f\x1e\x14\xbb\xd1\xa3\x13\x9a\x12\x26\x89\x24\xda\xde\xef\xc9\x46\xe7\x4f\xb8\x83\x7e\xc4\xab\x75\x42\xf5\xb1\x02\xe8\x14\xd8\x48\xc3\xa1\xfe\x54\xbf\x89\x3e\x0b\x69\xf2\x42\xf1\x4d\x67\xc6\xcf\xc0\xeb\x68\x54\x3c\x81\xd0\x4d\xd0\x90\x0e\x51\x89\x0b\x1e\x7d\xf0\x04\x7d\x88\x86\x22\xf6\x41\x24\x4c\xc5\x0f\xa0\x4b\xb3\xf2\x0f\xaf\x21\xe1\x17\x18\xae\x98\x68\x72\x7d\x60\x53\x1f\xbf\x2b\xa1\x24\xa0\x03\x33\xe9\x55\xa9\x6b\xd7\x07\x25\x6b\xab\x4a\x7b\x02\xc0\x2e\x09\x8e\x31\x9c\x67\xf8\x3f\xe2\xc0\xe7\x64\x5f\xcd\x21\xe9\x18\x98\x55\x24\xb5\xe5\x28\x65\x59\xa2\xf2\x23\x41\xf3\x93\xae\x89\x92\xe4\x7c\xcf\xf8\xa0\xd2\x5a\x99\xb8\x5d\x96\xe6\xc7\xf8\xed\x6b\xce\xcd\xba\xd6\x1f\xcd\x0b\x26\x78\xdb\x51\x95\xfc\x30\xff\x59\x41\x59\x14\x37\x5a\xd9\xcc\x5f\x7f\xa1\x03\xe3\x7e\x07\x53\x8f\x7b\x48\x93\x62\xae\x8f\xcb\xd2\xfc\x76\xc2\x04\x7f\xe1\x97\x1a\x28\xe7\xae\x35\xee\xb2\x11\xbc\x73\x58\xd3\x93\x37\x39\x57\x85\x4d\x39\x6f\xed\x5f\x00\x00\x00\xff\xff\xa5\xaf\x15\x5c\x27\x02\x00\x00"

func sidebarSidebar_newTmplBytes() ([]byte, error) {
	return bindataRead(
		_sidebarSidebar_newTmpl,
		"sidebar/sidebar_new.tmpl",
	)
}

func sidebarSidebar_newTmpl() (*asset, error) {
	bytes, err := sidebarSidebar_newTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sidebar/sidebar_new.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sosoHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x8b\xdb\x30\x10\x3d\xdb\xbf\x62\xd0\x71\x41\x71\x96\x6d\xa1\xb4\x8e\xa1\xb7\x2d\x85\x12\xd8\xf6\x07\x8c\xed\x89\xad\xae\x2c\x05\x69\x9c\x0f\xbc\xfe\xef\x45\xfe\x48\x9c\xb4\x7b\x6b\x0e\xb1\x34\x33\x3c\x3d\x3d\x3d\x5e\xd7\x31\x35\x7b\x8d\x4c\x20\x72\xf4\x94\xd4\x84\x25\x39\x01\xab\xbe\x8f\xe3\x38\x2d\xd5\x01\x0a\x8d\xde\x6f\x44\x61\x0d\xa3\x32\xe4\x44\x16\x77\x5d\xf2\x00\x9e\xd1\x31\x3c\x24\x61\x72\x39\xe8\xec\x11\x2a\xf9\x51\x64\x31\xc0\x2d\x80\x96\x4d\x29\x3f\x89\x0c\x00\x62\x08\xbf\x68\xd9\x2f\xe5\x4e\xd3\x09\xc2\x9f\x2c\xac\x6e\x1b\x03\xa8\x55\x65\xa4\x62\x6a\xbc\xf4\xec\x88\x8b\x7a\xec\xfb\xda\x29\xf3\x2a\xd7\x90\x57\xf2\x58\x2b\x26\x01\x9e\xcf\x9a\x36\xe2\xa8\x4a\xae\x3f\x03\xb6\x6c\xbf\x88\x2c\x8e\xa3\x70\xce\xf2\x18\xad\x3c\xcb\xca\xd9\x76\x0f\xd7\xa5\xdc\xe9\xd6\xd7\x90\x5b\x57\x92\x93\xb9\x65\xb6\x0d\xf8\xc2\x59\xad\xd1\x11\x8a\x6c\xc4\x89\xba\xce\xa1\xa9\x08\x56\x5f\x1d\xab\x42\x93\xef\xfb\xb1\x03\x90\x22\xd4\x8e\x76\x1b\x91\x74\xdd\xea\xc5\xb6\xae\xa0\xbe\x0f\xeb\x5f\x9e\xdc\xb8\xfa\x56\xf6\xfd\xaa\xe6\x46\x8b\xbf\xc9\x0c\xb7\x84\xbb\xbd\xc4\x82\x95\x35\xb0\x3f\xcb\x27\xd0\xb5\x64\x55\xd5\x2c\x00\x9d\x42\x59\xb4\xce\x91\xe1\x8d\x60\xd7\xd2\x4c\xf0\xfe\xb2\x93\xa6\x47\xf9\xb8\x5e\xdf\xa8\x59\x90\x61\x72\xf0\xbb\xf5\xac\x76\x67\x19\x1e\x97\x0c\xcb\x9c\xf8\x48\x64\x96\x70\x00\xa9\x67\x67\x4d\x35\x63\x36\xb9\x7c\x14\x59\xd7\xad\x7e\x2a\xd6\x04\x6f\xf0\x82\x3b\xea\xfb\x34\x19\xc7\x96\x4c\x92\x52\x1d\xde\x61\x16\xdc\xf0\xb8\x86\x00\x06\xbe\x41\xad\x07\xc8\x67\x6e\x34\xbc\xc1\x33\x61\x39\xa0\x07\xd8\x25\x46\x9a\xe0\x65\xdd\x75\x64\xca\x59\xff\xeb\x58\x1c\x45\xd1\xbc\x8b\x6e\x0c\xd6\xe4\xf2\x49\x64\x53\x6f\x68\x1a\x3c\x8c\x62\x6a\xcc\x49\x6f\xc4\x16\x2b\x02\x83\x07\x55\xe1\xa0\x3b\x9d\xb0\xd9\xeb\x41\xdd\x28\x0a\x76\x6d\xf5\x0c\xb6\xc7\x4a\x99\x71\xea\x5e\xc4\x51\x5b\x31\x9e\x01\x10\x45\x5d\xa7\x76\xb0\xda\x3a\xda\xda\xc1\x30\xd1\x68\xa6\x54\xab\x05\x1a\x0d\x2f\x23\xb2\x14\x6f\x8a\x5a\x99\x57\x31\x3b\xcb\xdb\x60\xa3\xef\x74\x3e\x5a\x57\xf6\x7d\xb2\x77\x74\x08\x95\x19\x7a\x34\x57\xb6\x75\x74\x08\x4a\xa5\x89\x56\xd9\xe5\xb8\x59\xaf\x45\x21\xd0\xfa\x41\x27\xfe\xdf\xbc\x0c\x9d\x38\x54\x2e\xd8\x13\xb1\xb0\x7f\x9f\xd8\xa4\x71\xd2\xea\xf1\xe9\x12\x83\x87\x21\x44\xa6\x17\xfb\x67\x9c\x7c\x18\x72\x26\xb8\xe1\x9a\x64\x5e\x95\x94\xa3\x4b\xa6\xef\x98\x66\x17\x9c\x19\x6e\x08\x31\x32\xe5\x18\x61\xd7\xea\x5d\x22\xee\xac\xe5\x29\x11\xff\x04\x00\x00\xff\xff\x2d\x28\x77\xd5\x2e\x05\x00\x00"

func sosoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_sosoHtml,
		"soso.html",
	)
}

func sosoHtml() (*asset, error) {
	bytes, err := sosoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "soso.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"about.tmpl":               aboutTmpl,
	"base/footer.tmpl":         baseFooterTmpl,
	"base/header.tmpl":         baseHeaderTmpl,
	"home.tmpl":                homeTmpl,
	"page/content.tmpl":        pageContentTmpl,
	"sidebar/sidebar.tmpl":     sidebarSidebarTmpl,
	"sidebar/sidebar_new.tmpl": sidebarSidebar_newTmpl,
	"soso.html":                sosoHtml,
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
	"about.tmpl": &bintree{aboutTmpl, map[string]*bintree{}},
	"base": &bintree{nil, map[string]*bintree{
		"footer.tmpl": &bintree{baseFooterTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{baseHeaderTmpl, map[string]*bintree{}},
	}},
	"home.tmpl": &bintree{homeTmpl, map[string]*bintree{}},
	"page": &bintree{nil, map[string]*bintree{
		"content.tmpl": &bintree{pageContentTmpl, map[string]*bintree{}},
	}},
	"sidebar": &bintree{nil, map[string]*bintree{
		"sidebar.tmpl":     &bintree{sidebarSidebarTmpl, map[string]*bintree{}},
		"sidebar_new.tmpl": &bintree{sidebarSidebar_newTmpl, map[string]*bintree{}},
	}},
	"soso.html": &bintree{sosoHtml, map[string]*bintree{}},
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