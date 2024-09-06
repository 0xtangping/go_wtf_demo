package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleFieldNames          = builder.RawFieldNames(&Article{})
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsExpectAutoSet   = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	articleRowsWithPlaceHolder = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	articleModel interface {
		Insert(ctx context.Context, data *Article) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Article, error)
		Update(ctx context.Context, data *Article) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context, page, pageSize int64) ([]*Article, error)
	}

	defaultArticleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Article struct {
		Id        int64     `db:"id"`
		Author    string    `db:"author"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func newArticleModel(conn sqlx.SqlConn) *defaultArticleModel {
	return &defaultArticleModel{
		conn:  conn,
		table: "`article`",
	}
}

func (m *defaultArticleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (*Article, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
	var resp Article
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleModel) Insert(ctx context.Context, data *Article) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, articleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Author, data.Title, data.Content)
	return ret, err
}

func (m *defaultArticleModel) Update(ctx context.Context, data *Article) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Author, data.Title, data.Content, data.Id)
	return err
}

func (m *defaultArticleModel) FindAll(ctx context.Context, pageSize,page int64) ([]*Article, error) {
    query := fmt.Sprintf(`
        SELECT %s FROM %s 
        ORDER BY created_at DESC 
        LIMIT ? OFFSET ?`, articleRows, m.table)
    
    var resp []*Article
	offset := (page - 1) * pageSize

    err := m.conn.QueryRowsCtx(ctx, &resp, query, pageSize, offset)
    if err != nil {
        return nil, err
    }
    return resp, nil
}


func (m *defaultArticleModel) tableName() string {
	return m.table
}
