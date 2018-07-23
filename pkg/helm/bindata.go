// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/values.yaml
package helm

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

var _valuesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x98\xcd\x6e\xe3\x36\x10\xc7\xef\x7e\x0a\x01\x7b\x69\x0e\xf5\xa9\xe8\xa1\x37\xd9\xf5\x36\x41\x36\x8d\xb0\x0e\x50\xf4\x48\x51\x63\x99\x58\x8a\x54\x87\x64\xba\xde\x6d\xde\xbd\xa0\x64\xd9\x24\x45\xc9\xa1\xda\x93\x01\xcd\xfc\x7f\x33\xfc\x1a\x0e\xfd\x21\xdb\xe6\x6a\xb5\xd3\xb4\xda\x92\x2d\xa0\xfe\x25\xfb\xfe\x3d\xdb\x10\x05\x3f\xff\xb4\x13\x54\x56\x90\xfd\x60\x3f\xe7\x6a\x73\xd2\xa0\xb2\xf5\x56\x8a\x03\xab\xd7\x57\xc1\x5d\xf6\x4f\xf6\x97\x91\x1a\xb2\xb7\xb7\xd5\x96\x3c\xc2\x29\x82\x28\x90\xbd\x12\x0d\x8f\x70\x0a\x41\x9d\x22\x60\x24\xe4\x11\xc9\xe1\x23\x4a\xa1\x0b\x94\x5f\x4f\x49\xa4\x50\xe6\x31\xf7\x80\xaf\x8c\xc2\x9e\xd5\x82\x89\x3a\x7d\x94\x11\xfd\x2c\x3f\x21\xef\x98\x34\xc6\xde\x12\x4d\xb8\x5c\xc4\xf6\xa4\x1e\x7b\xf5\x21\xa3\x52\x68\xc2\x04\x60\xc6\x1a\x52\x83\xca\x0e\x12\xb3\x56\x56\x6a\xf5\x44\x94\x06\xb4\x3b\xe5\xc1\x5a\xba\x98\x03\x39\xb0\x79\xeb\x2f\x85\x46\xc9\x0b\x4e\x04\x8c\x85\x23\xab\x2b\xfd\x5d\x56\x11\xc9\xe5\xeb\xf4\xac\x8c\x45\x11\xbb\x27\x3f\x09\x1a\x11\x0d\x5f\x5d\xd7\x02\x65\x03\xfa\x08\x46\xd9\x44\x76\x5f\x5b\x89\x1a\x70\x2c\x9e\xf1\x73\x71\x9f\xa1\x66\x4a\xe3\x69\x0c\xf0\x2c\x9e\x44\x9a\x68\x44\xe7\xbb\xeb\x9e\x7f\x33\x08\xdb\x4f\x0f\x63\x81\x67\x89\x25\xb5\x95\x42\x49\x1e\x59\x83\x98\x83\x17\x53\x28\x56\x72\x38\xcf\xfa\x06\xe5\x97\x58\xc2\x93\x5e\x2e\xea\x05\x9a\x96\x13\x7d\x8b\x35\xed\xe6\xc2\xfe\x80\x72\x72\x4c\x81\xcd\x95\x3d\xe7\x46\x1f\xbb\x62\x32\x96\x05\xb6\xf8\x56\x99\xdb\x1e\x33\xb2\x9c\x03\xea\x8d\x39\x1c\xe6\x37\x58\xe8\x36\x03\x7b\x22\x82\xd4\xef\xa0\xb9\x7e\x41\x85\x00\x4d\xab\x8c\x02\x6a\x76\x60\x94\x68\xe8\xaf\x1b\x3b\xef\x80\xa9\x85\xd4\x53\x7a\xa5\xe8\x6a\x49\xbc\xc6\xae\xa2\x11\xaf\x80\x65\x19\x9e\x75\x51\x5e\x62\x76\x83\x64\xc4\xda\x72\x06\x42\x2f\xc9\xee\xa2\x9c\x60\xa6\xb6\x01\x17\x51\xec\x6a\x40\xc9\xb3\xd6\x16\x6b\x7f\x07\xf4\xf5\x7f\xd9\x1e\x08\xb4\x5e\x54\xd7\x96\x30\x8e\x50\xe6\x31\xf3\xba\x46\xa8\x89\x96\x78\xed\x11\x52\x73\x9e\x60\xdc\x8c\x93\x30\x86\x29\x79\x64\x7e\x1e\x4d\x09\x1c\xf4\xb2\x2d\x14\x47\xdc\x8a\x92\xbc\x18\x23\x75\x24\x42\x3f\xc6\xff\x30\x0a\x1f\x30\x1f\x21\x79\x04\x81\xd6\x3f\x1d\x7e\x7f\xb1\xec\x20\x4c\x30\x66\xba\xbf\xe4\x83\x31\x25\x0f\x8f\x7a\xd3\x8d\xf8\x47\xda\xa9\xb2\xfe\xc7\x20\x29\x39\xa8\xae\xfd\xb0\xad\x69\xf7\x71\x6f\xca\x4a\x36\x84\x89\xcb\x65\xd2\x77\x8f\xe7\x40\xf6\x5a\x69\x6d\xa1\x00\xb5\x7e\x46\x7a\x04\xa5\xd1\x6e\xe9\x02\xe5\x81\x71\x58\x3f\xb7\x20\xf6\x47\x76\xd0\x4e\x0b\x33\x66\x7b\xf7\x99\x29\x39\xa3\xf7\x52\x69\x41\x1a\xf8\x7f\xa2\xfa\x4c\x37\x5a\x77\xff\xf5\x5e\x1f\x25\x36\x44\x7b\x77\xe6\xc8\x1a\xce\x22\x53\x34\xac\x9a\xde\x54\x9e\xf3\xcd\x29\x95\xc6\x1e\xf1\x61\x67\x2c\xd9\x36\x67\xc8\xc4\x86\x19\x42\x74\x23\x9d\x88\x30\xd8\x6e\x06\x58\x5f\x5c\x83\x50\x4a\x31\x29\xf6\x40\x11\xb4\xed\x8a\xc6\x41\xae\xc8\xc0\x75\x92\xb3\x13\xf4\x9d\x98\x9d\xa0\x2e\xe5\x5e\x17\x44\xa9\xbf\xab\x69\xf5\xe0\x11\xac\xda\x17\x53\x42\xbf\x4c\x6a\x95\x57\x0d\x13\x8f\x97\x0f\x91\x59\xfb\x93\x34\xfc\x89\xa0\x3a\x12\x7e\x2d\xdd\xbe\x6a\xa2\x96\x26\x21\x43\xd9\x4c\x51\xc8\x8b\x87\x64\xfc\x0c\xc1\x8b\xb4\x91\x52\xdb\xc3\xd4\xe6\x46\x4b\xd2\xb6\x28\x5f\x17\x0c\xe6\x06\xe5\x2e\x7c\x99\xa5\x0f\xc7\x13\x85\xc5\x8d\xd8\x67\xcf\xf9\x24\xae\x5e\x40\x10\xa1\x1f\x7e\xbd\x5d\x49\xf2\x6f\x43\xfd\x18\x34\x5e\x9e\xa6\x54\x14\x59\xab\x99\x14\x69\x34\x5f\xe9\x3d\x9d\xbb\xcb\xe6\x3d\xb4\xf3\xa7\x02\x99\xa0\xac\x25\x7c\x60\x0f\x84\x31\xb5\x3f\x33\xcb\xc9\xbd\xde\x7f\x34\x2a\x69\x90\xc2\x6f\x28\x4d\x9b\x32\x01\x9e\xd0\x25\x7e\x92\x94\xd8\x59\x99\x80\x0d\xe6\x60\x79\x05\x40\x05\x55\x56\x9e\x32\xd6\xd8\x47\xf7\xe5\x35\xbb\xd7\x12\x49\x3d\xd4\xb1\xe8\x7b\xd6\x77\x89\x3d\x89\xef\x5f\x5e\x0a\x67\xee\xa2\x85\x65\xec\xeb\xf5\x85\xce\xfb\xaa\xeb\x29\xbc\x2a\x36\x4d\x9d\xd7\x8d\x22\xa8\x64\x76\x44\x11\x7f\x49\xa6\x90\xe7\x54\x33\x7f\x6b\x71\x63\xeb\x9d\xb3\xf5\x23\xff\x5f\x0d\x3e\xeb\xbd\x46\x26\xea\xd8\x6a\xa5\xde\xa5\x8e\xee\x2e\xfa\x87\xc8\xfb\x5b\x2d\x57\x72\x37\xfe\xfb\x26\x39\xb3\x41\x15\x61\xa5\x64\x75\x11\xf8\x55\xf1\xb3\x11\xb6\x68\xda\x23\xc5\xfd\x93\xe1\x18\x5c\xc5\xbf\x01\x00\x00\xff\xff\x8f\x5d\x61\xdf\x69\x16\x00\x00")

func valuesYamlBytes() ([]byte, error) {
	return bindataRead(
		_valuesYaml,
		"values.yaml",
	)
}

func valuesYaml() (*asset, error) {
	bytes, err := valuesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "values.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"values.yaml": valuesYaml,
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
	"values.yaml": {valuesYaml, map[string]*bintree{}},
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
