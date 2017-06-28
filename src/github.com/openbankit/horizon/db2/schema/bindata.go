// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_aggregate_expenses_for_accounts.sql
// migrations/7_account_limits.sql
// migrations/8_account_limits_two_way.sql
// migrations/9_1_assets.sql
// migrations/9_2_options.sql
// migrations/9_commission.sql
// DO NOT EDIT!

package schema

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5c\xeb\x6f\xe3\xb8\x11\xff\x9e\xbf\x82\xb8\x2f\x4e\x50\x27\x8d\x37\xd9\x3c\x71\x07\xf8\x12\x5f\xd7\x68\x62\xef\xc5\x4e\x77\x17\x45\x21\xc8\x16\x6d\xab\x2b\x4b\x5a\x49\x4e\xe2\x2b\xfa\xbf\x77\x48\xbd\x28\x89\x2f\x3d\x72\xbd\x2f\xb7\xb1\x86\x3f\xfe\x66\x38\x1c\x0e\x87\x94\x8e\x8f\x0f\x8e\x8f\xd1\x67\x2f\x8c\xd6\x01\x9e\xfd\xfe\x80\x2c\x33\x32\x17\x66\x88\x91\xb5\xdb\xfa\xf0\xec\x80\x3c\xbf\x87\x7f\x63\x0b\xad\x02\x6f\x9b\x0b\xbc\xe0\x20\xb4\x3d\x17\x5d\x9f\x7c\x3c\x39\x65\xa4\x16\x7b\xe4\xaf\x0d\xd2\xbc\x24\x72\x30\x1b\xcd\x51\x18\x99\x11\xde\x62\x37\x32\x22\x7b\x8b\xbd\x5d\x84\x7e\x46\xa7\xb7\xf4\x91\xe3\x2d\xbf\x57\x7f\x5d\x3a\x36\x91\xc6\xee\xd2\xb3\x6c\x77\x0d\x0f\x7a\xcf\xf3\xdf\xae\x7a\xb7\x29\x9c\x6b\x99\x81\x65\x2c\x3d\x77\xe5\x05\x5b\x90\x30\xc2\x28\x80\xff\x85\x20\xe9\xb9\x09\xc6\x06\x03\xf4\x6a\xe7\x2e\x23\xa0\x63\x2c\x00\x09\x93\xe7\x2b\xd3\x09\x71\xa1\x1b\x00\x30\xb6\x38\x0c\xcd\x35\x15\x78\x35\x03\x17\xb0\x62\x91\xc0\x7b\x35\x42\xbc\xdc\x05\x76\xb4\x27\xe0\xab\xd5\x6d\xa2\x13\x36\x83\xe5\xc6\xf0\xcd\x68\x03\xbf\xfb\xbb\x85\x63\x2f\xfb\xc4\x08\x4b\xb0\x95\xe3\x41\xf3\x83\xfb\xa7\xe9\x67\x34\x9e\xdc\x8f\xbe\xa2\xf1\x6f\x68\xf4\x75\x3c\x9b\xcf\x12\xc9\x93\x28\x30\x2d\x6c\xe0\xd5\x0a\x2f\xa3\xd0\x58\xec\x0d\x2f\xb0\x70\x00\x2c\xbd\xef\xb7\xd2\x86\xb6\x6b\xe1\x37\x63\x63\x87\x91\x17\xec\x0d\x80\x71\x43\x93\x6a\x18\x1a\xa0\xa5\x6d\xd5\x69\xed\xf9\x38\x30\xb3\xb6\xd1\xde\xc7\x2d\x5a\xe7\x4c\x5a\xb1\xa8\xd7\xd6\xc1\xd6\x1a\xfc\x8d\x34\x0c\xf1\x8f\x1d\x38\x4c\x2d\x15\x98\xe6\x7e\x80\x5f\x6c\x6f\x17\x26\xbf\x19\x1b\x33\xdc\x34\x84\x6a\x8f\x60\x6f\x7d\x2f\x88\x00\x23\x99\x4c\x4d\x61\x9a\xda\x72\xe9\x78\x21\xb6\x0c\x33\xaa\xd3\x3e\x75\xe6\x06\xae\x64\x2e\x97\xde\xce\x85\xb6\xaf\x76\xb4\x21\xae\x64\x47\x61\xa3\xf6\xb5\x95\x66\x5b\x9a\x96\x15\x40\x18\x90\x37\xdf\x44\x3e\x99\xae\x9b\x48\xd5\xcf\x26\x2c\xcc\x09\x68\xa3\xd1\x22\x71\x1d\x1d\x61\x2f\xe6\xe1\x29\x05\x41\x53\x23\x7a\x33\x7c\x35\x24\x91\x04\x58\x4d\x49\xac\x2b\x96\x46\x37\xb9\xf0\xd2\xdb\x6e\xed\x30\x4c\x6c\xa5\x9e\x3c\x45\x79\x33\x0c\xb1\xc2\x5b\x4b\x0d\xe2\x81\xd7\x70\x55\x6e\x3b\x79\x93\x45\x3a\x9b\x94\x62\x6a\x3d\x75\xfb\xa4\x16\x08\x61\x4d\x84\x75\x05\xe8\xee\xc0\x8d\xd4\xba\xa5\x56\x20\x2b\x34\x0c\x96\xbd\x0c\xd3\x59\x00\x83\xfb\x76\x7b\x30\x7c\x98\x8f\x9e\xd0\x7c\xf8\xeb\xc3\x88\x69\x3c\x9d\x3c\x7c\x63\xc7\xb8\xb4\x12\xc1\xa2\x18\x00\x94\xed\x9b\x30\xb1\x10\xed\xfe\x6e\x3a\x99\xcd\x9f\x86\xe3\xc9\x9c\x81\x51\x35\x35\xfc\xef\x78\x5f\x87\x43\xb6\x92\xd4\x65\xc0\x6f\xa8\xdd\xff\xda\x0b\x7c\xc8\x22\xd6\xc9\x32\x26\xe9\xb0\x24\xa9\xdd\x43\xee\x83\x12\x70\xc6\x51\x75\x71\xa9\xd3\x48\x20\xe9\x73\x7d\xb4\x8a\x37\xc9\xa0\xab\xae\x57\xb7\x1f\xc7\xde\xda\xd2\xf1\x2d\x0a\x4a\xf1\x75\xdd\x39\x6e\x7d\x37\x7d\x78\x7e\x9c\x20\xdb\x8a\x3b\xbf\x1f\xfd\x36\x7c\x7e\x98\x6b\x62\x0b\xdc\xb4\x05\x32\xe3\x1e\x2d\x50\x62\x67\x90\x03\xd0\xbf\xf4\x6d\x97\x2e\xa6\xb3\xd1\xef\xcf\xa3\xc9\x5d\x03\x83\x43\x1c\x22\xa9\x5d\xed\x9e\x0b\x20\x7a\xad\xf3\x44\x54\x9b\xb5\x20\x70\xd4\xe1\xcc\x87\xd0\x6b\x9b\xa4\x6c\x7a\xc2\x49\x7e\xa6\x27\x9c\xe6\x45\x72\xe9\x52\x38\x53\x9a\x8d\x89\x50\x3a\x26\xca\xc5\x95\xc8\x71\xa0\xd2\x01\x65\x33\x05\x91\x48\x25\x34\xe9\xc9\xc7\x61\x26\x91\x1d\x7d\x9d\x8f\x26\xb3\xf1\x74\xc2\xae\x37\xc4\xb8\x58\x22\xe0\x3b\xfe\x3a\xfc\xe1\xa4\xea\xde\x7d\x1a\x3d\x0e\x2b\xfd\xdd\x1e\xc4\x3b\xf8\x89\xb9\xc5\x37\xe9\x6f\x68\x0e\x8b\xfd\x4d\xd2\xe4\x16\xcd\x60\xfb\xbb\x35\x6f\xd0\xf1\x2d\x9a\xbe\xba\x38\x80\x7f\xd1\x8d\xfd\xdd\xd3\x68\x38\x1f\xa5\xc8\x29\xde\x41\x11\x31\x21\x91\x40\x66\x3c\x95\xa8\x05\x8d\x26\xd3\x79\x49\x2b\xf4\x65\x3c\xff\x94\x75\xcd\xee\x94\x0b\xdd\xe7\x28\x25\x22\x77\xd3\xc7\xc7\xd1\x64\x2e\xa1\x11\x0b\xc0\x5a\x51\x05\x41\xe3\x19\xea\x7d\x7e\xf8\xab\xbf\x26\x15\x0f\x3f\xf0\x96\xd8\xda\x05\xa6\x83\x1c\xd3\x5d\xef\x60\xeb\xdf\x2b\xf3\x48\x06\xab\x33\x2b\xc4\x78\x45\x23\x70\xed\x9f\x03\x14\x29\x34\xd3\x3f\xe9\x96\xa8\x4f\xca\x38\x88\x24\x85\x68\xe5\x05\x88\xfc\x4e\x8a\x2b\x24\x6d\x44\xde\x0a\x1d\xc2\xea\xd8\x47\x2f\xa6\xb3\xc3\x47\xc8\x37\xed\x20\xa4\x26\xd1\x2c\x76\x10\x31\x0b\xaf\xcc\x9d\x03\x39\xb5\xb9\x70\x70\xe8\x9b\x4b\x4c\x2a\x37\xbd\xd2\x53\xba\xc7\x83\x6d\x0b\x53\x8c\x29\xa8\x5f\x9a\x4d\x89\xf2\x74\xea\xe5\xaa\xa7\x5e\xcf\x1b\x80\x78\x96\x96\x92\x84\xc3\x03\x04\xff\x25\xc9\x2d\x5a\x6e\xcc\x00\xd6\x09\x1c\x80\xbe\xc1\x1e\xac\x70\x78\x71\x7e\x44\x07\x6b\xf2\xfc\xf0\xd0\x8f\x65\x69\x48\x21\xf9\x34\x47\x7c\xf0\xa1\x2c\xbe\x35\xdf\x98\x58\x4e\xca\x59\x0b\x7b\x6d\xbb\x51\xba\x76\xa2\xd3\x52\x03\xcb\xb4\x9d\xbd\x41\x9b\xa9\x85\xb7\x9e\x1b\x6d\x6a\x88\x17\xc8\xd8\x6e\x59\xbe\x77\x3c\xe8\xdd\xdc\xc0\x2f\x18\xd6\x0f\x21\xaf\x7a\xed\x58\x8a\xba\x2d\x0f\x8e\xca\xce\xcf\x89\xbd\x6d\x3d\x80\x49\x47\xdf\xdd\x0b\x68\x8f\x38\x20\x4b\xf9\x9e\xee\xbf\x50\xb8\x35\x1d\x47\xed\x07\xb6\x0b\xab\x1d\xd6\xf3\x19\x70\x00\x1d\xe1\x57\x8c\xbf\x6b\x23\x27\xc2\x9a\xd0\xe9\x58\xeb\x61\xa7\xd2\x9a\xe0\xa6\xeb\xee\x4c\x47\x13\x3b\x11\xd6\x84\xde\xf9\x10\x03\x69\x65\x0b\x91\xa2\x33\x78\xc6\xd6\x47\x24\x20\xd1\x3f\xd1\x1f\x9e\x8b\x65\xbe\x49\x53\x87\xc6\xee\x48\xd3\xeb\xd8\x03\x21\xaf\x4e\x98\x16\xf9\x51\x8f\xe1\x4f\x2f\x6d\x17\x8c\x37\xff\x5a\xce\x6d\xc3\x76\xdf\xf5\xdc\xfd\xd6\xdb\x85\x68\xe1\x79\x0e\x36\x5d\x95\xfe\x69\x92\x95\x26\x1c\x49\x4a\xa6\x67\x89\x2c\x81\x63\xa1\x28\x95\xd9\x7c\xf8\x34\x8f\x17\xc7\x01\xfd\x61\x3c\x81\x36\x74\x39\xfb\xf5\x5b\xf2\xd3\x64\x8a\x1e\xc7\x93\x7f\x0c\x1f\x9e\x47\xd9\xdf\xc3\xaf\xf9\xdf\x77\x43\x58\x56\xd1\xa0\x0e\x6d\x34\xfd\x32\x19\xdd\x43\x17\x0a\xfe\xf1\xae\x88\x4b\x3f\x83\x88\x7f\x3d\x21\x55\xb1\x22\x01\x26\x8f\x6d\xea\x3c\xcc\x0e\x4f\xee\x41\xb0\x88\xd3\xa2\x52\x3e\xfe\x9c\x71\x27\x42\x74\xa1\x47\xff\x0e\x3d\x77\x51\x7a\xba\x72\xcc\xc8\x58\x61\xe5\x64\x82\xf5\x65\x49\xce\x4f\xa4\xa2\x55\x2f\xaa\x6e\x02\xda\xb9\x52\x05\xef\xbd\xfd\x49\xa9\x40\x43\xa7\xaa\xe0\xe6\x9e\x95\x3f\xe2\xb8\x57\x79\x17\xd6\xd4\xc7\xca\x65\xac\xcc\xd1\x22\xfc\x56\x76\x33\xd3\xf7\x1d\x5b\x1e\x48\xab\x23\x5f\xd9\x5c\x36\x65\x5a\x06\x52\xcc\x09\xe9\x7a\x9f\x88\x30\xe5\x60\x41\x00\x5e\xd0\xb3\x4a\xba\x2a\x91\x13\x47\xdf\xdc\x93\x23\xcd\x3c\x6e\xa6\xbe\x4f\x73\x5a\x6e\xdb\x78\x91\xaa\xdd\x98\x66\xb0\xc4\xd6\xb4\xc2\x1b\x4f\x59\xb1\x71\xd3\x6d\x7e\x5b\xdb\x26\x38\x89\x69\x4b\x16\x37\x44\xa6\xae\x56\x35\x44\x92\x3f\xd1\x33\x81\x9f\x04\xc6\x96\x8c\x83\x85\x23\xc8\x82\x94\x76\x48\x6b\x23\x6d\xed\x90\xe0\x24\x76\x48\x4f\x19\x05\xdc\x98\xa3\x3f\xad\x05\x98\x77\xea\x28\x73\x53\xb6\xc0\x45\x07\x22\xe3\x21\x0a\xd2\xf9\x40\xe8\xc9\x67\x47\x7f\xa5\x79\x4d\x36\x1d\xd5\x1c\x29\x69\x13\x60\x7e\x56\x55\x68\xa4\xc8\xc0\x38\xb2\x99\xeb\x24\x7f\x96\x4e\x45\x2b\xba\x0c\xca\x4e\xe4\xc1\xee\x14\xf4\xb6\x21\x98\x71\x7d\x10\x56\x2e\xc3\x87\x19\xc8\x7f\x4a\x6e\x3c\xd0\xc5\x4d\x10\x0f\xc8\x63\x88\x2b\x38\x78\x11\x89\x90\xad\x50\xf4\x66\x90\x5c\x21\xb4\xff\xa8\x4a\x89\xbd\x57\x50\x15\x6c\xeb\xcc\x82\xd2\x73\x16\x3e\xf9\x6a\xe8\x4f\x6a\x75\x98\xa8\xab\x72\x37\x39\x82\x56\x1f\xef\x9d\x37\x34\x52\xb4\x61\x2e\xa1\xd5\x57\x9e\x5f\xc8\xc5\x39\x39\x07\xa7\x66\xde\x99\x6f\xaa\x96\xf3\xe2\x55\x13\xc1\x92\x4f\xf2\x93\x65\x52\x8b\x21\x0b\x4d\xcb\x75\x26\xfe\x29\xf4\x76\x90\xf2\xa6\xde\x2d\x88\xf0\x59\xc5\xa3\x77\x73\x53\x91\xd0\x98\x07\xc2\x43\x8c\xb6\x06\x16\x9e\x69\x69\x4e\x7f\x1d\xbb\xb7\x09\x00\xaa\x23\xa0\x6e\x42\x80\xa2\x97\x3f\x2b\x08\xd4\x54\xb6\x65\x18\x50\xf4\x56\x0d\x04\xa2\x06\x92\x50\x50\x38\xf6\xeb\xd0\x57\x53\xff\x64\x29\x69\x27\x58\x49\x5e\xa5\x48\xdb\x74\xa3\x85\x7c\xe2\x73\x65\xf3\xae\xc5\x19\x88\x29\x9c\x7a\xa2\xec\xed\xff\x92\x7f\x41\x26\x83\xdd\x17\xec\x00\x29\xde\x96\x10\x1e\x43\x36\xb4\x73\x22\xc1\xc3\x2d\x26\xa7\x0f\xdc\x47\xc4\x0a\xa2\xc7\xa1\xbd\x76\xcd\x68\x07\xd0\x1c\xb3\x5f\x5f\x1c\xfd\xf3\x5f\x79\xc4\xfd\xcf\x7f\x79\x31\x17\x24\x4a\x69\x19\xde\x7a\xf1\x4e\xaf\x1a\x9f\x33\x2c\x17\xcc\x20\x8d\xe0\x39\x56\x15\x26\xd1\x0c\xcc\x69\x2c\x60\xe0\xac\x90\x8c\xdc\x15\x38\xf0\x9a\xb3\x2d\x86\x29\x95\x4c\x97\xf4\x98\x5d\x67\x8e\xc7\xf3\x85\x5e\x8b\xe0\x1f\xdc\x93\x33\x97\x54\x1b\x17\xec\xfa\x62\x3a\x87\x3d\xb6\x70\x05\xda\x05\x78\xbd\x74\xe0\xb7\xee\x39\x49\xae\x24\x70\x89\x55\x8a\x1f\xef\xca\xae\xe6\x55\x0c\x2e\x63\xad\x14\xeb\x4f\xd1\x42\xfb\xb2\x8a\x54\x0f\xc5\x1a\xc1\xd7\xe4\x9e\x1c\x29\x92\xd3\x44\xe5\xd9\x1d\xba\x1f\xce\x87\x0a\x0d\x15\xa8\x82\x33\xa1\x36\xc8\x95\x8a\x7e\x1d\x30\x8d\xf2\x32\x58\x5c\x01\x36\x1b\x3d\x8c\xee\xe6\xcc\x61\xea\x09\xc0\x55\xe7\x6a\x1f\x0d\xfa\x71\x75\x48\x6c\x7d\x41\x9d\xb9\xbe\x4a\xea\x0a\x67\x1b\xbd\xaa\x53\x5d\x47\x39\x59\x95\x53\x47\xc3\xf1\x64\x36\x82\xa4\x6e\x3c\x99\x4f\x2b\x95\x4e\x9a\xb5\xcd\xd0\x61\x6f\x60\xd8\xae\x1d\xd9\xa6\x63\x84\x14\xeb\x24\xfc\xe1\x00\xbb\xde\x87\xd3\xc1\xc5\xf1\xe9\xd5\xf1\xd9\x29\x1a\x0c\x6e\x3e\x5e\xdd\x7c\x38\x3f\x19\x9c\x5e\x0f\x2e\xaf\xff\x72\x7a\xd6\x03\xd2\x5a\xe8\x1f\x8c\xf8\xf2\x71\x61\x76\x2d\x60\xe6\x79\xb6\x25\xeb\xe9\xc3\xf9\xf5\xd5\x60\x50\xa7\xa7\x33\xc3\x5c\xaf\x61\xba\xc2\x52\x6f\xe0\x37\x1f\xbb\x21\x0e\x0d\xb0\x65\x56\x31\x95\x75\x77\x7e\x71\xf5\xf1\xf2\xa2\x4e\x77\x97\x46\x71\xe2\xcb\xd0\x3f\x9e\x0d\x4e\x2f\xaf\xea\xa0\x5f\x95\xd0\x8d\xe8\xd5\x33\x5e\xcd\xbd\xac\x97\x8b\xab\xb3\xc1\xe0\xbc\x4e\x2f\xd7\xc6\x20\xa9\xb0\xca\x70\x2f\x2f\x2f\xae\x2e\x2e\xeb\xe1\x32\xc5\x7b\x09\xf2\xf5\xc5\xf9\xd9\xc5\xc7\x04\x59\x30\x07\xa4\x05\xf4\x16\x61\x50\x56\x3b\xee\x00\x96\x57\x8a\xed\x00\x56\xa3\x46\x56\x3f\xf4\x35\x2b\xd2\xb4\x09\x87\x7a\x79\x84\x4e\x88\x54\x14\x65\x3a\x30\xb9\x56\x6d\xa2\xb9\xd1\xeb\x6e\x8a\xbb\x30\xbb\x2a\xed\xa9\x63\x78\xe1\x16\xb8\x41\x56\xc1\xb9\xba\x9c\xdd\x03\x4b\xaf\x3a\xd7\xde\x28\x14\x40\xe9\x1e\x65\x78\x7f\xcf\xde\x9d\xe6\x74\x8b\x3e\x3f\x8d\x1f\x87\x4f\xdf\xd0\xdf\x47\xdf\xd0\x61\x72\x96\xd6\x67\x2e\xc6\x68\xdc\xe1\xe9\x98\x7f\x0e\x2c\xd3\xa1\xd4\xbd\x52\x8f\x7e\xf5\xf6\x8e\xe0\x0a\x44\x47\xda\x10\x2c\xae\x02\x59\x27\x45\xce\xb6\x25\x3b\x4c\xef\x86\x54\x0e\xc8\x63\x56\xea\x4e\x49\x8f\xfb\xbe\x43\x6b\x8e\x25\x54\x1e\x51\x5e\xc7\x4a\xb6\x3a\xaf\x83\xb4\x26\x2f\xef\x84\xa7\x8b\x06\x2d\x6d\xd5\xe4\xef\xda\x74\xa6\x9c\xa8\x1b\x99\x7a\x52\x6a\x4a\x05\x15\x6f\x32\x25\x9a\xd1\xd7\xa0\xf4\xea\x91\xf1\x1b\x53\x72\x58\x72\xd3\x96\x73\xcb\xf0\x79\x36\x9e\xfc\x0d\x2d\xa2\x00\xe3\x2c\xd0\xf0\x23\x09\xe7\x7d\xad\xfa\x4c\x9f\x27\x63\x58\x0f\x53\xc2\x7c\x58\xca\x94\x96\x89\x0a\xe4\xe2\xb0\x17\xcb\xf5\x11\x37\xe2\x31\xef\x9f\x35\x35\x62\x0e\x41\x68\x70\x4b\xbc\x45\x93\xc5\xc2\xfd\x4a\x0d\x95\x47\x8e\xbe\x41\xd7\x82\x19\x2d\x25\x6b\xd1\x2a\x17\xa0\x79\x6c\x92\xd7\xfe\x5a\xf0\x89\x11\xf4\x18\x95\xaa\xdb\xfd\x6a\x21\x5b\xb6\x60\x74\x30\xb2\x5c\x34\xc2\x9d\x29\xff\x15\x18\x1f\x1e\xe6\x77\xcf\x8e\x7f\xf9\x05\xf5\xc8\xf7\x05\x7a\x37\x37\xa4\xf0\x7b\x74\xd4\x47\x95\xe7\x91\x97\x3d\xd5\xd3\xa5\xe9\x2c\x92\x28\x94\xcd\x20\xb1\x56\x3c\xb5\x68\xb3\x8c\x7d\x76\xf9\x99\x6a\x59\x55\x53\x24\xad\xd2\x9a\xad\x60\xb5\x55\x97\x06\x88\x3a\xa3\x17\x67\x2a\x05\xe6\x9c\x31\xcc\x53\x2c\xb5\x54\x1c\x8b\x74\xc7\xbc\xe1\xe4\x2f\x44\xcc\x2a\xa2\xcc\x04\xe9\xfd\x4a\xee\x0a\xcb\xbe\xec\xdc\x92\x55\x09\x8e\x8d\x07\xe9\xcd\xac\x02\x2f\xde\x1d\x8d\x7e\x7a\xc9\x4a\x44\x36\xaf\x40\xb7\xa4\x69\x5b\xda\x04\xf3\x93\xe1\x3e\xf7\x62\x89\x82\x74\xfa\x7e\x7a\x17\xbc\x13\x2c\x96\xba\xe0\x40\xa0\x91\x26\x7c\x05\xd2\x57\xf1\xbb\x50\x20\xc1\x12\x2c\x16\x0d\x55\x28\x1e\xf3\x57\x95\x60\x3e\x3c\xd0\x34\xec\x30\x18\x4d\x8d\x2f\x37\x74\xe9\x4b\x0a\x6d\x6d\x5d\x84\x63\x29\xa7\xb7\x03\x0b\x1c\xf9\x8c\xaa\x5f\x83\x68\x4f\xab\x82\xa9\x97\x37\xf0\x08\x32\xdf\xb5\x68\x3c\xac\x39\x46\x73\x97\x54\xb8\x9f\xfa\xf3\x1d\x2d\xad\xaa\xec\x80\x55\x2d\xbb\x80\xac\x95\xf2\x4b\x3f\x5a\xf2\x6e\xb4\x8b\x83\xc1\x67\xac\x6f\x68\xf6\x0b\x2d\x4d\xfd\x44\x0d\xad\xc5\x18\x7d\xf9\x34\x7a\x1a\x41\x22\x21\xba\x99\xfd\x33\x8a\x02\xf2\x32\xe3\xf4\x09\x1d\x0a\x6f\x60\x27\x42\x0a\xfd\xcb\x1f\xb7\xe9\x46\xf5\x12\xaa\x72\x0d\xe5\x6e\xd0\x34\xbe\xe2\xd3\x0d\x5b\x1e\xb4\x32\x16\x66\x92\xfa\xbc\xbb\x9e\x0c\x05\xe8\x26\xc1\x5b\xff\x3b\x4d\x9d\x1b\xba\x72\xe7\x59\x49\xbf\xd4\x40\x5f\x19\xf6\xb3\x55\xef\x65\x7f\xf6\x9a\xbb\x4a\x13\x46\x56\x5f\x09\xee\x67\xbc\xde\x4b\x1b\xee\xed\x7d\x95\x5a\xbc\x46\xfa\xfa\x65\x5f\x39\x7b\x2f\x9d\xb2\xeb\x6a\x2a\x3d\x84\x35\x19\xc5\xd7\xdd\x3a\x25\x5e\x46\xe7\x66\x93\x75\x27\xb8\xf4\xc3\x76\xdd\xcc\x70\x59\x17\x3a\x3a\xd4\x4a\x92\x38\x9f\xf9\x7b\x17\x2d\x4a\x2b\x98\x90\xbb\x7a\x11\xe3\x7c\xd6\xb0\x53\xb7\xa9\xe2\x37\xce\x9b\x65\x1f\x72\x6c\x6a\x65\x09\xa6\x32\x45\x38\x3c\x4c\xef\xad\xd3\xa2\x4a\xe8\x39\xc9\x8b\x63\xd5\x2a\x8d\x48\xb0\x52\xa8\x11\x09\x96\x6a\x35\x15\xd1\x85\xb7\x5b\x6f\x22\xad\xee\x0b\xa2\x72\x02\x05\xd1\x72\xb9\x28\xcd\x09\xa9\x33\xfe\x8c\xce\xce\x98\x01\x13\x7d\xf1\x94\xd4\x7a\x7c\x07\x47\x98\x8e\xc4\xff\x02\x00\x00\xff\xff\x19\x99\x73\x7a\x1e\x55\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 21790, mode: os.FileMode(420), modTime: time.Unix(1472739164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x6d\x6f\xdb\x38\x12\xfe\x9e\x5f\x41\xf4\x8b\x1d\x9c\x73\xd7\xa2\x87\xa2\x97\xa0\x05\xdc\x44\xb9\x1a\xe7\xc8\xad\x2d\x5f\x5b\x1c\x0e\x02\x2d\xd1\xb6\x36\xb2\xa8\x92\x54\x5e\xba\xd8\xff\xbe\x43\xea\xc5\x7a\xa3\x24\x27\x72\x76\xb1\xc0\x6e\xa4\xe1\xcc\x3c\x33\xc3\x67\x86\x94\xcf\xce\xd0\xdf\x76\xde\x86\x61\x41\xd0\x32\x3c\x39\x3b\x83\x7f\xd1\x17\xca\xc5\x86\x91\xc5\xd7\x29\x72\xb1\xc0\x2b\xcc\x09\x72\xa3\x9d\x7a\x7d\xb2\x30\x2c\xc4\x05\xc8\xef\x48\x20\x6c\xe1\xed\x08\x8d\x04\xfa\x80\x5e\x5f\xa8\x57\x3e\x75\x6e\xab\x4f\x1d\xdf\x93\xd2\x24\x70\xa8\xeb\x05\x1b\x78\x31\x58\x5a\xd7\xef\x07\x17\xa9\xba\xc0\xc5\xcc\xb5\x1d\x1a\xac\x29\xdb\x81\x84\xcd\x05\x83\xff\x70\x90\xa4\x41\xa2\x63\x4b\x40\xf5\x3a\x0a\x1c\xe1\xd1\xc0\x5e\x81\x26\x22\xdf\xaf\xb1\xcf\x49\xc1\x0c\x28\xb0\x77\x84\x73\xbc\x51\x02\xf7\x98\x05\xa0\xeb\xe2\x24\x81\x67\xe2\x1d\x39\x47\xa1\x1f\x6e\xf8\x4f\xff\x02\x59\x8f\x21\xfc\x69\x7c\xb7\x0c\x73\x31\x99\x99\x17\x68\x01\x96\x76\xf8\x1c\x9d\x5d\xa0\xd9\x7d\x40\x18\xfc\x9f\x42\x7e\x39\x37\xc6\x96\xb1\x97\x44\x93\x6b\x64\xce\x2c\x78\x30\x59\x58\x8b\x54\x21\xfa\x36\xb1\x3e\xa3\xc5\xe5\x67\xe3\x66\x8c\xc2\x8d\xed\x40\x04\x7d\x2a\xad\x17\xcc\xef\xb5\x94\x1c\xb9\x9c\xdd\xdc\x18\xa6\xd5\xe0\x46\x2c\x80\x60\x69\x45\x09\x9a\x2c\xd0\xe0\xcb\xf4\x1f\xe1\x46\x26\x2f\x64\xd4\x21\x6e\xc4\xb0\x8f\x7c\x1c\x6c\x22\x88\xc7\xa0\xec\xc7\x96\x0b\xca\x48\x7f\x51\x88\xf5\x15\x83\x10\xad\x7c\xcf\xd1\x07\xa0\xe8\xc2\xd3\xf0\x27\x66\x25\x7c\x59\xb2\x48\x80\x2e\x04\xb5\x84\xe4\x73\x59\x71\x9c\x08\x8e\xe8\x1a\x0d\x6f\xc9\xe3\x08\xdd\x61\x3f\x22\xa7\x28\xc4\x1e\xe3\x2a\x24\xaa\x0c\x09\x66\xce\xd6\x0e\xb1\xd8\x42\xd5\xc4\x5e\x8f\x8a\x29\x94\x62\x2e\x59\xe3\xc8\x87\xd2\xc7\x2b\x9f\xf0\x10\x3b\x44\x96\xf3\xa0\xf4\xf6\xde\x13\x5b\x9b\x7a\x6e\xae\x42\x8b\x71\xf7\xa4\x67\x8f\x36\x76\x1c\x1a\x05\x82\xa7\xf0\xad\xf1\xa7\xa9\xb1\x07\x9f\xc4\x2e\x8b\x00\x88\x65\x66\xcf\xf3\xf9\x50\xeb\x2a\x5a\xd1\xf0\x04\xc1\x3f\x9e\x8b\x56\xde\xc6\x0b\x84\xca\x94\xb9\x9c\x4e\x47\xea\x39\x76\x5d\x06\xfb\x04\xb6\x16\x66\xd8\x11\x84\x41\x60\xd8\x23\x84\x6b\xf8\xee\x9f\xa7\x89\x48\xac\xc9\x56\x01\x05\x0d\x64\x03\x52\x45\x2d\x2b\xb5\xe7\x3d\xd8\xdb\x6a\xe7\x86\xf8\x51\x52\x03\x47\x2b\x4a\x7d\x82\x83\x4c\x1a\x5d\x19\xd7\xe3\xe5\xd4\x42\xd7\xe3\xe9\xc2\xc8\xaf\x05\xae\x78\xca\x62\xdf\xdb\x79\x82\xb8\x36\xe6\x2a\xbb\xbf\x71\x1a\xac\x4e\x4e\x2b\x15\x9e\xc4\x84\xac\xd7\xc4\xe9\x3b\xd0\x89\xd2\x24\xce\xa5\xf0\xdb\xba\xb8\xa7\x72\x34\x24\xc0\xbc\x92\xcd\x74\x92\xaf\x28\x73\x09\x7b\xa5\x89\x7c\x43\x52\x5c\x22\xb0\xe7\xb7\x06\xc5\x27\x2e\xac\xed\x39\x28\x89\xd2\x24\x28\x9c\xfc\x8c\x80\xf8\x75\x8e\xc6\xc2\xf6\x16\xf3\x6d\x7d\x1d\x96\xe4\x43\x46\xee\x3c\x1a\x71\xbb\x75\x61\x12\x23\x86\x03\x8e\xe3\x9e\xa1\xb2\x92\xf9\x91\x56\xd4\xeb\x92\x85\x7d\x56\xba\xc9\x3b\x3e\xe5\xb2\x0a\x05\x92\x7d\x0f\x9a\xd9\x2e\x44\x72\xfb\xcb\x0e\x28\x9f\xa0\x5f\x34\x20\xe5\x35\x8c\x60\xd1\xba\x28\x96\x8d\x42\xb7\xb3\x6c\x56\x47\xc9\x9f\xbb\x90\x32\x08\x8b\x7d\x07\xf9\x00\x44\x15\x2c\x6f\xca\x15\x45\x81\xea\x00\xb7\x17\xf0\xfa\x82\x5c\x13\x62\x87\xb0\x37\xeb\xdf\xca\x51\xc1\x06\x11\x1d\x53\xc8\xd7\xc0\x38\x84\xdd\xe9\x44\x76\xf8\xc1\x16\x0f\x36\x6c\x68\x9b\x7b\xbf\xaa\x52\xfa\x52\xde\xa7\x2d\xc4\x4c\x78\x8e\x17\xe2\xde\x79\xb5\xde\xc6\x9e\x65\xeb\x31\x75\xdf\xee\xed\x04\x72\x28\x7e\x50\x01\xc1\xfc\x99\x86\x61\x61\x7c\x5d\x1a\xe6\x65\x43\x24\xf2\xe0\x53\xe9\x6e\x36\x14\x82\x85\x35\x9e\x5b\x71\xfb\x7f\xa3\x1e\x4c\x4c\x50\xa6\x1a\xf6\xa7\x1f\xc9\x23\x73\x86\x6e\x26\xe6\x7f\xc7\xd3\xa5\x91\xfd\x3d\xfe\xbe\xff\xfb\x72\x0c\x83\x03\x7a\xd3\x0b\x50\x34\xfb\x66\x1a\x57\x60\xbb\x05\xf1\x78\x6a\x19\xf3\x03\x01\x67\xba\x5b\xc4\xff\xee\xb9\xad\x58\x8e\x55\xa8\x6d\x23\x40\x9e\x1e\xb5\x63\x42\x18\x82\x0f\x31\x2e\xd5\x8f\x9e\xd9\x8e\xe2\x47\x9c\x46\xcc\x21\x69\xa9\x6b\xb8\x3f\xe5\xa9\xc1\xe0\xfc\xbc\x22\xd1\x61\x53\xe4\xe1\x1d\x8f\x16\x74\x56\x54\xec\x35\xb4\x50\xb7\xb6\x3e\x01\xcf\x21\x05\x9d\x67\xfd\xd2\x42\x8b\x95\x97\x22\x86\x03\xc1\x3e\x93\x1a\x5a\xac\x55\xc9\x41\xb7\xa0\x81\x1e\x72\x4b\x8e\x57\xb2\x29\x45\xe4\xfd\xeb\x3c\x8e\x25\x53\x58\xcb\x90\xd7\x95\x41\x9a\xc9\xa0\x56\x76\x6f\x5a\x3f\xaf\x60\x6d\x6b\xd6\xcd\x7a\x7f\xc9\xb4\x06\x73\x0f\x09\xee\x88\x0f\x4e\x21\x41\x1e\x2a\x54\xfd\x20\x67\x27\x38\x5c\x6a\x5e\xee\x88\x3c\xf8\xd6\xbe\x92\x51\xd0\xbd\xe6\xde\x26\xc0\x22\x02\xd5\x35\x61\xff\xd7\xbb\xd3\xff\xfd\x7f\xcf\xc2\xbf\xff\x51\xc7\xc3\x20\x51\x1a\xe2\xc8\x8e\xc6\x27\xc6\x2a\x67\x67\xba\x02\x08\x43\x23\xab\xef\x75\x55\xd5\x24\xc8\x20\x9c\xf6\x0a\x12\x07\xc7\x6c\x88\xe2\x7b\x28\xe0\x0d\x51\x64\x98\xdf\x4c\xb0\xbd\x92\xad\x93\xd8\xee\xb4\xdf\xe3\xed\x32\x33\xa7\x6d\xdd\x1d\xc5\xf2\x97\xb3\xe9\xf2\xc6\x94\x29\x95\xd7\x00\x29\xca\x00\xe2\x7d\x87\xfd\xe1\xa0\xd3\x40\x01\xe1\x60\x64\xe3\xf8\x70\xa0\xad\x30\x7a\x6f\x28\xb4\xcd\xea\x20\x1c\x2d\xec\xd7\x84\xa4\x25\x14\xe1\x2d\x79\xdc\x5f\x06\x99\x0b\x6b\x3e\x9e\x98\x0d\x68\xab\x84\x77\x60\x02\x55\x29\x8d\xaf\xae\x72\xd6\xba\xf8\x88\xbe\xcc\x27\x37\xe3\xf9\x0f\xf4\x1f\xe3\x07\x1a\x7a\xee\xe1\x3d\xf8\x88\x48\x75\x36\x9b\xb0\x36\xfa\xd9\x8a\x76\x95\x0d\x28\x29\xa4\x89\x79\x65\x7c\x7f\x42\xa3\x52\xeb\x72\xfa\xe4\x4d\x5f\x6d\xdb\x5a\x2e\x26\xe6\xbf\xd1\x4a\x30\x38\x70\x0e\x13\xe1\x51\xa5\x2f\xd4\x79\x2a\xdb\x5b\x6f\x6e\xaa\x5e\xd9\xc9\xc7\x72\x87\xad\x73\x2d\x6e\xa8\xbd\x39\x17\xab\xeb\xe6\x5e\xa9\x97\x8f\xaa\x6d\xbb\xb6\xc6\x6d\xe0\xe0\xc7\xf8\xfd\x73\xdd\x5e\x9a\x13\x98\xb2\x12\xef\x4b\xba\xf3\x18\xd2\x6b\xb7\x82\xfb\x75\xc7\xec\x51\x7a\x83\xa6\xf3\x7c\x4f\xab\x7d\xfa\x0c\xec\xd9\xd5\xdb\xfd\x54\x3f\xaa\xbd\x28\x68\x41\x40\x43\x3b\x3c\x0a\x88\x44\x71\x1e\x87\xa6\xff\x3d\x09\x56\x15\x4d\x76\xa3\x07\x09\xef\x1b\x50\x51\x77\x1e\x53\x7a\x57\x59\x00\x51\xef\x5e\x7e\xf7\x1e\xc5\xc7\x8a\x81\x6e\xdb\xb6\xc6\x5b\x2f\x70\xc9\x83\x5d\xfe\x1a\x60\x83\xde\xe4\xca\xbf\x57\xd7\x5b\xad\xe5\x71\x64\x9f\x26\x8a\xec\x1d\x0b\x1e\x00\xa4\xe7\xf0\x37\x19\x6a\x77\x3f\x4e\x41\x81\x7b\x35\x0a\xd5\x77\x21\xc8\xa5\x27\x3a\x44\x05\xd4\xa2\x6f\x9f\x8d\xb9\xa1\xfd\xc6\xf2\x01\x4e\x6d\x11\x41\xb3\xb9\xfe\x4b\x4a\x2c\xd2\x1c\xd8\x84\xa1\x24\x5c\x39\xb6\xf7\xd3\x7d\x1a\x4d\xb4\xf2\xa3\x14\x6a\x29\x87\x64\xef\x4a\x95\xd9\x1d\xfc\x31\x5c\xaf\xb3\xd3\xca\x21\x99\x64\x77\x10\x47\x2d\xe9\x82\x9d\xa7\x30\xa0\x5e\x5d\xe9\x23\xc3\x91\x53\x50\xf9\xa6\xd1\x8a\xa5\xb4\xa0\x3b\xb2\xdc\x27\xa6\x97\xc9\x4c\xfe\x9b\x56\x1b\xac\x9c\x6c\x77\x44\x75\x5f\xcf\x5e\x06\x5a\xed\x77\xbb\x36\x8c\x75\x8b\xba\x83\x4d\x07\xd9\x97\x01\x98\xdd\x43\xb5\x81\xd2\x1e\x4c\x8a\xaa\xf7\x57\xf8\x47\xe7\x86\xb2\xa9\xda\xa1\xef\x50\x86\x28\x2a\x2d\x5e\x73\x1f\x83\x22\x9a\xec\x75\x01\x54\x5c\x71\x18\xb8\x23\xf5\xcc\xaa\x95\x4e\x40\xea\x3a\xa7\x9a\xe9\xc5\xc3\x91\x0e\x0b\x89\x62\xcd\xbc\xfa\xc4\xe3\x42\x35\x21\xfa\x7c\xe4\xa7\xe3\xa3\x6f\x97\xaa\xb1\x27\x0f\xea\x20\xec\x92\x6c\x36\x4a\x8f\xba\xf6\x8a\xd2\xdb\x7e\x0a\xaa\xc1\x40\xeb\x08\x36\x1c\xa6\x9f\xed\xce\x3e\x7e\x44\x03\x4e\xfd\xe4\xb7\x36\xaa\x14\x07\xe7\xe7\xf2\x36\xf9\xf4\x74\x84\xf4\x82\x0e\x75\xbb\x09\x7a\x9c\x47\x84\xe9\x45\x57\x34\xda\x6c\x45\x27\xf3\x05\xd1\x66\x07\x0a\xa2\x25\x17\xd2\xd1\x5b\xed\x27\x98\xa2\xdf\xbe\xcd\x65\x4f\xf7\x13\x49\x04\xd3\x77\xe8\x13\x41\x54\x26\xf2\xbf\xae\xbc\xa2\xf7\xc1\x89\xcb\x68\x88\xd4\x0f\xc7\xea\xcb\xc5\xc1\xdc\x81\x7c\x5d\xb4\x08\x16\x37\x54\xd3\xa2\x1c\x47\x74\x12\xeb\xae\x39\x6d\x6d\x4d\x32\x69\x55\x35\xc9\x64\x27\x9f\x4c\xe8\xcf\x00\x00\x00\xff\xff\x47\xfc\xd6\x1f\x94\x2a\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10900, mode: os.FileMode(420), modTime: time.Unix(1472739164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_index_participants_by_toidSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xb1\x0a\xc2\x30\x10\x86\xf7\x7b\x8a\x1b\x15\xe9\x13\x74\x12\x1b\xa4\x4b\x2a\xd5\x82\x5b\x48\xdb\x60\x6e\x30\x17\x92\x03\xe9\xdb\x2b\x3a\xd8\xda\xc5\xf5\xf8\xf8\xfe\xfb\x8a\x02\x77\x77\xba\x25\x2b\x0e\xbb\x08\x70\x68\xd5\xfe\xa2\xb0\xd6\x95\xba\xa2\xe7\x68\xfa\xc9\x78\xa6\x11\x1b\x8d\x9e\xb2\x70\x9a\x0c\x47\xf7\xe2\x89\x83\x89\x36\x09\x0d\x14\x6d\x90\x8c\xdd\xb9\xd6\x47\xec\x25\x39\x87\x9b\x35\x4b\xe3\xb6\xfc\xd1\xcb\x47\x2f\x4b\xbd\x24\x1b\xb2\x1d\xfe\x1c\x98\xd3\xef\x09\x98\x27\x55\xfc\x08\x00\x55\xdb\x9c\xd6\x49\xe5\xe2\xfe\xfd\xa5\x84\x67\x00\x00\x00\xff\xff\x33\xec\x54\x7a\x15\x01\x00\x00")

func migrations2_index_participants_by_toidSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_index_participants_by_toidSql,
		"migrations/2_index_participants_by_toid.sql",
	)
}

