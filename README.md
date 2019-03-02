# definitionGame

AWS Lambda handler for the Alexa skill "Definition Game", an exercise in using Go with Alexa.

## What I want it to do

### The MVP
This game picks a entry from an array of words, calls the Oxford English Dictionary API to get the first definition of that word, and then send the text of that definition to the Alexa-enabled device. The user will then guess what word is being defined. If the answer is correct, they will receive congratulations.

### Future upgrade wish list.
- Keeping score (current and all-time)
- Leaderboard
- Q&A abilities, particularly ability for user to get hints (in the form of first letter of word, or additional definitions)
- "Spelling Bee" mode

## Dumb notes
- This is a spare time/de-stressing project. Hopefully it fulfills that function.
- Trying to do a lot of testing with this as well, because more testing is always good.