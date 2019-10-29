# Scraping-school

高校名とコース，偏差値をスクレイピングするプログラムです．

スクレイピングには [goquery](https://github.com/PuerkitoBio/goquery) を使用しました．


## スクレイピングしたサイト

スクレイピングしたサイトは欲しい情報によって分けました

高校名、コース、偏差値→.gitignoreで隠したサイト

Url→グーグル検索結果


## Directory
### check
エラーチェックの処理の関数

### csv-name-course
都道府県ごとに高校名，偏差値，コースの情報を持つ csv ファイルが保存してます．
.gitignore で隠してます．

### csv-nam-url
csv-name-course に高校の公式ホームページのURLを追加した csv ファイルが保存されている．
これも.gitignore で隠しています．

### env
関数の引数などで使われる固有名詞やURLの変数があります．
スクレイピングしたサイトのURLは隠したほうが良さそうなので .gitignore で隠してます．

### prefectures
スクレイピングするサイトのURLで使用したり，csv ファイルの名前に使ったりしています．

### scrape
スクレイピングする際に使用する関数をまとめたものになっています．

## 動かすには
実行するには次のコマンドを打てばいいです

``` go
go run main.go
```


