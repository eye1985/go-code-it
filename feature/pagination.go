package feature

import (
	"codepocket/database"
	"math"
)

func Pagination(count float64, hitPerPage float64, start float64, userAndCodes []database.UserAndCode) *database.Pagination {
	totalPage := math.Ceil(count / hitPerPage)
	currentPage := math.Ceil(start / hitPerPage)

	next := start + hitPerPage
	prev := start - hitPerPage

	if len(userAndCodes) == 0 {
		totalPage = 0
		currentPage = 0
	}

	if prev < 0 {
		prev = 0
	}

	return &database.Pagination{
		Codes:       userAndCodes,
		CurrentPage: int16(currentPage),
		NextStart:   int16(next),
		PrevStart:   int16(prev),
		TotalPage:   int16(totalPage),
	}
}
