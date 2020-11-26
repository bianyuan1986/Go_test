package main

import (
	"fmt"
)

type Tiger struct {
	name string
}

type Tiger1 struct {
	name string
}

type I interface {
	Dump()
}

func (t *Tiger) GetName() {
	fmt.Println(t.name)
}

/*
func (t *Tiger) Dump() {
	fmt.Println(t.name)
}
*/

func (t Tiger) Dump() {
	fmt.Println(t.name)
}

func checkInterfaceImplementI(object interface{}) {
	if _, ok := object.(I); !ok {
		fmt.Println("Type doesn't implement interface I")
		return
	}
	fmt.Println("Type implement interface I")
}

func _checkInterfaceImplementI(object *interface{}) {
	if _, ok := (*object).(I); !ok {
		fmt.Println("Type doesn't implement interface I")
		return
	}
	fmt.Println("Type implement interface I")
}

/*保存某个对象实例的方法，记录方法地址的同时关联了对象*/
func main() {
	data := make(map[string]func())
	a := &Tiger{name: "Tiger1"}
	b := &Tiger{name: "Tiger2"}
	data["t1"] = a.GetName
	data["t2"] = b.GetName

	for _, v := range data {
		v()
	}
	t := Tiger{"2020"}
	t1 := Tiger1{"2021"}
	fmt.Println("Check struct [Tiger]")
	checkInterfaceImplementI(t)
	fmt.Println("Check struct [Tiger1]")
	checkInterfaceImplementI(t1)
	fmt.Println("Check struct [Tiger]")
	checkInterfaceImplementI(&t)
	fmt.Println("Check struct [Tiger1]")
	checkInterfaceImplementI(&t1)
}

