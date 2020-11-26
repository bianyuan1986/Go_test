package main

import (
	"fmt"
)

type A struct {
	name string
}

type B struct {
	member A
}

func (A) Amethod() {
	fmt.Println("A method!")
}

func testInterface(v *interface{}) {
	d := (*v).(A)
	d.Amethod()
}

func main() {
	var a = A{name: "A name"}
	//B{a}.member.Amethod()
	fmt.Println("Test Interface")
	var tmp interface{} = a
	testInterface(&tmp)
}

/*
(gdb) disassemble
Dump of assembler code for function main.main:
   0x0000000000491880 <+0>:	mov    %fs:0xfffffffffffffff8,%rcx
   0x0000000000491889 <+9>:	lea    -0x18(%rsp),%rax
   0x000000000049188e <+14>:	cmp    0x10(%rcx),%rax
   0x0000000000491892 <+18>:	jbe    0x49197c <main.main+252>
   0x0000000000491898 <+24>:	sub    $0x98,%rsp
   0x000000000049189f <+31>:	mov    %rbp,0x90(%rsp)
   0x00000000004918a7 <+39>:	lea    0x90(%rsp),%rbp
   0x00000000004918af <+47>:	xorps  %xmm0,%xmm0
   0x00000000004918b2 <+50>:	movups %xmm0,0x48(%rsp)
   0x00000000004918b7 <+55>:	lea    0x2e56d(%rip),%rax        # 0x4bfe2b
   0x00000000004918be <+62>:	mov    %rax,0x48(%rsp)
   0x00000000004918c3 <+67>:	movq   $0x6,0x50(%rsp)
   0x00000000004918cc <+76>:	xorps  %xmm0,%xmm0
   0x00000000004918cf <+79>:	movups %xmm0,0x68(%rsp)
   0x00000000004918d4 <+84>:	lea    0x68(%rsp),%rax
   0x00000000004918d9 <+89>:	mov    %rax,0x30(%rsp)
   0x00000000004918de <+94>:	test   %al,(%rax)
   0x00000000004918e0 <+96>:	lea    0xe259(%rip),%rcx        # 0x49fb40
   0x00000000004918e7 <+103>:	mov    %rcx,0x68(%rsp)
   0x00000000004918ec <+108>:	lea    0x4870d(%rip),%rcx        # 0x4da000
   0x00000000004918f3 <+115>:	mov    %rcx,0x70(%rsp)
   0x00000000004918f8 <+120>:	test   %al,(%rax)
   0x00000000004918fa <+122>:	jmp    0x4918fc <main.main+124>
   0x00000000004918fc <+124>:	mov    %rax,0x78(%rsp)
   0x0000000000491901 <+129>:	movq   $0x1,0x80(%rsp)
   0x000000000049190d <+141>:	movq   $0x1,0x88(%rsp)
   0x0000000000491919 <+153>:	mov    %rax,(%rsp)
   0x000000000049191d <+157>:	movq   $0x1,0x8(%rsp)
   0x0000000000491926 <+166>:	movq   $0x1,0x10(%rsp)
   0x000000000049192f <+175>:	callq  0x48b190 <fmt.Println>
=> 0x0000000000491934 <+180>:	mov    0x50(%rsp),%rax   //字符串对象的长度
   0x0000000000491939 <+185>:	mov    0x48(%rsp),%rcx   //字符串地址
   0x000000000049193e <+190>:	mov    %rcx,0x58(%rsp)   //拷贝字符串地址到栈空间
   0x0000000000491943 <+195>:	mov    %rax,0x60(%rsp)   //拷贝字符串长度
   0x0000000000491948 <+200>:	lea    0x199f1(%rip),%rax        # 0x4ab340
   0x000000000049194f <+207>:	mov    %rax,0x38(%rsp)  //全局变量地址，第一个参数，存放到栈$rsp+0x38处，tmp变量地址入栈？
   0x0000000000491954 <+212>:	lea    0x58(%rsp),%rax  //传递拷贝对象首地址
   0x0000000000491959 <+217>:	mov    %rax,0x40(%rsp)  //把拷贝的对象的首地址写入到栈$rsp+0x40处
   0x000000000049195e <+222>:	lea    0x38(%rsp),%rax  //把存储全局变量地址的栈地址保存到rax中
   0x0000000000491963 <+227>:	mov    %rax,(%rsp)      //地址的地址入栈，tmp变量的栈地址？
   0x0000000000491967 <+231>:	callq  0x4917d0 <main.testInterface>
   0x000000000049196c <+236>:	mov    0x90(%rsp),%rbp
   0x0000000000491974 <+244>:	add    $0x98,%rsp
   0x000000000049197b <+251>:	retq
   0x000000000049197c <+252>:	callq  0x459d70 <runtime.morestack_noctxt>
   0x0000000000491981 <+257>:	jmpq   0x491880 <main.main>
End of assembler dump.
(gdb) print a
$1 = {name = 0x4bfe2b "A name"}
(gdb) print &a
$2 = (main.A *) 0xc0000a2f30
(gdb) print /x $rsp+0x50
$3 = 0xc0000a2f38
(gdb) x /1xg $3
0xc0000a2f38:	0x0000000000000006
(gdb) stepi
0x0000000000491939	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x000000000049193e	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x0000000000491943	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x0000000000491948	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x000000000049194f	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x0000000000491954	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x0000000000491959	28	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
29	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
0x0000000000491963	29	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb)
29	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb) print $rsp
$4 = (void *) 0xc0000a2ee8
(gdb) x /2xg $4 //第一个参数为栈空间地址，该地址处保存了全局变量地址
0xc0000a2ee8:	0x000000c0000a2f20	0x0000000000000001
(gdb) stepi
main.testInterface (v=0xc0000a2f20) at /Users/bianyuan/Desktop/workspace/test/Go_test/test.go:19
19	in /Users/bianyuan/Desktop/workspace/test/Go_test/test.go
(gdb) print /x $rsp+0x38
$5 = 0xc0000a2f18
(gdb) x /1xg v  //可以看到全局变量的地址，参看指令0x0000000000491948 <+200>:	lea    0x199f1(%rip),%rax        # 0x4ab340
0xc0000a2f20:	0x00000000004ab340
(gdb)
*/
