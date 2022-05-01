// Code generated for package conf by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../conf/app.conf
package conf

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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x5b\x6b\xe4\x3a\x12\x7e\xd7\xaf\x28\xcc\x0e\x24\xd0\x71\xdb\xc9\x92\xcc\x76\xe3\x87\x65\x96\xb9\x40\x96\x1d\x26\x19\x06\x26\x04\x47\x6d\x95\xdd\xda\x96\x55\x46\x92\xfb\x92\x5f\x7f\x28\xd9\xee\x74\xce\xe4\xc0\xbc\x9c\x7e\x69\x4b\xaa\xfa\xaa\x54\x97\xaf\x24\xbb\xae\xb4\xb2\x45\x28\x20\xd9\xe2\xf3\x05\xda\x44\xac\x9c\xb4\xea\x97\x5d\xd7\xdb\xb2\xf7\xe8\x78\xcf\x11\x85\x61\xa7\x25\x15\xa5\x14\x6e\x13\x21\x1e\x14\xae\xfa\xe6\x51\x74\xe4\x42\x91\xe7\x59\x7e\x29\x84\x78\x30\xd4\x3c\x8a\x25\xdc\xaf\x11\x0c\x35\x50\x93\x6b\x65\x00\xd4\x61\x8d\x0e\x92\xff\x7b\xb2\x09\x90\x83\x24\xe0\x3e\x24\x62\x3c\x2e\xa6\x35\xdb\x2a\x3b\x19\xd6\xbc\x65\xa8\xf1\x6c\xc7\xa1\xd2\xfe\x11\xc4\x12\x7c\x20\x87\xe0\xd1\x7b\x4d\x56\xa0\x95\x2b\xc3\x0e\xd5\xd2\x78\x14\x52\x29\x87\xde\xb3\x66\x7e\x79\x93\x66\x69\x96\xe6\x8b\xeb\xab\x9b\x7f\x25\xa2\x93\xde\xef\xc8\x29\x3e\x4b\x84\x5a\xf1\x7f\xc6\xd0\x23\xd4\xe4\xf1\xb8\x84\xce\xd1\x56\x2b\x74\xb3\xa3\xe3\x2d\xb6\xe4\x0e\xc9\x0c\x92\x5a\x1b\x4c\x66\xf1\x0e\xd1\xb1\x24\x15\x93\x38\x14\x30\xc8\x8d\x70\x15\xd9\x5a\x37\xbd\x93\x81\x41\x6b\x72\xe0\xd0\x77\x58\x05\xbd\xc5\xa3\x8d\x85\x58\xc2\xc5\xa8\xb7\x00\x45\xe8\xc1\x52\x00\x8b\xa8\x40\xda\xc3\x08\x02\x07\x0c\x51\x90\xcd\x2f\x8e\x8e\xf2\x0a\x38\x5e\x33\xc0\xb4\x49\xe1\x49\xc9\x20\xe7\xe3\xa9\x7f\x8a\x1a\xd1\xcb\x05\x58\x0c\x3b\x72\x9b\x22\x54\xdd\x8c\x43\x55\xc4\xd8\xcc\xa6\xd0\x14\xad\xac\xa4\x23\x3b\x53\xab\x22\x9b\x75\x44\xa6\xf4\xfa\x19\x8b\x3c\xcb\x66\x5a\x19\x2c\x83\x6e\x91\xfa\x50\xe4\xef\xb3\xe3\x85\xcb\xd1\xb9\x02\x5e\xd9\x3d\x5e\x9f\x36\x1a\x21\xd6\x56\xa0\x31\x79\xe1\x24\xcc\x5a\xa1\x0d\xba\xd6\xe8\x52\x31\x08\x4f\x85\xa8\x4b\xa3\x37\x58\xea\x56\x6a\x23\x96\xf0\x63\x8d\x31\x0d\x8c\x82\x61\x02\xd6\x16\x3e\xdf\xdf\x7f\xbd\x03\xb2\xe6\x70\x44\xf0\x58\xf5\x8e\x31\x82\xeb\x71\xf4\xe4\xd3\x07\xd0\x36\xa0\xdb\x4a\xc3\x5a\x1e\x2b\xb2\xca\xc7\x8c\x4c\xbe\xf0\x05\x52\xd1\x54\xe5\x51\xb0\x80\xab\xeb\x2c\x1b\x11\x5a\xb9\xd7\x6d\xdf\x82\xd1\x35\x02\x87\xe2\xcf\x38\x72\x42\x4a\x45\x2b\xf7\x25\xcb\xc5\x90\x41\x01\xef\xaf\xff\x79\xc4\x39\x8d\x09\xab\x7d\xb8\xfb\xf6\x11\x02\x6d\xd0\xa6\xa2\xf2\xae\x2e\x5f\xc7\xa1\xe4\x3d\x6e\xab\x4a\x56\x6b\x9c\xca\x34\x2e\x40\x2a\xd9\x85\xb7\x6b\x74\x28\xcc\xa1\x48\x5b\x6c\xa3\x7c\x92\x8a\x51\xe3\xb4\x4c\x3f\xd2\x8b\x62\x8c\xe3\xec\x2f\x82\x95\x8a\x93\xc0\xe4\x97\xd9\xa4\x3b\x98\x02\x69\xd5\x89\xa9\x19\x57\xad\xe5\x42\x27\x0b\x6b\xf2\x01\xc6\xe6\x5c\xfc\x5d\x15\x39\xf6\x50\x34\xbf\x80\xa7\x17\x02\xc8\xf3\xcb\x3c\x7f\x12\xd1\x09\x76\xfc\xd5\xbe\x10\x0f\x3b\x5c\x4d\x61\xed\x1c\x05\xaa\xc8\x40\x58\xcb\x00\xda\x43\xef\x51\x0d\x35\xe7\xb6\x08\x4a\x3b\xac\x02\x04\x27\xeb\x5a\x57\xbc\xcf\xb5\x2c\xbb\xce\xe8\x2a\x76\x78\x2a\x96\xf0\xa1\x77\x0e\x6d\x30\x07\xf0\x7d\xc7\xcc\xe8\x21\x59\x87\xd0\x71\x56\xf8\xdf\x47\x0a\xa9\x1a\x3d\x86\xac\xb7\x7a\x3f\x30\xc8\x60\xba\x00\x96\x1a\x1d\x9a\x18\x2d\x10\xac\x10\x8c\xf6\x01\x2d\x2a\x58\x1d\x7e\xb5\xcc\x5a\x25\xcb\x43\x01\x59\xbc\x62\x26\xa6\x6b\x91\x0b\x60\xfb\x76\x35\x34\xd0\x6f\x21\x45\x1d\x26\xd2\x3c\xcb\xf3\x44\x30\xd4\x27\xb4\xe8\x64\x40\xf0\x01\xbb\x98\xc8\xff\x39\xa8\x1d\xb5\x20\x21\xed\xea\xfd\xc0\x45\xb8\x67\x55\x54\xc3\x09\x83\xff\xd0\x56\xd1\xce\x43\x85\x8e\x7b\xbd\x1a\x20\x98\x0a\xce\x14\x89\x65\x24\xbb\x9a\x5c\x83\x81\xbd\x1b\xf4\xa3\x62\xe7\xf4\x96\x85\x37\x78\x38\x67\x73\xff\x00\xea\xd0\x7a\x6f\xa0\xdb\x54\x3e\xbf\x84\x0b\x6d\x23\x6a\xb4\x7e\x41\x7d\x18\x57\xd8\xc2\x85\xa5\x0d\x1e\xfc\xef\x69\x6d\xf0\x30\x29\xf1\x81\xe7\x0f\x85\x5e\xf0\xa2\x8c\xb7\x2a\xa0\xea\x7d\xa0\x76\x1e\x53\x38\x9f\xcc\x88\x0d\x1e\xde\x14\x18\x11\xa7\x0c\xb4\xda\x46\xf2\x90\xc6\xd0\x0e\x15\xdc\xdf\xde\xc1\x16\x1d\xb3\xc5\x0c\xaa\x37\x2a\xe6\xfe\xf6\x2e\xcf\xb8\x52\xf8\x23\x9f\x3e\x2e\x93\xd9\x50\x34\xbc\xb8\x4a\x52\x11\x8c\x2f\x5b\x6d\xcb\x11\x0b\x0a\x88\x62\xdc\x9a\x71\x2e\xa0\x6b\xf5\xc0\x6e\xbb\x35\xda\xb1\x8c\xa7\xfa\xdd\x6a\x09\xdf\xad\xde\x83\xa2\x56\x72\x9f\x53\xb5\xc1\x90\x0a\xae\xc8\x72\x58\x94\x27\x08\x05\x5c\x5f\x5f\x8b\x25\xdc\x52\x25\x0d\x9c\xfd\xe7\xbf\x3f\xcf\xe1\xfb\xb7\xdb\x48\x62\xdc\xc5\xe8\x3c\x9c\xc5\x19\x74\x77\xf7\x19\xfa\x4e\xc9\x80\xe7\x20\xab\x8a\x59\xd1\x36\xb0\xc3\x55\x74\x40\x57\xc8\x7d\xf2\xc5\x42\xcb\x2d\x59\x49\x8f\x1e\x0e\xd4\x83\xa2\x97\xb9\x17\x08\xaa\xb5\xb4\xcd\x30\x2e\x14\xd6\xb2\x37\x01\xb6\xd2\xf4\x51\xf9\xdf\x86\x59\x4c\x87\x48\x58\xa0\x6b\xd6\x77\xd1\x6e\xbc\xa2\x03\xce\x1f\x37\x30\x03\xc6\x81\xc3\x6c\x2a\x7d\x1c\x17\xf1\x30\x15\x86\xef\x51\xc6\xf7\x46\xef\xb8\xef\xde\x9d\x4d\x4d\x78\xee\x17\xf3\xf9\xbb\xb3\x63\x4b\x9d\xfb\xc5\xb8\xe2\xf4\x9c\xfb\xb9\x78\x3d\x92\x94\xf6\xf1\x21\x62\xa8\x69\xf8\xaa\xda\x82\xa3\x3e\xf0\x54\x1b\x8f\xca\x61\x5d\xf2\x8b\xe8\x38\x98\x4e\x10\xc6\x97\xcc\x49\x1f\x82\xc1\x2d\x1a\xf8\xf4\xf3\xcb\x57\xa8\xa8\xed\xdc\x34\x5c\x06\xd1\xb2\x79\xd6\xdd\xf1\xe5\xf3\x26\x58\x20\x30\x24\x15\x48\xef\x31\x78\x38\xd3\x29\xa6\x90\xf0\xc0\xe6\x7a\x0a\xd8\x76\x46\x06\x8c\x7c\xd4\xf5\x2b\xa3\xab\xe4\x7c\x68\x5a\xa5\xfd\x06\xb4\xf5\x01\xa5\x02\xaa\x01\xdb\x15\x2a\xc5\x6c\xa1\xed\x30\x25\x19\xb7\x1c\x70\x4b\x56\x29\xa3\xca\x89\x37\xf1\xf5\x27\xad\xe2\x68\x74\xb2\xc1\x63\xa5\x48\x4b\xf6\xd0\x52\x1f\xb9\xd5\xf9\x59\xcc\x4e\xcc\x2b\xf8\x35\xf5\x46\xc5\x94\x69\x5b\x99\x5e\x71\x8c\x7c\xbf\x8a\xaf\xc1\x89\x91\xd7\xd2\x2a\xf3\x42\x5c\x0e\xb9\xf4\x23\x73\xef\x0f\xa9\x18\x4d\x8e\x09\x9d\x0b\xf1\xd0\x92\x6d\x48\xad\x1e\xe3\xe3\xb0\x78\x61\xff\xcb\x9b\x2c\xbf\x11\x6a\x55\x6c\xf1\x99\xc7\xab\x6e\x65\x83\xa3\x14\x67\x7a\x31\x9f\x8f\x2c\xba\xb8\xba\xba\xba\x9a\xeb\xb9\xe8\xb4\x6d\xde\x3c\xe3\x83\x78\x5a\xf2\xeb\x8e\xac\xc7\x82\x36\x42\x3c\x10\x27\xe4\x51\xf8\x35\xed\xca\x9a\x88\xf3\x1f\x1f\xdb\x1c\x94\x29\x56\x4b\xb8\x5b\xd3\x6e\x78\x50\x50\x0d\x53\x56\x00\xf7\x58\xf5\xb1\x0e\xb4\x8d\x37\x1d\x00\x5e\x81\x4d\xc2\x65\x4c\xc7\xf8\xd6\x88\xb5\xc5\xbf\x3f\x02\x00\x00\xff\xff\x38\xc1\xe6\xe1\xec\x0b\x00\x00"

func confAppConfBytes() ([]byte, error) {
	return bindataRead(
		_confAppConf,
		"conf/app.conf",
	)
}

func confAppConf() (*asset, error) {
	bytes, err := confAppConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"conf/app.conf": confAppConf,
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
	"conf": &bintree{nil, map[string]*bintree{
		"app.conf": &bintree{confAppConf, map[string]*bintree{}},
	}},
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