func migrations2_index_participants_by_toidSql() (*asset, error) {
	bytes, err := migrations2_index_participants_by_toidSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1463067478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_aggregate_expenses_for_accountsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\x41\x6b\xea\x40\x10\xc7\xcf\xee\xa7\x98\xa3\xf2\x14\xde\x7b\x94\x5e\x3c\xa5\x4d\x0a\xd2\x34\x4a\x88\x50\x4f\xcb\x74\x77\xd1\xa5\xc9\x6e\xd8\x9d\xd4\xa6\x9f\xbe\x51\x43\x10\xb5\x4d\xe6\x96\xf0\xfb\xff\x19\x92\xdf\xcc\x66\xf0\xa7\xd0\x5b\x87\xa4\x60\x5d\x32\xf6\x98\x46\x41\x16\x41\x16\x3c\xc4\x11\xa0\x10\xb6\x32\xc4\x3d\x21\x69\x4f\x5a\x78\x18\x33\x68\x06\xa5\x74\xca\x7b\x38\x1f\xb1\x43\x87\x82\x94\x83\x0f\x74\xb5\x36\xdb\xf1\xfd\xdd\x04\x92\x65\x06\xc9\x3a\x8e\xa7\xa7\x9c\xf7\x8a\xb8\xb0\x52\xfd\x96\xfb\xf7\xff\x32\x77\x5c\x43\xb9\x12\x1d\xd5\x9c\xea\xf2\x10\xf7\x05\xe6\xb9\x36\xd4\xa1\x10\x46\x4f\xc1\x3a\xce\xe0\xef\x29\x24\x51\xe7\x35\xd7\x46\xd8\x42\xc1\x68\xf4\xa6\xb7\xfd\xb4\xad\x68\x18\xbe\x57\xea\xfd\xba\x7d\xd4\x83\xb7\xf5\xbd\xed\x85\x35\xb4\xeb\xea\x07\xe3\xdd\xf6\x3d\x3c\x1a\x53\x61\x3e\xb4\xbd\xa5\x87\xee\x5e\x95\xb2\x91\x49\x72\xa4\xe6\xb3\x1c\x5e\x90\x2e\x54\x63\x50\x51\xc2\x5e\xd3\xee\xf8\x08\x5f\xd6\xa8\x8b\x7f\xbc\x4a\x17\x2f\x41\xba\x81\xe7\x68\x33\x6e\xfd\x9a\x9e\x09\x33\xbd\x96\x60\xc2\x26\xf3\xce\xd8\x45\x12\x46\xaf\x37\x8c\xe5\x6d\x17\xd7\xf2\x13\x96\xc9\x4d\xa7\x5b\xe4\xd0\x76\x7e\x0f\xa1\xdd\x1b\xc6\xc2\x74\xb9\x1a\xd4\x3e\x3f\xa1\x3f\x9d\xce\x9c\x7d\x07\x00\x00\xff\xff\x26\xb0\x63\x72\x6c\x03\x00\x00")

func migrations3_aggregate_expenses_for_accountsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_aggregate_expenses_for_accountsSql,
		"migrations/3_aggregate_expenses_for_accounts.sql",
	)
}