/*
接口方法实现的接收者为指针类型，当参数为interface{}类型时如果通过值传递则
判断当前对象未实现该接口，如果传递指针类型则判断实现该接口
func (t *Tiger) Dump() {
	fmt.Println(t.name)
}
Tiger1
Tiger2
Check struct [Tiger]
Type doesn't implement interface I
Check struct [Tiger1]
Type doesn't implement interface I
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I

接口方法实现的接收者为对象类型，当参数为interface{}类型时如果通过值传递或者
通过指针传递则都判断实现该接口
func (t Tiger) Dump() {
	fmt.Println(t.name)
}
Tiger1
Tiger2
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I

Go语言的函数调用通过栈传递参数而不是通过寄存器传递参数
(gdb) break method.go:63
Breakpoint 1 at 0x491e3e: file /Users/bianyuan/Desktop/workspace/test/Go_test/method.go, line 63.
(gdb) run
Starting program: /home/baohu.rbh/method
Tiger1
Tiger2
Check struct [Tiger]

Breakpoint 1, main.main () at /Users/bianyuan/Desktop/workspace/test/Go_test/method.go:63
63	/Users/bianyuan/Desktop/workspace/test/Go_test/method.go: No such file or directory.
(gdb) disassemble
Dump of assembler code for function main.main:
   0x0000000000491d79 <+857>:	jmp    0x491d30 <main.main+784>
   0x0000000000491d7b <+859>:	xorps  %xmm0,%xmm0
   0x0000000000491d7e <+862>:	movups %xmm0,0xe0(%rsp)
   0x0000000000491d86 <+870>:	lea    0x2f6bc(%rip),%rax        # 0x4c1449
   0x0000000000491d8d <+877>:	mov    %rax,0xe0(%rsp)
   0x0000000000491d95 <+885>:	movq   $0x4,0xe8(%rsp)
   0x0000000000491da1 <+897>:	xorps  %xmm0,%xmm0
   0x0000000000491da4 <+900>:	movups %xmm0,0xd0(%rsp)
   0x0000000000491dac <+908>:	lea    0x2f69a(%rip),%rax        # 0x4c144d
   0x0000000000491db3 <+915>:	mov    %rax,0xd0(%rsp)
   0x0000000000491dbb <+923>:	movq   $0x4,0xd8(%rsp)
   0x0000000000491dc7 <+935>:	xorps  %xmm0,%xmm0
   0x0000000000491dca <+938>:	movups %xmm0,0xf0(%rsp)
   0x0000000000491dd2 <+946>:	lea    0xf0(%rsp),%rax
   0x0000000000491dda <+954>:	mov    %rax,0x78(%rsp)
   0x0000000000491ddf <+959>:	test   %al,(%rax)
   0x0000000000491de1 <+961>:	lea    0xef98(%rip),%rcx        # 0x4a0d80
   0x0000000000491de8 <+968>:	mov    %rcx,0xf0(%rsp)
   0x0000000000491df0 <+976>:	lea    0x49c39(%rip),%rcx        # 0x4dba30
   0x0000000000491df7 <+983>:	mov    %rcx,0xf8(%rsp)
   0x0000000000491dff <+991>:	test   %al,(%rax)
   0x0000000000491e01 <+993>:	jmp    0x491e03 <main.main+995>
   0x0000000000491e03 <+995>:	mov    %rax,0x168(%rsp)
   0x0000000000491e0b <+1003>:	movq   $0x1,0x170(%rsp)
   0x0000000000491e17 <+1015>:	movq   $0x1,0x178(%rsp)
   0x0000000000491e23 <+1027>:	mov    %rax,(%rsp)
   0x0000000000491e27 <+1031>:	movq   $0x1,0x8(%rsp)
   0x0000000000491e30 <+1040>:	movq   $0x1,0x10(%rsp)
   0x0000000000491e39 <+1049>:	callq  0x48b190 <fmt.Println>   //调用fmt.Println
=> 0x0000000000491e3e <+1054>:	mov    0xe8(%rsp),%rax          //把t对象的字符串地址和字符串长度保存到寄存器中
   0x0000000000491e46 <+1062>:	mov    0xe0(%rsp),%rcx
   0x0000000000491e4e <+1070>:	mov    %rcx,0x110(%rsp)         //把寄存器中地址复制到栈中其它地址处，相当于拷贝了对象
   0x0000000000491e56 <+1078>:	mov    %rax,0x118(%rsp)
   0x0000000000491e5e <+1086>:	lea    0x1a97b(%rip),%rax        # 0x4ac7e0
   0x0000000000491e65 <+1093>:	mov    %rax,(%rsp)
   0x0000000000491e69 <+1097>:	lea    0x110(%rsp),%rax         //把复制后的对象地址存入rax寄存器然后入栈作为参数传递给函数
   0x0000000000491e71 <+1105>:	mov    %rax,0x8(%rsp)
   0x0000000000491e76 <+1110>:	callq  0x4918c0 <main.checkInterfaceImplementI>
   0x0000000000491e7b <+1115>:	xorps  %xmm0,%xmm0
   0x0000000000491e7e <+1118>:	movups %xmm0,0xf0(%rsp)
   0x0000000000491e86 <+1126>:	lea    0xf0(%rsp),%rax
   0x0000000000491e8e <+1134>:	mov    %rax,0x70(%rsp)
   0x0000000000491e93 <+1139>:	test   %al,(%rax)
   0x0000000000491e95 <+1141>:	lea    0xeee4(%rip),%rcx        # 0x4a0d80
   0x0000000000491e9c <+1148>:	mov    %rcx,0xf0(%rsp)
   0x0000000000491ea4 <+1156>:	lea    0x49b95(%rip),%rcx        # 0x4dba40
   0x0000000000491eab <+1163>:	mov    %rcx,0xf8(%rsp)
   0x0000000000491eb3 <+1171>:	test   %al,(%rax)
   0x0000000000491eb5 <+1173>:	jmp    0x491eb7 <main.main+1175>
   0x0000000000491eb7 <+1175>:	mov    %rax,0x150(%rsp)
   0x0000000000491ebf <+1183>:	movq   $0x1,0x158(%rsp)
   0x0000000000491ecb <+1195>:	movq   $0x1,0x160(%rsp)
   0x0000000000491ed7 <+1207>:	mov    %rax,(%rsp)
   0x0000000000491edb <+1211>:	movq   $0x1,0x8(%rsp)
   0x0000000000491ee4 <+1220>:	movq   $0x1,0x10(%rsp)
   .................
   0x0000000000491fa1 <+1409>:	callq  0x48b190 <fmt.Println>
   0x0000000000491fa6 <+1414>:	lea    0xe0(%rsp),%rax     //如果传递的是对象地址，则此处没有复制对象的过程，此处直接把对象地址存入rax然后入栈
   0x0000000000491fae <+1422>:	mov    %rax,0xc8(%rsp)
   0x0000000000491fb6 <+1430>:	lea    0x166c3(%rip),%rcx        # 0x4a8680
   0x0000000000491fbd <+1437>:	mov    %rcx,(%rsp)
   0x0000000000491fc1 <+1441>:	mov    %rax,0x8(%rsp)      //对象地址入栈
   0x0000000000491fc6 <+1446>:	callq  0x4918c0 <main.checkInterfaceImplementI>
Quit
(gdb) break *0x491e76
Breakpoint 2 at 0x491e76: file /Users/bianyuan/Desktop/workspace/test/Go_test/method.go, line 63.
(gdb) c
Continuing.

Breakpoint 2, main.main () at /Users/bianyuan/Desktop/workspace/test/Go_test/method.go:63
63	in /Users/bianyuan/Desktop/workspace/test/Go_test/method.go
(gdb) print /x $rsp
$1 = 0xc000110c98
(gdb) x /2xg $1   //查看入栈的两个参数，第二个参数为拷贝的对象的地址0xc000110da8
0xc000110c98:	0x00000000004ac7e0	0x000000c000110da8
(gdb) print t
$2 = {name = 0x4c1449 "2020"}
(gdb) print t1
$3 = {name = 0x4c144d "2021"}
(gdb) print &t  //原始对象地址
$4 = (main.Tiger *) 0xc000110d78
(gdb) print &t1
$5 = (main.Tiger1 *) 0xc000110d68
(gdb) x /2xg $4  //原始对象中数据
0xc000110d78:	0x00000000004c1449	0x0000000000000004
(gdb) x /2xg 0x000000c000110da8 //拷贝对象的数据
0xc000110da8:	0x00000000004c1449	0x0000000000000004
(gdb) print /x $rsp+0xe0 //原始对象地址
$6 = 0xc000110d78
*/
