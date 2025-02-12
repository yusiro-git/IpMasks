package IpMasks

import (
	"fmt"
	"math"
	"strings"
)

type IPv4 uint32

func (ip IPv4) Decimal() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip>>16&0b11111111, ip>>8&0b11111111, ip&0b11111111)
}

func (ip IPv4) Binary() string {
	return fmt.Sprintf("%08b.%08b.%08b.%08b", ip>>24, ip>>16&0b11111111, ip>>8&0b11111111, ip&0b11111111)
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

func DetectClass(ip IPv4) (string, error) {
	if ip&0b11110000_00000000_00000000_00000000 == 0b11110000_00000000_00000000_00000000 {
		return "E", nil
	} else if ip&0b11100000_00000000_00000000_00000000 == 0b11100000_00000000_00000000_00000000 {
		return "D", nil
	} else if ip&0b11000000_00000000_00000000_00000000 == 0b11000000_00000000_00000000_00000000 {
		return "C", nil
	} else if ip&0b10000000_00000000_00000000_00000000 == 0b10000000_00000000_00000000_00000000 {
		return "B", nil
	} else if ip&0b10000000_00000000_00000000_00000000 == 0 {
		return "A", nil
	} else {
		return "X", fmt.Errorf("invalid IP address")
	}
}

func GetBordersClass(class string) (IPv4, IPv4, error) {
	switch class {
	case "A": //"127.255.255.255"
		return IPv4(0), IPv4(0b01111111_11111111_11111111_11111111), nil
	case "B":
		return IPv4(0b10000000_00000000_00000000_00000000), IPv4(0b10111111_11111111_11111111_11111111), nil
	case "C":
		return IPv4(0b11000000_00000000_00000000_00000000), IPv4(0b11011111_11111111_11111111_11111111), nil
	case "D":
		return IPv4(0b11100000_00000000_00000000_00000000), IPv4(0b11101111_11111111_11111111_11111111), nil
	case "E":
		return IPv4(0b11110000_00000000_00000000_00000000), IPv4(0b11111111_11111111_11111111_11111111), nil
	default:
		return IPv4(0), IPv4(0), fmt.Errorf("invalid class")
	}
}

func ParseIPv4Adress(str string) (IPv4, error) {
	var ip IPv4 = 0
	if len(str) == 35 && str[8] == '.' && str[17] == '.' && str[26] == '.' {
		if parts := strings.Split(str, "."); len(parts) == 4 {
			for index, component := range parts {
				if num, err := BinaryToDecimal(component); err == nil {
					ip = ip | IPv4(num)<<(24-8*index)
				} else {
					return 0b0, fmt.Errorf("invalid IP address")
				}
			}
		}

	} else if len(str) > 6 && len(str) < 16 {
		components := strings.Split(str, ".")
		if len(components) != 4 {
			return 0b0, fmt.Errorf("invalid IP address")
		}
		for index, component := range components {
			if num, err := ToDecimal(component); err == nil && num >= 0 && num <= 255 {
				ip = ip | IPv4(num)<<(24-8*index)
			} else {
				return 0b0, fmt.Errorf("invalid IP address")
			}
		}
	} else {
		return 0b0, fmt.Errorf("invalid IP address")
	}
	return ip, nil
}

func SetMask(ip IPv4) IPv4 {
	mask := IPv4(0b11111111_11111111_11111111_11111111)
	for i := 0; i < 32; i++ {
		if ip&mask != ip {
			return 0b10000000_00000000_00000000_00000000 | (mask >> 1)
		}
		mask = mask << 1
	}
	return mask
}

type network_metadata struct {
	ip            IPv4   // done
	mask          IPv4   // done
	broadcast     IPv4   // done
	class         string // done
	ipQuantity    int    // done
	hostsQuantity int    // done
	top5          []IPv4 //         TO DO Т^Т
	low5          []IPv4
}

func GetNetworkInfo(ip IPv4, networks int, hosts int) (network_metadata, error) {
	var this_network network_metadata
	this_network.class, _ = DetectClass(ip)
	this_network.ip = ip
	func() {
		if networks < 0 || hosts < 0 {
			panic("invalid input")
		}
		if int64(networks)+int64(hosts)+int64(SetMask(ip)) > 0b11111111_11111111_11111111_11111111 {
			panic("invalid input")
		}
	}()

	this_network.mask = AddBits(SetMask(ip), numberOfBits(networks))
	this_network.broadcast = ip | ^this_network.mask
	this_network.ipQuantity = int(^SetMask(ip))
	this_network.hostsQuantity = this_network.ipQuantity - 2
	var num int
	if this_network.hostsQuantity*this_network.ipQuantity < 5 {
		num = this_network.hostsQuantity * this_network.ipQuantity
	} else {
		num = 5
	}
	for i := range num {
		this_network.top5 = append(this_network.top5, this_network.ip+IPv4(i))
		this_network.low5 = append(this_network.low5, this_network.broadcast-IPv4(i))
	}

	return this_network, nil
}