func migrations3_aggregate_expenses_for_accountsSql() (*asset, error) {
	bytes, err := migrations3_aggregate_expenses_for_accountsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_aggregate_expenses_for_accounts.sql", size: 876, mode: os.FileMode(420), modTime: time.Unix(1472211234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations7_account_limitsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x90\x3d\x4f\xc3\x30\x10\x40\xe7\xdc\xaf\xb8\x31\x11\xad\x04\x08\xb1\x74\x0a\xc4\x48\x88\xd0\x56\x51\x32\x74\x8a\x0e\xc7\x4a\x4f\xaa\xed\xca\xbe\x16\xfa\xef\x09\x50\x21\x0f\x7c\x78\x7e\xcf\x77\xf7\xe6\x73\xbc\xb0\x3c\x06\x12\x83\xdd\x1e\xe0\xbe\x51\x65\xab\xb0\x2d\xef\x6a\x85\xa4\xb5\x3f\x38\xe9\x77\x6c\x59\x22\xe6\x90\xd1\x30\x04\x13\x23\xa6\x4f\x6f\x29\x90\x16\x13\xf0\x48\xe1\xc4\x6e\xcc\x6f\x6f\x0a\x5c\xae\x5a\x5c\x76\x75\x3d\x83\x0f\x86\x62\x34\xd2\x6b\x3f\x98\xbf\xbc\xab\xeb\xc4\x9b\x88\x4f\xd5\xd2\x5b\xef\xf7\x66\xda\x90\xbd\xc3\x2c\x7b\xe1\x91\x9d\x7c\x73\x58\xa9\x87\xb2\xab\x5b\xbc\xfc\x9a\x34\x10\xef\x4e\xbd\x1c\x82\xf3\xc7\xe9\xeb\x7f\x79\xeb\x9d\x6c\x13\xe3\xcc\x67\xbf\xf1\xeb\xe6\xf1\xb9\x6c\x36\xf8\xa4\x36\xf9\xb9\xc6\x2c\x39\xaf\x80\x62\x01\x90\x56\xad\xfc\xab\x03\xa8\x9a\xd5\xfa\xc7\xaa\x0b\x78\x0f\x00\x00\xff\xff\x79\x84\x25\x20\x83\x01\x00\x00")

