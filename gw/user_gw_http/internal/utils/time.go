package utils

import (
	"dolott_user_gw_http/internal/types"
	"fmt"
	"strconv"
	"time"
)

func GetTimebyTimestamp(timestamp int64) (*time.Time, *types.Error) {
	i, err := strconv.ParseInt(fmt.Sprintf("%d", timestamp), 10, 64)
	if err != nil {
		return nil, types.NewBadRequestError("internal issue, error code #1")
	}
	tm := time.Unix(i, 0)
	return &tm, nil
}

func ParseTime(timeStr string, errMsg string) (time.Time, *types.Error) {
	timeInt64, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, types.NewInternalError(errMsg)
	}
	return time.Unix(timeInt64, 0), nil
}
