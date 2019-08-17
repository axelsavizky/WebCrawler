# Web Crawler for AgileEngine

### How to run
You only need to execute the main inside the src folder. The first argument is the original file path, the second argument is the diff file path and the third argument is the id to find.

#### Clarifications
  - In order to compare html nodes, I converted them to string and compare with levenshtein distance
  - This code has no tests! That's not a quality code. I know that, but i ran out of time, and the code works, so i decided to commit this anyway.
  - I decided to panic in an error like "opening a file". This maybe is not the best choice, but i think this is not the most important part of the exercise.
  - Maybe there are many things to improve, like the amount of arguments of the search function. I couldn't refactor it because of the time, but i marked some parts with TODOs
  - For levenshtein distance I used a library
  - I did all the search algorithm in the main file. This is not a good way to organize a Golang project, but again, i ran out of time so i decided to let all in the main file and commit a code that works!

  For any doubts you can send me an email to: axel.savizky@gmail.com
