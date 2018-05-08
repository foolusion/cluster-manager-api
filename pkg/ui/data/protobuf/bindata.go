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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdf\x6f\xdb\x38\x12\x7e\xf7\x5f\x31\xd0\xcb\x39\x87\xc6\x4a\xd3\x3d\xe0\xe0\x20\xd8\xcb\xb9\xbd\xd6\xb8\xd6\x09\x62\x6f\x83\x7d\x0a\x68\x6a\x2c\xf1\x4a\x91\x5c\x92\xb2\xeb\x3b\xf4\x7f\x3f\xf0\x87\x64\x49\x96\x53\x6c\x36\x40\xd7\x0f\x6d\x44\x0e\x3f\xce\xf7\x71\x38\x33\x52\x9a\xc2\x4c\xaa\xbd\x66\x79\x61\xe1\xf2\xe2\xf5\xdf\x61\x49\x4a\x53\x89\x1c\x96\x6f\x97\x30\xe3\xb2\xca\x60\x41\x2c\xdb\x22\xcc\x64\xa9\x2a\xcb\x44\x0e\x2b\x24\x25\x90\xca\x16\x52\x9b\xc9\x28\x4d\x47\x69\x0a\x1f\x19\x45\x61\x30\x83\x4a\x64\xa8\xc1\x16\x08\x37\x8a\xd0\x02\xeb\x99\x57\xf0\x19\xb5\x61\x52\xc0\xe5\xe4\x02\xc6\xce\x20\x89\x53\xc9\xd9\x95\x83\xd8\xcb\x0a\x4a\xb2\x07\x21\x2d\x54\x06\xc1\x16\xcc\xc0\x86\x71\x04\xfc\x4a\x51\x59\x60\x02\xa8\x2c\x15\x67\x44\x50\x84\x1d\xb3\x85\xdf\x27\xa2\x38\x4f\xe0\xd7\x88\x21\xd7\x96\x30\x01\x04\xa8\x54\x7b\x90\x9b\xb6\x21\x10\x1b\x9d\x76\xbf\xc2\x5a\x35\x4d\xd3\xdd\x6e\x37\x21\xde\xe1\x89\xd4\x79\xca\x83\xa9\x49\x3f\xce\x67\xef\x16\xcb\x77\xe7\x97\x93\x8b\xb8\xe8\x17\xc1\xd1\x18\xd0\xf8\x5b\xc5\x34\x66\xb0\xde\x03\x51\x8a\x33\x4a\xd6\x1c\x81\x93\x1d\x48\x0d\x24\xd7\x88\x19\x58\xe9\x9c\xde\x69\xe6\x74\x7b\x05\x46\x6e\xec\x8e\x68\x74\x30\x19\x33\x56\xb3\x75\x65\x3b\x9a\xd5\x2e\x32\xd3\x31\x90\x02\x88\x80\xe4\x66\x09\xf3\x65\x02\xff\xbc\x59\xce\x97\xaf\x1c\xc8\xc3\x7c\xf5\xe1\xf6\x97\x15\x3c\xdc\xdc\xdf\xdf\x2c\x56\xf3\x77\x4b\xb8\xbd\x87\xd9\xed\xe2\xed\x7c\x35\xbf\x5d\x2c\xe1\xf6\x5f\x70\xb3\xf8\x15\xfe\x3d\x5f\xbc\x7d\x05\xc8\x6c\x81\x1a\xf0\xab\xd2\x8e\x81\xd4\xc0\x9c\x9a\x98\x79\xe9\x96\x88\x1d\x17\x36\x32\xb8\x64\x14\x52\xb6\x61\x14\x38\x11\x79\x45\x72\x84\x5c\x6e\x51\x0b\x17\x09\x0a\x75\xc9\x8c\x3b\x55\x03\x44\x64\x0e\x86\xb3\x92\x59\x62\xfd\xd0\x11\xaf\xc9\x68\x64\xf6\xc2\x92\xaf\x70\x0d\x89\xd2\xd2\xca\x37\xc9\xd5\x68\xa4\x08\xfd\xe2\x80\x29\xaf\x8c\x45\xfd\x58\x12\x41\x72\xd4\x8f\x44\xb1\xab\xd1\x88\x95\x4a\x6a\x0b\x49\x2e\x65\xce\x31\x25\x8a\xa5\x44\x08\x19\x37\x99\x78\x98\xe4\xaa\x31\xf3\xcf\xf4\x3c\x47\x71\x6e\x76\x24\xcf\x51\xa7\x52\x79\xd3\xc1\x65\xa3\x51\x98\x85\x71\xae\x15\x9d\xe4\xc4\xe2\x8e\xec\xc3\x34\x7d\xcc\x51\x3c\x46\x94\x49\x44\x99\x48\x85\x82\x28\xb6\xbd\xac\x67\xce\xe0\x1a\xfe\x37\x02\x60\x62\x23\xa7\xfe\x2f\x00\xcb\x2c\xc7\x29\x24\xb3\x40\x09\x3e\x05\x4a\x70\x73\x37\x4f\xae\xbc\xc5\x36\x5c\x87\x29\x24\xdb\x8b\xc9\xeb\xc9\x45\x1c\xa6\x52\x58\x42\x6d\x8d\xe3\x7e\x82\x94\x0e\xea\x13\xa3\x05\x41\x0e\x9f\x51\xe0\x7f\x19\x89\xf6\xee\x57\x69\x3e\x85\xc4\x45\xb2\x99\xa6\x69\xce\x6c\x51\xad\x27\x54\x96\xe9\xf6\xc8\x14\x4b\xc2\x9c\x71\x19\xa7\xfe\x91\xbb\x01\x67\x1c\x8d\xbe\xb9\xff\xfc\x3f\xf8\xd5\xa2\x16\x84\x3f\x66\x92\x9a\xda\x9f\x93\x5b\x39\xf5\xce\x91\x4a\xb3\x37\x16\xe3\x63\x14\x33\x22\x67\x68\xa8\x66\x5e\x44\xc7\x46\x6a\x04\xb2\x96\x95\x85\xfc\xfe\x6e\x76\xfe\xbe\x65\xfb\x6d\x04\x60\x68\x81\x25\x9a\x29\x7c\x58\xad\xee\xae\xfa\x03\x4b\x37\x42\xa5\x30\x95\x1f\x4a\xe2\x0d\x74\xd8\xe9\x7f\x8c\x14\x1e\x46\x69\x99\x55\xf4\xd4\xfc\xb7\xab\xd1\xc8\xa0\xde\x32\x8a\x50\x9f\x52\xe0\xe8\x2e\x83\xbb\x19\x08\x1f\x90\x73\x09\x0f\x52\xf3\x0c\x96\xd1\xf6\x1c\x76\x8c\x73\xd0\xa8\x90\x58\x20\xe0\x6e\xba\x4f\x8b\x56\xfa\x50\x77\xa7\xe5\xb6\xde\xb2\x0c\x33\x8f\xa7\x15\x0d\x48\x01\x68\x7c\xf8\xfb\x93\xc9\xcf\x40\xa3\xad\xb4\x30\xed\xf1\x7b\x54\x7c\x7f\xd6\x0a\x81\x26\x46\xfd\x1d\x98\x10\xc5\x26\xee\x0c\xea\xc8\x73\x3f\x25\x8d\x85\x29\x24\xfe\x82\x6c\x5f\xa7\x85\x43\xdb\x39\xb4\x24\x5a\xac\x65\xb6\x9f\x42\xf2\xd7\xe4\x70\xd0\x41\xeb\x36\x65\x25\x33\xa0\xb2\x12\x16\x34\x1a\x25\xdd\xa5\x05\x78\x08\x8c\xdd\x73\x76\x48\xbc\xa2\x2a\xd7\xa8\x5d\x76\x55\x32\x33\x2e\xd3\xd5\xfc\x8d\x22\x74\x40\x84\xf7\x68\xef\x64\x36\xf3\xe8\xe3\xd6\x43\x57\x86\xd6\xc4\x73\x74\x18\x56\x23\x47\xab\x64\xe6\x89\x25\x1d\x43\x27\x0a\x1c\x54\x19\x52\xc6\xb3\xf7\x64\x7c\xfd\x22\x75\x9e\x6a\x78\xcd\x34\x12\x8b\x75\x10\x8d\x3b\x8f\x5d\x6e\x9d\xa9\x3f\xc0\xae\xea\x90\x8b\xfe\x3c\x8f\x98\x46\xab\x19\x6e\x43\xf2\x37\x96\xd8\xca\xb8\x23\x6d\x58\xba\xc4\x0e\xcc\x1a\xf8\x52\xad\x91\x4a\xb1\x61\xb9\xaf\x0d\x54\x0a\x81\xd4\xb2\x2d\xb3\xfb\xf6\x09\x37\x32\x1c\xfe\x3e\x3a\xdf\x3f\x2c\x40\x8e\x4f\x0b\x30\xc8\x34\x43\x8e\x16\x07\xce\xef\xad\x9f\x68\x1c\xef\x3c\x76\x7d\xef\x4c\x3d\xdf\xfd\xe8\xc9\xef\x66\xc0\x84\xb1\x84\x73\x18\x4b\x0d\x1a\xe3\xd3\x19\x58\xc6\x79\x8b\xce\x5d\x1d\xaa\x2b\x3f\x0e\xe3\xde\x40\x97\x52\x6f\xf2\xe5\xae\x5c\xf0\xea\x79\x41\x79\x82\x68\x81\xbc\x04\x5a\x10\x6d\x6b\xeb\x95\x6b\x12\x7d\x42\x5e\xa3\xab\x30\x56\x57\xd4\xb7\xab\xcc\x87\xb0\x33\x85\x82\x18\x20\x5c\x23\xc9\xf6\xb0\x46\x14\x90\xa1\xe2\x72\x8f\xad\x54\x66\x5c\xd2\x76\x99\xab\x11\x71\x1e\xf6\xfc\x80\xbc\x9c\x79\x94\x71\x7f\xa4\x2b\x63\x7f\xf6\xc5\x2e\xb7\xe3\xfc\x3c\x11\x63\x94\x35\x6c\x7b\xea\x1d\x22\xbf\x45\xb2\x37\x30\x14\xfd\x2f\x40\xf1\x38\xfe\xbb\x2c\x1b\x3a\xdf\x46\x23\x7f\xc8\xdd\x52\xec\x7a\x6f\x34\x76\x54\xa2\x31\xae\x6f\xec\x14\xd4\xb8\x95\x6b\x9e\x45\x1e\xaa\xf1\x35\xbc\xbe\x6a\x41\xd5\x85\xcd\x55\xec\x16\xec\x00\x9c\x67\xd8\x05\xac\x8d\x6a\xcc\xfa\xb9\x5b\xce\x0e\x9d\xc4\xa2\x29\x87\x56\xc2\x06\x2d\x0d\x01\xd7\x94\xd9\xda\xee\x23\x92\x2d\x02\x96\xca\xee\x9d\xe5\x6f\x15\xea\x3d\xb8\x2b\xd0\xd4\x53\xd3\xe7\x15\x60\x9f\x70\xa4\xed\xbe\x73\xe5\x3b\xf5\xda\x5d\xb6\xee\x8e\x67\x7e\x29\x13\xf6\xcd\x65\x58\xd3\xdf\xac\x5f\xea\xba\xbc\xeb\x17\xaf\xba\x92\x58\xe9\x6e\x69\x53\x49\x63\x67\x30\x74\x52\x83\xf8\x3d\x3a\x0f\x05\xfa\xf7\x19\xa9\xfd\x2b\x63\x7b\xa3\x1d\x31\xed\x6d\xdc\x3b\x9a\x7f\x9b\xac\x03\x27\xdc\x22\xc9\x41\x7e\x09\x5b\x36\xd9\x24\x94\xc0\x23\xcf\x73\x14\xa8\x7d\xfb\xd8\x76\x39\x56\xcb\x6b\xb8\xec\x9f\xc0\xef\x53\x84\x4b\xf9\xc5\xbd\x06\xaa\xef\xea\xd1\xab\x9d\x07\xf0\xb9\xe9\xe0\xc6\xb3\x0d\xcd\xf8\x49\xba\x0f\x05\xb1\xee\x55\xb3\x5b\xf9\x5b\x38\xa7\xc8\x0e\xac\x6f\x75\x07\x56\xd6\xcd\x41\xdd\x13\x0f\xc0\xb5\xec\xaf\xe1\x4d\x87\x64\xbf\xfe\xb6\x8f\xfc\xb0\x61\xc4\xfc\x8b\x09\x4a\x59\x19\xd2\xbf\xdc\x7f\x57\xc4\xe3\x22\x7e\xd8\x61\x26\x2b\x9e\x75\xa4\xac\xeb\x8a\x4b\xa1\x27\x95\x5c\x76\xd4\x6b\x87\xd9\xd3\xa1\x72\x5c\x99\x5b\xae\xd4\x31\x12\xaa\xb8\x29\xbc\x6b\x6b\xac\x8b\xa3\xff\x2e\xd0\xde\xa3\xf6\xb8\xed\x58\x2b\xfd\x9c\x86\x61\xe2\x54\x6a\x39\x9c\xb5\xff\x7a\x23\xa4\xa7\xe8\x91\x7c\xca\x0e\x5f\x38\x42\xad\x4e\xa1\x52\xb9\x26\x99\x3b\x8b\x36\x5e\x7c\xd1\x0d\xc7\xdc\x0d\xd7\xe8\x53\xd3\x91\x9d\xef\x58\x56\x8f\xfe\xdc\x68\x1b\x3c\x66\xae\x90\x6f\xb1\x6b\x4a\xb2\x92\x09\x50\x9a\x6d\x19\xc7\x1c\xcd\xcf\x87\x13\xaa\x3f\x26\x78\xbb\x6b\xf8\xe9\x58\x12\xe7\x83\x8b\x27\xdb\x12\xc5\x7f\xc4\xb1\x32\x02\x47\x7d\xc3\xbb\x1e\x66\x35\x23\x3f\xf9\x78\x48\x93\x70\x0d\x7f\x7b\xea\x58\xfb\x89\x8b\x04\xf6\x52\xc5\x9c\x02\xa6\xa2\x14\x8d\xd9\x54\xfc\xe9\xbb\x1a\xf1\x0d\xec\x50\x23\xe4\x6c\x8b\x62\xb8\x36\x75\xa3\x6c\xa0\x73\x79\xe1\x30\x8b\x9d\xa6\x41\xeb\x5e\x85\x43\xa5\x7a\xef\x72\x26\xa3\x61\x6a\x19\x66\xea\x5d\xda\x91\x15\x3a\x8f\xe5\xc0\xd2\x43\x5f\x12\x1a\xb9\x7e\xa2\x18\xec\xb9\x7e\xa0\xce\xc7\xcd\xd3\x9f\x4e\xe6\x45\xdd\xe3\xd6\xbb\x0c\x0a\x3b\xd4\xe8\xfd\x40\x5d\x07\x29\x0e\x16\x85\xc8\xbb\xb9\x9a\x4f\xb4\x4c\x03\xab\xeb\x4c\xd5\xcd\x71\xc3\xb9\x6c\xd0\xc3\x43\xc0\x0e\x7a\x27\x5a\x1d\x40\xe8\xca\x4b\x14\x76\xb8\x5e\x9d\x58\xde\xb4\x92\x61\xbd\x1f\x6f\x7d\xdc\xea\xa5\xde\xe1\x54\xde\xa9\xa2\xde\x5b\x8d\x4a\x1a\x66\xa5\xee\x14\x4f\x37\xda\xc9\xda\xc7\x0b\xc5\x89\x70\xfa\xa9\xbb\x86\x68\x8c\x11\x13\x3e\x0c\x8f\x05\x1a\x97\x4f\xf7\xa4\xe4\x70\xee\xa7\x3e\x13\x5e\xa1\x99\xf8\x11\x2a\x85\x45\x61\x63\xfb\x59\x4b\xef\x0d\x62\xaa\xfd\x7f\x00\x00\x00\xff\xff\x83\xe3\xe9\xe7\xb8\x18\x00\x00")

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

	info := bindataFileInfo{name: "api.proto", size: 6328, mode: os.FileMode(420), modTime: time.Unix(1525750671, 0)}
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
