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

		if time+T > N {
			continue
		}

		isAppended := false

		for i := range schedules {
			if schedules[i].time > 0 {
				schedules[i].time--
			}

			if schedules[i].time <= 0 {
				cpy := make([]int, len(schedules[i].pay))
				copy(cpy, schedules[i].pay)

				schedules = append(schedules, Schedule{time: schedules[i].time, pay: cpy}) // 추가하지 않는 경우
				schedules[i].pay = append(schedules[i].pay, P)                             // 추가하는 경우
				schedules[i].time += T
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
		// fmt.Println(schedules)
	}

	ans := 0
	for _, schedule := range schedules {
		ret := SumOfArray(schedule.pay)
		if ret > ans {
			ans = ret
		}
	}

	fmt.Println(ans)
}
