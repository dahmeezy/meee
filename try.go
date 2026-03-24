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

func ToJadenCase(str string) string {
  res:=""
  if str==""{
    return " "
  } else{
    words:=strings.Fields(str)
    for _,word:=range words{
      for i,c :=range word{
        if i==0 && (c>='a'&&c<='z'){
          res+=string(int(c)-32)
        } else{
          res+=string(c)
        }
      }
    }
  }  
  return res // your code here...
  
  
}
