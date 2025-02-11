package main

import (
	"fmt"
	"math"
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
		return fmt.Errorf("invalid index")
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
			return -1, fmt.Errorf("invalid IP address")
		} else {
			if c&1 == 1 {
				number += math.Pow(2, float64(len(str)-deg-1))
			}
		}
	}
	return int(number), nil
}

func ToDecimal(str string) (int, error) {
	if len(str) > 1 && str[0] == '0' {
		return -1, fmt.Errorf("invalid IP address")
	} else if len(str) == 0 {
		return -1, fmt.Errorf("invalid IP address")
	}

	number := 0
	for index, symbol := range str {
		if symbol < '0' || symbol > '9' {
			return -1, fmt.Errorf("invalid IP address")
		}
		number += int(symbol-'0') * int(math.Pow(10, float64(len(str)-index-1)))
	}
	return number, nil
}

func ParseIPv4Adress(str string) (IPV4Adress, error) {
	ip := IPV4Adress{}
	if len(str) == 35 && str[8] == '.' && str[17] == '.' && str[26] == '.' {
		for index, component := range strings.Split(str, ".") {
			if num, err := BinaryToDecimal(component); err == nil {
				ip.Set(index, num)
			} else {
				return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
			}
		}
	} else if len(str) > 6 && len(str) < 16 {
		// counter := 0
		// index := 0
		// components := make([]strings.Builder, 4)
		// for _, symbol := range str {
		// 	if symbol == '.' && counter != 0 {
		// 		if num, err := strconv.Atoi(components[index].String()); err == nil && num >= 0 && num <= 255 {
		// 			if len(components[index].String()) > 1 && components[index].String()[0] == '0' {
		// 				return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// 			}
		// 			ip.Set(index, num)
		// 			index++
		// 			counter = 0
		// 		} else {
		// 			return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// 		}
		// 	} else {
		// 		counter++
		// 		if counter > 3 {
		// 			return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// 		}
		// 		components[index].WriteRune(symbol)
		// 	}
		// }
		// if counter < 4 {
		// 	if num, err := strconv.Atoi(components[index].String()); err == nil {
		// 		ip.Set(index, num)
		// 		index++
		// 	} else {
		// 		return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// 	}
		// } else {
		// 	return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// }

		// if index != 4 {
		// 	return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		// }
		components := strings.Split(str, ".")
		if len(components) != 4 {
			return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
		}
		for index, component := range components {
			if num, err := ToDecimal(component); err == nil && num >= 0 && num <= 255 {
				ip.Set(index, num)
			} else {
				return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
			}
		}
	} else {
		return IPV4Adress{0, 0, 0, 0}, fmt.Errorf("invalid IP address")
	}
	return ip, nil
}
