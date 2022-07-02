package _752_open_the_lock

func openLock(deadends []string, target string) int {
	return BFS("0000", target, deadends)
}

func BFS(start string, target string, deadends []string) int {
	q := make([]string, 0)
	visited := map[string]bool{}
	deads := map[string]bool{}
	for _, deadend := range deadends {
		deads[deadend] = true
	}

	step := 0

	q = append(q, start)
	visited[start] = true

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if _, exist := deads[cur]; exist {
				continue
			}

			if cur == target {
				return step
			}

			for j := 0; j < len(cur); j++ {

				s := []byte(cur)
				if cur[j] == '0' {
					s[j] = '9'
				} else {
					s[j]--
				}

				if _, exist := visited[string(s)]; !exist {
					q = append(q, string(s))
					visited[string(s)] = true
				}

				s = []byte(cur)
				if cur[j] == '9' {
					s[j] = '0'
				} else {
					s[j]++
				}

				if _, exist := visited[string(s)]; !exist {
					q = append(q, string(s))
					visited[string(s)] = true
				}

			}
		}
		step++
	}
	return -1
}
