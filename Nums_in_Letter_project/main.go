package main

import "fmt"

func Sonlar(n int) {
	if n >= 1 && n < 1000000000000 {

		if n >= 1000000000 {
			var yuzmilliard = (n % 1000000000000) / 100000000000
			switch yuzmilliard {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz  ")
			case 5:
				fmt.Print(" besh yuz")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz")
			case 9:
				fmt.Print(" to'qqiz yuz ")

			}

			var onmilliard = (n % 100000000000) / 10000000000
			switch onmilliard {
			case 1:
				fmt.Print(" on ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" toqson ")

			}

			var birmilliard = (n % 10000000000) / 1000000000
			switch birmilliard {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var milliard string = " milliard "
			fmt.Print(milliard)
		}

		if n >= 1000000 {
			var yuzmillion = (n % 1000000000) / 100000000
			switch yuzmillion {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz  ")
			case 5:
				fmt.Print(" besh yuz")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz")
			case 9:
				fmt.Print(" to'qqiz yuz ")

			}

			var onmillion = (n % 100000000) / 10000000
			switch onmillion {
			case 1:
				fmt.Print(" on ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" to'qson ")

			}

			var birmillion = (n % 10000000) / 1000000
			switch birmillion {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var million string = " million "
			fmt.Print(million)
		}
		if n >= 1000 {
			var yuzming = (n % 1000000) / 100000
			switch yuzming {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz ")
			case 5:
				fmt.Print(" besh yuz ")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz ")
			case 9:
				fmt.Print(" to'qqiz yuz ")
			}
			var onming = (n % 100000) / 10000
			switch onming {
			case 1:
				fmt.Print(" on  ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" toqson ")

			}

			var minglar = (n % 10000) / 1000
			switch minglar {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var ming string = " ming "
			fmt.Print(ming)
		}

		var yuz = (n % 1000) / 100
		switch yuz {
		case 1:
			fmt.Print(" bir yuz ")
		case 2:
			fmt.Print(" ikki yuz ")
		case 3:
			fmt.Print(" uch yuz ")
		case 4:
			fmt.Print(" to'rt yuz ")
		case 5:
			fmt.Print(" besh yuz ")
		case 6:
			fmt.Print(" olti yuz ")
		case 7:
			fmt.Print(" yetti yuz ")
		case 8:
			fmt.Print(" sakkiz yuz ")
		case 9:
			fmt.Print(" to'qqiz yuz ")
		}
		var onlar = (n % 100) / 10
		switch onlar {
		case 1:
			fmt.Print(" on ")
		case 2:
			fmt.Print(" yigirma ")
		case 3:
			fmt.Print(" o'ttiz ")
		case 4:
			fmt.Print(" qirq ")
		case 5:
			fmt.Print(" ellik ")
		case 6:
			fmt.Print(" oltmish ")
		case 7:
			fmt.Print(" yetmish ")
		case 8:
			fmt.Print(" sakson ")
		case 9:
			fmt.Print(" to'qson ")
		}
		var birlar int = n % 10
		switch birlar {
		case 1:
			fmt.Print(" bir ")
		case 2:
			fmt.Print(" ikki ")
		case 3:
			fmt.Print(" uch ")
		case 4:
			fmt.Print(" to'rt ")
		case 5:
			fmt.Print(" besh ")
		case 6:
			fmt.Print(" olti ")
		case 7:
			fmt.Print(" yetti ")
		case 8:
			fmt.Print(" sakkiz ")
		case 9:
			fmt.Print(" to'qqiz ")
		}

	} else {
		fmt.Println("Stop!!!")
	}

}

func main() {
	Sonlar(123223124)
}
