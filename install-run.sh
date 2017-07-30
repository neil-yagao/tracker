#! /bin/bash

#kill last process
kill -9 `cat pid`
rm save_pid.txt
BASE_DIR=`pwd`
#checkout the latest code
git pull origin master 
go get "github.com/beego/bee"
go get "github.com/astaxie/beego"
go get "github.com/go-sql-driver/mysql"

#compile font end
cd $BASE_DIR/src/web-mobile
npm run install

#compile back end 
cd $BASE_DIR/src/backend/src
go build main.go
nohup my_command > run.log 2>&1 &
echo $! > $BASE_DIR/pid