// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/filipkroca/b2n"
	"github.com/filipkroca/teltonikaparser/teltonikajson"
)

// HAvlData represent human readable set of a pointer to an AvlIO Decoding key and a pointer to IO element with RAW data
type HAvlData struct {
	AvlIO     *AvlIO
	IOElement *IOElement
}

// HumanDecoder is responsible for decoding
type HumanDecoder struct {
	elements map[string]map[uint16]AvlIO
}

// AvlIO represent parsed element values from JSON
type AvlIO struct {
	No              string `json:"No"`
	PropertyName    string `json:"PropertyName"`
	Bytes           string `json:"Bytes"`
	Type            string `json:"Type"`
	Min             string `json:"Min"`
	Max             string `json:"Max"`
	Multiplier      string `json:"Multiplier"`
	Units           string `json:"Units"`
	Description     string `json:"Description"`
	HWSupport       string `json:"HWSupport"`
	ParametrGroup   string `json:"Parametr Group"`
	FinalConversion string `json:"FinalConversion"`
}

// Human takes a pointer to IOElement, device type ["FMBXY", "FM64"] and return a pointer to decoding key
func (h *HumanDecoder) Human(el *IOElement, device string) (*HAvlData, error) {
	//init decoding key
	if len(h.elements) == 0 {
		h.loadElements()
	}

	// check if IOElement is valid
	if !((*el).Length > 0 && (*el).IOID > 0 && len((*el).Value) > 0) {
		log.Fatal("Unable to decode empty element")
		return nil, errors.New("Unable to decode empty element")
	}

	// find decode key and pair it
	avl, ok := h.elements["FMBXY"][(*el).IOID]
	if !ok {
		log.Fatalf("Unknown element %v", (*el).IOID)
		return nil, fmt.Errorf("Unknown element %v", (*el).IOID)
	}

	// return pointer to merged struct with decode key AvlIO and data IOElement
	havl := HAvlData{
		AvlIO:     &avl,
		IOElement: el,
	}
	return &havl, nil
}

// loadElements parses ./decoding/.. into slice
func (h *HumanDecoder) loadElements() {
	// make map
	h.elements = make(map[string]map[uint16]AvlIO)

	// read our opened xmlFile as a byte array.
	byteValue := []byte(teltonikajson.FMBXY)
	fmbxy := make(map[uint16]AvlIO)
	//h.elements["FMBXY"] = make(map[uint16]AvlIO)
	err := json.Unmarshal(byteValue, &fmbxy)
	if err != nil {
		log.Panic(err)
	}
	h.elements["FMBXY"] = fmbxy

	// read our opened xmlFile as a byte array.
	byteValue = []byte(teltonikajson.FM64)
	fm64 := make(map[uint16]AvlIO)
	//h.elements["FMBXY"] = make(map[uint16]AvlIO)
	err = json.Unmarshal(byteValue, &fm64)
	if err != nil {
		log.Panic(err)
	}
	h.elements["FM64"] = fm64

}

// GetFinalValue return decimal value, if necesarry with float, return should be empty interface because there is many values to return
func (h *HAvlData) GetFinalValue() (interface{}, error) {

	if h.AvlIO.FinalConversion == "toUint8" {
		if h.AvlIO.Bytes != "1" || h.AvlIO.Type != "Unsigned" || len(h.IOElement.Value) != 1 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Uint8 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Uint8(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toUint16" {
		if h.AvlIO.Bytes != "2" || h.AvlIO.Type != "Unsigned" || len(h.IOElement.Value) != 2 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Uint16 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Uint16(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toUint32" {
		if h.AvlIO.Bytes != "4" || h.AvlIO.Type != "Unsigned" || len(h.IOElement.Value) != 4 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Uint32 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Uint32(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toUint64" {
		if h.AvlIO.Bytes != "8" || h.AvlIO.Type != "Unsigned" || len(h.IOElement.Value) != 8 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Uint64 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Uint64(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toInt8" {
		if h.AvlIO.Bytes != "1" || h.AvlIO.Type != "Signed" || len(h.IOElement.Value) != 1 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Int8 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Int8TwoComplement(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toInt16" {
		if h.AvlIO.Bytes != "2" || h.AvlIO.Type != "Signed" || len(h.IOElement.Value) != 2 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Int16 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Int16TwoComplement(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toInt32" {
		if h.AvlIO.Bytes != "4" || h.AvlIO.Type != "Signed" || len(h.IOElement.Value) != 4 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr to Int32 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Int32TwoComplement(&h.IOElement.Value, 0), nil
	}

	if h.AvlIO.FinalConversion == "toInt64" {
		if h.AvlIO.Bytes != "8" || h.AvlIO.Type != "Signed" || len(h.IOElement.Value) != 8 {
			return nil, fmt.Errorf("Unable to convert %vBytes long parametr, %vBytes real long parametr, to Int64 %v", h.AvlIO.Bytes, len(h.IOElement.Value), h.AvlIO.PropertyName)
		}
		return b2n.ParseBs2Int64TwoComplement(&h.IOElement.Value, 0), nil
	}

	return nil, nil
}
