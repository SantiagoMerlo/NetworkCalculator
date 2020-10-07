package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese address")
	textAddress, _ := reader.ReadString('\n')

	address := stringToBinary(textAddress)

	fmt.Printf("Ingrese Netmask \n")
	textMask, _ := reader.ReadString('\n')

	mask := stringToBinary(textMask)

	wildCard := inverse(mask[:])

	broadcast := orBinary(address[:], wildCard[:])

	firstHost := addAddress(3, address[:], 1)

	lastHost := maxHost(broadcast[:])

	fmt.Println("\t\t-----------------------------------")
	fmt.Println("Clase: \t", clase(address))
	fmt.Println("Address: \t", address, "\t valor: ", binaryToInt(address[:]))
	fmt.Println("Mask:    \t", mask, "\t valor: ", binaryToInt(mask[:]))
	fmt.Println("WildCard: \t", wildCard, "\t valor: ", binaryToInt(wildCard[:]))
	fmt.Println("Broadcast: \t", broadcast, "\t valor: ", binaryToInt(broadcast[:]))
	fmt.Println("FirstHost: \t", stringToBinary(firstHost), "\t valor: ", firstHost)
	fmt.Println("LastHost: \t", stringToBinary(lastHost), "\t valor: ", lastHost)
	fmt.Println("\t\t----------------------------------- ")

	//fmt.Printf("Ingrese Categoria (A,B,C) \n")
	// A = al segundo octeto, B = 3, C = 4
	//categoria, _ := reader.ReadString('\n')

	fmt.Printf("\nIngrese subred 1 \n")
	net1, _ := reader.ReadString('\n')
	salto := 0
	salto, address = host(net1, address, 0, 2, false)

	fmt.Printf("Ingrese subred 2 \n")
	net2, _ := reader.ReadString('\n')
	salto, address = host(net2, address, salto, 2, false)

	fmt.Printf("Ingrese subred 3-1 \n")
	net3, _ := reader.ReadString('\n')
	if net2 == net3 {
		salto, address = host(net3, address, salto, 3, true)
		fmt.Printf("El Ultimo host y broadcast estan mal. se le tiene que sumar al first host la cantidad de host disponibles")
		fmt.Printf("A partir de este punto pueden estar mal los calculos")
	} else {
		salto, address = host(net3, address, salto, 3, false)
	}

	fmt.Printf("Ingrese subred 3-2 \n")
	net4, _ := reader.ReadString('\n')
	if net3 == net4 {
		salto, address = host(net4, address, salto, 3, true)
		fmt.Printf("El Ultimo host y broadcast estan mal. se le tiene que sumar al first host la cantidad de host disponibles")
	} else {
		salto, address = host(net4, address, salto, 3, false)
	}

}

func host(net string, address []string, salto int, octeto int, carry bool) (int, []string) {
	fmt.Println("\n\t\t-----------------------------------")
	addressSubred := addAddress(2, address, salto)
	number, _ := strconv.Atoi(net[:len(net)-1]) //eliminacion del '\n'
	subnet := math.RoundToEven(math.Log2(float64(number + 2)))
	fmt.Println("Cantidad de host disponibles: \t", (math.Pow(2, subnet) - 2))
	fmt.Println("Cantidad de bits apagados: \t", subnet)
	mask := maskForNumber(int(subnet))
	preJump := valueForArray(octeto, mask[:])
	jump := 256 - preJump
	fmt.Println("Salto: \t\t\t\t", jump)
	wildCard := inverse(mask[:])
	if carry {
		addressSubred = addAddress(3, stringToBinary(addressSubred), int((math.Pow(2, subnet))-1))
	}
	addressInString := stringToBinary(addressSubred)
	broadcast := orBinary(addressInString[:], wildCard[:])
	firstHost := addAddress(3, addressInString[:], 1)
	ultimoHost := maxHost(broadcast[:])
	fmt.Println("Address: \t", address, "\t valor: ", binaryToInt(addressInString[:]))
	fmt.Println("Mask:    \t", mask, "\t valor: ", binaryToInt(mask[:]))
	fmt.Println("WildCard: \t", wildCard, "\t valor: ", binaryToInt(wildCard[:]))
	fmt.Println("Broadcast: \t", broadcast, "\t valor: ", binaryToInt(broadcast[:]))
	fmt.Println("First Host: \t", stringToBinary(firstHost), "\t valor: ", firstHost)
	fmt.Println("Last Host: \t", stringToBinary(ultimoHost), "\t valor: ", ultimoHost)
	return int(jump), addressInString
}

