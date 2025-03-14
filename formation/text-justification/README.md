Given an array of words and a fixed line length maxWidth, format the text so that each line is exactly maxWidth characters and fully justified (left and right aligned). 

Requirements:
- Pack as many words as possible in each line.
- Distribute extra spaces between words as evenly as possible. If the spaces don't divide evenly, place the extra spaces in the leftmost gaps.
- If a line contains only one word, left-align it and pad the rest of the line with spaces to reach maxWidth.
- For the last line, words should be left-aligned, with no extra spaces between words. Pad the end with spaces to make the line maxWidth long.


EXAMPLE(S)
["the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"], k = 16

returns
["the  quick brown", // (2 spaces, 1 space)
"fox  jumps  over", // (2 spaces, 2 spaces)
"the lazy dog    "]  // (left justify, fill the end with 4 spaces)

# Instructions
1. `git clone https://github.com/asungy/sr`
2. `cd sr/formation/text-justification`
3. Implement quicksort in `playground/main.go`
4. Run `go run main.go` to verify output is as expected.
5. Reference `answer/` for a solution if needed.
