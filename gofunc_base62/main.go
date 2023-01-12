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
