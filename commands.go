// Copyright 2022 Gábor Nyíri. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package teltonikaparser is an implementation of https://wiki.teltonika.lt/view/Codec Codec12 for UDP packets in GO Lang
// implemented https://wiki.teltonika-gps.com/view/Codec#Codec_12
package teltonikaparser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/basvdlei/gotsmart/crc16"
)

const (
	CodecID             = 0x0C
	RequestPreamble     = 0x00000000
	ResponsePreamble    = 0x00000000
	CommandTypeRequest  = 0x05
	CommandTypeResponse = 0x06
)

type commandRequestPre struct {
	// Preamble - the packet starts with four zero bytes.
	Preamble uint32
	// Data Size - size is calculated from the Codec ID field to the second command quantity field.
	DataSize uint32
	// Codec ID - in Codec12 it is always 0x0C.
	CodecID byte
	// Command Quantity 1 - it is ignored when parsing the message.
	CommandQuantity1 byte
	// Type - it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Command Size – command length.
	CommandSize uint32
}

type commandRequestPost struct {
	// Command Quantity 2 - a byte which defines how many records (commands)
	// are in the packet. This byte will not be parsed but it’s recommended
	// that it should contain the same value as Command Quantity 1.
	CommandQuantity2 byte
	// Calculated from Codec ID to the Command Quantity 2. CRC (Cyclic
	// Redundancy Check) is an error-detecting code using for detect
	// accidental changes to RAW data. For calculation we are using CRC-16/IBM.
	CRC uint32
}

type CommandRequest struct {
	commandRequestPre
	// Command - command  in HEX.
	Command []byte // dynamic long type. Needs to read separately.
	commandRequestPost
}

type commandResponsePre struct {
	// Preamble - the packet starts with four zero bytes.
	Preamble uint32
	// Data Size - size is calculated from the Codec ID field to the second command or response quantity field.
	DataSize uint32
	// Codec ID - in Codec12 it is always 0x0C.
	CodecID byte
	// Response Quantity 1 - it is ignored when parsing the message.
	ResponseQuantity1 byte
	// Type - it can be 0x05 to denote command or 0x06 to denote response.
	Type byte
	// Response Size – command or response length.
	ResponseSize uint32
}

type commandResponsePost struct {
	// Response Quantity 2 - a byte which defines how many records (responses) are in the packet.
	// This byte will not be parsed but it’s recommended that it should contain the same value as Response Quantity 1.
	ResponseQuantity2 byte
	// calculated from Codec ID to the Command Quantity 2. CRC (Cyclic Redundancy Check) is an error-detecting
	// code using for detect accidental changes to RAW data. For calculation we are using CRC-16/IBM.
	CRC uint32
}

type CommandResponse struct {
	commandResponsePre
	// Response – response in HEX.
	Response []byte // dynamic long type. Needs to read separately.
	commandResponsePost
}

func EncodeCommandRequest(command string) ([]byte, error) {
	buffer := new(bytes.Buffer)

	commandRequest := CommandRequest{
		commandRequestPre: commandRequestPre{
			Preamble:         RequestPreamble,
			DataSize:         uint32(7 + len(command) + 1), // 7 header bytes + actual command text + 1 byte more
			CodecID:          CodecID,
			CommandQuantity1: 0x01,
			Type:             CommandTypeRequest, // 0x05 for command request, 0x06 for command response,
			CommandSize:      uint32(len(command)),
		},
		Command: []byte(command),
		commandRequestPost: commandRequestPost{
			CommandQuantity2: 0x01,
			CRC:              0, // let's add it separately when rest of the package is ready
		},
	}

	err := binary.Write(buffer, binary.BigEndian, commandRequest.commandRequestPre)
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("%v", err)
	}

	err = binary.Write(buffer, binary.BigEndian, commandRequest.Command)
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("%v", err)
	}

	err = binary.Write(buffer, binary.BigEndian, commandRequest.commandRequestPost)
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("%v", err)
	}

	// Calculate CRC by my own and check if it is equal to the got CRC
	raw := buffer.Bytes()
	d := raw[8 : len(raw)-4] // drop first 8 bytes and last 4 bytes and calculate CRC for it
	crc := uint32(crc16.Checksum(d))

	buffer = new(bytes.Buffer)
	_, err = buffer.Write(raw[:len(raw)-4])
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("%v", err)
	}
	err = binary.Write(buffer, binary.BigEndian, crc)
	if err != nil {
		return buffer.Bytes(), fmt.Errorf("%v", err)
	}

	return buffer.Bytes(), nil
}

func DecodeCommandRequest(rawCommand []byte) (CommandRequest, error) {
	var decoded CommandRequest

	reader := bytes.NewReader(rawCommand)

	// Read first part of the record until the dynamic sized section.
	err := binary.Read(reader, binary.BigEndian, &decoded.commandRequestPre)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}

	// Allocate memory for the dynamic sized section. Actual size is defined in the first block.
	decoded.Command = make([]byte, decoded.CommandSize)

	size, err := reader.Read(decoded.Command)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}
	if uint32(size) != decoded.CommandSize {
		return decoded, fmt.Errorf("%d bytes were expected but got %d", decoded.CommandSize, size)
	}

	// Read the remaining part which is again fixed in size.
	err = binary.Read(reader, binary.BigEndian, &decoded.commandRequestPost)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}

	// Calculate CRC by my own and check if it is equal to the got CRC
	d := rawCommand[8 : len(rawCommand)-4] // drop first 8 bytes and last 4 bytes and calculate CRC
	crc := crc16.Checksum(d)

	expected := decoded.CRC
	if uint32(crc) != expected {
		return decoded, fmt.Errorf("CRC check failed! ACTUAL: %x EXPECTED: %x", crc, expected)
	}

	return decoded, nil
}

func DecodeCommandResponse(rawResponse []byte) (CommandResponse, error) {
	var decoded CommandResponse

	reader := bytes.NewReader(rawResponse)

	// Read first part of the record until the dynamic sized section.
	err := binary.Read(reader, binary.BigEndian, &decoded.commandResponsePre)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}

	// Allocate memory for the dynamic sized section. Actual size is defined in the first block.
	decoded.Response = make([]byte, decoded.ResponseSize)

	err = binary.Read(reader, binary.BigEndian, &decoded.Response)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}

	err = binary.Read(reader, binary.BigEndian, &decoded.commandResponsePost)
	if err != nil {
		return decoded, fmt.Errorf("%v", err)
	}

	if decoded.Preamble != ResponsePreamble {
		return decoded, fmt.Errorf("wrong preamble: %v", decoded.Preamble)
	}

	if decoded.CodecID != CodecID {
		return decoded, fmt.Errorf("wrong CodecID: %v", decoded.CodecID)
	}

	if decoded.Type != CommandTypeResponse {
		return decoded, fmt.Errorf("wrong type: %v", decoded.Type)
	}

	// Calculate CRC by my own and check if it is equal to the got CRC
	d := rawResponse[8 : len(rawResponse)-4] // drop first 8 bytes and last 4 bytes and calculate CRC
	calculatedCrc := crc16.Checksum(d)

	expected := decoded.CRC
	if uint32(calculatedCrc) != expected {
		return decoded, fmt.Errorf("wrong CRC! Calculated: %x Received: %x", calculatedCrc, expected)
	}

	return decoded, nil
}
