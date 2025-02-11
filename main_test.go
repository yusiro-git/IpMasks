package main

import (
	"fmt"
	"testing"
)

type pair struct {
	ip  string
	err error
}

type test struct {
	want pair
	data string
}

func TestParseIPv4Adress(t *testing.T) {
	var tests = []test{
		{pair{"1.1.1.1", nil}, "1.1.1.1"},
		{pair{"10.100.200.0", nil}, "10.100.200.0"},
		{pair{"234.255.255.255", nil}, "234.255.255.255"},
		{pair{"0.0.0.0", nil}, "0.0.0.0"},
		{pair{"255.255.255.255", nil}, "255.255.255.255"},
		{pair{"70.230.2.7", fmt.Errorf("invalid IP address")}, "70.230.2.7"},
		{pair{"9.10.111.13", nil}, "9.10.111.13"},
		{pair{"192.34.8.183", nil}, "192.34.8.183"},

		{pair{"", fmt.Errorf("invalid IP address")}, "237.123.983.1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "87313.38.3.0"},
		{pair{"", fmt.Errorf("invalid IP address")}, "982.2.23.256"},
		{pair{"", fmt.Errorf("invalid IP address")}, "23.-1.87.111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "34.98.123.-123"},
		{pair{"", fmt.Errorf("invalid IP address")}, "45.123.2147483649.32"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.72.4294967298.1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "-1284761.-4294967298.1.1"},

		{pair{"", fmt.Errorf("invalid IP address")}, "1.1.1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "1.1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "243.123.255"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.123"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123"},

		{pair{"", fmt.Errorf("invalid IP address")}, "dsj.asd.qe.q"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.13.df.1"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.13.1.df"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.13.1qe.1d"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123sdf.13.1.1d"},
		{pair{"", fmt.Errorf("invalid IP address")}, "123.13yg"},
		{pair{"", fmt.Errorf("invalid IP address")}, "asdouasdgai"},
		{pair{"", fmt.Errorf("invalid IP address")}, "aiudaduiasgdiuagduasgduasdgaiudgsudgasudasudg"},

		{pair{"0.0.0.0", nil}, "00000000.00000000.00000000.00000000"},
		{pair{"1.1.1.1", nil}, "00000001.00000001.00000001.00000001"},
		{pair{"255.255.255.255", nil}, "11111111.11111111.11111111.11111111"},
		{pair{"98.12.45.1", nil}, "01100010.00001100.00101101.00000001"},
		{pair{"200.0.34.87", nil}, "11001000.00000000.00100010.01010111"},

		{pair{"", fmt.Errorf("invalid IP address")}, "00000000.00000000.00000000.00000000.00000000"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00000000.00000000.00000000.00000000.00000000.00000000"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00011.01111.00000000.00001111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "000011110000.000011110000.11110000"},
		{pair{"", fmt.Errorf("invalid IP address")}, "000011110000.000011110000.11110000.00001111.00001111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "010.010.010.010"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.00002222.00001111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.00009111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.123"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.00001111.4562.00001111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.05555511.00001111"},
		{pair{"", fmt.Errorf("invalid IP address")}, "00001111.00001111.aaaaa111.000011l1"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			ptr, err := ParseIPv4Adress(test.data)
			got := ""
			if err == nil {
				got = ptr.Decimal()
			}
			if got != test.want.ip || (err != nil && err.Error() != test.want.err.Error()) {
				t.Errorf("got %v, %v; want %v, %v", got, err, test.want.ip, test.want.err)
			}
		})
	}
}

func TestDetectClass(t *testing.T) {
	var tests = []test{
		{pair{"A", nil}, "0.0.0.0"},
		{pair{"A", nil}, "98.91.255.76"},
		{pair{"A", nil}, "127.98.12.1"},
		{pair{"A", nil}, "0.35.123.255"},
		{pair{"A", nil}, "127.255.255.255"},
		{pair{"B", nil}, "128.0.0.0"},
		{pair{"B", nil}, "191.255.255.255"},
		{pair{"B", nil}, "141.45.182.99"},
		{pair{"B", nil}, "165.255.255.89"},
		{pair{"C", nil}, "192.0.0.0"},
		{pair{"C", nil}, "223.255.255.255"},
		{pair{"C", nil}, "200.200.200.200"},
		{pair{"C", nil}, "204.255.0.45"},
		{pair{"D", nil}, "224.0.0.0"},
		{pair{"D", nil}, "239.255.255.255"},
		{pair{"D", nil}, "230.56.98.251"},
		{pair{"D", nil}, "225.255.45.1"},
		{pair{"E", nil}, "240.0.0.0"},
		{pair{"E", nil}, "255.255.255.255"},
		{pair{"E", nil}, "245.56.123.245"},
		{pair{"E", nil}, "250.255.98.234"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			ptr, _ := ParseIPv4Adress(test.data)
			got, err := DetectClass(ptr)
			if err != nil || got != test.want.ip {
				t.Errorf("got %v, %v; want %v, %v", got, err, test.want.ip, test.want.err)
			}
		})
	}
}

type tripple struct {
	min string
	max string
	err error
}

type tptr struct {
	want tripple
	data string
}

func TestGetBordersClass(t *testing.T) {
	var tests = []tptr{
		{tripple{"0.0.0.0", "127.255.255.255", nil}, "A"},
		{tripple{"128.0.0.0", "191.255.255.255", nil}, "B"},
		{tripple{"192.0.0.0", "223.255.255.255", nil}, "C"},
		{tripple{"224.0.0.0", "239.255.255.255", nil}, "D"},
		{tripple{"240.0.0.0", "255.255.255.255", nil}, "E"},
		{tripple{"", "", fmt.Errorf("invalid class")}, "X"},
		{tripple{"", "", fmt.Errorf("invalid class")}, "F"},
	}
	for _, test := range tests {
		t.Run(test.data, func(t *testing.T) {
			min, max, err := GetBordersClass(test.data)
			if (err == nil && test.want.err != nil) || (err != nil && test.want.err == nil) || (err != nil && err.Error() != test.want.err.Error()) || min != test.want.min || max != test.want.max {
				t.Errorf("got '%v', '%v', '%v'; want '%v', '%v', '%v'", min, max, err, test.want.min, test.want.max, test.want.err)
			}
		})
	}
}
