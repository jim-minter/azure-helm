// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\xdd\x73\xdb\x36\x12\x7f\x2e\xff\x8a\x2d\xe5\xe6\xab\x01\x69\xa7\xed\xa4\xa3\x56\x99\x71\x1c\x25\x97\x8b\x1b\xeb\xe4\x64\x3a\x37\x69\x26\x03\x11\x4b\x0a\x11\x08\xb0\x00\x28\x5b\x75\xf4\xbf\xdf\x2c\x48\x7d\xcb\x76\x32\xb9\xbe\xd4\x0f\xb2\x04\x2c\x76\x17\xfb\xf1\xdb\x5d\x74\xbe\x4d\x47\x52\xa7\x23\xee\xc6\xc0\xf0\x32\x8a\x3a\xf0\xdc\x58\xf0\xe8\xbc\xd4\x45\x17\x94\x29\x80\x6b\x01\xc2\x9a\x0a\xb8\x52\xe0\x2d\xcf\x73\x99\x81\x1f\x73\x0f\x17\xa6\x56\x02\xac\xa9\x3d\xc2\x54\x72\xf0\x63\x84\x92\x3b\x8f\x16\xfa\xa7\x4f\xa3\x0e\x0c\xfb\xe7\x67\x6f\x87\x27\xfd\x17\xc3\xb3\xb7\x83\x5e\x3c\x33\xb5\x65\x16\x9d\xa9\x6d\x86\xac\xb0\xa6\xae\xe2\xa8\x03\x67\xe7\x1f\x9e\xff\xe7\xd9\xeb\x5e\x6c\x2a\xd4\x6e\x2c\x73\x9f\x1c\x6c\x9c\x4c\x8c\xe3\x02\xa7\x49\xa6\x4c\x2d\xe2\xa8\x13\x75\x40\x56\x9e\x8f\x14\x3a\x60\x2f\xe1\xe5\xeb\xc1\xdb\x37\xc0\x1c\x1c\xdc\x13\xb2\x80\xef\xdd\xd8\x58\x0f\xf1\x41\xcb\x37\x86\x4f\xe0\xb9\x54\xc0\x8e\xee\x03\xfb\x08\xa7\x67\x2f\x80\x31\x65\x0a\x56\x59\xcc\xe5\x25\xc4\xaf\xde\x3e\xed\x03\x91\xc2\xb3\xe1\xd9\xa0\x1b\x7f\x1d\x7f\xe2\x11\x45\x57\x57\x20\x73\x48\x4e\x8c\xce\x65\x91\x9c\x63\x56\x5b\xe9\x67\x03\xee\xb3\xf1\x80\x67\x13\x5e\xa0\x83\xf9\x3c\x52\xa6\x28\xd0\x02\xf3\xad\xe1\x98\xf3\xdc\xfa\xba\x4a\xdc\x18\x62\xa9\x9d\xe7\x4a\x49\x5d\x80\x45\x01\x64\xf2\x4c\x68\xc8\x02\xcf\xda\x72\x2f\x8d\x06\xa3\xe1\xe0\xde\xd8\x38\xaf\x79\x89\xf7\xe3\x28\xe3\x1e\x9e\xa4\x53\x6e\x53\x25\x47\xe9\xac\x2e\xd3\x4c\x49\xd4\x9e\x65\x68\x7d\x52\x61\x09\xbf\xfe\x7a\xb7\x7f\xf6\xfc\x2e\xa9\x78\x82\xd6\x1f\xbb\xa7\x33\x8f\x6e\xa9\x2b\xad\xc9\x5c\x66\xdc\xa3\x4b\x5a\x5d\x87\x58\x19\x27\xbd\xb1\xb3\xb0\x0d\x9f\xe0\xdc\x5b\xd2\x6b\x3e\x8f\xfa\x67\xcf\xaf\x17\x3a\xc1\xd9\xb6\xcc\x81\x95\x53\xee\xf1\x15\xce\xbe\x50\xf2\x2b\x9c\xed\x08\xee\xc0\x9b\xb3\x67\x67\x5d\x10\xa8\xd0\x63\x88\xc0\xdc\x28\x65\x2e\x88\xc6\x61\x16\x4c\xc4\x73\x0a\x49\x0a\xdf\x4c\xd5\x64\x65\x07\xdc\x22\xd8\x5a\xc3\x85\xf4\x63\xe0\x30\x2d\x41\x96\xbc\xc0\xe6\xf7\x44\x66\x93\xe0\x87\xc4\x62\x65\x60\xc4\x27\x28\x40\xea\xe6\x96\x90\xa2\xcf\xe8\x8a\x61\xd3\x25\x22\xdd\x22\x5f\x5c\xf5\x9d\x1d\xa3\x62\x8f\x99\x43\x3b\x45\xcb\x6c\x55\xba\xf7\x11\x39\xa9\x37\x44\x01\xff\xe2\x1e\xfa\xda\xa3\xad\xac\x74\x08\xa7\x52\xd7\x97\xf0\x18\xce\x03\x31\xdc\x1b\x0e\x7e\x73\xf7\xa3\x11\x77\x58\x5b\xd5\x1b\x7b\x5f\xb9\x6e\x9a\x66\x42\x27\x16\xc5\x98\xfb\x24\x33\x65\x9a\x19\xed\x51\xfb\x54\x48\xe7\x53\x92\x96\x36\xb2\xd2\xc7\xe9\xe3\x86\x51\x7a\x40\x2c\xb8\xcd\xc6\xa9\x71\x51\x51\x15\x13\x9c\xf5\x72\xa9\xb0\x9b\xa6\xe1\x1e\xd5\x44\xa6\xb6\x2a\x59\x51\x15\xe9\x70\xf0\x1b\x7b\x31\x78\xc1\x5e\xf5\xff\xcb\x1a\x29\xcc\xa2\x42\xee\x30\x72\x4e\x65\x9c\x02\xa8\x17\x4e\xd9\xb1\x2b\xd3\x8c\xa7\x2d\x55\x8d\x15\x39\x39\x50\x05\xb7\x37\x94\x37\xc4\xe0\x8a\x94\x14\xba\x21\x70\x22\xd4\x94\x87\xa2\x37\x43\x17\x6d\x9b\x14\x2f\xbd\xe5\xee\xcb\x2c\xcb\xa0\x1f\x4e\xfd\x1d\x26\x6e\xf4\xf9\x27\x5a\xda\x38\x64\x3f\x24\x47\x47\xfb\x6c\x7d\x56\xa1\x3e\x27\xd4\x86\x13\xa3\x3d\x97\x1a\x2d\x0c\x14\xf7\xb9\xb1\x25\xd0\xa1\xbf\x27\x9a\x31\x25\xde\xff\x04\x63\x13\x8e\x7d\x76\x21\x38\x1e\x9e\x11\xae\x85\x6a\x02\x75\x25\x08\x2b\xe1\xdd\xd5\x55\x8b\x8b\xee\xdf\x46\xea\x5b\xca\x4e\xfc\x10\x62\x98\xcf\xdf\xef\x94\x8e\xdc\x58\xe0\xde\x63\x59\x79\x90\x1a\xae\x8e\x92\xe4\xa7\xf9\x2f\x20\x4c\x04\x30\xab\x4b\x68\xd5\x00\x36\x03\xf6\x27\x7c\x99\xcc\x20\x12\xee\xdc\x81\x91\x45\x3e\x89\x00\x6e\xbc\xf0\xbb\x85\x1a\x07\x57\xed\xb7\xf9\xfb\xfd\x57\x6f\x75\x6a\x6a\x61\xce\xa5\x42\x11\x47\x40\xb5\xf7\xdd\xbb\xb5\xd3\xc0\x94\x87\x9f\xe0\xfd\xfb\x5f\xa8\x46\x68\x70\x0a\xb1\x82\xa3\x5f\x00\x95\x43\xc0\x4b\xe9\xe9\x47\x2e\x23\x61\x34\xde\xe2\x0d\x8b\xa5\x99\x7e\x59\x51\x26\xeb\x65\x0a\xb9\xa6\x2a\x14\xd9\x12\x98\xcd\xe1\xc6\x22\x7d\x53\xf0\x5c\x5d\xa1\x16\xf3\x79\x14\xc9\x1c\xbe\x85\xc2\x62\xb5\x22\x17\x26\x9b\xa0\x6d\x8a\x54\xee\x3c\x1f\x35\x17\x8e\x00\xdc\xcc\x79\x2c\x33\xaf\xc0\x79\x53\x41\x43\xc8\x82\x56\x75\x95\x78\x59\xa2\xbd\x95\x8a\x32\x52\x66\x78\x1d\xdd\xda\x7e\x39\xc9\x5d\x72\x99\x3b\x60\x39\xa4\x02\xa7\x94\xd6\x93\x94\xff\x55\x5b\x4c\x97\x2d\x60\xc5\xad\x3f\x8a\x00\x30\x1b\x1b\xb8\x7b\x33\x19\xec\xdc\x11\x88\x3d\x14\xb6\xfa\xb3\x36\x9e\x03\x1c\xc2\xe1\x5d\x78\xf2\x64\x75\x75\x52\xc3\xd4\xda\x6f\x9f\x8c\x00\x2c\x3a\x6f\x2c\x66\x46\x03\x1b\xee\xd9\x6f\x1a\x1a\xe2\xd4\x2c\xa5\x82\x63\x69\x74\xf2\xd1\x19\xbd\x6a\x66\x22\x80\x98\x1a\x49\x61\xe5\x14\x6d\xdc\x85\xf8\xa3\xa9\xad\xe6\x4a\xc4\x0f\x69\x4f\x48\x47\x69\xce\x14\x16\x3c\x9b\x31\x8b\x85\x74\xde\xce\xe2\x2e\x78\x5b\x63\xd4\xb4\x30\x9b\xb6\xe4\xd6\xef\x1a\x73\x3f\xc1\x96\xef\x72\x19\x45\xad\x65\xaa\x5a\x29\xca\xcf\x45\x4e\xbe\xa4\xde\xc6\x25\xaf\x8d\xc0\x90\x85\x4f\x82\xa9\x35\x51\xdd\xd9\x1b\x45\xe8\x33\xb1\x2f\x86\x96\x5e\xdd\xf6\x95\xcb\x9c\x3c\x4a\x55\xad\x0f\xe1\xd3\xa7\xe6\x76\xd7\xb9\x75\x8d\x74\x4b\x60\xe3\x50\x81\x39\xaf\x95\x77\x9f\xe5\x50\x3a\x77\xbd\x3b\xc3\x2e\xd9\x85\x90\x4d\xb8\x80\x6a\x3e\xab\x1e\xfe\xfc\xe3\x8f\x3f\x06\x5c\xfb\xa6\xb2\xc6\x9b\xde\xc1\x95\x70\xfe\xbb\xef\x1e\x3e\x98\x47\xdf\x54\xc6\xfa\x66\xa1\xd3\x79\xf0\x70\x1e\x7d\xb3\x1a\x05\x8e\xc3\xa8\xf2\x72\xd8\xff\xfd\xf8\xf4\xf4\xc3\xf1\xe9\xe9\xd9\xef\xc0\x2a\x38\x08\x4c\x80\x95\xe4\x1d\x8f\xc0\x58\xf3\xff\x75\xff\x77\x5a\x5c\x6c\x33\x41\xac\xe1\x20\x7c\xb2\x8f\x70\x7c\x72\xd2\x1f\xbc\x01\x76\xd1\x62\xce\x42\x0e\x73\x7c\x8a\x6d\xf0\xb9\x99\x6b\x80\x25\x5d\xec\xd2\x88\x76\x11\x10\x8c\x22\x81\x50\x48\x93\x57\x2f\x38\x2f\x50\xfb\x30\xac\x69\xf4\x17\xc6\x4e\xa0\xf6\x52\x49\x2f\xd1\x41\x61\x02\x52\x7a\x03\x96\x67\x48\x58\x25\x24\xe1\x54\x42\x93\x4e\xbe\x3c\x6c\x6b\xed\x60\x84\xb9\xb1\x08\x42\x3b\x90\x0e\x26\xda\x5c\x68\xf0\x26\x74\xd6\x8d\x24\x04\xd4\x02\xea\xaa\xe9\x95\x09\x5d\x67\xe0\x42\x21\x88\x2e\xc6\x52\x61\x00\xde\x25\xf8\x01\x13\xf7\xa1\xd7\x83\x38\x0e\xe0\x2b\xcc\x0a\x7a\x9b\x6b\x37\x67\xbe\x85\x9b\x43\xf7\xbc\xc1\x5f\x98\x2f\xea\x51\xcb\xa5\xb1\x9d\x43\x0f\xdf\x5f\x46\x78\x19\x6c\x7b\x7e\x7c\xfe\x76\xf8\xb2\x77\x77\x8d\xcb\x6f\x01\xc5\x5b\x26\xcd\x3e\xcc\xe7\x77\xc3\x41\x76\xb9\x48\x1b\x1a\x08\x18\xab\xac\x9c\x4a\x85\x05\x0a\x60\x8c\xa0\x9a\x2d\x0c\x4a\x77\x02\x36\x85\xb4\x9b\xd2\xd7\xee\x5f\xc0\xb0\x95\x76\xa3\xca\xd0\x56\x8f\xa8\xd6\x24\xb0\x39\x11\x45\x4d\x09\x63\x19\x67\xde\xd6\xce\x93\x67\x07\x52\xc3\xa4\x1e\x61\xe3\x74\x47\x86\xaf\x1d\x82\x32\x19\x57\xc0\x2b\xd9\xf6\x63\x91\x23\xe5\x24\x30\x8b\x10\xbb\xce\x3d\x78\xd0\xac\x77\xe1\x7e\xf2\xa0\xf3\xc7\xd1\xa2\xc3\x5a\xab\x41\x9d\xb8\xc9\x67\x63\x65\x21\x75\xda\x94\xb5\x74\x39\x71\xb3\x66\x21\x59\x09\xff\x7a\x19\x14\x2e\xe1\xe3\xff\xcf\xd5\x09\xfd\xf5\x4c\xad\x31\x3e\x0d\x6c\xd2\x96\x4f\x74\x75\xc5\x28\x21\x34\xc2\x41\xf2\x94\x67\x93\xba\x7a\xaa\xcc\xe8\x35\xc5\x71\x1c\xdf\x3a\xaf\x2f\x53\x92\x90\x68\x8a\x76\xb6\xd3\x07\x44\x1d\x70\x9e\xe2\x16\x0a\xf4\x21\xa7\x46\x41\x4a\x68\x09\x86\xf9\x26\x72\xa5\x0f\x22\x6a\x5e\x48\x8f\x67\xd2\xf6\x36\xf7\xda\x73\xe5\x44\x48\x0b\x07\x6b\x74\xb7\x34\x2f\xc2\x5c\x68\x65\xb8\x20\x35\x1b\x1e\xf1\x67\xa6\x60\xdf\x67\xa2\xb1\xc9\x35\x59\xb8\x91\x45\xbb\x89\xf3\x47\x04\x21\x79\x76\xa2\xb0\xbb\xbb\xb4\x8f\x38\xbc\xfe\x54\xd6\x4c\xa5\x40\x9b\x76\xd3\x0f\x82\x7b\x9e\x7e\x30\xf5\x92\xf5\xba\x19\xba\xa9\xa9\x29\x43\x69\xeb\x96\xbb\x00\x19\xb4\xb1\x45\xc3\x89\x8d\x5a\xa7\xf7\xe8\xe4\x56\x1c\xcc\xe7\x2d\x91\x08\x8f\x64\xa1\xdf\xeb\x91\xb0\xd6\x25\x89\x18\xb5\x04\x3c\xbc\x3e\xf4\x16\x16\xbf\xd9\x2f\xad\xfc\x05\x31\x75\xb1\x8b\x60\x79\xb4\x28\x6f\xb7\xb5\xa5\x44\x14\xde\x3d\x34\xaf\xdc\xd8\xf8\xcf\xf5\x6c\x83\x8e\x64\x93\xaf\xf7\x2c\xd9\xb2\xbb\xfc\xb6\xdc\x5a\x8f\xdd\xee\xe6\xaf\xc6\x47\x0c\xa1\xff\xe6\xe4\xd9\xc9\x9b\xd3\x0f\xc7\x83\x97\xbd\xf8\x87\xf8\x1a\xd7\x6d\x28\x1b\x68\x88\x4b\xe8\x8d\xda\x6b\x2f\xcc\xb5\x11\x0f\x3b\xde\xa1\xe8\x61\x94\x3c\x9b\x79\xa5\xf1\xa2\x25\x08\xd5\x6b\x2d\x7b\xdb\x65\xa9\xa5\x97\x5c\xb1\xf6\x29\x09\xe2\xd6\x15\x87\xe1\x6f\x39\xdc\x6e\xac\x76\x1f\xfd\xf0\xf3\xe1\xc3\xf5\xa5\xa3\xbd\x84\x47\xbb\x84\x8f\xf6\x12\x3e\x0a\x84\xf1\x7e\x95\x98\x37\x13\xd4\xc1\x2c\x2c\x37\x96\x85\xb6\x6b\x8b\x94\x8b\x29\x5a\x2f\x1d\xb2\x0a\xd1\xb2\xda\x2a\x07\x7b\x60\x32\x88\x89\xa2\x72\xba\x6b\xa5\xf4\xc1\xd6\xda\xce\x58\xb3\xb4\xe7\x06\x3c\x6d\xb4\x6a\x5b\x7c\x3f\x27\xc0\x31\xf4\x0d\x71\x80\x6a\x6a\x44\x68\x0e\xf2\xb5\x46\xc1\xb8\x28\xa1\xb2\x26\xa7\x90\x5f\x15\xb6\xcc\x68\x6f\x8d\x62\x95\xe2\xd4\x6e\x74\xa8\x81\xe1\xca\x19\xd0\x88\x62\x45\x97\x84\x1a\x9b\x4c\x8d\xaa\x4b\x74\x40\x81\x91\x59\xe4\x1e\xc5\xa2\x23\xa2\xe6\xb4\x99\x34\x32\xea\x83\xa8\x59\x6a\xe1\x97\x95\x70\xf8\xf8\xa7\x43\xea\x05\x97\xd7\x69\x41\xeb\x1a\xfe\xa4\x47\x33\x1c\x84\x72\xe1\x66\x4e\x99\x02\x9c\xd4\x59\x68\xb0\x4a\xae\x79\x81\x80\x54\x43\xfc\x98\x48\xfc\xd8\x9a\xba\x18\xc3\x62\xbe\x88\x56\x23\x41\x3b\x64\x2c\xb8\x2c\xc7\x86\xad\x09\x6d\x7b\x3b\xea\x80\x36\x1e\xbb\xc0\xbd\x29\x65\xc6\x56\x16\x0b\x6d\x5e\x66\xb9\x1b\x83\x32\xa6\x72\x50\x6b\x2f\xd5\xe2\x65\x5f\x3a\xa8\xab\x35\xe6\xcd\x4b\xc6\x7e\x2e\x7b\x75\xa1\x01\xe6\x46\x6a\xb8\xf3\xb9\xaf\xe7\xe4\x42\xc7\x28\x2e\xda\xc2\xcb\x98\xc5\x91\x31\x64\x33\xbf\x31\xe6\x7c\xfa\x04\x57\xb7\x3d\x3a\x34\x47\x89\xcd\x7a\xd6\x7b\x03\x99\x29\xab\xf0\xbc\xbc\xef\x05\x22\x26\xb0\x1c\xd7\x9e\x90\x7b\x29\x1f\xb4\xb9\x88\xe6\xab\x41\xfd\x7f\x01\x00\x00\xff\xff\x38\xd1\xf9\x7b\x68\x19\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\x5f\x6f\xdc\xc6\x11\x7f\xdf\x4f\x31\x3e\x19\x56\x8c\x66\x49\x0b\x85\x21\x40\x8e\x0c\xb8\xae\x92\xa6\x6e\x2a\x41\x6a\x51\x14\x82\x1e\xf6\x76\x87\xe4\xe6\x96\xbb\xf4\xec\xf0\x24\xe6\x7c\xdf\x3d\x18\xf2\xce\x52\x24\xf9\x1c\x3d\xe4\x25\x6f\x3a\xed\xfc\x9f\xf9\xfd\x66\xb8\xf7\xac\x9c\xfb\x58\xce\x4d\x6e\x40\xe3\x8d\x52\xab\x15\xf8\x0a\x8a\xf7\x29\x56\xbe\x2e\x2e\xd0\xf6\xe4\x79\x38\x33\x6c\x9b\x33\x63\x17\xa6\xc6\x0c\xeb\xb5\x0a\xa9\xae\x91\x40\x33\xc4\xe4\x50\x67\x36\xc4\x7d\x57\xe4\x06\x66\x3e\x66\x36\x21\xf8\x58\x03\xa1\x83\xc6\x30\x58\x17\xc1\x8e\x16\x7b\x32\xec\x53\x84\x14\xe1\xf9\x37\x4d\xca\x1c\x4d\x8b\x2f\x67\xca\x1a\x86\xb7\xe5\xd2\x50\x19\xfc\xbc\x1c\xfa\xb6\xb4\xc1\x63\x64\x6d\x91\xb8\xe8\xb0\x85\xef\xbe\xdb\x3f\x39\xfd\x7e\x5f\x02\x7c\x8f\xc4\xef\xf2\xdf\x06\xc6\xfc\x39\x52\xf9\x9f\xaf\xbc\x35\x8c\xb9\xd8\x44\x7a\x8e\x5d\xca\x9e\x13\x0d\xe3\x33\x7c\x82\x0b\x26\x89\x6b\xbd\x56\x27\xa7\xdf\x7f\xd9\xe9\x02\x87\xfb\x3e\xcf\xc8\x2f\x0d\xe3\x07\x1c\x9e\xe8\xf9\x03\x0e\x0f\x1c\xef\xc1\x7f\x4e\xff\x7e\x7a\x04\x0e\x03\x32\x02\x37\x08\x55\x0a\x21\x5d\x8b\x4c\x46\x3b\x96\xc8\x54\x8c\x04\x26\x04\xb0\xa1\xcf\x8c\x94\xc1\x10\x02\xf5\x11\xae\x3d\x37\x60\x60\xd9\x82\x6f\x4d\x8d\xd3\xef\x85\xb7\x8b\xb1\x0f\x05\x61\x97\x60\x6e\x16\xe8\xc0\xc7\x29\x4b\x28\x91\xad\xa4\x38\x3e\xe6\xc2\x95\xf7\xc4\xb7\xa9\x5e\x52\x83\x41\x1f\xea\x8c\xb4\x44\xd2\xd4\xb5\xf9\x4a\x49\x93\x8e\xcf\xd1\xc1\x3f\x0c\xc3\x49\x64\xa4\x8e\x7c\x46\xf8\x97\x8f\xfd\x0d\x1c\xc2\xc5\x28\x0c\xdf\x9c\x9f\xfd\x94\x5f\xaa\xb9\xc9\xd8\x53\x38\x6e\x98\xbb\x7c\x54\x96\xd6\xc5\x82\xd0\x35\x86\x0b\x9b\xda\xd2\xa6\xc8\x18\xb9\x74\x3e\x73\x29\xde\xca\xc9\x57\x79\x58\x1e\x4e\x86\xca\xe7\x62\xc2\x90\x6d\xca\x94\x55\xdd\xd5\x0b\x1c\x8e\x2b\x1f\xf0\xa8\x2c\xc7\x3c\xba\x85\x2f\xa9\x6b\x75\xdd\xd5\xe5\xf9\xd9\x4f\xfa\x87\xb3\x1f\xf4\x87\x93\xff\xeb\xc9\x8b\x26\x0c\x68\x32\xaa\x9c\x83\x35\x32\x40\xc7\xa3\x16\x35\xb9\x2d\xad\x29\x37\x52\x3d\x76\xd2\xe4\x51\x6a\x6c\xfb\x24\xb9\x63\x06\x6f\x45\x25\xa0\x1d\x83\xa3\x30\x9a\x79\x40\x77\x3c\x60\x56\xf7\x4b\x8a\x37\x4c\x26\x3f\xad\xb2\x1a\x4e\x46\xad\x3f\xa2\xc4\x53\x3c\x7f\xc6\x4a\xa7\x8c\xfa\xaf\xc5\xc1\xc1\x63\xb5\x3e\xed\x30\x5e\x34\xbe\x62\x78\x9f\x22\x1b\x1f\x91\xe0\x2c\x18\xae\x12\xb5\x20\x4a\x7f\xcc\x34\x63\x29\xb6\xff\x0c\xc5\x16\x1e\xfb\x9d\x6b\xe0\xdd\xf9\xa9\xb0\xda\xb8\x49\xa0\xef\x9c\x30\x25\x5c\xae\x56\x1b\x56\xcc\xff\x4c\x3e\x7e\x65\xe5\xcc\xbe\x85\x19\xac\xd7\x57\x0f\x16\x47\x95\x08\x0c\x33\xb6\x1d\x83\x8f\xb0\x3a\x28\x8a\xd7\xeb\x37\xe0\x92\x02\x18\xfa\x16\x36\x61\x80\x1e\x40\x7f\x84\xa7\xf9\x1c\x5d\xc2\x8b\x17\x30\x27\x34\x0b\x05\xb0\x23\xdd\xcb\x6d\x10\xcf\x57\x9b\xbf\xd6\x57\x8f\x27\xbe\x89\x68\xda\x83\x95\xf1\x01\xdd\x4c\x81\x6c\xdd\xcb\xcb\x3b\xda\xa0\x03\xc3\x6b\xb8\xba\x7a\x23\xfb\x21\x42\x0e\x88\x1d\x1c\xbc\x01\x0c\x19\x01\x6f\x3c\xcb\x8f\xca\x2b\x97\x22\xee\xec\x04\x61\x9b\x96\x4f\x5b\xc7\x52\x39\x1b\xd0\x44\xd9\x3f\x8a\x5a\xd0\x54\xc1\xce\xf5\xbc\x6b\x6c\x56\x2b\x8c\x6e\xbd\x56\xca\x57\xf0\x0c\x6a\xc2\xee\x56\xdc\x25\xbb\x40\x9a\xd6\x53\x95\xd9\xcc\xa7\x74\x15\x40\x1e\x32\x63\x6b\x39\x40\xe6\xd4\xc1\x24\xa8\xc7\xa8\xfa\xae\x60\xdf\x22\x7d\x55\x4a\xb0\xe8\x2d\x7e\x49\xee\xce\x7b\xbb\xa8\x72\x71\x53\x65\xd0\x15\x94\x0e\x97\x02\xe8\x45\x69\x7e\xe9\x09\x4b\xc2\x9c\x7a\xb2\xa8\x3b\x43\x7c\xa0\x00\xd0\x36\x09\xf6\x77\x8b\xc1\x83\x1c\x41\xcc\x43\x4d\xdd\xc7\x3e\xb1\x01\x78\x05\xaf\xf6\xe1\xed\xdb\xdb\xd4\x25\x8c\xd4\x47\xbe\xaf\xa9\x00\x08\x33\x27\x42\x9b\x22\xe8\xf3\x07\xef\xab\x95\x96\xf9\xc1\x8f\x50\x9c\xa7\x80\x02\xbe\x8a\x8c\x4c\xaf\x02\x98\xae\x1c\x71\x32\x49\x97\xce\x60\x9b\x62\xf1\x73\x4e\xf1\xf6\xc2\x51\x00\xb3\x90\x6a\xed\xc8\x2f\x91\x66\x47\x30\xfb\x39\xf5\x14\x4d\x70\xb3\x6f\xe5\xcd\xf9\x2c\xd8\xd7\x01\x6b\x63\x07\x4d\x58\xfb\xcc\x34\xcc\x8e\x80\xa9\x47\x35\xdd\x35\x12\x07\x46\x37\xf9\xbd\x5b\x71\x43\xfc\xb0\xe4\x8f\x0b\xdc\xeb\x70\xe5\x95\xda\xd4\xaf\xeb\x43\x10\x04\x6f\x51\xfb\xa3\xdc\x3e\xb9\xf8\x77\x72\x38\xe2\xf4\xed\xd8\x90\x28\x52\x2f\x94\xda\x83\xeb\x11\x36\x62\x5a\x86\x5f\x60\x01\xd7\xc6\xd4\x18\x19\x4c\x74\x10\x91\xaf\x13\x2d\xa0\x67\x1f\x3c\x7b\xcc\x50\xa7\x11\x9e\x9c\x80\x8c\x45\x81\x88\xf3\x02\x8f\x42\xed\x49\x79\xb7\xca\xd4\xc7\x0c\x73\xac\x12\x21\xb8\x98\xc1\x67\x58\xc4\x74\x1d\x81\xd3\x78\xca\x4d\x9e\x70\xac\x44\xdf\x4d\xc7\x99\x40\x7a\x80\x3c\x72\x8f\xba\x6e\x7c\xc0\x11\xed\x9f\x31\x07\xda\xbd\x84\xe3\x63\x98\xcd\x46\xc4\xbb\x74\x8b\xf7\x09\xdf\x93\xce\x33\xd8\x5d\x8b\x8b\x09\xf6\xb0\xde\x52\xe0\xc6\xca\x44\x12\x19\x19\xfe\x72\xa3\xf0\xa6\x4b\xc4\x70\xf1\xee\xe2\xbf\xe7\x3f\x1e\xef\xdf\xb1\xf2\xbf\x44\x0b\xa4\x8d\x91\xe9\x1d\xd6\xeb\xfd\x51\x51\xdf\x6c\xfb\x20\x17\xa8\xd6\x1d\xf9\xa5\x0f\x58\xa3\x03\xad\x85\x21\xf4\xb6\xa0\x92\x13\xe8\x25\x94\x47\xa5\xfc\x79\xf4\x0b\x68\xdc\x78\xdb\x19\x32\x6c\x48\x4b\xf5\x51\x1c\x4e\x1a\x4a\x4d\xbc\xa9\xad\xd1\x4c\x7d\x66\xb5\x6b\xda\xb9\x8f\xe8\xb4\x71\x2d\x74\x94\x64\xb5\x42\xea\x30\x66\xd9\xf3\x5a\x36\x35\xa5\xa0\xbb\x60\x22\x4e\xa3\x2a\x4c\xfa\x15\x2d\x69\xe6\xdd\xb9\x96\xb9\x42\x30\x21\x27\x88\x88\xee\x56\xb2\x08\xc9\x9a\x50\x2c\x53\xe8\x5b\xcc\xe0\x3c\x81\x25\x34\x8c\x6e\x3b\x2b\x02\xff\x09\xfa\x56\x26\x44\xc6\x48\xb5\x0b\x11\xd4\x2d\xbc\x3a\x7c\xfd\x0a\xf4\x1d\x62\x4c\xe4\x6b\x1f\xcb\x2f\xd8\x97\x38\x26\x48\x8e\xc4\x9e\x87\x1c\x52\x0d\xd9\x47\x3b\x8e\x5e\x6b\xa2\x7c\x17\xe0\x12\x69\xe0\x46\x44\xb8\xa1\xd4\xd7\x0d\x6c\x51\xad\x6e\xd1\xb7\x81\xf6\xd6\xca\x67\x84\xde\xa3\xcc\xfb\xcf\x6a\x0f\x62\x62\x3c\x02\xc3\xa9\xf5\x56\xff\xb6\x66\x60\x49\xbe\x29\x43\x4a\x5d\x86\x3e\xb2\x0f\xd0\x1a\xf9\x8a\x11\xb8\xf4\xdd\x1d\xe3\xd3\x51\xf1\xb8\x95\x47\x63\x11\xae\xd8\x29\x0d\x2f\x7e\xef\x67\xac\xb4\x30\x6b\x21\xd7\x0d\x4b\x68\x4d\x38\x4f\x49\x6a\xc6\xbf\x61\x94\x4f\x9f\x60\xb5\xfb\x02\x98\x14\xc5\xc8\x9d\x55\x2a\x9c\x60\x53\xdb\x8d\xdf\x79\x8f\x9d\x03\xb2\xfa\x73\xd3\xb3\x13\x02\xd9\x7a\x87\x98\xae\xd5\xfa\x76\x6f\xfe\x1a\x00\x00\xff\xff\x93\xe0\xdc\x12\xa6\x0f\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
