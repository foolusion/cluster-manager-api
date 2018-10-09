// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3b\x5d\x73\xdb\x38\x92\xef\xfe\x15\x5d\x7e\x39\xf9\x2a\x96\x62\x27\x99\x99\xb3\x37\x37\xab\x95\x33\x89\x2a\x89\xed\xb2\x9c\xa4\xe6\x49\x05\x91\x2d\x0a\x6b\x12\xe0\x00\xa0\x14\xcd\x56\xfe\xfb\x15\xbe\x48\x80\xa4\xe4\xc4\x71\x6e\xea\xb6\xce\x55\x33\x11\x89\x46\xa3\xd1\xdd\xe8\x2f\x34\x47\x23\x98\xf0\x72\x2b\x68\xb6\x52\x70\xfa\xf4\xe4\x17\x98\x91\x42\x56\x2c\x83\xd9\xc5\x0c\x26\x39\xaf\x52\xb8\x24\x8a\xae\x11\x26\xbc\x28\x2b\x45\x59\x06\xb7\x48\x0a\x20\x95\x5a\x71\x21\x87\x07\xa3\xd1\xc1\x68\x04\xef\x68\x82\x4c\x62\x0a\x15\x4b\x51\x80\x5a\x21\x8c\x4b\x92\xac\xd0\x8f\x3c\x81\x8f\x28\x24\xe5\x0c\x4e\x87\x4f\x61\xa0\x01\x0e\xdd\xd0\xe1\xd1\xb9\x46\xb1\xe5\x15\x14\x64\x0b\x8c\x2b\xa8\x24\x82\x5a\x51\x09\x4b\x9a\x23\xe0\xe7\x04\x4b\x05\x94\x41\xc2\x8b\x32\xa7\x84\x25\x08\x1b\xaa\x56\x66\x1d\x87\x45\x53\x02\xbf\x3b\x1c\x7c\xa1\x08\x65\x40\x20\xe1\xe5\x16\xf8\x32\x04\x04\xa2\x1c\xd1\xfa\x6f\xa5\x54\x79\x36\x1a\x6d\x36\x9b\x21\x31\x04\x0f\xb9\xc8\x46\xb9\x05\x95\xa3\x77\xd3\xc9\xab\xcb\xd9\xab\xe3\xd3\xe1\x53\x37\xe9\x03\xcb\x51\x4a\x10\xf8\x47\x45\x05\xa6\xb0\xd8\x02\x29\xcb\x9c\x26\x64\x91\x23\xe4\x64\x03\x5c\x00\xc9\x04\x62\x0a\x8a\x6b\xa2\x37\x82\x6a\xbe\x3d\x01\xc9\x97\x6a\x43\x04\x6a\x34\x29\x95\x4a\xd0\x45\xa5\x22\x9e\x79\x12\xa9\x8c\x00\x38\x03\xc2\xe0\x70\x3c\x83\xe9\xec\x10\xfe\x31\x9e\x4d\x67\x4f\x34\x92\x4f\xd3\xdb\x37\x57\x1f\x6e\xe1\xd3\xf8\xe6\x66\x7c\x79\x3b\x7d\x35\x83\xab\x1b\x98\x5c\x5d\x5e\x4c\x6f\xa7\x57\x97\x33\xb8\xfa\x0d\xc6\x97\xbf\xc3\xdb\xe9\xe5\xc5\x13\x40\xaa\x56\x28\x00\x3f\x97\x42\xef\x80\x0b\xa0\x9a\x9b\x98\x1a\xd6\xcd\x10\x23\x12\x96\xdc\x92\x24\x4b\x4c\xe8\x92\x26\x90\x13\x96\x55\x24\x43\xc8\xf8\x1a\x05\xd3\x9a\x50\xa2\x28\xa8\xd4\x52\x95\x40\x58\xaa\xd1\xe4\xb4\xa0\x8a\x28\xf3\xaa\xb3\xaf\xe1\x81\x06\x79\x4f\x93\x15\xc1\x1c\x3e\x22\xc3\x3f\x29\x81\xbf\x15\x6b\xfb\xeb\xef\x59\x41\x68\x3e\x4c\x78\xf1\xdf\x1a\x6e\x9c\xd3\x3b\x02\xef\x88\x90\xc8\xe0\x6f\x44\x3f\x0d\x73\xf3\x14\x02\x1e\xc8\x2d\x53\xe4\x33\xbc\x84\xc3\x52\x70\xc5\x9f\x1d\x9e\x1f\x1c\x94\x24\xb9\xd3\xa4\x26\x79\x25\x15\x8a\x79\x41\x18\xc9\x50\xcc\x49\x49\xcf\x0f\x0e\x68\x51\x72\xa1\xe0\x30\xe3\x3c\xcb\x71\x44\x4a\x3a\x22\x8c\x71\x47\xf6\xd0\xa0\x39\x3c\xaf\xc1\xcc\x73\x72\x9c\x21\x3b\x96\x1b\x92\x65\x28\x46\xbc\x34\xa0\xbd\xd3\x0e\x0e\xec\x28\x0c\x32\x51\x26\xc3\x8c\x28\xdc\x90\xad\x1d\x4e\xe6\x19\xb2\xb9\xc3\x32\x74\x58\x86\xbc\x44\x46\x4a\xba\x3e\xf5\x23\x47\xf0\x12\xfe\x75\x00\x40\xd9\x92\x9f\x99\x5f\x00\x8a\xaa\x1c\xcf\xe0\x70\x62\xb7\x04\xef\xed\x96\x60\x7c\x3d\x3d\x3c\x37\x10\x6b\x7b\xc0\xce\xe0\x70\xfd\x74\x78\x3a\x7c\xea\x5e\x27\x9c\x29\x92\x28\x8f\x47\xff\x31\x52\x68\x54\xfe\xac\x4f\x2e\x27\xb7\x0e\x58\xff\x55\x22\x3f\x83\x43\x7d\x30\xe4\xd9\x68\x94\x51\xb5\xaa\x16\x9a\xd7\x23\x69\xe1\x8f\x13\x96\xa8\x91\x63\xed\xb1\x63\xed\x31\x29\x69\x80\x03\xb5\x80\xce\xe0\x90\xa4\x05\x65\x7f\x0f\x27\x0e\x29\x77\x70\x5f\xf4\x3f\xe6\x7f\xf8\x59\xa1\x60\x24\x9f\xa7\x3c\x91\x9e\xd0\xef\x25\x23\x45\x99\x08\x6a\x58\x7c\x06\x87\xef\xb9\x40\x20\x0b\x5e\x29\xd8\xc5\xc1\x2f\x07\x00\x32\x59\x61\x81\xf2\x0c\xde\xdc\xde\x5e\xcf\xce\xdb\x6f\xf4\x8b\x84\x33\x59\x99\x37\x87\xee\xe0\xeb\x25\x46\xff\x94\x9c\x19\x34\xa5\xe0\x69\x95\xec\x1a\xff\x72\x7e\x70\x20\x51\xac\x69\x82\x35\x21\x76\xbf\xfa\x3c\xd3\x3c\xd7\xf3\xd7\xd4\x58\x4a\xe2\xf5\xd7\x8c\x8b\x32\x81\x89\x40\xa2\xd0\xcf\x1b\x44\x8f\xef\x65\x76\x04\x02\x55\x25\x98\x6c\x0d\xdd\x60\x99\x6f\x8f\x02\x05\xa8\x35\xd4\x9c\x80\x21\x29\xe9\x50\x33\xda\xeb\x5d\xf3\x57\x72\xa9\xe0\x0c\x0e\xcd\x21\x59\x9f\x78\x76\x1f\x46\x40\x0b\x9e\x6e\x35\xd0\x7f\x36\xaf\xbf\x38\x11\x47\x3b\x13\xa8\x04\xc5\xb5\xb5\x33\x52\x11\x55\x49\x6d\x9b\xeb\x6d\x6a\x1b\x02\x54\x49\xb8\xab\x16\x98\x70\xb6\xa4\x99\x31\x43\x09\x67\x0c\x13\x45\xd7\x54\x6d\x6b\x56\xbc\x46\x55\xf3\xa1\xf9\x1d\x33\xa1\x79\xff\x70\x0e\x64\xb8\x9f\x01\xbd\x3b\x4d\x31\x47\x85\x3d\x02\xbc\x30\x03\x35\xe1\xd1\x63\x4c\x7b\x34\xf4\x70\xf2\x1d\x25\xdf\xbc\x03\x92\xfe\xb3\x92\x0a\xc8\x5e\x75\x1c\x1b\x20\x47\xe3\x25\x4f\x51\xc2\x20\x7a\x17\x6f\x29\x1a\xfa\x0e\x9d\xac\x7e\x80\x4a\x12\xc8\xa9\x54\x5a\x1d\x1d\x3e\xd9\xa3\x69\xef\x34\xc8\x20\x7e\xde\xa5\x71\x7a\xec\xb1\xb5\x6e\xa4\x69\xbc\x67\x47\x94\x49\x45\xf2\x1c\x06\x5c\x80\x40\xf7\x74\x04\x8a\xe6\x79\x20\xb9\x6b\x2f\xd5\x5b\xf3\x1e\x06\xad\x17\xf1\xae\x5a\x83\x8f\x67\x4e\x2c\x55\x0f\x13\xdd\x8e\x8d\xae\x30\x2f\x20\x59\x11\xa1\x3c\xf4\xad\x0e\x24\x37\x7a\xca\x02\xb5\x57\x50\xa2\x4a\x4c\x48\x4b\x8d\xed\xd1\xa0\xb0\x22\x12\x48\x2e\x90\xa4\x5b\x58\x20\x32\x48\xb1\xcc\xf9\x16\xd3\x26\xce\x94\xa4\x40\xe3\x39\x6b\x26\x4e\xed\x9a\x6f\x30\x2f\x26\x06\xcb\xa0\xfd\x26\x66\x63\x7b\xf4\xf1\xf8\xa8\x37\xfd\x30\x2e\x3a\xfb\x50\x6f\xb7\xc5\xbe\xc6\x66\x05\xbb\x6c\xbd\xe8\xb3\x5b\x8f\xb0\xc7\xae\xe5\x8a\x77\xb9\xeb\x3c\x57\x82\xf9\x50\xc8\x04\x50\xa2\x30\x0e\xd8\x39\x7f\x52\x52\xd0\xfe\x37\x38\x0b\xaf\x51\xb9\xdc\x64\x1a\x80\x0f\x9a\xd7\x9d\x23\xee\xde\x3f\xda\xf1\x76\xe4\x7e\xdd\xde\xaa\x32\x13\x24\x45\xb7\x98\x34\x1e\x92\x40\x46\xd7\xc8\x3a\x06\xfa\x35\xaa\x0f\x16\xdc\x59\xa5\xf6\x0e\x77\x8e\x76\xf6\xbc\x13\xf2\xd1\x8d\x9c\xdb\xe0\x7d\x0e\x4a\x29\x2c\x4a\xa5\x73\x2c\xcf\x91\xae\x83\x8a\x89\x86\x41\xfc\x1c\xef\x31\x1e\x7b\x6c\xf7\xd4\xdd\xd5\xb7\x9c\xd2\xaa\x4c\x89\x8e\x27\x36\x12\x12\x81\x29\x32\x45\x49\x2e\x75\x9a\x9c\x3a\xf9\x77\x77\xae\x67\x8c\x3f\xcd\x26\x01\xfc\xa0\xef\x6d\x9b\x0b\x5d\x88\xc7\xe3\x45\x83\x73\x44\x36\xf2\xfb\x78\xf1\x67\x25\xf0\x1b\xb9\xa1\xa7\xf4\xf1\xa3\xf5\xbe\x97\x23\x2d\x98\x1f\xc3\x13\xbd\xc8\xd7\x73\xe5\xcb\xc1\x01\xb2\xaa\xb0\xce\x3c\xad\x13\x09\x9d\x76\x2f\x29\xc3\x14\x5e\xc2\x53\x0b\xab\x15\xe7\x25\x9c\xb8\x07\xc3\xb9\x97\x70\xea\xd2\xc6\x62\x43\xcc\xf3\xb3\xf3\x1a\xa3\x3b\x04\x33\x1b\xa3\xd7\xf9\xc9\x25\x57\x20\xd1\x7a\x86\xd9\xed\xf8\xf6\xc3\x6c\xfe\xe1\x72\x76\xfd\x6a\x32\xfd\x6d\xfa\xea\xa2\x59\xce\x38\x5d\x84\xeb\x9b\xab\x8f\xd3\xd9\xf4\xea\x72\x7a\xf9\xda\xc4\xfb\x08\x94\xa5\x3a\x21\x42\x69\x5c\xaa\x8f\xfb\xa9\x84\x05\x52\x96\x69\x71\x12\x85\xe9\xd0\x60\x89\xa6\xd7\xd4\x3b\xdc\x37\x1f\x2e\xef\x45\xab\x9d\xba\x71\xe6\x0e\xad\xcd\x2f\x24\x2c\xab\x3c\xdf\x42\x25\xc9\x22\x47\xbb\x94\xc7\x56\x33\xc5\xaf\xf2\x6a\x72\x75\x39\x99\xbe\xeb\x5f\x89\x28\x90\xbc\x40\xd8\x70\x71\xa7\xf1\x12\x13\x56\xe4\x5b\xb7\x99\x94\x33\x04\xce\x42\x92\x9e\x80\xac\x92\x15\x10\xe9\x6c\x96\x06\xd3\xc3\x05\x31\x04\x73\x01\x8c\xa7\x58\x17\x8a\x1c\x71\x01\x11\x46\x4a\x01\x81\xb3\xdb\xab\xeb\xeb\xaf\x66\xaf\x75\xa9\xa9\x93\x9f\x9b\xf9\x12\x9e\x47\x28\x5f\xdd\xdc\x5c\xdd\xec\xc5\x57\x10\xbd\x45\xa8\x98\x65\xa1\x99\x6c\x67\xbd\x84\x17\x11\xae\x8b\x57\xaf\x6f\xc6\x17\xaf\x2e\xf6\xa2\x73\xa5\x34\x73\x80\x85\x61\xa2\x66\x1a\x07\x81\x52\xe9\xac\x5d\x8b\x0b\x96\x15\x33\x03\x24\xf7\x49\x60\x8d\xfb\x25\xfc\x64\x34\xb7\x40\x29\x49\x86\x30\x21\x79\xbe\x20\xc9\x5d\xa3\xb7\x9a\x94\x0f\x37\xef\x34\xd2\x44\x87\x8c\x66\x54\x71\x33\x2c\x95\xd0\xac\xa9\x44\xde\xd1\xb1\xe9\x85\x2f\x1d\x6a\x12\x51\xaa\x70\x82\x7b\x35\xa7\xa9\xd5\x9a\x90\x80\x56\x5a\x1e\x1c\x20\x1d\x43\x3a\x9c\x7e\xfb\x8a\x6b\x66\xd6\x69\x96\x13\x8f\x5b\x45\xc7\x9b\x1d\xba\x4a\x7f\xde\x7d\x6d\xce\xf8\x61\x03\x11\x2d\xed\xed\xc2\xac\xc4\xa4\x99\x14\xea\x78\xcd\xaa\x20\x50\xb2\x78\xfc\x40\xe2\x7f\x78\xfb\xd0\xbb\x49\x63\x0d\x83\x3a\xc6\x0a\x4d\xa9\xd1\xe8\xb3\x8a\x76\xbb\x21\x32\xdc\x2b\x2c\xb6\xb6\xd0\x1b\x72\x78\xc1\x79\x0e\xfc\xae\xb3\xef\x14\x15\xa1\xb9\x6c\x33\xd0\x4d\xd5\xfa\x52\x72\x26\xad\x3e\xfa\x38\x45\x61\x51\x03\xb6\xe5\x14\x95\x0d\xbe\x46\x48\x39\xe7\x77\x98\x42\x55\xee\x15\xd1\xf8\xd3\x0c\x02\x4f\x61\x5e\xb7\x1c\xb1\xb5\xc7\x8d\x1c\x8c\x77\xe9\x4e\x6a\xfb\x2b\x6f\xb9\x1b\x1b\x10\xd2\xea\x25\x2c\x61\x40\x36\xce\x99\x8c\xac\x75\x1f\xa1\x4a\x8e\xac\x51\xf5\x6a\x10\xe8\xc3\xf3\x1d\x5c\x69\x49\x75\xda\xb2\x2a\xd6\xb2\xc9\xad\x54\x58\x74\xe5\x16\x4a\xe1\xc2\x08\x6e\xaf\x2c\xda\x95\x90\x50\x99\x88\xd2\x26\x2c\x58\xfb\x3f\xa4\xe5\xba\xe2\x36\x9f\xe3\xdb\x7f\x3b\x81\x3c\xe4\x80\xbe\xd8\xcd\xd1\x96\x28\x27\xbc\xca\xd3\x48\x9a\x3e\x35\xd6\x49\xe0\xce\x43\x38\xab\x6b\x76\x3b\x8c\xa2\x2b\xea\xed\x3e\x68\xae\x62\xd2\x50\xf2\xd5\xb2\x39\x79\xa8\x6c\x4e\xff\x77\x0e\x4b\x5d\xef\x79\xe8\x81\x71\x93\xde\xf5\x96\xa1\xb0\xb4\x11\x4c\x8f\x59\xeb\xb2\x3b\x04\x6a\x88\xb9\x68\xd9\xb4\x50\x6e\xc6\x89\x9d\xf4\x32\xaa\x07\xda\x1d\xb0\xe0\xb4\xa4\x29\xb5\xce\xb9\x27\xe3\x8e\x8b\xbd\x3b\x50\x5a\x80\xb9\xdf\x41\x3b\xcc\xd9\x3d\x3f\x8e\x54\x6b\xfd\x7b\xde\xc7\x90\xc0\x0c\xfd\xdf\x67\x4b\x68\x16\x83\x7a\xb9\x0e\x71\x6c\xb9\x5c\xff\xdc\x81\x36\x80\x6f\x47\x7f\xdf\xcc\xe9\x17\xbb\xa3\x82\x28\xfe\xe8\xb5\xe7\xf5\xb9\x3a\x86\xa4\x12\x02\x99\xca\x5d\x3c\x40\xed\xb9\xe7\x3a\xda\x24\xf2\xde\x90\xc8\x17\x7c\xf8\x12\xde\x56\x0b\x14\x0c\x15\x46\xb3\xee\x7e\x91\x73\x0f\x64\x64\x64\x06\x39\x43\xbe\xac\xa9\x98\x87\x52\x6a\x92\x36\xb7\x84\xb6\x54\xdd\x80\xab\x13\x74\x8d\x3f\xcd\xcc\x7e\xad\xd5\x7a\x76\xbe\x03\xea\xad\x83\x72\x46\xea\xf9\x0e\xb8\x8f\xef\x3f\x11\x81\x06\xb4\xce\xd3\x7e\xf6\xf9\xdf\x57\x84\x5b\x54\xc2\x9b\x71\x63\x6b\x56\x34\x5b\xcd\xc9\x9a\xd0\x9c\x2c\xa8\x8e\xa4\x3b\x11\xfb\x92\x2c\x04\x4d\x5c\xbc\xa3\x33\xea\x88\xf5\xa8\x74\xae\x33\x77\x40\x26\xf0\xde\x29\x7d\xcf\x88\x28\x0a\xd7\x4c\xbc\x20\x8a\xc0\x04\x99\x57\xac\xf1\xa7\x99\x7e\x65\xdf\x40\x4a\x14\x99\x27\xf6\x77\x28\xe5\xd0\xb4\x6b\xea\x2a\x1a\x7b\xb0\x3e\xcf\x11\x56\x07\xc2\x83\xf9\x8f\xdf\xaf\x80\x2a\x2c\x6a\x77\x73\x2d\xdc\x79\xa8\x04\xa6\xda\x40\xe8\x80\x54\xf2\x4a\x24\x28\xa3\x53\x67\x4a\xb8\x2c\x41\xc8\x04\xaf\xca\x96\x71\x1e\x7f\x9a\xf9\xf1\xd7\x7a\xd8\x96\xa8\x59\x82\x73\x0b\x6d\xe5\xdc\x48\x8d\x26\xab\x0e\x33\x3c\x2b\x63\xa6\x44\xca\x68\x27\x0a\xcc\x4c\xbd\xa1\x92\xc7\x48\xa4\x3a\x3e\x79\x02\xa8\x92\xe1\x51\x0d\x59\xe7\x29\x99\xd5\xf9\x93\xf3\x2e\x92\x48\x15\xfe\xe4\x0c\x65\x80\x70\xf1\x04\xfc\xef\xd3\xc4\xfc\xde\xa0\xfe\x9d\xb6\x57\xaa\x19\xe0\x96\x0c\xb1\xce\x2d\xd6\x9a\xfb\x8d\xd6\xfe\xc6\x05\x6c\x56\xc8\x6c\x0e\xad\x56\x94\x65\x12\xb4\x86\xfb\x3a\xbc\xcb\xda\xdb\x6c\xe9\x91\x55\xe7\xac\x7e\xbc\x9e\x00\x4d\x9f\xc0\x22\x27\xec\xce\xd4\x84\xf4\x7f\x87\x16\xa3\x3e\xf5\xe6\x79\xcb\xab\xc3\x27\xb0\xa4\x79\x8e\x29\xd0\xa5\xe9\x40\xd1\x04\x68\xf5\xf8\x78\x3d\x69\x73\x72\x5d\x26\xf3\xd0\x23\xf8\xa0\x08\x93\x4a\x68\xfe\x19\x21\xb7\x27\x49\x37\x6a\x55\xa0\xce\x16\x5b\xf4\x4e\xc7\xef\x41\xf0\xbc\xe9\xba\xf0\xe7\x77\x40\x04\x3b\xf2\x07\x4a\x4a\x9e\x50\x63\x7a\xd2\xb4\xbd\x0e\x25\xc5\x5c\x63\x98\x13\xc1\x1a\x8d\x6d\xb8\xdd\x52\x5c\xb0\x25\x22\x20\xa0\xb6\x25\x9a\xc2\x08\xab\x8a\x85\x36\x22\xcb\x5a\x6d\x65\x9b\xf5\xb1\x76\x47\x5c\xaf\xf1\x1b\x7c\x83\xe2\xc5\x30\x27\x22\xc3\x1d\x4a\x69\x80\xda\x8c\x7c\x4f\x19\x2d\xaa\xa2\x8f\x10\x18\xa4\xb8\x24\x55\xae\xcc\xd9\xff\x13\x05\x6f\x50\x52\xa6\x9e\x9d\x42\x41\xd9\xfc\x8f\x8a\x30\x65\x8d\x5a\xcc\xe2\xf7\xe4\xf3\x77\x60\x26\x9f\x43\xcc\xcf\x82\xea\xdb\x68\xa4\x03\xc0\xd0\xf9\x8c\xaf\xa7\x30\xb3\xd7\x0b\x41\x88\xd8\xdc\x23\xc0\xbf\xdc\x3c\x1b\x2a\x6a\x71\xfb\xd9\x3e\x9c\xee\xce\x6b\x87\x95\x4b\xe0\x25\x0a\xeb\xa6\x74\x3e\x7d\xf5\x76\x47\xfa\xe5\x51\xf5\x5c\x6f\x74\x8e\x8c\x22\x99\x2f\x56\x65\x54\x27\xd3\x25\x97\x54\x71\xb1\x6d\xcb\x2e\xa3\x2a\xf0\xa4\x27\x1d\x5d\x5e\x11\xb9\xf2\xf1\x83\xc6\x94\xf0\xa2\xa0\xaa\x0f\x8b\x1d\xe9\x48\xab\xc7\x9d\x29\x81\x68\xb6\x9a\xe4\x48\x98\x35\x1b\xda\x03\xf4\xa2\xd5\xc0\x73\x5b\x70\x8a\x1c\xf0\x68\xa4\xad\xad\x89\xe0\x8c\xf7\x68\xcf\x35\x2f\xe7\xa9\x9d\xf7\x3c\x9a\xf7\xb1\x91\x70\xc6\x6d\xa1\xd9\x04\x5a\x45\x49\x5d\xfd\x2b\xa4\x81\x07\xfc\x79\x11\xe1\x99\xd8\x19\xa2\xf1\xac\xc1\xbc\xc4\x0f\x1a\xbf\x1a\xcc\xba\xce\x89\xd2\x92\x03\xaa\x2c\x13\x2c\xa0\x2d\x76\x8f\x40\x54\xcc\x74\x63\x05\x11\x89\xc3\x58\xfa\x89\x3d\x41\x83\xdf\xd2\xb4\x95\x58\xf6\xe8\x8a\xdb\x4d\x14\x1e\xf9\x9c\xc3\x09\x3d\x89\xbd\x73\xe5\x1a\xc8\x8c\xc1\xf5\x45\x4e\xef\xa8\x03\x83\x12\x3a\xeb\x28\x4e\x98\x61\x22\x50\xbd\xc5\xed\xd4\xee\x52\x9f\xab\x71\x92\xa0\x8c\x82\x3a\x69\xa0\xe6\x77\xb8\x9d\xb7\x42\xf5\x06\x87\x9d\xf5\x16\xb7\x35\x1e\xb2\x0b\x8f\x1d\xd0\xe8\xba\x85\x60\xeb\x45\x77\xa3\xa8\xdd\xac\xad\x92\x7d\x17\x5f\xda\x99\x6c\x1c\x41\x95\xe5\x5e\x9e\x90\xb2\xec\x63\xc6\x2d\x32\xc2\xd4\x9e\x0d\x28\x0b\xd0\xde\xf8\x35\x91\x72\xc3\x45\xba\x67\x66\xe9\x41\x3a\xc5\xe9\x6a\x51\xb7\x60\xed\xe3\x7e\x00\x66\x49\x7f\xee\x55\x4b\xb3\x87\x24\x09\xaf\x98\xb2\xa7\xce\x94\x0b\xa3\xb2\x9c\xf3\xe9\x36\x90\x6e\x82\xb6\x01\x7e\x3e\x83\x9c\x93\x14\x16\x24\xd7\x96\x5e\x1c\xb5\x18\xec\xd2\x19\xdb\x8b\x35\x76\x8b\x44\xac\x9e\xe4\x14\x99\x9a\xa6\x30\x20\x77\xe4\xcc\x70\xfe\xe2\x28\xa4\x3c\x31\x00\x7d\xec\xb6\x53\xad\x06\xba\xe9\x9e\x4d\x7d\x18\xac\x02\xfa\x43\xd5\x1f\x4e\xbf\xed\x0b\xa7\xcd\xae\xdb\x31\xa4\xc3\x9d\xf3\xc4\x9f\xd5\x6f\x0d\xa3\xdb\x1a\xb8\x2b\x90\xde\x1f\x0e\xbf\xfd\x96\x70\xf8\x11\x02\x94\xf6\x7a\x1d\x1f\xc7\x82\x04\xbe\x37\x5a\x8b\x73\xcb\x36\x65\x36\xb4\x99\x29\xc2\x52\x22\xd2\xf9\xc5\xe9\x7c\x7d\xba\x3f\xc0\x39\xfd\x61\x01\xce\xb3\x1f\x16\xe0\x3c\x0f\x02\x9c\x5e\x4d\x0c\x72\xd2\x5a\x19\xdf\x93\x64\x45\x75\xa0\xbf\x31\x79\x85\xf6\x50\x82\x4a\xec\x68\x56\xad\x1c\x16\x89\x9b\x66\x70\x25\x9c\x29\xc1\xf3\x79\x99\x13\x86\x73\x66\xba\xc9\x42\x8b\xf2\x08\x4b\xe8\xe4\x15\x45\x8d\x3b\xac\x7c\x50\x09\x72\x65\xaa\xa2\x0b\xad\x74\x6b\x92\x57\x08\x39\xbd\x43\xa0\xe5\x99\xe9\x03\x36\xf7\x7d\xbe\x8d\x88\xc0\x9a\x0a\x55\x91\x1c\xa6\xd7\x23\x3d\xec\x31\x69\x7b\xa9\xed\x94\xbd\x61\xaa\x3b\x5b\x21\xa9\xa4\xe2\x05\x0a\xe9\x52\x6b\xd3\x2d\x6e\x63\x88\xa2\x62\xe6\x56\x6c\x4f\xc1\x86\x94\x74\x8e\x2c\x2d\x39\x35\x26\x3a\xae\xb9\x34\x05\x8f\x77\x64\x81\x79\x6c\x21\xbc\xca\x13\xc8\xf5\xe0\xfd\x85\x14\xb3\xf1\xfe\x09\x76\xac\xe5\xf9\xa3\x92\x88\xbb\x7e\xaf\x7b\xc6\x35\xa7\x6b\x3a\xbb\xf2\x88\x28\xad\xa4\xe6\x54\x61\x1d\xe4\x6c\xf6\xa6\xc7\x53\xd4\x20\x6d\xaa\x57\x5c\xaa\x3d\xf3\xcc\x70\xdb\xb3\x19\xa1\xf6\xcc\xb1\x87\xc2\x8c\xb6\x1d\xda\xdd\x2f\xb2\x0e\xd7\xeb\x4c\xcd\xaa\x2d\x18\xb5\x1d\x5a\x4d\x32\x97\xb8\x54\x02\x67\xf9\x16\x88\xbf\xdd\xa5\x4b\xff\x4d\x03\xe6\xe6\x1e\xda\x5d\xd3\x0f\x77\x38\xd3\xb8\x30\x67\xa4\x21\xeb\x65\x0b\x77\x1a\x62\x95\x6f\xab\x82\x9b\xd4\xd6\x98\x6e\xb3\x5f\xff\xed\x97\x6d\xcf\x83\x11\xf8\x7e\x7a\xd7\x65\xdd\xa7\x45\xcd\xd5\xbf\xef\xb6\x71\xb3\x9b\x63\xe5\x7a\xf4\xcc\x07\x0d\xb1\x13\xf4\x77\x42\x71\xad\x55\x96\x24\xc1\x7d\x68\x28\x6b\x53\x62\xa7\x84\xa2\xa6\x85\x26\x7c\x40\x97\x75\x32\xe1\xec\xe1\x91\xe5\xa5\xeb\x7f\xfc\xa3\x22\xdb\x21\xe5\x23\xc9\x0b\xcc\xaa\xad\xeb\x4c\x8c\x7c\xb5\xc5\xf4\x12\x4e\x1a\x2a\x83\xbc\xc0\x4e\x30\x3d\x6a\xf6\xbb\x0f\xdb\x9d\x38\xaa\xbb\x93\x14\x87\xc1\xfa\x74\x78\x72\x32\x7c\x6a\x9c\x46\x84\xbb\x49\x18\x82\x5a\x93\x0c\x85\x50\x37\xb5\x1c\x6f\x68\xea\xdf\xfe\xea\x61\x67\x96\x3b\x54\xc1\x8a\xac\x31\x06\x35\x3d\xf8\x50\x0a\xba\xa6\x39\x66\x28\x7f\x6d\xb2\x45\xff\x7d\x84\x81\x0b\x35\xae\x66\xbf\xeb\x73\xa0\x2a\x34\x8f\xce\x76\x59\xc4\x4e\x96\x9d\x22\x90\x1e\x9c\xb3\x06\x4f\x58\x69\xfc\xea\x5b\x9f\x9f\x1e\x7a\xeb\xf3\xf3\xe3\xde\xfa\xfc\xf2\xb0\x1b\xb9\xff\xda\x77\xec\xda\x97\xe6\xc4\xca\xbb\xc9\xed\x65\x65\xac\xd2\xb2\xca\xbb\xf9\x3d\x04\xc5\x74\x87\x5f\xc2\x06\x05\xda\x4e\xc0\x50\xb7\x9a\xdb\x83\xf8\x9a\xa8\xa7\x59\x35\xb8\x21\x7c\x8c\x43\x6c\xa8\xeb\x60\x68\x57\x95\x9d\x71\x99\x17\x84\x65\xed\xfe\x04\xd3\x70\x3a\x43\xa5\x93\x25\x2b\xf1\xd7\xc8\x50\xd0\xa4\x69\x47\xb5\x0d\xbc\xe1\xc1\xf9\x6a\xed\x7a\xfe\x50\xed\x7a\xf1\xb8\xda\xf5\xd3\xc3\xb4\xeb\xe7\xbd\xe2\xfc\xab\xd5\xab\xdb\x25\xfc\x57\x6b\x17\xe9\x57\xaf\x4b\xdf\xd3\xed\xd7\xfa\x77\x52\xa8\x31\x54\x8c\xfe\x51\x21\x50\xf7\x39\x62\xea\x62\xcd\x76\xd3\x4d\x5f\x5b\xd3\xcf\x0f\xd3\xcb\x5f\xf6\xe9\xc1\x5f\xad\x96\x1d\xfb\xd1\x7b\x31\x19\x66\x8a\xb6\x37\xbe\x40\xa6\xf6\x46\xcf\xed\xe9\x2e\x7a\xe1\x6e\xbe\xad\x65\x34\x1f\x84\xb5\xfa\xcf\xfa\x83\x97\xa8\xf7\xc5\x50\xdb\xaa\xcb\x7a\x96\xba\x77\x66\x78\xe7\x35\xb1\xc5\xc0\x76\xa8\xfb\xf3\x3d\x73\x5c\x78\xd2\x1f\xb1\xbc\x88\x27\x12\x81\x4e\xa2\xb6\x3f\x7d\xc0\x50\xea\xb8\x60\x4b\x8a\x1c\x8e\xcd\xd0\x47\x9d\x48\xc8\xa1\x79\xa3\xe3\x67\x64\x4a\x1e\x75\x32\x8d\xd8\xff\x9b\x90\x9d\x6c\x4d\x49\x87\x37\x1f\x6e\x0c\x4c\xc6\x41\x45\x46\xca\x52\x07\x1f\x3c\xaf\xcc\xb2\x41\x88\xb5\xd8\x2a\x94\x16\x7c\xee\x31\x18\xdd\x8e\xaa\x07\x6d\x4e\xb6\xae\xdc\x1a\x8e\x7c\xb8\x79\xd7\x04\xff\xb5\x4c\xda\x05\x80\xa8\x9f\xb0\x07\x8b\x11\x54\x2d\x8a\xae\x5e\x9d\xf6\x64\xe1\xf7\xf4\xeb\xf7\xc7\xf0\x0f\xb6\x21\x7f\x4d\xff\x54\x4f\xb7\x97\xcb\xbc\xd1\x68\x56\xc2\x99\xd4\x04\x6a\x56\x69\x29\xb8\xe8\xba\xe7\x64\x76\x1a\x75\xf6\x7f\xc1\x10\x38\x27\xc2\xda\xad\x51\x6e\x95\xdd\x9d\x51\x86\x6c\x77\x24\xec\x2d\x66\xc9\xa5\xa4\x3a\x50\xb6\x1f\xd7\x33\xbe\xe9\x0d\x93\xeb\x39\x6d\x3b\xd5\xf9\x54\xe1\xff\xc5\xdb\x16\x6f\x0f\xef\x0d\x92\x0d\x86\xed\xd5\xfc\xd7\xfb\xcc\xd6\x37\x39\xb7\x9f\xf6\x48\xa9\xd7\xb7\x51\x09\x24\x70\x6a\xbb\x3b\x5d\x03\xb4\xed\x8f\x28\xf7\x76\x43\x3e\x90\x81\xfd\x7d\x4d\x0f\xd6\xa7\xd3\x6f\xd7\xa7\x67\x0f\xd5\xa7\x46\xf6\x35\x25\xc9\x4a\x67\x0f\x76\x5e\xab\xab\xc7\x0d\x7d\x4b\x47\xcf\x68\x14\x49\x20\xee\xe9\x89\xae\xf4\x62\xb8\x56\x57\x4f\x70\x8d\x17\xc1\xed\xed\xea\x79\x60\xc4\x15\x39\xb3\x5d\xab\x45\x2c\x68\x15\x54\x37\x08\x1b\xc2\x94\x4d\xec\xd3\xbe\x62\x24\xec\xaf\xae\x92\x34\xad\x4b\xab\xad\x9b\xfd\x9d\x2b\x09\x2c\xf8\x1a\x61\x29\x78\xf1\x35\xcb\xdd\x18\xf0\x70\x51\x8b\xa0\x5e\x37\xee\x32\x89\x2b\x8f\xdd\xc9\x1d\x8d\xd8\x55\x47\x84\x4e\x2d\xf1\xa4\xc7\x3f\x77\xbf\x6f\x7e\x0c\x5b\xb0\xe3\x9b\xad\x6e\x32\xd5\x0e\xec\x3a\xe7\xfc\x51\x5a\xa8\xee\xa1\xec\xb1\xb7\xdd\xfd\x34\xeb\x87\x6d\xfc\xde\x4b\xaf\x7b\xa9\xfb\xee\xcd\xff\x4f\x00\x00\x00\xff\xff\xd8\xb2\x08\xb4\x90\x47\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 18320, mode: os.FileMode(420), modTime: time.Unix(1539059833, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
