import "sort"

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := make([]string, 0)
	current := folder[0]
	res = append(res, current)
	if len(folder) > 1 {
		for i := 1; i < len(folder); i++ {
			if current != folder[i] && !strings.HasPrefix(folder[i], current+"/") {
				res = append(res, folder[i])
				current = folder[i]
			}
		}
	}
	return res
}
