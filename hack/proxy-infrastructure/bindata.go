// Code generated by go-bindata.
// sources:
// data/start.sh
// DO NOT EDIT!

package proxyinfrastructure

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _startSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xdd\x4e\x1b\x3d\x10\xbd\xf7\x53\xf8\xe3\xab\xca\x95\x6d\xa5\x45\xa8\x42\x2c\x12\x84\x54\x54\x55\x0b\x52\xfa\x73\x81\x72\xe1\xf5\x4e\x36\x16\x5e\x7b\x65\x8f\x69\x36\x34\xef\xde\xd9\x2c\x25\x69\x50\x2b\x6e\xe2\x9f\xe3\x39\x39\x73\xe6\xec\xff\xff\xa9\xd2\x7a\x55\xea\xb4\xe0\x02\x96\x8c\xc1\x12\x0c\x7f\x73\xf6\x7a\xd4\x6f\xdb\x10\x91\x5f\x5d\x7f\x9a\x14\x2a\x86\x80\xcc\x54\x8c\x75\xb9\xe1\xd6\x27\xd4\xce\xf1\xda\x22\x17\x1d\x63\x26\x47\xc7\x45\xe2\x0b\xc4\x36\x9d\x28\x95\x30\x44\x5d\x83\xac\x43\xa8\x1d\xe8\xd6\x26\x69\x42\xa3\xea\xe0\xb4\xaf\x69\x19\xc9\xd1\x48\x1e\x4b\x67\x7d\x5e\x0a\xdd\x54\xc7\x47\x12\x75\x94\xf5\x8a\xff\xe4\xb4\xe1\x62\xcc\x55\x4e\x51\xb9\x60\x34\x11\x2f\x57\xec\xe6\xfc\xcb\x55\xf1\xaa\xff\x3d\xd9\x22\xc4\xd4\xab\x67\xac\x0e\xbc\x06\x92\x92\x7b\x45\x8b\x5c\x6e\xfe\x2d\xb4\xe0\xd3\xc2\xce\x71\xbb\x13\x7a\x95\x23\x28\xd3\x54\xaa\x8d\x61\x49\xca\x9b\xbb\xca\x46\x5e\x69\xd4\xcc\x68\xe4\x67\x9b\x6b\x61\x20\xa2\x6c\xa1\xe1\xa7\xa7\x87\x93\xeb\xf7\x87\xec\xe1\x81\xcb\x71\xf0\x73\x5b\xcb\x29\xc4\x7b\x88\x63\x7a\x41\x62\xfb\xe5\x3c\x5d\x74\x08\x89\x4e\x53\x8c\xd6\xd7\x7c\xbd\x66\x54\xc4\x76\x09\xef\xa0\xfb\x37\xdf\x47\xe8\x88\xe0\x26\xda\x7b\x8d\x40\x87\x97\x90\x1a\xfd\x37\xce\xb1\x7e\x91\x36\x05\x68\x54\xea\x12\x02\x19\x32\xac\x4a\xc7\x20\x36\xf4\x32\x91\x30\x6b\xe0\x89\xff\xf6\xab\xb7\x38\x63\x97\x90\x4c\xb4\x2d\xda\xe0\x8b\xa7\xc7\x8c\xdd\x4e\x87\xe7\x33\x36\xa1\x04\x4d\x69\x8a\x38\x84\xe6\x71\x48\x83\xe1\x7c\x63\x2d\x1f\x80\x3d\xaf\x7b\x93\xfe\x40\x7e\x9b\x46\x8d\x3e\xaf\x1a\x7a\x17\x29\x97\x9e\x26\x7f\xb0\xd3\xfc\x67\xc0\x4b\x98\x5b\x12\x4b\x12\xe5\xb7\x1e\x5e\xaf\x0f\x48\xe1\x87\x21\xb5\x33\xf6\x5d\x7b\x84\xea\xa2\x2b\x9a\xec\xd0\x8a\x4c\x9d\xf6\xf9\xa3\x08\x0d\xe6\xa4\x5c\x05\x3e\xf8\x61\xd0\x51\x3a\xa0\x09\x5e\x44\x70\x41\x57\xcf\x50\xf0\xba\x74\xc0\xb7\x56\xec\xe1\xa9\xb7\x62\x07\x66\x73\x1b\xe1\x07\xe9\x10\x14\x43\x2e\xc4\x2a\x78\x28\xda\x5c\x3a\x6b\xe8\xa4\xab\x4a\xf4\x5f\x5d\xf1\xee\xe8\xe8\xad\x42\xd3\xd2\x5d\x0b\xb1\xd1\x1e\x3c\xee\x97\x3e\x2a\xfa\x15\x00\x00\xff\xff\x2f\xc9\x8a\x4a\xc4\x03\x00\x00")

func startShBytes() ([]byte, error) {
	return bindataRead(
		_startSh,
		"start.sh",
	)
}

func startSh() (*asset, error) {
	bytes, err := startShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "start.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"start.sh": startSh,
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
	"start.sh": {startSh, map[string]*bintree{}},
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
