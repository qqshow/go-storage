// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/context.tmpl (594B)
// tmpl/header.tmpl (711B)
// tmpl/meta.tmpl (63B)
// tmpl/pairs.tmpl (2.015kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _contextTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\xcd\x6a\xe3\x30\x10\xbe\xfb\x29\x86\xe0\x83\x0d\x5e\xf9\xbe\x90\x53\x96\xb0\xa7\x52\x9a\x43\x7b\x9d\xca\x63\x47\x34\x1e\x09\x69\xec\x1a\x82\xdf\xbd\x48\x8e\x5b\x97\x40\x6e\xf3\xfd\xcc\xcf\x27\x5d\xaf\x7f\xc0\x23\x77\x04\x39\x63\x4f\x15\xe4\x0d\x0a\xc2\xdf\x3d\xa8\x7f\xb1\x98\xe7\x6c\x63\xf9\xa8\x20\x1f\xa3\xb8\xb8\x92\x08\x79\xcb\x89\x1a\xd5\x71\x60\x1d\xd6\x16\xd3\x26\xe5\x17\x52\xff\x31\x1c\x2c\x0b\x4d\x12\x85\xba\x86\xa5\x5f\x3d\xa3\x27\x8e\xdc\xab\x91\xf3\xea\xc0\xa6\x09\xa0\x6f\x20\x0c\xce\x59\x2f\xd0\x5a\x7f\xd7\xa4\xb2\x76\x60\x0d\xc5\x8d\x7f\x21\x4d\x66\x24\x0f\xf3\x5c\x3e\x5a\x50\xfc\x68\xd8\x87\xe4\x2e\xe0\x7b\x86\x0c\x9e\x17\xf2\x9a\x01\x00\x04\x87\x5c\x81\x96\x29\x86\xb5\x8e\x58\x3c\x6a\xc3\x9d\x3a\x09\x7a\x39\x39\xe4\xa3\xb7\xfd\x3a\x5a\xcb\x54\xc1\xae\x33\x72\x1e\xde\x95\xb6\x7d\xfd\x36\x20\x7f\xda\x3a\x88\xf5\xd8\x51\x1d\xc8\x8f\x46\x53\xa8\xe3\x3e\xf5\x84\x3d\xc5\x1c\x11\xf0\xa6\xde\x1e\xbe\x2b\xd3\x19\x0d\xb5\xe4\xd3\x31\xea\x68\xd8\x84\x73\x51\x66\x49\x70\x68\x7c\x80\x3d\xa0\x73\xc4\x4d\x91\x60\x05\x2e\xa8\x6d\x62\x2d\x53\xb9\xcc\xf1\x29\x20\x84\xbb\x3d\xeb\xab\x1c\xf0\x72\x59\xde\x30\x5b\xbe\x90\xb8\x59\x7f\xf3\x51\xf9\x15\x00\x00\xff\xff\x5b\x12\x0d\xa2\x52\x02\x00\x00")

func contextTmplBytes() ([]byte, error) {
	return bindataRead(
		_contextTmpl,
		"context.tmpl",
	)
}

func contextTmpl() (*asset, error) {
	bytes, err := contextTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "context.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3e, 0xff, 0xae, 0xbb, 0x1e, 0xec, 0xa, 0x55, 0x23, 0xfd, 0x59, 0x78, 0x59, 0xcc, 0xbf, 0xf8, 0x3e, 0x85, 0x3a, 0xd9, 0xc3, 0x32, 0x76, 0x1d, 0x57, 0xd2, 0xbc, 0x22, 0xb, 0xa5, 0xf3, 0x7a}}
	return a, nil
}

var _headerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x6e\x2a\x31\x0c\x85\xf7\x79\x0a\x6b\x56\xf7\x2e\x9a\x3c\x40\x97\x85\x45\x37\x50\x09\x16\xdd\x55\x26\xb1\xd2\x08\x26\x8e\x1c\x33\x2d\x42\xbc\x7b\xa5\x86\xa1\x3f\xaa\x34\xb3\xf2\xf1\xf1\x17\x27\x39\xce\xc1\x03\x07\x82\x48\x99\x04\x95\x02\xec\x4e\x10\xf9\xd6\xc3\x90\x10\x52\x56\x92\x8c\x07\xe7\xfb\xe0\x2a\xc9\x90\x3c\xdd\xc3\x62\x0d\xab\xf5\x16\x96\x8b\xc7\xad\x35\x05\xfd\x1e\x23\xc1\xf9\x0c\x76\x85\x3d\xc1\xe5\x62\x4c\xea\x0b\x8b\xc2\x3f\x03\x00\xd0\x79\xce\x4a\xef\xda\xb5\x2e\x71\x67\x9a\x8a\x49\x5f\x8f\x3b\xeb\xb9\x77\x5c\x28\xab\xa0\x4f\x39\x7e\xd7\x77\xf1\x2f\xf8\xf9\x88\xf9\x8d\x5d\x55\x16\x8c\xd4\x4d\xcc\x5d\xd9\x47\x57\x29\xf6\x94\x75\x16\x4b\x39\x14\x4e\x33\x61\x2f\x14\x28\x6b\xc2\xc3\xbc\x77\x34\xed\x0f\x58\xeb\xe4\x81\x6b\xe0\xd3\xa0\x9e\xca\x48\x95\x3a\x09\xba\x82\x49\x66\x2e\x75\x3d\x29\x06\x54\xec\xcc\x7f\x63\x06\x14\x78\x81\xaf\x2f\xdb\x27\xe1\x21\x05\x92\xeb\x64\x4c\xee\xb7\x7f\x4d\xdf\x6e\x5a\x1d\xdd\x76\x91\xdd\xb4\x2a\x3f\xed\xcf\x88\xec\xf6\x54\xe8\xb6\xa4\xa5\x61\x37\x4d\x2c\x45\x58\xcc\x47\x00\x00\x00\xff\xff\x22\xba\x87\x02\xc7\x02\x00\x00")

