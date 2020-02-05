# fizzbuzz

This is a Rest API for an enhanced version of FizzBuzz algorithm.

Everything is customizable (
- the limit, 
- the numbers whose multipliers should be replaced with words 
- and the words).

## Routes
POST /fizzbuzz

body : 
```json
{
	"int1" : 3,
	"int2" : 5,
	"limit" : 100,
	"str1" : "fizzz",
	"str2" : "buzzz"
}
```
