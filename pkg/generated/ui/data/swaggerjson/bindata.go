// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x6d\x6f\xe3\x36\x12\xfe\x9e\x5f\x31\xd0\x1d\x70\xbb\x40\x36\xde\xee\xe1\x80\x22\x9f\x2e\x4d\x70\xad\xb1\x4d\x1b\xd4\x8b\x06\xc5\xb5\x30\xc6\xd4\xd8\x66\x4d\x91\x5a\x92\xb2\xd7\x7b\xc8\x7f\x3f\x90\x92\x6c\x49\x96\x6c\x45\xca\x26\xee\xc2\x05\x8a\x78\xcd\x97\x79\xe5\xf0\xe1\x0c\xe9\xff\x9d\x01\x04\x66\x85\xb3\x19\xe9\xe0\x12\x82\x77\x17\x6f\x83\x73\xf7\x1d\x97\x53\x15\x5c\x82\x6b\x07\x08\x2c\xb7\x82\x5c\xfb\xb5\x48\x8c\x25\x0d\xb7\x28\x71\x46\x1a\xae\xee\x86\xbe\x3f\x40\xb0\x24\x6d\xb8\x92\xae\xd7\xf2\xed\x45\x3e\x11\x40\xc0\x94\xb4\xc8\xec\x66\x36\x80\x40\x62\xe4\xa7\x1b\x61\x64\x12\x39\x83\xeb\x9f\xae\x3f\x64\xdd\x01\x82\x44\x0b\xd7\x38\xb7\x36\x36\x97\x83\xc1\x8c\xdb\x79\x32\xb9\x60\x2a\x1a\x98\xb4\xff\x1b\x26\x99\x1d\xb0\x94\x97\x37\x51\xca\xcb\x1b\x8c\xf9\x76\x0e\x8a\x90\xfb\x59\x30\x8c\xb8\xfc\x77\x71\xe0\x05\x57\x81\xef\xf6\x70\x06\xf0\xe0\xa5\x35\x6c\x4e\x11\x99\xe0\x12\xfe\x9b\xf2\xec\x69\xe7\x02\xb8\x7f\xb8\x11\x7f\xf8\xbe\x4c\x49\x93\x94\x3a\x63\x1c\x0b\xce\xd0\x72\x25\x07\x7f\x1a\x25\xb7\x7d\x63\xad\xc2\x84\xb5\xec\x8b\x76\x6e\xb6\x2a\x1f\x60\xcc\x07\xcb\x6f\x72\x29\x8b\xda\x9b\x51\x51\x99\x8e\xfd\x24\x8a\x50\xaf\x9d\xb8\xf7\x5c\x08\xd0\x64\x35\xa7\x25\x81\x9d\x13\x18\x8b\x36\x31\xa0\xa6\x80\x90\x4d\x06\x28\x43\xe0\xd6\xc0\x22\x99\x10\x53\x72\xca\x67\x30\x55\x1a\x98\x92\x92\x98\xe5\x4b\x6e\xd7\x1b\x55\x02\x04\x2a\x26\xed\x59\x1e\x86\x8e\xc6\xf7\x64\x33\x3f\x28\x76\xd2\x64\x62\x25\x0d\x99\x12\x6f\x00\xc1\xbb\xb7\x6f\x2b\x5f\x01\x04\x21\x19\xa6\x79\x6c\x33\x8f\x29\x4c\x94\x4a\xe4\x0c\x82\x3b\xc3\x00\x82\xbf\x6b\x9a\xba\x11\x7f\x1b\x84\x34\xe5\x92\xbb\x19\x4c\xae\xa5\x71\xe6\x0b\x63\x8c\xf9\x96\xcb\x5f\x28\x16\xeb\xa0\x34\xd1\xc3\x59\xdd\xe7\x87\x82\x38\x31\x6a\x8c\xc8\x92\xde\x1a\x2f\xfd\xaf\x22\x48\xee\xca\xfe\xef\xf9\x5e\x21\x7f\xc2\x88\x9c\x1d\x9c\x55\x72\x4b\x58\x05\x13\x02\xa1\xd4\x82\x42\x48\xe2\x8b\xea\x14\xdc\x8f\xfc\x98\x90\x5e\x57\x9b\x34\x7d\x4c\xb8\x26\x67\x92\x29\x0a\x43\x95\x66\xbb\x8e\x3d\x63\xc6\x6a\x2e\x67\x45\xf1\x1f\xce\x0f\x8b\x83\x2b\x73\x61\x88\x69\xb2\xe3\x05\xad\xc7\x3c\x3c\x20\xdb\x87\x39\xc1\xc8\xf7\x7f\x4f\xeb\x61\xe8\xdd\xe9\xea\x6e\x08\x57\x8c\x91\x31\xc7\x28\x16\x7a\xce\x9c\x74\xad\x45\x4b\x85\x79\x4f\xeb\x8d\x78\x78\x7c\xe2\x69\x9a\x39\xc6\x0f\xcb\xf4\x8b\xef\x78\xb4\xa2\x7c\x4e\x34\x5d\x60\x1c\xb7\xf3\xbd\xab\x38\x3e\x62\xaf\xf3\xb2\x58\x92\x28\x6d\x0b\x59\x3e\xf8\x8e\xc7\x6d\x98\x18\x8d\x59\x29\xdd\xc6\x34\x77\x59\xd7\xe3\x16\xc8\x24\x93\x0d\xe7\x2d\xc3\x5d\x61\xc4\xb1\xca\x16\x6b\xb5\xe4\x61\x69\xa3\xae\x13\xa7\xb8\x33\xe5\x43\x0c\xbc\xc2\x95\x19\xe0\xc2\x0c\x96\xd1\x0a\x35\x0d\xc8\xb2\xd7\xcf\x23\xd9\xe6\xf3\x1f\x85\x1d\xd9\xe2\xac\xba\x17\xe7\x78\x74\x3b\xf8\x8f\xb3\x8a\x72\x82\x90\x04\x59\xda\x0f\x98\xd2\x3e\x5b\x80\xb4\x07\xfc\xdc\xf8\xae\x7f\x01\xfc\x53\x62\xf4\x58\x20\xd0\xfd\x1c\x2d\x70\x53\x84\x40\xff\x30\xe0\x06\x3a\x24\x14\x92\xb1\x5a\xad\x8f\x66\xf5\x9c\x40\xd0\x09\x04\xbd\xac\x28\x27\x10\x74\xcc\x86\x39\x81\xa0\x13\x08\xfa\x72\x92\x3d\x25\x08\x8a\x93\x03\x29\x23\x2f\xb0\x71\x46\x6c\x03\x82\xae\x35\xe1\x5f\x02\x04\x95\x18\x7d\x16\x10\x34\x51\xe1\x8e\x0f\xa4\xee\x51\xd7\x52\xf0\x0e\xab\x93\xaa\x73\x3c\xb5\x02\x6e\xcd\xac\x8d\xf8\xdd\xfd\xed\xac\xa0\xbd\x6a\x06\x73\x20\xb8\xb1\xdd\xd2\x98\x08\x6e\xac\x5b\x9b\xd9\x5c\xa6\x55\x76\xf2\x47\x47\xf0\x88\x9d\xb3\xcc\xe9\xb3\x78\xe7\x09\xd1\x9e\x10\xed\xcb\x8a\x72\x42\xb4\xc7\x6c\x98\x13\xa2\x3d\x21\xda\x2f\x27\xd9\x97\x42\x18\xdb\xd2\xf3\xa3\xc0\x45\xa2\x25\x64\x43\x81\xcb\xa9\xd2\x91\x87\x11\x80\x13\x95\x58\xc0\x98\x83\x21\xbd\xdc\x0b\x84\xbf\x27\xfb\x6b\x3a\xc3\x70\x3b\xc1\x91\x63\x8e\x8c\xe1\x4e\x78\xa3\x8b\xb1\x36\x55\xf6\x02\x6b\xdb\x3a\x77\x09\xa0\x5e\xbd\x1f\x8d\x62\x62\x57\xef\x47\x43\x69\x2c\x4a\x46\xdf\x6b\x95\xc4\x45\xc3\xe6\xee\xa5\x26\x7f\x12\xdb\x86\x5d\xb7\x24\x62\xd2\x96\x57\x34\x9d\x2f\x99\x92\xee\x2b\x2e\x7a\x5e\x6a\xcb\xaf\x3a\xb8\x58\x20\x0b\x0b\x68\xe6\x39\xa9\xd7\x4a\x3a\xdf\xe3\x69\xe4\x52\x82\xeb\x0d\xaf\x46\x16\x65\x88\x3a\x1c\xdf\xbc\x1b\x2f\xdf\x9d\x03\x59\x76\xf1\xba\x9e\x64\xc4\xe5\xf8\x63\x82\xd2\x72\xbb\x6e\x22\xcd\xa5\xa5\x59\x25\x48\x04\xa9\x97\x66\xcd\xff\x7c\xd7\xc0\xd8\x2d\x97\x3c\x4a\x22\x90\x49\x34\x21\xed\x54\xc0\x33\x56\x0d\xbc\x0a\x69\x8a\x89\xb0\x06\xac\x82\xcf\xa4\x55\x13\x8b\xf8\xe9\x8b\xb2\x88\x9f\xba\xb1\x78\x56\x61\xb5\xc6\x1c\xde\xd6\x06\xbc\xc3\xba\xe3\x87\x37\x0f\xca\xb0\x8e\x58\x50\x0a\x49\x65\x77\xbe\x4f\xdd\xf9\x7e\x74\x83\x16\xaf\x49\x56\x2e\x72\x3c\xd6\x97\x33\x0c\xd8\xc5\xd3\xee\xe7\x9c\xcd\x21\x9d\x00\x5e\x25\xe6\x0d\xa1\xb1\x6f\xbe\xd9\xeb\x63\xb8\x44\x2e\x70\xc2\x05\xb7\xeb\xf1\x67\x25\x77\x83\x58\x4e\x1a\xb5\xc6\xf2\x3e\x11\x70\x4b\x51\xb5\x7f\xfb\x7d\xaf\xca\x78\x91\x15\xf0\xac\x14\x84\x98\x9c\x43\xfe\xf9\x1d\xf3\x9f\x57\xe4\x3e\x87\xbb\xd2\x35\xdb\x3e\xa5\xe3\xec\x04\x99\xa1\xda\xd8\xf5\xc9\xc2\xd4\x53\x85\x90\xe8\x5f\x17\x02\xf5\x8c\x4e\xc1\xe3\xeb\x09\x1e\x77\x3a\xbb\xa4\x95\x68\x0a\x87\x95\x75\xf5\x68\x4f\x5b\xc6\xcc\x01\xe1\xae\x5b\xe2\xaf\x77\xd7\xc0\xc3\x73\x98\x08\x94\x0b\x8f\x8f\xdd\xff\xbf\x07\xcc\xf3\x0e\x4a\x92\xff\x62\xad\x92\xdf\x83\x73\x98\x72\x21\x28\x04\x3e\x75\x5f\x00\x6a\x82\xef\x7e\xfb\xd9\xcd\x51\x6f\x75\x43\x2c\xd1\x2e\xda\x78\x1d\x76\x65\x73\x94\xcd\xb2\x6f\xcb\xe6\x18\x8d\xb5\x12\x34\x46\xdd\x2d\xa0\x3a\x5d\x0c\xaf\x6e\xc1\x4d\xe2\x25\x2e\xde\xea\x7a\x85\x5a\xbe\xce\xed\x68\x8c\x62\xdc\x03\xc3\x30\x6c\xe5\x49\xff\x51\x1a\x56\x73\x92\x60\x54\x44\x60\xe7\x5c\xce\x8c\xd7\x1d\x0a\x4d\x18\xae\x21\xd5\x75\xb8\xc7\x79\x7e\xbd\xbd\x47\x4d\xce\x7f\xde\x27\x13\xd2\x92\x2c\x99\x1f\x71\x42\xe2\xc5\xa1\x14\x82\xf0\x7c\xd4\x5a\x65\x89\x22\xe9\x4e\xc2\x8f\xae\xa7\x51\x7b\x7c\x68\xd0\x58\xfa\xe9\x16\xd9\x9c\x4b\xff\x45\x1f\x9d\x25\x86\x74\x2f\xbd\xe5\x13\x78\x1f\x1b\x8d\x7e\xc8\x8e\xa2\xf5\xfa\x9b\x2b\x63\x3b\x53\x72\x83\x5b\x51\x89\x95\x6e\xa4\xd2\x23\x26\x3b\x1e\xdc\xd4\xed\x78\xc8\x13\x15\x8f\x90\xb6\xe6\xb0\xbf\xf8\xd6\x6c\x8e\x80\x9b\x45\xac\xa4\xd5\x4a\x40\x2c\x50\xd2\x05\x7c\x98\x73\x03\x52\x85\x04\xdc\x80\x92\x62\x0d\x08\x11\xfa\x65\xce\xdd\xc9\x80\x1b\x98\x72\x12\xa1\x6b\x4e\x63\x7e\x78\x51\xcf\xb3\x77\xca\x27\xc0\x50\xf5\x87\xbd\x96\xeb\xff\x30\xe2\x72\x7a\x49\x59\xdd\xa8\x24\x4a\x17\x83\x69\x15\xbe\xdc\x78\x13\x13\xe3\xd3\xec\xa2\xb3\x9f\x06\x37\xdf\x79\x65\x96\x63\x57\xe5\x54\x5a\x73\xa4\xee\xb1\x04\x67\xdc\x8e\x77\x33\x04\x07\xbd\xa5\x24\x8f\xc5\x19\x28\x99\x9e\x03\xb9\x05\x4d\xb1\x32\xdc\x2a\xbd\xae\xb7\xb4\x23\xc9\x54\x14\xf1\x1e\xab\x11\xcd\x7c\x73\xf4\xe4\x16\xb2\xe9\x1a\xc9\x59\x4d\x34\x36\x16\x6d\xb7\x50\x73\x3f\x27\x3b\x77\xe0\x44\x83\x54\xd6\x53\x75\x33\xc2\x0a\x0d\x30\x41\x28\xd3\x1d\x69\x92\x70\xd1\xc0\x84\x6b\x0a\xc7\x61\x57\x06\x6e\x3c\x7a\x98\x7a\x0a\x61\x83\x98\xaa\x97\x1d\x33\xaf\x72\x44\x66\xca\xc5\xd5\xd0\x01\x3c\xa6\xa2\x98\x0b\xaa\xa7\x98\x35\xea\x4e\xf4\xae\xb3\xc1\x9e\x54\x43\x14\x13\x68\x9d\x8f\x77\x9a\xff\x2e\x1b\x0c\xdc\xa6\x66\x4a\xe9\xa5\x09\xd9\x01\xe8\x44\x4a\x2e\x9d\xdb\x1e\xda\x07\x6b\xd2\x44\x57\xf7\xa3\x6b\x4d\x21\x49\xcb\x51\xf4\x82\x9b\xe5\x6a\x53\xd7\xd5\xd0\x5c\x83\x6a\x04\x92\x95\x62\x50\x3f\xca\x4d\x25\xa2\x7a\xea\x3d\xce\xe9\x8d\x15\x9c\xd6\xa1\x97\x6d\xed\xe6\xfc\x3b\x31\x29\x74\xf0\x98\xd1\xf9\x43\x01\xa7\x06\x07\xdd\xe0\x73\xa2\xf3\x2d\x65\x44\x7a\xc9\x19\x5d\x31\xa6\x12\x69\xfb\xb8\x04\x13\x9c\xa4\xed\xe3\x0e\xd7\x7e\x86\x61\x08\xaf\x70\x81\x97\xbe\x4a\x74\xd3\x70\x90\xcc\x88\xa5\x0e\xd1\x93\x60\xea\x0b\x19\xd1\x1c\x81\xb4\x3b\x1d\x3a\xad\x63\xaa\xba\x34\xf4\x4c\xd6\xd5\x77\x20\xd9\x09\xca\xd7\x2d\x40\x93\x51\x89\xf6\xc7\x51\xfa\x74\x09\x42\x61\x08\x13\x14\xee\xd0\xa8\x5f\xb7\xb4\xdb\xd3\x2c\xe0\xac\x56\xd7\x55\x73\x75\x15\xbc\x86\x2c\x6a\x5a\x48\xeb\x4a\xa8\xb6\xbc\xf6\x74\xe0\xb1\x44\xab\xa1\xf6\xd5\x10\x8b\x2a\x25\xa8\xce\x91\xa8\xb9\x30\xf5\x22\xc1\x21\x8b\x0b\x37\x64\x91\x8b\xa1\xa5\xa8\x8f\x97\x75\xd4\xcb\xf0\xa6\xf2\xa0\xaa\xde\x04\x9d\x4f\x60\x35\x4f\xb6\x1a\x8c\xec\x5f\xd9\x75\x84\x5e\xdb\x3b\xd1\xdb\xc7\x7a\x07\x29\x6e\xdf\xee\xf5\xa6\x5a\x78\x06\xe8\x11\x91\x7f\x05\xe8\x3e\xd6\x33\xd1\x16\x44\x64\x0e\x72\x72\x8d\x67\x71\x8d\xd6\x56\xa9\x29\xb8\xf5\x31\x8f\x50\xac\x7a\x44\x6b\x2f\xa0\xdf\x21\xfc\x86\xb7\x93\x79\xaf\xaa\x91\xd5\x6e\x67\xd0\x7c\x20\x6e\xb3\x2f\x36\x20\xf7\x72\x94\xf4\x47\x92\xc3\x76\xce\x1a\x77\x01\x52\x17\x1e\x6b\x31\x57\x03\xb7\x19\x84\x30\x69\xe7\x32\xcc\x70\x8b\x18\xc5\xc2\xfd\x4d\x81\x85\x87\x13\x9b\x78\x5f\x87\x2b\xaa\x72\xe5\xc9\xea\x34\x29\xfb\x3c\xf9\x8b\xa6\x42\x70\x61\xa2\x86\xec\x45\x25\x0d\xdf\x7b\x81\xdc\xf7\x5e\x20\x21\x5a\x1c\xb3\x6a\xf1\xaf\xb5\x26\xea\x6a\x88\x7b\x16\xd3\xfd\xe8\x99\x96\x52\xf9\x84\xf8\x94\x0b\x69\x83\x7d\x7b\xa9\xab\xa6\x6a\x52\xcf\xe4\x77\xbf\xfd\x0c\xa9\xa7\x1e\x8f\xfb\x37\x14\x18\x9f\xdd\xfd\x6f\xcd\xec\x45\x2a\x06\xcd\x4f\xe5\x37\x17\xd5\x1b\x73\x2a\xf9\xdd\xa8\xae\x9e\x5d\x52\xc0\x5d\x36\x9b\x0f\x02\x7b\xf2\xd6\x59\xb7\x72\xe2\xb3\xaf\xf6\x4b\xc4\x5f\xc2\x0c\x45\xfc\xb1\x11\xf1\x0d\xb0\x44\x6b\x92\x56\xac\xd3\xf4\x37\x37\x80\x2b\x03\x4a\x43\x84\xd8\xb0\x8a\x16\xdf\x9a\xde\x49\xd8\xe5\x36\x81\xb7\xcd\x6a\x37\x5c\x5d\x58\x75\x8f\x6c\xb5\xf1\x7f\x7f\xc4\x6d\xb2\x7a\x91\x23\xb7\xfb\x3e\x11\x4f\x19\x68\xab\xa5\x93\x5e\xea\x7b\x1a\x42\xdb\x2a\x42\x43\xb9\x89\xcf\xe6\xe3\xe2\xd5\x8c\x26\xd3\x4e\x94\x12\x84\xb2\xa9\x2a\x54\xdb\xbc\x2f\x35\x9d\x87\x04\x6e\xe0\x87\xab\x06\x5c\x4f\x76\xa5\xf4\x62\x3c\xc5\x89\xe6\xac\xb3\xcf\xa5\xc3\xb3\xd8\x53\x49\xe4\x76\x5a\xd3\xe9\xb5\xbb\x1e\x8b\x59\x2d\x9e\x5b\xcb\x2b\x34\xc5\xb0\x9b\xa6\xae\xb8\x01\x4d\x1f\x13\x32\x0d\xe5\x80\xdd\xdf\xad\x79\xa4\x23\x16\x0e\x8f\xcd\xe6\x09\x7d\xf2\xa1\x7a\x2a\xca\x39\x83\xfc\xf6\x65\x5f\xa3\x15\x16\x42\x9f\xbc\x67\x5a\x58\x1c\xfb\xc2\xe2\x58\xaa\xf0\x29\x2e\x54\x3d\xaa\x18\xb8\x5b\xda\x3e\x8c\x28\xb2\xde\x06\x56\xfe\x86\x14\x53\x51\xac\xb9\xa1\xc3\x30\xce\x2d\x3f\xd2\x5f\xbf\x9c\x18\xf3\x31\xc9\x30\x56\xbc\x73\xf6\x92\x1b\x30\x73\x95\x88\xd0\x45\x19\xcc\x2e\x32\x08\xbe\x20\xe0\xf1\xa5\xaf\x89\x5b\xb7\x13\xaf\xb8\x10\x59\x0f\xae\x6d\x82\x02\x86\x77\x03\xd7\xfc\xbb\xbc\x43\xe3\xb3\xca\xc8\xfc\x59\x8f\x3e\x59\xd2\x12\x05\xb0\xc4\x58\x15\x91\x36\x59\x08\xc3\x89\xa0\xac\xec\x15\x25\xd2\xed\x58\xd4\x3f\xd1\x53\xf3\x53\x03\xc7\x17\xe0\xae\xbd\x7e\x8b\x51\x62\x42\xf9\x2f\x0e\x34\x61\xc9\x1e\xa9\x9b\x51\x29\x5d\xb3\x1b\x2a\xdb\xea\xb6\xee\x91\xd8\xf1\x29\x77\x68\xca\xfb\x72\x5a\x2a\x37\x6b\xe3\xc2\xf7\xbe\xfd\xe1\x8b\x45\x86\x03\x5b\xca\xe1\x78\xf0\x63\xf5\xc9\x61\x1f\xdb\x7d\x75\x76\xeb\xbb\xaf\x17\xaa\x06\x9d\xf4\x5a\x7a\xc4\x70\x84\x7a\x9d\xc2\xe6\xa5\x88\xc7\x4f\x3f\xbf\x6f\x80\xeb\xa9\x1c\x63\x5e\x7b\xd7\x65\x8f\x6a\x0f\x5f\x99\xd9\x7f\x0b\xa2\xd8\x73\xd7\x02\xbb\xd5\x23\x4f\xc5\x17\x8b\xf2\x63\x58\x29\xa6\x6d\x5f\x79\xe4\x5b\xcf\x8d\x62\x85\x67\x1e\x95\xbb\x57\xb7\x4a\x53\xf6\xce\x66\xcf\x4f\x4a\xf6\xf8\x21\x48\xc7\xce\xd9\xc3\xd9\xff\x03\x00\x00\xff\xff\xce\x4a\xc3\xf3\xe0\x52\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 21216, mode: os.FileMode(420), modTime: time.Unix(1537158883, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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
