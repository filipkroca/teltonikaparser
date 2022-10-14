// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func ExampleDecode() {
	stringData := `01e4cafe0128000f333532303934303839333937343634080400000163c803eb02010a2524c01d4a377d00d3012f130032421b0a4503f00150051503ef01510052005900be00c1000ab50008b60006426fd8cd3d1ece605a5400005500007300005a0000c0000007c70000000df1000059d910002d33c65300000000570000000064000000f7bf000000000000000163c803e6e8010a2530781d4a316f00d40131130031421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005426fcbcd3d1ece605a5400005500007300005a0000c0000007c70000000ef1000059d910002d33b95300000000570000000064000000f7bf000000000000000163c803df18010a2536961d4a2e4f00d50134130033421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542702bcd3d1ece605a5400005500007300005a0000c0000007c70000001ef1000059d910002d33aa5300000000570000000064000000f7bf000000000000000163c8039ce2010a25d8d41d49f42c00dc0123120058421b0a4503f00150051503ef01510052005900be00c1000ab50009b60005427031cd79d8ce605a5400005500007300005a0000c0000007c700000019f1000059d910002d32505300000000570000000064000000f7bf000000000004`

	bs, _ := hex.DecodeString(stringData)
	// decode a raw data byte slice
	parsedData, err := Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}
	fmt.Printf("Decoded packet codec 8:\n%+v\n", parsedData)

	// test with Codec8 Extended packet
	stringData = `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ = hex.DecodeString(stringData)

	// decode a raw data byte slice
	parsedData, err = Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}
	fmt.Printf("Decoded packet codec 8 extended:\n%+v\n", parsedData)

	// Output:
	// Decoded packet codec 8:
	// {IMEI:352094089397464 CodecID:8 NoOfData:4 Data:[{UtimeMs:1528069090050 Utime:1528069090 Priority:1 Lat:491403133 Lng:170206400 Altitude:211 Angle:303 VisSat:19 Speed:50 EventID:66 Elements:[{Length:1 IOID:69 Value:[3]} {Length:1 IOID:240 Value:[1]} {Length:1 IOID:80 Value:[5]} {Length:1 IOID:21 Value:[3]} {Length:1 IOID:239 Value:[1]} {Length:1 IOID:81 Value:[0]} {Length:1 IOID:82 Value:[0]} {Length:1 IOID:89 Value:[0]} {Length:1 IOID:190 Value:[0]} {Length:1 IOID:193 Value:[0]} {Length:2 IOID:181 Value:[0 8]} {Length:2 IOID:182 Value:[0 6]} {Length:2 IOID:66 Value:[111 216]} {Length:2 IOID:205 Value:[61 30]} {Length:2 IOID:206 Value:[96 90]} {Length:2 IOID:84 Value:[0 0]} {Length:2 IOID:85 Value:[0 0]} {Length:2 IOID:115 Value:[0 0]} {Length:2 IOID:90 Value:[0 0]} {Length:2 IOID:192 Value:[0 0]} {Length:4 IOID:199 Value:[0 0 0 13]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 45 51 198]} {Length:4 IOID:83 Value:[0 0 0 0]} {Length:4 IOID:87 Value:[0 0 0 0]} {Length:4 IOID:100 Value:[0 0 0 247]} {Length:4 IOID:191 Value:[0 0 0 0]}]} {UtimeMs:1528069089000 Utime:1528069089 Priority:1 Lat:491401583 Lng:170209400 Altitude:212 Angle:305 VisSat:19 Speed:49 EventID:66 Elements:[{Length:1 IOID:69 Value:[3]} {Length:1 IOID:240 Value:[1]} {Length:1 IOID:80 Value:[5]} {Length:1 IOID:21 Value:[3]} {Length:1 IOID:239 Value:[1]} {Length:1 IOID:81 Value:[0]} {Length:1 IOID:82 Value:[0]} {Length:1 IOID:89 Value:[0]} {Length:1 IOID:190 Value:[0]} {Length:1 IOID:193 Value:[0]} {Length:2 IOID:181 Value:[0 8]} {Length:2 IOID:182 Value:[0 5]} {Length:2 IOID:66 Value:[111 203]} {Length:2 IOID:205 Value:[61 30]} {Length:2 IOID:206 Value:[96 90]} {Length:2 IOID:84 Value:[0 0]} {Length:2 IOID:85 Value:[0 0]} {Length:2 IOID:115 Value:[0 0]} {Length:2 IOID:90 Value:[0 0]} {Length:2 IOID:192 Value:[0 0]} {Length:4 IOID:199 Value:[0 0 0 14]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 45 51 185]} {Length:4 IOID:83 Value:[0 0 0 0]} {Length:4 IOID:87 Value:[0 0 0 0]} {Length:4 IOID:100 Value:[0 0 0 247]} {Length:4 IOID:191 Value:[0 0 0 0]}]} {UtimeMs:1528069087000 Utime:1528069087 Priority:1 Lat:491400783 Lng:170210966 Altitude:213 Angle:308 VisSat:19 Speed:51 EventID:66 Elements:[{Length:1 IOID:69 Value:[3]} {Length:1 IOID:240 Value:[1]} {Length:1 IOID:80 Value:[5]} {Length:1 IOID:21 Value:[3]} {Length:1 IOID:239 Value:[1]} {Length:1 IOID:81 Value:[0]} {Length:1 IOID:82 Value:[0]} {Length:1 IOID:89 Value:[0]} {Length:1 IOID:190 Value:[0]} {Length:1 IOID:193 Value:[0]} {Length:2 IOID:181 Value:[0 8]} {Length:2 IOID:182 Value:[0 5]} {Length:2 IOID:66 Value:[112 43]} {Length:2 IOID:205 Value:[61 30]} {Length:2 IOID:206 Value:[96 90]} {Length:2 IOID:84 Value:[0 0]} {Length:2 IOID:85 Value:[0 0]} {Length:2 IOID:115 Value:[0 0]} {Length:2 IOID:90 Value:[0 0]} {Length:2 IOID:192 Value:[0 0]} {Length:4 IOID:199 Value:[0 0 0 30]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 45 51 170]} {Length:4 IOID:83 Value:[0 0 0 0]} {Length:4 IOID:87 Value:[0 0 0 0]} {Length:4 IOID:100 Value:[0 0 0 247]} {Length:4 IOID:191 Value:[0 0 0 0]}]} {UtimeMs:1528069070050 Utime:1528069070 Priority:1 Lat:491385900 Lng:170252500 Altitude:220 Angle:291 VisSat:18 Speed:88 EventID:66 Elements:[{Length:1 IOID:69 Value:[3]} {Length:1 IOID:240 Value:[1]} {Length:1 IOID:80 Value:[5]} {Length:1 IOID:21 Value:[3]} {Length:1 IOID:239 Value:[1]} {Length:1 IOID:81 Value:[0]} {Length:1 IOID:82 Value:[0]} {Length:1 IOID:89 Value:[0]} {Length:1 IOID:190 Value:[0]} {Length:1 IOID:193 Value:[0]} {Length:2 IOID:181 Value:[0 9]} {Length:2 IOID:182 Value:[0 5]} {Length:2 IOID:66 Value:[112 49]} {Length:2 IOID:205 Value:[121 216]} {Length:2 IOID:206 Value:[96 90]} {Length:2 IOID:84 Value:[0 0]} {Length:2 IOID:85 Value:[0 0]} {Length:2 IOID:115 Value:[0 0]} {Length:2 IOID:90 Value:[0 0]} {Length:2 IOID:192 Value:[0 0]} {Length:4 IOID:199 Value:[0 0 0 25]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 45 50 80]} {Length:4 IOID:83 Value:[0 0 0 0]} {Length:4 IOID:87 Value:[0 0 0 0]} {Length:4 IOID:100 Value:[0 0 0 247]} {Length:4 IOID:191 Value:[0 0 0 0]}]}] Response:[0 5 202 254 1 40 4]}
	//Decoded packet codec 8 extended:
	//{IMEI:352093085698206 CodecID:142 NoOfData:1 Data:[{UtimeMs:1545914096000 Utime:1545914096 Priority:2 Lat:0 Lng:0 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:252 Elements:[{Length:1 IOID:239 Value:[0]} {Length:1 IOID:240 Value:[0]} {Length:1 IOID:21 Value:[5]} {Length:1 IOID:200 Value:[0]} {Length:1 IOID:69 Value:[2]} {Length:1 IOID:1 Value:[0]} {Length:1 IOID:113 Value:[0]} {Length:1 IOID:252 Value:[0]} {Length:2 IOID:181 Value:[0 0]} {Length:2 IOID:182 Value:[0 0]} {Length:2 IOID:66 Value:[48 86]} {Length:2 IOID:205 Value:[67 42]} {Length:2 IOID:206 Value:[96 100]} {Length:2 IOID:17 Value:[0 9]} {Length:2 IOID:18 Value:[255 34]} {Length:2 IOID:19 Value:[3 209]} {Length:2 IOID:15 Value:[0 0]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 0 0 0]}]}] Response:[0 5 202 254 1 1 1]}
}

func ExampleHumanDecoder_Human() {

	// test with Codec8 Extended packet
	stringData := `01e4cafe0126000f333532303934303839333937343634080400000163c803b420010a259e1a1d4a057d00da0128130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427025cd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32c85300000000570000000064000000f7bf000000000000000163c803ac50010a25a9d21d4a01b600db0128130056421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542702ecd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32b05300000000570000000064000000f7bf000000000000000163c803a868010a25b5581d49fe5400db0127130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427039cd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32995300000000570000000064000000f7bf000000000000000163c803a4b2010a25cc861d49f75c00db0124130058421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542703ccd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32695300000000570000000064000000f7bf000000000004`

	/*01e4cafe0126000f333532303934303839333937343634080400000163c803b420010a259e1a1d4a057d00da0128130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427025cd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32c85300000000570000000064000000f7bf000000000000000163c803ac50010a25a9d21d4a01b600db0128130056421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542702ecd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32b05300000000570000000064000000f7bf000000000000000163c803a868010a25b5581d49fe5400db0127130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427039cd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32995300000000570000000064000000f7bf000000000000000163c803a4b2010a25cc861d49f75c00db0124130058421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542703ccd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32695300000000570000000064000000f7bf000000000004*/

	bs, _ := hex.DecodeString(stringData)

	// decode a raw data byte slice
	parsedData, err := Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}

	// initialize a human decoder
	humanDecoder := HumanDecoder{}

	// loop over raw data
	for _, val := range parsedData.Data {
		// loop over Elements
		for _, ioel := range val.Elements {
			// decode to human readable format
			decoded, err := humanDecoder.Human(&ioel, "FM11XY") // second parameter - device family type ["FMBXY", "FM64", "FM36", "FM11XY"]
			if err != nil {
				log.Panicf("Error when converting human, %v\n", err)
			}

			// get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
			if val, err := (*decoded).GetFinalValue(); err != nil {
				log.Panicf("Unable to GetFinalValue() %v", err)
			} else if val != nil {
				// print output
				fmt.Printf("Property Name: %v, Value: %v\n", decoded.AvlEncodeKey.PropertyName, val)
			}
		}
	}

	// Output:
	// Property Name: GPS Status, Value: 3
	// Property Name: Movement Sensor, Value: 1
	// Property Name: Data Mode, Value: 5
	// Property Name: GSM level, Value: 3
	// Property Name: Ignition, Value: 1
	// Property Name: LVCAN Speed, Value: 0
	// Property Name: LVCAN Accelerator Pedal Position, Value: 0
	// Property Name: LVCAN Fuel Level (percentage), Value: 0
	// Property Name: LVC CNG Status, Value: 0
	// Property Name: LVC Oil level, Value: 0
	// Property Name: PDOP, Value: 8
	// Property Name: HDOP, Value: 5
	// Property Name: External Power Voltage, Value: 28709
	// Property Name: Cell ID, Value: 31192
	// Property Name: Area Code, Value: 24666
	// Property Name: LVCAN Fuel Level (liters), Value: 0
	// Property Name: LVCAN Engine RPM, Value: 0
	// Property Name: LVC Engine Temperature, Value: 0
	// Property Name: LVCAN Door Status, Value: 0
	// Property Name: LVC CNG Level, Value: 0
	// Property Name: Odometer/Trip Distance, Value: 24
	// Property Name: GSM Operator Code, Value: 23001
	// Property Name: Total distance, Value: 2962120
	// Property Name: LVCAN Total Fuel Used, Value: 0
	// Property Name: LVCAN Vehicle Distance, Value: 0
	// Property Name: LVCAN Program Number, Value: 247
	// Property Name: LVC CNG Used, Value: 0
	// Property Name: GPS Status, Value: 3
	// Property Name: Movement Sensor, Value: 1
	// Property Name: Data Mode, Value: 5
	// Property Name: GSM level, Value: 3
	// Property Name: Ignition, Value: 1
	// Property Name: LVCAN Speed, Value: 0
	// Property Name: LVCAN Accelerator Pedal Position, Value: 0
	// Property Name: LVCAN Fuel Level (percentage), Value: 0
	// Property Name: LVC CNG Status, Value: 0
	// Property Name: LVC Oil level, Value: 0
	// Property Name: PDOP, Value: 8
	// Property Name: HDOP, Value: 5
	// Property Name: External Power Voltage, Value: 28718
	// Property Name: Cell ID, Value: 31192
	// Property Name: Area Code, Value: 24666
	// Property Name: LVCAN Fuel Level (liters), Value: 0
	// Property Name: LVCAN Engine RPM, Value: 0
	// Property Name: LVC Engine Temperature, Value: 0
	// Property Name: LVCAN Door Status, Value: 0
	// Property Name: LVC CNG Level, Value: 0
	// Property Name: Odometer/Trip Distance, Value: 23
	// Property Name: GSM Operator Code, Value: 23001
	// Property Name: Total distance, Value: 2962096
	// Property Name: LVCAN Total Fuel Used, Value: 0
	// Property Name: LVCAN Vehicle Distance, Value: 0
	// Property Name: LVCAN Program Number, Value: 247
	// Property Name: LVC CNG Used, Value: 0
	// Property Name: GPS Status, Value: 3
	// Property Name: Movement Sensor, Value: 1
	// Property Name: Data Mode, Value: 5
	// Property Name: GSM level, Value: 3
	// Property Name: Ignition, Value: 1
	// Property Name: LVCAN Speed, Value: 0
	// Property Name: LVCAN Accelerator Pedal Position, Value: 0
	// Property Name: LVCAN Fuel Level (percentage), Value: 0
	// Property Name: LVC CNG Status, Value: 0
	// Property Name: LVC Oil level, Value: 0
	// Property Name: PDOP, Value: 8
	// Property Name: HDOP, Value: 5
	// Property Name: External Power Voltage, Value: 28729
	// Property Name: Cell ID, Value: 31192
	// Property Name: Area Code, Value: 24666
	// Property Name: LVCAN Fuel Level (liters), Value: 0
	// Property Name: LVCAN Engine RPM, Value: 0
	// Property Name: LVC Engine Temperature, Value: 0
	// Property Name: LVCAN Door Status, Value: 0
	// Property Name: LVC CNG Level, Value: 0
	// Property Name: Odometer/Trip Distance, Value: 23
	// Property Name: GSM Operator Code, Value: 23001
	// Property Name: Total distance, Value: 2962073
	// Property Name: LVCAN Total Fuel Used, Value: 0
	// Property Name: LVCAN Vehicle Distance, Value: 0
	// Property Name: LVCAN Program Number, Value: 247
	// Property Name: LVC CNG Used, Value: 0
	// Property Name: GPS Status, Value: 3
	// Property Name: Movement Sensor, Value: 1
	// Property Name: Data Mode, Value: 5
	// Property Name: GSM level, Value: 3
	// Property Name: Ignition, Value: 1
	// Property Name: LVCAN Speed, Value: 0
	// Property Name: LVCAN Accelerator Pedal Position, Value: 0
	// Property Name: LVCAN Fuel Level (percentage), Value: 0
	// Property Name: LVC CNG Status, Value: 0
	// Property Name: LVC Oil level, Value: 0
	// Property Name: PDOP, Value: 8
	// Property Name: HDOP, Value: 5
	// Property Name: External Power Voltage, Value: 28732
	// Property Name: Cell ID, Value: 31192
	// Property Name: Area Code, Value: 24666
	// Property Name: LVCAN Fuel Level (liters), Value: 0
	// Property Name: LVCAN Engine RPM, Value: 0
	// Property Name: LVC Engine Temperature, Value: 0
	// Property Name: LVCAN Door Status, Value: 0
	// Property Name: LVC CNG Level, Value: 0
	// Property Name: Odometer/Trip Distance, Value: 24
	// Property Name: GSM Operator Code, Value: 23001
	// Property Name: Total distance, Value: 2962025
	// Property Name: LVCAN Total Fuel Used, Value: 0
	// Property Name: LVCAN Vehicle Distance, Value: 0
	// Property Name: LVCAN Program Number, Value: 247
	// Property Name: LVC CNG Used, Value: 0
}

func ExampleHumanDecoder_AvlDataToHuman() {
	// test with Codec8 Extended packet
	stringData := `01e4cafe0126000f333532303934303839333937343634080400000163c803b420010a259e1a1d4a057d00da0128130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427025cd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32c85300000000570000000064000000f7bf000000000000000163c803ac50010a25a9d21d4a01b600db0128130056421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542702ecd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32b05300000000570000000064000000f7bf000000000000000163c803a868010a25b5581d49fe5400db0127130057421b0a4503f00150051503ef01510052005900be00c1000ab50008b60005427039cd79d8ce605a5400005500007300005a0000c0000007c700000017f1000059d910002d32995300000000570000000064000000f7bf000000000000000163c803a4b2010a25cc861d49f75c00db0124130058421b0a4503f00150051503ef01510052005900be00c1000ab50008b6000542703ccd79d8ce605a5400005500007300005a0000c0000007c700000018f1000059d910002d32695300000000570000000064000000f7bf000000000004`

	bs, _ := hex.DecodeString(stringData)

	// decode a raw data byte slice
	parsedData, err := Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}

	// initialize a human decoder
	humanDecoder := HumanDecoder{}

	decoded, err := humanDecoder.AvlDataToHuman(&parsedData.Data)

	fmt.Printf("Property Name: %v\n", decoded)

}

// Property Name: LVCAN Program Number, Value: 247
// Property Name: LVC CNG Used, Value: 0

func BenchmarkDecode(b *testing.B) {
	stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ := hex.DecodeString(stringData)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := Decode(&bs)
		if err != nil {
			log.Panicf("Error when decoding a bs, %v\n", err)
		}
	}
}

func BenchmarkHuman(b *testing.B) {
	stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ := hex.DecodeString(stringData)
	// initialize a human decoder
	humanDecoder := HumanDecoder{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		// decode a raw data byte slice
		parsedData, err := Decode(&bs)
		if err != nil {
			log.Panicf("Error when decoding a bs, %v\n", err)
		}

		// loop over raw data
		for _, val := range parsedData.Data {
			// loop over Elements
			for _, ioel := range val.Elements {
				// decode to human readable format
				decoded, err := humanDecoder.Human(&ioel, "FMBXY") // second parameter - device family type ["FMBXY", "FM64"]
				if err != nil {
					log.Panicf("Error when converting human, %v\n", err)
				}

				// get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
				if _, err := (*decoded).GetFinalValue(); err != nil {
					log.Panicf("Unable to GetFinalValue() %v", err)
				}
			}
		}
	}
}
