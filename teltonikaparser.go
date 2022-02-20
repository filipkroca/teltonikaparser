// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package teltonikaparser is an implementation of https://wiki.teltonika.lt/view/Codec Codec08 and Codec08Extended for UDP packets in GO Lang
// implemented https://wiki.teltonika.lt/view/Codec#Codec_8
// implemented https://wiki.teltonika.lt/view/Codec#Codec_8_Extended
package teltonikaparser

import (
	"fmt"

	"github.com/filipkroca/b2n"
)

// Decoded struct represent decoded Teltonika data structure with all AVL data as return from function Decode
type Decoded struct {
	IMEI     string    // IMEI number, if len==15 also validated by checksum
	CodecID  byte      // 0x08 (codec 8) or 0x8E (codec 8 extended)
	NoOfData uint8     // Number of Data
	Data     []AvlData // Slice with avl data
	Response []byte 	 // Slice with a response
}

// AvlData represent one block of data
type AvlData struct {
	UtimeMs  uint64    // Utime in mili seconds
	Utime    uint64    // Utime in seconds
	Priority uint8     // Priority, 	[0	Low, 1	High, 2	Panic]
	Lat      int32     // Latitude (between 850000000 and -850000000), fit int32
	Lng      int32     // Longitude (between 1800000000 and -1800000000), fit int32
	Altitude int16     // Altitude In meters above sea level, 2 bytes
	Angle    uint16    // Angle In degrees, 0 is north, increasing clock-wise, 2 bytes
	VisSat   uint8     // Satellites Number of visible satellites
	Speed    uint16    // Speed in km/h
	EventID  uint16    // Event generated (0 – data generated not on event)
	Elements []Element // Slice containing parsed IO Elements
}

// Element represent one IO element, before storing in a db do a conversion to IO datatype (1B, 2B, 4B, 8B)
type Element struct {
	Length uint16 // Length of element, this should be uint16 because Codec 8 extended has 2Byte of IO len
	IOID   uint16 // IO element ID
	Value  []byte // Value of the element represented by slice of bytes
}

