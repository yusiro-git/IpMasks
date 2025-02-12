package IpMasks

import (
	"fmt"
	"testing"
)

type result interface{}

type IPv4_with_error struct {
	ip  []IPv4
	err error
}

type class_with_error struct {
	class string
	err   error
}

type test struct {
	want result
	data string
}

func TestParseIPv4Adress(t *testing.T) {
	var tests = []test{
		{IPv4_with_error{[]IPv4{0b00000001_00000001_00000001_00000001}, nil}, "1.1.1.1"},
		{IPv4_with_error{[]IPv4{0b00001010_01100100_11001000_00000000}, nil}, "10.100.200.0"},
		{IPv4_with_error{[]IPv4{0b11101010_11111111_11111111_11111111}, nil}, "234.255.255.255"},
		{IPv4_with_error{[]IPv4{0b00000000_00000000_00000000_00000000}, nil}, "0.0.0.0"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11111111_11111111}, nil}, "255.255.255.255"},
		{IPv4_with_error{[]IPv4{0b00001001_00001010_01101111_00001101}, nil}, "9.10.111.13"},
		{IPv4_with_error{[]IPv4{0b11000000_00100010_00001000_10110111}, nil}, "192.34.8.183"},

		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "237.123.983.1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "87313.38.3.0"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "982.2.23.256"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "23.-1.87.111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "34.98.123.-123"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "45.123.2147483649.32"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.72.4294967298.1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "-1284761.-4294967298.1.1"},

		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "1.1.1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "1.1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "243.123.255"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.123"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123"},

		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "dsj.asd.qe.q"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.13.df.1"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.13.1.df"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.13.1qe.1d"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123sdf.13.1.1d"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "123.13yg"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "asdouasdgai"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "aiudaduiasgdiuagduasgduasdgaiudgsudgasudasudg"},

		{IPv4_with_error{[]IPv4{0b00000000_00000000_00000000_00000000}, nil}, "00000000.00000000.00000000.00000000"},
		{IPv4_with_error{[]IPv4{0b00000001_00000001_00000001_00000001}, nil}, "00000001.00000001.00000001.00000001"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11111111_11111111}, nil}, "11111111.11111111.11111111.11111111"},
		{IPv4_with_error{[]IPv4{0b01100010_00001100_00101101_00000001}, nil}, "01100010.00001100.00101101.00000001"},
		{IPv4_with_error{[]IPv4{0b11001000_00000000_00100010_01010111}, nil}, "11001000.00000000.00100010.01010111"},

		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00000000.00000000.00000000.00000000.00000000"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00000000.00000000.00000000.00000000.00000000.00000000"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00011.01111.00000000.00001111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "000011110000.000011110000.11110000"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "000011110000.000011110000.11110000.00001111.00001111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "010.010.010.010"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.00002222.00001111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.00009111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.123"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.4562.00001111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.05555511.00001111"},
		{IPv4_with_error{[]IPv4{0}, fmt.Errorf("invalid IP address")}, "00001111.00001111.aaaaa111.000011l1"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			got, err := ParseIPv4Adress(test.data)
			want := test.want.(IPv4_with_error)
			if got != want.ip[0] || (err != nil && err.Error() != want.err.Error()) {
				t.Errorf("got %v, %v; want %v, %v", got, err, want.ip, want.err)
			}
		})
	}
}

func TestDetectClass(t *testing.T) {
	var tests = []test{
		{class_with_error{"A", nil}, "0.0.0.0"},
		{class_with_error{"A", nil}, "98.91.255.76"},
		{class_with_error{"A", nil}, "127.98.12.1"},
		{class_with_error{"A", nil}, "0.35.123.255"},
		{class_with_error{"A", nil}, "127.255.255.255"},
		{class_with_error{"B", nil}, "128.0.0.0"},
		{class_with_error{"B", nil}, "191.255.255.255"},
		{class_with_error{"B", nil}, "141.45.182.99"},
		{class_with_error{"B", nil}, "165.255.255.89"},
		{class_with_error{"C", nil}, "192.0.0.0"},
		{class_with_error{"C", nil}, "223.255.255.255"},
		{class_with_error{"C", nil}, "200.200.200.200"},
		{class_with_error{"C", nil}, "204.255.0.45"},
		{class_with_error{"D", nil}, "224.0.0.0"},
		{class_with_error{"D", nil}, "239.255.255.255"},
		{class_with_error{"D", nil}, "230.56.98.251"},
		{class_with_error{"D", nil}, "225.255.45.1"},
		{class_with_error{"E", nil}, "240.0.0.0"},
		{class_with_error{"E", nil}, "255.255.255.255"},
		{class_with_error{"E", nil}, "245.56.123.245"},
		{class_with_error{"E", nil}, "250.255.98.234"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			got, err := DetectClass(func() IPv4 { ptr, _ := ParseIPv4Adress(test.data); return ptr }())
			want := test.want.(class_with_error)
			if err != nil || got != want.class {
				t.Errorf("got %v, %v; want %v, %v", got, err, want.class, want.err)
			}
		})
	}
}

func TestGetBordersClass(t *testing.T) {
	var tests = []test{
		{IPv4_with_error{[]IPv4{0b00000000_00000000_00000000_00000000, 0b01111111_11111111_11111111_11111111}, nil}, "A"},
		{IPv4_with_error{[]IPv4{0b10000000_00000000_00000000_00000000, 0b10111111_11111111_11111111_11111111}, nil}, "B"},
		{IPv4_with_error{[]IPv4{0b11000000_00000000_00000000_00000000, 0b11011111_11111111_11111111_11111111}, nil}, "C"},
		{IPv4_with_error{[]IPv4{0b11100000_00000000_00000000_00000000, 0b11101111_11111111_11111111_11111111}, nil}, "D"},
		{IPv4_with_error{[]IPv4{0b11110000_00000000_00000000_00000000, 0b11111111_11111111_11111111_11111111}, nil}, "E"},
		{IPv4_with_error{[]IPv4{0b0, 0b0}, fmt.Errorf("invalid class")}, "X"},
		{IPv4_with_error{[]IPv4{0b0, 0b0}, fmt.Errorf("invalid class")}, "F"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			min, max, err := GetBordersClass(test.data)
			want := test.want.(IPv4_with_error)
			if (err == nil && want.err != nil) || (err != nil && want.err == nil) || (err != nil && err.Error() != want.err.Error()) || min != want.ip[0] || max != want.ip[1] {
				t.Errorf("got '%v', '%v', '%v'; want '%v', '%v', '%v'", min, max, err, want.ip[0], want.ip[1], want.err)
			}
		})
	}
}

func TestSetMask(t *testing.T) {
	var tests = []test{
		{IPv4_with_error{[]IPv4{0b11111111_11000000_00000000_00000000}, nil}, "255.192.0.0"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11111110_00000000}, nil}, "251.123.34.0"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11111111_11111110}, nil}, "204.0.0.2"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11110000_00000000}, nil}, "33.8.16.0"},
		{IPv4_with_error{[]IPv4{0b11111111_11111111_11111111_11110000}, nil}, "243.0.14.16"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			got := SetMask(func() IPv4 { ptr, _ := ParseIPv4Adress(test.data); return ptr }())
			want := test.want.(IPv4_with_error)
			if got != want.ip[0] {
				t.Errorf("got %v; want %v", got, want.ip[0])
			}
		})
	}
}
