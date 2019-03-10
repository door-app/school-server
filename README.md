# school-server
GRPC

## DB
# gooseをinstall
```
$ go get bitbucket.org/liamstask/goose/cmd/goose
$ goose
なんか出れば良いと思う
```

### マイグレーションファイル作成
```
$ goose create create_users_table sql
```
プロジェクトルートで実行すれば、db/migrations 以下にファイルが生成されるので、
SQL書く。

### マイグレーション実行
```
$ goose up
$ goose down
```


### マイグレーションファイル名サンプル
```
▽ テーブル作成
create_table_****

▽ テーブル削除
drop_table_**** / remove_table_****

▽ カラム追加
add_column_****_tablename

▽ カラム削除
delete_column_****_tablename / remove_column_****_tablename

▽ インデックス追加
add_index_****_tablename

▽ インデックス削除
delete_index_****_tablename / remove_index_****_tablename
```