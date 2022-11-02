package seedwork

type DeferredResult struct {
	lastInsertId int64
	rowsAffected int64
}

func (r *DeferredResult) SetLastInsertId(v int64) {
	r.lastInsertId = v
}

func (r *DeferredResult) SetRowsAffected(v int64) {
	r.rowsAffected = v
}

func (r DeferredResult) LastInsertId() (int64, error) {
	return r.lastInsertId, nil
}
func (r DeferredResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}
