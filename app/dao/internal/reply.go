// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// ReplyDao is the manager for logic model data accessing and custom defined data operations functions management.
type ReplyDao struct {
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns ReplyColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ReplyColumns defines and stores column names for table gf_reply.
type ReplyColumns struct {
	Id         string // 回复ID
	ParentId   string // 回复对应的上一级回复ID(没有的话默认为0)
	Title      string // 回复标题
	Content    string // 回复内容
	TargetType string // 评论类型: content, reply
	TargetId   string // 对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id
	UserId     string // 网站用户ID
	ZanCount   string // 赞
	CaiCount   string // 踩
	CreatedAt  string // 创建时间
	UpdatedAt  string //
}

//  replyColumns holds the columns for table gf_reply.
var replyColumns = ReplyColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Title:      "title",
	Content:    "content",
	TargetType: "target_type",
	TargetId:   "target_id",
	UserId:     "user_id",
	ZanCount:   "zan_count",
	CaiCount:   "cai_count",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewReplyDao creates and returns a new DAO object for table data access.
func NewReplyDao() *ReplyDao {
	return &ReplyDao{
		Group:   "default",
		Table:   "gf_reply",
		Columns: replyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReplyDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReplyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
