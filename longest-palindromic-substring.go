package main

import "fmt"

func expandAroundCenter(s string, left, right int) string {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left+1 : right]
}

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }

    var result string
    for i := 0; i < len(s); i++ {
        palindrome1 := expandAroundCenter(s, i, i)
        palindrome2 := expandAroundCenter(s, i, i+1)
        
        if len(palindrome1) > len(result) {
            result = palindrome1
        }
        if len(palindrome2) > len(result) {
            result = palindrome2
        }
    }
    return result
}

func main() {
    str := "babad"
    fmt.Println("Longest Palindromic Substring:", longestPalindrome(str))
}
