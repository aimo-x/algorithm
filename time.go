package algorithm

import (
	"time"
)

// Time algorithm
type Time struct {
}

// GetZeroTime 获取某一天的0点时间
func (t Time) GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetFirstDateOfMonth 获取某月的第一天
func (t Time) GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return t.GetZeroTime(d)
}

// GetLastDateOfMonth 获取某月的最后一天
func (t Time) GetLastDateOfMonth(d time.Time) time.Time {
	return t.GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// ConsecutiveDays 判断日期中是否连续了 nd 天,并返回 <= nd 的最打连续天数
func (t Time) ConsecutiveDays(ds []time.Time, nd int) (ok bool, lx int) {
	lx = 1
	for i, va := range ds {
		if i != 0 {
			if va.Month() == ds[i-1].Month() { // 月份相同
				if va.Day() == ds[i-1].Day()+1 { // 日期是否连续
					lx = lx + 1 // 增加
				} else { // 不满足连续
					if lx >= nd { // 判断是否连续nd 天， 满足就跳出
						ok = true
						return
					}
					lx = 1 // 重置
				}
			} else if int(va.Month()) == int(ds[i-1].Month())+1 { // 月份连续
				if t.GetLastDateOfMonth(ds[i-1]).Day() == ds[i-1].Day() && t.GetFirstDateOfMonth(va).Day() == va.Day() { // 同时满足 最后一天，和第一天
					lx = lx + 1
				} else {
					if lx >= nd { // 判断是否连续nd 天 满足就跳出
						ok = true
						return
					}
					lx = 1 // 重置
				}
			} else {
				lx = 1 // 重置
			}
		}
	}
	if lx >= nd { // 判断是否连续nd 天
		ok = true
		return
	}
	return
}
