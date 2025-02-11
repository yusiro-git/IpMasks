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

func DetectClass(ip IPV4Adress) (string, error) {
	if ip.FirstOctet&0b11110000 == 0b11110000 {
		return "E", nil
	} else if ip.FirstOctet&0b11100000 == 0b11100000 {
		return "D", nil
	} else if ip.FirstOctet&0b11000000 == 0b11000000 {
		return "C", nil
	} else if ip.FirstOctet&0b10000000 == 0b10000000 {
		return "B", nil
	} else if ip.FirstOctet&0b10000000 == 0 {
		return "A", nil
	} else {
		return "X", fmt.Errorf("invalid IP address")
	}
}

func GetBordersClass(class string) (string, string, error) {
	switch class {
	case "A":
		return "0.0.0.0", "127.255.255.255", nil
	case "B":
		return "128.0.0.0", "191.255.255.255", nil
	case "C":
		return "192.0.0.0", "223.255.255.255", nil
	case "D":
		return "224.0.0.0", "239.255.255.255", nil
	case "E":
		return "240.0.0.0", "255.255.255.255", nil
	default:
		return "", "", fmt.Errorf("invalid class")
	}
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
