package strsub

import (
	"testing"
	"fmt"
)


func TestConvAvoidChar(t *testing.T) {

	type Input struct {
		Str string
		Expect string
	}

	tests := []Input{
		Input{"やまだたろう","やまだたろう"},
		Input{"やまダ太郎","やまダ太郎"},
		Input{"ﾔﾏﾀﾞ､ﾀﾛｳ","ヤマダ、タロウ"},
		Input{"ｶﾞﾗﾊﾟｺﾞｽｹｰﾀｲ","ガラパゴスケータイ"},
		Input{"､｡;,'","、。；，’"},
	}

	for _,tt := range tests {

		result := ConvAvoidChar(tt.Str)

		if result != tt.Expect {
			t.Error(fmt.Sprintf("Error 実際：%s  理想：%s \n",result,tt.Expect))
		} //else {
			// t.Log("",result)
			// t.Log(result)

		//}
	}
}
