package noppe

import (
	"testing"
	"fmt"
)


func TestIsNotEmpty(t *testing.T) {

	var np *interface{}

	result := IsNotEmpty(np)

	if result {
		t.Error(fmt.Sprintf("Error null 実際：%v  理想：%v \n",result,false))
	}


	str := interface{}("")

	result = IsNotEmpty(&str)

	if result {
		t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",str.(string),result,false))
	}


	str = interface{}("teststring")

	result = IsNotEmpty(&str)

	if !result {
		t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",str.(string),result,true))
	}



	nbr := interface{}(0)

	result = IsNotEmpty(&nbr)

	if !result {
		t.Error(fmt.Sprintf("Error %d 実際：%v  理想：%v \n",nbr.(int),result,true))
	}


}




func TestIsAllowedChar(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"やまだたろう",true,true},
		Input{"aaa,",true,false},
		Input{"'bbb",true,false},
		Input{"ｶﾀｶﾅ",true,false},
		Input{"",false,true},
		Input{"ｶ高菜",false,false},
	}

	for _,tt := range tests {

		result := IsAllowedChar(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error 実際：%v  理想：%v \n",result,tt.Expect))
		} //else {
			// t.Log("",result)
			// t.Log(result)

		//}
	}


}

func TestIsFullKana(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"やまだたろう",true,false},
		Input{"カタカナ",true,true},
		Input{"aaa,",true,false},
		Input{"'bbb",true,false},
		Input{"ｶﾀｶﾅ",true,false},
	}

	for _,tt := range tests {

		result := IsFullKana(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}


func TestIsNbrTel(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"",true,false},
		Input{"0251234567",false,true},
		Input{"0251234567",true,true},
		Input{"025-123-4567",true,false},
		Input{"025-123-4567",false,false},
		Input{"08011223344",true,true},
		Input{"080-1122-3344",true,false},
		Input{"0255-12-3456",true,false},
		Input{"025-121-345",true,false},
		Input{"02-1211-3451",true,false},
		Input{"0244-121-34511",true,false},
		Input{"a24-1222-1111",true,false},
	}

	for _,tt := range tests {

		result := IsNbrTel(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}


func TestIsHyphenedTel(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"",true,false},
		Input{"0251234567",true,false},
		Input{"0251234567",false,false},
		Input{"025-123-4567",true,true},
		Input{"025-123-4567",false,true},
		Input{"08011223344",true,false},
		Input{"080-1122-3344",true,true},
		Input{"0255-12-3456",true,true},
		Input{"025-121-345",true,false},
		Input{"02-1211-3451",true,false},
		Input{"0244-121-34511",true,false},
		Input{"a24-1222-1111",true,false},
	}

	for _,tt := range tests {

		result := IsHyphenedTel(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}


func TestIsNbrZipcode(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"",true,false},
		Input{"9501122",true,true},
		Input{"9501122",false,true},
		Input{"950-1122",true,false},
		Input{"950-1122",false,false},
		Input{"95-11122",true,false},
		Input{"950a1122",true,false},
	}

	for _,tt := range tests {

		result := IsNbrZipcode(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}


func TestIsHyphenedZipcode(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"",true,false},
		Input{"9501122",true,false},
		Input{"9501122",false,false},
		Input{"950-1122",true,true},
		Input{"950-1122",false,true},
		Input{"95-11122",true,false},
		Input{"950a1122",true,false},
	}

	for _,tt := range tests {

		result := IsHyphenedZipcode(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}


func TestIsMailadr(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"taro.yamada123@example.com",true,true},
		Input{"taro.yamada123@example.co.jp",true,true},
		Input{"taro.yamada123.-test.@example.co.jp",true,true},
		Input{"taro.yamada123.@example.c",true,false},
		Input{"taro.yamada123.-example.c",true,false},
		Input{"taro.yamada123.@@example.c",true,false},
	}

	for _,tt := range tests {

		result := IsMailadr(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}

func TestIsNumeric(t *testing.T) {

	type Input struct {
		Str string
		Req bool
		Expect bool
	}

	tests := []Input{
		Input{"",false,true},
		Input{"123",true,true},
		Input{"123.45",true,true},
		Input{"0",true,true},
		Input{"0.5",true,true},
		Input{"09",true,false},
		Input{"09.",true,false},
		Input{"09.5",true,false},
		Input{"abc",true,false},
		Input{"0..5",true,false},
		Input{"0.5.5",true,false},
	}

	for _,tt := range tests {

		result := IsNumeric(tt.Str,tt.Req)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error %s 実際：%v  理想：%v \n",tt.Str,result,tt.Expect))
		} 
	}
}
