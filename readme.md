___________________________________________________________________________________________________________
How to Run this Project
___________________________________________________________________________________________________________
Install Go (>=1.16)
Download This Project
Run go build in the project root
This would create an exeutable file by the name of mobydick
We can run main file using the command go run main.go "pathoffile" or
Run the executable ./mobydick "pathoffile"
if pathoffile is not provided code will take default file "mobydick.txt" as input
___________________________________________________________________________________________________________
To run Test Cases:-
Run all test cases :-go test ./...
check coverage of all testcases:- go test -coverprofile fmtcoverage.html fmt

___________________________________________________________________________________________________________
Ouput:-

Currently, we have set the number of unique words for which we want to find frequencies to 20. As an extension we can also pass it as an argument along with the file path to change the number at run time.

The output of the program would be frequency and the word of first 20 words with highest frequencies in their descending order

Format:
<freq> <word>