/*
Released as open source by NCC Group Plc - http://www.nccgroup.com/

Developed by Jose Selvi, jose dot selvi at nccgroup dot com

http://www.github.com/nccgroup/wstalker

Released under AGPL see LICENSE for more information
*/

package filedump

import (
	"fmt"
	"os"
)

// FileDump is ...
type FileDump struct {
	filename string
	file     *os.File
}

// NewFileDump function...
func NewFileDump(filename string) (*FileDump, error) {
	var e error
	f := &FileDump{}

	f.filename = filename
	f.file, e = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if e != nil {
		return nil, e
	}

	return f, nil
}

func (f *FileDump) Write(method, url, request, response string) error {
	_, e := fmt.Fprintf(f.file, request+","+response+","+method+","+url+"\n")
	return e
}

// Close function...
func (f *FileDump) Close() error {
	return f.file.Close()
}
