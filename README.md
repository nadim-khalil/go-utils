# Go Utils

This repository contains the `utils` library, a set of useful tools to be use in go projects.

## Getting Started

currently the utils has the below utility functions:
* **NewLog** = it is used to save logs in a log file. Allowing you to choose the log file destination.
* **CopyFile** = it copies a file from src to dst.
* **MsSQLSendAlarm** = it can be used to send predefined alarms in the form of inserting text to tables in MSSQL DB used for Alarm notifications.


### Installing

install using:

```
go get github.com/nadim-khalil/go-utils
```

### How to use
#### NewLog
NewLog is used to save logs in a log file.

*example:*
```
if err != nil {
	logpath := "errors.log"
	utils.NewLog(logpath)
	utils.Log.Printf("error in xxx func: %v \r\n", err)
}
os.Exit(1)
```

#### CopyFile
CopyFile copies a file from src to dst. If src and dst files exist, and are the same, then return success. Otherise, attempt to create a hard link between the two files. If that fail, copy the file contents from src to dst.

*example:*
```
//to := "source file"
//from := "destination file"

fmt.Printf("Copying %s from %s\n", to, from)
err := utils.CopyFile(from, to)
if err != nil {
	fmt.Printf("CopyFile failed %q\n", err)
	logpath := "errors.log"
	utils.NewLog(logpath)
	utils.Log.Printf("CopyFile failed")
} else {
	fmt.Printf("CopyFile succeeded\n")
}
```
#### MsSQLSendAlarm
MsSQLSendAlarm can be used to send predefined alarms in the form of inserting text to tables in MSSQL DB used for Alarm notifications

*example:*
```

//fill the variables in case you want to send Alarms to an SQL DB used for Alarms notifications
iniConnString := "server=x.x.x.x;user id=sa;password=xxxxx;port=1433"
iniSQLQuery := "insert into [x.x.x.x].Alarms.dbo.xxx_tbl (number,date,text,flag_sent) select 'xx',getdate(),'Failed xxxx',0"

if err != nil {

		fmt.Printf("error in getCurrentDir func: %v \r\n", err)
		if iniConnString != "" && iniSQLQuery != "" {
			utils.MsSQLSendAlarm(iniConnString, iniSQLQuery)
		}
		os.Exit(1)
```


## Author

* **Nadim Khalil** - *My personal site* - [nadim.tk](http://nadim.tk)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used


