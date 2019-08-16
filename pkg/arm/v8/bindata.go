// Package arm Code generated by go-bindata. (@generated) DO NOT EDIT.
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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x5d\x73\xdb\xb6\xd2\xbe\x2e\x7e\xc5\x96\x72\x9b\x8f\x06\xa4\x9d\xb6\x93\x8e\x5a\x65\xc6\x71\x94\xbc\x7e\xe3\xc4\x3a\x72\x72\x3a\x67\xd2\x4c\x06\x02\x96\x14\x22\x10\x60\x01\x50\xb6\xa2\xe8\xbf\x9f\x01\x48\xc9\xb2\x24\xdb\xc9\x49\x7b\xd3\x5c\x38\x26\xb0\xd8\x5d\xec\xc7\xb3\xbb\x70\xe7\xdb\x6c\x24\x75\x36\x62\x6e\x0c\x14\x2f\x08\xe9\xc0\x33\x63\xc1\xa3\xf3\x52\x17\x5d\x50\xa6\x00\xa6\x05\x08\x6b\x2a\x60\x4a\x81\xb7\x2c\xcf\x25\x07\x3f\x66\x1e\xce\x4d\xad\x04\x58\x53\x7b\x84\xa9\x64\xe0\xc7\x08\x25\x73\x1e\x2d\xf4\x4f\x9e\x90\x0e\x0c\xfb\x67\xa7\x6f\x86\x47\xfd\xe7\xc3\xd3\x37\x83\x5e\x32\x33\xb5\xa5\x16\x9d\xa9\x2d\x47\x5a\x58\x53\x57\x09\xe9\xc0\xe9\xd9\xfb\x67\xff\x7a\xfa\xaa\x97\x98\x0a\xb5\x1b\xcb\xdc\xa7\x7b\x57\x4e\xa6\xc6\x31\x81\xd3\x94\x2b\x53\x8b\x84\x74\x48\x07\x64\xe5\xd9\x48\xa1\x03\x7a\x0c\xc7\xaf\x06\x6f\x5e\x03\x75\xb0\x77\x57\xc8\x02\x7e\x70\x63\x63\x3d\x24\x7b\x2d\xdf\x04\x3e\x81\x67\x52\x01\x3d\xb8\x07\xf4\x03\x9c\x9c\x3e\x07\x4a\x95\x29\x68\x65\x31\x97\x17\x90\xbc\x78\xf3\xa4\x0f\x81\x14\x9e\x0e\x4f\x07\xdd\xe4\xeb\xf8\x07\x1e\x84\xcc\xe7\x20\x73\x48\x8f\x8c\xce\x65\x91\x9e\x21\xaf\xad\xf4\xb3\x01\xf3\x7c\x3c\x60\x7c\xc2\x0a\x74\xb0\x58\x10\x65\x8a\x02\x2d\x50\xdf\x1a\x8e\x3a\xcf\xac\xaf\xab\xd4\x8d\x21\x91\xda\x79\xa6\x94\xd4\x05\x58\x14\x10\x4c\xce\x85\x06\x1e\x79\xd6\x96\x79\x69\x34\x18\x0d\x7b\x77\xc7\xc6\x79\xcd\x4a\xbc\x97\x10\xce\x3c\x3c\xce\xa6\xcc\x66\x4a\x8e\xb2\x59\x5d\x66\x5c\x49\xd4\x9e\x72\xb4\x3e\xad\xb0\x84\xdf\x7e\xbb\xd3\x3f\x7d\x76\x27\xa8\x78\x84\xd6\x1f\xba\x27\x33\x8f\x6e\xa5\x6b\x58\x93\xb9\xe4\xcc\xa3\x4b\x5b\x5d\x87\x58\x19\x27\xbd\xb1\xb3\xb8\x0d\x9f\xe0\xcc\xdb\xa0\xd7\x62\x41\xfa\xa7\xcf\xae\x17\x3a\xc1\xd9\xa6\xcc\x81\x95\x53\xe6\xf1\x05\xce\xbe\x50\xf2\x0b\x9c\x6d\x09\xee\xc0\xeb\xd3\xa7\xa7\x5d\x10\xa8\xd0\x63\x8c\xc0\xdc\x28\x65\xce\x03\x8d\x43\x1e\x4d\xc4\xf2\x10\x92\x21\x7c\xb9\xaa\x83\x95\x1d\x30\x8b\x60\x6b\x0d\xe7\xd2\x8f\x81\xc1\xb4\x04\x59\xb2\x02\x9b\xef\x89\xe4\x93\xe8\x87\xd4\x62\x65\x60\xc4\x26\x28\x40\xea\xe6\x96\x90\xa1\xe7\xe1\x8a\x71\xd3\xa5\x22\xdb\x20\x5f\x5e\xf5\xad\x1d\xa3\xa2\x8f\xa8\x43\x3b\x45\x4b\x6d\x55\xba\x77\x24\x38\xa9\x37\x44\x01\xff\xc7\x3c\xf4\xb5\x47\x5b\x59\xe9\x10\x4e\xa4\xae\x2f\xe0\x11\x9c\x45\x62\xb8\x3b\x1c\xbc\x74\xf7\xc8\x88\x39\xac\xad\xea\x8d\xbd\xaf\x5c\x37\xcb\xb8\xd0\xa9\x45\x31\x66\x3e\xe5\xa6\xcc\xb8\xd1\x1e\xb5\xcf\x84\x74\x3e\x0b\xd2\xb2\x46\x56\xf6\x28\x7b\xd4\x30\xca\xf6\x02\x0b\x66\xf9\x38\x33\x8e\x14\x55\x31\xc1\x59\x2f\x97\x0a\xbb\x59\x16\xef\x51\x4d\x64\x66\xab\x92\x16\x55\x91\x0d\x07\x2f\xe9\xf3\xc1\x73\xfa\xa2\xff\x1f\xda\x48\xa1\x16\x15\x32\x87\xc4\x39\xc5\x59\x08\xa0\x5e\x3c\x65\xc7\xae\xcc\x38\xcb\x5a\xaa\x1a\xab\xe0\xe4\x48\x15\xdd\xde\x50\xde\x10\x83\x97\xa4\x41\xa1\x1b\x02\x87\xa0\x0e\x79\x28\x7a\x33\x74\x64\xd3\xa4\x78\xe1\x2d\x73\x5f\x66\x59\x0a\xfd\x78\xea\xef\x30\x71\xa3\xcf\x3f\xd1\xd2\xc6\x21\xfd\x31\x3d\x38\xd8\x65\xeb\xd3\x0a\xf5\x59\x40\x6d\x38\x32\xda\x33\xa9\xd1\xc2\x40\x31\x9f\x1b\x5b\x42\x38\xf4\xf7\x44\x33\x66\x81\xf7\x3f\xc3\xd8\xec\x63\x6d\x91\x1b\x8b\xad\x69\x57\xdf\x5b\x36\xab\xda\xda\x91\x96\x92\x5b\xe3\x4c\xde\xd8\x6e\x56\x97\x11\x8d\xb2\xcb\x93\xeb\x12\x8a\xaa\xe0\x63\xe4\x93\x9e\x36\x11\x35\x3f\xbb\xec\x1c\x0e\x4f\x03\x8a\xc6\xda\x05\x75\x25\x02\x32\xc3\xdb\xf9\xbc\x45\x61\xf7\xff\x46\xea\x5b\x8a\x5c\xf2\x00\x12\x58\x2c\xde\x6d\x15\xaa\xdc\x58\x60\xde\x63\x59\x79\x90\x1a\xe6\x07\x69\xfa\xf3\xe2\x57\x10\x86\x00\xcc\xea\x12\x5a\x35\x80\xce\x80\xfe\x09\x5f\x26\x33\x8a\x84\xef\xbf\x87\x91\x45\x36\x21\x00\x37\x5e\xf8\xed\x52\x8d\xbd\x79\xfb\xdb\xe2\xdd\xee\xab\xb7\x3a\x35\x95\x37\x67\x52\xa1\x48\x08\x84\x4a\xff\xf6\xed\xda\x69\xa0\xca\xc3\xcf\xf0\xee\xdd\xaf\xa1\x22\x69\x70\x0a\xb1\x82\x83\x5f\x01\x95\x43\xc0\x0b\xe9\xc3\x47\x2e\x89\x30\x1a\x6f\xf1\x86\xc5\xd2\x4c\xbf\xac\x05\x08\xd6\xe3\x0a\x99\x0e\x35\x8f\xd8\x12\xa8\xcd\xe1\xc6\x96\xe0\xa6\x50\x9d\xcf\x51\x8b\xc5\x82\xdc\x5c\x6b\x63\x37\x07\xdc\x22\x5b\x2b\xb9\xd2\x83\x74\xab\xfa\xe9\x0d\x98\xda\xc2\xbf\x5f\x36\xb5\xd6\x91\x78\x86\x09\x01\xc2\xf0\x09\x5a\x42\x64\x0e\xdf\x42\x61\xb1\xba\x54\xa8\xd9\x6a\x8a\x6e\xee\x3c\x1b\x35\x26\x25\x00\x6e\xe6\x3c\x96\xdc\x2b\x70\xde\x54\x2d\x0f\x1a\xef\x5d\x57\xa9\x97\x25\xda\x5b\xa9\x02\xc2\x48\x8e\xd7\xd1\xad\xed\x97\x93\xdc\xa5\x17\xb9\x03\x9a\x43\x26\x70\x1a\x60\x6a\xd2\xa4\x5b\xb6\x6a\x69\x2b\x66\xfd\x01\x01\x40\x3e\x36\x70\xe7\x66\x32\xd8\xba\x23\x04\xf6\x50\xd8\xea\xcf\xda\x78\x06\xb0\x0f\xfb\x77\xe0\xf1\xe3\xcb\xab\x07\x35\x4c\xad\xfd\xe6\x49\x02\x60\xd1\x79\x13\x32\x5f\x03\x1d\xee\xd8\x6f\x1a\xb4\xc0\xa9\x59\xca\x04\xc3\xd2\xe8\xf4\x83\x33\xfa\xb2\x39\x23\x00\x49\x68\x8c\x85\x95\x53\xb4\x49\x17\x92\x0f\xa6\xb6\x9a\x29\x91\x3c\x08\x7b\x42\xba\x00\x2a\x54\x61\xc1\xf8\x8c\x5a\x2c\xa4\xf3\x76\x96\x74\xc1\xdb\x1a\x49\xd3\x92\x5d\xb5\x25\xb3\x7e\xdb\x98\xbb\x09\x36\x7c\x97\x4b\x42\x5a\xcb\x54\xb5\x52\x01\x01\x96\x59\x7f\x1c\xe3\x27\x7d\x65\x04\xc6\x3c\x7f\x1c\x4d\xad\x03\xd5\xf7\x3b\xa3\x08\x3d\x17\xbb\x62\x68\xe5\xd5\x4d\x5f\x39\xee\xe4\x41\xa6\x6a\xbd\x0f\x9f\x3e\x35\xb7\xbb\xce\xad\x6b\xa4\x1b\x02\x1b\x87\x0a\xcc\x59\xad\xbc\xfb\x2c\x87\x86\x73\xd7\xbb\x33\xee\x06\xbb\x04\xec\x14\x2e\xe2\xa6\xe7\xd5\x83\x5f\x7e\xfa\xe9\xa7\x88\x9c\xdf\x54\xd6\x78\xd3\xdb\x9b\x0b\xe7\xbf\xfb\xee\xc1\xfd\x05\xf9\xa6\x32\xd6\x37\x0b\x9d\xce\xfd\x07\x0b\xf2\xcd\xe5\x68\x73\x18\x47\xaf\xe3\x61\xff\xf7\xc3\x93\x93\xf7\x87\x27\x27\xa7\xbf\x03\xad\x60\x2f\x32\x01\x5a\x06\xef\x78\x04\x4a\x9b\xff\x5f\xf5\x7f\x0f\x8b\xcb\x6d\x2a\x02\x6b\xd8\x8b\x3f\xe9\x07\x38\x3c\x3a\xea\x0f\x5e\x03\x3d\x6f\x51\x6d\x29\x87\x3a\x36\xc5\x36\xf8\xdc\xcc\x35\xd0\x95\x2d\x77\x03\xac\x9c\x47\x8c\x0c\x91\x10\x90\x44\x07\xaf\x9e\x33\x56\xa0\xf6\x71\xf8\xd4\xe8\xcf\x8d\x9d\x40\xed\xa5\x92\x5e\xa2\x83\xc2\x44\x2c\xf6\x06\x2c\xe3\x18\xd0\x50\xc8\x00\x3b\x69\x98\xdc\xf2\xd5\x61\x5b\x6b\x07\x23\xcc\x8d\x45\x10\xda\x05\x2c\x9a\x68\x73\xae\xc1\x9b\x88\x5e\x8d\x24\x04\xd4\x02\xea\xaa\xe9\xfd\x03\x7e\xcf\xc0\xc5\x52\x43\xce\xc7\x52\x61\x84\xf6\x15\xbc\x02\x15\xf7\xa0\xd7\x83\x24\x89\xf0\x2e\xcc\x25\xb8\x37\xd7\x6e\xce\x7c\x0b\x37\x87\xee\x59\x83\xf0\xb0\x58\x56\xbc\x96\x4b\x63\x3b\x87\x1e\x7e\xb8\x20\x78\x11\x6d\x7b\x76\x78\xf6\x66\x78\xdc\xbb\xb3\xc6\xe5\x65\xac\x13\x2d\x93\x66\x1f\x16\x8b\x3b\xf1\x20\xbd\x58\xa6\x4d\x18\x70\x28\xad\xac\x9c\x4a\x85\x05\x0a\xa0\x34\x14\x03\xba\x34\x68\xb8\x13\xd0\x29\x64\xdd\x2c\xfc\xda\xfd\x08\x14\x5b\x69\x37\xaa\x0c\x6d\x7d\x22\xb5\x0e\x02\x9b\x13\x84\x34\x45\x92\x72\x46\xbd\xad\x9d\x0f\x9e\x1d\x48\x0d\x93\x7a\x84\x8d\xd3\x5d\x30\x7c\xed\x10\x94\xe1\x4c\x01\xab\x64\xdb\x5f\x12\x17\x94\x93\x40\x2d\x42\xe2\x3a\x77\xe1\x7e\xb3\xde\x85\x7b\xe9\xfd\xce\x1f\x07\xcb\xee\x67\xad\xca\x75\x92\x26\x9f\x8d\x95\x85\xd4\x59\x53\x38\xb3\xd5\x0b\x02\x6d\x16\xd2\x4b\xe1\x5f\x2f\x23\x84\x4b\xfc\xf1\xd7\x73\x75\x42\x7f\x3d\x53\x6b\x8c\xcf\x22\x9b\xac\xe5\x43\xe6\x73\x1a\x12\x42\x23\xec\xa5\x4f\x18\x9f\xd4\xd5\x13\x65\x46\xaf\x42\x1c\x27\xc9\xad\xef\x0f\xab\x94\x0c\x48\x34\x45\x3b\xdb\xea\x34\x48\x07\x9c\x0f\x71\x0b\x05\xfa\x98\x53\xa3\x28\x25\x36\x1d\xc3\xfc\x2a\x72\x65\xf7\x49\x68\x8f\x82\x1e\x4f\xa5\xed\x5d\xdd\x6b\xcf\x95\x13\x21\x2d\xec\xad\xd1\xdd\xd2\x1e\x09\x73\xae\x95\x61\x22\xa8\xd9\xf0\x48\x3e\x33\x05\xfb\x9e\x8b\xc6\x26\xd7\x64\xe1\x95\x2c\xda\x4e\x9c\x3f\x08\xc4\xe4\xd9\x8a\xc2\xee\xf6\xd2\x2e\xe2\xf8\x9a\x55\x59\x33\x95\x02\x6d\xd6\xcd\xde\x0b\xe6\x59\xf6\xde\xd4\x2b\xd6\xeb\x66\xe8\x66\xa6\x0e\x19\x1a\xb6\x6e\xb9\x0b\x04\x83\x36\xb6\x68\x38\xd1\x51\xeb\xf4\x5e\x38\xb9\x11\x07\x8b\x45\x4b\x24\xe2\xa3\x5f\x6c\xdf\x7a\x41\x58\xeb\x92\x54\x8c\x5a\x02\x16\x5f\x53\x7a\x4b\x8b\xdf\xec\x97\x56\xfe\x92\x38\xf4\xc9\xcb\x60\x79\xb8\x2c\x6f\xb7\x35\xbe\x81\x28\xbe\xe3\x68\x56\xb9\xb1\xf1\x9f\xeb\xd9\x06\x1d\x83\x4d\xbe\xde\xb3\xc1\x96\xdd\xd5\x6f\xab\xad\xf5\xd8\xed\x5e\xfd\x6a\x7c\x44\x11\xfa\xaf\x8f\x9e\x1e\xbd\x3e\x79\x7f\x38\x38\xee\x25\x3f\x26\xd7\xb8\xee\x8a\xb2\x91\x26\x70\x89\xbd\x51\x7b\xed\xa5\xb9\xae\xc4\xc3\x96\x77\x42\xf4\xd0\x90\x3c\x57\xf3\x4a\xe3\x79\x4b\x10\xab\xd7\x5a\xf6\xb6\xcb\x52\x4b\x2f\x99\xa2\xed\xd3\x18\x24\xad\x2b\xf6\xe3\xbf\xd5\xe0\x79\x65\xb5\xfb\xf0\xc7\x5f\xf6\x1f\xac\x2f\x1d\xec\x24\x3c\xd8\x26\x7c\xb8\x93\xf0\x61\x24\x4c\x76\xab\x44\xbd\x99\xa0\x8e\x66\xa1\xb9\xb1\x34\xb6\x5d\x1b\xa4\x4c\x4c\xd1\x7a\xe9\x90\x56\x88\x96\xd6\x56\x39\xd8\x01\x93\x51\x0c\x21\xe5\x74\xdb\x4a\xd9\xfd\x8d\xb5\xad\xc1\x69\x65\xcf\x2b\xf0\x74\xa5\x55\xdb\xe0\xfb\x39\x01\x8e\xb1\x6f\x48\x22\x54\x87\x46\x24\x4c\x5a\xbe\xd6\x28\x28\x13\x25\x54\xd6\xe4\x21\xe4\x2f\x0b\x1b\x37\xda\x5b\xa3\x68\xa5\x58\x68\x37\x3a\xa1\x81\x61\xca\x19\xd0\x88\xe2\x92\x2e\x8d\x35\x36\x9d\x1a\x55\x97\xe8\x20\x04\x46\x1c\xce\x50\x2c\x3b\xa2\xd0\x9c\x36\x93\x06\x0f\x7d\x50\x68\x96\x5a\xf8\xa5\x25\xec\x3f\xfa\x79\x3f\xf4\x82\xab\xeb\xb4\xa0\x75\x0d\xff\xa0\x47\x33\x1c\xc4\x72\xe1\x66\x4e\x99\x02\x9c\xd4\x3c\x36\x58\x25\xd3\xac\x40\xc0\x50\x43\xfc\x38\x90\xf8\xb1\x35\x75\x31\x86\xe5\x7c\x41\x2e\x47\x82\x76\xc8\x58\x72\x59\x8d\x0d\x1b\x13\xda\xe6\x76\x40\x17\xf4\xb1\x5c\xd5\x15\x14\xa8\x71\xca\xe2\xd8\x1f\x11\xc4\x33\x3e\x59\xe3\x50\xeb\x92\xb9\x09\x94\xc2\x89\x25\x03\x60\x1f\x1d\xf2\x8d\xcf\xd2\x68\xb1\x43\x81\xe6\x6d\xe5\x7f\x3e\xde\x0c\x3d\x5f\x76\x9a\x74\x40\x1b\x8f\x5d\x60\xde\x94\x92\xd3\xcb\x80\x88\x5d\x2c\xb7\xcc\x8d\x41\x19\x53\x39\xa8\xb5\x97\x6a\xf9\x87\x18\xe9\xa0\xae\xb6\x55\xdf\xc9\x65\x25\xec\xaf\xf8\xe3\x85\xe3\x63\x14\x75\x8c\x87\x75\xd0\xb1\x38\x32\xc6\x87\x46\x90\x9b\xb2\x8a\x0f\x09\xbb\x5e\x5a\x12\xe2\xc6\xb5\x0f\xd5\x23\x20\x74\x73\xe6\x87\x87\x64\x3e\x47\xe5\x70\xb1\xd8\x32\xe6\x8d\xf7\x59\x8d\x70\xab\xa7\x8c\xff\x06\x00\x00\xff\xff\x14\x58\x90\x71\xf8\x1a\x00\x00")

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

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x5d\x6f\x1b\xb9\x15\x7d\xe7\xaf\x38\x2b\x07\xf1\x06\xbb\x9c\x89\x5b\x04\x06\x9c\x75\x80\x34\xf5\x6e\xb7\x69\x6a\xc3\xee\x07\x0a\xc3\x0f\x14\x79\x67\x86\x2b\x0e\x39\x21\xef\xc8\x56\x14\xfd\xf7\x82\x33\x92\xad\xda\x8e\x52\xb7\xd8\x97\x7d\x13\xcd\xfb\xcd\x7b\xcf\xb9\xe3\xbd\x6f\xca\xa9\xf5\xe5\x54\xa5\x06\x92\x6e\x84\x58\x2e\x61\x2b\x14\xef\x82\xaf\x6c\x5d\x5c\x90\xee\xa3\xe5\xc5\x99\x62\xdd\x9c\x29\x3d\x53\x35\x25\xac\x56\xc2\x85\xba\xa6\x08\xc9\xf0\xc1\x90\x4c\xac\x22\xf7\x5d\x91\x1a\x4c\xac\x4f\xac\x9c\xb3\xbe\x46\x24\x83\x46\x31\xb4\xf1\xd0\x83\xc5\x3e\x2a\xb6\xc1\x23\x78\x3c\xfb\xb6\x09\x89\xbd\x6a\xe9\xc5\x44\x68\xc5\x78\x53\xce\x55\x2c\x9d\x9d\x96\x8b\xbe\x2d\xb5\xb3\xe4\x59\x6a\x8a\x5c\x74\xd4\xe2\x87\x1f\xf6\x4f\x4e\x7f\xdc\xcf\x01\xbe\xa3\xc8\x6f\xd3\x1f\x16\x4c\xe9\x36\xd2\xfc\x37\x5b\x59\xad\x98\x52\xb1\x8e\xf4\x9c\xba\x90\x2c\x87\xb8\x18\xae\xf1\x19\x17\x1c\x73\x5c\xab\x95\x38\x39\xfd\xf1\xcb\x4e\x67\xb4\xb8\xef\xf3\x2c\xda\xb9\x62\x7a\x4f\x8b\x27\x7a\x7e\x4f\x8b\x07\x8e\xf7\xf0\xb7\xd3\x3f\x9e\x1e\xc1\x90\x23\x26\x70\x43\xa8\x82\x73\xe1\x3a\xcb\x24\xd2\x43\x89\x54\xc5\x14\xa1\x9c\x83\x76\x7d\x62\x8a\x09\x2a\x12\x62\xef\x71\x6d\xb9\x81\xc2\xbc\x85\x6d\x55\x4d\xe3\x79\x66\xf5\x6c\x78\x87\x22\x52\x17\x30\x55\x33\x32\xb0\x7e\xcc\x12\x25\xb1\xce\x29\x0e\x97\xa9\x30\xe5\x3d\xf1\x4d\xaa\x97\xb1\x21\x27\x0f\x65\xa2\x38\xa7\x28\x63\xd7\xa6\x2b\x91\x1f\xe9\xf8\x9c\x0c\xfe\xa4\x18\x27\x9e\x29\x76\xd1\x26\xc2\x5f\xac\xef\x6f\x70\x88\x8b\x41\x18\xdf\x9e\x9f\x7d\x48\x2f\xc4\x54\x25\xea\xa3\x3b\x6e\x98\xbb\x74\x54\x96\xda\xf8\x22\x92\x69\x14\x17\x3a\xb4\xa5\x0e\x9e\xc9\x73\x69\x6c\xe2\x32\x7b\x2b\x47\x5f\xe5\x61\x79\x38\x1a\x2a\x9f\x65\x13\x2a\xea\xa6\x0c\x49\xd4\x5d\x3d\xa3\xc5\x71\x65\x1d\x1d\x95\xe5\x90\x47\x37\xb3\x65\xec\x5a\x59\x77\x75\x79\x7e\xf6\x41\xfe\x74\xf6\x93\x7c\x7f\xf2\x2f\x39\x7a\x91\x91\x1c\xa9\x44\x22\x25\xa7\x55\x6e\xa0\xe3\x41\x2b\x36\xa9\x2d\xb5\x2a\xd7\x52\x3d\x75\xf9\x91\x07\xa9\xe1\xd9\x47\xc9\x1d\x3d\x78\x27\x9a\x03\xda\xd1\x38\x82\xbc\x9a\x3a\x32\xc7\x0b\x4a\xe2\x7e\x49\xe9\x86\xa3\x4a\x4f\xab\xac\xc4\xc9\xa0\xf5\x6b\x94\x78\x8c\xe7\xb7\x58\xe9\x90\x48\xfe\xbe\x38\x38\x78\xac\xd6\xa7\x1d\xf9\x8b\xc6\x56\x8c\x77\xc1\xb3\xb2\x9e\x22\xce\x9c\xe2\x2a\xc4\x16\x59\xe9\xd7\xe9\x66\x2a\xb3\xed\xdf\x46\xb1\xd5\xa7\x3e\x92\x0e\x91\xd6\xa5\xbd\x3d\x3f\xa8\x59\xb7\x66\x8e\xa2\xb5\x3a\x86\x14\xaa\xb1\x76\x8b\xbe\x1d\xd0\xa8\xbc\xd3\xdc\xf6\x50\x77\xb5\x6e\x48\xcf\x8e\x7d\x18\x50\xf3\xbf\x24\x9d\xb7\xe7\xa7\x19\x43\x07\xde\x42\xdf\x99\x8c\xcb\xb8\x5c\x2e\xd7\x18\x9c\xfe\x1c\xac\xff\x0a\xc1\x4d\xbe\xc7\x04\xab\xd5\xd5\x03\x9a\xaa\x42\x84\x62\xa6\xb6\x63\x58\x8f\xe5\x41\x51\xbc\x5a\xbd\x86\x09\x02\x58\xf4\x2d\xd6\x61\x40\x2e\x20\x3f\xe2\x69\x3e\x07\x97\x78\xfe\x1c\xd3\x48\x6a\x26\x80\x1d\xe9\x5e\x6e\x82\x78\xb6\x5c\xff\x5a\x5d\x3d\x9e\xf8\x3a\xa2\x91\x75\x2b\x65\x1d\x99\x89\x40\xe6\xf8\xcb\xcb\x2d\x6d\x48\xc7\x78\x85\xab\xab\xd7\x99\x8d\x3c\x92\x23\xea\x70\xf0\x1a\xe4\x12\x81\x6e\x2c\xe7\x43\x65\x85\x09\x9e\x76\xbe\x44\xa4\x36\xcc\x9f\x46\xfe\xb9\x72\xda\x91\xf2\x99\xed\x44\x6c\x21\x63\x85\x9d\xcb\xc0\xae\x26\x5d\x2e\xc9\x9b\xd5\x4a\xec\x66\xd9\x3a\x86\xbe\x83\x8e\xa4\xb6\xc8\xd6\x32\x6c\xba\x65\x4e\x0e\x08\x7d\xc4\x3f\x3e\x8c\x2c\x9b\xc4\xa0\xa3\x8c\x81\x09\x7a\x46\x51\x08\x5b\xe1\x1b\xd4\x91\xba\xbb\x80\xc6\xab\x91\x6e\xab\xc4\x6a\x3a\x16\x54\x00\x69\x91\x98\x5a\xcd\x0e\x89\x43\xb7\xb6\x21\x87\xbc\xfb\xae\x60\xdb\x52\xfc\xaa\x54\xc6\x16\xab\xe9\x4b\x72\x5b\xf7\xed\xac\x4a\xc5\x4d\x95\x20\x2b\x94\x86\xe6\x19\xa0\x66\xe3\xa0\x95\x91\x52\xe8\xa3\x26\xd9\xa9\xc8\x07\x02\x20\xdd\x04\xec\xef\x16\xc3\x83\x1c\x91\xcd\xa3\x8e\xdd\xc7\x3e\xb0\x02\x5e\xe2\xe5\x3e\xde\xbc\xb9\x4b\x3d\x87\x11\x7a\xcf\xf7\x35\x05\x10\x29\x71\xc8\x33\xef\x21\xcf\x1f\xdc\x2f\x97\x32\x77\x28\x7d\x44\x71\x1e\x1c\xe5\xf1\xae\xa2\xca\xf3\x21\x80\x71\x6b\xcb\x4e\x46\xe9\xd2\x28\x6a\x83\x2f\x7e\x49\xc1\xdf\x6d\x6c\x02\x98\xb8\x50\x4b\x13\xed\x9c\xe2\xe4\x08\x93\x5f\x42\x1f\xbd\x72\x66\xf2\x7d\xbe\x33\x36\x65\xa4\x91\x8e\x6a\xa5\x17\x32\x52\x6d\x13\xc7\xc5\xe4\x08\x1c\x7b\x12\xe3\x9e\x96\xe3\x20\x6f\x46\xbf\xdb\x15\x57\x91\x1f\x96\xfc\x71\x81\x7b\x2f\x5c\x59\x21\xd6\xf5\xeb\x7a\xe7\x32\x46\x6c\x70\xe1\xe7\xa1\xcb\x8a\xbf\x06\x43\x03\x12\xbc\x19\x1e\xc4\x67\xa9\xe7\xb9\x9b\xaf\x87\xc1\xcc\xa6\x73\x03\xe7\xc1\xc3\xb5\x52\x35\x79\x86\xf2\x06\x9e\xf8\x3a\xc4\x19\x7a\xb6\xce\xb2\xa5\x84\x3a\x0c\x00\xc0\x01\x51\x69\xca\x43\x68\x6c\xee\xf6\x42\xec\xe5\xf2\x6e\x94\x63\xef\x13\xa6\x54\x85\x48\x30\x3e\xe5\x11\x98\xf9\x70\xed\xc1\x61\x18\x9a\xd1\x13\x0d\x95\xe8\xbb\x71\xd9\xcc\xa0\xb1\x40\x1a\xd0\x4d\x5c\x37\xd6\xd1\x80\x27\xb7\x53\x0d\x69\x5e\xe0\xf8\x18\x93\xc9\x80\x29\x26\xdc\x21\xca\x88\x20\xa3\xce\x37\xd8\x5d\x8b\x8b\x11\x58\xb0\xda\x80\xec\xda\xca\x08\x43\x89\x18\xdf\xdd\x08\xba\xe9\x42\x64\x5c\xbc\xbd\xf8\xfb\xf9\xcf\xc7\xfb\x5b\x56\xfe\x19\xe2\x8c\xe2\xda\xc8\x78\x8f\xd5\x6a\x7f\x50\x94\x37\x9b\x77\xc8\x1b\xb5\x94\x5d\xb4\x73\xeb\xa8\x26\x03\x29\x33\x06\xc9\x4d\x41\x73\x4e\x90\x73\x94\x47\x65\xfe\x79\xf4\x09\x92\xd6\xde\x76\x86\x8c\x35\x2c\x8a\xde\x67\x87\xa3\x86\x10\x23\x32\x4b\xad\x24\xc7\x3e\xb1\xd8\xd5\xed\xdc\x7b\x32\x52\x99\x16\x5d\x0c\x79\x55\x40\xe8\xc8\xa7\xbc\xb7\xc8\xbc\x79\xc4\xe0\x64\xe7\x94\xa7\xb1\x55\x33\x56\x7f\x45\x2b\x3f\xe6\x76\x5f\xe7\xbe\x22\x28\x97\x02\x3c\x91\xb9\x93\x2c\x5c\xd0\xca\x15\xf3\xe0\xfa\x96\x12\x8c\x8d\x23\x5a\x92\xd9\xf4\x4a\x1e\xff\x71\xf4\x75\xee\x90\xdc\x46\xa2\x9d\x65\x41\xd9\xe2\xe5\xe1\xab\x97\x90\x5b\xc0\x18\xa2\xad\xad\x2f\xbf\x60\x3f\xc7\x31\x8e\xe4\x40\x1d\x69\x91\x5c\xa8\x91\xac\xd7\x43\xeb\xb5\xca\xe7\xef\x1c\x9a\x53\x5c\x70\x93\x45\xb8\x89\xa1\xaf\x1b\x6c\xa6\x5a\xdc\x4d\xdf\x7a\xb4\x37\x56\x6e\x27\xf4\x1e\x64\xde\xbf\x16\x7b\x48\xc4\xc3\x6c\xf5\x1d\x6a\xf2\x34\x57\x03\x07\x0f\x1f\x65\xac\xf4\x6c\xcb\x42\xef\x5b\x95\x66\x68\x4d\x32\x1b\x03\x50\x9f\x12\xe9\x7b\xc7\x36\x78\xf3\x48\x00\xe3\x9a\xf3\x3f\xab\x8f\xf8\xf2\x34\x6d\xb1\x07\x1f\x98\x8e\xa0\x38\xb4\x56\xcb\xff\x6c\x09\xe8\xa8\x52\x03\x17\x42\x97\xd0\x7b\xb6\x0e\xad\x4a\x03\x27\x26\xf4\xdd\xc3\xd0\x1f\xb5\x72\xeb\xec\xff\xff\x2f\x42\xd2\x0d\x99\x7e\xe8\x86\xad\x5d\x01\x91\xa6\x21\x70\x46\x26\x1d\xda\x6e\xe0\xf5\xc7\xd6\x9e\x89\x48\x4d\xcf\x26\x83\x98\x94\x6b\x9d\xef\x7e\x97\x57\x03\x97\x68\xb5\x7a\x50\xca\x9d\xd9\xe0\xf3\xe7\x91\x14\x6e\x37\x8b\x7f\x07\x00\x00\xff\xff\x5b\xbc\xde\x3a\x36\x11\x00\x00")

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
