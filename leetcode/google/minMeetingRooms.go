package leetcode

import (
	"sort"
)

func MinMeetingRooms(intervals [][]int) int {
	meetingRooms := []int{}
	intervals = sortByFirstValue(intervals)
	for _, interval := range intervals {
		if len(meetingRooms) == 0 {
			meetingRooms = append(meetingRooms, interval[1])
		} else {
			roomFound := false
			for key, room := range meetingRooms {
				if room <= interval[0] {
					meetingRooms[key] = interval[1]
					roomFound = true
					break
				}
			}
			if !roomFound {
				meetingRooms = append(meetingRooms, interval[1])
			}
		}
	}
	return len(meetingRooms)
}

func sortByFirstValue(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	return intervals
}
