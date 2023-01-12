# Video 11 show notes for Go Lang Base 62

* Link to this [page](https://bit.ly/rust_base62): `https://bit.ly/rust_base62`
* Link to related [video](https://youtu.be/1kbAUob3Bi0) TODO


## Tags

```
#golang #java #functionalprogramming #scala #base62

```



# Base62Encoder URLShortener



### Example URL shortener services:
* https://www.shorturl.at/
* https://bitly.com/
* https://tinyurl.com/

### URL shortener services and Base 62 explained:
* [WikiPedia URL Shortening](https://en.wikipedia.org/wiki/URL_shortening)
* [Base 62](https://en.wikipedia.org/wiki/Base62).

### URL shortener services and Base 62 tutorials with example code:
* [How a URL Shortening Application Works](https://dzone.com/articles/how-a-url-shortening-application-works)
* [Designing the Shortening URL system like Bit.ly, loading 6 billion clicks a month](https://itnext.io/designing-the-shortening-url-system-like-bit-ly-loading-6-billion-clicks-a-month-78b3e48eee8c)


### Videos in this series
* [First Video in Base62 language series 3rd Vlog video](https://www.youtube.com/watch?v=07Wkf9OZE3U)  [Show Notes]()
* [Second Video in Base62 language (Java) series 4th Vlog video](https://www.youtube.com/watch?v=sOhzb6OqyGA) [Show Notes](https://gist.github.com/RichardHightower/035fda0b65de540574e458dedf9dae6d)
* [Third Video in Base62 language (Scala) series 5th Vlog video](https://www.youtube.com/watch?v=TlQZn9MajlY)  [Show Notes](https://gist.github.com/RichardHightower/1d64d0c958a7643c8b0b573c08138e1f)
* [Fourth Video in Base62 language (Clojure) series 6th Vlog video](https://www.youtube.com/watch?v=1kbAUob3Bi0) [Show notes](https://gist.github.com/RichardHightower/5b45e5162cf8295f73e71d67ad4a442d)
* [Fifth Video in Base62 language (Python) series 7th Vlog video](https://www.youtube.com/watch?v=Sdfxs8UsK3w) [Show notes](https://github.com/RichardHightower/python_base62)
* [Sixth Video in Base62 language (JavaScript/TypeScript) series 8th Vlog video](https://www.youtube.com/TBD) [Show notes](https://github.com/RichardHightower/ts_js_base62)
* [Seventh Video in Base62 language (Kotlin) series 9th Vlog video](https://www.youtube.com/TBD) [Show notes](https://github.com/RichardHightower/ts_js_base62)
* [Eighth Video in Base62 language (Rust) series 10th Vlog video](https://www.youtube.com/TBD) [Show notes](https://github.com/RichardHightower/https://github.com/RichardHightower/rust_base62)
* [Ninth Video in Base62 language (Go Lang) series 11th Vlog video](https://www.youtube.com/TBD) [Show notes](https://github.com/RichardHightower/golang_base62)



# Related articles
* [Comparing Basic FP support part 1 --Rick Hightower](https://www.linkedin.com/pulse/comparing-basic-fp-support-part-1-rick-hightower/)
* [Is Java a good FP language? Comparing Basic FP support part 2 --Rick Hightower](https://www.linkedin.com/pulse/java-good-fp-language-comparing-basic-support-part-2-rick-hightower/)
* [Translating to Clojure: a learning task (Part 1) --Tom Hicks](https://hickst.hashnode.dev/translating-to-clojure-a-learning-task-part-1)


# Where is Rick?
* [LinkedIn](https://www.linkedin.com/in/rickhigh/)
* [Rick's YouTube Channel](https://www.youtube.com/channel/UCgCx8XtYUGW9aSfzXhP2m6Q)
* [Where Rick works](http://www.cloudurable.com/)


## Title: Porting Base62Encoder/Decoder from Scala to Go Lang

Porting Base62Encoder/Decoder from Scala to Go Lang.

[Link To video](https://youtu.be/1kbAUob3Bi0) TODO

## Scala to Go Lang

### Main Method

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
def main(args: Array[String]): Unit = {
  val id = 12345678910L
  val strId = convertToEncodedString(id)
  val newId = convertToLong(strId)
  println(s"$id $strId $newId")

  val longURL = "https://www.somewebiste.com/dp/..."
  val urlId = Math.abs(longURL.hashCode)
  val shortHandle = convertToEncodedString(urlId)
  println(s"$urlId $shortHandle ${convertToLong(shortHandle)}")

}
```
</sub>
</td>

<td >
<sub>

```go
func main() {
    fmt.Printf("%s %d\n", ConvertToEncodedString(12345678910), ConvertToLong("DTVD3O"))
    longUrl := "https://www.somewebiste.com/dp/0201616165/?_encoding=UTF8&pd_rd_w=vwEcs&content-id=amzn1.sym.8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_p=8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_r=BQ0KD40K57XG761DBNBA&pd_rd_wg=DtkHk&pd_rd_r=f94b60b7-9080-4065-b77f-6377ec854d17&ref_=pd_gw_ci_mcx_mi"
    urlId := hash(longUrl)
    shortHandle := ConvertToEncodedString(int64(urlId))
    fmt.Printf("url id %d short handle %s %d \n", urlId, shortHandle, ConvertToLong(shortHandle))
}
```

</sub>
</td>

</tr>
</table>



### convertToEncodedString

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
def convertToEncodedString(id: Long): String = {
  val builder: List[String] = List()
  val placeHolder = findStartBucket(id)
  val results = accumulateDigits(placeHolder, id, builder)
  val bucketValue = powDigitsBase(1)
  val digitIndex: Int = (results._2 / bucketValue).toInt
  val acc = results._2 - (bucketValue * digitIndex)
  val newBuilder: List[String] = appendSafeToList(results._3, digitIndex)
  //Put the remainder in the ones column
  val place1DigitIndex = (acc % bucketValue).toInt
  val finalBuilder = newBuilder ++ List(DIGITS(place1DigitIndex).toString)
  finalBuilder.mkString("")
}

private def accumulateDigits(placeHolder: Int, acc: Long,
                     builder: List[String]): (Int, Long, List[String]) = {
  if (!(placeHolder > 1)) {
    return (placeHolder, acc, builder)
  }
  val bucketValue = powDigitsBase(placeHolder)
  val digitIndex = (acc / bucketValue).toInt
  accumulateDigits(placeHolder - 1, acc - (bucketValue * digitIndex),
       appendSafeToList(builder, digitIndex))
}
```
</sub>
</td>

<td >
<sub>

```golang
func ConvertToEncodedString(id int64) string {
    var builder strings.Builder
    placeHolder := FindStartBucket(id)
    var bucketValue int64
    acc  := id
    var digitIndex int

    for placeHolder > 1 {
        bucketValue = PowDigitsBase(int64(placeHolder))
        digitIndex = int (acc / bucketValue)
        acc = acc - (bucketValue * int64(digitIndex))
        AppendSafe(&builder, digitIndex)
        placeHolder--
    }
    bucketValue = PowDigitsBase(1)
    digitIndex = (int) (acc / bucketValue)
    acc = acc - (bucketValue * int64(digitIndex))
    AppendSafe(&builder, digitIndex)

    //Put the remainder in the ones column
    digitIndex = (int) (acc % bucketValue)
    AppendSafe(&builder, digitIndex)
    return builder.String()
}
```

</sub>
</td>

</tr>
</table>


### findStartBucket

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
private def findStartBucket(value: Long): Int = {
  val i = Range(0, 15, 1).find(i => value < powDigitsBase(i.toLong))
  i.getOrElse(0)
}
```
</sub>
</td>

<td >
<sub>

```golang
func FindStartBucket(value int64) int{
    for i := 0; i < 15; i++ {
        if value < PowDigitsBase(int64(i)) {
            return i-1
        }
    }
    return 0
}
```

</sub>
</td>

</tr>
</table>



### powDigitsBase

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
private def powDigitsBase(exponent: Long): Long =
                        doPow(exponent, DIGITS.length)

private def doPow(exponent: Long, base: Int): Long = {
  if (exponent == 0) return 1
  doPow(exponent - 1, base) * base
}

```
</sub>
</td>

<td >
<sub>

```golang
func PowDigitsBase( exponent int64) int64 {
    return DoPow(exponent, int64(len(DIGITS)))
}

func DoPow( exponent int64, base int64) int64 {
    var result int64 = 1
    exp := exponent
    for exp != 0 {
        result *= base
        exp--
    }
    return result
}
```

</sub>
</td>

</tr>
</table>




### appendSafeToList

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
private def appendSafeToList(builder: List[String], digitIndex: Int): List[String] = {
  if (digitIndex != 0) builder ++ List((DIGITS(digitIndex)).toString)
  else if (builder.nonEmpty) builder ++ List((DIGITS(digitIndex)).toString)
  else builder
}
```
</sub>
</td>

<td >
<sub>

```golang
func AppendSafe( builder *strings.Builder,  digitIndex int) strings.Builder{
	char := DIGITS[digitIndex]

	if digitIndex != 0 {
		builder.WriteRune(char)
	} else {
		if builder.Len() > 0 {
			builder.WriteRune(char)
		}
	}
	return *builder
}
```

</sub>
</td>

</tr>
</table>




### convertToLong

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala

def convertToLong(strId: String): Long =
                            doConvertToLong(strId.toCharArray)

private def doConvertToLong(chars: Array[Char]): Long = {
  val (acc, _) = chars.reverse.foldLeft(0L, 0) { (pos, ch) =>
    val (acc, position) = pos
    val value = computeValue(ch, position)
    (acc + value, position + 1)
  }
  acc
}
```
</sub>
</td>

<td >
<sub>

```golang
func ConvertToLong( strId string) int64 {
	return DoConvertCharsToLong([]rune(strId))
}

func DoConvertCharsToLong(chars []rune) int64 {
 	var acc int64 = 0
	position  := 0
	for index := len(chars) -1; index > -1; index-- {
		c := chars[index]
		value :=  ComputeValue(c, position)
		acc += value
		position++
	}
	return acc
}
```
</sub>
</td>

</tr>
</table>




### computeValue

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
private def computeValue(c: Char, position: Int) = {
  val digitIndex = findIndexOfDigitInTable(c)
  val multiplier = powDigitsBase(position)
  digitIndex * multiplier
}
```
</sub>
</td>

<td >
<sub>

```golang
func ComputeValue(c rune, position int) int64 {
	digitIndex := FindIndexOfDigitInTable(c)
	multiplier  := PowDigitsBase(int64(position))
	return int64(digitIndex) * multiplier
}
```

</sub>
</td>

</tr>
</table>






----


## Go Lang full example

```golang
package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

var DIGITS =  [...]rune{
//0    1    2    3    4    5    6    7    8    9
'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
//10    11   12   13  14   15   16   17    18    19   20  21
'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
//22    23   24   25  26   27   28   29    30    31   32  33    34  35
'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
// 36  37  38   39   40   41    42    43   44  45   46    47
'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',  //Easy to add more characters if not using lookup tables.
// 48  49   50   51   52   53   54   55  56   57   58  59   60   61   // 62   63, 64
'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Printf("%s %d\n", ConvertToEncodedString(12345678910), ConvertToLong("DTVD3O"))
	longUrl := "https://www.somewebiste.com/dp/0201616165/?_encoding=UTF8&pd_rd_w=vwEcs&content-id=amzn1.sym.8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_p=8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_r=BQ0KD40K57XG761DBNBA&pd_rd_wg=DtkHk&pd_rd_r=f94b60b7-9080-4065-b77f-6377ec854d17&ref_=pd_gw_ci_mcx_mi"
	urlId := hash(longUrl)
	shortHandle := ConvertToEncodedString(int64(urlId))
	fmt.Printf("url id %d short handle %s %d \n", urlId, shortHandle, ConvertToLong(shortHandle))
}


func ConvertToEncodedString(id int64) string {
	var builder strings.Builder
	placeHolder := FindStartBucket(id)
	var bucketValue int64
	acc  := id
	var digitIndex int

	for placeHolder > 1 {
		bucketValue = PowDigitsBase(int64(placeHolder))
		digitIndex = int (acc / bucketValue)
		acc = acc - (bucketValue * int64(digitIndex))
		AppendSafe(&builder, digitIndex)
		placeHolder--
	}
	bucketValue = PowDigitsBase(1)
	digitIndex = (int) (acc / bucketValue)
	acc = acc - (bucketValue * int64(digitIndex))
	AppendSafe(&builder, digitIndex)

	//Put the remainder in the ones column
	digitIndex = (int) (acc % bucketValue)
	AppendSafe(&builder, digitIndex)
	return builder.String()
}

func FindIndexOfDigitInTable(c rune) int {
	index := -1
	for i := 0; i < len(DIGITS); i++ {
		if DIGITS[i] == c {
			index = i
			break
		}
	}
	return index
}

func FindStartBucket(value int64) int{
	for i := 0; i < 15; i++ {
		if value < PowDigitsBase(int64(i)) {
			return i-1
		}
	}
	return 0
}

func AppendSafe( builder *strings.Builder,  digitIndex int) strings.Builder{
	char := DIGITS[digitIndex]

	if digitIndex != 0 {
		builder.WriteRune(char)
	} else {
		if builder.Len() > 0 {
			builder.WriteRune(char)
		}
	}
	return *builder
}

func ConvertToLong( strId string) int64 {
	return ConvertCharsToLong([]rune(strId))
}

func ConvertCharsToLong(chars []rune) int64 {
 	var acc int64 = 0
	position  := 0
	for index := len(chars) -1; index > -1; index-- {
		c := chars[index]
		value :=  ComputeValue(c, position)
		acc += value
		position++
	}
	return acc
}

func ComputeValue(c rune, position int) int64 {
	digitIndex := FindIndexOfDigitInTable(c)
	multiplier  := PowDigitsBase(int64(position))
	return int64(digitIndex) * multiplier
}

func PowDigitsBase( exponent int64) int64 {
	return DoPow(exponent, int64(len(DIGITS)))
}

func DoPow( exponent int64, base int64) int64 {
 	var result int64 = 1
	exp := exponent
	for exp != 0 {
		result *= base
		exp--
	}
	return result
}

```






## Scala to Go Lang Functional

See https://github.com/robpike/filter


### Main Method

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
def main(args: Array[String]): Unit = {
  val id = 12345678910L
  val strId = convertToEncodedString(id)
  val newId = convertToLong(strId)
  println(s"$id $strId $newId")

  val longURL = "https://www.somewebiste.com/dp/..."
  val urlId = Math.abs(longURL.hashCode)
  val shortHandle = convertToEncodedString(urlId)
  println(s"$urlId $shortHandle ${convertToLong(shortHandle)}")

}
```
</sub>
</td>

<td >
<sub>

```go
func main() {
    fmt.Printf("%s %d\n", ConvertToEncodedString(12345678910), ConvertToLong("DTVD3O"))
    longUrl := "https://www.somewebiste.com/dp/0201616165/?_encoding=UTF8&pd_rd_w=vwEcs&content-id=amzn1.sym.8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_p=8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_r=BQ0KD40K57XG761DBNBA&pd_rd_wg=DtkHk&pd_rd_r=f94b60b7-9080-4065-b77f-6377ec854d17&ref_=pd_gw_ci_mcx_mi"
    urlId := hash(longUrl)
    shortHandle := ConvertToEncodedString(int64(urlId))
    fmt.Printf("url id %d short handle %s %d \n", urlId, shortHandle, ConvertToLong(shortHandle))
}
```

</sub>
</td>

</tr>
</table>



### convertToEncodedString

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala
def convertToEncodedString(id: Long): String = {
  val builder: List[String] = List()
  val placeHolder = findStartBucket(id)
  val results = accumulateDigits(placeHolder, id, builder)
  val bucketValue = powDigitsBase(1)
  val digitIndex: Int = (results._2 / bucketValue).toInt
  val acc = results._2 - (bucketValue * digitIndex)
  val newBuilder: List[String] = appendSafeToList(results._3, digitIndex)
  //Put the remainder in the ones column
  val place1DigitIndex = (acc % bucketValue).toInt
  val finalBuilder = newBuilder ++ List(DIGITS(place1DigitIndex).toString)
  finalBuilder.mkString("")
}

private def accumulateDigits(placeHolder: Int, acc: Long,
                     builder: List[String]): (Int, Long, List[String]) = {
  if (!(placeHolder > 1)) {
    return (placeHolder, acc, builder)
  }
  val bucketValue = powDigitsBase(placeHolder)
  val digitIndex = (acc / bucketValue).toInt
  accumulateDigits(placeHolder - 1, acc - (bucketValue * digitIndex),
       appendSafeToList(builder, digitIndex))
}
```
</sub>
</td>

<td >
<sub>

```golang

func ConvertToEncodedString(id int64) string {
	var builder strings.Builder
	placeHolder := FindStartBucket(id)
	var bucketValue int64
	acc  := id
	var digitIndex int

	results := AccumulateDigits(Acc{int64(placeHolder), id, &builder})

	bucketValue = PowDigitsBase(1)
	digitIndex = int(results.acc / bucketValue)
	acc = results.acc - (bucketValue * int64(digitIndex))
	AppendSafe(&builder, digitIndex)

	place1DigitIndex := int(acc % bucketValue)
	AppendSafe(&builder, place1DigitIndex)
	return builder.String()
}

type Acc struct {
	placeHolder int64
	acc int64
	builder *strings.Builder
}

func  AccumulateDigits(args Acc) Acc {
	if !(args.placeHolder > 1) {
		return args
	}
	bucketValue := PowDigitsBase(args.placeHolder)
	digitIndex := int(args.acc / bucketValue)
	return AccumulateDigits( Acc{args.placeHolder - 1,
		args.acc - (bucketValue * int64(digitIndex)),
		AppendSafe(args.builder, digitIndex)})
}

```

</sub>
</td>

</tr>
</table>





### convertToLong

<table>
<tr>
<th>
Scala
</th>
<th>
Go Lang
</th>
</tr>

<tr>
<td>
<sub>

```scala

def convertToLong(strId: String): Long =
                            doConvertToLong(strId.toCharArray)

private def doConvertToLong(chars: Array[Char]): Long = {
  val (acc, _) = chars.reverse.foldLeft(0L, 0) { (pos, ch) =>
    val (acc, position) = pos
    val value = computeValue(ch, position)
    (acc + value, position + 1)
  }
  acc
}
```
</sub>
</td>

<td >
<sub>

```golang

func ConvertToLong( strId string) int64 {
	return DoConvertCharsToLong([]rune(strId))
}

type Pos struct {
	acc int64
	position int
}

func reduce(initValue Pos, s []rune, f func(rune, Pos) Pos) Pos {
	acc := initValue
	for _, v := range s {
		acc = f(v, acc)
	}
	return acc
}

func DoConvertCharsToLong(chars []rune) int64 {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	pos := reduce(Pos{0, 0}, chars, func(r rune, pos Pos) Pos {
		value := ComputeValue(r, pos.position)
		return Pos{pos.acc + value, pos.position + 1 }
	})
	acc:= pos.acc
	return acc
}
```
</sub>
</td>

</tr>
</table>


----


## Go Lang Functional full example

```golang

package main

import (
	"fmt"
	"hash/fnv"
	"strings"
)

var DIGITS =  [...]rune{
	//0    1    2    3    4    5    6    7    8    9
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	//10    11   12   13  14   15   16   17    18    19   20  21
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
	//22    23   24   25  26   27   28   29    30    31   32  33    34  35
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	// 36  37  38   39   40   41    42    43   44  45   46    47
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',  //Easy to add more characters if not using lookup tables.
	// 48  49   50   51   52   53   54   55  56   57   58  59   60   61   // 62   63, 64
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Printf("%s %d\n", ConvertToEncodedString(12345678910), ConvertToLong("DTVD3O"))
	longUrl := "https://www.somewebiste.com/dp/0201616165/?_encoding=UTF8&pd_rd_w=vwEcs&content-id=amzn1.sym.8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_p=8cf3b8ef-6a74-45dc-9f0d-6409eb523603&pf_rd_r=BQ0KD40K57XG761DBNBA&pd_rd_wg=DtkHk&pd_rd_r=f94b60b7-9080-4065-b77f-6377ec854d17&ref_=pd_gw_ci_mcx_mi"
	urlId := hash(longUrl)
	shortHandle := ConvertToEncodedString(int64(urlId))
	fmt.Printf("url id %d short handle %s %d \n", urlId, shortHandle, ConvertToLong(shortHandle))
}



func ConvertToEncodedString(id int64) string {
	var builder strings.Builder
	placeHolder := FindStartBucket(id)
	var bucketValue int64
	acc  := id
	var digitIndex int

	results := AccumulateDigits(Acc{int64(placeHolder), id, &builder})

	bucketValue = PowDigitsBase(1)
	digitIndex = int(results.acc / bucketValue)
	acc = results.acc - (bucketValue * int64(digitIndex))
	AppendSafe(&builder, digitIndex)

	place1DigitIndex := int(acc % bucketValue)
	AppendSafe(&builder, place1DigitIndex)
	return builder.String()
}

type Acc struct {
	placeHolder int64
	acc int64
	builder *strings.Builder
}

func  AccumulateDigits(args Acc) Acc {
	if !(args.placeHolder > 1) {
		return args
	}
	bucketValue := PowDigitsBase(args.placeHolder)
	digitIndex := int(args.acc / bucketValue)
	return AccumulateDigits( Acc{args.placeHolder - 1,
		args.acc - (bucketValue * int64(digitIndex)),
		AppendSafe(args.builder, digitIndex)})
}


func FindIndexOfDigitInTable(c rune) int {
	index := -1
	for i := 0; i < len(DIGITS); i++ {
		if DIGITS[i] == c {
			index = i
			break
		}
	}
	return index
}

func FindStartBucket(value int64) int{
	for i := 0; i < 15; i++ {
		if value < PowDigitsBase(int64(i)) {
			return i-1
		}
	}
	return 0
}

func AppendSafe( builder *strings.Builder,  digitIndex int) *strings.Builder{
	char := DIGITS[digitIndex]

	if digitIndex != 0 {
		builder.WriteRune(char)
	} else {
		if builder.Len() > 0 {
			builder.WriteRune(char)
		}
	}
	return builder
}

func ConvertToLong( strId string) int64 {
	return DoConvertCharsToLong([]rune(strId))
}

// see https://github.com/robpike/filter

type Pos struct {
	acc int64
	position int
}

func reduce(initValue Pos, s []rune, f func(rune, Pos) Pos) Pos {
	acc := initValue
	for _, v := range s {
		acc = f(v, acc)
	}
	return acc
}

func DoConvertCharsToLong(chars []rune) int64 {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	pos := reduce(Pos{0, 0}, chars, func(r rune, pos Pos) Pos {
		value := ComputeValue(r, pos.position)
		return Pos{pos.acc + value, pos.position + 1 }
	})
	acc:= pos.acc
	return acc
}

func ComputeValue(c rune, position int) int64 {
	digitIndex := FindIndexOfDigitInTable(c)
	multiplier  := PowDigitsBase(int64(position))
	return int64(digitIndex) * multiplier
}

func PowDigitsBase( exponent int64) int64 {
	return DoPow(exponent, int64(len(DIGITS)))
}

func DoPow( exponent int64, base int64) int64 {
	if exponent == 0 { return 1 }
	return DoPow(exponent - 1, base) * base
}

```
