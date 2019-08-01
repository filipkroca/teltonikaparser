// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"fmt"

	"github.com/filipkroca/b2n"
)

// DecodeElements take pointer to a byte slice with raw data, start Byte position and Codec ID, and returns slice of Element
func DecodeElements(bs *[]byte, start int, codecID byte) ([]Element, int, error) {

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
	ElementsBS := make([]Element, 0, totalElements)

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
		ElementsBS = append(ElementsBS, cutIO(bs, nextByte, codecLenDel, 1))
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
		ElementsBS = append(ElementsBS, cutIO(bs, nextByte, codecLenDel, 2))
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
		ElementsBS = append(ElementsBS, cutIO(bs, nextByte, codecLenDel, 4))
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
		ElementsBS = append(ElementsBS, cutIO(bs, nextByte, codecLenDel, 8))
		nextByte += codecLenDel + 8
		totalElementsChecksum++
	}

	if codecID == 0x8e {
		//parse variableByte ios, only Codec 8 extended

		noOfElements = int(b2n.ParseBs2Uint16(bs, nextByte))

		nextByte = nextByte + codecLenDel

		for ioB := 0; ioB < noOfElements; ioB++ {
			//append element to the returned slice
			ElementsBS = append(ElementsBS, cutIOxLen(bs, nextByte))
			nextByte += codecLenDel + 2
			totalElementsChecksum++
		}
	}

	if totalElementsChecksum != totalElements {
		//log.Fatalf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
		return []Element{}, 0, fmt.Errorf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
	}

	return ElementsBS, nextByte, nil

}

func cutIO(bs *[]byte, start int, idLen int, length int) Element {
	curIO := Element{}
	//determine length of this sized elements (num. of 1Bytes elements, num. of 2Bytes elements ...)
	curIO.Length = uint16(length)

	//parse element ID according to the length of ID [1, 2] Byte
	if idLen == 1 {
		curIO.IOID = uint16(b2n.ParseBs2Uint8(bs, start))
	} else if idLen == 2 {
		curIO.IOID = b2n.ParseBs2Uint16(bs, start)
	}

	if (start+idLen+length) < len(*bs) {
		//log.Fatalf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
		//return Element{}, fmt.Errorf("cutIO error, want minimum length of bs %v, got %v", start+idLen+length, len(*bs))
		return Element{}
	}

	curIO.Value = (*bs)[start+idLen : start+idLen+length]

	return curIO
}

func cutIOxLen(bs *[]byte, start int) Element {
	curIO := Element{}

	//parse element ID according to the length of ID [1, 2] Byte
	curIO.IOID = b2n.ParseBs2Uint16(bs, start)

	//determine length of this variable element
	curIO.Length = b2n.ParseBs2Uint16(bs, start+2)

	curIO.Value = (*bs)[start+4 : start+4+int(curIO.Length)]

	return curIO
}
