#! /bin/bash

#kill last process
kill -9 `cat pid`
rm -f pid
BASE_DIR=`pwd`
#checkout the latest code
git pull origin master 
go get "github.com/beego/bee"
go get "github.com/astaxie/beego"
go get "github.com/go-sql-driver/mysql"

#compile font end
cd $BASE_DIR/src/web-mobile
npm run build

#compile back end 
cd $BASE_DIR/src/backend/src
nohup  bee run > run.log 2>&1 &
echo $! > $BASE_DIR/pid