func headerTmplBytes() ([]byte, error) {
	return bindataRead(
		_headerTmpl,
		"header.tmpl",
	)
}

func headerTmpl() (*asset, error) {
	bytes, err := headerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0x4b, 0xc, 0xb8, 0x76, 0x1b, 0x35, 0x8c, 0x66, 0xc6, 0x14, 0x3, 0xa3, 0x76, 0x39, 0x28, 0x76, 0x32, 0xe2, 0x80, 0xe9, 0xf8, 0xe3, 0xd2, 0x39, 0xf8, 0xf7, 0xd8, 0xba, 0xa6, 0xee, 0x82}}
	return a, nil
}

var _metaTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\x08\xa9\x2c\x48\x55\xc8\x2c\x56\x28\xc9\x48\x55\x28\x01\xb1\xd3\xf2\x8b\x14\xaa\xab\x15\xf4\xfc\x12\x73\x53\x15\x6a\x6b\xb9\x92\xf3\xf3\x8a\x4b\x20\xca\x6c\x15\x94\x90\x64\x94\xb8\x00\x01\x00\x00\xff\xff\xfc\x49\xcc\x09\x3f\x00\x00\x00")

func metaTmplBytes() ([]byte, error) {
	return bindataRead(
		_metaTmpl,
		"meta.tmpl",
	)
}

func metaTmpl() (*asset, error) {
	bytes, err := metaTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x70, 0x58, 0x9b, 0x31, 0x99, 0x67, 0xe2, 0x68, 0xbb, 0xf9, 0x82, 0x93, 0xde, 0x50, 0xb3, 0xeb, 0x16, 0xaf, 0xb0, 0xb9, 0x15, 0x44, 0x49, 0xe6, 0xc9, 0xb4, 0xf8, 0xd5, 0xc1, 0x4d, 0xb5, 0x23}}
	return a, nil
}

