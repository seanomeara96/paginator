package paginator

import (
	"net/http"
	"strconv"
)

/*
Paginate pulls page, limit, and offset values from request params and calculates the page number using the limit and offset if not supplied.
if strconv.Atoi errors on either limit, offset or page then this will return defaultLimit, 0, 0
*/
func Paginate(r *http.Request, defaultLimit int) (limit int, offset int, page int) {
	q := r.URL.Query()

	limit = defaultLimit
	offset = 0
	page = 1

	_limit := q.Get("limit")
	if _limit != "" {
		l, err := strconv.Atoi(_limit)
		if err != nil {
			return defaultLimit, 0, 0
		} else {
			limit = l
		}
	}

	_offset := q.Get("offset")
	if _offset != "" {
		o, err := strconv.Atoi(_offset)
		if err != nil {
			return defaultLimit, 0, 0
		} else {
			offset = o
		}
	}

	_page := q.Get("page")
	if _page != "" {
		p, err := strconv.Atoi(_page)
		if err != nil {
			return defaultLimit, 0, 0
		} else {
			page = p
			offset = (limit * (p - 1))
		}
	}

	return limit, offset, page
}