// Decode takes a pointer to a slice of bytes with raw data and return Decoded struct
func Decode(bs *[]byte) (Decoded, error) {
	decoded := Decoded{}
	var err error
	var nextByte int

	// check for minimum packet size
	if len(*bs) < 45 {
		return Decoded{}, fmt.Errorf("Minimum packet size is 45 Bytes, got %v", len(*bs))
	}

	// check for teltonika packet ID
	if (*bs)[2] != 0xca || (*bs)[3] != 0xfe {
		return Decoded{}, fmt.Errorf("Probably not Teltonika packet, trashed")
	}

	// determine bit number where start data, it can change because of IMEI length
	imeiLenX, err := b2n.ParseBs2Uint8(bs, 7)
	if err != nil {
		return Decoded{}, fmt.Errorf("Decode error, %v", err)
	}
	imeiLen := int(imeiLenX)

	if imeiLen != 15 && imeiLen != 16 {
		//log.Fatalf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
		return Decoded{}, fmt.Errorf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
	}

	// decode and validate IMEI
	decoded.IMEI, err = b2n.ParseIMEI(bs, 8, imeiLen)
	if err != nil {
		return Decoded{}, fmt.Errorf("Decode error, %v", err)
	}

	// count start bit for data
	startByte := 8 + imeiLen

	// decode Codec ID
	decoded.CodecID = (*bs)[startByte]
	if decoded.CodecID != 0x08 && decoded.CodecID != 0x8e {
		return Decoded{}, fmt.Errorf("Invalid Codec ID, want 0x08 or 0x8E, get %v", decoded.CodecID)
	}

	// initialize nextByte counter
	nextByte = startByte + 1

	// determine no of data in packet
	decoded.NoOfData, err = b2n.ParseBs2Uint8(bs, nextByte)
	if err != nil {
		return Decoded{}, fmt.Errorf("Decode error, %v", err)
	}

	// increment nextByte counter
	nextByte++

	// make slice for decoded data
	decoded.Data = make([]AvlData, 0, decoded.NoOfData)
	// go through data
	for i := 0; i < int(decoded.NoOfData); i++ {

		decodedData := AvlData{}

		// time record in ms has 8 Bytes
		decodedData.UtimeMs, err = b2n.ParseBs2Uint64(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}

		decodedData.Utime = uint64(decodedData.UtimeMs / 1000)
		nextByte += 8

		// parse priority
		decodedData.Priority, err = b2n.ParseBs2Uint8(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		if !(decodedData.Priority <= 2) {
			return Decoded{}, fmt.Errorf("Invalid Priority value, want priority <= 2, got %v", decodedData.Priority)
		}

		nextByte++

		// parse and validate GPS
		decodedData.Lng, err = b2n.ParseBs2Int32TwoComplement(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		if !(decodedData.Lng > -1800000000 && decodedData.Lng < 1800000000) {
			return Decoded{}, fmt.Errorf("Invalid Lat value, want lat > -1800000000 AND lat < 1800000000, got %v", decodedData.Lng)
		}
		nextByte += 4

		decodedData.Lat, err = b2n.ParseBs2Int32TwoComplement(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}

		if !(decodedData.Lat > -850000000 && decodedData.Lat < 850000000) {
			return Decoded{}, fmt.Errorf("Invalid Lat value, want lat > -850000000 AND lat < 850000000, got %v", decodedData.Lat)
		}
		nextByte += 4

		// parse Altitude
		decodedData.Altitude, err = b2n.ParseBs2Int16TwoComplement(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		if !(decodedData.Altitude > -5000 && decodedData.Altitude < 12000) {
			return Decoded{}, fmt.Errorf("Invalid Altitude value, want Altitude > -5000 AND Altitude < 12000, got %v", decodedData.Altitude)
		}
		nextByte += 2

		// parse Angle
		decodedData.Angle, err = b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		if decodedData.Angle > 360 {
			return Decoded{}, fmt.Errorf("Invalid Angle value, want Angle <= 360, got %v", decodedData.Angle)
		}
		nextByte += 2

		// parse num. of vissible sattelites VisSat
		decodedData.VisSat, err = b2n.ParseBs2Uint8(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		nextByte++

		// parse Speed
		decodedData.Speed, err = b2n.ParseBs2Uint16(bs, nextByte)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}
		nextByte += 2

		// parse EventID
		if decoded.CodecID == 0x8e {
			// if Codec 8 extended is used, Event id has size 2 bytes
			decodedData.EventID, err = b2n.ParseBs2Uint16(bs, nextByte)
			if err != nil {
				return Decoded{}, fmt.Errorf("Decode error, %v", err)
			}

			nextByte += 2
		} else {
			x, err := b2n.ParseBs2Uint8(bs, nextByte)
			if err != nil {
				return Decoded{}, fmt.Errorf("Decode error, %v", err)
			}
			decodedData.EventID = uint16(x)
			nextByte++
		}

		decodedIO, endByte, err := DecodeElements(bs, nextByte, decoded.CodecID)
		if err != nil {
			return Decoded{}, fmt.Errorf("Decode error, %v", err)
		}

		nextByte = endByte
		decodedData.Elements = decodedIO

		decoded.Data = append(decoded.Data, decodedData)

	}

	if int(decoded.NoOfData) != len(decoded.Data) {
		return Decoded{}, fmt.Errorf("Error when counting number of parsed data, want %v, got %v", int(decoded.NoOfData), len(decoded.Data))
	}

	// check if packet was corretly parsed
	endNoOfData := (*bs)[nextByte]
	if decoded.NoOfData != endNoOfData {
		return Decoded{}, fmt.Errorf("Unexpected byte representing control num. of data on end of parsing, want %#x, got %#x", decoded.NoOfData, endNoOfData)
	}

	// create response packet
	decoded.Response = []byte{0x00, 0x05, 0xCA, 0xFE, 0x01, (*bs)[4], decoded.NoOfData}

	return decoded, nil
}