func migrations7_account_limitsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations7_account_limitsSql,
		"migrations/7_account_limits.sql",
	)
}

func migrations7_account_limitsSql() (*asset, error) {
	bytes, err := migrations7_account_limitsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/7_account_limits.sql", size: 387, mode: os.FileMode(420), modTime: time.Unix(1464705819, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations8_account_limits_two_waySql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x91\xcd\x4e\xc4\x20\x14\x85\xf7\x3c\xc5\xdd\x3b\x5d\xb8\xee\x0a\x05\x57\x0c\x98\x06\xd6\x04\xc7\x66\xbc\xc9\x00\x93\x4a\xfd\x79\x7b\xad\xda\x94\xe2\x5f\xdb\xed\x21\xe7\xcb\xe5\x7c\x55\x05\x17\x1e\x8f\x9d\x4b\x2d\x98\x33\x21\x54\x68\xde\x80\xa6\x57\x82\x83\x3b\x1c\x62\x1f\x92\x3d\xa1\xc7\xf4\x08\x0d\x97\x74\xcf\xe1\x5a\x09\xb3\x97\xe0\xdd\x8b\x8d\xe7\xf6\xbd\x88\x31\x80\x56\xf3\xc0\xc6\x3e\xd5\xcb\x61\xf7\x0e\x4f\xaf\x36\xf5\x5d\x88\x4f\x6d\x37\xd0\x3e\x93\x0f\xe6\x2a\x92\x8f\x21\x3d\x14\xac\x31\x5b\x40\x23\x00\x94\xb1\x1f\xff\x68\x31\xc0\x1d\x1e\x31\x24\x90\x4a\x83\x34\x42\x00\xe3\x37\xd4\x08\x0d\xd5\xe5\x6e\xde\x9c\xce\x5f\xd3\xca\x0f\xfd\xb3\x57\x13\x92\x8b\x63\xf1\x39\x6c\x55\x37\x2c\xf2\x4d\xdf\x6a\x75\x5f\xd3\x4e\xe6\xc6\xfd\x37\xa8\xcb\x58\xa5\xcd\x7f\xd4\xb1\x46\xdd\xfe\xe6\x6e\x57\xbc\xe7\x86\xca\xb7\xb9\x87\x9a\xbc\x05\x00\x00\xff\xff\xa7\x41\x23\x51\x25\x03\x00\x00")

