// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"fmt"
	"log"
)

//DecodeIOElements take pointer to a byte slice with raw data, start Byte position and Codec ID, and returns slice of IOElement
func DecodeIOElements(bs *[]byte, start int32, codecID byte) ([]IOElement, uint32, error) {

	var codecLenDel int32 = 1
	if codecID == 0x8e {
		//if Codec 8 extended is used, Event id has size 2 bytes
		//Codec ID	0x08	0x8E
		//AVL Data IO element length	1 Byte	2 Bytes
		//AVL Data IO element total IO count length	1 Byte	2 Bytes
		//AVL Data IO element IO count length	1 Byte	2 Bytes
		//AVL Data IO element AVL ID length	1 Byte	2 Bytes
		codecLenDel = 2
	}

	//parse number of elements and prepare array
	totalElements := int(ParseHex2Uint64(bs, int32(start), int32(start+codecLenDel)))
	totalElementsChecksum := 0
	//make array
	ioElementsBS := make([]IOElement, 0, totalElements)

	nextByte := start + codecLenDel

	//parse 1Byte ios
	noOfElements := int(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+codecLenDel)))
	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 1))
		nextByte += codecLenDel + 1
		totalElementsChecksum++
	}

	//parse 2Byte ios
	noOfElements = int(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+codecLenDel)))
	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 2))
		nextByte += codecLenDel + 2
		totalElementsChecksum++
	}

	//parse 4Byte ios
	noOfElements = int(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+codecLenDel)))
	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 4))
		nextByte += codecLenDel + 4
		totalElementsChecksum++
	}

	//parse 8Byte ios
	noOfElements = int(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+codecLenDel)))
	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 8))
		nextByte += codecLenDel + 8
		totalElementsChecksum++
	}

	if codecID == 0x8e {
		//parse variableByte ios, only Codec 8 extended
		noOfElements = int(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+codecLenDel)))
		nextByte = nextByte + codecLenDel

		for ioB := 0; ioB < noOfElements; ioB++ {
			//append element to the returned slice
			ioElx := cutIOxLen(bs, nextByte)
			ioElementsBS = append(ioElementsBS, ioElx)
			nextByte += 2 + int32(ioElx.Length)
			totalElementsChecksum++
		}
	}

	if totalElementsChecksum != totalElements {
		log.Fatalf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
		return []IOElement{}, 0, fmt.Errorf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
	}

	return ioElementsBS, uint32(nextByte), nil

}

func cutIO(bs *[]byte, start int32, idLen int32, len int32) IOElement {
	curIO := IOElement{}
	//determine length of this sized elements (num. of 1Bytes elements, num. of 2Bytes elements ...)
	curIO.Length = uint16(len)
	curIO.IOID = uint16(ParseHex2Uint64(bs, start, start+idLen))
	curIO.Value = (*bs)[start+idLen : start+idLen+len]

	return curIO
}

func cutIOxLen(bs *[]byte, start int32) IOElement {
	curIO := IOElement{}

	//determine length of this variable element
	curIO.Length = uint16(ParseHex2Uint64(bs, start+2, start+4))

	curIO.IOID = uint16(ParseHex2Uint64(bs, start, start+2))
	curIO.Value = (*bs)[start+4 : start+4+int32(curIO.Length)]

	return curIO
}
