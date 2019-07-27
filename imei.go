// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"errors"
	"fmt"
	"log"
)

//ParseIMEI takes teltonica UDP packet as a byte slice and return uint64 IMEI number or error
func ParseIMEI(bs *[]byte) (uint64, error) {
	//validate Teltonika packet ID on byte no. 3 and 4
	if len(*bs) < 22 || (*bs)[2] != 0xca || (*bs)[3] != 0xfe {
		return 0, errors.New("Invalid udp packet")
	}

	//parse IMEI length on byte 6 and 7
	imeiLen := ParseHex2Uint64(bs, 6, 8)
	if imeiLen != 15 && imeiLen != 16 {
		log.Panicf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
		return 0, fmt.Errorf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
	}

	//parse IMEI, IMEI number starts at 8 byte and stop on 8 + len
	imei, err := ParseHexString2Uint64(bs, 8, (8 + int32(imeiLen)))
	if err != nil {
		log.Printf("Error from HEX parser %v", err)
		return 0, errors.New("Not possible to parse IMEI from HEX utf-8")
	}

	//validate IMEI with Luhn algorithm
	if imeiLen == 15 {
		if ValidateIMEI(&imei) != true {
			return 0, errors.New("IMEI checksum is not valid")
		}
	}

	//count imei len and compary with len in the packet
	if iterativeDigitsCount(imei) != imeiLen {
		return 0, fmt.Errorf("IMEI len is not corresponding with len in packet, want %v got %v", imeiLen, iterativeDigitsCount(imei))
	}

	return imei, nil
}