func migrations8_account_limits_two_waySqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations8_account_limits_two_waySql,
		"migrations/8_account_limits_two_way.sql",
	)
}

func migrations8_account_limits_two_waySql() (*asset, error) {
	bytes, err := migrations8_account_limits_two_waySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/8_account_limits_two_way.sql", size: 805, mode: os.FileMode(420), modTime: time.Unix(1464875699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations9_1_assetsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x31\x4f\xc3\x30\x10\x85\x67\xfc\x2b\x6e\x4c\x44\x3a\x80\x10\x4b\xa7\x40\x3c\x44\x04\xa7\x58\xb1\x44\x27\xeb\x9a\x58\xe1\xa4\xd6\xae\x6c\x17\x94\x7f\x4f\xdd\x42\x5b\x15\x71\xeb\x7b\xdf\xe9\xbd\x37\x9b\xc1\xed\x86\x46\x8f\xd1\x80\xda\x32\xf6\x2c\x79\xd9\x71\xe8\xca\xa7\x86\x03\x86\x60\x22\x64\x0c\xf6\x47\x03\x9c\x6f\x45\x63\x30\x9e\x70\x5d\xb0\x9b\x38\x6d\xcd\x49\x20\x1b\x41\xb4\x1d\x08\xd5\x34\xc5\x81\xeb\xdd\x70\x96\xfb\x0f\xf4\xd8\x47\xe3\xe1\x13\xfd\x44\x76\xcc\xee\xee\xf3\x2b\x80\x42\xd8\xed\x0d\xff\x01\x8f\x0f\x7f\x01\x8d\xd6\xd9\x69\xe3\x76\x01\x56\xce\xad\x0d\xda\x2b\xcb\x42\xd6\xaf\xa5\x5c\xc2\x0b\x5f\x66\x34\xe4\x2c\x9f\x9f\x9a\x2a\x51\xbf\x29\x0e\xb5\xa8\xf8\xfb\xb1\x70\xd0\x29\xb3\x3e\xe6\xd0\x87\x7a\xad\xf8\xdd\x22\x49\xc5\x4f\xc6\x02\x92\x98\x7e\x5d\xae\x58\xb9\x2f\xcb\x58\x25\xdb\xc5\xe5\x8a\xf3\xef\x00\x00\x00\xff\xff\x34\x6b\x2a\x7c\x69\x01\x00\x00")

