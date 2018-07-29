// see https://leetcode.com/problems/jewels-and-stones/description/

package main

/* use strings */
// import (
//     "fmt"
//     "strings"
// )

// func numJewelsInStones(J string, S string) int {
//     count := 0;
//     for _, v := range J {
//         count += strings.Count(S, string(v))
//     }
//     return count;
// }

/* use map */
import "fmt"

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	for _, v := range J {
        m[v] = true
    }
    count := 0;
    for _, v := range S {
        if m[v] {
			count ++
		}
    }
    return count;
}


func main() {
    S := "aaAsssfasBA";
    J := "aA";
    num := numJewelsInStones(J, S)
    fmt.Println(num);
}
