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

// GetAllManyCategoryHasManySection is a function to get a slice of record(s) from many_category_has_many_section table in the path-system database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllManyCategoryHasManySection(ctx context.Context, page, pagesize int, order string) (results []*model.ManyCategoryHasManySection, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ManyCategoryHasManySection{})
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

// GetManyCategoryHasManySection is a function to get a single record from the many_category_has_many_section table in the path-system database
// error - ErrNotFound, db Find error
func GetManyCategoryHasManySection(ctx context.Context, argIDCategory int32, argIDSection int32) (record *model.ManyCategoryHasManySection, err error) {
	record = &model.ManyCategoryHasManySection{}
	if err = DB.First(record, argIDCategory, argIDSection).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddManyCategoryHasManySection is a function to add a single record to many_category_has_many_section table in the path-system database
// error - ErrInsertFailed, db save call failed
func AddManyCategoryHasManySection(ctx context.Context, record *model.ManyCategoryHasManySection) (result *model.ManyCategoryHasManySection, RowsAffected int64, err error) {
	db := DB.Create(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateManyCategoryHasManySection is a function to update a single record from many_category_has_many_section table in the path-system database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateManyCategoryHasManySection(ctx context.Context, argIDCategory int32, argIDSection int32, updated *model.ManyCategoryHasManySection) (result *model.ManyCategoryHasManySection, RowsAffected int64, err error) {

	result = &model.ManyCategoryHasManySection{}
	db := DB.First(result, "id_category = ?", argIDCategory, "id_section = ?", argIDSection)
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

// DeleteManyCategoryHasManySection is a function to delete a single record from many_category_has_many_section table in the path-system database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteManyCategoryHasManySection(ctx context.Context, argIDCategory int32, argIDSection int32) (rowsAffected int64, err error) {

	record := &model.ManyCategoryHasManySection{}
	db := DB.First(record, "id_category = ?", argIDCategory, "id_section = ?", argIDSection)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
