# unixxd
unixxd is xxd app which can display also unicode code points instead of hex bytes

## install
```
go install github.com/dankox/unixxd@latest
```

## example of hex display
```
unixxd main.go
```
<img width="570" alt="image" src="https://user-images.githubusercontent.com/39275087/174286072-3a805482-c80d-4ec3-8059-4112f042d3f9.png">


## example of unicode display
```
unixxd --unicode main.go
```
<img width="905" alt="image" src="https://user-images.githubusercontent.com/39275087/174286153-255bfdc2-0907-45ba-bae3-570921521ee5.png">

## build on mainframe
Create `syncz_vars.yml` file for synchronizing
```yml
host: mainframe.host
user: user_name
remotedir: /uss/path/to/the/project/unixxd
```

If you don't have `syncz` you can upload the source code to USS using `sftp` in binary mode. All *.go files needs to be utf-8.

Build it using `go build` or `go install`.