// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"errors"
	"log"
)

//HAvlData represent human readable set of AVL Data
type HAvlData struct {
	PropertyName string
}

//Human convert all AVL data into human readable format
func Human(el *IOElement) (HAvlData, error) {
	//check if
	if !(el.Length > 0 && el.IOID > 0 && len(el.Value) > 0) {
		log.Fatal("Unable to decode empty element")
		return HAvlData{}, errors.New("Unable to decode empty element")
	}

	return HAvlData{}, nil
}
