FizzBuzz Server
===

### requirement
go version go1.10.2 or greater

### get started
```
$>go get github.com/mdebelle/fizzbuzz
$>cd $GOPATH/src/github.com/mdebelle/fizzbuzz
$>go install
$>$GOPATH/bin/fizzbuzz
```

The server package handles requests on `http://localhost:8080/fizzbuzz`

### example

GET `http://localhost:8080/fizzbuzz?string1=fizz&int1=3&string2=buzz&int2=5&limit=100`

response `200 OK`
```json
[
    "1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
    "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz",
    "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz",
    "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz",
    "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz",
    "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz",
    "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz",
    "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz",
    "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz",
    "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"
]
```

if you want you can change the number of motifs with

OPTIONS `http://localhost:8080/fizzbuzz?motifs=1`

response `200 OK`

GET `http://localhost:8080/fizzbuzz?string1=fizz&int1=3&limit=100`

response `200 OK`
```json
[
    "1", "2", "fizz", "4", "5", "fizz", "7", "8", "fizz", "10",
    "11", "fizz", "13", "14", "fizz", "16", "17", "fizz", "19", "20",
    "fizz", "22", "23", "fizz", "25", "26", "fizz", "28", "29", "fizz",
    "31", "32", "fizz", "34", "35", "fizz", "37", "38", "fizz", "40",
    "41", "fizz", "43", "44", "fizz", "46", "47", "fizz", "49", "50",
    "fizz", "52", "53", "fizz", "55", "56", "fizz", "58", "59", "fizz",
    "61", "62", "fizz", "64", "65", "fizz", "67", "68", "fizz", "70",
    "71", "fizz", "73", "74", "fizz", "76", "77", "fizz", "79", "80",
    "fizz", "82", "83", "fizz", "85", "86", "fizz", "88", "89", "fizz",
    "91", "92", "fizz", "94", "95", "fizz", "97", "98", "fizz", "100"
]
```


