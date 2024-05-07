package misc

type Pagination interface {
	GetSkip() uint
	GetLimit() uint
}

type page struct {
	Skip  int `json:"skip" form:"skip"`
	Limit int `json:"limit" form:"limit"`
}

func NewPage(skip, limit int) Pagination {
	return page{Skip: skip, Limit: limit}
}

func NewOffsetPage(page, size int64) Pagination {
	offset := (page - 1) * size
	return NewPage(int(offset), int(size))
}

func (p page) GetSkip() uint {
	if p.Skip < 0 {
		return 0
	}
	return uint(p.Skip)
}

func (p page) GetLimit() uint {
	if p.Limit <= 0 {
		return 10
	}
	return uint(p.Limit)
}

type CursorPagination interface {
	GetLimit() uint
	GetCursor() string
	After() bool
}

type cursorPagination struct {
	Limit   int    `json:"limit" form:"limit"`
	Cursor  string `json:"cursor" form:"id"`
	IsAfter bool   `json:"after" form:"after"`
}

func NewCursorPagination(limit int, cursor string, isAfter bool) cursorPagination {
	return cursorPagination{Limit: limit, Cursor: cursor, IsAfter: isAfter}
}

func (p cursorPagination) GetLimit() uint {
	if p.Limit <= 0 {
		return 10
	}
	return uint(p.Limit)
}

func (c cursorPagination) GetCursor() string {
	return c.Cursor
}

func (c cursorPagination) After() bool {
	return c.IsAfter
}
