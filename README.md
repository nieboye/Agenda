# Agenda

## Install

```shell
go get -u github.com/Binly42/agenda-go
```


## Methods

```shell
Usage:
  agenda [flags]
  agenda [command]

Available Commands and local flags:
  add        -s startTime //create Meeting 
  			-e endTime 
  			-t title 
  			-p participator 
  			
  delete     -u           // delete user or meeting
  			-m 			// note that -u & -m can't appear at the same time
  			-t title         
            
  edit       -i           //edit meeting, invite a participator or delete participator
  			-d 
			-p participator 
  help        Help about any command
  
  login      -u username  //login 
  			-p password
  
  logout     			 //logout
  
  register   -u username  //register for further use
  			-p password
  			[-e] email
  			[-t] phone
  			
  search     -u           //search users or meetings
  			-m			// note that -u & -m can't appear at the same time
  			-s startTime
  			-e endTime

Root Flags:
  -h, --help                  help for agenda
  -l, --license licensetext   Name of license for the project (can provide licensetext in config)
  -b, --projectbase string    base project directory eg. github.com/spf13/
      --viper                 Use Viper for configuration (default true)
  -a, --author string         Author name for copyright attribution (default "YOUR NAME")
        --config string         config file (default is $HOME/.cobra.yaml) (default "./.cobra.yaml")
Use "agenda [command] --help" for more information about a command.

## 样例

register 

```shell
$ ./agenda-go.exe register -u lyb -p 12345 -e lyb@gmail.com -t 2080290
register called
register sucessfully!

```

login

```shell
$ ./agenda-go.exe login -u lyb -p12345
login called
login called by lyb
login with info password: 12345
login sucessfully!

```

create meeting

```shell
$ ./agenda-go.exe add -s "2011-01-01 10:00:34" -e "2011-01-02 08:00:34" -t MatrixShareMeeting -p lrd
create Meeting called
start: 2011-01-01 10:00:34 +0000 UTC end: 2011-01-02 08:00:34 +0000 UTC
sucessfully create meeting
```

search user

```shell
$ ./agenda-go.exe search -u
search called
+-------+---------------+---------+
| NAME  |     EMAIL     |  PHONE  |
+-------+---------------+---------+
| sky   |               |         |
| lrd   |               |         |
| lyb   | lyb@gmail.com | 2080290 |
| binly | binly@git.com | 1234567 |
+-------+---------------+---------+

```

search meeting

```shell
$ ./agenda-go.exe search -s "2011-01-01 10:00:34" -e "2011-01-02 08:00:34"
search called
+--------------------+---------+----------------------+----------------------+---------------+
|       TITLE        | SPONSOR |      STARTTIME       |       ENDTIME        | PARTICIPATORS |
+--------------------+---------+----------------------+----------------------+---------------+
| MatrixShareMeeting | lyb     | 2011-01-01T10:00:34Z | 2011-01-02T08:00:34Z |  lyb          |
+--------------------+---------+----------------------+----------------------+---------------+

```

add participator

```shell
$ ./agenda-go.exe edit -i -u sky -t MatrixShareMeeting
edit called
$ ./agenda-go.exe search -s "2011-01-01 10:00:34" -e "2011-01-02 08:00:34"
search called
+--------------------+---------+----------------------+----------------------+---------------+
|       TITLE        | SPONSOR |      STARTTIME       |       ENDTIME        | PARTICIPATORS |
+--------------------+---------+----------------------+----------------------+---------------+
| MatrixShareMeeting | lyb     | 2011-01-01T10:00:34Z | 2011-01-02T08:00:34Z |  sky lyb      |
+--------------------+---------+----------------------+----------------------+---------------+

```

 logout

```shell
$ ./agenda-go.exe logout
logout called
logout sucessfully!

```

delete meeting

```shell
$ ./agenda-go.exe delete -m -t MatrixShareMeeting
delete called
$ ./agenda-go.exe search -s "2011-01-01 10:00:34" -e "2011-01-02 08:00:34"
search called

```

delete current user

```shell
$ ./agenda-go.exe delete -u
delete called
CancelAccount successfully
$ ./agenda-go.exe login -u lyb -p12345
login called
login called by lyb
login with info password: 12345
Error[login]： a nil user/*user is to be used

```






