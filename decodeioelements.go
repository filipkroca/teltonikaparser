// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"fmt"
	"log"

	"github.com/filipkroca/b2n"
)

//DecodeIOElements take pointer to a byte slice with raw data, start Byte position and Codec ID, and returns slice of IOElement
func DecodeIOElements(bs *[]byte, start int, codecID byte) ([]IOElement, int, error) {

	var totalElements int
	codecLenDel := 1
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
	if codecID == 0x8e {
		totalElements = int(b2n.ParseBs2Uint16(bs, start))
	} else if codecID == 0x08 {
		totalElements = int(b2n.ParseBs2Uint8(bs, start))
	}
	totalElementsChecksum := 0
	//make a slice
	ioElementsBS := make([]IOElement, 0, totalElements)

	//start parsing data
	nextByte := start + codecLenDel

	//parse 1Byte ios
	noOfElements := int(b2n.ParseBs2Uint8(bs, nextByte))
	if codecID == 0x8e {
		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 1))
		nextByte += codecLenDel + 1
		totalElementsChecksum++
	}

	//parse 2Byte ios
	noOfElements = int(b2n.ParseBs2Uint8(bs, nextByte))
	if codecID == 0x8e {
		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 2))
		nextByte += codecLenDel + 2
		totalElementsChecksum++
	}

	//parse 4Byte ios
	noOfElements = int(b2n.ParseBs2Uint8(bs, nextByte))
	if codecID == 0x8e {
		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 4))
		nextByte += codecLenDel + 4
		totalElementsChecksum++
	}

	//parse 8Byte ios
	noOfElements = int(b2n.ParseBs2Uint8(bs, nextByte))
	if codecID == 0x8e {
		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		//append element to the returned slice
		ioElementsBS = append(ioElementsBS, cutIO(bs, nextByte, codecLenDel, 8))
		nextByte += codecLenDel + 8
		totalElementsChecksum++
	}

	if codecID == 0x8e {
		//parse variableByte ios, only Codec 8 extended

		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))

		nextByte = nextByte + codecLenDel

		for ioB := 0; ioB < noOfElements; ioB++ {
			//append element to the returned slice
			ioElementsBS = append(ioElementsBS, cutIOxLen(bs, nextByte))
			nextByte += codecLenDel + 2
			totalElementsChecksum++
		}
	}

	if totalElementsChecksum != totalElements {
		log.Fatalf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
		return []IOElement{}, 0, fmt.Errorf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
	}

	return ioElementsBS, nextByte, nil

}

func cutIO(bs *[]byte, start int, idLen int, len int) IOElement {
	curIO := IOElement{}
	//determine length of this sized elements (num. of 1Bytes elements, num. of 2Bytes elements ...)
	curIO.Length = uint16(len)

	//parse element ID according to the length of ID [1, 2] Byte
	if idLen == 1 {
		curIO.IOID = uint16(b2n.ParseBs2Uint8(bs, start))
	} else if idLen == 2 {
		curIO.IOID = b2n.ParseBs2Uint16(bs, start)
	}

	curIO.Value = (*bs)[start+idLen : start+idLen+len]

	return curIO
}

func cutIOxLen(bs *[]byte, start int) IOElement {
	curIO := IOElement{}

	//parse element ID according to the length of ID [1, 2] Byte
	curIO.IOID = b2n.ParseBs2Uint16(bs, start)

	//determine length of this variable element
	curIO.Length = b2n.ParseBs2Uint16(bs, start+2)

	curIO.Value = (*bs)[start+4 : start+4+int(curIO.Length)]

	return curIO
}
