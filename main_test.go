package main

import (
	"testing"
	"fmt"
)

type pair struct {
	ip string
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
		{pair{"70.230.2.7", fmt.Errorf("invalid IP address")}, ".70.230.2.7"},
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
			got, err := ParseIPv4Address(test.data)
			if got != test.want.ip || (err != nil && err.Error() != test.want.err.Error()) {
				t.Errorf("got %v, %v, want %v, %v", got, err, test.want.ip, test.want.err)
			}
		})
	}

