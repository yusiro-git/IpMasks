package IpMasks

import (
	"fmt"
	"testing"
)

type network_metadata_with_error struct {
	metadata network_metadata
	err      error
}

type testt struct {
	want     network_metadata_with_error
	ip       string
	networks int
	hosts    int
}

func TestGetNetworkInfo(t *testing.T) {
	var testts = []testt{
		{
			network_metadata_with_error{
				network_metadata{
					ip:            0b00000001_00000001_00000001_00000001,
					mask:          0b11111111_11111111_11111111_11111111,
					broadcast:     0b00000001_00000001_00000001_00000001,
					class:         "A",
					ipQuantity:    1,
					hostsQuantity: -1,
					top5:          []IPv4{0b00000001_00000001_00000001_00000001},
					low5:          []IPv4{0b00000001_00000001_00000001_00000001},
				},
				nil,
			},
			"1.1.1.1",
			1,
			1,
		},
		{
			network_metadata_with_error{
				network_metadata{
					ip:            0b11000000_10101000_00000000_00000001,
					mask:          0b11111111_11111111_11111111_00000000,
					broadcast:     0b11000000_10101000_00000000_11111111,
					class:         "C",
					ipQuantity:    255,
					hostsQuantity: 253,
					top5:          []IPv4{0b11000000_10101000_00000000_00000001, 0b11000000_10101000_00000000_00000010, 0b11000000_10101000_00000000_00000011, 0b11000000_10101000_00000000_00000100, 0b11000000_10101000_00000000_00000101},
					low5:          []IPv4{0b11000000_10101000_00000000_11111111, 0b11000000_10101000_00000000_11111110, 0b11000000_10101000_00000000_11111101, 0b11000000_10101000_00000000_11111100, 0b11000000_10101000_00000000_11111011},
				},
				nil,
			},
			"192.168.0.1",
			1,
			254,
		},
		{
			network_metadata_with_error{
				network_metadata{},
				fmt.Errorf("invalid input"),
			},
			"192.168.0.1",
			-1,
			254,
		},
		{
			network_metadata_with_error{
				network_metadata{},
				fmt.Errorf("invalid input"),
			},
			"192.168.0.1",
			1,
			-1,
		},
		{
			network_metadata_with_error{
				network_metadata{},
				fmt.Errorf("invalid input"),
			},
			"192.168.0.1",
			1,
			4294967296,
		},
	}
	for _, testt := range testts {
		t.Run(testt.ip, func(t *testing.T) {
			ip, _ := ParseIPv4Adress(testt.ip)
			got, err := GetNetworkInfo(ip, testt.networks, testt.hosts)
			want := testt.want.metadata

		})
	}
}
