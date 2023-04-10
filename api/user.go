package api

import (
	"database/sql"
	"log"
	"net/http"

	"golang-orm/db/orm/ent"
	db "golang-orm/db/orm/sqlc"
	"golang-orm/util"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type createUserRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type userResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// sqlc////////////////////////////////////////////
// 標準でトランザクションが用意されていない
// sql文をそのまま定義できるが、別ファイルに切り出して、sqlcコマンドを叩き、専用のメソッドが作成される
// レスポンスで登録した結果が返ってくる
// 生のsql文を理解していないと書けない
// 複雑なクエリ作成で困ることはない
// sqlc////////////////////////////////////////////
func sqlcCreateUser(ctx *gin.Context, config util.Config) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal("invalid param", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer conn.Close()

	arg := db.CreateUserParams{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
	}

	store := db.NewStore(conn)
	user, err := store.CreateUser(ctx, arg)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type createUserRequestSqlx struct {
	Username string `db:"username"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
}

func newUserResponseSqlx(user createUserRequestSqlx) userResponse {
	return userResponse(user)
}

// sqlx////////////////////////////////////////////
// 標準でトランザクションが用意されている
// sql文をそのまま定義できる
// レスポンスで登録した結果が返ってこない
// 生のsql文を理解していないと書けない
// 複雑なクエリ作成で困ることはない
// sqlx////////////////////////////////////////////
func sqlxCreateUser(ctx *gin.Context, config util.Config) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal("invalid param", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dbx, err := sqlx.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer dbx.Close()

	sql := `INSERT INTO users (username, full_name, email) VALUES (:username, :full_name, :email);`
	in := createUserRequestSqlx(req)

	tx := dbx.MustBegin()
	_, err = tx.NamedExec(sql, in)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	tx.Commit()

	rsp := newUserResponseSqlx(in)
	ctx.JSON(http.StatusOK, rsp)
}

func newUserResponseEnt(user *ent.User) userResponse {
	return userResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

// sqlc////////////////////////////////////////////
// 機能が多い(学習コストが高い)
// わかりやすい公式ドキュメントがあるhttps://entgo.io/ja/docs/getting-started
// トランザクションはないため実装する必要がある
// ここ最近急激にスターを伸ばしている
// sqlc////////////////////////////////////////////
func entCreateUser(ctx *gin.Context, config util.Config) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal("invalid param", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	conn, err := ent.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer conn.Close()

	user, err := conn.User.
		Create().
		SetUsername(req.Username).
		SetFullName(req.FullName).
		SetEmail(req.Email).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponseEnt(user)
	ctx.JSON(http.StatusOK, rsp)
}

func gormCreateUser(ctx *gin.Context, config util.Config) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Fatal("invalid param", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)
}
