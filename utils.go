package utils

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	// Import go-mssqldb strictly for side-effects
	_ "github.com/denisenkom/go-mssqldb"
)

//Log
var (
	Log *log.Logger
)

//NewLog is used to save logs in a log file
func NewLog(logpath string) {

	file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)

}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

////this is how you call the copy
// fmt.Printf("Copying %s to %s\n", os.Args[1], os.Args[2])
// err := CopyFile(os.Args[1], os.Args[2])
// if err != nil {
// 	fmt.Printf("CopyFile failed %q\n", err)
// } else {
// 	fmt.Printf("CopyFile succeeded\n")
// }

//MsSQLSendAlarm can be used to send predefined alarms in the form of inserting text to tables in MSSQL DB used for Alarm notifications
//_connString sample: "server=x.x.x.x;user id=sa;password=xxxxx;port=1433"
//_query sample: "insert into [x.x.x.x].Alarms.dbo.xxx_tbl (number,date,text,flag_sent) select 'xx',getdate(),'Failed xxxx',0"
func MsSQLSendAlarm(_connString string, _query string) {

	conn, err := sql.Open("mssql", _connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare(_query)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	stmt.Exec()
	defer stmt.Close()

}
