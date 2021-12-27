package main

import (
	"fmt"
)

type Schedule struct {
	time int
	pay  []int
}

func SumOfArray(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func main() {
	var N int

	fmt.Scan(&N)

	schedules := make([]Schedule, 0)

	for time := 0; time < N; time++ {
		var T, P int

		fmt.Scanf("%d %d\n", &T, &P)

		// You can't choice scehdule after quit your job. But all schedule's time ticking
		if time+T > N {
			for i := range schedules {
				if schedules[i].time > 0 {
					schedules[i].time--
				}
			}
			continue
		}

		// If this Schedule cannot be appended in any schedule
		isAppended := false

		// Iterate all Schedules and Modify or Create Schedule
		for i := range schedules {
			if schedules[i].time > 0 {
				schedules[i].time--
			}
			// fmt.Println(schedules[i], P, T)
			if schedules[i].time == 0 {
				// Copy current pay array
				pay_cpy := make([]int, len(schedules[i].pay))
				copy(pay_cpy, schedules[i].pay)

				// case 1. Not append this schedule 
				schedules = append(schedules, Schedule{time: schedules[i].time, pay: pay_cpy})

				// case 2. Append this schedule
				schedules[i].pay = append(schedules[i].pay, P)                             
				schedules[i].time += T

				// This schedule is appended to exists schedule. So it is no need to be created as new schedule
				isAppended = true
			} else {
				continue
			}
		}
		
		if !isAppended {
			schedules = append(schedules, Schedule{time: T})
			last_idx := len(schedules) - 1
			schedules[last_idx].pay = append(schedules[last_idx].pay, P)
		}
	}
	
	// Get bset schedule that gives biggest money 
	ans := 0
	for _, schedule := range schedules {
		ret := SumOfArray(schedule.pay)
		if ret > ans {
			ans = ret
		}
	}

	fmt.Println(ans)
}
