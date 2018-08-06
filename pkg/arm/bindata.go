// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// data/azuredeploy.json
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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x1b\x5d\x6f\xdb\x38\xf2\xbd\xbf\x82\xd0\x15\x48\x53\xc4\x5f\x49\xda\x3b\x04\xb8\x87\xa4\x69\x37\x46\x93\xc6\x88\xd2\xbd\x87\x45\x50\xd0\xe4\xd8\xe6\x46\x22\x75\x24\xe5\xd4\xf5\xf9\xbf\x1f\xa8\x0f\x5b\xb2\x29\x53\xb2\xbd\x68\x97\xfb\x10\xaf\x34\x33\xe4\x0c\xe7\x7b\xd4\xf9\x2b\x84\x10\xf2\x5e\x2b\x32\x81\x10\x7b\x17\xc8\x9b\x68\x1d\xa9\x8b\x4e\x27\x7d\xd2\x0e\x31\xc7\x63\x08\x81\xeb\x36\xfe\x11\x4b\x68\x13\x11\x66\xef\x54\xe7\xb4\xdb\x7b\xd7\xea\xf6\x5a\xdd\x5e\x87\x42\x14\x88\x99\x81\x7b\x84\x30\x0a\xb0\x86\xf6\x9f\x4a\xf0\x7f\x78\x27\xe9\x0e\x44\x70\x0d\x5c\xff\x0e\x52\x31\xc1\xcd\x46\xbd\x76\xd7\xfc\x97\x03\x44\x58\xe2\x10\x34\x48\xe5\x5d\xa0\xf9\x22\x7b\x3a\xc5\x92\xe1\x61\x00\xa5\x87\x12\x94\x88\x25\x49\x1e\xfe\xf1\x6a\x3e\x6f\x21\x36\x42\x5c\x68\xf4\x86\x71\x0a\xdf\x51\xfb\x83\xe0\x1a\x33\x0e\xd2\x07\x39\x65\x04\xda\x03\x29\x22\x90\x9a\x81\x6a\x5f\x8e\x81\xeb\x81\x10\xc1\x40\x8a\x11\x0b\x40\xa1\xee\x71\xfb\x77\x0e\xda\x8f\x87\x1c\x74\xff\x1a\x2d\x16\xc9\x36\x66\xcd\x97\xbf\x92\x8d\xf5\x2c\x02\x73\xf4\x3b\x46\xa4\x50\x62\xa4\xdb\x5f\x40\xbf\x08\xf9\xdc\x99\x32\xa9\x63\x1c\x64\xff\xab\x32\xa6\x96\x88\x38\x62\x05\xce\x4f\xbb\xbd\xf7\xad\xee\x59\xeb\xac\xbb\x0e\x17\x08\x82\x75\x06\x35\x9f\x5b\x18\xb9\xcd\x00\xd0\xff\xd0\x9f\x0a\x2d\x16\xeb\x04\x38\x0e\x93\x13\x4e\x39\xe8\xf5\x77\xd1\x52\x08\x46\x98\xa5\x77\xe9\x21\x29\x95\xa0\x94\x1f\x61\x02\x56\x88\x22\xd4\x40\xc2\x88\x7d\xcf\xae\xc0\x06\x98\x00\xf7\xba\xe9\x25\x77\xfe\xe5\x59\x81\x9e\x36\x9e\x2e\x4e\x36\x0f\xa6\x92\x9b\xa9\xde\xca\x7e\x52\x54\x94\x07\x85\x11\x8e\x83\x75\x91\x94\x40\x1d\xe2\x29\xc1\x96\x84\x90\x28\x73\xce\xe8\xe9\xb9\x9d\xd3\x84\x37\xeb\x9b\xcd\xa7\x65\xa9\xac\xde\x2f\x4e\x12\x65\x07\x4e\x8d\x8e\xce\xe7\x9d\xb7\xe8\xf1\xfe\xfa\xfe\x02\xbd\x00\x0a\xf1\x0c\x0d\x01\x19\x5b\x41\x5a\x20\x3c\x15\x8c\x22\x3d\x01\xd4\x1f\x20\xcc\x29\xba\xbd\x42\x13\x90\x60\x2c\xe5\x05\x90\xd2\x22\x42\xb1\x62\x7c\x8c\xde\x76\x72\x5a\xd7\x18\x42\xc1\x7d\xd0\x0a\x8d\x84\x44\xfe\xf5\x97\x36\x42\x8f\x13\xe0\xe8\x0e\x2b\x0d\x12\x71\x41\x41\x21\x35\x11\x71\x40\x11\x11\x21\xa0\x38\x42\x58\xa1\x07\xc0\x74\x56\x20\x84\x63\x2d\x42\xac\x19\xc1\x41\x30\x3b\x49\xb6\x9f\x00\x27\x60\x0e\x08\x5c\x83\x04\x8a\x18\xd7\x02\x11\xc1\x15\xa3\x20\x53\x95\x36\x9b\xe2\x02\x99\x5b\x81\xe9\x15\x0e\x30\x27\x20\x51\x66\x00\x68\x24\x05\xd7\xe6\xdc\x86\xb7\xcb\x41\x1f\x29\x90\x53\x90\x29\x5a\x2e\xa8\xda\x66\x1b\xc5\xc3\x80\x91\xfe\xe0\x32\xbd\x4f\xf8\xf9\x86\xcb\xa2\x16\x8e\x58\xca\x54\x53\x03\x5e\x72\x13\xe4\xc7\xb9\x03\x3d\x11\xd4\xd0\xbd\x9e\x71\x1c\x32\x62\x31\x00\x8f\xd1\x00\x1e\x59\x08\x22\xd6\x7d\x7e\xc7\x78\xac\x93\x0d\x7a\xef\x2c\xb0\x94\x2b\x1f\xb4\xb9\x80\x6a\x2b\xf1\xa8\x08\x31\xe3\x5f\x70\x08\xb7\x78\x08\x41\x41\x26\x23\x36\x6e\x3f\x64\xde\xfb\x37\x29\xe2\xc8\x48\x62\xd3\x07\x94\xf5\x7f\x4d\x0a\xea\x39\xb6\xb3\x9f\x8b\xf0\x0a\x2b\x46\xbc\x6a\x1b\xca\x7f\xd6\xd6\x92\xa0\xa0\x88\x3b\x6b\x08\x85\x08\x38\x55\xf7\xdc\xea\xc6\xbc\x3f\xf2\xa0\xd6\xa7\x6f\x8e\x6a\x28\xea\xd1\x09\x3a\x2a\xea\xca\xd1\xf1\x53\x99\xe5\xa7\xbf\x4a\x43\x83\xe1\xee\x1a\x3a\xc4\xe4\x19\x38\xcd\xb8\x30\x51\x78\x2f\xaf\x9e\x91\xb3\x7b\x5c\x8b\x5f\xb5\x28\x74\xe2\x4f\x80\xd3\xfe\x20\x55\xcf\x38\x75\x46\x7b\x1d\x2b\xa7\x79\xa8\x68\x13\x49\x36\xc5\x1a\x9a\x1a\x76\x99\x46\x59\x81\x9c\x9b\xa2\xd4\x33\x98\x1d\x0e\xae\x9a\xeb\xcb\x1e\x17\xab\xdf\xd4\xbb\x59\xc6\x87\x22\xe6\xf4\x0b\xd6\x4b\x3d\xdb\x0e\xf6\x10\xa7\x79\xa6\x15\x6c\xe5\x03\x18\x1f\x2f\x21\x77\xd5\x90\x48\x48\xdd\x3a\x3f\x3f\x3b\x94\x86\x6c\xda\x55\xa3\x0b\x26\x82\x13\xac\xdf\x6c\xbf\xe7\x92\x17\x34\x77\x5c\x74\x04\x47\xc7\x27\xe8\xa8\x63\x31\xef\xfc\x99\x5b\x09\x1c\x0a\x9c\xd1\x19\x08\xa9\xbd\x0b\x74\x7e\x7e\xe6\x80\x07\x6e\x72\xa1\x4f\x81\xc0\x26\x5a\xf5\x07\xde\x05\x1a\xe1\x40\x81\x03\xad\xc2\x1f\xfc\x14\x71\x56\xf9\xa6\xe5\x8b\xbd\x85\x9a\x13\xaa\x2d\xd5\x06\xb9\x42\x09\xcf\x70\x7b\xcd\x94\x96\x6c\x18\xe7\x51\xe8\xda\x99\x93\xa3\xcc\x0e\x86\xd5\xf5\xc8\xda\xe9\x0e\x2a\xff\x64\x67\xd5\xc9\x8d\x75\x6f\x69\x47\x52\x68\x41\x12\xdb\xf4\x1e\x49\x74\x80\x42\xc1\xe2\xa8\x44\xac\x6b\x39\xb4\x94\xb9\x5f\xc9\x89\x31\x53\x1a\x4c\x71\xd0\xe7\x3e\x10\xc1\xa9\x41\x71\xe9\x15\x8f\xc3\x21\xc8\xfb\xd1\x20\xe7\xe6\xd4\x75\x07\x75\x35\xfd\xf0\x97\xf5\x8b\x64\xb5\xbe\x16\x12\x8f\xa1\xa3\xd2\xbf\x97\x84\x88\x98\x6b\x77\x5e\xfb\xae\xd5\x7d\xdf\xea\xbd\xfb\xcb\x2a\x9f\x52\x89\x30\x36\xbe\x62\xe6\x97\x8e\x58\x45\xc0\xd9\xd7\x48\xd1\x1f\x33\x49\xf8\x1a\x73\x8a\x25\xfd\x76\xfb\xe0\x57\xca\x33\x29\xb4\x25\xe6\x63\x40\xaf\x71\x14\xa1\x8b\x7f\x37\xed\x2a\x2d\x76\xa9\x34\x78\xfa\xd7\x07\x12\x4b\xa6\x67\x49\x89\x74\xa8\x9a\xf4\xf5\x96\xab\xa9\xbe\x15\xae\xc6\x2d\x83\x8c\xa3\xa8\x6d\x2a\xba\x1d\xe4\xaf\x32\x6e\xf6\xce\x9a\x70\x10\x88\x97\x6f\x4a\x4d\x0e\xd6\xc6\x21\x24\xcd\x85\x3d\x93\x59\xbf\xb8\x22\x11\x05\x45\x24\x8b\x72\x99\x26\x38\xc8\xf7\x6f\x90\x96\x78\x34\x72\xe7\xe0\x14\x94\x66\x3c\x91\xf8\xe5\x7a\x03\xe9\x6d\x03\x64\x13\xad\x1f\x8c\x6e\x26\xf7\x7f\xda\x3a\x3d\x75\x22\x33\x09\x24\x3f\x77\x3f\xcd\x78\xdd\x71\x97\x09\x73\x6d\x26\xbc\x77\x7b\x0d\xbd\xa5\x03\x3c\x0d\xcc\xcd\x85\x90\xe2\x95\xf8\x7f\xdb\xdc\x33\x67\x1d\x63\xf8\x6f\xaa\xd6\x0f\x22\x00\xe4\x85\x49\x8f\xcb\x2b\xd9\x6d\x71\xd5\xd6\xcf\xa4\x81\xfe\x2b\x69\xe8\xcd\xe3\xe3\xc0\xff\xa9\x3a\x7a\x7e\x7e\xe6\x48\x14\xd0\x41\xb4\xd4\x19\xfc\xff\x6e\x5a\x9a\xb5\x7a\xd7\x5f\x6e\x69\x10\xe7\x3f\x9d\x21\xe7\x83\x08\xa3\x58\x43\x3e\xb9\xb8\xc3\x64\xc2\x38\xf8\x04\x07\xe0\x43\x8d\x74\xe0\x9f\xad\xde\x69\xab\xdb\xdb\x23\xe8\x54\x84\xf3\x72\xa3\xac\x30\xdd\x49\x8c\xb5\x6a\x52\xb3\x44\x77\x74\x2b\xd6\x06\x35\x26\xf1\x9f\x72\xd0\x26\xb7\x5f\xeb\xae\x6f\xf5\x11\x8d\xf7\x75\x94\x1b\x6b\xbb\x37\xa5\x6e\xcd\x1b\xcc\x2e\x96\xf0\xed\xec\x14\xe6\xee\x4c\x29\x57\xe4\xaf\x4c\x5a\x35\x03\x59\xcc\xb5\x6c\x6d\x67\x82\x23\x4c\x52\xcb\xf5\xf2\x7d\x3e\x6c\xc9\xf3\xd0\x5a\xb2\x98\xaa\xc3\x9d\xcf\x7e\xc0\x12\x63\x3d\xc5\x2e\xdc\x62\x9e\x5b\xf6\x43\x3c\x86\xbc\x07\x9d\xf0\xe5\x6d\x5c\xa8\x17\x05\xd8\x5e\xf2\x97\x0e\x50\x22\xe9\x7f\xfe\xba\xed\xe0\x49\x97\x4c\x4d\x52\xb1\x6c\x20\x0f\xf2\xb7\x5b\x49\x48\x41\x63\xa2\xad\x04\xee\x47\xa3\x02\xb2\x4d\x0c\x16\xcd\x72\x66\x6e\x62\x0a\x32\x92\x62\xca\x32\xbb\xaf\x68\xa0\x78\x71\x34\x96\x98\xc2\x40\x04\x8c\xcc\xaa\xe7\x02\xa1\xa0\xa9\x13\xc2\x3c\xc6\x81\xa5\xef\x6f\x21\x5d\x76\x4f\x59\x76\x5d\xbd\x85\x50\x2e\x10\x94\xce\xef\x42\xc6\xbf\x2a\x90\xf9\x75\x92\x40\xc4\xb4\x15\xab\x8d\x96\x76\x09\x8d\xa4\x3e\x53\x1a\xb5\x59\xc5\x83\x35\x23\x69\x6d\xa3\x10\x30\x1e\x7f\x6f\xd6\x59\xf2\x28\x53\x78\x18\xc0\x00\x2b\xf5\x22\x24\xbd\x8c\xf5\x04\xb8\x66\x4b\x3f\xab\x65\xec\xea\x6a\x99\x9c\xb9\x56\x07\x25\x6d\xe6\x7e\x86\xd9\xf6\xc1\x6e\x71\xb9\xa9\x2e\xa9\x3f\xc3\xec\x1a\x6b\x9c\x09\xcd\xf7\x6f\x06\xf9\x76\x97\xca\xd7\x92\xf1\xf1\x4a\xad\x7d\xff\xe6\x33\xcc\xda\x4b\x08\xbb\x51\x54\x33\x82\xb5\x61\xd9\xeb\x4c\x44\x08\x9d\xd5\xf5\x76\xda\x4a\x4d\x3a\x38\xd6\x13\x21\xd9\x0f\xa0\xdf\x9e\x0d\xaf\xb5\xe8\x56\xb7\xa9\xf3\xb5\x39\xce\xae\x87\x5f\x91\x02\xd8\xd9\xf5\xb2\xc2\xbd\x96\xa2\xb3\xd4\xdb\x8d\x40\x02\xcf\xa6\xfa\x3b\xfa\xc4\x0d\xd2\xc2\xb8\x9c\x1a\xce\xa8\xce\x60\x62\x2f\xc7\x58\x16\x4e\x12\x92\x9a\xba\xe7\x12\x89\xe9\x2a\xcf\xd9\x20\x93\xe5\x40\x2b\x4f\x9b\xb8\xd6\x40\x81\x53\x5c\x96\x59\xca\x06\xf5\xf2\x68\x34\xdb\xc3\x84\xf1\xcd\xb4\x2d\xb9\xd8\x24\xc4\x57\x52\x49\x2e\x32\x27\x62\x82\xfe\x96\x04\x23\x5f\x5b\x5a\x98\x9e\x50\xd7\x4c\x3d\xbb\xfd\x15\x49\x7c\xf5\xd8\xb0\xfb\x00\x98\xfe\x47\x32\x0d\x2e\x99\x13\x09\x58\xc3\xfd\xb2\x72\xf9\x24\x45\x98\x30\xe3\x42\x4c\xbf\x4f\xa2\xb5\x4e\x86\x0a\xd6\x73\x59\x6e\x0a\x0d\x24\x84\x2c\x0e\x37\x7b\x42\xeb\x6b\x9b\x11\xef\x54\x57\x26\x87\xa2\x58\x63\xc3\x82\xdb\xeb\xd6\xe0\x90\x32\xf5\x6c\xb2\xa2\xdf\xae\xbc\x0b\x74\xe6\xa8\x89\x90\x4d\xfa\x1f\xc3\x48\xcf\x6a\x78\x5b\x2f\x88\x0d\x7c\x77\x47\x89\x3d\xb9\x34\xb2\xca\x03\x66\xe9\x6e\x2d\x0f\x98\xc1\xf6\xb9\x06\x39\xc2\x04\x6a\x8e\x79\xf3\x55\x43\xde\xcb\x56\x99\xb3\xba\x46\x0d\xcb\xff\x02\x0e\x0b\xb1\x9c\xd5\x0a\xf6\x4b\xa4\x74\x02\xd6\x1f\x7c\x12\xf2\x05\x4b\x9a\x9a\x64\x03\x7c\x5b\x51\x51\xfb\xc8\xa8\xf6\xfc\x78\xaf\xda\xa5\x6a\x39\x46\x31\xab\x13\x46\x0d\xf5\xa1\xb8\xea\x4b\x02\xa1\xd2\x27\x3e\x24\xd9\xb3\x41\x3a\x83\x76\x54\x9c\x35\xfc\xe6\x4a\x54\x22\x90\x7e\x87\xb7\xd3\xe6\x68\xa5\x0e\xf3\x39\x1b\x6d\xd6\xf2\x8b\xc5\x7c\x6e\x7d\x68\x82\xeb\x62\x51\x6f\xa6\x57\x59\xdc\x9f\xa0\xa3\x4e\xf6\x15\x61\x27\xfb\x14\xf0\xe8\xf8\x69\x3e\x07\x4e\x6d\x1f\x20\xb9\x56\x4d\xed\x2a\x31\x9f\x7d\x26\x11\x65\xed\xa4\xe6\xd3\x65\x2b\xd5\x95\x4e\x35\xd4\xa6\xd5\xb9\xf6\xd3\xaa\x25\x9d\xca\xc9\xf0\x4e\x24\xdd\xb9\xf6\x26\xc6\xce\xe1\xb7\x6a\x79\xc5\xa6\xcd\x55\x83\x4f\x97\x5c\x6b\x2f\x31\xff\x32\x5f\x6c\x54\xad\xe6\x77\xe7\x4c\x04\xf6\xdf\xaa\x1e\xe4\xf6\x02\xce\x4d\x67\x5b\xae\x63\x85\xaf\x48\x70\xe0\xbb\x06\x6e\x0a\x8d\x5a\x29\xce\x12\xfa\xa0\xe9\x0c\x51\xae\xe4\x1b\xed\x9a\xce\xe0\x58\x8b\xaf\x69\xcf\xe8\x8e\x71\x21\x57\x9d\xe5\x06\xe9\x49\x24\x85\x06\xa2\x81\x3a\x3f\x48\xb5\xa2\xa7\x83\x92\xac\xc8\xbb\xc2\x0a\xde\x9f\x7f\xe4\x44\x50\x40\x6f\x7c\x8d\xa5\x8e\xa3\x82\x1b\x39\xb6\x7e\xa7\x6a\x5b\x75\x13\x8f\x52\xd9\xbb\x32\xdd\xcb\xe4\xdf\x58\x7c\x5c\x5d\x68\x4d\x72\xaa\x20\x83\xba\x47\xc8\x67\x02\x1f\x62\xa5\x45\xe8\xa7\xf2\x68\x80\x7b\x83\x39\x0d\x40\x16\xc7\x02\xed\xae\x5b\x4a\x07\xb6\xa0\xcd\x2e\x62\xd5\x70\x64\xdd\xc9\x64\x7d\x6f\x4f\xc4\x3a\x8a\x75\x2a\xba\x57\x8b\x57\xff\x0f\x00\x00\xff\xff\xbf\x72\x5a\x79\x13\x33\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xfd\x73\x1a\x39\x96\x3f\x6f\xff\x15\x3a\x3c\x7b\xd9\x4d\xa5\x01\x3b\x4e\x26\x61\xe2\xad\xc2\x18\x27\x54\x88\xcd\x02\xc9\xdc\xdc\xde\x94\x4b\xb4\x1e\xa0\x71\xb7\xd4\x91\xd4\xc4\x4c\x26\xff\xfb\x95\x3e\xba\x5b\x0d\x8d\x8d\xed\xec\x5e\xdd\xd5\x65\xaa\x3c\x20\xbd\x2f\xe9\x7d\xe8\xbd\xa7\x6e\x0e\xfe\xad\x35\xa3\xac\x35\xc3\x72\x89\x42\xb8\x09\x82\x03\x34\xbd\x3c\xbb\xec\xa0\x16\xa8\xa8\x45\x98\x4c\xb0\xfc\xdc\x24\x2d\x2e\xe8\x82\xb2\x30\x4b\xa5\x12\x80\x93\x90\x30\xd9\x8c\x38\x9b\x23\x2a\x51\x94\x09\x01\x4c\xc5\x6b\xb4\xc4\x82\x44\x9c\x00\xf9\x09\x51\x15\x1c\xa0\x54\xf0\x19\x9e\xc5\x6b\x24\x97\x3c\x8b\x09\x7b\xa2\xd0\x0c\x82\x60\xd2\x1f\x7f\x1a\xf4\xfa\x57\xd3\x5f\x46\xfd\x13\x4b\x39\xa0\x73\xf4\x0f\x14\xce\x51\xc3\x30\x96\x6b\xa9\xa9\xd3\x45\x0b\x2b\x9e\xd0\x28\xe4\x29\x30\xb9\xa4\x73\x15\x32\x4e\xa0\x81\x7e\xfd\x09\xa9\x25\xb0\x00\x21\x84\x2a\xe4\x36\xe1\x83\x39\x0d\x82\x2c\xe1\x19\x53\xa8\x95\x30\xd5\x12\x20\x79\x26\x22\x40\x7f\xfc\x81\x94\xc8\x20\x48\xae\xe7\xb2\x79\x33\x97\x9a\x7b\x8b\xc0\xaa\x45\xa8\xbc\x6e\xe1\xdf\x33\x01\x05\x70\x98\x62\xa1\x0e\x03\x88\x96\x1c\x3d\xb9\x1d\x08\xa1\xd6\x0a\x8b\x56\x4c\x67\x2d\xc2\xa3\x6b\x10\x08\x69\xe2\x68\x21\xd2\xcf\x19\x57\x18\xa1\x36\x6a\x3f\x41\x7f\xfb\x9b\x59\xe8\x5c\x2a\x3c\x0b\xe4\x5a\x2a\x48\x22\x15\x23\xa9\x78\x8a\x2c\x5e\x53\x82\x58\xd1\x08\x02\x27\x7c\x95\x6a\x20\x40\x2a\x2e\x20\xe2\x0c\x85\xe3\xad\x59\x9f\x22\x16\x6a\x93\x64\xa1\xe6\x88\x33\x49\x09\x08\x34\xc7\x91\x42\x6a\x89\xd5\xd6\x1e\xc8\x48\xd2\xc3\x56\x9c\xb1\xf6\x2e\x65\x6b\x4d\x63\xa1\xa8\xa2\x9c\xdd\x82\xfe\x93\xc7\x2d\x53\x99\x00\x24\x95\xc0\x0a\x16\x6b\x34\xe7\x02\x09\x90\xf4\x77\x90\x88\xce\x83\x03\xc4\x00\x08\x90\x42\x4c\xc6\x95\xc7\x99\x32\x05\x8c\x50\xb6\x30\x78\x6a\x49\x25\x92\x91\xa0\xa9\x42\x8a\xa3\x19\x20\x01\xc0\x94\xc0\x4c\x3d\x43\xb3\x4c\x19\x20\x89\xe7\xa0\xd6\x9a\x30\x57\x48\xa6\x10\xd1\xf9\x5a\x13\xf0\x95\xbf\x04\xe1\x59\xc3\x2d\xdb\x90\x5b\x4e\xbd\x39\x78\x80\xa5\x5a\x40\x45\xc4\x19\x02\x81\x39\xce\x62\x25\xeb\x0c\x61\x43\xd5\x1a\x6b\x97\xa2\xcd\x9c\x15\xa1\x71\x7a\x79\x39\x9d\x4c\xc7\xdd\xd1\x55\xef\xf2\xe2\x7c\xf0\xf6\xea\xa2\xfb\xa1\x7f\xa2\x3d\x25\xb4\x6e\x14\x26\x58\x2a\x10\x8d\x9c\x5b\xe9\x5f\x3f\x7c\xf5\xdd\xe7\x9b\x71\xaf\x20\x90\x40\x50\x48\x51\x08\xa8\x21\x0f\xce\xfa\xa7\x1f\xdf\x5e\x0d\x2f\xdf\x0e\xfb\x9f\xfa\xc3\x93\xa3\xcd\x81\xe3\x83\x06\xda\x8b\xaa\x56\x04\x91\x0a\x51\x86\x54\x94\x3e\x3b\x7a\xfe\xaa\xfd\x13\x22\x3c\x38\xd8\x9a\xf8\xf1\x75\x01\x61\x3e\xbc\x3a\x3e\x7e\x9e\x7f\x38\xb6\x1f\xda\x2f\x9e\xa3\x8c\xb8\x0f\x7a\xe4\x75\xfb\xb5\x25\xf7\xa7\x54\x70\xc5\x4f\x7e\xf8\x4a\xa4\xfa\xf3\x9f\x9f\x3d\xfd\x16\xfc\x29\xe5\x42\xd9\x81\x83\x83\xa7\xcf\xbe\x05\x7f\xa2\xa9\xc2\xb3\x18\x24\x0a\xbb\xe8\x72\x72\x75\x3e\x18\xf7\x7f\xee\x0e\x87\x57\xdd\xe1\xf0\xf2\x67\x14\xa6\xe8\x07\x43\x04\x85\x89\x76\x20\x05\x28\x0c\xed\xff\x2f\xfa\x3f\xeb\xc1\x7c\x3a\x24\x9a\x34\xfa\xc1\xfc\x0d\x7f\x43\xdd\x5e\xaf\x3f\x9a\x06\x84\x33\x08\x82\x9c\x49\x28\xf1\x0a\xd0\xe6\xce\xe7\xb3\x41\x20\x12\x14\x8a\xb9\xdd\x43\xad\xd6\xd6\x53\xfb\xd9\x86\xc6\x96\xd5\x5d\xeb\x69\x10\xcc\xb0\x84\x97\xc7\x28\x24\xe8\xcd\x9b\x37\xe8\xeb\x57\xd4\x03\xa1\xba\xf2\x74\xad\x40\xa2\x66\xcf\xd0\x6d\xea\x31\x3a\xa7\x11\x56\x20\x9b\x7d\x15\x91\x1e\x36\x63\xe8\x0f\x74\x6a\xf0\xfb\x4c\xbb\x2d\xfa\xf6\xcd\x89\x64\x58\x46\xb8\x19\x09\xf5\x40\x0e\x13\x10\x2b\x10\x7b\x70\x91\x16\xb0\x96\xd3\x48\xd0\x15\x56\xf0\x1e\xd6\xfb\xf2\x7b\x0f\xeb\xbd\xd8\x5d\xc3\xfa\x81\x0b\x1b\xc1\x5e\xcb\x4a\xe1\x3b\x2c\xca\xf0\xba\x73\x49\x86\x95\x5e\xd0\x36\xaf\x5f\x70\x12\x7f\xc0\x42\x2e\x71\x5c\x70\xe9\x92\x84\xb2\xf7\xd9\x0c\xac\xd1\xed\xa4\xed\x4c\x4d\xfb\xa9\xf9\xd3\xbc\x2e\x70\x1e\xb0\x75\x77\x59\x9c\xcf\xcd\x19\x5e\x90\x5c\x13\x2a\xb4\xe7\xd5\x98\x3e\xc3\x09\x90\x7f\x92\xf5\x57\x39\xd9\xff\x35\xf5\x56\x87\x0f\x76\x89\xfb\xb1\xdc\xc5\x66\x4f\xd3\xe9\xe1\x5b\x8d\x66\x8b\xd7\xc3\x9c\xe1\x32\xcf\xa7\x7a\x9c\x49\x1e\xc3\x7d\x16\x68\xb4\xd7\x8a\x1c\xe2\x63\xd6\xba\x25\xc5\xfe\x2b\xaf\x0a\xf1\xb0\x4d\x38\x17\x9c\xa9\x91\xe0\x37\xeb\xfb\x69\x78\xae\xf1\xc2\x54\x23\x3e\xdc\xa8\x26\x36\x75\x9b\xd0\x05\xa3\x6c\x71\x3f\x01\x5c\xda\x17\x4a\xba\x60\x8f\x8c\x54\x5b\x62\xec\xaf\x82\x0d\x29\x1e\x1e\x95\x7b\x31\x05\xa6\x1e\xec\xd6\x16\xfb\xb1\xe1\xda\x09\xb1\xff\xf2\x6b\x64\x78\xd8\x16\x74\x17\x0b\x01\x0b\xac\xb8\x28\x0d\xf2\x3e\x9b\x81\x0b\xfc\xd0\xb3\xcc\x47\x6d\x48\xad\x48\xfb\x6f\xcd\x0e\x89\x1e\xb6\x3d\x1f\x0c\x4d\x7d\xe6\xc5\xa0\x1e\x6c\x2a\xd7\x16\xff\x7b\x58\x4b\x9d\x40\xf7\x36\x9b\x0d\x79\x1e\xb3\x35\x36\x82\x3d\x74\x63\x5c\x18\xfb\x5e\xdb\xe2\x0b\x73\xef\x4d\xa9\xc8\xf2\x98\x2d\xd9\x27\x81\xad\x95\xe0\x3b\x24\xb4\x15\x09\xee\xbd\x05\xb7\xa5\xb8\xa3\x6c\x16\xd3\xa8\x86\xbf\x0b\xe2\xdd\x28\xd2\xf5\xe6\x7b\x58\x37\x0b\xd0\xfb\xc5\x72\x6c\x29\xc8\x66\x6a\xf0\x77\x88\xb1\x73\x1f\xb6\xe4\x78\x28\x77\xcb\xa1\x9e\x7d\xce\xec\x9d\x1a\x61\x29\xbf\x90\x3d\x79\x2c\x55\x6a\xc0\xf7\xcb\xb3\x4b\x1f\xdf\x2f\xd1\x76\x4c\xca\xb6\x96\xe7\xe6\x2e\xed\x0e\x22\xac\xfc\xcc\x5f\xff\xb1\x3d\xb7\x37\x6f\xfa\x97\xe7\x41\x7f\xda\x3b\xbb\xea\x9e\x7d\xea\x8f\xa7\x83\x49\xff\xaa\x37\x1c\xf4\x2f\xa6\x57\x1f\xc7\xc3\xc9\xc9\x52\xa9\x54\x76\x5a\xad\x1f\xfe\xb2\xe4\x52\xe9\xd4\xe7\xaf\x1d\x5d\x54\x5b\x9c\x5e\x7f\x3c\xbd\x3a\x1f\x0c\xfb\x27\xb5\x85\x99\x85\xb1\xd4\x0c\x68\xf7\xe3\xf4\xdd\x89\x69\x78\x98\xa9\xb3\xee\xb4\x7b\x75\x36\x18\x9f\x54\x5b\x11\x66\xae\x3f\xec\xf7\xa6\x83\xcb\x8b\xab\xe9\xe0\x43\xff\xf2\xe3\xf4\xe4\xe8\x45\xbb\x6d\xa7\xde\xf5\xbb\xe3\xe9\x69\xbf\x3b\xbd\x1a\x5c\x4c\xfb\xe3\x4f\xdd\xe1\x49\x31\x37\xb8\x18\x4c\x07\xdd\xa1\xb7\x9a\x51\xbf\x3f\xbe\x6d\x2d\xaf\x36\x30\x7b\xc3\x8f\x93\x69\x7f\x7c\x62\xb7\x31\x6c\x9b\x7f\x05\x6e\x65\xd4\x60\x3f\xf3\x87\x0e\x6b\x01\x0f\xb7\x01\x8f\x6a\x01\x8f\x3c\x79\xde\xf7\x7f\xd9\xb1\xb5\xda\x36\x0d\xc8\x70\x30\x99\xf6\x2f\x6a\xf5\xd5\x6e\x9a\xff\x3c\x5d\x39\xe0\xed\xed\x28\x41\x73\xd6\xa6\xe7\xe3\xed\x92\x1d\x35\x98\x75\x1a\x2f\x6a\x56\x0f\x6c\xb7\xd2\xcd\x7c\xcd\xe2\x8a\x72\xb4\x84\x9a\x8e\xb5\x2a\xce\xae\x7a\xdd\x4d\x60\x97\xfb\x1a\xd0\xbf\x7f\xbc\x9c\x76\xaf\x4e\xbb\xbd\xf7\xfd\x8b\xb3\xab\xd3\x5f\xa6\xfd\xc9\xc9\xf1\xd1\xeb\xe3\xd7\x2f\x7f\x3c\x7a\xfd\xd2\xc2\xdc\x4d\xe9\xf2\x3c\x08\xa2\x6a\xc9\xe8\x15\x95\x35\xe3\xe6\xa8\xc8\x93\xf0\xbb\x30\xd3\x6b\xda\x8a\x70\xa8\x44\x26\x55\xcb\xf6\x78\x5b\x98\x45\x4b\x2e\xa4\xe7\xb9\x8e\x58\x96\x12\xac\x20\xcc\xe1\x7d\xf7\xad\x0b\xdc\xae\x31\xd7\x5c\xe3\x24\x76\x0e\x8d\x49\x42\xa5\xa4\x9c\xd9\x98\xd2\x09\x10\x4a\xe3\x6c\x41\xbd\xef\x08\x9d\x66\x34\x26\x67\xae\x85\x68\x87\x10\xb2\xb4\x32\x81\x15\xe5\x2c\x1f\x44\x08\xa7\xf4\x13\x08\x4d\xb1\x83\x56\x87\xc5\xf0\x35\x65\xa4\x53\x25\x64\x39\x94\x0c\x2e\x57\x20\x04\x25\xf0\x78\x0e\x05\x25\x8f\xc5\x88\x93\x91\x00\x09\xea\x31\xd4\x9d\xe8\xdd\xea\xa6\x19\xa0\x42\x37\x4d\xca\x5b\x83\x04\x2f\x60\xc4\x63\x1a\xad\x1f\xc6\x0e\x6e\x20\xca\x34\xe4\x38\x8b\xcb\x0d\x41\x28\x44\x09\x56\xd1\xd2\xd0\xef\x32\xc6\x95\x21\xe7\x01\x68\x90\x6b\x58\x77\x10\xd5\x20\xb2\x59\x11\x8b\x00\x5b\x87\x05\x69\x0f\x07\xa1\x15\x8e\x33\xe8\xa0\x86\xf6\xbe\x86\x37\xa3\xbd\xba\x53\x8a\x13\x12\x60\x14\x88\x07\xc0\xd9\xd8\x5d\x45\x6c\x48\x91\xdf\x50\x74\x50\xca\x89\xdc\x31\x35\xd3\xea\xf2\x27\x05\xfc\x06\x91\xea\xd8\x5e\x77\x39\x2c\xaf\x69\x7a\x69\x38\xc5\x46\x8e\x73\x4c\xe3\x4c\xc0\x06\x9c\x55\x92\xb7\xf9\x4e\x3f\x65\xaa\xef\x59\x79\x99\x03\x0e\xd8\x9c\x5b\xd9\x23\x10\xea\x9c\xc6\xd0\x41\xb7\xd4\x2b\x86\x13\xac\x6f\x85\xd3\xd1\x09\xa7\x74\x08\x2b\x88\x65\x27\x08\xb5\x6e\x37\x54\x8d\x33\xb5\x2c\xc5\x11\xf0\x39\x03\xa9\xde\x01\x26\x20\x9c\x30\x46\xb8\x5e\xb7\x83\x6a\x6a\x79\x0f\x80\x27\x09\x67\x17\x38\xc9\x15\x10\xee\x10\x2a\xb0\x86\xa5\x04\xb6\x5c\x46\x02\xe6\xf4\xa6\xc4\xfa\x8f\x70\x0c\x09\x57\x10\xf6\x35\x4c\x68\x46\x17\x82\x67\xa9\x05\xdf\x86\x7b\xab\x27\xcd\x60\x26\x41\x68\x4b\xd9\x05\xf9\x51\x82\x08\x22\xce\x94\xe0\x71\x0c\x9e\x16\x20\x86\xa8\x74\x88\x98\x47\xd7\x17\xc6\xe0\x36\x73\x94\xb0\x44\xd6\xd6\xe2\x72\x31\x93\xca\xb1\x85\xce\x6c\x2d\x01\x5b\xe5\x17\x2e\x57\x68\xb3\xa6\x17\xe1\x2c\x26\xd7\x63\x4d\x9f\xc0\x63\xd9\x41\x8d\xa7\x8d\x20\xe2\x42\x76\xe3\x98\x7f\x01\x72\x69\xa2\xab\xec\x04\x84\xc9\x72\x35\x33\xca\x48\x97\x10\x01\x52\x76\x50\x7e\x58\xbe\x6a\xbf\x78\xee\xe6\x2e\x40\x7d\xe1\xe2\xba\x83\x54\x94\x1e\x07\x50\x14\xf4\xb9\x01\x46\xb8\x83\x6a\x9a\x81\xfe\x4a\x76\x34\x15\xbc\x95\xec\x28\xf9\x11\xca\x44\x6c\x34\x13\xa2\x9d\xa9\x9a\x46\x9a\x28\x2e\xf0\x02\xca\x55\xe9\xf4\x50\x30\x50\x20\xdd\x94\x35\x9c\x8e\x37\xd1\xa4\xbc\x0e\xb0\x1a\xd9\xb4\x4e\x27\x5a\xa7\x1b\x64\xfc\x10\x55\x03\xe6\x13\x31\x41\xad\x94\x6c\xce\x45\x82\x55\xc7\x4f\xb8\x07\x25\xc4\xb9\x99\x45\x7f\xa0\xcf\x19\x57\x3a\x1b\xb6\xe8\x7e\x64\xd0\x44\x28\x53\xda\x78\xe3\x31\x2c\xa8\x54\x62\xfd\xce\x6d\x49\xc7\xdd\x5f\x86\xc2\x4d\x34\xdd\x25\x5a\x53\xae\xa2\xce\x8b\x76\xbb\x1d\xd8\x78\x63\xd3\x70\x17\x6a\xae\xfd\x8a\xdb\xd7\xeb\x6e\x5d\xd6\x54\xfd\xdb\xea\xac\x29\xc5\x11\x4a\xb9\x50\x1d\x74\xd8\x3e\x7a\xd1\x0e\xca\xcd\xf7\xe5\xd1\xdc\x71\x4a\x6d\xa1\xd7\x15\x8b\x2c\x01\x96\x1f\xe1\x51\xcc\x33\xe2\x52\x82\xdc\x63\xfd\xd4\xc1\xcc\xa7\x82\xaf\x28\x01\x61\xef\x19\x4d\x29\xe0\x21\xe7\xb3\x45\xe0\xd1\x40\xe6\xb3\xc8\x98\xa2\x09\x6c\x90\x97\xa0\x14\x65\x0b\xd9\xbc\x7e\xa5\x6d\xa6\xb5\x3a\xc4\x71\xba\xc4\x87\x27\x45\x18\x97\x56\xe9\xe1\x0c\x47\xd7\xc0\x48\x8e\xa8\x0d\xf3\x79\x05\x20\x01\x42\x71\xa8\xd6\x29\x14\xcc\xd3\x34\xd6\xe5\x2d\xe5\xac\xb5\x62\xa4\xe9\x99\xa7\xb9\x38\x9b\x65\x5a\xf4\xd2\xab\xff\x95\xdb\x11\xc5\x99\x09\x63\xd2\xb6\x30\x43\x6d\x04\xe1\x5c\x2b\xb8\x86\x53\xb5\x47\x5f\x87\x7e\x0d\xeb\x3d\xb0\xad\x91\xd8\xef\x83\x51\x07\x1d\x1e\xfd\x68\x62\xd2\xe1\xdd\xc7\xdf\xae\xbe\x4b\x25\x66\xee\x6a\x88\x20\x24\xa3\x25\x90\xac\x88\xf4\x16\xbc\xae\xa6\xce\xe1\x9a\xbf\x49\x93\x90\xb8\x30\x2c\x27\xd9\x8c\x81\xb6\xed\x1f\x8f\x9a\xcf\x4d\x1c\x6d\x1d\xbe\x0c\x2c\x96\x95\xda\x68\xad\x08\x1d\x43\xce\x53\x6d\x32\x3d\x77\x24\x32\x66\xcf\x95\x8d\x9c\x12\x47\x11\xa4\x7a\x5a\x01\x53\xd3\x75\x0a\xb2\xb3\x8f\xd9\x3c\xf3\x61\x9c\xa4\x08\xcd\x32\x21\x55\x07\xbd\x6c\xb7\x03\x97\xe0\xe5\x54\xf7\x22\x6a\x90\x3e\xa7\xb2\x83\x9e\x1b\x0a\x5b\x6b\xd1\x85\xbd\xf3\xe2\xad\xf3\xd0\xaf\xd9\xed\x88\x6d\xa6\x7c\x1c\x0f\x4d\x38\x4c\x05\x65\x0a\x35\xf2\x40\xdf\x30\xf1\x51\x61\xca\x6c\xd3\x87\x46\xd0\x1c\x09\x9e\x82\x50\x14\x64\xf3\x52\x44\x4b\x30\xcf\x45\x70\x31\x12\x5c\x5b\x96\xb9\xf6\x98\xb8\x6b\x0f\x1d\x58\x2d\xfd\x3c\x38\xfa\x51\x95\xd9\x53\xad\x0c\x38\xce\x5a\xdd\x69\xe7\xce\x9c\x88\x12\xa1\x63\x55\xf3\xf0\xe8\x95\xd5\xe7\xb1\xd9\x01\x7d\x02\x59\x6d\x0f\x81\x2d\xd4\xb2\x83\x5e\x07\x26\x4d\x31\x41\x79\x30\x72\x54\x7a\x83\xb3\xb1\xa3\xe4\x0e\xd6\x96\xde\x34\xc7\x7b\x64\x4a\x16\x9b\x3a\x08\x20\x4b\xac\xbc\x5a\x89\xaf\x64\x28\x0d\x87\xd2\xc0\x3c\xaa\x9b\x46\xc6\xab\x89\x19\x96\x12\xd4\xbf\x78\x73\x1b\xf9\x9d\x51\xab\xe1\x6f\xb4\x4e\xc7\x30\x53\x7e\x65\x96\x80\x5a\x72\xd2\x41\x38\x53\xfa\xec\xa4\x04\x98\xa2\x6a\x3d\x72\x81\xc8\xed\x58\xcc\x17\x94\x79\xb9\x72\x82\xd3\x94\xb2\xc5\x07\x87\x1c\xc5\x98\x26\x41\x99\xed\x77\x75\xd8\x42\xdd\x33\x33\x54\x8d\x69\x3b\x0a\x16\x43\xc1\xcb\xff\x21\xc1\x34\xf6\xab\x16\x33\x50\x7c\xa7\xc4\x9f\x93\xd9\x2c\xa8\x94\x1b\xde\x9c\xfe\x5e\x7c\x4d\x05\xcc\x41\x08\x20\x1f\x5d\xba\xe9\x43\x66\x8c\x7e\xce\xe0\xca\x43\xb0\x11\x69\x70\x56\xe4\x07\x3b\x55\xe4\x86\x46\x82\xb2\x88\xa6\x38\xce\xd5\xe4\xe2\xe3\x59\x55\x09\x25\xed\x09\x44\x02\xd4\xc3\xe9\x5b\xfc\x6d\xea\x36\xaf\xd0\x46\x32\x38\x1b\x6c\xa8\xd4\x81\xe4\xc9\x9c\x53\x4a\xa6\x96\x5c\xd0\xdf\xa1\xce\x3e\x8d\xf6\x9b\x09\x8d\x04\x97\x7c\xae\x38\x8b\x29\xd3\x47\x57\xe2\x2c\x57\x1b\xe0\x14\x18\x36\x2b\x6d\xb4\x8c\xfd\x1f\xb5\x0a\x92\x8d\x6d\xf9\x10\x52\xfc\x1a\xd8\xf7\x63\x66\xc8\x6d\x30\x0a\x51\xb4\xc4\x71\x0c\x6c\xe1\x97\x79\xf7\x34\xe5\x21\x8f\x70\x8c\x4c\x53\x95\x0b\xb2\xbf\x41\xcf\x77\x9d\x58\x45\x87\xd6\xd7\xd4\xbb\xa9\xed\xf2\x8e\x1c\x9f\x1a\x9d\xb9\xa3\xab\xeb\x25\x82\xff\xd3\x71\x3b\x97\xe0\x61\xbc\x8b\x9b\x14\xc3\xf5\xfc\xef\x67\x17\x55\xda\x12\x36\xda\x4a\xc5\xd0\x07\x7c\xd3\x5d\xc0\x44\x1f\x5f\x44\x9f\x7e\xf9\x01\xea\xa6\x6d\x14\x97\x92\xf9\x83\xd6\x53\xe4\xee\x44\xc2\x82\x85\xd2\xc2\x99\x0e\x57\xe0\xcc\xd4\x17\x41\x27\x00\x52\x4e\xf5\xf0\x86\x18\xaf\x5e\x1e\x3b\x39\x0a\xcb\xaf\x03\x7b\xd1\x6e\x07\xa9\xe0\xbf\x41\xe4\xc5\x61\x57\x18\x5c\x70\x02\x13\x53\xd1\x72\xd1\x41\xe6\x49\x38\x61\x2e\xfe\xfd\x3a\xa9\x15\xf1\x24\xcd\x14\xe4\x29\xaf\x84\x28\x13\x54\xad\x75\x61\x19\x69\x45\xba\xc0\x1e\xc9\x62\x64\x8c\x8d\x0f\xc8\x76\xa7\x75\x94\x4f\x0e\xf1\x0c\x62\x39\x32\x1a\xb0\x1d\x93\x17\xb6\x12\xa7\x64\x13\xef\xb0\x9d\xff\x0b\x0f\x5f\xe7\xff\x5a\x66\x34\x10\x3c\xd3\x09\x79\xb9\x14\x99\xcd\x08\x4f\x30\xb5\xce\xfd\x17\xca\x08\xdc\x3c\xde\x08\xc7\x3c\x2b\x8d\x45\xa2\xf6\x5f\x9d\x59\x4e\x72\x6e\xbe\xed\xc8\xca\x05\x4d\x29\x5a\x82\x19\x5e\x00\x29\x3a\x1d\x61\xbe\xef\xe6\xb3\xe9\x24\x19\x57\xd3\xe3\x69\xcc\xd7\xbb\xfc\x2e\x2d\xae\x86\x2a\xa5\x7f\xed\xc5\x0e\x42\x69\x7e\x49\xa5\x81\x1d\xdf\x5b\x2e\xa2\xa4\x6d\x49\xe4\x39\x75\x6d\x47\xe0\xf8\xb8\xbe\x21\x50\x93\x80\x7b\x57\x24\x7e\x43\xa8\x58\xcb\x66\x2e\xee\xf5\xfd\xf5\xda\x6f\xc6\xb6\xad\x24\x07\xec\x3c\xa6\x8b\xa5\xb2\x06\x5c\xb4\x9b\xa6\x34\x01\x9e\xa9\x4d\x5f\x34\xcf\xad\xf8\xf7\x85\xc6\x26\x43\x4f\xbc\xbd\x9e\xb8\xa9\x76\x58\xf6\x7a\x3e\xa6\x48\x02\x8a\xc3\x2d\xbc\xf3\x74\x7d\x5c\x14\x5c\xf1\x38\x4b\xbc\x9e\x02\x59\x33\x9c\xd0\xc8\x84\x6e\x1d\x52\x28\x5b\xf4\x19\x9e\xc5\x40\xdc\xa9\x63\xfb\xff\x3b\xfa\xec\x75\x71\x08\xbd\x79\xf3\xa4\x7f\x79\xfe\x64\xb3\x01\x68\xcf\x8e\x49\x25\xc0\x05\x0e\xb1\x13\x84\x26\x0e\xe9\x73\xc4\x96\x11\x95\x2e\x47\x05\xa9\x9b\xa9\xe5\xe6\xd5\x5f\x25\x14\x03\x8b\xc4\x3a\xbd\x9d\x48\x9f\x45\xb7\xd0\xb8\x7d\xcd\x95\x22\xae\x58\xed\xd7\x00\xa1\x46\xb9\xe4\x46\x07\x35\x56\x87\x8d\x67\x7a\x54\xaf\x5c\x7f\xb7\xad\x18\x3b\x96\x0a\x20\xd6\xd8\x1a\x1d\xf4\x0f\xa3\xfc\xaf\xce\x04\x1a\x5a\x67\x1a\xfe\x82\x7f\x32\xda\xfa\x4f\xce\x8c\xc6\x62\x1a\x29\xdb\xac\xfe\xf6\xac\x1e\xe3\x03\xbe\xe9\x9f\x4e\x3e\x39\x1d\x67\xec\x6e\xf0\xb7\xbd\xfe\xe8\xec\x3e\x08\x26\x5b\x3e\xa3\xf2\xfa\x1e\x48\x2a\x5a\x0e\x98\x8e\x88\x9c\x74\xe7\x73\xca\xa8\x5a\xdf\x8e\x72\xc1\x35\x87\xfd\xd6\xfc\x16\x18\x08\x9d\x64\x16\x1b\x7a\x2b\xf8\x88\x93\x29\x8f\x41\x68\x48\x7d\x82\x4d\x31\x65\xea\x0e\x9c\xde\x12\xa2\x6b\x0d\xfc\x01\x12\x2e\xd6\x23\x1d\xdf\x32\x01\x7b\x22\xe9\xa5\xdc\x03\xc5\xee\xeb\x29\x35\x4f\xdf\xd7\xc3\x63\xd7\xca\x69\x74\x8a\x31\x84\x1a\xf9\x41\x92\xef\xb0\x3f\x89\x50\x23\x36\x47\x68\x61\x6d\xc5\xb8\x80\x85\xb6\x57\x6f\xf0\xd7\xe2\x73\x9e\x01\x3b\x01\x3c\x51\xc7\x1e\x96\x86\xfa\xd5\x59\x35\xe5\x82\xea\x30\xb5\xdb\xaa\xf3\x94\x61\x92\x0a\xc0\x64\x64\x31\xac\x57\x18\xb8\x2f\xa0\x23\x77\xa3\x83\x0e\x6f\xdd\xab\x4d\x83\x7a\x30\xa1\x21\x60\xa9\xdc\xc1\x01\x0f\x97\xe7\x14\xc7\x98\x45\x40\xf2\xeb\x21\x97\x96\xe8\x4d\xba\x2f\x29\x6d\x34\x23\x53\xfc\x75\x57\x9c\x92\x11\x27\xf2\x36\xb1\x4c\x66\x73\x17\xbd\x47\x6f\x93\xf1\x13\xe7\x39\x94\xb3\x7b\xd3\xb9\xdd\x66\x99\xa2\xb7\xda\xad\x96\xe0\x77\xce\xbc\xbb\xba\xdd\xa6\xa9\xc3\x65\x8d\x58\x47\xa5\xad\x06\x2e\xc4\xdb\x97\x4b\xcc\xf9\x6b\x12\x09\x74\xf8\xf2\x55\xf3\xe5\xf3\xe6\xe1\xd1\xeb\xe6\xe1\xcb\x27\x35\x4f\x52\x0b\x90\x3c\x5e\xd9\x46\x68\xed\xd3\xd4\x95\xb6\x69\xf9\x0e\x90\x79\x81\x86\x4a\x44\x32\xdb\x28\x03\xf2\x0c\x61\x46\xcc\xdb\x40\x4f\xa4\x79\x63\x86\x12\xc0\xf1\xf6\xa9\xb3\xab\x0f\x5b\x1c\x3c\xca\x96\x97\xa4\x72\xd8\x15\x35\xa7\x9f\x68\x66\x33\xfb\x0a\x0f\xe5\x6c\x03\x7c\xe2\x4f\x55\x90\x30\xce\x2f\x6d\xc8\xf7\x6d\x2c\x14\x84\xbf\x7f\x57\x01\x63\x32\xdd\x77\x53\xf2\xbb\x59\x73\xb9\x57\x81\x1e\xfb\x33\x3e\x4a\xee\xd7\x3b\x44\x1e\xba\xe9\x6a\x8e\x6f\xab\x1e\x43\xcb\x56\x7c\x4c\x2e\x42\x57\x1b\x05\xa9\xa0\x09\x16\xeb\x49\x84\x63\x98\x80\xca\x4b\xc2\x62\x7e\x95\xd8\x36\xeb\x2a\x91\xd2\x5a\x6e\x6e\x57\x94\xad\x40\x2a\xba\xc0\x0a\x90\x5a\x02\x0a\xc3\x04\x33\x3a\x07\xa9\xc2\x4c\xc4\xc8\x3d\x10\x89\x52\x2c\x70\x02\x0a\x84\x31\x3a\x09\x80\xe8\x1c\x51\x85\x12\xed\x1b\xc1\x01\x5a\x42\x9c\xa2\x4c\x22\xac\x10\x8e\x6b\x8c\xd0\xd8\x7e\xca\x89\xb4\x4f\x48\xf9\x0f\x54\xd4\x25\x79\x23\x4e\x82\x04\x14\x26\x58\xe1\x4e\x50\xdc\xaa\xab\x88\xb8\x2f\x32\xc5\x11\xd8\x1b\xb5\xd0\xbe\x46\x17\xc8\x14\xa2\x8e\xbb\xb0\x30\xfb\xe9\x8a\x0f\x2c\x16\xc5\x05\xeb\x1f\xce\xa9\x25\x28\x14\x62\xf7\xa5\x89\x6a\x9e\xdf\x72\x73\x70\x03\x51\xce\x57\x93\x4e\x12\x5c\x5e\xb2\x98\xb7\x32\xe5\xd2\x7d\x0b\x23\xf3\xc1\xdc\x99\x55\xec\xc0\x56\xfe\x7d\x15\x11\x73\xd9\xb6\xd9\x1e\xb2\x97\x6c\x59\x1c\xbb\xe7\x1f\x50\x37\xfe\x82\xd7\xf6\x8e\x3f\xa6\x2b\x60\x20\xe5\x48\xf0\x59\xd1\xbd\xd3\x32\x95\xad\xac\x8a\x4c\xa8\xb8\xfc\x89\x54\xec\x8d\x84\x61\x84\xcd\x05\x88\x37\xb6\xf5\xa0\x4e\x05\x3c\xbf\x6f\xa9\x45\x28\x1e\x4b\xf2\x51\xf2\x3b\x96\xdd\x18\x45\xc1\xe2\x30\x80\x91\x94\xeb\xe4\xc9\x1b\xdd\x79\xd5\x5a\x82\xe4\x97\x3b\x4b\xc0\xb1\x5a\xba\x09\x1d\xf9\x29\x8e\xcf\x20\xc6\xeb\xa2\x3c\x3b\x7e\xe1\xb5\xb6\x0a\x2d\xe6\xbe\x64\xae\x35\x6e\x8a\xe7\x5b\x74\x19\x4b\x63\x58\x14\x85\x8b\x1e\xb4\xd5\xce\x07\x53\xb6\xe6\x6a\x37\xef\xff\x8d\xb0\x5a\x76\xca\x05\x3a\x1a\x25\x27\x77\x31\xe6\xc6\x75\xa2\x72\xc9\xe2\xb5\x47\xb9\x4a\xa7\xf2\x94\xde\x16\x2d\xed\x06\x66\x54\x17\xbf\x94\x2d\xce\xa8\xd8\xc6\xd1\xfb\x55\xd6\xc7\x96\x8d\x15\x3f\xbf\xc5\xe6\xd2\xb2\xcb\x17\x5c\xb3\x84\xba\x05\xec\xc4\xdc\x12\x7a\x53\xe4\xfa\x32\xa8\x8c\x05\x38\xa5\x7b\x54\x7c\x9b\xc1\xc0\xe6\xa1\x56\x96\xca\xe3\x3a\x3a\xda\x71\x06\x4c\x75\x10\x4e\x69\x11\x36\xca\xcf\x0f\x8c\x1a\xe6\xed\xdc\x5c\x69\xc6\x99\xdd\x17\x4b\xd9\x79\x8c\xd9\xae\x93\xbd\x9e\x25\x2b\x90\x62\xbe\x88\x61\x05\xf1\xc9\x71\x5d\x7c\x29\xdf\x8f\xae\x0f\x2c\x3d\x7b\x35\x3b\x8a\x31\x83\xef\x13\x5a\xb4\xf7\xbd\x2d\x1f\xf9\xca\x35\x6d\x3d\xed\xf7\x72\xd4\xdc\xa3\xdb\xb6\x8c\x0b\xa8\xd1\x12\xf4\x5e\xbf\x9b\x4e\x47\x93\x3d\x5c\x12\x21\xb5\xd1\x4a\x39\x6c\x7b\x26\x94\x6f\xac\xf6\x1b\x7a\x4f\x29\x5b\x1a\x69\xfd\x3d\x64\x75\x22\xed\x94\xf5\x7b\x07\x92\x8a\xd9\x54\xa2\x40\xc5\x84\xf6\x8d\x29\x3b\x93\xca\x3a\xca\x95\xfb\xf8\x5d\x1c\x1e\x13\x62\xb6\x17\x57\xbf\xb4\x7d\x88\x6c\x2f\xe7\xb6\xc5\xdc\x15\x85\xbc\xe7\x96\x1e\x10\x8d\x2c\xe7\xea\xe3\x56\xff\x8c\x58\x53\xe5\xf0\x98\x98\x43\xa5\x02\xb6\xf5\x8c\xf2\xf1\xf1\xf1\xff\x95\xb0\x74\x7c\x8b\xab\xd7\xa9\xeb\xff\x3d\xf9\x7f\x93\x27\xe7\x3f\x7b\x20\xcc\xef\x1e\xfc\x05\x3d\xb5\xa5\x77\x07\xfd\xb5\xf9\xf4\xe0\xbf\x0e\x6b\x12\xc8\xfc\xa7\x0f\xf6\x7f\xad\xe3\xd1\x3c\x6a\xdf\xd1\x3e\x40\xef\xba\xbd\xf7\xda\xfa\xd2\x35\xda\x98\x34\xbf\x8a\xc1\xb9\x92\x4a\xe0\xd4\x1f\x97\xdc\xfe\xe2\x47\x21\xae\xfb\xb1\x10\x8d\x8f\x32\x2d\x28\x65\xa6\x72\x93\x6b\x16\x05\x07\x88\x60\x48\x38\xd3\x05\xce\x17\x1a\xc7\xa6\x31\x30\xc7\x34\x76\x75\x9b\x32\xa0\x76\xc1\x96\x84\xcd\xa6\x50\xc4\x85\x80\x48\xc5\xeb\x66\xed\x13\xf7\x9b\xd2\x6e\x01\xd4\xc9\x1e\xd8\x9f\xf2\x80\x0e\xaa\xfb\xa9\x09\x14\x09\x2c\x97\x28\xe6\x3c\x95\x28\x63\x8a\xc6\xb9\x5c\x54\xa2\x2c\xf5\x7e\x1d\x05\xcc\x7d\x42\x2d\x91\xe2\xc7\x52\x36\x7f\x4b\xe5\x36\x60\xf4\xef\xfe\x9b\xec\x82\x73\xd5\x32\x52\x6f\xae\xfc\x6e\x43\xf1\xb1\x5b\x6e\xd1\xff\x1d\x00\x00\xff\xff\x9e\xdc\xed\xe6\xae\x47\x00\x00")

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

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x71\x4f\x23\xb7\x13\xfd\xdf\x9f\x62\x7e\xf0\x53\xd3\x9e\xba\x6b\xb8\x9e\x50\xcb\x01\x12\x84\x80\x10\x1c\x89\x48\x7a\xd2\xa9\xaa\x90\xd7\x9e\x10\x2b\xbb\xb6\x6f\x3c\x1b\xb1\x05\xbe\x7b\xb5\xbb\x21\x24\x81\xe3\x68\xfb\x5f\x34\xf3\xe6\xcd\xdb\x19\xfb\x39\x9b\xff\x93\x99\x75\x32\x53\x71\x02\x09\xde\x0a\xb1\x09\xa3\xfe\x71\x7f\x17\x24\xb2\x96\xc6\xc5\x42\xc5\xaf\xa9\x91\x9e\xec\x8d\x75\x49\x19\x22\x13\xaa\x22\x31\x2e\xa6\xda\xbb\x31\xd8\x08\xba\x24\x42\xc7\x79\x05\x13\x45\x46\x7b\x83\xe6\x23\x58\x16\x9b\x10\xc8\x67\x2a\xcb\x2b\x88\x13\x5f\xe6\xc6\x75\x18\x32\x14\x62\xd8\xbb\xfa\x7c\xd6\xed\x5d\x8f\xbe\x0c\x7a\xfb\x2d\xb3\xb0\x63\xf8\x03\x92\x31\x6c\x34\x8d\x63\x15\x6b\x76\x7b\x23\x15\xfb\xc2\xea\xc4\x07\x74\x71\x62\xc7\x9c\x38\x6f\x70\x03\xfe\xfc\x08\x3c\x41\x27\x00\x00\x56\xe8\xd6\xf1\x62\x6c\x85\x28\x0b\x5f\x3a\x06\x59\x38\x96\x84\xd1\x97\xa4\x11\xee\xef\x81\xa9\x44\x51\x4c\xc7\x31\xbd\x1d\xc7\xba\xbb\x34\x38\x93\xc6\xc6\xa9\x54\x7f\x95\x84\x0b\x70\x12\x14\xf1\xb6\x40\x3d\xf1\xd0\x79\x1d\x04\x20\x67\x8a\x64\x6e\x33\x69\xbc\x9e\x22\x01\xd4\xe4\x70\x43\xe1\x6b\xe9\x59\x01\x6c\xc1\x56\x07\x0e\x0e\x9a\x0f\x1d\x47\x56\x99\x88\x55\x64\x2c\x34\xe7\x10\xd9\x07\x68\xeb\xd2\x88\x34\xb3\x1a\xc5\x5c\xfc\x2a\xab\x20\x8c\xec\x09\xb5\x77\x90\x5c\x3d\xcb\x2e\x33\x2a\xe2\x75\xca\xf9\x97\x1c\xf5\xfb\xa3\xe1\xe8\xea\x70\x70\xdd\xed\x5f\x9e\x9c\x9d\x5e\x5f\x1e\x7e\xea\xed\xd7\x13\x4e\xda\xf1\x27\x77\x77\x90\xf6\x6e\x99\x54\x7a\xe5\x73\x84\x87\x87\x85\xf2\xa7\x15\xfd\xff\x6e\x79\x03\x0f\xcd\x86\x84\x88\x68\x20\xb1\x90\x20\x6c\xc4\xcd\xe3\xde\xd1\xef\xa7\xd7\x17\xfd\xd3\x8b\xde\xe7\xde\xc5\xfe\xfb\xf5\xc0\x87\xcd\x0d\x78\x13\x2b\x15\x90\xd0\xb8\xc5\x22\x6b\x23\xdf\xb5\xbf\xdb\x53\x24\x0b\x15\x19\x49\xbe\x13\x22\x53\x11\x77\x3e\x40\x62\x60\x6f\x6f\x0f\xee\xee\xe0\xa8\x09\xf4\x5c\x7d\x3e\xe1\xc7\x2f\xaa\xc8\x3f\x29\x8a\x13\x95\x43\xda\x6d\x3a\xa6\x97\xde\xe0\x91\xf7\x1c\x99\x54\x38\x2f\x33\x6c\x95\xfc\x04\x0f\x0f\x70\xb0\xdc\xa5\x96\x22\xb3\x47\x64\x3a\x5d\x40\xbf\xd7\xb5\x8b\xc4\x87\xf1\xa8\x62\x8c\x8b\xae\x75\xcc\x8e\xad\x56\x8c\x71\x55\x42\x93\xfa\x46\xf7\x66\x47\x0b\x09\x01\x29\xd5\xc4\xdf\x6b\x3f\x20\x3b\x53\x8c\xe7\x58\xfd\x03\x11\xe7\x58\xbd\x59\xc3\x14\x2b\xa1\x27\x85\x37\xb0\xb5\xb3\xb5\x05\x6f\xab\x78\x0e\x7b\x71\xb4\xff\x79\xb6\x5d\xf5\xda\x40\xb5\x6a\x26\xa8\xc3\x73\x39\x6d\xaa\x8d\x87\xa9\x95\x5a\x25\x4c\x65\x64\xd9\x5e\x7b\xa9\x9c\x9e\x78\x8a\xf2\xc9\xa3\xe6\x64\x65\x30\x8a\x31\x79\xc4\x3f\xde\x3a\xa7\x0a\xac\x2f\x22\x12\x6c\xef\xfc\x9a\xee\xfc\x92\x6e\xbf\xff\x2d\xdd\xde\xe9\xbc\x20\xab\xf6\x96\x7c\xd6\x58\xad\x28\xa6\xc6\x12\x24\xab\x0a\x75\xee\x4b\x13\xc8\xcf\xac\x41\x7a\xb2\x6f\x9e\xd8\x58\x7b\xb3\x29\x43\xde\x7c\xbf\xf9\x19\x94\x33\xc0\x13\xc5\x9d\x08\xce\x33\x58\x83\x2a\x17\x5a\xf1\x6a\xdb\x15\xc2\xd6\xe1\x5a\xa7\xdf\xdb\xeb\xf4\xfa\x27\x1d\xc1\xe8\x94\xe3\x33\xb3\x5b\xef\xe0\x71\xd2\xa3\x36\x78\x0c\xf7\x50\xbb\x5c\xed\x14\x22\x96\x59\xd4\x64\x03\x5b\xef\xd6\xe0\xc3\xe5\xd4\x4a\x91\x52\xa6\x9b\x5b\x5c\x69\xc0\xca\x3a\xa4\x61\xeb\x5d\xe9\x80\x7c\xa8\xf7\x8a\x31\x9d\x87\x06\x64\x9d\xb6\x41\xe5\x03\xf2\x63\x9b\x63\x3a\x67\x78\x99\x78\x88\x9a\x90\xff\x3d\x79\x5b\xbf\x46\x3d\x7a\xeb\x50\x1e\x1f\x8b\x53\xf2\x65\x58\x41\x5f\x2d\x67\x96\x4b\x72\xaf\x55\x3d\xa8\x6f\x48\xbe\x98\xa7\x57\x46\x8f\xba\x24\xcb\x55\xc3\x75\xa9\x0a\xdc\x05\x17\x6f\x12\xed\x8b\x50\x32\x8a\x40\xb6\x50\x54\x0d\xb5\xca\x71\x88\xdc\x02\x62\x5c\xe4\x67\xc5\xa8\x0a\xb8\x0b\xb3\x22\x46\xd1\xeb\x9f\xd4\xe7\xca\x79\xc6\x5d\x78\xc9\x95\x41\x53\xfd\xf7\x21\xf7\x3e\x44\x28\x1d\xdb\x1c\x5a\x1f\xae\x0f\x60\x19\x96\xde\x22\x74\x2a\xcb\xf1\x45\x92\xc5\xd3\xb4\xfe\x72\xbd\x06\x86\x1f\xc4\xdf\x01\x00\x00\xff\xff\xe5\x33\xaa\xbb\xc1\x08\x00\x00")

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
	"azuredeploy.json":  azuredeployJson,
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
	"azuredeploy.json":  {azuredeployJson, map[string]*bintree{}},
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
