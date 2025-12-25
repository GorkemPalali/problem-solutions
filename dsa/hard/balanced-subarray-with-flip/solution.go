package balancedsubarraywithflip

func findMaxSubarray(nums []int32) int32 {
	n := len(nums)
	offset := n
	// Use a large negative sentinel for "not seen".
	const INF = int(-1 << 30)

	first := make([]int, 2*n+1)
	for i := range first {
		first[i] = INF
	}
	first[0+offset] = -1

	prefix := 0
	best := 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			prefix -= 1
		} else {
			prefix += 1
		}

		// 1) Sum == 0 → look up earliest same prefix
		if idx := first[prefix+offset]; idx != INF {
			if i-idx > best {
				best = i - idx
			}
		}

		// 2) Sum == +2 → earliest prefix == prefix-2
		if (prefix-2)+offset >= 0 && (prefix-2)+offset < len(first) {
			if idx := first[prefix-2+offset]; idx != INF {
				if i-idx > best {
					best = i - idx
				}
			}
		}

		// 3) Sum == -2 → earliest prefix == prefix+2
		if (prefix+2)+offset >= 0 && (prefix+2)+offset < len(first) {
			if idx := first[prefix+2+offset]; idx != INF {
				if i-idx > best {
					best = i - idx
				}
			}
		}

		if first[prefix+offset] == INF {
			first[prefix+offset] = i
		}
	}
	return int32(best)
}
