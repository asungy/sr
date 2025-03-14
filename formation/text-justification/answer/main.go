package main

import "fmt"
import "strings"

// Given an array of words and a fixed line length maxWidth, format the text so that each line is exactly maxWidth characters and fully justified (left and right aligned). 
// 
// Requirements:
// - Pack as many words as possible in each line.
// - Distribute extra spaces between words as evenly as possible. If the spaces don't divide evenly, place the extra spaces in the leftmost gaps.
// - If a line contains only one word, left-align it and pad the rest of the line with spaces to reach maxWidth.
// - For the last line, words should be left-aligned, with no extra spaces between words. Pad the end with spaces to make the line maxWidth long.
// 
// 
// EXAMPLE(S)
// ["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"], k = 16
// 
// returns
// ["the  quick brown", // (2 spaces, 1 space)
// "fox  jumps  over", // (2 spaces, 2 spaces)
// "the lazy dog    "]  // (left justify, fill the end with 4 spaces)

func textJustification(words []string, k int) []string {
    type Info struct {
         // words per line
        words []string
        // number of spaces between each word by default
        spaceCount int
        // number of extra spaces that have to be distributed (this is less than len(words) - 1)
        remainingSpaces int
    }

    includeWord := func(words []string, nextWord string, k int) bool {
        if len(words) == 0 {
            // assuming all words are less than k
            return true
        }

        wordsLength := len(nextWord)
        for _, word := range words {
            wordsLength += len(word)
        }

        return wordsLength + len(words) <= k
    }

    toInfo := func(words []string, k int) Info {
        wordsLength := 0
        for _, word := range words {
            wordsLength += len(word)
        }

        spaceBlocks := len(words) - 1
        if spaceBlocks == 0 {
            return Info {
                words, 0, 0,
            }
        }

        spaceCount := (k - wordsLength) / spaceBlocks
        remainingSpaces := (k - wordsLength) % spaceBlocks

        return Info {
            words,
            spaceCount,
            remainingSpaces,
        }
    }

    fullyJustify := func(info Info) string {
        var sb strings.Builder
        for i, word := range info.words[:len(info.words)-1] {
            sb.WriteString(word)
            for _ = range info.spaceCount {
                sb.WriteRune(' ')
            }
            if i < info.remainingSpaces {
                sb.WriteRune(' ')
            }
        }
        sb.WriteString(info.words[len(info.words)-1])
        return sb.String()
    }

    leftJustify := func(info Info, k int) string {
        var sb strings.Builder
        charCount := 0
        for _, word := range info.words {
            sb.WriteString(word)
            sb.WriteRune(' ')
            charCount += len(word) + 1
        }
        for _ = range (k - charCount) {
            sb.WriteRune(' ')
        }
        return sb.String()
    }

    lines := []Info{}
    currentWords := []string{}
    for _, word := range words {
        if includeWord(currentWords, word, k) {
            currentWords = append(currentWords, word)
        } else {
            lines = append(lines, toInfo(currentWords, k))
            currentWords = []string{word}
        }
    }
    
    if len(currentWords) > 0 {
        lines = append(lines, toInfo(currentWords, k))
    }
    
    result := []string{}
    for _, line := range lines[:len(lines)-1] {
        result = append(result, fullyJustify(line))
    }
    result = append(result, leftJustify(lines[len(lines)-1], k))

    return result
}

func main() {
    // words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
    words := []string{
        "Build", "a", "man", "a", "fire,", "and", "he'll", "be", "warm", "for", "a", "day.",
        "Set", "a", "man", "on", "fire,", "and", "he'll", "be", "warm", "for", "the", "rest", "of", "his", "life.",
    }
    k := 20
    lines := textJustification(words, k)
    for _, line := range lines {
        fmt.Printf("|%s|\n", line)
    }
}