func migrations9_1_assetsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations9_1_assetsSql,
		"migrations/9_1_assets.sql",
	)
}

func migrations9_1_assetsSql() (*asset, error) {
	bytes, err := migrations9_1_assetsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/9_1_assets.sql", size: 361, mode: os.FileMode(420), modTime: time.Unix(1469797719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations9_2_optionsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xc8\x2f\x28\xc9\xcc\xcf\x2b\xe6\xd2\xe0\x52\x50\xc8\x4b\xcc\x4d\x55\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x36\xd2\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xd1\x01\x4a\xa6\x24\x96\x24\x2a\x94\xa4\x56\x94\x80\x38\x01\x41\x9e\xbe\x8e\x41\x91\x0a\xde\xae\x91\x1a\x20\x5d\x9a\x5c\x9a\xd6\x5c\x5c\x5c\xc8\x36\xb9\xe4\x97\xe7\x71\x71\xb9\x04\xf9\x07\xa0\xda\x64\x0d\x08\x00\x00\xff\xff\x26\x18\xd1\xf9\x8f\x00\x00\x00")

func migrations9_2_optionsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations9_2_optionsSql,
		"migrations/9_2_options.sql",
	)
}

func migrations9_2_optionsSql() (*asset, error) {
	bytes, err := migrations9_2_optionsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/9_2_options.sql", size: 143, mode: os.FileMode(420), modTime: time.Unix(1484155113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations9_commissionSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\x4d\x4f\xc2\x30\x18\x3e\xd3\x5f\xf1\xde\xd8\x22\x24\x1e\x8c\x07\x48\x48\xa6\xab\x66\x71\x0e\x9c\x5b\x22\xa7\xa5\x2b\x2f\x50\x65\x2b\x69\x8b\xca\xbf\xb7\x7c\x38\x16\x40\x81\x9d\xd6\x3c\x4f\x9f\xaf\xb4\xdd\x86\xab\x42\x4c\x14\x33\x08\xe9\x9c\x90\xfb\x98\x7a\x09\x85\xc4\xbb\x0b\x29\x70\x59\x14\x42\x6b\x21\x4b\x70\x48\x43\x8c\xa0\xfa\x72\x31\xd1\xa8\x04\x9b\xb5\xc8\xea\xf8\x81\xcb\x6c\xca\xf4\x74\xf5\xcf\xa7\x4c\x31\x6e\x50\x39\xb7\x37\x2e\x44\xfd\x04\xa2\x34\x0c\x77\xbc\x4f\x36\x5b\x20\x34\xde\xb5\x2c\xf3\x3d\x78\x3c\x63\x26\x1b\xa3\x45\xad\xbe\x28\x4d\x05\x83\x4f\x1f\xbc\x34\x4c\xe0\x7a\x43\x9c\xa3\xe2\x58\x6e\xb8\x27\xa8\x83\x38\x78\xf6\xe2\x21\x3c\xd1\xa1\x23\x46\x2e\x71\xbb\x55\xc9\x34\x0a\x5e\x52\x0a\x41\xe4\xd3\xb7\x5a\xd7\x2c\xdf\x96\xe9\x47\xf5\x05\xd2\xd7\x20\x7a\x84\xdc\x28\xeb\xe9\xfc\x16\xb6\x6a\x5b\xb1\x63\x2a\x8c\x73\xb9\xb0\xd9\xfe\x11\x72\x9c\xdd\x26\xed\x5e\xaf\x39\x56\xb2\x68\xba\x9d\x8e\xc1\x6f\xe3\xb6\x60\x0f\x86\xa6\x91\x15\x7a\x96\x77\x66\x96\x73\xbc\x34\xc0\xfa\xd2\xca\xc7\x0e\x8b\x13\x54\xc7\x83\x1c\xb0\x4e\x04\xd2\x1a\x2f\x9a\x62\x7d\xa1\x32\xf9\x6b\x90\x0d\x8b\xcb\xd1\x19\x2c\xeb\xba\x40\x55\x1f\x90\xd4\xdf\xbf\x2f\xbf\x4a\x42\xfc\xb8\x3f\x38\x78\xff\x5d\xf2\x13\x00\x00\xff\xff\xad\xa9\xe9\xc7\x29\x03\x00\x00")

