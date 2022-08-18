package o

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 找出两个人共同的空闲时间
 * Sample input:
 * 甲的已预定时间段：[["9:30", "10:30"], ["12:00", "13:00"], ["16:00", "18:00"]]
 * 甲的上下班时间：["9:00", "20:00"]
 * 乙的已预定时间段：[["10:00", "11:30"], ["12:30", "14:30"], ["14:30", "15:00"], ["16:00", "17:00"]]
 * 乙的上下班时间：["10:00", "18:30"]
 * 会议时长：30
 */
func BookMeeting(aSchedule, bSchedule [][]string, aBound, bBound []string, minutes int) [][]string {
	// [["0:00", "9:00"], ["9:30", "10:30"], ["12:00", "13:00"], ["16:00", "18:00"], ["20:00", "24:00"]]
	aUsed := getUsedTimes(aSchedule, aBound)
	// [["0:00", "10:00"], ["10:00", "11:30"], ["12:30", "14:30"], ["14:30", "15:00"], ["16:00", "17:00"], ["18:30", "24:00"]]
	bUsed := getUsedTimes(bSchedule, bBound)

	i, j := 1, 1
	available := [][]string{}
	for i < len(aUsed) || j < len(bUsed) {
		if aUsed[i][0]-aUsed[i-1][1] < minutes {
			i++
			continue
		}

		if bUsed[j][0]-bUsed[j-1][1] < minutes {
			j++
			continue
		}

		if aUsed[i][0]-bUsed[j-1][1] < minutes {
			i++
			continue
		}

		if bUsed[j][0]-aUsed[i-1][1] < minutes {
			j++
			continue
		}

		start := max(aUsed[i-1][1], bUsed[j-1][1])
		end := min(aUsed[i][0], bUsed[j][0])

		available = append(available, []string{convertMinutesToTime(start), convertMinutesToTime(end)})
		i++
		j++
	}

	return available
}

func getUsedTimes(schedule [][]string, bound []string) [][]int {
	used := [][]int{}
	start := []int{0, convertTimeToMinutes(bound[0])}
	used = append(used, start)
	for _, item := range schedule {
		time := []int{convertTimeToMinutes(item[0]), convertTimeToMinutes(item[1])}
		used = append(used, time)
	}

	end := []int{convertTimeToMinutes(bound[1]), 1440}
	used = append(used, end)

	return used
}

func convertTimeToMinutes(time string) int {
	parts := strings.Split(time, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])

	return hour*60 + minute
}

func convertMinutesToTime(minutes int) string {
	hour := minutes / 60
	minute := minutes % 60

	return fmt.Sprintf("%02d:%02d", hour, minute)
}
