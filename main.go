package main

import (
    "github.com/01-edu/z01"
)

// Split function splits the string s by the separator sep and returns a slice of strings.
func Split(s, sep string) []string {
    var result []string
    var currentWord []rune
    sepLen := len(sep)

    for i := 0; i < len(s); {
   	 if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
   		 // If the separator is found, add the current word to the result
   		 if len(currentWord) > 0 {
   			 result = append(result, string(currentWord))
   			 currentWord = nil // Reset current word for the next word
   		 }
   		 i += sepLen // Skip the separator
   	 } else {
   		 // If not a separator, add the character to the current word
   		 currentWord = append(currentWord, rune(s[i]))
   		 i++
   	 }
    }

    // Add the last word if there's any remaining
    if len(currentWord) > 0 {
   	 result = append(result, string(currentWord))
    }

    return result
}

func printStringSlice(slice []string) {
    for i, word := range slice {
   	 for _, r := range word {
   		 z01.PrintRune(r)
   	 }
   	 if i < len(slice)-1 {
   		 z01.PrintRune(' ')
   	 }
    }
    z01.PrintRune('\n')
}

func main() {
    s := "HelloHAhowHAareHAyou?"
    result := Split(s, "HA")
    printStringSlice(result)
}
