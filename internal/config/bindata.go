// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package config generated by go-bindata.// sources:
// configs/config.ini
// configs/migrations/01_initial.up.sql
package config

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xef\x8f\xda\x46\x10\xfd\xbe\x7f\xc5\x7c\xcc\x55\x8a\x43\x9a\xb6\x8a\x8a\xf8\x60\xc0\x21\xd6\x19\x1b\x61\xee\xfa\xe3\x54\x59\x8b\x3d\xd8\x2b\x96\x1d\x67\x67\x0d\xa5\x7f\x7d\xb5\x36\x5c\x49\x94\xa4\x07\x5f\xac\xf1\x9b\xf7\xc6\xfb\xde\xec\x13\xa3\x3d\xa2\xfd\x4b\x8c\x21\xac\x2a\x8b\xcc\x40\x06\x4e\x8d\x2a\x1b\x70\x0d\xc2\xf0\x1a\xb4\x62\x87\x86\x03\x91\xc4\xf9\x26\x4a\x8b\x70\x3e\x5f\x47\x79\x0e\x13\x18\x05\xfd\x5f\x88\x31\xac\xc8\xba\xef\x77\xaf\xb2\xf5\x06\x26\xf0\x7e\xf4\x7e\x24\xc4\x53\x25\x9d\xdc\x4a\x46\x2f\x3e\xbf\x3c\x03\xa3\x73\xca\xd4\x1c\xc0\x07\xb2\x60\xe8\x04\x64\xf4\x19\x5a\x62\x57\x5b\xe4\x4f\x1a\x14\x03\x77\x6d\x4b\xd6\x61\x15\x88\x8f\x59\xee\x29\x35\x95\x52\x37\xc4\xee\xaa\xf1\xf3\x4f\xef\x7e\x14\x0f\x79\xb4\x86\xc9\x73\xb3\x58\x85\x79\xfe\x5b\xb6\x9e\xdf\xd6\xe6\xd3\x22\x0d\x97\xd1\x6d\x49\x8c\x21\xcf\x13\x38\x50\x85\xe0\x08\xb6\x08\x1d\x63\x05\x27\xe5\x9a\x9b\x41\x02\x78\x94\x5a\x55\x3d\x8c\x41\x5a\xfc\x55\x8c\xe1\x07\xa8\x14\xcb\xad\x46\x78\x0d\x29\x79\x9a\xbe\x68\xf1\x53\xa7\xac\x2f\x86\xfa\x24\xcf\xdc\xf3\xbf\xe2\xbd\x6a\xe1\x88\x56\xed\x54\x29\x9d\x22\x73\xd7\x83\xfb\xca\xf9\x75\x29\xbf\x80\x0f\x75\x70\x8d\x74\xfd\xe9\x96\x68\xdd\xd0\x8a\xd0\x5a\x64\x34\x0e\x2b\xd8\x7a\x04\x8a\x31\xc0\xf5\xf8\x4f\x92\x81\x55\x6d\x86\x97\x12\x9c\xed\xd8\x23\x67\xe1\x67\x7a\xbb\x4e\xeb\x17\x2b\x2a\x32\x9f\x69\xf6\x7a\x37\x96\x7f\x5b\x13\xa4\xa9\x6e\x91\xde\x34\x30\xf2\x30\x8c\x7c\x90\xae\x6c\x90\x7b\x00\x19\x04\x65\xbe\xfc\xd4\x3b\x91\xe7\x49\xb1\xcc\xe6\xde\xb2\xcb\x61\x0b\xf1\xa4\xa9\x66\x1f\xa4\xa5\x32\xea\xd0\x1d\x40\x53\x0d\x1a\x8f\xa8\x03\x91\x64\x8b\x22\x89\x1e\xa3\xc4\xc7\xb5\x37\xb7\xa1\x4e\x57\x1e\xc2\xde\xdc\x93\x55\xce\xa1\xf1\x56\xef\x94\xc6\x40\x7c\x88\x93\xa8\x48\xb2\xc5\x22\x4e\x17\x30\xf1\xb3\xe3\xf7\xdb\x4a\x32\x4c\x1a\xe1\x15\xbb\x8a\x3a\x77\x17\x88\x59\x96\xe6\xd9\x57\x59\x56\xd2\x35\x57\x29\x50\xc6\xd1\xcd\xbe\xf4\xd4\xd2\x3e\x73\x0f\xb3\xaf\xc2\xcd\xc7\x3e\xe2\x35\xbf\xe9\x18\x2d\x07\x9a\x6a\x4f\x35\xf9\xd6\xcf\x0f\x7b\xd9\x23\xd8\x91\x1d\xb4\x34\xd5\xb5\x32\xf5\xff\xf4\x2d\xe5\xdf\xc0\xea\x1f\x04\xda\x5d\x47\x1a\xda\xb7\xb8\x23\x8b\xa0\x1c\x83\x25\x27\xbd\x99\xca\xc0\x72\x1a\x88\x65\xf8\x7b\x91\xc7\x7f\x7a\x3f\xde\x8e\x2e\x14\xb2\x7e\x29\x43\x25\xcf\x3c\x70\x84\x8b\x5b\x8a\xde\x45\xd3\x1d\xb6\x68\x3d\xd3\xb5\xe3\xca\xc6\xfe\x0c\xf7\x88\xed\xd0\x3b\x0d\x67\xf7\x0f\xab\x7c\xe8\xef\xd3\xa0\x8c\x8f\x43\xfa\x4c\xc0\x58\x92\xa9\x78\x08\xb3\x2c\x4b\x7f\xd5\x39\xda\xa3\xf1\x17\xca\xd1\x6f\x72\xe0\xef\x21\xdc\xc9\x4e\x3b\x5f\x7b\x0b\x0d\x75\x36\x10\x9b\xec\x3e\x4a\x8b\xcd\xc6\xc7\xe7\xdd\x2f\xa3\x3e\x41\x09\x5d\x76\xc0\x13\xab\xda\x28\x53\xc3\x1e\xfd\x87\xe4\xf1\x22\x8d\xd3\x45\xb1\x5a\xc7\x8f\xe1\x26\x2a\xee\xa3\x3f\x8a\x24\x9b\x85\x9b\x38\x4b\x61\xe2\x93\xb2\x7b\xd3\x5a\x75\xdc\xe3\x39\x68\xf1\xf0\x1f\xfe\x61\x9a\xc4\xb3\xaf\xc3\xbb\xed\x15\xfd\x6f\x00\x00\x00\xff\xff\xc7\x8b\x24\xf3\xaf\x05\x00\x00")

