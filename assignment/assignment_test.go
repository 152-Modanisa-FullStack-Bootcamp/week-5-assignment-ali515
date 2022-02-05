package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Obj interface{}

//Test AddUint32
type TestAddUint32Model struct {
	input1, input2, result1 uint32
	result2                 bool
}

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	cases := [7]TestAddUint32Model{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, true},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
	}

	for i := 0; i < len(cases); i++ {
		_case := cases[i]
		sum, overflow := AddUint32(_case.input1, _case.input2)
		msg := []Obj{_case, sum, overflow}
		assert.Equal(t, _case.result1, sum, msg)
		assert.Equal(t, _case.result2, overflow, msg)
	}
}

//Test CeilNumber
type TestCeilNumberModel struct {
	input, result float64
}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/

	cases := []TestCeilNumberModel{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for i := 0; i < len(cases); i++ {
		_case := cases[i]

		point := CeilNumber(_case.input)
		msg := []Obj{_case, point}
		assert.Equal(t, _case.result, point, msg)
	}
}

//Test AlphabetSoup
type TestAlphabetSoupModel struct {
	input, result string
}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/
	cases := []TestAlphabetSoupModel{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for i := 0; i < len(cases); i++ {
		_case := cases[i]
		result := AlphabetSoup(_case.input)
		msg := []Obj{_case, result}
		assert.Equal(t, _case.result, result, msg)
	}
}

//Test StringMask
type TestStringMaskModel struct {
	input1 string
	input2 uint
	result string
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/
	cases := []TestStringMaskModel{
		{"!mysecret*", 2, "!m********"},
		{"", 2, "*"},
		{"", 3, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"string", 8, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for i := 0; i < len(cases); i++ {
		_case := cases[i]
		result := StringMask(_case.input1, _case.input2)
		msg := []Obj{_case, result}
		assert.Equal(t, _case.result, result, msg)
	}
}

//Test WordSplit
type TestWordSplitModel struct {
	input  string
	result string
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	cases := []TestWordSplitModel{
		{"hellocat", "hello,cat"},
		{"catbat", "cat,bat"},
		{"yellowapple", "yellow,apple"},
		{"", "not possible"},
		{"notcat", "not possible"},
		{"bootcamprocks!", "not possible"},
	}
	for i := 0; i < len(cases); i++ {
		_case := cases[i]
		result := WordSplit([2]string{_case.input, words})
		msg := []Obj{_case, result}
		assert.Equal(t, _case.result, result, msg)
	}
}

//Test VariadicSet
func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks!" => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/

	set := VariadicSet(4, 2, 5, 4, 2, 4)
	assert.Equal(t, []interface{}{4, 2, 5}, set)

	set = VariadicSet("bootcamp", "rocks!", "really", "rocks!")
	assert.Equal(t, []interface{}{"bootcamp", "rocks!", "really"}, set)

	set = VariadicSet(1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first")
	assert.Equal(t, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}, set)
}
