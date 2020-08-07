# command-line-arguments
"".main STEXT size=74 args=0x0 locals=0x10
	0x0000 00000 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	TEXT	"".main(SB), ABIInternal, $16-0
	0x0000 00000 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$0, $-2
	0x000d 00013 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	JLS	67
	0x000f 00015 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$0, $-1
	0x000f 00015 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	SUBQ	$16, SP
	0x0013 00019 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	MOVQ	BP, 8(SP)
	0x0018 00024 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	LEAQ	8(SP), BP
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$0, $-2
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$1, $-2
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	PCDATA	$0, $0
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	PCDATA	$1, $0
	0x001d 00029 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	CALL	runtime.printlock(SB)
	0x0022 00034 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	MOVQ	$3, (SP)
	0x002a 00042 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	CALL	runtime.printint(SB)
	0x002f 00047 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	CALL	runtime.printnl(SB)
	0x0034 00052 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:4)	CALL	runtime.printunlock(SB)
	0x0039 00057 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:5)	MOVQ	8(SP), BP
	0x003e 00062 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:5)	ADDQ	$16, SP
	0x0042 00066 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:5)	RET
	0x0043 00067 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:5)	NOP
	0x0043 00067 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$1, $-1
	0x0043 00067 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$0, $-2
	0x0043 00067 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	PCDATA	$0, $-1
	0x0048 00072 (/Users/bianyuan/Desktop/workspace/test/Go_test/x.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 34 48  eH..%....H;a.v4H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 c7 04 24 03 00 00 00 e8 00 00 00 00 e8  ..H..$..........
	0x0030 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 08 48 83  .........H.l$.H.
	0x0040 c4 10 c3 e8 00 00 00 00 eb b6                    ..........
	rel 5+4 t=17 TLS+0
	rel 30+4 t=8 runtime.printlock+0
	rel 43+4 t=8 runtime.printint+0
	rel 48+4 t=8 runtime.printnl+0
	rel 53+4 t=8 runtime.printunlock+0
	rel 68+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=35
	0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 00                                         ...
	rel 11+8 t=1 "".main+0
	rel 19+8 t=1 "".main+74
	rel 29+4 t=30 gofile../Users/bianyuan/Desktop/workspace/test/Go_test/x.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=16
	0x0000 04 02 11 0a a5 9c 06 41 06 f6 71 04 01 03 7e 01  .......A..q...~.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