func configIniBytes() ([]byte, error) {
	return bindataRead(
		_configIni,
		"config.ini",
	)
}

func configIni() (*asset, error) {
	bytes, err := configIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.ini", size: 1455, mode: os.FileMode(420), modTime: time.Unix(1604775496, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4f\x4f\x83\x40\x10\xc5\xef\x7c\x8a\x77\x6c\x93\xe2\x17\xf0\x84\x74\x1a\x89\x94\xd6\xed\x12\xad\x17\xb2\xb2\xd3\xb0\xb1\x85\x66\x17\x34\x7e\x7b\x03\xfd\x67\xab\xc6\x34\x61\x6f\x93\x0c\x8f\xf7\xde\x2f\x13\x0a\x0a\x24\x41\x06\x77\x31\xa1\x71\x6c\x9d\x37\xf0\x00\xa0\x31\x1a\xc7\x27\xe9\x59\x62\x2e\xa2\x69\x20\x96\x78\xa0\xe5\xc8\xeb\x76\x78\xa3\xcc\xfa\x6c\xe7\xba\x97\xcc\x24\x92\x34\x8e\x91\x26\xd1\x63\x4a\xa3\x4e\xb4\x54\x1b\x46\x0f\xa2\x7b\x8f\x5b\xe5\xdc\x47\x65\x75\x4f\x72\x5a\xd5\x9c\xe5\x96\x55\xcd\x1a\x32\x9a\xd2\x42\x06\xd3\x39\x9e\x22\x79\xdf\x8d\x78\x99\x25\x84\x31\x4d\x82\x34\x96\x08\x53\x21\x28\x91\xd9\x69\xf1\x20\xe7\x0d\x6f\x3d\xcf\xf7\x21\x78\x65\xd9\x15\xa8\xab\x37\x2e\x1d\x7c\xdf\x3b\x23\xa2\x9a\xba\xe0\xb2\x36\x79\xfb\xbf\x4c\xf3\xbb\xc9\xf9\x48\xc8\xb1\xcd\x0e\x94\xba\x64\x82\x26\x24\x28\x09\x69\xb1\x43\x89\x41\x63\xf4\x10\xb3\x04\x63\x8a\x49\x12\xc2\x60\x11\x06\x63\xba\x4c\xd6\x4e\xbe\x0f\xd9\x7a\x80\x2b\xaa\x66\xad\xf1\xca\x28\x94\x2b\x58\x63\x55\x59\x28\xad\x59\xc3\x71\xde\x58\x53\x7f\xde\xb4\x3e\xdb\x8f\x3a\xd7\xfd\xb1\xea\xab\xdc\x5d\xa2\xb5\x72\x75\xd6\x38\xee\x1a\xfa\x53\xed\x7f\x73\x1d\xaa\x1f\x77\x92\x6d\xd9\x6e\x8c\x73\xa6\x2a\x7f\x01\x72\x1d\x8e\x73\xdf\x27\xe1\x2b\x2b\xbd\x68\x73\x77\x54\x18\xec\x6d\x8d\xbe\x09\x0f\xdb\x4c\x5f\x01\x00\x00\xff\xff\x0f\xfc\xa0\x7b\xfb\x03\x00\x00")

func migrations01_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialUpSql,
		"migrations/01_initial.up.sql",
	)
}

func migrations01_initialUpSql() (*asset, error) {
	bytes, err := migrations01_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 1019, mode: os.FileMode(420), modTime: time.Unix(1604771397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"config.ini":                   configIni,
	"migrations/01_initial.up.sql": migrations01_initialUpSql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"config.ini": &bintree{configIni, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.up.sql": &bintree{migrations01_initialUpSql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
