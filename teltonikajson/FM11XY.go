// this file is used to store JSON

package teltonikajson

// FM11XY holds JSON representation of AVL IO elements for devices family FM11XY
const FM11XY string = `{
	"1":{
	   "PropertyName":"Digital Input Status 1",
	   "Bytes":"1",
	   "Description":"Logic: 0 / 1",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"2":{
	   "PropertyName":"Digital Input Status 2",
	   "Bytes":"1",
	   "Description":"Logic: 0 / 1",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"3":{
	   "PropertyName":"Digital Input Status 3",
	   "Bytes":"1",
	   "Description":"Logic: 0 / 1",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"9":{
	   "PropertyName":"Analog Input 1",
	   "Bytes":"2",
	   "Description":"Voltage: mV, 0 – 30 V",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"16":{
	   "PropertyName":"Total distance",
	   "Bytes":"4",
	   "Description":"Total distance: m",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"21":{
	   "PropertyName":"GSM level",
	   "Bytes":"1",
	   "Description":"GSM signal level value in scale 1 – 5",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"24":{
	   "PropertyName":"Speed",
	   "Bytes":"2",
	   "Description":"Value in km/h, 0 – xxx km/h",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"66":{
	   "PropertyName":"External Power Voltage",
	   "Bytes":"2",
	   "Description":"Voltage: mV, 0 – 30 V",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"69":{
	   "PropertyName":"GPS Status",
	   "Bytes":"1",
	   "Description":"States: 0 – GPS module is turned off, 2 – working, but no fix, 3 – working with GPS fix, 4 – GPS module is in sleep state, 5 – antenna is short circuit",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"71":{
	   "PropertyName":"Dallas Temperature ID 4",
	   "Bytes":"8",
	   "Description":"Dallas sensor ID number",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"72":{
	   "PropertyName":"Dallas Temperature 1",
	   "Bytes":"4",
	   "Description":"10 * Degrees ( °C ), -55 - +115, if 3000 – Dallas error",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"73":{
	   "PropertyName":"Dallas Temperature 2",
	   "Bytes":"4",
	   "Description":"10 * Degrees ( °C ), -55 - +115, if 3000 – Dallas error",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"74":{
	   "PropertyName":"Dallas Temperature 3",
	   "Bytes":"4",
	   "Description":"10 * Degrees ( °C ), -55 - +115, if 3000 – Dallas error",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"75":{
	   "PropertyName":"Dallas Temperature 4",
	   "Bytes":"4",
	   "Description":"10 * Degrees ( °C ), -55 - +115, if 3000 – Dallas error",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"76":{
	   "PropertyName":"Dallas Temperature ID 1",
	   "Bytes":"8",
	   "Description":"Dallas sensor ID number",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"77":{
	   "PropertyName":"Dallas Temperature ID 2",
	   "Bytes":"8",
	   "Description":"Dallas sensor ID number",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"78":{
	   "PropertyName":"iButton ID",
	   "Bytes":"8",
	   "Description":"iButton ID number",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"79":{
	   "PropertyName":"Dallas Temperature ID 3",
	   "Bytes":"8",
	   "Description":"Dallas sensor ID number",
	   "Parametr Group":"M",
	   "FinalConversion":"to[]byte"
	},
	"80":{
	   "PropertyName":"Data Mode",
	   "Bytes":"1",
	   "Description":"0 – home on stop, 1 – home on move, 2 – roaming on stop, 3 – roaming on move, 4 – unknown on stop, 5 – unknown on move",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"179":{
	   "PropertyName":"Digital Output 1 state",
	   "Bytes":"1",
	   "Description":"Logic: 0 / 1",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"180":{
	   "PropertyName":"Digital Output 2 state",
	   "Bytes":"1",
	   "Description":"Logic: 0 / 1",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"181":{
	   "PropertyName":"PDOP",
	   "Bytes":"2",
	   "Description":"Probability * 10; 0-500",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"182":{
	   "PropertyName":"HDOP",
	   "Bytes":"2",
	   "Description":"Probability * 10; 0-500",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"199":{
	   "PropertyName":"Odometer/Trip Distance",
	   "Bytes":"4",
	   "Description":"Distance between two records: m",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"200":{
	   "PropertyName":"Deep Sleep",
	   "Bytes":"1",
	   "Description":"0 – not deep sleep mode, 1 – deep sleep mode",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"205":{
	   "PropertyName":"Cell ID",
	   "Bytes":"2",
	   "Description":"GSM base station ID",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"206":{
	   "PropertyName":"Area Code",
	   "Bytes":"2",
	   "Description":"Location Area code (LAC), it depends on",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"239":{
	   "PropertyName":"Ignition",
	   "Bytes":"1",
	   "Description":"0 – ignition off, 1 – ignition on",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"240":{
	   "PropertyName":"Movement Sensor",
	   "Bytes":"1",
	   "Description":"0 – not moving, 1 – moving",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"241":{
	   "PropertyName":"GSM Operator Code",
	   "Bytes":"4",
	   "Description":"Currently used GSM Operator code",
	   "Parametr Group":"M",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"81":{
	   "PropertyName":"LVCAN Speed",
	   "Bytes":"1",
	   "Description":"Value in km/h",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"82":{
	   "PropertyName":"LVCAN Accelerator Pedal Position",
	   "Bytes":"1",
	   "Description":"Value in persentages, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"83":{
	   "PropertyName":"LVCAN Total Fuel Used",
	   "Bytes":"4",
	   "Description":"Value in liters multiplied by 10, L*10",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"84":{
	   "PropertyName":"LVCAN Fuel Level (liters)",
	   "Bytes":"2",
	   "Description":"Value in liters, L",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"85":{
	   "PropertyName":"LVCAN Engine RPM",
	   "Bytes":"2",
	   "Description":"Value in rounds per minute, rpm",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"87":{
	   "PropertyName":"LVCAN Vehicle Distance",
	   "Bytes":"4",
	   "Description":"Value in meters, m",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"89":{
	   "PropertyName":"LVCAN Fuel Level (percentage)",
	   "Bytes":"1",
	   "Description":"Value in percentages, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"90":{
	   "PropertyName":"LVCAN Door Status",
	   "Bytes":"2",
	   "Description":"Door status value: Min – 0, Max – 16128 Door status is represented as bitmask converted to decimal value. Possible values: 0 – all doors closed, 0x100 (256) – front left door is opened, 0x200 (512) – front right door is opened, 0x400 (1024) – rear left door is opened, 0x800 (2048) – rear right door is opened, 0x1000 (4096) – hood is opened, 0x2000 (8192) – trunk is opened, 0x3F00 (16128) – all doors are opened, or combinations of values",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"100":{
	   "PropertyName":"LVCAN Program Number",
	   "Bytes":"4",
	   "Description":"Value: Min – 0, Max – 999",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"101":{
	   "PropertyName":"LVC ModuleID",
	   "Bytes":"8",
	   "Description":"Module ID",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"102":{
	   "PropertyName":"LVC Engine Work Time",
	   "Bytes":"4",
	   "Description":"Engine work time in minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"103":{
	   "PropertyName":"LVC Engine Work Time (counted)",
	   "Bytes":"4",
	   "Description":"Total Engine work time in minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"105":{
	   "PropertyName":"LVC Total Mileage (counted)",
	   "Bytes":"4",
	   "Description":"Total Vehicle Mileage, m",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"107":{
	   "PropertyName":"LVC Fuel Consumed (counted)",
	   "Bytes":"4",
	   "Description":"Total Fuel Consumed,liters * 10",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"110":{
	   "PropertyName":"LVC Fuel Rate",
	   "Bytes":"2",
	   "Description":"Fuel Rata, liters *10",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"111":{
	   "PropertyName":"LVC AdBlue Level (percent)",
	   "Bytes":"1",
	   "Description":"AdBlue, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"112":{
	   "PropertyName":"LVC AdBlue Level (liters)",
	   "Bytes":"2",
	   "Description":"AdBlue level, L",
	   "Parametr Group":"A2",
	   "Type":"Signed",
	   "FinalConversion":"toInt16"
	},
	"114":{
	   "PropertyName":"LVC Engine Load",
	   "Bytes":"1",
	   "Description":"Engine load, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"115":{
	   "PropertyName":"LVC Engine Temperature",
	   "Bytes":"2",
	   "Description":"Engine Temperature, 10 * Degrees ( °C ),",
	   "Parametr Group":"A2",
	   "Type":"Signed",
	   "FinalConversion":"toInt16"
	},
	"118":{
	   "PropertyName":"LVC Axle 1 Load",
	   "Bytes":"2",
	   "Description":"Axle 1 load, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"119":{
	   "PropertyName":"LVC Axle 2 Load",
	   "Bytes":"2",
	   "Description":"Axle 2 load, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"120":{
	   "PropertyName":"LVC Axle 3 Load",
	   "Bytes":"2",
	   "Description":"Axle 3 load, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"121":{
	   "PropertyName":"LVC Axle 4 Load",
	   "Bytes":"2",
	   "Description":"Axle 4 load, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"122":{
	   "PropertyName":"LVC Axle 5 Load",
	   "Bytes":"2",
	   "Description":"Axle 5 load, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"123":{
	   "PropertyName":"LVC Control State Flags",
	   "Bytes":"4",
	   "Description":"Control state flags Byte0 (LSB): 0x01 – STOP 0x02 – Oil pressure / level 0x04 – Coolant liquid temperature / level 0x08 – Handbrake system 0x10 – Battery charging 0x20 – AIRBAG Byte1:0x01 – CHECK ENGINE 0x02 – Lights failure 0x04 – Low tire pressure 0x08 – Wear of brake pads 0x10 – Warning 0x20 – ABS 0x40 – Low Fuel Byte2:0x01 – ESP 0x02 – Glow plug indicator 0x04 – FAP 0x08 – Electronics pressure control 0x10 – Parking lights 0x20 – Dipped headlights 0x40 – Full beam headlights Byte3: 0x40 – Passenger's seat belt 0x80 – Driver's seat belt",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"124":{
	   "PropertyName":"LVC Agricultural Machinery Flags",
	   "Bytes":"8",
	   "Description":"Agricultural machinery flags Byte0 (LSB): 0x01 – Mowing 0x02 – Grain release from hopper 0x04 – First front hydraulic turned on 0x08 – Rear Power Take-Off turned on Byte1: 0x01 – Excessive play under the threshing drum 0x02 – Grain tank is open 0x04 – 100% of Grain tank 0x08 – 70% of Grain tank 0x10 – Drain filter in hydraulic system of drive cylinders is plugged 0x20 – Pressure filter of drive cylinders hydraulic system is plugged 0x40 – Alarm oil level in oil tank 0x80 – Pressure filter of brakes hydraulic system is plugged Byte2: 0x01 – Oil filter of engine is plugged 0x02 – Fuel filter is plugged 0x04 – Air filter is plugged 0x08 – Alarm oil temperature in hydraulic system of chasis 0x10 – Alarm oil temperature in hydraulic system of drive cylinders 0x20 – Alarm oil pressure in engine 0x40 – Alarm coolant level 0x80 – Overflow chamber of hydraulic unit Byte3: 0x01 – Unloader drive is ON. Unloading tube pivot is in idle position 0x02 – No operator! 0x04 – Straw walker is plugged 0x08 – Water in fuel 0x10 – Cleaning fan RPM 0x20 – Trashing drum RPM Byte4:0x02 – Low water level in the tank 0x04 – First rear hydraulic turned on 0x08 – Standalone engine working 0x10 – Right joystick moved right 0x20 – Right joystick moved left 0x40 – Right joystick moved front 0x80 – Right joystick moved back Byte5: 0x01 – Brushes turned on 0x02 – Water supply turned on 0x04 – Vacuum cleaner  0x08 – Unloading from the hopper 0x10 – High Pressure washer (Karcher) 0x20 – Salt (sand) disperser ON 0x40 – Low salt (sand) level Byte6: 0x01 – Second front hydraulic turned on 0x02 – Third front hydraulic turned on 0x04 – Fourth front hydraulic turned on 0x08 – Second rear hydraulic turned on 0x10 – Third rear hydraulic turned on 0x20 – Fourth rear hydraulic turned on 0x40 – Front three-point Hitch turned on 0x80 – Rear three-point Hitch turned on Byte7:0x01 – Left joystick moved right 0x02 – Left joystick moved left 0x04 – Left joystick moved front 0x08 – Left joystick moved back 0x10 – Front Power Take-Off turned on",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"125":{
	   "PropertyName":"LVC Harvesting Time",
	   "Bytes":"4",
	   "Description":"",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"126":{
	   "PropertyName":"LVC Area of Harvest",
	   "Bytes":"4",
	   "Description":"Area of Harvest, m^2",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"127":{
	   "PropertyName":"LVC Mowing Efficiency",
	   "Bytes":"4",
	   "Description":"Mowing efficiency, (m^2)/h",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"128":{
	   "PropertyName":"LVC Grain Mown Volume",
	   "Bytes":"4",
	   "Description":"Mown Volume, kg",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"129":{
	   "PropertyName":"LVC Grain Moisture",
	   "Bytes":"2",
	   "Description":"Grain Moisture in proc, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"130":{
	   "PropertyName":"LVC Harvesting Drum RPM",
	   "Bytes":"2",
	   "Description":"Harvesting Drum RPM, RPM",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"131":{
	   "PropertyName":"LVC Gap Under Harvesting Drum",
	   "Bytes":"1",
	   "Description":"Gap Under Harvesting Drum, mm",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"132":{
	   "PropertyName":"LVC Security State Flags",
	   "Bytes":"8",
	   "Description":"Security State Flag Byte0 (LSB): Every two bits in this byte correspond to a different CAN bus number. 00 – CAN not connected, connection not required 01 – CAN connected, but currently module not received data 10 – CAN not connected, require connection 11 – CAN connectedExample: Byte0 - 0F hex – 00001111 binary CAN4, CAN3, CAN2, CAN1 Byte1: Not used Byte2: 0x20 – bit appears when any operate button in car was put 0x40 – bit appears when immobilizer is in service mode 0x80 – immobiliser, bit appears during introduction of a programmed sequence of keys in the car. Byte3: 0x01 – the key is in ignition lock 0x02 – ignition on 0x04 – dynamic ignition on 0x08 – webasto 0x20 – car closed by factory's remote control 0x40 – factory-installed alarm system is actuated (is in panic mode) 0x80 – factory-installed alarm system is emulated by module Byte4: 0x01 – parking activated (automatic gearbox) 0x10 – handbrake is actuated (information available only with ignition on) 0x20 – footbrake is actuated (information available only with ignition on) 0x40 – engine is working (information available only when the ignition on) 0x80 – revers is on Byte5: 0x01 – Front left door opened 0x02 – Front right door opened 0x04 – Rear left door opened 0x08 – Rear right door opened 0x10 – engine cover opened 0x20 – trunk door opened Byte6: 0x01 – car was closed by the factory's remote control 0x02 – car was opened by the factory's remote control 0x03 – trunk cover was opened by the factory's remote control 0x04 – module has sent a rearming signal 0x05 – car was closed three times by the factory's remote control - High nibble (mask 0xF0 bit) 0x80 – CAN module goes to sleep mode Byte7: Not used",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"133":{
	   "PropertyName":"LVC Tacho Total Vehicle Distance",
	   "Bytes":"4",
	   "Description":"Tacho Total Vehicle Distance, m",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"134":{
	   "PropertyName":"LVC Trip Distance",
	   "Bytes":"4",
	   "Description":"Trip Distance, m",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"135":{
	   "PropertyName":"LVC Tacho Vehicle Speed",
	   "Bytes":"2",
	   "Description":"Tacho Vehicle Speed, km/h",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"136":{
	   "PropertyName":"LVC Tacho Driver Card Presence",
	   "Bytes":"1",
	   "Description":"Tacho Driver Card Presence 0x00 – No driver card 0x01 – Driver1 card presence 0x02 – Driver2 card presence 0x03 – Driver1 and driver2 cards present",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"137":{
	   "PropertyName":"LVC Driver1 States",
	   "Bytes":"1",
	   "Description":"Driver1 States 0xX0 – break/rest 0xX1 – availability 0xX2 – work 0xX3 – driving 0x0X – no time-related warning detected 0x1X – limit #1: 15 min before 4 1/2 h 0x2X – limit #2: 4 1/2 h reached (continuous driving time exceeded) 0x3X – limit #3: 15 minutes before optional warning 1 0x4X – limit #4: optional warning 1 reached 0x5X – limit #5: 15 min before optional warning 0x6X – limit #6: optional warning 2 reached",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"138":{
	   "PropertyName":"LVC Driver2 States",
	   "Bytes":"1",
	   "Description":"Driver2 States 0xX0 – break/rest 0xX1 – availability 0xX2 – work 0xX3 – driving 0x0X – no time-related warning detected 0x1X – limit #1: 15 min before 4 1/2 h 0x2X – limit #2: 4 1/2 h reached (continuous driving time exceeded) 0x3X – limit #3: 15 minutes before optional warning 1 0x4X – limit #4: optional warning 1 reached 0x5X – limit #5: 15 min before optional warning 0x6X – limit #6: optional warning 2 reached",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"139":{
	   "PropertyName":"LVC Driver1 Continuous Driving Time",
	   "Bytes":"2",
	   "Description":"Driver1 Continuous Driving Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"140":{
	   "PropertyName":"LVC Driver2 Continuous Driving Time",
	   "Bytes":"2",
	   "Description":"Driver2 Continuous Driving Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"141":{
	   "PropertyName":"LVC Driver1 Cumulative Break Time",
	   "Bytes":"2",
	   "Description":"Driver1 Cumulative Break Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"142":{
	   "PropertyName":"LVC Driver2 Cumulative",
	   "Bytes":"2",
	   "Description":"Driver2 Cumulative Break Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"143":{
	   "PropertyName":"LVC Driver1 Duration Of Selected Activity",
	   "Bytes":"2",
	   "Description":"Driver1 Duration Of Selected Activity, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"144":{
	   "PropertyName":"LVC Driver2 Duration Of Selected Activity",
	   "Bytes":"2",
	   "Description":"Driver2 Duration Of Selected Activity, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"145":{
	   "PropertyName":"LVC Driver1 Cumulative Driving Time",
	   "Bytes":"2",
	   "Description":"Driver1 Cumulative Driving Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"146":{
	   "PropertyName":"LVC Driver2 Cumulative Driving Time",
	   "Bytes":"2",
	   "Description":"Driver2 Cumulative Driving Time, minutes",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"147":{
	   "PropertyName":"LVC Driver1 ID High",
	   "Bytes":"8",
	   "Description":"Driver1 ID High",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"148":{
	   "PropertyName":"LVC Driver1 ID Low",
	   "Bytes":"8",
	   "Description":"Driver1 ID Low",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"149":{
	   "PropertyName":"LVC Driver2 ID High",
	   "Bytes":"8",
	   "Description":"Driver2 ID High",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"150":{
	   "PropertyName":"LVC Driver2 ID Low",
	   "Bytes":"8",
	   "Description":"Driver2 ID Low",
	   "Parametr Group":"A2",
	   "FinalConversion":"to[]byte"
	},
	"151":{
	   "PropertyName":"LVC Battery Temperature",
	   "Bytes":"2",
	   "Description":"10* Degrees, ( °C )",
	   "Parametr Group":"A2",
	   "Type":"Signed",
	   "FinalConversion":"toInt16"
	},
	"152":{
	   "PropertyName":"LVC Battery Level (percent)",
	   "Bytes":"1",
	   "Description":"Value in percentages, %",
	   "Parametr Group":"A2",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"160":{
	   "PropertyName":"LVC DTC Errors",
	   "Bytes":"1",
	   "Description":"DTC faults count",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"161":{
	   "PropertyName":"LVC Slope Of Arm",
	   "Bytes":"2",
	   "Description":"Value in o",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"162":{
	   "PropertyName":"LVC Rotation Of Arm",
	   "Bytes":"2",
	   "Description":"Value in o",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"163":{
	   "PropertyName":"LVC Eject Of Arm",
	   "Bytes":"2",
	   "Description":"Value in m * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"164":{
	   "PropertyName":"LVC Horizontal Distance Arm Vechicle",
	   "Bytes":"2",
	   "Description":"Value in m * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"165":{
	   "PropertyName":"LVC Height Arm Above Ground",
	   "Bytes":"2",
	   "Description":"Value in m * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"166":{
	   "PropertyName":"LVC Drill RPM",
	   "Bytes":"2",
	   "Description":"-",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"167":{
	   "PropertyName":"LVC Amount Of Spread Salt Square Meter",
	   "Bytes":"2",
	   "Description":"Value in g/m2",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"168":{
	   "PropertyName":"LVC Battery Voltage",
	   "Bytes":"2",
	   "Description":"Value in V * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"169":{
	   "PropertyName":"LVC Amount Spread Fine Grained Salt",
	   "Bytes":"4",
	   "Description":"Value in tons * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"170":{
	   "PropertyName":"LVC Amount Spread Coarse Grained Salt",
	   "Bytes":"4",
	   "Description":"Value in tons * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"171":{
	   "PropertyName":"LVC Amount Spread DiMix",
	   "Bytes":"4",
	   "Description":"Value in tons * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"172":{
	   "PropertyName":"LVC Amount Spread Coarse Grained Calcium",
	   "Bytes":"4",
	   "Description":"Value in m3 * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"173":{
	   "PropertyName":"LVC Amount Spread Calcium Chloride",
	   "Bytes":"4",
	   "Description":"Value in m3 * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"174":{
	   "PropertyName":"LVC Amount Spread Sodium Chloride",
	   "Bytes":"4",
	   "Description":"Value in m3 * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"176":{
	   "PropertyName":"LVC Amount Spread Magnesium Chloride",
	   "Bytes":"4",
	   "Description":"Value in m3 * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"177":{
	   "PropertyName":"LVC Amount Spread Gravel",
	   "Bytes":"4",
	   "Description":"Value in tons * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"178":{
	   "PropertyName":"LVC Amount Spread Sand",
	   "Bytes":"4",
	   "Description":"Value in tons * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"183":{
	   "PropertyName":"LVC Width Pouring Left",
	   "Bytes":"2",
	   "Description":"Value in m * 100",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"184":{
	   "PropertyName":"LVC Width Pouring Right",
	   "Bytes":"2",
	   "Description":"Value in m * 100",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"185":{
	   "PropertyName":"LVC Salt Spreader Work",
	   "Bytes":"4",
	   "Description":"Value in h * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"186":{
	   "PropertyName":"LVC Distance During Salting",
	   "Bytes":"4",
	   "Description":"Value in km * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"187":{
	   "PropertyName":"LVC Load Weight",
	   "Bytes":"4",
	   "Description":"Value in kg",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"188":{
	   "PropertyName":"LVC Retarder Load",
	   "Bytes":"1",
	   "Description":"Value in % Valid range: 0 – 125%",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"189":{
	   "PropertyName":"LVC Cruise Time",
	   "Bytes":"4",
	   "Description":"Value in min",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"190":{
	   "PropertyName":"LVC CNG Status",
	   "Bytes":"1",
	   "Description":"0 – engine not on CNG 1 – engine on CNG",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"191":{
	   "PropertyName":"LVC CNG Used",
	   "Bytes":"4",
	   "Description":"Value in kg * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint32"
	},
	"192":{
	   "PropertyName":"LVC CNG Level",
	   "Bytes":"2",
	   "Description":"Value in % * 10",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint16"
	},
	"193":{
	   "PropertyName":"LVC Oil level",
	   "Bytes":"1",
	   "Description":"0 – Oil level/pressure warning off 1 – Oil level/pressure warning on",
	   "Parametr Group":"O",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"155":{
	   "PropertyName":"Geofence zone 01",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"156":{
	   "PropertyName":"Geofence zone 02",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"157":{
	   "PropertyName":"Geofence zone 03",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"158":{
	   "PropertyName":"Geofence zone 04",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"159":{
	   "PropertyName":"Geofence zone 05",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"175":{
	   "PropertyName":"Auto Geofence",
	   "Bytes":"1",
	   "Description":"Event: 0 – target left zone, 1 – target entered zone",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"249":{
	   "PropertyName":"Jamming",
	   "Bytes":"1",
	   "Description":"1 – jamming start, 0 – jamming stop",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"250":{
	   "PropertyName":"Trip",
	   "Bytes":"1",
	   "Description":"1 – trip start, 0 – trip stop",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"251":{
	   "PropertyName":"Immobilizer",
	   "Bytes":"1",
	   "Description":"1 – iButton connected",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"252":{
	   "PropertyName":"Authorized driving",
	   "Bytes":"1",
	   "Description":"1 – authorized iButton connected",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"253":{
	   "PropertyName":"Green driving type",
	   "Bytes":"1",
	   "Description":"1 – harsh acceleration, 2 – harsh braking, 3 – harsh cornering",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"254":{
	   "PropertyName":"Green driving value",
	   "Bytes":"1",
	   "Description":"Depending on green driving type: if harsh acceleration or braking – g*100 (value 123 -> 1.23g), if harsh cornering – degrees (value in radians)",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	},
	"255":{
	   "PropertyName":"Over Speeding",
	   "Bytes":"1",
	   "Description":"At over speeding start km/h, at over speeding end km/h",
	   "Parametr Group":"ME",
	   "Type":"Unsigned",
	   "FinalConversion":"toUint8"
	}
 }`
