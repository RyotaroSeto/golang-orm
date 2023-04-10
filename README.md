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
