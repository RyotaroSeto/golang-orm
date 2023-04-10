# golang-orm

## migration
### golang-migrate
- `brew install golang-migrate`


## orm
### sqlc
1. `brew install sqlc`
2. `touch sqlc.yaml`
3. sqlc.yaml記入
4. 作成したいSQLのファイル作成(db/orm/sqlc/query/user.sql)
5. `sqlc generate`実行

### sqlx
1. 特に作業なし。sql文を直書きするのみ


### ent
1. `go get -d entgo.io/ent/cmd/ent`
2. `go run -mod=mod entgo.io/ent/cmd/ent new User(新しいスキーマの初期化)`
3. 作成された`user.go`にFieldを定義する
4. `go generate ./ent(アセットの生成)`
5. CRUDファイルが作成される
