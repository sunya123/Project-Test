package main

import (
	"fmt"
	"time"
)

const (
	RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
	RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
	MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
	UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
)

const (
	surrogateMin = 0xD800
	surrogateMax = 0xDFFF
)
const (
	t1 = 0x00 // 0000 0000
	tx = 0x80 // 1000 0000
	t2 = 0xC0 // 1100 0000
	t3 = 0xE0 // 1110 0000
	t4 = 0xF0 // 1111 0000
	t5 = 0xF8 // 1111 1000

	maskx = 0x3F // 0011 1111
	mask2 = 0x1F // 0001 1111
	mask3 = 0x0F // 0000 1111
	mask4 = 0x07 // 0000 0111

	rune1Max = 1<<7 - 1
	rune2Max = 1<<11 - 1
	rune3Max = 1<<16 - 1
)


//func commons(s, t string) string {
//	var str []rune
//	for s != "" && t != "" {
//		// Extract first rune from each string.
//		var sr, tr rune
//		if s[0] < utf8.RuneSelf {
//			sr, s = rune(s[0]), s[1:]
//		} else {
//			r, size := utf8.DecodeRuneInString(s)
//			sr, s = r, s[size:]
//		}
//		if t[0] < utf8.RuneSelf {
//			tr, t = rune(t[0]), t[1:]
//		} else {
//			r, size := utf8.DecodeRuneInString(t)
//			tr, t = r, t[size:]
//		}
//
//		// If they match, keep going; if not, return false.
//
//		// Easy case.
//		if tr == sr {
//			str = append(str, sr)
//			continue
//		}
//
//		return string(str)
//	}
//
//	// One string is empty.  Are both?
//	return string(str)
//}

func common(s, t string) string {
	var str []rune
	var sb,tb [4]byte
	var rs,rt  rune
	ind:=0
	if s!=""&&t!=""{
		for len(s)-ind>0&&len(t)-ind>0{
			sb[0]=s[ind]
			tb[0]=t[ind]
			if sb[0] < tx&& tb[0]<tx && sb[0]==tb[0] {
				  str=append(str,rune(sb[0]))
				  ind++
				continue
			}
			// unexpected continuation byte?
			if sb[0] < t2||tb[0]< t2 {
				break
			}
			if len(s)-ind < 2 ||len(t)-ind<2 {
				break
			}
			sb[1] = s[ind+1]
			tb[1] = t[ind+1]
			if sb[1]< tx || t2 <= sb[1] ||tb[1]<tx||tb[1]<tx||t2<=tb[1]{
				break
			}

			// 2-byte, 11-bit sequence?
			if sb[0] < t3 && tb[0]<t3 {
				rs = rune(sb[0]&mask2)<<6 | rune(sb[1]&maskx)
				rt = rune(tb[0]&mask2)<<6 | rune(tb[1]&maskx)
				if rt <= rune1Max||rs <= rune1Max {
					break
				}
				if rs==rt{
					str=append(str,rs)
					ind+=2
					continue
				}
			}


			// need second continuation byte
			if len(s)-ind < 3|| len(t)-ind<3{
				break
			}
			sb[2]= s[ind+2]
			tb[2]= t[ind+2]
			if sb[2] < tx || t2 <= sb[2] ||tb[2] < tx || t2 <= tb[2]{
				break
			}

			// 3-byte, 16-bit sequence?
			if sb[0] < t4 && tb[0]<t4 {
				rs = rune(sb[0]&mask3)<<12 | rune(sb[1]&maskx)<<6 | rune(sb[2]&maskx)
				rt = rune(tb[0]&mask3)<<12 | rune(tb[1]&maskx)<<6 | rune(tb[2]&maskx)
				if rs <= rune2Max ||rs <= rune2Max {
					break
				}
				if surrogateMin <= rs && rs <= surrogateMax ||surrogateMin <= rt && rt <= surrogateMax {
					break
				}
				if rs==rt{
					str=append(str,rs)
					ind+=3
					continue
				}
			}


			// need third continuation byte
			if len(s)-ind < 4||len(t)-ind<4 {
				break
			}
			sb[3] = s[ind+3]
			tb[3] = t[ind+3]
			if sb[3] < tx || t2 <= sb[3] ||tb[3] < tx || t2 <= tb[3]{
				break
			}

			// 4-byte, 21-bit sequence?
			if sb[0] < t5&&tb[0]<t5 {
				rs = rune(sb[0]&mask4)<<18 | rune(sb[1]&maskx)<<12 | rune(sb[2]&maskx)<<6 | rune(sb[3]&maskx)
				rt = rune(tb[0]&mask4)<<18 | rune(tb[1]&maskx)<<12 | rune(tb[2]&maskx)<<6 | rune(tb[3]&maskx)
				if rs <= rune3Max || MaxRune < rs ||rt <= rune3Max || MaxRune < rt {
					break
				}
				if rs==rt{
					str=append(str,rs)
					ind+=4
					continue
				}

			}
		}


	}
	return string(str)
}

func main() {
	stri := "几个大盘那个好"
	stri2 := "几个大盘那个好代码规范咖啡店老"
	stri4 := "abcd"
	stri5 := "abcd7"
	str6 := "123"
	str7 := "123456"
	str8 := "abc哈124"
	str9 := "aBc哈124而899"
	str10 := "aBc哈124而89"
	t:=time.Now()
	for i:=0;i<100000000;i++{
		common(stri, stri2)
		//fmt.Println(str)
		common(stri4, stri5)
		//fmt.Println(st)
		 common(str6, str7)
		//fmt.Println(sts)
		common(str8, str9)
		//fmt.Println(sts2)
		common(str10, str9)
		//fmt.Println(common(str10, str9))
	}
	dt:=time.Since(t)
	fmt.Println(dt)

}
