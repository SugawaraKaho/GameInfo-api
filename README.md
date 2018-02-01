# GameInfo-api

## 作業する時
作業するときは,
```
cd $GOPATH/src/[プロジェクト名]
```
$GOPATHがわからない時は
```
go env GOPATH
```
というコマンドで出てきたパスがGOPATH

## goパッケージのインストールの仕方

`go get ~~~~`て書かれているやつは以下のコマンドに置き換える
```
glide get example.com/example 
glide install
```

## ファイルの実行
```
$ go run example.go
```
