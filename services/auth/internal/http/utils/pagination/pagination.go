package pagination

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"auth/pkg/lib/utils/consts"
)

func ParsePaginationQuery(r *http.Request) (Query, error) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "20"
	}

	if page == "0" || limit == "0" {
		return Query{}, consts.ErrInvalidQueryParam
	}

	pageUInt, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		return Query{}, consts.ErrInvalidQueryParam
	}

	limitUInt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		return Query{}, consts.ErrInvalidQueryParam
	}

	return Query{
		Page:  pageUInt,
		Limit: limitUInt,
	}, nil
}

func parseFullURL(r *http.Request, query map[string]string) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	q := url.Values{}
	for k, v := range query {
		q.Set(k, v)
	}

	return fmt.Sprintf("%s://%s%s?%s", scheme, r.Host, r.URL.Path, q.Encode())
}

func BuildPaginationLinks(r *http.Request, p Query, totalPages uint64) Links {
	firstPage := parseFullURL(
		r, map[string]string{
			"page":  "1",
			"limit": strconv.FormatUint(p.Limit, 10),
		},
	)
	lastPage := parseFullURL(
		r, map[string]string{
			"page":  strconv.FormatUint(totalPages, 10),
			"limit": strconv.FormatUint(p.Limit, 10),
		},
	)

	var prevPage, nextPage *string
	if p.Page > 1 {
		prev := parseFullURL(
			r, map[string]string{
				"page":  strconv.FormatUint(p.Page-1, 10),
				"limit": strconv.FormatUint(p.Limit, 10),
			},
		)
		prevPage = &prev
	}
	if p.Page < totalPages {
		next := parseFullURL(
			r, map[string]string{
				"page":  strconv.FormatUint(p.Page+1, 10),
				"limit": strconv.FormatUint(p.Limit, 10),
			},
		)
		nextPage = &next
	}

	return Links{
		Prev:  prevPage,
		Next:  nextPage,
		First: firstPage,
		Last:  lastPage,
	}
}
