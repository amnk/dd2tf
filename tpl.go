// Code generated by go-bindata.
// sources:
// tmpl/monitor.tmpl
// tmpl/screenboard.tmpl
// tmpl/timeboard.tmpl
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

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xdf\x6e\xdb\x20\x14\x87\xef\xfd\x14\x47\xbe\xde\xf2\x04\xed\xc5\xd4\x66\x6a\xa4\x2d\xd1\xaa\x48\xbd\x98\x26\x84\xcc\x89\x8d\x42\x20\xc5\x38\x56\xc4\x78\xf7\x89\x3f\x76\x6c\x97\x74\x77\xe6\x3b\x3f\x3e\x30\xc7\x58\x63\xab\x3a\x5d\x21\x94\x8c\x1a\xca\x54\x4d\x4e\x4a\x72\xa3\x74\x09\x25\xb1\x16\x56\x1b\x06\xce\x95\x60\x0b\x00\x49\x4f\x08\x8f\x50\x7a\xbc\xf5\xcf\xce\x95\x05\x80\xb9\x9e\x47\xbc\xf7\xcf\x11\xbf\x77\xa8\xaf\xf0\x08\x1e\xff\x0a\xcf\x7f\x01\xdb\x8a\x9e\xf1\xa9\xa1\x9a\x56\x06\x75\x0b\xce\x15\x00\xd6\x7e\x85\x9e\x9b\x06\x56\xbb\xb3\xe1\x4a\x4e\x30\x3f\xc0\x6a\xab\x0c\x3f\x5c\xb7\xea\x99\x1a\x1a\x2b\x32\x10\x22\x15\xf1\x9b\x06\x80\xb4\xcc\xc7\xa4\x77\xa0\x64\x73\xe1\x0f\x55\x1d\x31\x31\x11\x9f\xe3\xfc\x69\x21\x3b\xf3\x15\xe3\xd2\x1b\x69\x50\x5f\xa8\x88\x55\x9d\x28\xe1\x03\x8e\xba\x7c\x3a\x2b\x8e\x3b\xff\xd6\x31\x6e\x66\xaf\x48\x03\x99\xbe\xdd\x24\x73\xc7\xe4\xdf\x7e\xcf\x4f\x78\xd0\xb1\x47\xc1\x16\x4e\x8a\x98\x11\x0f\xca\x4c\x38\xab\xf5\x19\xd5\x99\x97\x48\x4d\x1c\x91\x26\x79\xe6\xd5\xac\x60\x23\x2b\xd1\x31\xdc\xd3\x3a\xb5\x97\x47\x40\x8c\x27\x51\x33\xcd\x7c\xda\x85\xf7\x8e\x6b\xfc\xde\x09\xf1\xc6\x25\x53\xfd\xd0\x86\x80\xc9\xa1\x13\x82\xf4\xb1\x30\x34\x22\x3b\x21\x7f\x7e\xd8\xbf\xa8\xd6\x3c\xa3\xa0\xd7\x74\x78\xd8\x93\x46\xb5\x86\xb0\xc0\xd2\xc9\x7d\x88\x65\x6d\xeb\x0b\x15\x1d\xf5\x1f\xf5\x24\x89\x23\x9c\x29\xb3\xd9\x7c\x33\x1a\x8d\x6d\xa3\x04\xf3\x47\x39\xb9\x3e\x33\xee\xdb\x74\x1b\x3f\x86\x1b\x7c\x53\xec\x8e\x31\x02\xa0\x8e\x69\xfd\x1b\x9a\x2d\x7a\x9b\xf3\x46\xb5\xe4\xb2\x1e\x70\x9f\x86\x71\xf6\xa2\xf8\xb9\xe2\x15\x2b\x75\xf1\xbf\x84\xb9\x8a\xe8\x81\xcf\x9c\xcb\xf4\x1d\xf7\x93\xe6\x86\x57\xc3\x35\x03\xa8\x86\x71\x94\x2d\xcb\xff\xb1\x2c\x17\x1d\x6c\xcb\x3d\xde\xcb\xcf\xf4\xd9\x4e\xde\x1f\x84\x1e\x8f\x17\x25\x5d\x90\xdf\xd6\x6a\x2a\x6b\x8c\x25\xe7\x4a\x6b\x57\xce\x95\x5f\xac\x45\xc9\x9c\xfb\xb3\x14\x9d\xb0\x6d\x69\xed\x6f\xfa\xc3\xc3\x7a\xb7\x2f\xfc\x66\x7f\x26\xe6\x5c\xe1\x11\x84\x1f\xb2\x88\x9f\x62\x2e\x9f\xfe\xc7\xab\xf5\x18\x5b\x18\x5c\x51\xfc\x0b\x00\x00\xff\xff\x2f\x6d\x59\x4e\x42\x06\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1602, mode: os.FileMode(420), modTime: time.Unix(1597429847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\xdb\x72\xdb\x38\x12\x7d\xf7\x57\xa0\xf4\xbc\x71\xbe\x20\x0f\x8e\x9d\x8b\xab\xec\x8d\xd7\x72\x9c\xdd\x9d\x9a\x62\xc1\x64\x8b\x42\x99\x24\x14\x00\x94\x2d\x6b\xf8\xef\x53\xb8\x37\x28\xde\x32\x35\xe3\x17\x0b\xe7\x34\x4e\xe3\xd2\x6c\xdc\x04\x48\xde\x8a\x1c\xc8\xaa\xa0\x8a\x16\xbc\xcc\x64\x2e\x00\x9a\x27\x4e\x45\xb1\x22\xab\xe3\x91\x9c\x5f\x17\xa4\xeb\x56\xe4\x78\x46\x88\x62\xaa\x02\xf2\xc1\xe2\x0f\xa6\xd0\x75\xab\x33\x42\x8e\xc7\x77\x84\x6d\xc8\xf9\x3d\xd0\xe2\x5b\x53\x1d\x48\xd7\x9d\x11\x22\x80\x16\x19\xd7\xc5\x0f\x44\xd7\x48\x59\x5d\x07\x9a\x22\x16\xb4\xc0\x7a\x4b\x05\x38\x4c\xda\xdf\xb6\x2e\x26\x06\x6b\x3e\x40\xbd\xab\xa8\x82\x47\x2a\x18\x7d\xaa\x40\x46\xfa\x85\xa9\xed\xa4\x81\xa0\x4d\x09\xe4\xdc\x02\xca\xd9\x65\x7b\x67\x68\xba\x4e\x48\x43\x6b\xd0\xff\x5d\xf7\xff\xad\x8b\xb6\xf7\x84\xec\x04\x6c\xd8\x6b\xe0\xee\x6c\xd1\xb3\x05\x6c\x68\x5b\x29\xcf\x5e\xb9\xa2\x1e\xd5\x33\x42\x06\x7b\x34\x54\x48\x7a\xf3\x83\x15\x25\xa8\xb1\x3e\xbc\x18\xd6\x35\x5c\x1d\x76\x71\xd2\xf4\x6f\xdf\xae\x57\x8f\xfe\x37\x40\x07\x0f\xfd\x2f\x40\x61\x80\xdd\x7c\x5b\xcd\x93\x48\x78\x80\x57\x95\xd4\x09\x5d\xe8\x49\x5c\x54\xac\x6c\x12\x9d\x8c\x1a\x08\xab\x79\xa3\x79\xb9\x35\x7b\x4b\x5b\x95\x49\x8d\x60\x31\x67\x32\xa9\xf5\x15\x58\xb9\x55\x1e\xdd\xda\x92\x13\x09\xdc\xa4\xc2\x0f\x56\xa8\xad\x07\x5f\x4c\xc1\xd5\xf7\xcc\x74\x67\xec\xf8\xd9\x7e\xe8\xdf\xbe\x07\x0b\xc6\xf5\x92\x57\x5c\x78\x30\x37\x05\x57\xdb\x33\x33\xe3\x58\x03\xc6\xdc\xf7\x82\x50\xa5\x7f\xdb\x68\x22\xa4\x62\x7b\xc8\xe4\x8e\x86\x19\xbb\x61\x7b\x58\xeb\xb2\xf7\xd3\x8d\x79\x1b\x74\x5e\xc1\x15\x6c\x86\xfc\x27\x84\x62\x15\x64\x05\x6c\x42\x33\xbc\xc0\x23\x7b\xf3\x36\x84\xec\xd9\x9b\x6f\x95\xc5\x57\xc8\x1a\x79\x47\x63\xd7\x4a\xc5\xeb\xef\x0d\x53\x91\xcb\x0d\x96\xb5\x1a\xf4\x03\x89\xcd\x66\x45\x2f\x5a\xc5\x65\x4e\xe3\xf7\x42\x08\x0d\x90\x53\xc4\x36\xb3\x82\x3a\x0a\x92\x0f\xc7\x06\x49\xef\xcb\x41\x46\xb3\x8a\xf7\xf0\xb3\x05\xe9\x33\x48\x3a\xf4\xc3\x5c\x92\x61\xf4\x9f\xb0\x66\x61\x46\xa2\xf8\x7f\x5a\x10\x87\x68\x48\xc8\x4f\xdf\x46\xcf\xac\x92\x3a\x49\x0b\x51\xaf\x6d\xb6\x0a\xf8\x68\x26\x9b\xd1\x31\x4e\xfb\x62\x3f\x35\x98\x61\x49\x6c\xb6\x48\xf7\x16\x94\x60\x39\x66\x6a\x8b\x38\xc1\xc0\x2f\xeb\x2d\xbc\xaa\xcf\xac\x52\x20\x92\x3e\xeb\x69\xde\x58\x18\xcd\x73\x30\x5c\x24\x7d\xc3\x6a\x1c\xdf\xfa\x1b\xae\x63\x6c\x7b\x76\x91\xd4\x45\x59\x0a\x28\xa9\xe2\x49\x2b\x69\x44\x7d\x78\x63\xbb\x45\xca\x97\xbc\xde\x51\x01\x0f\x1c\x93\xb9\x05\x33\xc5\x63\x4a\x8b\x66\xcb\x74\xb7\x3a\x70\xfb\xd3\x9f\x1b\x34\x99\xff\xc4\x70\x91\xf4\x37\x51\x80\xf8\x98\x44\x3a\xd7\x50\xf6\x14\x16\xd2\x68\xb2\x5c\xf1\x8a\x89\x53\xc9\x82\x89\x44\xd3\x1a\x2d\x12\xfd\xf4\xaa\x04\xbd\xe4\x15\xe6\x40\x63\x59\xce\x2b\x2f\x8a\x8c\x16\x89\x5e\x37\xb9\x00\x2a\xe1\x0b\xe7\x09\xcf\x1c\x9e\x95\x9a\x70\xe2\x3d\xe3\x85\xf1\xd0\x14\x4c\x31\xde\xd0\xea\x33\x17\x35\xc5\xf9\x08\x67\xab\x01\xbb\x77\x3d\xc3\x93\xd4\xa5\x03\x2b\xd4\xca\x36\xa6\x1a\xca\x63\x23\x4b\xab\xaf\x39\xba\xc0\x8e\xf6\x29\x0a\xde\xd1\x0a\x94\x82\x94\xdc\x39\xd0\xef\x1f\x83\xcd\x62\x59\xfb\x55\xf4\x3f\x4b\xff\xfd\xe0\x0f\x33\xb1\x5c\xac\x7f\xdd\xec\x41\xa8\x94\x63\x16\x0b\x53\xec\x2c\x16\x6b\x3e\xd2\xaa\xed\x0d\xc4\xde\x40\x7e\x05\x77\xfc\xf2\x46\xd6\xb4\x84\xef\xf7\x37\xbd\x66\x6a\x34\x6b\x45\x88\x74\x64\x36\x23\xdd\x0d\xc4\x29\x99\x8b\xdd\x91\x70\x5e\xab\x43\x05\xc3\x11\x7c\x42\x49\x03\x0c\xc5\xe3\x3f\x14\x3e\xfd\xec\x38\xb3\xd6\xce\xa8\x25\x5b\x61\xfb\x37\xb1\x21\x1e\x1f\x7d\xf2\xfe\xbd\x19\x89\xbf\x38\xe2\x0f\xb4\xb4\xcb\x64\x92\x38\x14\x2d\xdd\x7a\x2a\xc9\x07\xf2\xdb\xf1\xe8\x92\x43\xb4\xee\xba\xd5\xf1\x78\xde\x75\xab\x7f\x1d\x8f\xd0\x14\x5d\xf7\xfb\xb8\x37\xdd\x46\xb7\x0f\x9a\xda\x71\xcd\x03\x26\x4f\xef\xa1\x49\xb2\x57\x0c\x91\x41\xea\x24\xad\x81\xb6\x42\x71\x33\xba\xed\xd2\xcd\x36\xc6\xbd\x26\xbd\x3b\x69\xe4\x30\xd2\x6b\xf8\x2d\x15\xcf\xc9\x30\xc7\x86\x0f\x52\x27\x0d\xaf\x8d\xd5\xc0\x4e\xf2\xef\xda\x02\xde\xd0\x27\x48\xd6\xbf\xca\x00\x7e\x07\xe4\xd8\x45\x52\x27\x69\x6b\x26\x69\x0d\xc6\xcc\xd0\xe0\x2f\x8e\x18\x2d\xe0\xcf\x43\x67\x03\xfc\xc4\x11\x33\x39\x2e\xeb\xad\x65\x72\x5a\x8e\x06\x93\x07\x46\x7c\x5a\xc2\x87\xa4\xe4\x78\x34\x52\xf7\x4e\x40\xce\x24\xe3\xe1\x24\xb3\x0b\x40\xbc\x3d\x09\x16\xb3\x47\xe6\xf4\x36\x61\xc1\x91\x68\x44\xeb\x33\x6f\x92\xb1\xd9\xf0\x26\x1d\x1b\x64\x30\x29\x74\x51\x81\x50\xd7\x57\x1e\xa6\xba\x98\xb1\xb0\x11\x8a\xf4\xb4\x4a\xab\xf8\x3d\x6c\x04\xc8\x90\x47\xf5\x01\x32\x13\x0e\x43\x67\xc8\x68\x36\xa9\x78\x03\x25\x42\x2b\x5b\xf2\xe1\xef\xb9\x05\x0a\x78\x90\xac\x4a\x32\x4c\x89\xd1\xa4\x5c\x72\x46\x34\x87\xb1\xe1\x64\x35\x16\x83\x71\x99\x47\x4b\x3b\x5e\xd5\xc7\x2e\x4b\xb6\x90\x3f\x87\xcb\x12\x53\x08\xfb\x7f\xcb\x4c\x56\xff\x22\x78\xbb\x63\x4d\xe9\xf1\xd2\x97\x9d\x08\xe2\xe7\x75\x12\x91\x44\x61\xfe\xce\x26\x7f\xbe\xe3\x32\xde\x8e\xe4\xcf\xd9\x8e\xcb\x78\xef\xe5\xe9\x59\x95\x4f\x45\x09\x89\x0c\x68\x00\xe9\x38\x83\xe9\xeb\xb3\x87\xdb\x30\x1d\x5b\x55\x87\xf9\x70\xf8\x6c\x23\x70\x03\xb0\xef\xb9\xba\x1f\xcb\x1c\x6f\xd0\x9f\xca\x64\x73\x1e\xd9\x49\x11\xb3\xb6\xe2\xb8\x36\xc9\x39\x09\x6b\x6c\x32\xa9\xb5\x66\x6f\x28\x38\xa4\x2d\x39\x91\xc0\x4d\x2a\xdc\x52\x51\xb2\x90\xd2\x6a\x5b\xf2\x17\x0a\x9e\x9b\xee\x4f\xb3\x0f\x3d\x69\xf6\xa1\x0b\x06\x9d\x6e\x3c\x88\x3d\xcb\xc1\xfd\x0b\x9d\xb0\xc5\xcc\xfd\x0f\xbd\xe9\x1b\x2f\x91\x76\xb7\xe6\x89\xae\xb9\x58\x4f\x45\x93\xcb\xf5\xf1\x91\x86\x47\x10\x78\x29\xd1\x33\x96\xed\x1d\x16\x07\x1d\x99\x4d\x27\x38\x7a\xe0\xad\xea\x69\x56\x06\xec\xab\xf6\x4d\xa7\x67\xb4\x95\x6a\xbd\xe5\x2f\x5f\x59\x3c\xc2\xd6\xad\x54\x99\xdc\xf2\x97\x6c\xab\x51\x3f\xbf\xa9\xe5\x22\xd5\x4f\x42\x70\x31\xa0\x0b\x16\xef\x29\x07\xeb\x45\xda\x37\x54\x41\x93\x1f\x4e\xc5\x2b\x47\xf4\xd4\xa3\xfd\x22\xf9\x8f\x02\xe8\x73\xc1\x5f\x9a\x53\x07\x4f\x81\xea\xb9\xc0\x75\x16\x39\xb9\x62\x52\x09\xf6\xd4\x2a\x34\xa9\xd1\x4f\x81\xd9\x9e\xab\x5e\xcd\x45\xde\xee\xdd\x8b\xda\x0d\x93\xea\xd4\x9b\x7f\x6f\xcb\x2a\x4d\xf7\xdc\xf5\xaa\x4e\xba\xbb\x62\x72\x57\xd1\x83\xbd\xef\xf0\x64\x61\x41\x7f\x9d\xe1\x9f\x9b\x7a\xa6\xf3\x2f\x09\x77\x02\x36\x20\xa0\x89\x29\xc0\xa4\xd1\x6c\x17\x71\x7c\xfb\x91\x98\x4f\xaf\x12\xac\x80\xff\x83\xe0\x97\xbc\x6d\xe2\xa7\xb0\x65\x05\x64\x6f\x20\x78\x96\x5b\xdc\xaf\x1d\x7d\xeb\x99\xb4\xd9\xd0\x12\xd6\x8a\xaa\x56\xea\xd1\x4c\x9e\xab\x6a\x43\x66\xd2\xb0\x76\x2a\x92\x17\xac\xb1\xca\x8b\x3d\xe2\x27\xb0\x01\x8f\xf6\x61\x0a\x3f\xeb\x8c\x55\xfe\x35\x8f\x78\xd1\x1a\xf2\x88\x97\xb0\xb1\xca\xbf\xe6\x31\xd9\x6f\x0f\xb9\x4c\xb6\xdf\xa3\xd5\xa7\x8f\x08\x54\xd0\x5a\x62\xd4\x1e\x25\x53\x7c\x67\x4b\xfd\xe7\xa0\x35\xc7\xd7\x54\x92\xc7\x0b\x2a\xc7\x2c\x7a\x6b\x89\xe8\xe8\x5b\xdc\x84\x82\x89\xd8\x08\x9b\xb0\x8e\x9f\x8c\xe5\x66\x45\xd6\x8a\x26\x3d\x31\x45\xdf\x15\xc7\x8d\x88\x8c\x9e\x02\x87\xbf\xf8\xb6\x6e\x24\xfa\xd2\x4d\xd1\xbe\xba\x7b\xf2\x0f\x02\x32\xa7\x3b\xb8\xdc\x52\x41\x73\x7c\xa3\x32\xb6\x94\xf2\x52\x42\x68\x7d\x65\x4b\x7e\xed\xf4\xdc\x40\x10\xe8\x93\xad\x7d\xbc\x9e\x78\x08\xef\xce\xfe\x0c\x00\x00\xff\xff\xd2\xc2\xd5\xc0\xb8\x20\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8376, mode: os.FileMode(420), modTime: time.Unix(1597338736, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcd\x6e\xdb\x38\x10\xbe\xfb\x29\x06\x42\x0e\xbb\x40\xec\x07\x58\x20\x87\x6c\x02\x07\x0b\x6c\xdb\x34\x09\xd2\x43\x51\x08\xb4\x34\x52\x88\x52\x3f\xa1\xa8\x24\x0e\xc1\x77\x2f\x38\xfc\x93\x6c\xc5\x3d\xd4\x27\xce\x37\xdf\xfc\x72\x86\x96\xc4\xa1\x1b\x65\x81\x90\x95\x4c\xb1\xb2\xab\x73\xc5\x1b\xdc\x75\x4c\x96\x19\x64\x5a\xc3\xe6\xbf\x12\x8c\xc9\x40\xaf\x00\x14\x57\x02\xc1\xfd\x2e\x9c\xf6\x81\x20\x63\xb2\x15\x40\x89\x43\x21\x79\xaf\x78\xd7\x06\xf5\xf5\x04\x72\x24\x89\xac\xcc\xbb\x56\xec\xc9\x87\xe5\xdc\x21\x2b\xbf\x58\x60\x6d\xcc\x0a\x40\xeb\x57\xae\x9e\x60\x73\x23\x59\xff\x34\x44\x50\xb2\xb6\x46\xd8\x00\x89\xb5\xd5\x51\x4a\x21\xa9\x85\x74\x92\xab\x6b\xac\x78\xcb\x29\x09\xe7\x0e\xe0\x85\xbf\xa7\x22\x1e\xf9\xbb\x55\x04\xa3\x35\xf0\x0a\x36\x97\xa3\xea\x86\x82\x09\xb4\x2a\x16\x05\x6f\x92\xb4\xc6\x64\xd6\x04\xdb\xd2\xbb\x0e\x0e\x6e\x25\x16\x7c\xf0\x41\xfb\x28\x78\x07\x49\xfb\xa1\x83\x1b\xd9\x8d\x3d\x75\xa0\xb6\x27\xb8\x80\xef\x5a\x9f\xd5\x0e\xfd\xe7\x22\x10\x8c\x09\xdd\x39\xe3\x6d\x89\x6f\xe7\x70\x86\x02\x9b\x03\x06\xaf\xbc\xda\x98\x73\xad\x29\x58\xa6\x35\x31\xe9\x44\xc8\x8f\xe5\x44\xee\x8b\xae\x47\x4a\x64\xb0\x27\x9f\xc8\xe0\x50\x1b\xc6\x11\x4e\x25\x92\x18\x7f\x94\x88\xda\xbb\xfb\x20\x74\x20\xc9\x4d\x01\x40\xcf\x04\x2a\x85\xb3\xe9\x24\xfe\xe6\xd6\x6b\xc2\x0d\x47\x6e\x5e\x09\xde\x2f\x72\xb7\x56\x11\xf8\x66\x39\x99\x4f\x4c\xfe\x44\x49\x6d\xb1\x90\x1b\xb5\x19\xe8\x67\xd6\xdb\x35\xa4\x8a\xf9\xaa\x7d\x8f\x69\x8d\xac\x90\xf2\x7b\x61\x62\x8c\xa3\xf6\x48\x42\x52\x6a\x4d\xe1\xff\x67\x3b\x14\x36\x8e\xa0\x83\x27\x3b\xf4\x68\xa4\x52\x09\xc7\x87\x58\x94\xab\xe0\x0e\x9f\x47\x1c\xd4\x62\x09\xd2\xe9\x62\x0d\xcf\x93\x5e\x7f\x1d\x51\xee\xd3\x1a\xc5\x3c\xa9\xb4\xb5\x31\x54\xef\x41\xb9\x5a\xdb\x14\xc0\x7b\x8f\x26\x97\x75\x2d\xb1\x66\xaa\x93\x2e\x09\x0b\xb6\x08\x59\x06\x7f\x5d\xe3\x1d\x56\xf7\x4a\xf2\xb6\x9e\xf2\xfe\xa6\x25\x4d\x66\x61\x4b\x13\x12\xa3\x91\x43\x1b\x75\x7d\x18\x36\x4e\x97\xd6\xbe\x17\x0e\x89\xbc\xf9\xc0\x59\x3b\xda\x73\x3f\x5e\x76\xcb\xfd\x31\xec\x78\x1a\xbc\x79\xb7\x93\xf1\x37\x5e\xaa\x27\x6b\xfa\x4a\x07\x6f\xe8\xd0\x13\x66\xa7\x9b\x3a\xb7\x99\x54\xe9\xcb\x0e\x8d\x88\x0a\x37\xd0\x57\x5d\x5b\xd2\x2b\xc9\xc4\xb6\x93\x0d\x53\x03\x4c\x47\xfb\x43\x75\x78\x9a\x53\x43\x8b\x44\xcd\x2b\xe2\xce\xba\x06\xd3\xb6\x9d\xee\xda\x7c\x3e\xac\xe7\xa6\x67\x72\x7a\xc7\x57\x09\x49\xb3\x97\x8a\x0a\xdb\xb3\xbc\x52\xc7\x01\x62\x33\xc6\x41\x75\xcd\xbf\xf5\x55\x27\xc8\x73\x41\x72\xbe\xab\xf3\x82\x90\x10\xfd\x80\xf6\x5b\x8f\xdb\x43\x8f\xd5\xa2\xc7\xed\xc7\x1e\xa7\xd7\xe6\x14\x4b\xa7\x55\xa2\x1e\x6f\xfa\xb2\x64\xfc\xe3\x3b\xc5\xd2\xf5\x3f\x60\xd3\x0b\xa6\xf0\x91\x49\xce\x76\x02\xe3\xc3\x37\xf9\x67\xb6\xdf\x08\x9e\x96\xbf\x78\x9e\xbf\xf9\x96\x35\x38\x79\x2d\x3e\x5b\x31\xdc\x57\x2f\xb1\xe2\x6f\x30\xf9\x6b\xb4\x62\xd0\x96\x58\xb1\x51\xa8\xf4\x49\xe1\x44\xfb\x51\xb2\x98\xb4\x59\xfd\x0a\x00\x00\xff\xff\xe4\xf8\xa0\x0f\xd5\x08\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 2261, mode: os.FileMode(420), modTime: time.Unix(1597338736, 0)}
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
	"tmpl/monitor.tmpl": tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"monitor.tmpl": &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
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

