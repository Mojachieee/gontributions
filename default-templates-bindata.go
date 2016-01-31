// Code generated by go-bindata.
// sources:
// templates/default.html
// templates/detailed.html
// DO NOT EDIT!

package main

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

var _templatesDefaultHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\xc1\x0e\x82\x30\x10\x44\xcf\xf8\x15\xeb\x0f\x40\xb8\x9a\x86\x8b\x57\x63\xfc\x05\x68\x57\xa9\x81\x5d\x52\xb6\x07\x42\xf8\x77\xb7\xad\x5e\x3c\x4d\x3b\xd3\x79\x1d\x73\x76\x6c\x65\x5b\x10\x46\x99\xa7\xee\x64\x8a\x54\x66\xc4\xde\x25\x6d\x7e\x87\x81\xdd\xa6\xaa\x49\xdb\xdd\xfc\x2a\xc0\x4f\xe0\x05\x09\x56\x8e\xc1\x22\x58\x26\x09\x7e\x88\xe2\x99\x56\xad\xb5\xf9\x71\x4c\xb0\x6a\xdf\x21\xf4\xf4\x42\xa8\xe1\x38\x92\x3d\xf9\x4e\xbd\xfa\x11\xf8\x8d\x56\xea\x7b\x3f\xa3\x26\x17\x48\xe6\x95\x23\x89\xde\xfe\x89\xda\x29\x28\x24\xf7\xc5\x34\x19\x6f\x9a\xb2\x4d\x3f\xcd\xe3\x3f\x01\x00\x00\xff\xff\x50\x3e\x7e\x46\xd4\x00\x00\x00")

func templatesDefaultHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultHtml,
		"templates/default.html",
	)
}

func templatesDefaultHtml() (*asset, error) {
	bytes, err := templatesDefaultHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.html", size: 212, mode: os.FileMode(420), modTime: time.Unix(1454240192, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDetailedHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x4d\x4e\xc4\x30\x0c\x85\xd7\x70\x0a\x33\x07\x68\x34\xfb\x90\x0d\x2c\x47\x08\x21\x71\x80\xfc\xb8\x34\x68\x26\xae\x92\x74\x51\x55\xbd\x3b\x4e\x43\xab\x82\x2a\xcd\xca\xce\xcb\xf7\xe2\xf8\xc9\x27\x47\x36\x8f\x3d\x42\x97\x6f\x57\xf5\x28\x6b\x79\x90\x1d\x6a\x57\xaa\x58\x1b\x43\x6e\xe4\xca\x37\x67\x75\xf1\x29\x03\xb5\x40\x3d\x06\x48\x34\x44\x8b\x60\x29\xe4\xe8\xcd\x90\x3d\x85\xc4\xb6\x73\x81\xa7\x09\xa2\x0e\x5f\x08\x0d\xcc\x73\x3d\xfb\x16\x9a\xf7\x48\xdf\x68\x73\xf3\xf9\x71\xa9\xba\x34\x4a\x6a\xe8\x22\xb6\xcf\x27\x66\xfe\x01\x27\xb5\xd7\xde\xf4\x0d\x59\x94\x42\x2b\x29\xcc\xef\x14\xbc\x26\xdc\x9e\x3a\xa4\x37\x32\xb8\x0a\x16\x34\x0a\xb5\x74\x7b\xc7\x2b\x26\x1b\x7d\x5f\xf6\x60\xf0\xef\x5e\x07\xb6\x17\x1a\x42\xbe\x07\xae\xdd\x6e\x3e\xff\x68\x49\x94\xa3\x5a\x22\xff\x09\x00\x00\xff\xff\x45\x98\x2d\xc5\x8a\x01\x00\x00")

func templatesDetailedHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDetailedHtml,
		"templates/detailed.html",
	)
}

func templatesDetailedHtml() (*asset, error) {
	bytes, err := templatesDetailedHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/detailed.html", size: 394, mode: os.FileMode(420), modTime: time.Unix(1454240192, 0)}
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
	"templates/default.html": templatesDefaultHtml,
	"templates/detailed.html": templatesDetailedHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"default.html": &bintree{templatesDefaultHtml, map[string]*bintree{}},
		"detailed.html": &bintree{templatesDetailedHtml, map[string]*bintree{}},
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

