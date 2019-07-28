// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	teltonikaparser "github.com/filipkroca/teltonikaparser/decoding"
)

//HAvlData represent human readable set of AVL Data
type HAvlData struct {
	PropertyName string
}

//HumanDecoder is responsible for decoding to human-readable format
type HumanDecoder struct {
	avlKey map[uint16]AvlIO //this should be inicialized with New method
}

//AvlIO represent parsed element values from JSON
type AvlIO struct {
	No            string `json:"No"`
	PropertyName  string `json:"PropertyName"`
	Bytes         string `json:"Bytes"`
	Type          string `json:"Type"`
	Min           string `json:"Min"`
	Max           string `json:"Max"`
	Multiplier    string `json:"Multiplier"`
	Units         string `json:"Units"`
	Description   string `json:"Description"`
	HWSupport     string `json:"HWSupport"`
	ParametrGroup string `json:"Parametr Group"`
}

//New parses ./decoding/FmbAVLID.json and return HumanDecoder struct with mapped AVL Keys
func (h *HumanDecoder) New() HumanDecoder {

	new := HumanDecoder{}
	// read our opened xmlFile as a byte array.
	byteValue := []byte(teltonikaparser.AVLJson)

	err := json.Unmarshal(byteValue, &new.avlKey)
	if err != nil {
		log.Panic(err)
	}

	return new
}

//Human return a pointer to decoding key
func (h *HumanDecoder) Human(el *IOElement) (*AvlIO, error) {
	//check if human is inicialized
	if len(h.avlKey) == 0 {
		log.Fatal("You should create a HumanDecoder by New method on it because it is necessary to initialize decoder Map")
		return nil, errors.New("You should create a Human by New method on it because it is necessary to initialize decoder Map")
	}

	//check if
	if !((*el).Length > 0 && (*el).IOID > 0 && len((*el).Value) > 0) {
		log.Fatal("Unable to decode empty element")
		return nil, errors.New("Unable to decode empty element")
	}

	//find decode key
	i, ok := h.avlKey[(*el).IOID]
	if !ok {
		log.Fatalf("Unknown element %v", (*el).IOID)
		return nil, fmt.Errorf("Unknown element %v", (*el).IOID)
	}

	return &i, nil
}
