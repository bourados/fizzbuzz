# fizzbuzz

This is a Rest API for an enhanced version of FizzBuzz algorithm written in Go Lang.

Everything is customizable (
- the limit, 
- the numbers whose multipliers should be replaced with words 
- and the words).

## Routes
#### POST /fizzbuzz

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

returns an array of numbers with from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2. 

#### GET /stat
returns the most used request and the number of hits
