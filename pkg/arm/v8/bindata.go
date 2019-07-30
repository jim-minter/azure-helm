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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x6d\x73\x14\x37\x12\xfe\x1c\xfd\x8a\xce\xac\x13\x5e\x82\x66\x6c\x92\x1c\xa9\x4d\x96\x2a\x63\x16\x8e\xc3\x60\x9f\x0d\x97\xba\x22\x14\xa5\x1d\xf5\xcc\x8a\x9d\x91\x26\x92\x66\xed\x8d\xd9\xff\x7e\xd5\x9a\xd9\xf7\xf5\x82\x43\xb8\x0f\xe1\x83\xf1\x48\xad\xee\x56\xeb\xd1\xd3\xdd\x72\xe7\xeb\x64\xa0\x74\x32\x10\x6e\x08\x1c\x2f\x19\xeb\xc0\x13\x63\xc1\xa3\xf3\x4a\xe7\x5d\x28\x4c\x0e\x42\x4b\x90\xd6\x54\x20\x8a\x02\xbc\x15\x59\xa6\x52\xf0\x43\xe1\xe1\xc2\xd4\x85\x04\x6b\x6a\x8f\x30\x56\x02\xfc\x10\xa1\x14\xce\xa3\x85\xfe\xf1\x23\xd6\x81\xb3\xfe\xf9\xc9\xeb\xb3\xa3\xfe\xd3\xb3\x93\xd7\xa7\xbd\x68\x62\x6a\xcb\x2d\x3a\x53\xdb\x14\x79\x6e\x4d\x5d\x45\xac\x03\x27\xe7\xef\x9e\xfc\xfb\xf1\xcb\x5e\x64\x2a\xd4\x6e\xa8\x32\x1f\xef\xad\xac\x8c\x8d\x13\x12\xc7\x71\x5a\x98\x5a\x46\xac\xc3\x3a\xa0\x2a\x2f\x06\x05\x3a\xe0\xcf\xe0\xd9\xcb\xd3\xd7\xaf\x80\x3b\xd8\xbb\x2d\x55\x0e\xdf\xb9\xa1\xb1\x1e\xa2\xbd\x56\x6f\x04\x1f\xc0\x0b\x55\x00\x3f\xb8\x03\xfc\x3d\x1c\x9f\x3c\x05\xce\x0b\x93\xf3\xca\x62\xa6\x2e\x21\x7a\xfe\xfa\x51\x1f\x48\x14\x1e\x9f\x9d\x9c\x76\xa3\xcf\xd3\x4f\x3a\x18\xbb\xba\x02\x95\x41\x7c\x64\x74\xa6\xf2\xf8\x1c\xd3\xda\x2a\x3f\x39\x15\x3e\x1d\x9e\x8a\x74\x24\x72\x74\x30\x9d\xb2\xc2\xe4\x39\x5a\xe0\xbe\x0d\x1c\x77\x5e\x58\x5f\x57\xb1\x1b\x42\xa4\xb4\xf3\xa2\x28\x94\xce\xc1\xa2\x04\x0a\x79\x2a\x35\xa4\x41\x67\x6d\x85\x57\x46\x83\xd1\xb0\x77\x7b\x68\x9c\xd7\xa2\xc4\x3b\x11\x4b\x85\x87\x87\xc9\x58\xd8\xa4\x50\x83\x64\x52\x97\x49\x5a\x28\xd4\x9e\xa7\x68\x7d\x5c\x61\x09\xbf\xfc\x72\xab\x7f\xf2\xe4\x16\xb9\x78\x84\xd6\x1f\xba\x47\x13\x8f\x6e\xee\x2b\x8d\xa9\x4c\xa5\xc2\xa3\x8b\x5b\x5f\xcf\xb0\x32\x4e\x79\x63\x27\x61\x1a\x3e\xc0\xb9\xb7\xe4\xd7\x74\xca\xfa\x27\x4f\xae\x37\x3a\xc2\xc9\xba\xcd\x53\xab\xc6\xc2\xe3\x73\x9c\xdc\xd0\xf2\x73\x9c\x6c\x18\xee\xc0\xab\x93\xc7\x27\x5d\x90\x58\xa0\xc7\x80\xc0\xcc\x14\x85\xb9\x20\x19\x87\x69\x08\x91\xc8\x08\x92\x04\xdf\xb4\xa8\x29\xca\x0e\x84\x45\xb0\xb5\x86\x0b\xe5\x87\x20\x60\x5c\x82\x2a\x45\x8e\xcd\xf7\x48\xa5\xa3\x70\x0e\xb1\xc5\xca\xc0\x40\x8c\x50\x82\xd2\xcd\x2e\x21\x41\x9f\xd2\x16\xc3\xa4\x8b\x65\xb2\x26\x3e\xdb\xea\x1b\x3b\xc4\x82\x3f\xe0\x0e\xed\x18\x2d\xb7\x55\xe9\xde\x32\x3a\xa4\xde\x19\x4a\xf8\xa7\xf0\xd0\xd7\x1e\x6d\x65\x95\x43\x38\x56\xba\xbe\x84\x07\x70\x1e\x84\xe1\xf6\xd9\xe9\x0b\x77\x87\x0d\x84\xc3\xda\x16\xbd\xa1\xf7\x95\xeb\x26\x49\x2a\x75\x6c\x51\x0e\x85\x8f\x53\x53\x26\xa9\xd1\x1e\xb5\x4f\xa4\x72\x3e\x21\x6b\x49\x63\x2b\x79\x90\x3c\x68\x14\x25\x7b\xa4\x42\xd8\x74\x98\x18\xc7\xf2\x2a\x1f\xe1\xa4\x97\xa9\x02\xbb\x49\x12\xf6\x51\x8d\x54\x62\xab\x92\xe7\x55\x9e\x9c\x9d\xbe\xe0\x4f\x4f\x9f\xf2\xe7\xfd\xff\xf2\xc6\x0a\xb7\x58\xa0\x70\xc8\x9c\x2b\x52\x41\x00\xea\x85\x55\x76\xe8\xca\x24\x15\x49\x2b\x55\x63\x45\x87\x1c\xa4\xc2\xb1\x37\x92\x3b\x30\xb8\x10\x25\x87\x76\x00\x87\xa1\xa6\x7b\x28\x7b\x13\x74\x6c\x3d\xa4\x78\xe9\xad\x70\x37\x8b\x2c\x87\x7e\x58\xf5\x25\x42\xdc\xf8\xf3\x77\x8c\xb4\x71\xc8\xbf\x8f\x0f\x0e\xb6\xc5\xfa\xa4\x42\x7d\x4e\xac\x0d\x47\x46\x7b\xa1\x34\x5a\x38\x2d\x84\xcf\x8c\x2d\x81\x16\x7d\x19\x34\x63\x42\xba\xff\x1e\xc1\x16\x7f\xd4\x16\x53\x63\xb1\x0d\xed\xfc\x7b\x23\x66\x55\x9b\x3b\xe2\x52\xa5\xd6\x38\x93\x35\xb1\x9b\xd4\x65\x60\xa3\x64\xb1\x72\xd9\x42\x5e\xe5\xe9\x10\xd3\x51\x4f\x9b\xc0\x9a\x9f\x9c\x76\x0e\xcf\x4e\x88\x45\x43\xee\x82\xba\x92\xc4\xcc\xf0\xe6\xea\xaa\x65\x61\xf7\x2f\xa3\xf4\x47\x92\x5c\x74\x0f\x22\x98\x4e\xdf\x6e\x24\xaa\xcc\x58\x10\xde\x63\x59\x79\x50\x1a\xae\x0e\xe2\xf8\xc7\xe9\xcf\x20\x0d\x03\x98\xd4\x25\xb4\x6e\x00\x9f\x00\xff\x1d\x6e\x66\x33\x98\x84\x6f\xbf\x85\x81\x45\x31\x62\x00\x3b\x37\xfc\x66\xe6\xc6\xde\x55\xfb\xdb\xf4\xed\xf6\xad\xb7\x3e\x35\x99\x37\x13\xaa\x40\x19\x31\xa0\x4c\xff\xe6\xcd\xd2\x6a\xe0\x85\x87\x1f\xe1\xed\xdb\x9f\x29\x23\x69\x70\x05\x62\x05\x07\x3f\x03\x16\x0e\x01\x2f\x95\xa7\x8f\x4c\x31\x69\x34\x7e\xe4\x34\x2c\x96\x66\x7c\xb3\x12\x80\xa2\x97\x16\x28\x34\xe5\x3c\x66\x4b\xe0\x36\x83\x9d\x25\xc1\x2e\xa8\x5e\x5d\xa1\x96\xd3\x29\xdb\x9d\x6b\x43\x35\x07\xa9\x45\xb1\x94\x72\x95\x07\xe5\xe6\xf9\xd3\x1b\x30\xb5\x85\xff\xbc\x68\x72\xad\x63\x61\x8d\x90\x12\x78\x06\xd2\xa4\x23\xb4\x8c\xa9\x0c\xbe\x86\xdc\x62\xb5\xf0\xa9\x99\x6a\xf2\x6e\xe6\xbc\x18\x34\x51\x65\x00\x6e\xe2\x3c\x96\xa9\x2f\xc0\x79\x53\xb5\x3a\x78\xd8\x7a\x5d\xc5\x5e\x95\x68\x3f\x2a\x45\x24\xa3\x52\xbc\x4e\x6e\x69\xbe\x1c\x65\x2e\xbe\xcc\x1c\xb9\x9b\x48\x1c\x13\x53\x8d\x9a\x1b\x97\xcc\xab\xda\x4a\x58\x7f\xc0\x00\x30\x1d\x1a\xb8\xb5\x5b\x0c\x36\xf6\x08\xa4\x1e\x72\x5b\xfd\x5e\x1b\x2f\x00\xf6\x61\xff\x16\x3c\x7c\xb8\xd8\x3a\xb9\x61\x6a\xed\xd7\x57\x32\x00\x8b\xce\x1b\xba\xfc\x1a\xf8\xd9\x96\xf9\xa6\x46\x23\x4d\xcd\x50\x22\x05\x96\x46\xc7\xef\x9d\xd1\x8b\xfa\x8c\x01\x44\x54\x1b\x4b\xab\xc6\x68\xa3\x2e\x44\xef\x4d\x6d\xb5\x28\x64\x74\x8f\xe6\xa4\x72\xc4\x2b\xbc\xc0\x5c\xa4\x13\x6e\x31\x57\xce\xdb\x49\xd4\x05\x6f\x6b\x64\x4d\x55\xb6\x1a\x4b\x61\xfd\x66\x30\xb7\x0b\xac\x9d\x5d\xa6\x18\x6b\x23\x53\xd5\x45\x41\x24\x30\xbb\xf8\xcf\x02\x84\xe2\x97\x46\x62\xb8\xea\x0f\x43\xa8\x35\x49\x7d\xbb\x15\x45\xe8\x53\xb9\x0d\x43\xf3\x53\x5d\x3f\x2b\x97\x3a\x75\x90\x14\xb5\xde\x87\x0f\x1f\x9a\xdd\x5d\x77\xac\x4b\xa2\x6b\x06\x9b\x03\x95\x98\x89\xba\xf0\xee\x93\x0e\x94\xd6\x5d\x7f\x9c\x61\x96\xe2\x42\xf4\x29\x5d\xa0\x4e\x9f\x56\xf7\x7e\xfa\xe1\x87\x1f\x02\x79\x7e\x55\x59\xe3\x4d\x6f\xef\x4a\x3a\xff\xcd\x37\xf7\xee\x4e\xd9\x57\x95\xb1\xbe\x19\xe8\x74\xee\xde\x9b\xb2\xaf\x16\xdd\xcd\x61\xe8\xbe\x9e\x9d\xf5\x7f\x3d\x3c\x3e\x7e\x77\x78\x7c\x7c\xf2\x2b\xf0\x0a\xf6\x82\x12\xe0\x25\x9d\x8e\x47\xe0\xbc\xf9\xff\x65\xff\x57\x1a\x9c\x4d\x73\x49\xaa\x61\x2f\xfc\xe4\xef\xe1\xf0\xe8\xa8\x7f\xfa\x0a\xf8\x45\x4b\x6c\x33\x3b\xdc\x89\x31\xb6\xe0\x73\x13\xd7\xb0\x57\x32\x9b\x25\x66\xb9\x08\x34\x49\x48\x20\x32\xd1\x74\xaa\x17\x42\xe4\xa8\x7d\xe8\x3f\x35\xfa\x0b\x63\x47\x50\x7b\x55\x28\xaf\xd0\x41\x6e\x02\x1d\x7b\x03\x56\xa4\x48\x84\x28\x15\x31\x4f\x4c\xcd\x5b\x36\x5f\x6c\x6b\xed\x60\x80\x99\xb1\x08\x52\x3b\xa2\xa3\x91\x36\x17\x1a\xbc\x09\x04\xd6\x58\x42\x40\x2d\xa1\xae\x9a\xf2\x9f\x28\x7c\x02\x2e\x64\x1b\x76\x31\x54\x05\x06\x76\x9f\x33\x2c\x70\x79\x07\x7a\x3d\x88\xa2\xc0\xf0\xd2\x2c\xf8\xfd\x13\xf8\x9c\x80\x4c\x7b\xdc\xc4\xf2\x79\x23\x05\xd3\xe9\xee\xe4\xb8\xfb\x46\x2c\xb4\x7c\x5e\xfe\xfb\x64\x2b\x37\x4d\x83\xff\xd8\xbf\x2e\x0f\x3a\xf4\xf0\xdd\x25\xc3\xcb\x80\xa7\xf3\xc3\xf3\xd7\x67\xcf\x7a\xb7\x96\x8c\xbf\x08\xbe\xb7\xb6\x9b\x79\x98\x4e\x6f\x85\x85\xfc\x72\x46\x15\xd4\xd7\x71\x5e\x59\x35\x56\x05\xe6\x28\x81\x73\xca\x81\x7c\x06\x22\x3a\x47\xe0\x63\x48\xba\x09\xfd\xda\xfd\x03\x38\xb6\xd6\x76\xef\xb4\x8d\x19\xab\x35\x19\x6c\x56\x30\xd6\xd4\x06\x3c\x15\xdc\xdb\xda\x79\x42\xf3\xa9\xd2\x30\xaa\x07\xd8\x00\xdd\x11\xd8\x6a\x87\x50\x98\x54\x14\x20\x2a\xd5\x96\xd5\xcc\x91\x73\x0a\xb8\x45\x88\x5c\xe7\x36\xdc\x6d\xc6\xbb\x70\x27\xbe\xdb\xf9\xed\x60\x56\xf4\x2d\x25\xf7\x4e\xd4\x70\x98\xb1\x2a\x57\x3a\x69\x0e\x33\x99\x3f\x9c\xf0\x66\x20\x5e\x18\xff\x7c\x1b\x74\x45\xc2\x8f\xbf\x5e\xab\x93\xfa\xf3\x95\x5a\x63\x7c\x12\xd4\x24\xad\x1e\x76\x75\xc5\x09\x8c\x1a\x61\x2f\x7e\x24\xd2\x51\x5d\x3d\x2a\xcc\xe0\x25\xdd\xdd\x28\xfa\xe8\xb3\xcb\x9c\x86\x88\x7d\xc7\x68\x27\x1b\x05\x16\xeb\x80\xf3\x74\xe3\x21\x47\x1f\x78\x64\x10\xac\x84\x5a\xeb\x2c\x5b\x65\xeb\xe4\x2e\xa3\xeb\x40\x7e\x3c\x56\xb6\xb7\x3a\xd7\xae\x2b\x47\x52\x59\xd8\x5b\x92\xfb\xd3\x2c\xd2\xf7\xa9\x6c\xf6\xfc\x99\x44\xb2\xa2\xe8\x4b\x72\xc9\xaa\xa1\xbf\x8e\x4e\x76\xfa\x29\xcd\x85\x2e\x8c\x90\x14\xc4\xe6\x10\xa2\x55\xfe\xd8\xa4\x8c\xdf\x18\x04\xda\xd8\xb8\x7f\xdd\xcd\xa1\x6d\xc2\xe1\xf9\xb2\xb2\x66\xac\x24\xda\xa4\x9b\xbc\x93\xc2\x8b\xe4\x9d\xa9\xe7\xaa\x97\x01\xd0\x4d\x4c\x4d\xdc\x44\x53\x1f\x8b\x19\x41\xa9\xd9\x44\xa3\x89\x0f\x5a\xb8\xf7\x68\xe5\xda\x0d\x98\x4e\x5b\x21\x19\x5e\x79\x43\xbd\xde\x23\x63\x2d\x18\x63\x39\x68\x05\x44\x78\x3e\xeb\xcd\x42\xb5\x3b\xa0\xad\xfd\x99\x30\x1d\xe1\xec\x9a\xdc\x9f\x15\x33\x7f\x16\xd3\x0d\xef\xd3\x9e\x3f\x13\xd3\x2b\x8a\xbe\x24\xa6\x57\x0d\xfd\x9f\x30\xdd\x44\x39\xbc\x7c\x6a\x51\xb9\xa1\xf1\x37\xc2\x34\xa1\xa8\x3b\xff\x6d\x3e\xb5\xcc\x57\xdd\xd5\xaf\x06\x9d\x1c\xa1\xff\xea\xe8\xf1\xd1\xab\xe3\x77\x87\xa7\xcf\x7a\xd1\xf7\xd1\x35\xa0\x5d\x0d\x0a\xc9\x90\x96\xd0\x03\xb4\xfe\xce\x80\xb2\x72\x13\x36\x70\x49\xf7\x86\x13\x61\xae\x72\xa9\xc6\x8b\x56\x20\x54\x69\x4b\x8c\xdd\x0e\x2b\xad\xbc\x12\x05\x6f\x5f\x81\x21\x6a\x63\xb8\x1f\xfe\xcd\xdf\x58\x56\x46\xbb\xf7\xbf\xff\x69\xff\xde\xf2\xd0\xc1\x56\xc1\x83\x4d\xc1\xfb\x5b\x05\xef\x07\xc1\x68\xbb\x4b\xdc\x9b\x11\xea\x10\x16\x9e\x19\xcb\x43\x7b\xb1\x26\x2a\xe4\x18\xad\x57\x0e\x79\x85\x68\x79\x6d\x0b\x07\x5b\x52\x63\x30\xc3\x58\x39\xde\x8c\x52\x72\x77\x6d\x6c\xe3\x8d\x60\x1e\xcf\x95\x94\xb4\xd2\x92\xac\xe9\xfd\x14\x64\x62\xa8\x8f\xa3\x90\x9e\xa9\xe0\x9e\x4e\x19\xf3\xb5\x46\xc9\x85\x2c\xa1\xb2\x26\xa3\x5a\x7b\x51\xcc\xa4\x46\x7b\x6b\x0a\x5e\x15\x82\xb0\xdf\xa1\x42\x5d\x14\xce\x80\x46\x94\x0b\xb9\x38\xd4\x55\xf1\xd8\x14\x75\x89\x0e\x08\x18\xe1\x1d\x02\xe5\xac\xf2\xa7\x26\xac\xe9\xa8\x53\xaa\xf7\xa9\x29\x68\x53\x2e\x2f\x61\xff\xc1\x8f\xfb\xd4\xf3\xcc\xb7\xd3\xd2\xf5\x35\xfa\xc9\x8f\xa6\x09\x0e\x25\x82\x9b\xb8\xc2\xe4\xe0\x94\x4e\x43\x23\x51\x0a\x2d\x72\x04\xa4\xba\xc1\x0f\x49\xc4\x0f\xad\xa9\xf3\x21\xcc\xfa\x68\xb6\x68\x7d\xdb\x66\x7a\xa6\x65\xde\x1e\xaf\xbd\x44\xac\x4f\x13\xaf\xa2\x0f\x25\x4a\x5d\x41\x8e\x1a\xc7\x22\x30\x58\xb8\xfa\x5e\xa4\xa3\x25\x0d\xb5\x2e\x85\x1b\x41\x29\x9d\x9c\x29\x00\xf1\x87\xc3\x74\xed\xb3\x34\x5a\x6e\x71\xa0\x79\x46\xfc\xd3\xcb\x9b\xe6\xfe\x66\xab\x59\x07\xb4\xf1\xd8\x05\xe1\x4d\xa9\x52\xbe\x00\x44\xe8\xd6\x52\x2b\xdc\x10\x0a\x63\x2a\x07\xb5\xf6\xaa\x98\xfd\xcd\x51\x39\xa8\xab\x4d\xd7\xb7\x6a\x99\x1b\xfb\x2b\xfe\x4e\xe7\xd2\x21\xca\x3a\xe0\x61\x99\x74\x2c\x0e\x8c\xf1\x54\xfc\xa7\xa6\xac\xc2\x9b\xd9\xb6\x47\xc5\x88\xb9\x61\xed\x29\x6f\x12\x43\x37\x6b\xbe\xbb\xcf\xae\xae\x28\x03\x4c\xa7\x1b\xc1\xdc\xb9\x9f\xf9\x53\xc5\xfc\xd5\xee\x7f\x01\x00\x00\xff\xff\x56\x00\x30\x67\xe3\x1d\x00\x00")

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

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x5f\x6f\x1c\xb7\x11\x7f\xe7\xa7\x98\x9c\x0c\x2b\x46\xc2\x5d\xa9\x85\x2b\x40\x8e\x0c\xb8\xae\x92\xa6\xae\x2b\x41\xea\x1f\x14\x82\x1e\x78\xe4\xec\x2e\x73\x5c\x72\x3d\x9c\x3d\x69\x7d\xbe\xef\x5e\x90\x7b\x27\x5d\x75\xf2\x39\x4a\x90\x97\xbc\x1d\xc5\xf9\xfb\xe3\xcc\x6f\x66\xb5\xf7\x55\x39\xb5\xbe\x9c\xaa\xd8\x80\xc4\x5b\x21\x16\x0b\xb0\x15\x14\x6f\x83\xaf\x6c\x5d\x5c\xa2\xee\xc9\xf2\x70\xae\x58\x37\xe7\x4a\xcf\x54\x8d\x11\x96\x4b\xe1\x42\x5d\x23\x81\x64\xf0\xc1\xa0\x8c\xac\x88\xfb\xae\x88\x0d\x4c\xac\x8f\xac\x9c\xb3\xbe\x06\x42\x03\x8d\x62\xd0\xc6\x83\xce\x16\x7b\x52\x6c\x83\x87\xe0\xe1\xd9\xd7\x4d\x88\xec\x55\x8b\x2f\x26\x42\x2b\x86\xd7\xe5\x5c\x51\xe9\xec\xb4\x1c\xfa\xb6\xd4\xce\xa2\x67\xa9\x91\xb8\xe8\xb0\x85\xef\xbe\xdb\x3f\x3d\xfb\x7e\x3f\x05\xf8\x16\x89\xdf\xc4\x3f\x0f\x8c\xf1\x2e\xd2\xf4\x37\x5b\x59\xad\x18\x63\xb1\x8a\xf4\x02\xbb\x10\x2d\x07\x1a\xf2\x35\x7c\x82\x4b\xa6\x14\xd7\x72\x29\x4e\xcf\xbe\xff\xbc\xd3\x19\x0e\x0f\x7d\x9e\x93\x9d\x2b\xc6\x77\x38\x3c\xd1\xf3\x3b\x1c\xb6\x1c\xef\xc1\x3f\xcf\xfe\x72\x76\x0c\x06\x1d\x32\x02\x37\x08\x55\x70\x2e\xdc\x24\x99\x88\x3a\x43\xa4\x2a\x46\x02\xe5\x1c\x68\xd7\x47\x46\x8a\xa0\x08\x81\x7a\x0f\x37\x96\x1b\x50\x30\x6f\xc1\xb6\xaa\xc6\xf1\x3c\xb3\x7a\x96\xdf\xa1\x20\xec\x02\x4c\xd5\x0c\x0d\x58\x3f\x66\x09\x25\xb2\x4e\x29\xe6\xcb\x58\x98\xf2\x81\xf8\x3a\xd5\x2b\x6a\xd0\xc9\x23\x19\x91\xe6\x48\x92\xba\x36\x5e\x8b\xf4\x48\x27\x17\x68\xe0\xaf\x8a\xe1\xd4\x33\x52\x47\x36\x22\xfc\xdd\xfa\xfe\x16\x8e\xe0\x32\x0b\xc3\xd7\x17\xe7\xef\xe3\x0b\x31\x55\x11\x7b\x72\x27\x0d\x73\x17\x8f\xcb\x52\x1b\x5f\x10\x9a\x46\x71\xa1\x43\x5b\xea\xe0\x19\x3d\x97\xc6\x46\x2e\x93\xb7\x72\xf4\x55\x1e\x95\x47\xa3\xa1\xf2\x59\x32\xa1\x48\x37\x65\x88\xa2\xee\xea\x19\x0e\x27\x95\x75\x78\x5c\x96\x39\x8f\x6e\x66\x4b\xea\x5a\x59\x77\x75\x79\x71\xfe\x5e\xfe\x70\xfe\x83\x7c\x77\xfa\x5f\x39\x7a\x91\x84\x0e\x55\x44\x11\xa3\xd3\x2a\x15\xd0\x49\xd6\xa2\x26\xb6\xa5\x56\xe5\x4a\xaa\xc7\x2e\x3d\x72\x96\xca\xcf\x3e\x4a\xee\xa8\xc1\x7b\xd1\x14\xd0\x8e\xc2\x11\xe8\xd5\xd4\xa1\x39\x19\x30\x8a\x87\x90\xe2\x2d\x93\x8a\x4f\x43\x56\xc2\x69\xd6\xfa\x2d\x20\x1e\xe3\xf9\x3d\x22\x1d\x22\xca\x3f\x16\x87\x87\x8f\x61\x7d\xd6\xa1\xbf\x6c\x6c\xc5\xf0\x36\x78\x56\xd6\x23\xc1\xb9\x53\x5c\x05\x6a\x21\x29\xfd\x36\xd5\x8c\x65\xb2\xfd\xfb\x00\x5b\x7d\xec\x09\x75\x20\x5c\x41\x7b\x77\xde\xc2\xac\x5b\x4d\x8e\xa2\xb5\x9a\x42\x0c\xd5\x88\xdd\xd0\xb7\x99\x8d\xca\x7b\xcd\x4d\x0f\x75\x57\xeb\x06\xf5\xec\xc4\x87\xcc\x9a\x3f\x73\xe8\xbc\xb9\x38\x4b\x1c\x9a\xe7\x16\xf4\x9d\x49\xbc\x0c\x57\x8b\xc5\x8a\x83\xe3\xdf\x82\xf5\x5f\x18\x70\x93\x6f\x61\x02\xcb\xe5\xf5\xd6\x98\xaa\x02\x81\x62\xc6\xb6\x63\xb0\x1e\x16\x87\x45\xf1\x72\xf9\x0a\x4c\x10\x00\x43\xdf\xc2\x2a\x0c\x90\x03\xc8\x0f\xf0\x34\x9f\xd9\x25\x3c\x7f\x0e\x53\x42\x35\x13\x00\x3b\xd2\xbd\x5a\x07\xf1\x6c\xb1\xfa\xb5\xbc\x7e\x3c\xf1\x55\x44\xe3\xd4\xad\x94\x75\x68\x26\x02\xd2\x8c\xbf\xba\xda\xd0\x06\xe9\x18\x5e\xc2\xf5\xf5\xab\x34\x8d\x3c\x44\x87\xd8\xc1\xe1\x2b\x40\x17\x11\xf0\xd6\x72\x3a\x54\x56\x98\xe0\x71\xe7\x4b\x10\xb6\x61\xfe\xb4\xe1\x9f\x90\xd3\x0e\x95\x4f\xd3\x4e\x50\x0b\x92\x2a\xd8\xb9\x0c\xec\x2a\xd2\xc5\x02\xbd\x59\x2e\xc5\xee\x29\x5b\x53\xe8\x3b\xd0\x84\x6a\x63\xd8\x5a\x06\x1b\xef\x26\x27\x07\x08\x3d\xc1\xbf\xdf\x8f\x53\x36\x8a\xac\xa3\x8c\x01\x59\x81\x09\x7a\x86\x24\x84\xad\xe0\x2b\xa8\x09\xbb\xfb\x98\xc6\xab\x71\xe2\x56\x91\xd5\x74\xc4\x54\x00\xc4\x21\x32\xb6\x9a\x1d\x44\x0e\xdd\xca\x86\xcc\xa9\xf7\x5d\xc1\xb6\x45\xfa\xa2\x54\xa2\x17\xab\xf1\x73\x72\x1b\xf7\xed\xac\x8a\xc5\x6d\x15\x53\xb8\xa5\xc1\x79\xe2\xa8\xd9\xd8\x6b\x25\x61\x0c\x3d\x69\x94\x9d\x22\x3e\x14\x00\xa8\x9b\x00\xfb\xbb\xc5\x60\x2b\x47\x48\xe6\xa1\xa6\xee\x43\x1f\x58\x01\x1c\xc0\xc1\x3e\xbc\x7e\x7d\x9f\x7a\x0a\x23\xf4\x9e\x1f\x6a\x0a\x00\xc2\xc8\x21\xb5\xbd\x07\x79\xb1\x75\xbf\x58\xc8\x54\xa4\xf8\x01\x8a\x8b\xe0\x30\x75\x78\x45\x2a\xb5\x88\x00\x18\x17\xb7\xe4\x64\x94\x2e\x8d\xc2\x36\xf8\xe2\xa7\x18\xfc\xfd\xd2\x26\x00\x26\x2e\xd4\xd2\x90\x9d\x23\x4d\x8e\x61\xf2\x53\xe8\xc9\x2b\x67\x26\xdf\xa6\x3b\x63\x63\x22\x1b\xe9\xb0\x56\x7a\x90\x84\xb5\x8d\x4c\xc3\xe4\x18\x98\x7a\x14\xe3\xaa\x96\xe2\x40\x6f\x46\xbf\x9b\x88\x2b\xe2\x6d\xc8\x1f\x17\x78\xf0\xc2\x95\x15\x62\x85\x5f\xd7\x3b\x97\x68\x62\x4d\x0d\x3f\xe6\x42\x2b\xfe\x11\x0c\x66\x32\x78\x9d\x1f\xc4\x27\xa9\xe7\xa9\xa0\x6f\x72\x6f\x26\xd3\xa9\x86\x53\xef\xc1\x8d\x52\x35\x7a\x06\xe5\x0d\x78\xe4\x9b\x40\x33\xe8\xd9\x3a\xcb\x16\x23\xd4\x21\x73\x00\x07\x20\xa5\x31\xf5\xa1\xb1\xa9\xe0\x0b\xb1\x97\xe0\x5d\x2b\x53\xef\x23\x4c\xb1\x0a\x84\x60\x7c\x4c\x5d\x30\xf3\xe1\xc6\x03\x87\xdc\x37\xa3\x27\xcc\x48\xf4\xdd\xb8\x6f\x26\xde\x18\x20\x66\x82\x13\x37\x8d\x75\x98\x29\xe5\xae\xb1\x41\x9a\x17\x70\x72\x02\x93\x49\xa6\x15\x13\xee\x49\xe5\x8b\x24\x92\x70\x49\x19\x6e\x43\x73\x39\x4a\xc1\x72\xb9\x9b\x8d\x77\x03\x7c\x6f\xe5\xd7\x10\xee\xcf\xf6\xf1\x54\xde\xfd\xd3\xc1\xe7\x88\x37\x22\xc3\x37\xb7\x02\x6f\xbb\x40\x0c\x97\x6f\x2e\xff\x75\xf1\xe3\xc9\xfe\x86\xf3\xff\x04\x9a\x21\xad\x7c\x8f\xf7\xb0\x5c\xee\x67\x45\x79\xbb\xae\xbb\xf4\x11\x21\x65\x47\x76\x6e\x1d\xd6\x68\x40\xca\x44\xbb\x72\x5d\x40\xe9\x0d\x41\xce\xa1\x3c\x2e\xd3\xcf\xe3\x8f\x20\x71\xe5\x6d\x77\xa6\x2b\xcc\x44\xef\x93\xc3\x51\x43\x88\x71\x18\x49\xad\x24\x53\x1f\x59\xec\xea\x6e\xee\x3d\x1a\xa9\x4c\x0b\x1d\x85\xb4\x1d\x41\xe8\xd0\xc7\xb4\xaa\xc9\xb4\x6c\x51\x70\xb2\x73\xca\xe3\xd8\x9a\x09\xa5\x2f\x68\xa5\xc7\xdc\xec\xe3\xd4\x47\x08\xca\xc5\x00\x1e\xd1\xdc\x4b\x16\x2e\x68\xe5\x8a\x79\x70\x7d\x8b\x11\x8c\xa5\x71\x40\xa0\x59\xf7\x46\xa2\xbb\x91\xea\x74\xea\x88\xd4\x36\xa2\x9d\x25\x41\xd9\xc2\xc1\xd1\xcb\x03\x90\x1b\x83\x20\x90\xad\xad\x2f\x3f\x63\x3f\xc5\x31\x52\x50\x9e\x96\x71\x88\x2e\xd4\x10\xad\xd7\xb9\xd5\x5a\xe5\xd3\xa7\x1d\xce\x91\x06\x6e\x92\x08\x37\x14\xfa\xba\x81\x35\x8b\x89\x7b\xb6\x59\x51\xd9\xda\xca\x1d\x23\x3d\x18\x11\x0f\xaf\xc5\x1e\x44\xe4\xcc\x25\x7d\x07\x35\x7a\x9c\xab\xdc\x05\xf9\x3b\x94\x95\x9e\x6d\x58\xe8\x7d\xab\xe2\x0c\x5a\x13\xcd\xda\x00\xa8\x8f\x11\xf5\x83\x63\x1b\xbc\x79\x24\x80\x71\xb3\xfb\xc5\xea\x23\x9f\x3e\x4d\x5b\xec\x81\x0f\x8c\xc7\xa0\x38\xb4\x56\xcb\xff\x2f\x09\xd0\xa4\x62\x03\x2e\x84\x2e\x42\xef\xd9\x3a\x68\x55\xcc\x6b\x40\x84\xbe\xdb\x0e\xfd\x51\x2b\x77\xce\x7e\xfd\x3f\x4e\xa2\x6e\xd0\xf4\xb9\x1a\x36\xd6\x23\x20\x9c\x86\xc0\x89\x89\x75\x68\xbb\xbc\xca\x3c\xb6\xe9\x4d\x44\x6c\x7a\x36\x89\xb4\xa5\x5c\xe9\x7c\xf3\x87\xb4\x0d\xb9\x88\xcb\xe5\x16\x94\x3b\xb3\x81\x4f\x9f\xc6\x21\x78\xb7\x4c\xfd\x2f\x00\x00\xff\xff\xe7\x16\x12\xf9\x29\x12\x00\x00")

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
