package piscine

import "strings"

// func RepeatAlpha(s string) string {
// 	result:=""

// 	for _,c:=range s{
// 		var count int
// 		if c>='a'&&c<='z'{
// 			count=int(c-'a'+1)
// 		} else if c>='A'&&c<='Z'{
// 			count=int(c-'A'+1)
// 		} else {
// 			continue

//			}
//			for i:=0;i<count;i++{
//				result+=string(c)
//			}
//		}
//		return result
//	}

func Sep(sen string) []string {
	words := strings.Split(sen, "e")
	return words
}
