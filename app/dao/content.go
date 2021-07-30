// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"focus/app/dao/internal"
)

// contentDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type contentDao struct {
	*internal.ContentDao
}

var (
	// Content is globally public accessible object for table gf_content operations.
	Content = contentDao{
		internal.NewContentDao(),
	}
)

// Fill with you ideas below.
