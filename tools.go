// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"math"
	"strconv"
)

//ParseHex2Uint64 takes a byte slice pointer, start byte, stop byte and returns Uint64 parsed as it was HEX number
func ParseHex2Uint64(bs *[]byte, start int32, stop int32) uint64 {

	var sum uint64
	var order uint32

	//convert hex byte slice to Uint64
	for i := stop - 1; i >= start; i-- {
		//shift to the left by 8 bits overy cycle
		sum += uint64((*bs)[i]) << order
		order += 8
	}

	return sum
}

//ParseHex2Int32TwoComplement takes a byte slice pointer, start byte, stop byte and returns Int32 parsed as it was HEX number coded with Two Complement Arithmetic
func ParseHex2Int32TwoComplement(bs *[]byte, start int32, stop int32) int32 {
	var sum int32
	var order uint32
	var signed bool

	//mask last Byte with mask (1000 0000) then shift by 7 bits and check sign bit
	if (*bs)[start]&0x80>>7 == 1 {
		signed = true
	}

	//convert hex byte slice to int32
	for i := stop - 1; i >= start; i-- {
		cb := (*bs)[i]
		//if signed do a XOR operation on every Byte
		if signed {
			cb ^= 0xFF
		}
		//shift to the left by 8 bits every cycle
		sum += int32(cb) << order
		order += 8
	}

	//if signed, increment with complement 1 and multiply by -1
	if signed {
		sum++
		sum = sum * -1
	}

	return sum
}

//ParseHexString2Uint64 takes a byte slice pointer, start byte, stop byte and returns Uint64 parsed as it was int string
func ParseHexString2Uint64(bs *[]byte, start int32, stop int32) (uint64, error) {
	var str string

	//convert hex byte slice to Uint64
	for _, x := range (*bs)[start:stop] {
		str += string(x)
	}

	sumUint64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return sumUint64, nil
}

func iterativeDigitsCount(number uint64) uint64 {
	var count uint64
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

//ValidateIMEI takes 15 digits long IMEI number and return bool validity
func ValidateIMEI(number *uint64) bool {
	checkSumDigit := digit(number, uint64(1))
	var checkSum uint64

	//make buffer array for Luhn algorithm with len 14 bytes and cap 31 bytes
	digits := make([]uint8, 14, 31)

	for i := 15; i > 1; i-- {

		digits[i-2] = digit(number, uint64(i))

		if i%2 == 0 {
			digits[i-2] = digits[i-2] * 2
		}

		if digits[i-2] >= 10 {
			digits = append(digits, 1)
			digits[i-2] = digits[i-2] % 10
		}
	}

	for _, val := range digits {
		checkSum += uint64(val)
	}

	//return divider to 10 is same as checkSumDigit
	return ((10 - checkSum%10) == uint64(checkSumDigit))
}

//digit takes number and return one digit on a certain position
func digit(num *uint64, place uint64) uint8 {
	r := *num % uint64(math.Pow(10, float64(place)))
	return uint8(r / uint64(math.Pow(10, float64(place-1))))
}
