// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package decoder

/*
// doTableTest is the universal decoder test sequence.
func doTableTest(t *testing.T, f newDecoder, endianness bool, teTa testTable) {
	lu := make(id.TriceIDLookUp)
	luM := new(sync.RWMutex)
	assert.Nil(t, lu.FromJSON([]byte(til)))
	lu.AddFmtCount(os.Stdout)
	buf := make([]byte, decoder.DefaultSize)
	dec := f(lu, luM, nil, endianness) // p is a new decoder instance
	for _, x := range teTa {
		in := ioutil.NopCloser(bytes.NewBuffer(x.in))
		dec.SetInput(in)
		lineStart := true
		var err error
		var n int
		var act string
		for nil == err {
			n, err = dec.Read(buf)
			if io.EOF == err && n == 0 {
				break
			}
			if "" != ShowID && lineStart {
				act += fmt.Sprintf(ShowID, decoder.LastTriceID)
			}
			// a := fmt.Sprint(string(buf[:n]))
			// if emitter.SyncPacketPattern != a { // todo: Handle ShowID in that case.
			// 	act += a // ignore sync packets
			// }
			lineStart = false
		}
		act = strings.TrimSuffix(act, "\\n")
		act = strings.TrimSuffix(act, "\n")
		assert.Equal(t, x.exp, act)
	}
}

var (
	// til is the trace id list content for tests
	til = `{
		"10052": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1 %d\\n"
		},
		"10333": {
			"Type": "Trice8i",
			"Strg": "rd:%d, %d\\n"
		},
		"10363": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1  -\u003e short  trice macro    (no   cycle counter) for everywhere %u\\n"
		},
		"10407": {
			"Type": "trice8_1",
			"Strg": "tst:trice8_1 %02X\\n"
		},
		"1047663": {
			"Type": "TRICE16_2",
			"Strg": "MSG: triceFifoMaxDepth = %d, select = %d\\n"
		},
		"10509": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_HAL_UART_NUCLEO-G474\\n"
		},
		"10524": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1 -\u003e short trice macro (no cycle counter) for everywhere 0x%X\\n"
		},
		"10650": {
			"Type": "trice8_4",
			"Strg": "tst:trice8_4 %02X %02X %02X %02X\\n"
		},
		"10732": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:   MDK-ARM_LL_UART_RTT0_BARE_STM32F070_NUCLEO-64   \\ns:                                                   \\n\\n"
		},
		"10761": {
			"Type": "TRICE16_1",
			"Strg": "WRN:warning     message, SysTick is %6u\\n"
		},
		"10813": {
			"Type": "TRICE16_1",
			"Strg": "time:TRICE16_1   message, SysTick is %6u\\n"
		},
		"11439": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1 %6u\\n"
		},
		"11804": {
			"Type": "TRICE0",
			"Strg": "Invalid wav file\\n"
		},
		"11862": {
			"Type": "Trice0",
			"Strg": "att:SysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_Handler\\n"
		},
		"11967": {
			"Type": "TRICE16_1",
			"Strg": "att:TRICE16_1   message, SysTick is %6u\\n"
		},
		"12051": {
			"Type": "TRICE64_2",
			"Strg": "tst:TRICE64_2 %u %u\\n"
		},
		"12093": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTTD_NUCLEO-F030R8\\n"
		},
		"12126": {
			"Type": "TRICE0",
			"Strg": "att:TASKING_GenericSTMF030R8_RTTD\\n"
		},
		"12345": {
			"Type": "Trice16_1",
			"Strg": "%b"
		},
		"12478": {
			"Type": "TRICE16_4",
			"Strg": "tst:TRICE16_4 %u %u %u %u\\n"
		},
		"12497": {
			"Type": "Trice8_1i",
			"Strg": "msg: Trice8_1i -\u003e short  trice macro    (no   cycle counter) for only inside critical section %d\\n"
		},
		"12927": {
			"Type": "Trice8_2",
			"Strg": "msg: Trice8_2  %x, %x\\n"
		},
		"13003": {
			"Type": "Trice0i",
			"Strg": "msg: Trice0i   short\\n"
		},
		"13083": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_RTTD_STM32F0300-DISCO\\n"
		},
		"13256": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i -\u003e short trice macro (no cycle counter) for only inside critical section %x\\n"
		},
		"13390": {
			"Type": "TRICE16_1",
			"Strg": "INFO:informal   message, SysTick is %6u\\n"
		},
		"13520": {
			"Type": "Trice16_1",
			"Strg": "rd: Trice16_1 %u\\n"
		},
		"13551": {
			"Type": "Trice16_1i",
			"Strg": "rd: Trice16_1i %u\\n"
		},
		"13668": {
			"Type": "Trice16_1i",
			"Strg": "int:SysTick: %5u\\n"
		},
		"13685": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_LL_UART_NUCLEO-F030R8\\n"
		},
		"13704": {
			"Type": "Trice16_1",
			"Strg": "int:Trice16_1 SysTick: %5u\\n"
		},
		"13850": {
			"Type": "Trice8_2",
			"Strg": "INT:%d %d\\n"
		},
		"13913": {
			"Type": "Trice0i",
			"Strg": "int:SysTick_Handler"
		},
		"14014": {
			"Type": "TRICE8_3",
			"Strg": "tst:TRICE8_3 %u %u %u\\n"
		},
		"14382": {
			"Type": "TRICE0",
			"Strg": "Successfully initialized audio service\\n"
		},
		"14522": {
			"Type": "TRICE16_1",
			"Strg": "tim:timing      message, SysTick is %6d\\n"
		},
		"14969": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_LL_UART_NUCLEO-F070RB\\n"
		},
		"15330": {
			"Type": "Trice0i",
			"Strg": "msg: Trice0i -\u003e short  trice macro    (no   cycle counter) for only inside critical section\\n"
		},
		"15380": {
			"Type": "Trice0i",
			"Strg": "msg: Trice0i    -\u003e short  trice macro    (no   cycle counter) for only inside critical section\\n"
		},
		"15428": {
			"Type": "Trice8_2i",
			"Strg": "int:%d %d\\n"
		},
		"15516": {
			"Type": "TRICE16_1",
			"Strg": "TIME:TRICE16_1   message, SysTick is %6u\\n"
		},
		"15852": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_RTTD_NUCLEO-F091\\n"
		},
		"16383": {
			"Type": "Trice8_2i",
			"Strg": "msg: Trice8_2i  %x, %x\\n"
		},
		"16417": {
			"Type": "TRICE32_1",
			"Strg": "ISR:alive time %d milliseconds\\n"
		},
		"16731": {
			"Type": "TRICE8_8",
			"Strg": "tst:TRICE8_8 %02X %02X %02X %02X %02X %02X %02X %02X\\n"
		},
		"16793": {
			"Type": "Trice0i",
			"Strg": "INT:Trice0i SysTick_Handler\\n"
		},
		"17147": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_UART_NUCLEO-F030R8\\n"
		},
		"17181": {
			"Type": "Trice8i",
			"Strg": "rd:Trice8i line %d\\n"
		},
		"1719": {
			"Type": "TRICE16_1",
			"Strg": "tim  :TRICE16_1   message, SysTick is %6u\\n"
		},
		"17536": {
			"Type": "TRICE16_1",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"176": {
			"Type": "trice0",
			"Strg": "msg:hi"
		},
		"17896": {
			"Type": "TRICE0",
			"Strg": "s:                                                          \\ns:   MDK-ARM_LL_UART_WRAP_RTT0_BARE_STM32F030R8-NUCLEO-64   \\ns:                                                          \\n\\n"
		},
		"18398": {
			"Type": "TRICE0",
			"Strg": "Complex number c1: "
		},
		"18561": {
			"Type": "TRICE0",
			"Strg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_TO_ESC_NUCLEO-F070RB   \\ns:                                        \\n\\n"
		},
		"19001": {
			"Type": "Trice0i",
			"Strg": "wr: Trice0i   short\\n"
		},
		"19074": {
			"Type": "Trice8_1i",
			"Strg": "INT:SysTick: %5d\\n"
		},
		"19305": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1 -\u003e short trice macro (no cycle counter) for everywhere 0x%X\\n"
		},
		"19307": {
			"Type": "Trice8_1",
			"Strg": "msg: Trice8_1 -\u003e short trice macro (no cycle counter) for everywhere %u\\n"
		},
		"20096": {
			"Type": "TRICE8_2",
			"Strg": "tst:TRICE8_2 %02x %02x\\n"
		},
		"20101": {
			"Type": "Trice16_1",
			"Strg": "rd: Trice16_1 %d\\n"
		},
		"20104": {
			"Type": "Trice16i",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"20419": {
			"Type": "Trice16_1i",
			"Strg": "int:Trice16_1i SysTick: %5u\\n"
		},
		"20663": {
			"Type": "Trice16_1i",
			"Strg": "time: Trice16_1i %d\\n"
		},
		"20679": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i -\u003e short trice macro (no cycle counter) for only inside critical section\\n"
		},
		"21036": {
			"Type": "Trice0",
			"Strg": "int:SysTick_Handler"
		},
		"21124": {
			"Type": "TRICE8_5",
			"Strg": "tst:TRICE8_5 %02X %02X %02X %02X %02X\\n"
		},
		"21201": {
			"Type": "TRICE8_1",
			"Strg": "tst:TRICE8_1 0x%02x\\n"
		},
		"21324": {
			"Type": "Trice8i",
			"Strg": "rd:Trice8i line %d, %d\\n"
		},
		"21394": {
			"Type": "TRICE16_1",
			"Strg": "tst:TRICE16_1 0x%04x\\n"
		},
		"21430": {
			"Type": "trice16_1",
			"Strg": "tst:trice16_1   message, SysTick is %6u\\n"
		},
		"21678": {
			"Type": "TRICE8_6",
			"Strg": "tst:TRICE8_6 %02x %02x %02x %02x %02x %02x\\n"
		},
		"21682": {
			"Type": "Trice0",
			"Strg": "msg: Trice0    short\\n"
		},
		"21880": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i -\u003e short trice macro (no cycle counter) for only inside critical section %x\\n"
		},
		"22308": {
			"Type": "TRICE8_7",
			"Strg": "tst:TRICE8_7 %02x %02x %02x %02x %02x %02x %02x\\n"
		},
		"22722": {
			"Type": "TRICE0",
			"Strg": "s:                                              \\ns:    ARM-MDK_BARE_STM32F03051R8Tx-DISCOVERY    \\ns:                                              \\n\\n"
		},
		"22909": {
			"Type": "Trice16_1i",
			"Strg": "time: Trice16_1i %6u\\n"
		},
		"23045": {
			"Type": "Trice16_1i",
			"Strg": "rd: Trice16_1i %u\\n"
		},
		"23242": {
			"Type": "Trice0",
			"Strg": "msg: Trice0     -\u003e short  trice macro    (no   cycle counter) for everywhere\\n"
		},
		"23553": {
			"Type": "TRICE0",
			"Strg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_TO_ESC_NUCLEO-F030R8   \\ns:                                        \\n\\n"
		},
		"24194": {
			"Type": "TRICE16_1",
			"Strg": "SIG:signal      message, SysTick is %6u\\n"
		},
		"24267": {
			"Type": "TRICE32_1",
			"Strg": "tst:TRICE32_1 %u\\n"
		},
		"24326": {
			"Type": "TRICE0",
			"Strg": "att:TASKING_RTTD_cpp_Example\\n"
		},
		"2457": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1  -\u003e short  trice macro    (no   cycle counter) for everywhere                   %u\\n"
		},
		"24615": {
			"Type": "TRICE8_6",
			"Strg": "tst:TRICE8_6 %u %u %u %u %u %u\\n"
		},
		"24616": {
			"Type": "Trice8",
			"Strg": "rd:Trice8 line %d, %d\\n"
		},
		"24626": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTT_NUCLEO-F030R8\\n"
		},
		"24926": {
			"Type": "Trice8_2",
			"Strg": "int:%d %d\\n"
		},
		"25013": {
			"Type": "Trice8_1i",
			"Strg": "msg: Trice8_1i  -\u003e short  trice macro    (no   cycle counter) for only inside critical section %d\\n"
		},
		"25023": {
			"Type": "Trice8_1",
			"Strg": "msg: Trice8_1 -\u003e short trice macro (no cycle counter) for everywhere\\n"
		},
		"25308": {
			"Type": "Trice8",
			"Strg": "rd:Trice8 line %d\\n"
		},
		"25373": {
			"Type": "Trice16_1",
			"Strg": "rd: Trice16_1 %u\\n"
		},
		"25382": {
			"Type": "TRICE32_1",
			"Strg": "time:ms = %d\\n"
		},
		"25765": {
			"Type": "Trice0i",
			"Strg": "int:SysTick_Handler\\n"
		},
		"25873": {
			"Type": "Trice8_2",
			"Strg": "diag: Trice8_2  0x%02X, 0x%02x\\n"
		},
		"25880": {
			"Type": "TRICE8_6",
			"Strg": "tst:TRICE8_6 %02X %02X %02X %02X %02X %02X\\n"
		},
		"26286": {
			"Type": "TRICE0",
			"Strg": "Play 'sound.wav'\\n-\u003e "
		},
		"26825": {
			"Type": "Trice8_1i",
			"Strg": "msg:  Trice8_1i -\u003e short  trice macro    (no   cycle counter) for only inside critical section %d\\n"
		},
		"26891": {
			"Type": "Trice0",
			"Strg": "wr: Trice0    short\\n"
		},
		"27064": {
			"Type": "Trice0i",
			"Strg": "wr: Trice0i   short\\n"
		},
		"27253": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:   MDK-ARM_LL_UART_RTT0_PACK_STM32F030_NUCLEO-64   \\ns:                                                   \\n\\n"
		},
		"27424": {
			"Type": "Trice0",
			"Strg": "wr: Trice0    short\\n"
		},
		"27489": {
			"Type": "TRICE16_1",
			"Strg": "MSG: triceFifoMaxDepth = %d\\n"
		},
		"27565": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTTD_NUCLEO-F070RB\\n"
		},
		"27590": {
			"Type": "TRICE16_2",
			"Strg": "MSG: triceFifoMaxDepth: Bare = %d, Esc = %d\\n"
		},
		"27624": {
			"Type": "TRICE0",
			"Strg": "att:TASKING_STM32F4DISC_Audio_Service_RTTD\\n"
		},
		"28071": {
			"Type": "TRICE32_1",
			"Strg": "time:TRICE32_1   message, SysTick is %6d\\n"
		},
		"28450": {
			"Type": "trice8_3",
			"Strg": "tst:trice8_3 %02X %02X %02X\\n"
		},
		"28835": {
			"Type": "Trice0",
			"Strg": "msg: Trice0  -\u003e short  trice macro    (no   cycle counter) for everywhere\\n"
		},
		"29832": {
			"Type": "Trice16_1i",
			"Strg": "rd: Trice16_1i %d\\n"
		},
		"29874": {
			"Type": "Trice16_1",
			"Strg": "msg: Trice16_1 -\u003e short trice macro (no cycle counter) for everywhere\\n"
		},
		"30070": {
			"Type": "Trice16_1i",
			"Strg": "INT:Trice16_1i SysTick: %5u\\n"
		},
		"30221": {
			"Type": "TRICE0",
			"Strg": "att:atollicTrueSTUDIO_RTTD_DISCOVERY-STM32F407VGTx\\n"
		},
		"30623": {
			"Type": "TRICE16_2",
			"Strg": "tst:TRICE16_2 %u %u\\n"
		},
		"30688": {
			"Type": "TRICE0",
			"Strg": "\\ns:                                                     \\ns:   ARM-MDK_LL_UART_RTT0_ESC_STM32F030R8_NUCLEO-64    \\ns:                                                     \\n\\n"
		},
		"30693": {
			"Type": "Trice8_2i",
			"Strg": "diag: Trice8_2i  %x, %x\\n"
		},
		"30902": {
			"Type": "Trice16_1",
			"Strg": "time: Trice16_1 %6u\\n"
		},
		"31050": {
			"Type": "Trice0",
			"Strg": "msg: Trice0  -\u003e short trice macro (no cycle counter) for everywhere\\n"
		},
		"31376": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4 %u %u %u %u\\n"
		},
		"31407": {
			"Type": "Trice16_1",
			"Strg": "time: Trice16_1 %d\\n"
		},
		"31424": {
			"Type": "Trice8",
			"Strg": "rd:%d\\n"
		},
		"31759": {
			"Type": "trice8_7",
			"Strg": "tst:trice8_7 %02X %02X %02X %02X %02X %02X %02X\\n"
		},
		"31820": {
			"Type": "TRICE16_1",
			"Strg": "ISR:interrupt   message, SysTick is %6u\\n"
		},
		"31978": {
			"Type": "Trice8i",
			"Strg": "rd:%d\\n"
		},
		"32337": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i -\u003e short  trice macro    (no   cycle counter) for only inside critical section %d\\n"
		},
		"32731": {
			"Type": "TRICE0",
			"Strg": "Calculate c1 + c2: "
		},
		"32742": {
			"Type": "TRICE8_1",
			"Strg": "tst:TRICE8_1 %u\\n"
		},
		"32800": {
			"Type": "TRICE64",
			"Strg": "rd:TRICE64 line %d,%d\\n"
		},
		"32834": {
			"Type": "TRICE64_1i",
			"Strg": "rd:%d\\n"
		},
		"32881": {
			"Type": "TRICE8_5",
			"Strg": "msg:  TRICE8_5  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"33067": {
			"Type": "TRICE16_1i",
			"Strg": "msg: TRICE16_1i -\u003e normal trice macro    (with cycle counter) for only inside critical section %d\\n"
		},
		"33175": {
			"Type": "TRICE32",
			"Strg": "rd:%d\\n"
		},
		"33244": {
			"Type": "TRICE32_2i",
			"Strg": "rd:%c%c"
		},
		"33304": {
			"Type": "TRICE8",
			"Strg": "rd:TRICE8 line %d\\n"
		},
		"3338": {
			"Type": "Trice0i",
			"Strg": "int:Trice0i SysTick_Handler"
		},
		"33417": {
			"Type": "TRICE32_4",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"33465": {
			"Type": "trice64_2",
			"Strg": "tst:trice64_2 %x %u\\n"
		},
		"33496": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_RTT_NUCLEO-F030R8\\n"
		},
		"33523": {
			"Type": "TRICE8_2",
			"Strg": "msg:  TRICE8_2  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X\\n"
		},
		"33574": {
			"Type": "TRICE64_2",
			"Strg": "msg: TRICE64_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"33633": {
			"Type": "trice32_1",
			"Strg": "tst:trice32_1   message, SysTick is %6d\\n"
		},
		"33655": {
			"Type": "trice32i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"33659": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"33682": {
			"Type": "TRICE16",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"337": {
			"Type": "Trice16",
			"Strg": "ATT:Trice16  attention   message, SysTick is %6u\\n"
		},
		"33810": {
			"Type": "TRICE8_5",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"33874": {
			"Type": "TRICE32_1",
			"Strg": "TIME:TRICE32_1   message, SysTick is %6d\\n"
		},
		"33896": {
			"Type": "TRICE8_4i",
			"Strg": "rd:%c%c%c%c"
		},
		"33899": {
			"Type": "trice16i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"33989": {
			"Type": "TRICE16i",
			"Strg": "rd:%d\\n"
		},
		"34064": {
			"Type": "TRICE32_3",
			"Strg": "msg: TRICE32_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"34130": {
			"Type": "trice0",
			"Strg": "a:c"
		},
		"34148": {
			"Type": "trice16i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"34176": {
			"Type": "trice8_5",
			"Strg": "msg:  trice8_5  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"34211": {
			"Type": "TRICE8_2i",
			"Strg": "rd:%d, %d\\n"
		},
		"34274": {
			"Type": "trice0",
			"Strg": "3"
		},
		"34417": {
			"Type": "trice32_3",
			"Strg": "msg: trice32_3  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"34481": {
			"Type": "trice0",
			"Strg": "msg:hi\\n"
		},
		"34736": {
			"Type": "trice8_2",
			"Strg": "msg: trice8_2   -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"34813": {
			"Type": "trice8_2",
			"Strg": "msg:  trice8_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"34814": {
			"Type": "trice16_3",
			"Strg": "msg: trice16_3  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"34868": {
			"Type": "TRICE16_1",
			"Strg": "MSG: triceBareFifoMaxDepth = %d\\n"
		},
		"34882": {
			"Type": "TRICE16_1",
			"Strg": "WR:write        message, SysTick is %6u\\n"
		},
		"34956": {
			"Type": "TRICE32",
			"Strg": "rd:%d\\n"
		},
		"35036": {
			"Type": "trice64_2",
			"Strg": "msg: trice64_2  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X\\n"
		},
		"35055": {
			"Type": "trice16",
			"Strg": "ATT:trice16 attention   message, SysTick is %6u\\n"
		},
		"35199": {
			"Type": "TRICE16",
			"Strg": "ATT:TRICE16 attention   message, SysTick is %6u\\n"
		},
		"35213": {
			"Type": "TRICE8_2",
			"Strg": "msg:  TRICE8_2  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X\\n"
		},
		"35217": {
			"Type": "TRICE8_3",
			"Strg": "msg: TRICE8_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"35238": {
			"Type": "trice16_1",
			"Strg": "DBG:debug       message, SysTick is %6u\\n"
		},
		"35245": {
			"Type": "TRICE16_2i",
			"Strg": "rd:%d, %d\\n"
		},
		"35386": {
			"Type": "TRICE8_8",
			"Strg": "rd:%c%c%c%c%c%c%c%c"
		},
		"35473": {
			"Type": "TRICE32_3i",
			"Strg": "rd:%c%c%c"
		},
		"35527": {
			"Type": "TRICE0",
			"Strg": "att:IAR_EWARM_LL_UART_NUCLEO-F070RB\\n"
		},
		"35624": {
			"Type": "TRICE16_1i",
			"Strg": "rd:%d\\n"
		},
		"35689": {
			"Type": "trice0",
			"Strg": "w:B"
		},
		"35740": {
			"Type": "TRICE0",
			"Strg": "att:IAR_EWARM_RTTD_DISCOVERY-STM32F407VGTx\\n"
		},
		"35857": {
			"Type": "TRICE8_6i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"35916": {
			"Type": "TRICE16i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"35941": {
			"Type": "trice8_1",
			"Strg": "msg:  trice8_1  -\u003e normal trice function (with cycle counter) for everywhere                   %u\\n"
		},
		"36033": {
			"Type": "trice8_4",
			"Strg": "msg:  trice8_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"36071": {
			"Type": "TRICE32_3",
			"Strg": "msg: TRICE32_3  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"36297": {
			"Type": "TRICE0",
			"Strg": "s:                                              \\ns:   ARM-MDK_LL_UART_RTT0_ESC_STM32F070RB_NUCLEO-64    \\ns:                                              \\n\\n"
		},
		"36350": {
			"Type": "TRICE16_2i",
			"Strg": "rd:%c%c"
		},
		"36351": {
			"Type": "trice16_1",
			"Strg": "MSG:normal      message, SysTick is %6u\\n"
		},
		"36379": {
			"Type": "trice0",
			"Strg": "time:i"
		},
		"36395": {
			"Type": "TRICE8_4",
			"Strg": "msg: TRICE8_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"36399": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_RTT1_NUCLEO-F091RC\\n"
		},
		"36461": {
			"Type": "trice8_1",
			"Strg": "msg: trice8_1   -\u003e normal trice function (with cycle counter) for everywhere %u\\n"
		},
		"36547": {
			"Type": "TRICE8_8",
			"Strg": "tst:TRICE8_8 %u %u %u %u %u %u %u %u\\n"
		},
		"36633": {
			"Type": "trice0",
			"Strg": "e:A"
		},
		"36646": {
			"Type": "trice0",
			"Strg": "--------------------------------------------------\\n"
		},
		"36836": {
			"Type": "TRICE16_3i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"36916": {
			"Type": "TRICE8_7",
			"Strg": "msg:  TRICE8_7  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"36993": {
			"Type": "TRICE8_4i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"37067": {
			"Type": "trice8_2i",
			"Strg": "msg : trice8_2i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X\\n"
		},
		"37121": {
			"Type": "TRICE32_2",
			"Strg": "rd:%c%c"
		},
		"37187": {
			"Type": "TRICE8_3",
			"Strg": "msg:  TRICE8_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"37205": {
			"Type": "TRICE32_3",
			"Strg": "msg: TRICE32_3  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"37347": {
			"Type": "trice8_8",
			"Strg": "msg:  trice8_8  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"37353": {
			"Type": "TRICE16i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"37386": {
			"Type": "TRICE8_8i",
			"Strg": "msg:  TRICE8_8i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d %x %X %u %d\\n"
		},
		"37489": {
			"Type": "TRICE64_2i",
			"Strg": "rd:%d, %d\\n"
		},
		"37503": {
			"Type": "TRICE8_1",
			"Strg": "rd:%d\\n"
		},
		"37519": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:   MDK-ARM_LL_UART_RTT0_BARE_STM32F091_NUCLEO-64   \\ns:                                                   \\n\\n"
		},
		"37676": {
			"Type": "trice8_7",
			"Strg": "msg:  trice8_7  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"37693": {
			"Type": "trice16_2",
			"Strg": "msg: trice16_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"37773": {
			"Type": "trice0",
			"Strg": "message:J"
		},
		"37880": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"37902": {
			"Type": "TRICE16_2",
			"Strg": "msg: TRICE16_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"37968": {
			"Type": "trice8_2",
			"Strg": "msg:  trice8_2  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X\\n"
		},
		"38013": {
			"Type": "trice0",
			"Strg": "e:7"
		},
		"38168": {
			"Type": "TRICE32_4",
			"Strg": "tst:TRICE32_4 %u %u %u %u\\n"
		},
		"38214": {
			"Type": "TRICE32_2",
			"Strg": "tst:TRICE32_2 %u %u\\n"
		},
		"3826": {
			"Type": "Trice16_1",
			"Strg": "int:SysTick: %5u\\n"
		},
		"38267": {
			"Type": "TRICE32_4i",
			"Strg": "rd:%c%c%c%c"
		},
		"38516": {
			"Type": "TRICE32_1",
			"Strg": "tim :TRICE32_1   message, SysTick is %6d\\n"
		},
		"38569": {
			"Type": "trice0",
			"Strg": "2"
		},
		"38669": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"38727": {
			"Type": "TRICE64i",
			"Strg": "rd:%d\\n"
		},
		"38736": {
			"Type": "TRICE64_2",
			"Strg": "msg: TRICE64_2  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X\\n"
		},
		"38748": {
			"Type": "trice16_4",
			"Strg": "msg: trice16_4  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"38802": {
			"Type": "TRICE0",
			"Strg": "Calculate c2 - c1: "
		},
		"38837": {
			"Type": "trice16_1",
			"Strg": "msg: trice16_1  -\u003e normal trice function (with cycle counter) for everywhere                   %u\\n"
		},
		"38839": {
			"Type": "TRICE8_2i",
			"Strg": "rd:%c%c"
		},
		"39062": {
			"Type": "TRICE64i",
			"Strg": "rd:%d, %d\\n"
		},
		"39086": {
			"Type": "TRICE8_4",
			"Strg": "msg:  TRICE8_4  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"39100": {
			"Type": "trice16_1",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"39163": {
			"Type": "TRICE64",
			"Strg": "rd:%d\\n"
		},
		"39197": {
			"Type": "trice32",
			"Strg": "rd:trice32 line %d\\n"
		},
		"39202": {
			"Type": "TRICE8_1i",
			"Strg": "rd:%c"
		},
		"39245": {
			"Type": "trice16_3",
			"Strg": "msg: trice16_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"39375": {
			"Type": "trice8_7i",
			"Strg": "msg:  trice8_7i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d %x %X %u\\n"
		},
		"39406": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:   MDK-ARM_LL_UART_RTT0_BARE_STM32F030_NUCLEO-64   \\ns:                                                   \\n\\n"
		},
		"39407": {
			"Type": "trice0",
			"Strg": "msg: trice0  -\u003e normal trice function for everywhere\\n"
		},
		"39412": {
			"Type": "TRICE0i",
			"Strg": "msg: TRICE0i    -\u003e normal trice macro    (with cycle counter) for only inside critical section\\n"
		},
		"39419": {
			"Type": "trice64_1i",
			"Strg": "msg: trice64_1i -\u003e normal trice function (with cycle counter) for only inside critical section %d\\n"
		},
		"39459": {
			"Type": "trice16",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"39533": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_RTTD_NUCLEO-G474\\n"
		},
		"39704": {
			"Type": "TRICE32_1",
			"Strg": "tim  :TRICE32_1   message, SysTick is %6d\\n"
		},
		"39892": {
			"Type": "TRICE16_1",
			"Strg": "msg: TRICE16_1  -\u003e normal trice macro    (with cycle counter) for everywhere                   %u\\n"
		},
		"39907": {
			"Type": "TRICE32i",
			"Strg": "rd:%d, %d\\n"
		},
		"39935": {
			"Type": "TRICE16i",
			"Strg": "rd:%d, %d\\n"
		},
		"39944": {
			"Type": "trice16_1",
			"Strg": "WRN:warning     message, SysTick is %6u\\n"
		},
		"39968": {
			"Type": "TRICE16_3",
			"Strg": "rd:%c%c%c"
		},
		"40040": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"40045": {
			"Type": "TRICE0",
			"Strg": "Negate previous  : "
		},
		"40168": {
			"Type": "TRICE32_2i",
			"Strg": "rd:%d, %d\\n"
		},
		"40236": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"40252": {
			"Type": "trice16_1",
			"Strg": "tst:trice16_1   message, SysTick is %6u\\n"
		},
		"40468": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTTB_NUCLEO-F070RB\\n"
		},
		"40494": {
			"Type": "TRICE16_1i",
			"Strg": "rd:%c"
		},
		"40865": {
			"Type": "TRICE16_1",
			"Strg": "tst:TRICE16_1 %u\\n"
		},
		"40953": {
			"Type": "TRICE0",
			"Strg": "msg: TRICE0  -\u003e normal trice macro    (with cycle counter) for everywhere\\n"
		},
		"41026": {
			"Type": "TRICE0",
			"Strg": "Failed to initialize audio service"
		},
		"41032": {
			"Type": "TRICE8_3",
			"Strg": "msg:  TRICE8_3  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"41105": {
			"Type": "TRICE8_2",
			"Strg": "msg:  TRICE8_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"41184": {
			"Type": "TRICE16_1",
			"Strg": "tim :TRICE16_1   message, SysTick is %6u\\n"
		},
		"41306": {
			"Type": "TRICE32_1",
			"Strg": "tim:TRICE32_1   message, SysTick is %6d\\n"
		},
		"41366": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"41399": {
			"Type": "TRICE64_2",
			"Strg": "rd:%d, %d\\n"
		},
		"41495": {
			"Type": "trice8_8",
			"Strg": "msg:  trice8_8  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"41511": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:   MDK-ARM_LL_UART_RTT0_FLEX_STM32F030_NUCLEO-64   \\ns:                                                   \\n\\n"
		},
		"41604": {
			"Type": "TRICE8_5",
			"Strg": "tst:TRICE8_5 %u %u %u %u %u\\n"
		},
		"41613": {
			"Type": "TRICE32i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"41644": {
			"Type": "TRICE8_4",
			"Strg": "msg:  TRICE8_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"41697": {
			"Type": "TRICE16_3i",
			"Strg": "msg: TRICE16_3i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"41736": {
			"Type": "trice32_3",
			"Strg": "msg: trice32_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"41764": {
			"Type": "TRICE32_4i",
			"Strg": "msg: TRICE32_4i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"41804": {
			"Type": "TRICE0",
			"Strg": "att:IAR_EWARM_HAL_UART_NUCLEO-F070RB\\n"
		},
		"41820": {
			"Type": "TRICE8_5",
			"Strg": "msg:  TRICE8_5  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"41845": {
			"Type": "trice64_2i",
			"Strg": "msg: trice64_2i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X\\n"
		},
		"41911": {
			"Type": "trice64i",
			"Strg": "rd:%d, %d\\n"
		},
		"41944": {
			"Type": "TRICE32_1",
			"Strg": "rd:%d\\n"
		},
		"42009": {
			"Type": "trice16",
			"Strg": "rd:%d\\n"
		},
		"4205": {
			"Type": "Trice0i",
			"Strg": "msg: Trice0i -\u003e short trice macro (no cycle counter) for only inside critical section\\n"
		},
		"42100": {
			"Type": "TRICE0",
			"Strg": "att:IAR_EWARM_LL_UART_NUCLEO-F030RB\\n"
		},
		"42133": {
			"Type": "trice64_2",
			"Strg": "msg: trice64_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"42137": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"42270": {
			"Type": "TRICE16_4",
			"Strg": "msg: TRICE16_4  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"42481": {
			"Type": "trice64",
			"Strg": "rd:%d, %d\\n"
		},
		"42601": {
			"Type": "trice16_1",
			"Strg": "RD:read         message, SysTick is %6u\\n"
		},
		"42602": {
			"Type": "TRICE8_6",
			"Strg": "msg:  TRICE6_5  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"42699": {
			"Type": "TRICE64_1i",
			"Strg": "rd:%c"
		},
		"42720": {
			"Type": "TRICE8_1",
			"Strg": "msg:  TRICE8_1  -\u003e normal trice macro    (with cycle counter) for everywhere                   %u\\n"
		},
		"42899": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_UART_NUCLEO-F070RB\\n"
		},
		"42903": {
			"Type": "trice16",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"42963": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTTB_NUCLEO-F030R8\\n"
		},
		"43030": {
			"Type": "TRICE32",
			"Strg": "rd:%d,%d\\n"
		},
		"43095": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d\\n"
		},
		"43124": {
			"Type": "TRICE16_1",
			"Strg": "rd:%d\\n"
		},
		"43141": {
			"Type": "trice8_3",
			"Strg": "msg: trice8_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"43310": {
			"Type": "TRICE0",
			"Strg": "att:IAR_EWARM_RTT_NUCLEO-F030R8\\n"
		},
		"43336": {
			"Type": "TRICE0",
			"Strg": "s:                                          \\ns:    ARM-MDK_RTT0_BARE_STM32F0308-DISCO    \\ns:                                          \\n\\n"
		},
		"43353": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d\\n"
		},
		"43451": {
			"Type": "trice32_1",
			"Strg": "msg: trice32_1  -\u003e normal trice function (with cycle counter) for everywhere                   %u\\n"
		},
		"43499": {
			"Type": "TRICE0",
			"Strg": "Complex number c2: "
		},
		"43675": {
			"Type": "trice0",
			"Strg": "m:123\\n"
		},
		"43733": {
			"Type": "trice16_1",
			"Strg": "ISR:interrupt   message, SysTick is %6u\\n"
		},
		"43767": {
			"Type": "TRICE32_3",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"43916": {
			"Type": "trice8_6",
			"Strg": "tst:trice8_6 %02X %02X %02X %02X %02X %02X\\n"
		},
		"43956": {
			"Type": "trice0",
			"Strg": "--------------------------------------------------\\n"
		},
		"44006": {
			"Type": "TRICE16_4i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"44117": {
			"Type": "trice16_1",
			"Strg": "DIA:diagnostics message, SysTick is %6u\\n"
		},
		"44137": {
			"Type": "TRICE0",
			"Strg": "att:TASKING_GenericSTMF070RB_RTTD\\n"
		},
		"44173": {
			"Type": "trice64",
			"Strg": "rd:%d,%d\\n"
		},
		"44262": {
			"Type": "TRICE16_4",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"44310": {
			"Type": "TRICE32_1",
			"Strg": "rd:%c"
		},
		"44324": {
			"Type": "TRICE8",
			"Strg": "rd:TRICE8 line %d, %d\\n"
		},
		"44361": {
			"Type": "TRICE16_1",
			"Strg": "rd:%c"
		},
		"44389": {
			"Type": "trice64_1",
			"Strg": "msg: trice64_1  -\u003e normal trice function (with cycle counter) for everywhere %u\\n"
		},
		"44414": {
			"Type": "TRICE0",
			"Strg": "att:TASKING_GenericSTMF030R8_RTTB\\n"
		},
		"44543": {
			"Type": "TRICE32_3",
			"Strg": "rd:%c%c%c"
		},
		"44713": {
			"Type": "TRICE0",
			"Strg": "msg: TRICE0  -\u003e normal trice macro for everywhere\\n"
		},
		"44719": {
			"Type": "TRICE8_4",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"44727": {
			"Type": "TRICE32_2",
			"Strg": "rd:%d, %d\\n"
		},
		"44749": {
			"Type": "TRICE16_3",
			"Strg": "msg: TRICE16_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"44778": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_RTT_NUCLEO-F070RB\\n"
		},
		"44807": {
			"Type": "TRICE8_2",
			"Strg": "rd:%c%c"
		},
		"44870": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_RTTD_DISCOVERY-STM32F051R8Tx\\n"
		},
		"44960": {
			"Type": "TRICE32_1",
			"Strg": "msg: TRICE32_1  -\u003e normal trice macro    (with cycle counter) for everywhere                   %u\\n"
		},
		"44967": {
			"Type": "trice16_1",
			"Strg": "SIG:signal      message, SysTick is %6u\\n"
		},
		"44974": {
			"Type": "TRICE8_7",
			"Strg": "msg:  TRICE8_7  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"4500": {
			"Type": "Trice8_2i",
			"Strg": "diag: Trice8_2  0x%02X, 0x%02x\\n"
		},
		"45064": {
			"Type": "TRICE16_1",
			"Strg": "DIA:diagnostics message, SysTick is %6u\\n"
		},
		"45086": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"45214": {
			"Type": "trice8_2",
			"Strg": "msg:  trice8_2  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X\\n"
		},
		"45250": {
			"Type": "TRICE64_1",
			"Strg": "tst:TRICE64_1 %u\\n"
		},
		"45309": {
			"Type": "trice8_5i",
			"Strg": "msg:  trice8_5i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %u %d\\n"
		},
		"45348": {
			"Type": "trice16_2",
			"Strg": "msg: trice16_2  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X\\n"
		},
		"45424": {
			"Type": "TRICE8_6",
			"Strg": "msg:  TRICE6_5  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"45461": {
			"Type": "TRICE16_2",
			"Strg": "rd:%c%c"
		},
		"45471": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_HAL_UART_NUCLEO-F070RB\\n"
		},
		"45598": {
			"Type": "trice32_4i",
			"Strg": "msg: trice32_4i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"45671": {
			"Type": "trice16_1",
			"Strg": "INFO:informal   message, SysTick is %6u\\n"
		},
		"45744": {
			"Type": "TRICE32",
			"Strg": "rd:%d, %d\\n"
		},
		"45795": {
			"Type": "TRICE16",
			"Strg": "rd:%d\\n"
		},
		"45848": {
			"Type": "TRICE8_7",
			"Strg": "msg:  TRICE8_7  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"45876": {
			"Type": "TRICE8_5i",
			"Strg": "rd:%c%c%c%c%c"
		},
		"46050": {
			"Type": "TRICE0",
			"Strg": "att:STM32CubeIDE_HAL_UART_NUCLEO-F070RB\\n"
		},
		"46430": {
			"Type": "TRICE32_1i",
			"Strg": "rd:%c"
		},
		"46494": {
			"Type": "trice0",
			"Strg": "wr:d"
		},
		"46624": {
			"Type": "trice0",
			"Strg": "msg: trice0  -\u003e normal trice function (with cycle counter) for everywhere\\n"
		},
		"46836": {
			"Type": "TRICE16_1",
			"Strg": "tst:TRICE16_1   message, SysTick is %6u\\n"
		},
		"47057": {
			"Type": "TRICE8_7",
			"Strg": "rd:%c%c%c%c%c%c%c"
		},
		"47069": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"47137": {
			"Type": "TRICE8_1",
			"Strg": "msg: TRICE8_1   -\u003e normal trice macro    (with cycle counter) for everywhere                   %u\\n"
		},
		"47283": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_HAL_UART_NUCLEO-F030R8\\n"
		},
		"47307": {
			"Type": "TRICE8_7",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"47359": {
			"Type": "TRICE16_4i",
			"Strg": "rd:%c%c%c%c"
		},
		"47395": {
			"Type": "trice16_4i",
			"Strg": "msg: trice16_4i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"47464": {
			"Type": "trice16_3i",
			"Strg": "msg: trice16_3i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"47472": {
			"Type": "trice8_8i",
			"Strg": "msg:  trice8_8i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d %x %X %u %d\\n"
		},
		"47530": {
			"Type": "TRICE32_2",
			"Strg": "msg: TRICE32_2  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X\\n"
		},
		"47663": {
			"Type": "TRICE16_2",
			"Strg": "MSG: triceFifoMaxDepth = %d, select = %d\\n"
		},
		"47796": {
			"Type": "trice8",
			"Strg": "rd:%d\\n"
		},
		"47820": {
			"Type": "trice8_2i",
			"Strg": "msg:  trice8_2i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X\\n"
		},
		"47920": {
			"Type": "TRICE8_1i",
			"Strg": "rd:%d\\n"
		},
		"48114": {
			"Type": "TRICE0",
			"Strg": "att:MDK-ARM_LL_UART_demoBoard_STM32F030F4F4P6\\n"
		},
		"48178": {
			"Type": "trice8",
			"Strg": "rd:%d, %d\\n"
		},
		"48203": {
			"Type": "TRICE16_4i",
			"Strg": "msg: TRICE16_4i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"48220": {
			"Type": "TRICE64_1",
			"Strg": "rd:%d\\n"
		},
		"48322": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"48379": {
			"Type": "trice16",
			"Strg": "tst:trice16_1   message, SysTick is %6u\\n"
		},
		"48411": {
			"Type": "TRICE16_3",
			"Strg": "msg: TRICE16_3  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"48416": {
			"Type": "TRICE32_4",
			"Strg": "msg: TRICE32_4  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"48424": {
			"Type": "trice8_8",
			"Strg": "msg:  trice8_8  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"48439": {
			"Type": "TRICE16_1",
			"Strg": "tim:TRICE16_1   message, SysTick is %6u\\n"
		},
		"48560": {
			"Type": "TRICE8_1",
			"Strg": "rd:%c"
		},
		"48574": {
			"Type": "TRICE16_3i",
			"Strg": "rd:%c%c%c"
		},
		"48889": {
			"Type": "trice16_2",
			"Strg": "msg: trice16_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"49133": {
			"Type": "trice16_2",
			"Strg": "msg: trice16_2  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X\\n"
		},
		"4929": {
			"Type": "Trice0i",
			"Strg": "int:Trice0i SysTick_Handler\\n"
		},
		"49344": {
			"Type": "TRICE16_2",
			"Strg": "rd:%d, %d\\n"
		},
		"49379": {
			"Type": "trice16_1",
			"Strg": "TIM:timing      message, SysTick is %6u\\n"
		},
		"49408": {
			"Type": "trice16",
			"Strg": "tst:trice16_1   message, SysTick is %6u\\n"
		},
		"49458": {
			"Type": "TRICE8_4i",
			"Strg": "msg:  TRICE8_4i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"49749": {
			"Type": "TRICE8_6",
			"Strg": "msg:  TRICE6_5  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"49757": {
			"Type": "TRICE32_1",
			"Strg": "att:TRICE32_1   message, SysTick is %6d\\n"
		},
		"49840": {
			"Type": "TRICE64_2",
			"Strg": "rd:%c%c"
		},
		"49913": {
			"Type": "TRICE64_1",
			"Strg": "rd:%c"
		},
		"5": {
			"Type": "Trice0",
			"Strg": "msg:Qick\\n"
		},
		"50010": {
			"Type": "TRICE8_3i",
			"Strg": "rd:%c%c%c"
		},
		"50184": {
			"Type": "TRICE8_3",
			"Strg": "rd:%c%c%c"
		},
		"50329": {
			"Type": "trice0",
			"Strg": "msg: trice0     -\u003e normal trice function (with cycle counter) for everywhere\\n"
		},
		"50424": {
			"Type": "trice64_2",
			"Strg": "msg: trice64_2  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X\\n"
		},
		"50437": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"50448": {
			"Type": "trice8_1i",
			"Strg": "msg:  trice8_1i -\u003e normal trice function (with cycle counter) for only inside critical section %d\\n"
		},
		"50478": {
			"Type": "trice8_5",
			"Strg": "msg:  trice8_5  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"50501": {
			"Type": "trice8_7",
			"Strg": "msg:  trice8_7  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"50571": {
			"Type": "TRICE64",
			"Strg": "rd:TRICE64 %d\\n"
		},
		"50621": {
			"Type": "TRICE8_6",
			"Strg": "rd:%c%c%c%c%c%c"
		},
		"50624": {
			"Type": "TRICE16_4",
			"Strg": "msg: TRICE16_4  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"50677": {
			"Type": "TRICE32i",
			"Strg": "rd:%d\\n"
		},
		"50757": {
			"Type": "TRICE16_2",
			"Strg": "msg: TRICE16_2  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X\\n"
		},
		"50759": {
			"Type": "TRICE16_4",
			"Strg": "rd:%c%c%c%c"
		},
		"50770": {
			"Type": "TRICE8_6i",
			"Strg": "msg:  TRICE6_5i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d %x %X\\n"
		},
		"50930": {
			"Type": "trice8i",
			"Strg": "rd:%d\\n"
		},
		"50995": {
			"Type": "TRICE32_2",
			"Strg": "msg: TRICE32_2  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X\\n"
		},
		"51003": {
			"Type": "trice0",
			"Strg": "t:H"
		},
		"51048": {
			"Type": "TRICE8_3",
			"Strg": "msg:  TRICE8_3  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"51084": {
			"Type": "trice0i",
			"Strg": "msg: trice0i    -\u003e normal trice function (with cycle counter) for only inside critical section\\n"
		},
		"51091": {
			"Type": "trice32i",
			"Strg": "rd:%d\\n"
		},
		"51112": {
			"Type": "trice32",
			"Strg": "rd:trice32 line %d,%d\\n"
		},
		"51119": {
			"Type": "TRICE8_3",
			"Strg": "tst:TRICE8_3 %02x %02x %02x\\n"
		},
		"51156": {
			"Type": "trice32i",
			"Strg": "rd:%d, %d\\n"
		},
		"51207": {
			"Type": "TRICE8_6",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"51208": {
			"Type": "trice0",
			"Strg": "rd:e\\n"
		},
		"5124": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i %6u\\n"
		},
		"51264": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d\\n"
		},
		"51293": {
			"Type": "TRICE32_4",
			"Strg": "msg: TRICE32_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"51298": {
			"Type": "trice8_6",
			"Strg": "msg:  trice6_5  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"51309": {
			"Type": "trice8_3",
			"Strg": "msg:  trice8_3  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"51375": {
			"Type": "trice0",
			"Strg": "1"
		},
		"51484": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"51658": {
			"Type": "TRICE32_1i",
			"Strg": "msg: TRICE32_1i -\u003e normal trice macro    (with cycle counter) for only inside critical section %d\\n"
		},
		"51718": {
			"Type": "trice0",
			"Strg": "--------------------------------------------------\\n"
		},
		"51750": {
			"Type": "trice8_4i",
			"Strg": "msg:  trice8_4i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d\\n"
		},
		"51896": {
			"Type": "TRICE8",
			"Strg": "rd:%d\\n"
		},
		"51953": {
			"Type": "TRICE16_2",
			"Strg": "msg: TRICE16_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"52162": {
			"Type": "trice32_2",
			"Strg": "msg: trice32_2  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X\\n"
		},
		"52174": {
			"Type": "TRICE8_1",
			"Strg": "tst:TRICE8_1 %02x\\n"
		},
		"52191": {
			"Type": "trice0",
			"Strg": "diag:f"
		},
		"52233": {
			"Type": "trice32i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"52308": {
			"Type": "TRICE8_4",
			"Strg": "rd:%c%c%c%c"
		},
		"52309": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"52433": {
			"Type": "trice32_2",
			"Strg": "msg: trice32_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"52613": {
			"Type": "trice32_3i",
			"Strg": "msg: trice32_3i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"52748": {
			"Type": "trice8_5",
			"Strg": "msg:  trice8_5  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"52964": {
			"Type": "TRICE8_8",
			"Strg": "msg:  TRICE8_8  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"52986": {
			"Type": "trice64i",
			"Strg": "rd:%d\\n"
		},
		"53022": {
			"Type": "trice64",
			"Strg": "rd:trice64 line %d,%d\\n"
		},
		"53090": {
			"Type": "TRICE8_1",
			"Strg": "msg: TRICE8_1  -\u003e normal trice macro    (with cycle counter) for everywhere %u\\n"
		},
		"53154": {
			"Type": "trice16i",
			"Strg": "rd:%d\\n"
		},
		"53187": {
			"Type": "trice8_1",
			"Strg": "msg: trice8_1   -\u003e normal trice function (with cycle counter) for everywhere                   %u\\n"
		},
		"53200": {
			"Type": "TRICE8_5",
			"Strg": "tst:TRICE8_5 %02x %02x %02x %02x %02x\\n"
		},
		"53204": {
			"Type": "trice32_4",
			"Strg": "msg: trice32_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"53209": {
			"Type": "trice32_1",
			"Strg": "msg: trice32_1  -\u003e normal trice function (with cycle counter) for everywhere %u\\n"
		},
		"53291": {
			"Type": "TRICE16",
			"Strg": "rd:%d, %d\\n"
		},
		"53546": {
			"Type": "TRICE8_7",
			"Strg": "tst:TRICE8_7 %02X %02X %02X %02X %02X %02X %02X\\n"
		},
		"53639": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"53783": {
			"Type": "TRICE32_1",
			"Strg": "msg: TRICE32_1  -\u003e normal trice macro    (with cycle counter) for everywhere %u\\n"
		},
		"54014": {
			"Type": "TRICE8_8",
			"Strg": "msg:  TRICE8_8  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"54079": {
			"Type": "TRICE8_2i",
			"Strg": "msg:  TRICE8_2i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X\\n"
		},
		"54250": {
			"Type": "trice32",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"54294": {
			"Type": "TRICE32",
			"Strg": "rd:TRICE32 line %d,%d\\n"
		},
		"54544": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"54569": {
			"Type": "trice8_4",
			"Strg": "msg:  trice8_4  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"54701": {
			"Type": "TRICE32_4",
			"Strg": "msg: TRICE32_4  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"54733": {
			"Type": "TRICE8",
			"Strg": "rd:%d\\n"
		},
		"55076": {
			"Type": "trice8_1",
			"Strg": "msg: trice8_1  -\u003e normal trice function (with cycle counter) for everywhere %u\\n"
		},
		"55218": {
			"Type": "TRICE8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"55241": {
			"Type": "trice16_4",
			"Strg": "msg: trice16_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"55290": {
			"Type": "trice32_1",
			"Strg": "tst:trice32_1   message, SysTick is %6d\\n"
		},
		"55372": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"55382": {
			"Type": "trice64",
			"Strg": "rd:trice64 %d\\n"
		},
		"55451": {
			"Type": "trice8_6",
			"Strg": "msg:  trice6_5  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"55537": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"55595": {
			"Type": "TRICE0",
			"Strg": "rd_:triceFunctions.c"
		},
		"55609": {
			"Type": "TRICE8i",
			"Strg": "rd:%d\\n"
		},
		"55637": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"55697": {
			"Type": "TRICE8_7",
			"Strg": "tst:TRICE8_7 %u %u %u %u %u %u %u\\n"
		},
		"55715": {
			"Type": "trice16_1",
			"Strg": "WR:write        message, SysTick is %6u\\n"
		},
		"55780": {
			"Type": "trice32_2i",
			"Strg": "msg: trice32_2i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X\\n"
		},
		"55892": {
			"Type": "trice16_4",
			"Strg": "msg: trice16_4  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"56175": {
			"Type": "TRICE16_1",
			"Strg": "msg: TRICE16_1  -\u003e normal trice macro    (with cycle counter) for everywhere %u\\n"
		},
		"56183": {
			"Type": "TRICE8_5i",
			"Strg": "msg:  TRICE8_5i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %u %d\\n"
		},
		"56218": {
			"Type": "TRICE8_5",
			"Strg": "msg:  TRICE8_5  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %u %d\\n"
		},
		"56325": {
			"Type": "trice8_6",
			"Strg": "msg:  trice6_5  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d %x %X\\n"
		},
		"56364": {
			"Type": "TRICE8_1",
			"Strg": "msg: TRICE8_1   -\u003e normal trice macro    (with cycle counter) for everywhere %u\\n"
		},
		"56427": {
			"Type": "trice0",
			"Strg": "sig:This ASSERT error is just a demo and no real error:\\n"
		},
		"56437": {
			"Type": "TRICE8_8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"56442": {
			"Type": "trice8_7",
			"Strg": "msg:  trice8_7  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d %x %X %u\\n"
		},
		"56529": {
			"Type": "TRICE8_3",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"56533": {
			"Type": "trice32",
			"Strg": "rd:%d, %d\\n"
		},
		"56542": {
			"Type": "trice32",
			"Strg": "rd:%d\\n"
		},
		"5678": {
			"Type": "Trice16_1i",
			"Strg": "rd: Trice16_1i %d\\n"
		},
		"56998": {
			"Type": "trice0",
			"Strg": "dbg:k\\n"
		},
		"57007": {
			"Type": "TRICE16_1",
			"Strg": "MSG:normal      message, SysTick is %6u\\n"
		},
		"57067": {
			"Type": "trice32_2",
			"Strg": "msg: trice32_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"57154": {
			"Type": "trice8_4",
			"Strg": "msg:  trice8_4  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"57162": {
			"Type": "trice16",
			"Strg": "rd:%d, %d\\n"
		},
		"57618": {
			"Type": "TRICE16",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"579": {
			"Type": "TRICE8_8",
			"Strg": "tst:TRICE8_1 %%d=%d, %%u=%u, 0x%%x=0x%x, 0x%%2x=0x%2x, 0x%%02x=0x%02x, 0x%%3x=0x%3x, 0x%%03x=0x%03x, %%b=%b\\n"
		},
		"57931": {
			"Type": "TRICE16_3",
			"Strg": "msg: TRICE16_3  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"57939": {
			"Type": "TRICE64_2i",
			"Strg": "msg: TRICE64_2i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X\\n"
		},
		"57957": {
			"Type": "TRICE32_2",
			"Strg": "msg: TRICE32_2  -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"58025": {
			"Type": "TRICE32_3i",
			"Strg": "msg: TRICE32_3i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"58039": {
			"Type": "TRICE32",
			"Strg": "rd:TRICE32 line %d\\n"
		},
		"58222": {
			"Type": "TRICE8_6i",
			"Strg": "rd:%c%c%c%c%c%c"
		},
		"58278": {
			"Type": "TRICE64",
			"Strg": "rd:%d\\n"
		},
		"58373": {
			"Type": "TRICE64_2",
			"Strg": "msg: TRICE64_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"58503": {
			"Type": "trice32",
			"Strg": "rd:%d,%d\\n"
		},
		"58541": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d\\n"
		},
		"58642": {
			"Type": "TRICE32_3i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"58787": {
			"Type": "TRICE64",
			"Strg": "rd:%d, %d\\n"
		},
		"58845": {
			"Type": "trice32_3",
			"Strg": "msg: trice32_3  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"59167": {
			"Type": "TRICE8_3i",
			"Strg": "msg:  TRICE8_3i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"59225": {
			"Type": "TRICE16_1",
			"Strg": "TIM:timing      message, SysTick is %6u\\n"
		},
		"59229": {
			"Type": "TRICE32_3",
			"Strg": "tst:TRICE32_3 %u %u %u\\n"
		},
		"59253": {
			"Type": "TRICE8_1i",
			"Strg": "msg:  TRICE8_1i -\u003e normal trice macro    (with cycle counter) for only inside critical section %d\\n"
		},
		"59425": {
			"Type": "TRICE8_3",
			"Strg": "tst:TRICE8_3 %02X %02X %02X\\n"
		},
		"59450": {
			"Type": "TRICE8_8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"59556": {
			"Type": "TRICE8_7i",
			"Strg": "rd:%c%c%c%c%c%c%c"
		},
		"59595": {
			"Type": "trice32_1i",
			"Strg": "msg: trice32_1i -\u003e normal trice function (with cycle counter) for only inside critical section %d\\n"
		},
		"59628": {
			"Type": "trice32_4",
			"Strg": "msg: trice32_4  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"59727": {
			"Type": "TRICE8_2",
			"Strg": "msg: TRICE8_2   -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"59745": {
			"Type": "trice64",
			"Strg": "rd:%d\\n"
		},
		"59763": {
			"Type": "trice64_2",
			"Strg": "msg: trice64_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"59913": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"59946": {
			"Type": "trice64",
			"Strg": "rd:%d\\n"
		},
		"59976": {
			"Type": "trice16",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"6": {
			"Type": "Trice16_1",
			"Strg": "msg:%d\\n"
		},
		"60092": {
			"Type": "trice8_4",
			"Strg": "msg: trice8_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"60145": {
			"Type": "TRICE32_2i",
			"Strg": "msg: TRICE32_2i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X\\n"
		},
		"60229": {
			"Type": "trice16_2i",
			"Strg": "msg: trice16_2i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X\\n"
		},
		"60271": {
			"Type": "trice16_1",
			"Strg": "dbg:12345 as 16bit is %#016b\\n"
		},
		"60288": {
			"Type": "TRICE64_1",
			"Strg": "msg: TRICE64_1  -\u003e normal trice macro    (with cycle counter) for everywhere %u\\n"
		},
		"60551": {
			"Type": "trice16_1i",
			"Strg": "msg: trice16_1i -\u003e normal trice function (with cycle counter) for only inside critical section %d\\n"
		},
		"60735": {
			"Type": "TRICE16_3",
			"Strg": "tst:TRICE16_3 %u %u %u\\n"
		},
		"61024": {
			"Type": "TRICE16_1",
			"Strg": "RD:read         message, SysTick is %6u\\n"
		},
		"61043": {
			"Type": "trice8_3",
			"Strg": "msg:  trice8_3  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"61262": {
			"Type": "trice16_3",
			"Strg": "msg: trice16_3  -\u003e normal trice function  (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"61420": {
			"Type": "TRICE64_2i",
			"Strg": "rd:%c%c"
		},
		"61463": {
			"Type": "TRICE8_5i",
			"Strg": "rd:%d, %d, %d, %d, %d\\n"
		},
		"61470": {
			"Type": "trice16_1",
			"Strg": "msg: trice16_1  -\u003e normal trice function (with cycle counter) for everywhere %u\\n"
		},
		"61488": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4 %02x %02x %02x %02x\\n"
		},
		"61495": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4 %02X %02X %02X %02X\\n"
		},
		"61530": {
			"Type": "TRICE8_2",
			"Strg": "rd:%d, %d\\n"
		},
		"61588": {
			"Type": "TRICE32_2",
			"Strg": "msg: TRICE32_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n\\n"
		},
		"61677": {
			"Type": "trice8i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"61708": {
			"Type": "TRICE32_4",
			"Strg": "rd:%c%c%c%c"
		},
		"61896": {
			"Type": "trice8",
			"Strg": "rd:%d\\n"
		},
		"62001": {
			"Type": "TRICE64_2",
			"Strg": "msg: TRICE64_2  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X\\n"
		},
		"62064": {
			"Type": "trice16_1",
			"Strg": "ERR:error       message, SysTick is %6u\\n"
		},
		"62069": {
			"Type": "TRICE8_5",
			"Strg": "rd:%c%c%c%c%c"
		},
		"62081": {
			"Type": "TRICE8_2",
			"Strg": "msg: TRICE8_2   -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"6214": {
			"Type": "trice8_2",
			"Strg": "tst:trice8_2 %02X %02X\\n"
		},
		"62238": {
			"Type": "trice8_8",
			"Strg": "tst:trice8_8 %02X %02X %02X %02X %02X %02X %02X %02X\\n"
		},
		"62317": {
			"Type": "trice0",
			"Strg": "m:12"
		},
		"62334": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d\\n"
		},
		"62351": {
			"Type": "TRICE32_2",
			"Strg": "msg: TRICE32_2  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X\\n"
		},
		"62360": {
			"Type": "TRICE32",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"62387": {
			"Type": "trice8_6i",
			"Strg": "msg:  trice6_5i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u %d %x %X\\n"
		},
		"62466": {
			"Type": "trice8_3",
			"Strg": "msg:  trice8_3  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u\\n"
		},
		"62499": {
			"Type": "TRICE16_2i",
			"Strg": "msg: TRICE16_2i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X\\n"
		},
		"62535": {
			"Type": "TRICE16_1",
			"Strg": "ERR:error       message, SysTick is %6u\\n"
		},
		"62537": {
			"Type": "TRICE32",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"62542": {
			"Type": "TRICE16_4",
			"Strg": "msg: TRICE16_4  -\u003e normal trice macro (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"62644": {
			"Type": "TRICE32i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"62651": {
			"Type": "trice8",
			"Strg": "rd:trice8 line %d, %d\\n"
		},
		"62708": {
			"Type": "TRICE8_7i",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d\\n"
		},
		"62924": {
			"Type": "TRICE8_7i",
			"Strg": "msg:  TRICE8_7i -\u003e normal trice macro    (with cycle counter) for only inside critical section %x %X %u %d %x %X %u\\n"
		},
		"62974": {
			"Type": "TRICE16_3",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"62987": {
			"Type": "TRICE32_4i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"630": {
			"Type": "Trice8_1i",
			"Strg": "int:SysTick: %5d\\n"
		},
		"63098": {
			"Type": "TRICE64_1i",
			"Strg": "msg: TRICE64_1i -\u003e normal trice macro    (with cycle counter) for only inside critical section %d\\n"
		},
		"63167": {
			"Type": "trice8_3i",
			"Strg": "msg:  trice8_3i -\u003e normal trice function (with cycle counter) for only inside critical section %x %X %u\\n"
		},
		"63361": {
			"Type": "TRICE64_1",
			"Strg": "msg: TRICE64_1  -\u003e normal trice macro    (with cycle counter) for everywhere                   %u\\n"
		},
		"63466": {
			"Type": "trice0",
			"Strg": "d:G"
		},
		"63518": {
			"Type": "TRICE8i",
			"Strg": "rd:%d, %d, %d, %d\\n"
		},
		"63610": {
			"Type": "TRICE64",
			"Strg": "rd:%d,%d\\n"
		},
		"63660": {
			"Type": "TRICE8_3i",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"63747": {
			"Type": "TRICE8_4",
			"Strg": "msg:  TRICE8_4  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"63971": {
			"Type": "trice16_2",
			"Strg": "msg: Trice16_1 -\u003e short trice macro (no cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"64042": {
			"Type": "TRICE16",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"64054": {
			"Type": "trice8",
			"Strg": "rd:%d, %d, %d, %d, %d, %d, %d, %d\\n"
		},
		"64238": {
			"Type": "trice8",
			"Strg": "rd:trice8 line %d\\n"
		},
		"64268": {
			"Type": "trice16i",
			"Strg": "rd:%d, %d\\n"
		},
		"64296": {
			"Type": "trice32_4",
			"Strg": "msg: trice32_4  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X %u %d\\n"
		},
		"64447": {
			"Type": "trice64_1",
			"Strg": "msg: trice64_1  -\u003e normal trice function (with cycle counter) for everywhere                   %u\\n"
		},
		"64622": {
			"Type": "trice32_2",
			"Strg": "msg: trice32_2  -\u003e normal trice function (with cycle counter) for everywhere                   %x %X\\n"
		},
		"64745": {
			"Type": "trice32",
			"Strg": "rd:%d\\n"
		},
		"64784": {
			"Type": "TRICE8_8i",
			"Strg": "rd:%c%c%c%c%c%c%c%c"
		},
		"648": {
			"Type": "Trice16_1",
			"Strg": "INT:Trice16_1 SysTick: %5u\\n"
		},
		"64819": {
			"Type": "trice8_5",
			"Strg": "tst:trice8_5 %02X %02X %02X %02X %02X\\n"
		},
		"64846": {
			"Type": "TRICE32_1i",
			"Strg": "rd:%d\\n"
		},
		"64855": {
			"Type": "trice8_2",
			"Strg": "msg: trice8_2   -\u003e normal trice macro (with cycle counter) for everywhere 0x%X 0x%x\\n"
		},
		"64908": {
			"Type": "trice32",
			"Strg": "rd:%d, %d, %d\\n"
		},
		"64938": {
			"Type": "trice0",
			"Strg": "--------------------------------------------------\\n"
		},
		"64951": {
			"Type": "TRICE8_8",
			"Strg": "msg:  TRICE8_8  -\u003e normal trice macro    (with cycle counter) for everywhere                   %x %X %u %d %x %X %u %d\\n"
		},
		"65002": {
			"Type": "TRICE0",
			"Strg": "rd_:triceFifo.c"
		},
		"65003": {
			"Type": "trice64_1",
			"Strg": "tst:trice64_1 %x\\n"
		},
		"65005": {
			"Type": "TRICE0",
			"Strg": "m:12"
		},
		"65010": {
			"Type": "TRICE0",
			"Strg": "sig:This ASSERT error is just a demo and no real error:\\n"
		},
		"65013": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4 %d %d %d %d\\n"
		},
		"65017": {
			"Type": "TRICE16_3",
			"Strg": "tst:TRICE16_3 %d %d %d\\n"
		},
		"65021": {
			"Type": "TRICE32_2",
			"Strg": "tst:TRICE32_2 %d %d\\n"
		},
		"65023": {
			"Type": "TRICE32_2",
			"Strg": "sig:%2d:%6d\\n"
		},
		"65029": {
			"Type": "TRICE0",
			"Strg": "--------------------------------------------------\\n\\n"
		},
		"65031": {
			"Type": "trice8_7",
			"Strg": "tst:trice8_7 %d %d %d %d %d %d %d\\n"
		},
		"65041": {
			"Type": "TRICE0",
			"Strg": "time:i"
		},
		"65042": {
			"Type": "TRICE0",
			"Strg": "w:B"
		},
		"65044": {
			"Type": "trice16_1",
			"Strg": "WR:write        message, SysTick is %6d\\n"
		},
		"65046": {
			"Type": "TRICE0",
			"Strg": "e:7"
		},
		"65048": {
			"Type": "TRICE32_1",
			"Strg": "tst:TRICE32_1 %08x\\n"
		},
		"65049": {
			"Type": "TRICE16_2",
			"Strg": "msg: TRICE16_2  -\u003e normal trice macro     (with cycle counter) for everywhere                   %x %X\\n"
		},
		"65051": {
			"Type": "TRICE16_1",
			"Strg": "tim: post decryption SysTick=%d\\n"
		},
		"65052": {
			"Type": "TRICE8_4",
			"Strg": "%c%c%c%c"
		},
		"65054": {
			"Type": "TRICE16_1",
			"Strg": "tim: pre encryption SysTick=%d\\n"
		},
		"65055": {
			"Type": "TRICE16_4",
			"Strg": "tst:TRICE16_4  %%05x -\u003e   %05x   %05x   %05x   %05x\\n"
		},
		"65057": {
			"Type": "TRICE8_3",
			"Strg": "%c%c%c"
		},
		"65060": {
			"Type": "TRICE16_4",
			"Strg": "tst:TRICE16_4   %%7o -\u003e %7o %7o %7o %7o\\n"
		},
		"65061": {
			"Type": "TRICE16_4",
			"Strg": "att: %d,%d,%d,%d\\n"
		},
		"65064": {
			"Type": "TRICE0",
			"Strg": "rd_:triceBareFifoToEscFifo.c"
		},
		"65066": {
			"Type": "TRICE32_4",
			"Strg": "tst:TRICE32_4 %%10d -\u003e     %10d     %10d     %10d    %10x\\n"
		},
		"65073": {
			"Type": "TRICE32_3",
			"Strg": "tst:TRICE32_3 %x %x %x\\n"
		},
		"65077": {
			"Type": "TRICE16_4",
			"Strg": "tst:TRICE16_4   %%6d -\u003e  %6d  %6d  %6d  %6d\\n"
		},
		"65083": {
			"Type": "TRICE8_8",
			"Strg": "msg: message = %03x %03x %03x %03x %03x %03x %03x %03x\\n"
		},
		"65088": {
			"Type": "TRICE8_5",
			"Strg": "%c%c%c%c%c"
		},
		"65089": {
			"Type": "trice16_2",
			"Strg": "tst:trice16_2 %d %d\\n"
		},
		"65090": {
			"Type": "TRICE0",
			"Strg": "diag:f"
		},
		"65094": {
			"Type": "trice32_1",
			"Strg": "tst:trice32_1 %08x\\n"
		},
		"65099": {
			"Type": "trice16_1",
			"Strg": "INFO:informal   message, SysTick is %6d\\n"
		},
		"65103": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4   %%4o -\u003e %4o %4o %4o %4o\\n"
		},
		"65105": {
			"Type": "trice16_1",
			"Strg": "ATT:attention   message, SysTick is %6d\\n"
		},
		"65107": {
			"Type": "trice8",
			"Strg": "rd:%d, %d\\n"
		},
		"65109": {
			"Type": "TRICE0",
			"Strg": "3"
		},
		"65112": {
			"Type": "trice8_1",
			"Strg": "tst:trice8_1 %d\\n"
		},
		"65115": {
			"Type": "trice8_3",
			"Strg": "tst:trice8_3 %d %d %d\\n"
		},
		"65117": {
			"Type": "TRICE16_4",
			"Strg": "tst:TRICE16_4 %d %d %d %d\\n"
		},
		"65118": {
			"Type": "TRICE0",
			"Strg": "rd_:triceFifoToBytesBuffer.c"
		},
		"65121": {
			"Type": "TRICE8_7",
			"Strg": "%c%c%c%c%c%c%c"
		},
		"65132": {
			"Type": "TRICE32_3",
			"Strg": "tst:TRICE32_3 %d %d %d\\n"
		},
		"65137": {
			"Type": "trice8_4",
			"Strg": "tst:trice8_4   %%4o -\u003e %4o %4o %4o %4o\\n"
		},
		"65143": {
			"Type": "trice32_4",
			"Strg": "tst:trice32_4 %x %x %x %x\\n"
		},
		"65144": {
			"Type": "trice16_4",
			"Strg": "tst:trice16_4 %d %d %d %d\\n"
		},
		"65159": {
			"Type": "trice8_4",
			"Strg": "tst:trice8_4 %d %d %d %d\\n"
		},
		"65161": {
			"Type": "TRICE0",
			"Strg": "2"
		},
		"65167": {
			"Type": "TRICE0",
			"Strg": "a:c"
		},
		"65168": {
			"Type": "TRICE16_1",
			"Strg": "dbg:12345 as 16bit is %#016b\\n"
		},
		"65184": {
			"Type": "TRICE0",
			"Strg": "e:A"
		},
		"65201": {
			"Type": "TRICE0",
			"Strg": "d:G"
		},
		"65211": {
			"Type": "trice16_1",
			"Strg": "SIG:signal      message, SysTick is %6d\\n"
		},
		"65213": {
			"Type": "TRICE_S",
			"Strg": "%s\\n"
		},
		"65219": {
			"Type": "TRICE0",
			"Strg": "4"
		},
		"65228": {
			"Type": "TRICE0",
			"Strg": "1"
		},
		"65235": {
			"Type": "TRICE0",
			"Strg": "wrn:TRICES_1(id, pFmt, dynString) macro is not supported in bare encoding.\\nmsg:See TRICE_RTS macro in triceCheck.c for an alternative or use a different encoding.\\n"
		},
		"65236": {
			"Type": "trice16_1",
			"Strg": "ERR:error       message, SysTick is %6d\\n"
		},
		"65239": {
			"Type": "TRICE8_5",
			"Strg": "tst:TRICE8_5 %d %d %d %d %d\\n"
		},
		"65246": {
			"Type": "TRICE8_7",
			"Strg": "tst:TRICE8_7 %d %d %d %d %d %d %d\\n"
		},
		"65251": {
			"Type": "trice64_1",
			"Strg": "att:trice64_1 %#b\\n"
		},
		"65254": {
			"Type": "TRICE0",
			"Strg": "rd_:triceCheck.c"
		},
		"65262": {
			"Type": "trice64_2",
			"Strg": "tst:trice64_2 %x %x\\n"
		},
		"65264": {
			"Type": "TRICE8_8",
			"Strg": "tst:TRICE8_8 %d %d %d %d %d %d %d %d\\n"
		},
		"65274": {
			"Type": "TRICE16_1",
			"Strg": "tst:TRICE16_1   message, SysTick is %6d\\n"
		},
		"65279": {
			"Type": "TRICE8_2",
			"Strg": "%c%c"
		},
		"65281": {
			"Type": "TRICE8_8",
			"Strg": "msg: messge = %03x %03x %03x %03x %03x %03x %03x %03x\\n"
		},
		"65283": {
			"Type": "TRICE32_4",
			"Strg": "tst:TRICE32_4 %x %x %x %x\\n"
		},
		"65287": {
			"Type": "TRICE32_1",
			"Strg": "tst:TRICE32_1   message, SysTick is %6d\\n"
		},
		"65299": {
			"Type": "trice64_1",
			"Strg": "tst:trice64_1 %d\\n"
		},
		"65300": {
			"Type": "trice32_1",
			"Strg": "tst:trice32_1 %d\\n"
		},
		"65304": {
			"Type": "TRICE0",
			"Strg": "--------------------------------------------------\\n"
		},
		"65305": {
			"Type": "trice32_4",
			"Strg": "tst:trice32_4 %%10d -\u003e     %10d     %10d     %10d    %10x\\n"
		},
		"65308": {
			"Type": "TRICE8_2",
			"Strg": "tst:TRICE8_2 %d %d\\n"
		},
		"65309": {
			"Type": "trice16_1",
			"Strg": "DBG:debug       message, SysTick is %6d\\n"
		},
		"65312": {
			"Type": "trice16_1",
			"Strg": "DIA:diagnostics message, SysTick is %6d\\n"
		},
		"65314": {
			"Type": "TRICE0",
			"Strg": "dbg:k\\n"
		},
		"65318": {
			"Type": "TRICE32_2",
			"Strg": "tst:TRICE32_2 %x %x\\n"
		},
		"65329": {
			"Type": "TRICE8_1",
			"Strg": "%c"
		},
		"65330": {
			"Type": "TRICE16_1",
			"Strg": "tim: post encryption SysTick=%d\\n"
		},
		"65331": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4   %%4d -\u003e %4d %4d %4d %4d\\n"
		},
		"65336": {
			"Type": "TRICE0",
			"Strg": "msg: TRICE0     -\u003e normal trice macro    (with cycle counter) for everywhere\\n"
		},
		"65344": {
			"Type": "TRICE0",
			"Strg": "message:J"
		},
		"65364": {
			"Type": "trice16_1",
			"Strg": "MSG:normal      message, SysTick is %6d\\n"
		},
		"65367": {
			"Type": "TRICE0",
			"Strg": "rd:e\\n"
		},
		"65369": {
			"Type": "trice32_4",
			"Strg": "tst:trice32_4 %d %d %d %d\\n"
		},
		"65370": {
			"Type": "trice16_1",
			"Strg": "tst:trice16_1   message, SysTick is %6d\\n"
		},
		"65372": {
			"Type": "TRICE8_6",
			"Strg": "tst:TRICE8_6 %d %d %d %d %d %d\\n"
		},
		"65385": {
			"Type": "TRICE0",
			"Strg": "m:123\\n"
		},
		"65388": {
			"Type": "trice16_1",
			"Strg": "WRN:warning     message, SysTick is %6d\\n"
		},
		"65391": {
			"Type": "TRICE64_2",
			"Strg": "tst:TRICE64_2 %d %d\\n"
		},
		"65396": {
			"Type": "TRICE64_1",
			"Strg": "att:TRICE64_1 %#b\\n"
		},
		"65400": {
			"Type": "trice8_4",
			"Strg": "tst:trice8_4   %%4d -\u003e %4d %4d %4d %4d\\n"
		},
		"65405": {
			"Type": "TRICE16_1",
			"Strg": "tim: pre decryption SysTick=%d\\n"
		},
		"65406": {
			"Type": "trice64_2",
			"Strg": "tst:trice64_2 %d %d\\n"
		},
		"65409": {
			"Type": "TRICE8_1",
			"Strg": "tst:TRICE8_1 %d\\n"
		},
		"65412": {
			"Type": "TRICE16_2",
			"Strg": "tst:TRICE16_2 %d %d\\n"
		},
		"65416": {
			"Type": "TRICE16_1",
			"Strg": "tst:TRICE16_1 %d\\n"
		},
		"65418": {
			"Type": "trice16_4",
			"Strg": "tst:trice16_4   %%7o -\u003e %7o %7o %7o %7o\\n"
		},
		"65422": {
			"Type": "TRICE0",
			"Strg": "s:                                                   \\ns:                     myProject                     \\ns:                                                   \\n\\n"
		},
		"65424": {
			"Type": "trice16_1",
			"Strg": "tst:trice16_1 %d\\n"
		},
		"65428": {
			"Type": "trice32_3",
			"Strg": "tst:trice32_3 %x %x %x\\n"
		},
		"65437": {
			"Type": "trice16_3",
			"Strg": "tst:trice16_3 %d %d %d\\n"
		},
		"65438": {
			"Type": "trice16_1",
			"Strg": "RD:read         message, SysTick is %6d\\n"
		},
		"65439": {
			"Type": "TRICE16_1",
			"Strg": "TIM:timing      message, SysTick is %6d\\n"
		},
		"65442": {
			"Type": "trice8_2",
			"Strg": "tst:trice8_2 %d %d\\n"
		},
		"65443": {
			"Type": "TRICE0",
			"Strg": "wr:d"
		},
		"65447": {
			"Type": "trice16_4",
			"Strg": "tst:trice16_4  %%05x -\u003e   %05x   %05x   %05x   %05x\\n"
		},
		"65450": {
			"Type": "TRICE32_4",
			"Strg": "tst:TRICE32_4 %d %d %d %d\\n"
		},
		"65454": {
			"Type": "TRICE8_3",
			"Strg": "tst:TRICE8_3 %d %d %d\\n"
		},
		"65459": {
			"Type": "trice0",
			"Strg": "4"
		},
		"65460": {
			"Type": "TRICE8_8",
			"Strg": "tst:TRICE8_8 %02x %02x %02x %02x %02x %02x %02x %02x\\n"
		},
		"65462": {
			"Type": "trice32_1",
			"Strg": "tst:trice32_1   message, SysTick is %6d\\n"
		},
		"65463": {
			"Type": "trice16_1",
			"Strg": "ISR:interrupt   message, SysTick is %6d\\n"
		},
		"65468": {
			"Type": "TRICE8_8",
			"Strg": "%c%c%c%c%c%c%c%c"
		},
		"65473": {
			"Type": "TRICE8_6",
			"Strg": "%c%c%c%c%c%c"
		},
		"65484": {
			"Type": "TRICE8_8",
			"Strg": "att: encrypted = %03x %03x %03x %03x %03x %03x %03x %03x\\n"
		},
		"65485": {
			"Type": "trice32_2",
			"Strg": "tst:trice32_2 %d %d\\n"
		},
		"65492": {
			"Type": "TRICE8_4",
			"Strg": "tst:TRICE8_4  %%03x -\u003e  %03x  %03x  %03x  %03x\\n"
		},
		"65493": {
			"Type": "trice8_4",
			"Strg": "tst:trice8_4  %%03x -\u003e  %03x  %03x  %03x  %03x\\n"
		},
		"65495": {
			"Type": "trice8_6",
			"Strg": "tst:trice8_6 %d %d %d %d %d %d\\n"
		},
		"65498": {
			"Type": "trice16_4",
			"Strg": "tst:trice16_4   %%6d -\u003e  %6d  %6d  %6d  %6d\\n"
		},
		"65503": {
			"Type": "trice8_5",
			"Strg": "tst:trice8_5 %d %d %d %d %d\\n"
		},
		"65507": {
			"Type": "trice32_4",
			"Strg": "tst:trice32_4 %%09x -\u003e      %09x      %09x       %09x     %09x\\n"
		},
		"65509": {
			"Type": "TRICE8_2",
			"Strg": "tst:TRICE8_2 %02X %02X\\n"
		},
		"65510": {
			"Type": "TRICE32_4",
			"Strg": "tst:TRICE32_4 %%09x -\u003e      %09x      %09x       %09x     %09x\\n"
		},
		"65517": {
			"Type": "TRICE0",
			"Strg": "t:H"
		},
		"65523": {
			"Type": "trice8_8",
			"Strg": "tst:trice8_8 %d %d %d %d %d %d %d %d\\n"
		},
		"65525": {
			"Type": "trice32_2",
			"Strg": "tst:trice32_2 %x %x\\n"
		},
		"65526": {
			"Type": "TRICE16_4",
			"Strg": "att: encrypted = %d,%d,%d,%d,"
		},
		"65528": {
			"Type": "TRICE64_1",
			"Strg": "tst:TRICE64_1 %d\\n"
		},
		"65529": {
			"Type": "TRICE32_1",
			"Strg": "tst:TRICE32_1 %d\\n"
		},
		"65533": {
			"Type": "trice32_3",
			"Strg": "tst:trice32_3 %d %d %d\\n"
		},
		"6721": {
			"Type": "Trice8_1i",
			"Strg": "msg: Trice8_1i  -\u003e short  trice macro    (no   cycle counter) for only inside critical section %d\\n"
		},
		"6726": {
			"Type": "Trice16",
			"Strg": "ATT:attention   message, SysTick is %6u\\n"
		},
		"6737": {
			"Type": "TRICE8_1",
			"Strg": "tst:TRICE8_1 %02X\\n"
		},
		"7": {
			"Type": "Trice8_2",
			"Strg": "msg:%x, %x\\n"
		},
		"7101": {
			"Type": "Trice8_1i",
			"Strg": "msg: Trice8_1i -\u003e short trice macro (no cycle counter) for only inside critical section %d\\n"
		},
		"7132": {
			"Type": "Trice16i",
			"Strg": "ATT:Trice16i attention   message, SysTick is %6u\\n"
		},
		"7279": {
			"Type": "Trice8_1",
			"Strg": "msg: Trice8_1   -\u003e short  trice macro    (no   cycle counter) for everywhere                   %u\\n"
		},
		"7431": {
			"Type": "Trice8_1",
			"Strg": "msg: Trice8_1  -\u003e short  trice macro    (no   cycle counter) for everywhere %u\\n"
		},
		"7473": {
			"Type": "TRICE8_2",
			"Strg": "tst:TRICE8_2 %u %u\\n"
		},
		"7500": {
			"Type": "Trice0i",
			"Strg": "INT:SysTick_Handler\\n"
		},
		"7815": {
			"Type": "Trice8_1",
			"Strg": "msg: Trice8_1   -\u003e short  trice macro    (no   cycle counter) for everywhere %u\\n"
		},
		"7908": {
			"Type": "Trice8_2",
			"Strg": "diag: Trice8_2  %x, %x\\n"
		},
		"7946": {
			"Type": "Trice8",
			"Strg": "rd:%d, %d\\n"
		},
		"796": {
			"Type": "Trice8_1i",
			"Strg": "msg: Trice8_1i -\u003e short trice macro (no cycle counter) for only inside critical section\\n"
		},
		"8": {
			"Type": "Trice16_1",
			"Strg": "msg:%6u\\n"
		},
		"8080": {
			"Type": "Trice16_1",
			"Strg": "rd: Trice16_1 %u\\n"
		},
		"8101": {
			"Type": "TRICE16_1",
			"Strg": "DBG:debug       message, SysTick is %6u\\n"
		},
		"8908": {
			"Type": "Trice0",
			"Strg": "att:SysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_Handler"
		},
		"8915": {
			"Type": "Trice0",
			"Strg": "att:SysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_HandlerSysTick_Handler"
		},
		"9484": {
			"Type": "Trice8_1",
			"Strg": "msg:  Trice8_1  -\u003e short  trice macro    (no   cycle counter) for everywhere                   %u\\n"
		},
		"9549": {
			"Type": "Trice16_1i",
			"Strg": "msg: Trice16_1i %d\\n"
		},
		"983": {
			"Type": "Trice8_2i",
			"Strg": "INT:%d %d\\n"
		}
	}
`
)

var glob *sync.RWMutex // tests changing global values need to exclude each other

func init() {
	glob = new(sync.RWMutex)
}

func _TestTranslate(t *testing.T) {
	glob.Lock()
	defer glob.Unlock()
	sw := emitter.New(os.Stdout)
	lu := make(id.TriceIDLookUp) // empty
	assert.Nil(t, lu.FromJSON([]byte(til)))
	m := new(sync.RWMutex) // m is a pointer to a read write mutex for lu
	Encoding = "FLEX"
	TargetEndianness = "littleEndian"
	receiver.Port = "BUFFER"
	ShowID = "%d"
	defer func() {
		ShowID = "" // reset to default
	}()
	verbose := true
	rc, err := receiver.NewReadCloser(os.Stdout, verbose, receiver.Port, "2, 124, 227, 255, 0, 0, 4, 0")
	assert.Nil(t, err)
	err = Translate(os.Stdout, sw, lu, m, rc)
	assert.Equal(t, io.EOF, err)
}
*/
