package noppe

import (
	// "fmt"
	// "strings"
	// "regexp"
)


var fullkanas = []rune{'。','「','」','、','・','ヲ','ァ','ィ','ゥ','ェ','ォ','ャ','ュ','ョ','ッ','ー',
					   'ア','イ','ウ','エ','オ','カ','ガ','キ','ギ','ク','グ','ケ','ゲ','コ','ゴ','サ',
					   'ザ','シ','ジ','ス','ズ','セ','ゼ','ソ','ゾ','タ','ダ','チ','ヂ','ツ','ヅ','テ',
					   'デ','ト','ド','ナ','ニ','ヌ','ネ','ノ','ハ','バ','パ','ヒ','ビ','ピ','フ','ブ',
					   'プ','ヘ','ベ','ペ','ホ','ボ','ポ','マ','ミ','ム','メ','モ','ヤ','ユ','ヨ','ラ',
					   'リ','ル','レ','ロ','ワ','ン','ヴ','″','°',
					}

var halfkanas =  []rune{'ｧ','ｱ','ｨ','ｲ','ｩ','ｳ','ｪ','ｴ','ｫ','ｵ',
						'ｶ','ｶ','ｷ','ｷ','ｸ','ｸ','ｹ','ｹ','ｺ','ｺ',
						'ｻ','ｻ','ｼ','ｼ','ｽ','ｽ','ｾ','ｾ','ｿ','ｿ',
						'ﾀ','ﾀ','ﾁ','ﾁ','ｯ','ﾂ','ﾂ','ﾃ','ﾃ','ﾄ',
						'ﾄ','ﾅ','ﾆ','ﾇ','ﾈ','ﾉ','ﾊ','ﾊ','ﾊ','ﾋ',
						'ﾋ','ﾋ','ﾌ','ﾌ','ﾌ','ﾍ','ﾍ','ﾍ','ﾎ','ﾎ',
						'ﾎ','ﾏ','ﾐ','ﾑ','ﾒ','ﾓ','ｬ','ﾔ','ｭ','ﾕ',
						'ｮ','ﾖ','ﾗ','ﾘ','ﾙ','ﾚ','ﾛ','ﾜ','ﾜ','ｲ',
						'ｴ','ｦ','ﾝ','ｳ','ｶ','ｹ',' ','ｰ','､','｡',
						'｢','｣','･',
					}






func ConvAvoidChar(input string) string {

	var out []rune

	xx := []rune(input)

	xl := len(xx)
	cur := 0

	// for i,x := range xx {
		for {
		
		if cur >= xl {
			break
		}

		x := xx[cur]
		c := x

		switch  {
			case x == 0x27:
				c = '’'
			case x == 0x2C:
				c = '，'
			case x == 0x3B:
				c = '；'
			case x == 0x3C:
				c = '＜'
			case x == 0x3E:
				c = '＞'
			case 0xFF06 < x && x < 0xFFA0 :

				var cc []rune

				cc = append(cc,x)

				if cur < xl-1 {
					cc = append(cc,xx[cur+1])
				}


				c = halfKanaToWide(cc,&cur)
		}

		// fmt.Printf("%d %c\n",cur,c)

		out = append(out,c)

		cur++
	}

	return string(out)
}


func halfKanaToWide(c2c []rune,cur *int) rune {

	shift := 0
	var c2 int

	c1 := int(c2c[0])

	if len(c2c) > 1 {
		c2 = int(c2c[1])
	}


	//濁点
	if c1 == 0xFF9E {
		return fullkanas[87]
	}

	//半濁点
	if c1 == 0xFF9F {
		return fullkanas[88]
	}

	if 0xFF61 <=c1 && c1 <= 0xFF75{
		//ア～オ 始点ア：0xFF61

		if c1== 0xFF73 && c2==0xFF9E {
			*cur += 1
			return fullkanas[86]
		} else {
			return fullkanas[c1-0xFF61]
		}

	} else if 0xFF76 <= c1 && c1 <=0xFF84 {
		//カ～ト 始点カ：0xFF76
		// shift=0xFF76-0xFF61
		shift=21

		cpos := 2*(c1-0xFF76) + shift

		if c2== 0xFF9E {
			*cur += 1
			cpos ++
		}

		return fullkanas[cpos]

	} else if 0xFF85 <= c1 && c1 <=0xFF89 {
		//ナ～ノ 始点ナ：0xFF85
		// shift=0xFF76-0xFF61+2*(0xFF85-0xFF76)

		shift=51

		return fullkanas[c1 - 0xFF85 + shift]

	} else if 0xFF8A <=c1 && c1<=0xFF8E {
		//ハ～ホ 始点ハ：0xFF8A
		// shift=0xFF76-0xFF61+2*(0xFF85-0xFF76)+0xFF8A-0xFF85

		shift=56

		cpos := 3*(c1-0xFF8A) + shift

		if c2==0xFF9E {
			*cur += 1
			cpos += 1
		} else if c2==0xFF9F {
			*cur += 1
			cpos += 2
		}

		return fullkanas[cpos]

	} else {
		//マ～ン 始点マ：0xFF8F
		// shift=0xFF76-0xFF61+2*(0xFF85-0xFF76)+0xFF8A-0xFF85+3*(0xFF8F-0xFF8A)
		shift=71

		return fullkanas[c1 - 0xFF8F + shift]
	}



}