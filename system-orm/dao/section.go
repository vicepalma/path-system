package dao

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/vicepalma/path-system/system-orm/model"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllSection is a function to get a slice of record(s) from section table in the path-system database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllSection(ctx context.Context, page, pagesize int, order string) (results []*model.Section, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Section{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetSection is a function to get a single record from the section table in the path-system database
// error - ErrNotFound, db Find error
func GetSection(ctx context.Context, argID int32) (record *model.Section, err error) {
	record = &model.Section{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddSection is a function to add a single record to section table in the path-system database
// error - ErrInsertFailed, db save call failed
func AddSection(ctx context.Context, record *model.Section) (result *model.Section, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateSection is a function to update a single record from section table in the path-system database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateSection(ctx context.Context, argID int32, updated *model.Section) (result *model.Section, RowsAffected int64, err error) {

	result = &model.Section{}
	db := DB.First(result, "id = ?", argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteSection is a function to delete a single record from section table in the path-system database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteSection(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Section{}
	db := DB.First(record, "id = ?", argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