func addAddress(pos int, address []string, value int) string {
	var temp [4]int64

	for i := 0; i < 4; i++ {
		number, _ := strconv.ParseInt(address[i], 2, 64)
		if i == pos {
			temp[i] = number + int64(value)
		} else {
			temp[i] = number
		}
	}
	return fmt.Sprint(temp[0], ".", temp[1], ".", temp[2], ".", temp[3])
}

func maxHost(address []string) string {
	value := valueForArray(3, address)
	flag := 3
	if value != 0 {
		value = value - 1
	} else {
		flag = 2
		value := valueForArray(2, address)
		if value != 0 {
			value = value - 1
		} else {
			flag = 1
			if value != 0 {
				value := valueForArray(1, address)
				value = value - 1
			}
		}
	}
	return newAddress(flag, address, int(value))
}

// suma un valor a nuestro address
func newAddress(pos int, address []string, value int) string {
	var temp [4]int64

	for i := 0; i < 4; i++ {
		number, _ := strconv.ParseInt(address[i], 2, 64)
		temp[i] = number
	}

	temp[pos] = int64(value)

	return fmt.Sprint(temp[0], ".", temp[1], ".", temp[2], ".", temp[3])
}

func valueForArray(pos int, address []string) int64 {
	var temp [4]int64

	for i := 0; i < 4; i++ {
		number, _ := strconv.ParseInt(address[i], 2, 64)
		temp[i] = number
	}

	return temp[pos]
}

func maskForNumber(number int) [4]string {
	var temp [4][8]string
	var response [4]string

	count := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if (32 - number) > count {
				temp[i][j] = "1"
			} else {
				temp[i][j] = "0"
			}
			count++
		}
	}

	response[0] = strings.Join(temp[0][:], "")
	response[1] = strings.Join(temp[1][:], "")
	response[2] = strings.Join(temp[2][:], "")
	response[3] = strings.Join(temp[3][:], "")
	return response
}

func stringToBinary(str string) []string {

	temp := ""
	var response [4]string
	i := 0

	for pos := range str {
		char := str[pos]

		if char == '.' || char == '\n' {
			number, _ := strconv.Atoi(temp)
			response[i] = fmt.Sprintf("%08b", byte(number))
			i++
			temp = ""
		} else if (len(str) - 1) == pos {
			temp = temp + string(char)
			number, _ := strconv.Atoi(temp)
			response[i] = fmt.Sprintf("%08b", byte(number))
		} else {
			temp = temp + string(char)
		}
	}
	return response[:]
}

func andBinary(str1 []string, str2 []string) [4]string {

	var temp [4][8]string
	var response [4]string

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if str1[i][j] == str2[i][j] && str1[i][j] == '1' {
				temp[i][j] = "1"
			} else {
				temp[i][j] = "0"
			}
		}
	}
	response[0] = strings.Join(temp[0][:], "")
	response[1] = strings.Join(temp[1][:], "")
	response[2] = strings.Join(temp[2][:], "")
	response[3] = strings.Join(temp[3][:], "")
	return response
}

func orBinary(str1 []string, str2 []string) [4]string {

	var temp [4][8]string
	var response [4]string

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if str2[i][j] == '1' || str1[i][j] == '1' {
				temp[i][j] = "1"
			} else {
				temp[i][j] = "0"
			}
		}
	}
	response[0] = strings.Join(temp[0][:], "")
	response[1] = strings.Join(temp[1][:], "")
	response[2] = strings.Join(temp[2][:], "")
	response[3] = strings.Join(temp[3][:], "")
	return response
}

func binaryToInt(str []string) string {

	var temp [4]string
	for i := 0; i < 4; i++ {
		number, _ := strconv.ParseInt(str[i], 2, 64)
		temp[i] = strconv.Itoa(int(number))
	}

	return fmt.Sprint(temp[0], ".", temp[1], ".", temp[2], ".", temp[3])
}

func inverse(mask []string) [4]string {

	var temp [4][8]string
	var response [4]string

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if mask[i][j] == '1' {
				temp[i][j] = "0"
			} else {
				temp[i][j] = "1"
			}
		}
	}

	response[0] = strings.Join(temp[0][:], "")
	response[1] = strings.Join(temp[1][:], "")
	response[2] = strings.Join(temp[2][:], "")
	response[3] = strings.Join(temp[3][:], "")
	return response
}

func clase(addres []string) string {
	value := valueForArray(0, addres[:])
	if value < 128 {
		return "A"
	} else if value < 192 {
		return "B"
	} else {
		return "C'"
	}
}
