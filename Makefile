DB_URL=postgresql://postgres:postgres@localhost:5432/test_db?sslmode=disable

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres test_db

#migration#########################################################################################################
#golang-migrate-----------------------------------------------------------------------------------------#
# 初期ファイル作成
golang-migrate:
	migrate create -ext sql -dir db/migration/golang-migrate -seq create_users_table

# 初期ファイル更新
golang-migrateup:
	migrate -path db/migration/golang-migrate -database "$(DB_URL)" -verbose up

# 初期ファイル切り戻し
golang-migratedown:
	migrate -path db/migration/golang-migrate -database "$(DB_URL)" -verbose down

# 追加ファイル作成
golang-migrate1:
	migrate create -ext sql -dir db/migration/golang-migrate -seq add_mood_to_users

# 追加ファイル更新
golang-migrateup1:
	migrate -path db/migration/golang-migrate -database "$(DB_URL)" -verbose up 1

# 追加ファイル切り戻し
golang-migratedown1:
	migrate -path db/migration/golang-migrate -database "$(DB_URL)" -verbose down 1
#golang-migrate-----------------------------------------------------------------------------------------#

#migration#########################################################################################################


#orm#########################################################################################################
#sqlc-----------------------------------------------------------------------------------------#
sqlc:
	sqlc generate
#sqlc-----------------------------------------------------------------------------------------#
#orm#########################################################################################################
