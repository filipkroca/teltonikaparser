// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Package teltonikaparser is an implementation of https://wiki.teltonika.lt/view/Codec Codec08 and Codec08Extended for UDP packets in GO Lang
// implemented https://wiki.teltonika.lt/view/Codec#Codec_8
// implemented https://wiki.teltonika.lt/view/Codec#Codec_8_Extended
package teltonikaparser

import (
	"fmt"
	"log"
)

//Decoded struct represent decoded Teltonika data structure with all AVL data as return from function Decode
type Decoded struct {
	IMEI     uint64    //IMEI number, if len==15 also validated by checksum
	CodecID  byte      //0x08 (codec 8) or 0x8E (codec 8 extended)
	NoOfData byte      //Number of Data
	Data     []AvlData //Slice with avl data
}

//AvlData represent one array of data
type AvlData struct {
	UtimeMs    uint64      //Utime in mili seconds
	Utime      uint64      //Utime in seconds
	Priority   uint8       //Priority, 	[0	Low, 1	High, 2	Panic]
	Lat        int32       //Latitude (between 850000000 and -850000000), fit int32
	Lng        int32       //Longitude (between 1800000000 and -1800000000), fit int32
	Altitude   int16       //Altitude In meters above sea level, 2 bytes
	Angle      uint16      //Angle In degrees, 0 is north, increasing clock-wise, 2 bytes
	VisSat     uint8       //Satellites Number of visible satellites
	Speed      uint16      //Speed in km/h
	EventID    uint16      //Event generated (0 – data generated not on event)
	IOElements []IOElement //Slice containing parsed IO Elements
}

//IOElement represent one IO element, before storing in a db do a conversion to IO datatype (1B, 2B, 4B, 8B)
type IOElement struct {
	Length uint16 //Length of element, this should be uint16 because Codec 8 extended has 2Byte of IO len
	IOID   uint16 //IO element ID
	Value  []byte //Value of the element represented by slice of bytes
}

//Decode takes a pointer to a slice of bytes with raw data and return Decoded struct
func Decode(bs *[]byte) (Decoded, error) {
	decoded := Decoded{}
	var err error
	var nextByte uint32

	//decode and validate IMEI
	decoded.IMEI, err = ParseIMEI(bs)
	if err != nil {
		log.Fatalf("Error when decoding IMEI, %v", err)
		return Decoded{}, fmt.Errorf("Error when decoding IMEI, %v", err)
	}

	//determine bit number where start data, it can change because of IMEI length
	imeiLen := ParseHex2Uint64(bs, 6, 8)
	if imeiLen != 15 && imeiLen != 16 {
		log.Fatalf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
		return Decoded{}, fmt.Errorf("Error when determining IMEI len want 15 or 16, got %v", imeiLen)
	}

	//count start bit for data
	startByte := 8 + uint32(imeiLen)

	//decode Codec ID
	decoded.CodecID = (*bs)[startByte]
	if decoded.CodecID != 0x08 && decoded.CodecID != 0x8e {
		log.Fatalf("Invalid Codec ID, want 0x08 or 0x8E, get %v", decoded.CodecID)
		return Decoded{}, fmt.Errorf("Invalid Codec ID, want 0x08 or 0x8E, get %v", decoded.CodecID)
	}

	//initialize nextByte counter
	nextByte = startByte + 1

	//determine no of data in packet
	decoded.NoOfData = (*bs)[nextByte]

	//increment nextByte counter
	nextByte++

	//make slice for decoded data
	decoded.Data = make([]AvlData, 0, decoded.NoOfData)

	//go through data
	for i := 0; i < int(decoded.NoOfData); i++ {
		decodedData := AvlData{}

		//time record in ms has 8 Bytes
		decodedData.UtimeMs = ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+8))
		decodedData.Utime = uint64(decodedData.UtimeMs / 1000)
		nextByte += 8

		//parse priority
		decodedData.Priority = uint8((*bs)[nextByte])
		nextByte++

		//parse and validate GPS
		decodedData.Lat = ParseHex2Int32TwoComplement(bs, int32(nextByte), int32(nextByte+4))
		if !(decodedData.Lat > -850000000 && decodedData.Lat < 850000000) {
			log.Fatalf("Invalid Lat value, want lat > -850000000 AND lat < 850000000, got %v", decodedData.Lat)
			return Decoded{}, fmt.Errorf("Invalid Lat value, want lat > -850000000 AND lat < 850000000, got %v", decodedData.Lat)
		}
		nextByte += 4
		decodedData.Lng = ParseHex2Int32TwoComplement(bs, int32(nextByte), int32(nextByte+4))
		if !(decodedData.Lng > -1800000000 && decodedData.Lng < 1800000000) {
			log.Fatalf("Invalid Lat value, want lat > -1800000000 AND lat < 1800000000, got %v", decodedData.Lng)
			return Decoded{}, fmt.Errorf("Invalid Lat value, want lat > -1800000000 AND lat < 1800000000, got %v", decodedData.Lng)
		}
		nextByte += 4

		//parse Altitude
		decodedData.Altitude = int16(ParseHex2Int32TwoComplement(bs, int32(nextByte), int32(nextByte+2)))
		nextByte += 2

		//parse Altitude
		decodedData.Angle = uint16(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+2)))
		if decodedData.Angle > 360 {
			log.Fatalf("Invalid Angle value, want Angle <= 360, got %v", decodedData.Angle)
			return Decoded{}, fmt.Errorf("Invalid Angle value, want Angle <= 360, got %v", decodedData.Angle)
		}
		nextByte += 2

		//parse num. of vissible sattelites VisSat
		decodedData.VisSat = uint8(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+1)))
		nextByte++

		//parse Speed
		decodedData.Speed = uint16(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+2)))
		nextByte += 2

		//parse EventID
		if decoded.CodecID == 0x8e {
			//if Codec 8 extended is used, Event id has size 2 bytes
			decodedData.EventID = uint16(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+2)))
			nextByte += 2
		} else {
			decodedData.EventID = uint16(ParseHex2Uint64(bs, int32(nextByte), int32(nextByte+1)))
			nextByte++
		}

		decodedIO, endByte, err := DecodeIOElements(bs, int32(nextByte), decoded.CodecID)
		if err != nil {
			log.Fatalf("Error when parsing IO Elements, %v", err)
			return Decoded{}, fmt.Errorf("Error when parsing IO Elements, %v", err)
		}

		nextByte = endByte
		decodedData.IOElements = decodedIO

		decoded.Data = append(decoded.Data, decodedData)

	}

	if int(decoded.NoOfData) != len(decoded.Data) {
		log.Fatalf("Error when counting number of parsed data, want %v, got %v", int(decoded.NoOfData), len(decoded.Data))
		return Decoded{}, fmt.Errorf("Error when counting number of parsed data, want %v, got %v", int(decoded.NoOfData), len(decoded.Data))
	}

	//check if packet was corretly parsed
	endNoOfData := (*bs)[nextByte]
	if decoded.NoOfData != endNoOfData {
		log.Fatalf("Unexpected byte representing control num. of data on end of parsing, want %#x, got %#x", decoded.NoOfData, endNoOfData)
		return Decoded{}, fmt.Errorf("Unexpected byte representing control num. of data on end of parsing, want %#x, got %#x", decoded.NoOfData, endNoOfData)
	}

	return decoded, nil
}
