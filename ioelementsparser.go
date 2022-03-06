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
		// if Codec 8 extended is used, Event id has size 2 bytes
		// Codec ID	0x08	0x8E
		// AVL Data IO element length	1 Byte	2 Bytes
		// AVL Data IO element total IO count length	1 Byte	2 Bytes
		// AVL Data IO element IO count length	1 Byte	2 Bytes
		// AVL Data IO element AVL ID length	1 Byte	2 Bytes
		codecLenDel = 2
	}
	// parse number of elements and prepare array
	if codecID == 0x8e {
		x, err := b2n.ParseBs2Uint16(bs, start)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements error %v", err)
		}

		totalElements = int(x)
	} else if codecID == 0x08 {
		x, err := b2n.ParseBs2Uint8(bs, start)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements error %v", err)
		}

		totalElements = int(x)
	}
	totalElementsChecksum := 0
	// make a slice
	ElementsBS := make([]Element, 0, totalElements)

	// start parsing data
	nextByte := start + codecLenDel

	// parse 1Byte ios
	x, err := b2n.ParseBs2Uint8(bs, nextByte)
	if err != nil {
		return []Element{}, 0, fmt.Errorf("DecodeElements error %v", err)
	}
	noOfElements := int(x)

	if codecID == 0x8e {
		z, err := b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements error %v", err)
		}
		noOfElements = int(z)
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		cutted, err := cutIO(bs, nextByte, codecLenDel, 1)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements 1B error %v", err)
		}
		//append element to the returned slice
		ElementsBS = append(ElementsBS, cutted)
		nextByte += codecLenDel + 1
		totalElementsChecksum++
	}

	// parse 2Byte ios
	noOfElementsX, err := b2n.ParseBs2Uint8(bs, nextByte)
	if err != nil {
		return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 2B error %v", err)
	}
	noOfElements = int(noOfElementsX)

	if codecID == 0x8e {
		noOfElementsX, err := b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 2B Extended Codec error %v", err)
		}
		noOfElements = int(noOfElementsX)
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		cutted, err := cutIO(bs, nextByte, codecLenDel, 2)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements 2B error %v", err)
		}
		// append element to the returned slice
		ElementsBS = append(ElementsBS, cutted)
		nextByte += codecLenDel + 2
		totalElementsChecksum++
	}

	//parse 4Byte ios
	noOfElementsX, err = b2n.ParseBs2Uint8(bs, nextByte)
	if err != nil {
		return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 4B error %v", err)
	}
	noOfElements = int(noOfElementsX)

	if codecID == 0x8e {
		noOfElementsX, err := b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 4B Extended Codec error %v", err)
		}
		noOfElements = int(noOfElementsX)
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		cutted, err := cutIO(bs, nextByte, codecLenDel, 4)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements 4B error %v", err)
		}
		// append element to the returned slice
		ElementsBS = append(ElementsBS, cutted)
		nextByte += codecLenDel + 4
		totalElementsChecksum++
	}

	//parse 8Byte ios
	noOfElementsX, err = b2n.ParseBs2Uint8(bs, nextByte)
	if err != nil {
		return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 8B error %v", err)
	}
	noOfElements = int(noOfElementsX)

	if codecID == 0x8e {
		noOfElementsX, err := b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements 8B Extended Codec error %v", err)
		}
		noOfElements = int(noOfElementsX)
	}

	nextByte = nextByte + codecLenDel

	for ioB := 0; ioB < noOfElements; ioB++ {
		cutted, err := cutIO(bs, nextByte, codecLenDel, 8)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements 8B error %v", err)
		}
		// append element to the returned slice
		ElementsBS = append(ElementsBS, cutted)
		nextByte += codecLenDel + 8
		totalElementsChecksum++
	}

	if codecID == 0x8e {
		//parse variableByte ios, only Codec 8 extended

		noOfElementsX, err := b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return []Element{}, 0, fmt.Errorf("DecodeElements noOfElements variableB Extended Codec error %v", err)
		}
		noOfElements = int(noOfElementsX)

		nextByte = nextByte + codecLenDel

		for ioB := 0; ioB < noOfElements; ioB++ {
			cutted, err := cutIOxLen(bs, nextByte)
			if err != nil {
				return []Element{}, 0, fmt.Errorf("DecodeElements 2B error %v", err)
			}
			// append element to the returned slice
			ElementsBS = append(ElementsBS, cutted)
			nextByte += 4 + int(cutted.Length)
			totalElementsChecksum++
		}

	}

	if totalElementsChecksum != totalElements {
		//log.Fatalf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
		return []Element{}, 0, fmt.Errorf("Error when counting parsed IO Elements, want %v, got %v", totalElements, totalElementsChecksum)
	}

	return ElementsBS, nextByte, nil

}

// cutIO cuts a static length elements
func cutIO(bs *[]byte, start int, idLen int, length int) (Element, error) {
	curIO := Element{}
	//determine length of this sized elements (num. of 1Bytes elements, num. of 2Bytes elements ...)
	curIO.Length = uint16(length)

	var err error
	var curIOX uint8
	//parse element ID according to the length of ID [1, 2] Byte
	if idLen == 1 {
		curIOX, err = b2n.ParseBs2Uint8(bs, start)
		curIO.IOID = uint16(curIOX)
	} else if idLen == 2 {
		curIO.IOID, err = b2n.ParseBs2Uint16(bs, start)
	}
	if err != nil {
		return Element{}, fmt.Errorf("cutIO error ParseBs2Uint8 or ParseBs2Uint16, %v", err)
	}

	if (start + idLen + length) > len(*bs) {
		return Element{}, fmt.Errorf("cutIO error, want minimum length of bs %v, got %v, packet %x", start+idLen+length, len(*bs), *bs)
	}

	curIO.Value = (*bs)[start+idLen : start+idLen+length]

	return curIO, nil
}

// cutIOxLen cuts a variable length elements
func cutIOxLen(bs *[]byte, start int) (Element, error) {
	curIO := Element{}

	var err error
	//parse element ID according to the length of ID [1, 2] Byte
	curIO.IOID, err = b2n.ParseBs2Uint16(bs, start)
	if err != nil {
		return Element{}, fmt.Errorf("cutIOxLen error, %v", err)
	}

	//determine length of this variable element
	curIO.Length, err = b2n.ParseBs2Uint16(bs, start+2)
	if err != nil {
		return Element{}, fmt.Errorf("cutIOxLen error, %v", err)
	}

	curIO.Value = (*bs)[start+4 : start+4+int(curIO.Length)]

	return curIO, nil
}
