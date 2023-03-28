package userstorage

import (
	"context"
	"fmt"
)

type NativeStore interface {
	Check(context context.Context, data CheckRecordData) bool
}

type NativeStoreBiz struct {
	NativeStore NativeStore
}

func NewNativeStoreBiz(nativeStore NativeStore) *NativeStoreBiz {
	return &NativeStoreBiz{NativeStore: nativeStore}
}

type CheckRecordData struct {
	TableName  string
	ColumnName string
	Value      string
}

func (biz *NativeStoreBiz) CheckRecordIsExist(context context.Context, data CheckRecordData) bool {
	return biz.NativeStore.Check(context, data)
}

type RawDB struct {
	Count int `json:"count" gorm:"not null;column:count"`
}

func (s *sqlStore) Check(context context.Context, data CheckRecordData) bool {
	var rawDb RawDB
	fmt.Println(fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE %s = ? LIMIT 1", data.TableName, data.ColumnName), data.Value)
	s.db.Raw(fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE %s = ? LIMIT 1", data.TableName, data.ColumnName), data.Value).Scan(&rawDb)
	return rawDb.Count > 0
}