var _pairsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x5d\x6b\xdb\x30\x14\x7d\xf7\xaf\xb8\x0b\x66\xd8\xc5\x75\xdf\x3b\xf2\xb2\x2f\x06\xa3\x25\x8c\x6d\x2f\x25\x14\xcd\xbe\xee\x84\x13\x49\x93\x64\x37\xc1\xd3\x7f\x1f\xfa\xb0\x63\x1b\x27\x5b\x61\x79\x8a\xa5\x73\xae\xce\x3d\xf7\xa3\xeb\x20\x7e\x4f\x34\x81\xdb\x35\xe4\x60\x4c\x14\x75\xdd\x35\x48\xc2\x9e\x10\x62\x46\xf6\x98\x41\x5c\xf6\xf7\x0e\x68\x8c\x83\xc4\x82\xd9\x33\x87\x81\xdf\xa0\xf9\x86\xa8\x82\xec\xfa\xeb\x10\xa1\xce\x20\x6e\x1d\xae\x9c\x70\x6b\x77\x56\xcf\x88\x2d\x91\x20\x08\x95\x56\x94\x60\x60\x8c\xfb\x53\x83\x31\x77\x44\xc0\x1a\xf6\x44\x3c\x28\x2d\x29\x7b\xda\x2a\x2d\x9b\x42\x77\xa6\x8b\x00\x00\x6e\x6e\x60\x23\xf1\xba\xc4\x8a\x32\x2c\x5d\x0c\xe5\x2e\x56\x05\x67\x1a\x0f\x7a\x75\x0b\x03\xc3\x64\x3d\xe7\x0e\x35\x59\x20\x8d\xf5\xe3\x31\x83\xf8\xd1\xa9\x6d\xf3\x8d\x85\x58\xa1\x2e\xb4\x15\x57\xe3\x11\x8c\x59\x88\x6e\x63\x20\x2b\x2d\xd8\x44\x91\x3e\x0a\x3c\x93\x59\x60\xc2\x5f\x12\xb1\xf1\x68\x65\x45\x7c\x6c\x58\x31\x88\x78\xe7\xd3\x83\x90\x66\x1e\xbe\xe7\x12\x5e\x9a\xb0\xc4\x5f\x0d\x95\x58\x2e\xe5\x1d\x94\x30\xae\x47\xb8\x70\xf7\x89\xa8\xde\x95\x49\x65\xe1\x07\xe7\xbb\xb9\x28\xff\xb9\x84\xf6\x17\x09\x65\x25\x1e\x7c\x7b\xe6\x5f\x8f\x02\x6d\x13\x58\x74\xea\xbe\xc6\x7a\x4e\x4e\x77\x9d\x15\xf7\x93\xa8\x8d\xc4\x8a\x1e\x9c\xc9\xab\x7b\x7c\x5e\xc1\xb5\x31\x51\xd5\xb0\x02\x04\x91\x0a\x87\x42\xd8\xdc\x5c\x53\xe2\x4e\xe1\x09\x94\x28\xb8\x1a\x30\x90\x7a\xd2\x09\xcb\x4a\x07\x1d\xaa\x98\x70\xa1\x15\xe4\x79\x7e\x65\x4b\xad\x9c\x63\x29\x24\x57\xcb\x35\xcf\x00\xa5\xe4\x32\x0d\x45\x97\xa8\x9a\x9d\xb6\x5e\xbf\x5e\xc6\x77\xa1\x82\x2d\xd9\x35\xa8\x2c\x70\x4f\x6a\x4c\x46\xe3\x40\x99\x46\x59\x91\x02\x3b\x93\x3a\x68\xc5\x25\x3c\x66\xe0\x66\xcf\x97\xd6\x29\xf4\x0f\xce\xca\x98\x2c\xd8\x95\xf6\xf6\xda\x1f\xad\x6c\x2c\xee\x86\xf6\xec\x7c\x3e\xb4\xf9\x67\x3c\x6e\xdf\xc0\x2b\x5e\x8f\xde\xf1\x09\xea\x46\x32\x60\x74\x97\x81\x42\xd9\xd2\x02\x55\x7e\x8f\xcf\xd6\xa5\x6f\x4c\x35\x42\x70\xa9\xb1\xfc\x60\x4d\x49\xda\x74\xe0\x9a\x89\xda\x51\xdb\x9c\xcc\x08\xaf\xc2\x1a\xda\xfc\xbb\x3d\x89\x3c\x31\xf8\x25\xa1\x85\x91\x37\xc3\x29\xaf\x7d\x4b\x0e\x43\x67\xcb\x0b\xe2\x25\xa3\xd7\x3a\x47\xd6\xbd\x0e\xa1\xfa\xe1\xdb\x46\xc1\xb3\x89\x0f\xbe\xc8\x3d\xc6\xe9\x4d\x66\x53\xeb\x13\x37\xbe\x13\x2f\x30\x7b\xda\x5b\x52\xd4\x4f\x92\x37\xac\x4c\x02\xf5\xdc\xdc\xfb\xf4\xf6\xff\x69\xfa\x63\x61\x47\xd6\x2d\xf0\xf9\xe8\x9e\x71\xc6\xf7\x89\xdb\x95\xdb\x89\xa7\xf3\x05\x42\xab\x59\xfb\x5c\x6a\x9d\x2f\x81\xec\xfb\x66\xfa\xcc\x19\x43\x96\x2a\x73\x69\xa1\x8d\xfc\x0f\xbb\x2d\xc4\x87\x35\x68\x19\xba\xed\x4c\x83\x06\xde\x94\xd4\xe6\xc9\x3f\xad\xb6\x4b\x05\x0d\x9e\xf8\xf8\x99\xf5\x26\x32\xd1\x08\x34\xfa\xfb\x27\x00\x00\xff\xff\xae\xd6\x77\xd5\xdf\x07\x00\x00")

func pairsTmplBytes() ([]byte, error) {
	return bindataRead(
		_pairsTmpl,
		"pairs.tmpl",
	)
}

func pairsTmpl() (*asset, error) {
	bytes, err := pairsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pairs.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x73, 0xb1, 0xf0, 0x2e, 0xcc, 0xe3, 0xe7, 0xde, 0x81, 0xcf, 0xb1, 0xd7, 0x6c, 0xc8, 0xe2, 0xa1, 0x87, 0x9, 0x9f, 0x7a, 0x62, 0xe9, 0x31, 0xe, 0x1c, 0x4f, 0x19, 0x6, 0x3f, 0x91, 0x57, 0x26}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"context.tmpl": contextTmpl,
	"header.tmpl":  headerTmpl,
	"meta.tmpl":    metaTmpl,
	"pairs.tmpl":   pairsTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"context.tmpl": &bintree{contextTmpl, map[string]*bintree{}},
	"header.tmpl":  &bintree{headerTmpl, map[string]*bintree{}},
	"meta.tmpl":    &bintree{metaTmpl, map[string]*bintree{}},
	"pairs.tmpl":   &bintree{pairsTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
