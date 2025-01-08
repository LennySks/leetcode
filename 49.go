package main

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}

//func groupAnagrams(strs []string) [][]string {
//	//hash := make(map[string]int)
//	anagrams := make(map[string][]string)
//	for _, str := range strs {
//		chars := strings.Split(str, "")
//		sort.Strings(chars)
//		sortedStr := strings.Join(chars, "")
//		anagrams[sortedStr] = append(anagrams[sortedStr], str)
//		fmt.Println(anagrams)
//	}
//
//	var groups [][]string
//	for _, group := range anagrams {
//		groups = append(groups, group)
//	}
//	return groups
//}

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
