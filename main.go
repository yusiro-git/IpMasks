package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type IPV4Adress struct {
	FirstOctet  int
	SecondOctet int
	ThirdOctet  int
	FourthOctet int
}

func MakeIPV4Adress(firstOctet, secondOctet, thirdOctet, fourthOctet int) IPV4Adress {
	return IPV4Adress{
		FirstOctet:  firstOctet,
		SecondOctet: secondOctet,
		ThirdOctet:  thirdOctet,
		FourthOctet: fourthOctet,
	}
}

func (ip *IPV4Adress) Decimal() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.FirstOctet, ip.SecondOctet, ip.ThirdOctet, ip.FourthOctet)
}

func (ip *IPV4Adress) Binary() string {
	return fmt.Sprintf("%08b.%08b.%08b.%08b", ip.FirstOctet, ip.SecondOctet, ip.ThirdOctet, ip.FourthOctet)
}

func (ip *IPV4Adress) Set(index int, value int) error {
	if index < 0 || index > 3 {
		return fmt.Errorf("Invalid index: %d", index)
	}
	switch index {
	case 0:
		ip.FirstOctet = value
	case 1:
		ip.SecondOctet = value
	case 2:
		ip.ThirdOctet = value
	case 3:
		ip.FourthOctet = value
	}
	return nil
}

func BinaryToDecimal(str string) (int, error) {
	number := 0.0
	for deg, c := range str {
		if c != '0' && c != '1' {
			return -1, fmt.Errorf("Строка '%s' не является представлением числа в двоичной системе счисления.", str)
		} else {
			if c&1 == 1 {
				number += math.Pow(2, float64(len(str)-deg-1))
			}
		}
	}
	return int(number), nil
}

func ParseIPv4Adress(str string) (error, IPV4Adress) {
	ip := IPV4Adress{}
	if len(str) == 35 && str[8] == '.' && str[17] == '.' && str[26] == '.' {
		for index, component := range strings.Split(str, ".") {
			if num, err := BinaryToDecimal(component); err == nil {
				ip.Set(index, num)
			} else {
				return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
			}
		}
	} else if len(str) > 6 && len(str) < 16 {
		counter := 0
		index := 0
		components := make([]strings.Builder, 4)
		for _, symbol := range str {
			if symbol == '.' && counter != 0 {
				if num, err := strconv.Atoi(components[index].String()); err == nil {
					ip.Set(index, num)
					index++
					counter = 0
				} else {
					return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
				}
			} else {
				counter++
				if counter > 3 {
					return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
				}
				components[index].WriteRune(symbol)
			}
		}
		if counter < 4 {
			if num, err := strconv.Atoi(components[index].String()); err == nil {
				ip.Set(index, num)
			} else {
				return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
			}
		} else {
			return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
		}
	} else {
		return fmt.Errorf("Invalid IP adress: %s", str), IPV4Adress{0, 0, 0, 0}
	}
	return nil, ip
}

func main() {
	if err, a := ParseIPv4Adress("11000000.10101000.00000001.00000001"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a.Decimal(), a.Binary())
	}
	if err, b := ParseIPv4Adress("1.145.255.1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b.Decimal(), b.Binary())
	}

}