func migrations9_commissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations9_commissionSql,
		"migrations/9_commission.sql",
	)
}

func migrations9_commissionSql() (*asset, error) {
	bytes, err := migrations9_commissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/9_commission.sql", size: 809, mode: os.FileMode(420), modTime: time.Unix(1470062144, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_aggregate_expenses_for_accounts.sql": migrations3_aggregate_expenses_for_accountsSql,
	"migrations/7_account_limits.sql": migrations7_account_limitsSql,
	"migrations/8_account_limits_two_way.sql": migrations8_account_limits_two_waySql,
	"migrations/9_1_assets.sql": migrations9_1_assetsSql,
	"migrations/9_2_options.sql": migrations9_2_optionsSql,
	"migrations/9_commission.sql": migrations9_commissionSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_aggregate_expenses_for_accounts.sql": &bintree{migrations3_aggregate_expenses_for_accountsSql, map[string]*bintree{}},
		"7_account_limits.sql": &bintree{migrations7_account_limitsSql, map[string]*bintree{}},
		"8_account_limits_two_way.sql": &bintree{migrations8_account_limits_two_waySql, map[string]*bintree{}},
		"9_1_assets.sql": &bintree{migrations9_1_assetsSql, map[string]*bintree{}},
		"9_2_options.sql": &bintree{migrations9_2_optionsSql, map[string]*bintree{}},
		"9_commission.sql": &bintree{migrations9_commissionSql, map[string]*bintree{}},
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

