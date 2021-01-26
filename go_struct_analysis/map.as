
/*
92 func mapassign_fast64(t *maptype, h *hmap, key uint64) unsafe.Pointer 

365 type maptype struct {
366     typ    _type 
367     key    *_type 
368     elem   *_type
369     bucket *_type // internal type representing a hash bucket
370     // function for hashing keys (ptr to key, seed) -> hash
371     hasher     func(unsafe.Pointer, uintptr) uintptr
372     keysize    uint8  // size of key slot
373     elemsize   uint8  // size of elem slot
374     bucketsize uint16 // size of bucket
375     flags      uint32
376 }

*/

0000000000491840 <main.main>:
  491840:	64 48 8b 0c 25 f8 ff 	mov    %fs:0xfffffffffffffff8,%rcx
  491847:	ff ff 
  491849:	48 8d 44 24 d8       	lea    -0x28(%rsp),%rax
  49184e:	48 3b 41 10          	cmp    0x10(%rcx),%rax
  491852:	0f 86 35 03 00 00    	jbe    491b8d <main.main+0x34d>
#分配栈空间，$rsp+0xa8为返回值，$rsp+0xb0为第一个参数
  491858:	48 81 ec a8 00 00 00 	sub    $0xa8,%rsp
  49185f:	48 89 ac 24 a0 00 00 	mov    %rbp,0xa0(%rsp)
  491866:	00 
  491867:	48 8d ac 24 a0 00 00 	lea    0xa0(%rsp),%rbp
  49186e:	00 
  49186f:	0f 57 c0             	xorps  %xmm0,%xmm0
  491872:	0f 11 84 24 90 00 00 	movups %xmm0,0x90(%rsp)
  491879:	00 
  49187a:	48 8d 05 1f e3 00 00 	lea    0xe31f(%rip),%rax        # 49fba0 <type.*+0xdba0>
  491881:	48 89 84 24 90 00 00 	mov    %rax,0x90(%rsp)
  491888:	00 
  491889:	48 8d 05 c0 88 04 00 	lea    0x488c0(%rip),%rax        # 4da150 <syscall.Syscall6.args_stackmap+0x168>
  491890:	48 89 84 24 98 00 00 	mov    %rax,0x98(%rsp)
  491897:	00 
  491898:	48 8b 05 51 19 0d 00 	mov    0xd1951(%rip),%rax        # 5631f0 <os.Stdout>
  49189f:	48 8d 0d da 9d 04 00 	lea    0x49dda(%rip),%rcx        # 4db680 <go.itab.*os.File,io.Writer>
  4918a6:	48 89 0c 24          	mov    %rcx,(%rsp)
  4918aa:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  4918af:	48 8d 84 24 90 00 00 	lea    0x90(%rsp),%rax
  4918b6:	00 
  4918b7:	48 89 44 24 10       	mov    %rax,0x10(%rsp)
  4918bc:	48 c7 44 24 18 01 00 	movq   $0x1,0x18(%rsp)
  4918c3:	00 00 
  4918c5:	48 c7 44 24 20 01 00 	movq   $0x1,0x20(%rsp)
  4918cc:	00 00 
  4918ce:	e8 5d 99 ff ff       	callq  48b230 <fmt.Fprintln>
#创建hmap结构体
  4918d3:	e8 a8 ae f7 ff       	callq  40c780 <runtime.makemap_small>
#把hmap保存到$rax寄存器中
  4918d8:	48 8b 04 24          	mov    (%rsp),%rax
  4918dc:	48 89 44 24 58       	mov    %rax,0x58(%rsp)
#从.rodata段加载maptype
  4918e1:	48 8d 0d b8 4e 01 00 	lea    0x14eb8(%rip),%rcx        # 4a67a0 <type.*+0x147a0>
#maptype参数
  4918e8:	48 89 0c 24          	mov    %rcx,(%rsp)
#hmap结构体指针，第二个参数
  4918ec:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
#第三个参数，即key=3 a[3]=4
  4918f1:	48 c7 44 24 10 03 00 	movq   $0x3,0x10(%rsp)
  4918f8:	00 00 
  4918fa:	e8 01 df f7 ff       	callq  40f800 <runtime.mapassign_fast64>
  4918ff:	48 8b 44 24 18       	mov    0x18(%rsp),%rax
  491904:	48 c7 00 04 00 00 00 	movq   $0x4,(%rax)
  49190b:	48 8d 05 8e 4e 01 00 	lea    0x14e8e(%rip),%rax        # 4a67a0 <type.*+0x147a0>
  491912:	48 89 04 24          	mov    %rax,(%rsp)
  491916:	48 8b 4c 24 58       	mov    0x58(%rsp),%rcx
  49191b:	48 89 4c 24 08       	mov    %rcx,0x8(%rsp)
  491920:	48 c7 44 24 10 04 00 	movq   $0x4,0x10(%rsp)
  491927:	00 00 
  491929:	e8 d2 de f7 ff       	callq  40f800 <runtime.mapassign_fast64>
  49192e:	48 8b 44 24 18       	mov    0x18(%rsp),%rax
  491933:	48 c7 00 05 00 00 00 	movq   $0x5,(%rax)
  49193a:	0f 57 c0             	xorps  %xmm0,%xmm0
  49193d:	0f 11 84 24 80 00 00 	movups %xmm0,0x80(%rsp)
  491944:	00 
  491945:	48 8d 05 54 4e 01 00 	lea    0x14e54(%rip),%rax        # 4a67a0 <type.*+0x147a0>
  49194c:	48 89 84 24 80 00 00 	mov    %rax,0x80(%rsp)
  491953:	00 
  491954:	48 8b 4c 24 58       	mov    0x58(%rsp),%rcx
  491959:	48 89 8c 24 88 00 00 	mov    %rcx,0x88(%rsp)
  491960:	00 
  491961:	48 8b 0d 88 18 0d 00 	mov    0xd1888(%rip),%rcx        # 5631f0 <os.Stdout>
  491968:	48 8d 15 11 9d 04 00 	lea    0x49d11(%rip),%rdx        # 4db680 <go.itab.*os.File,io.Writer>
  49196f:	48 89 14 24          	mov    %rdx,(%rsp)
  491973:	48 89 4c 24 08       	mov    %rcx,0x8(%rsp)
  491978:	48 8d 8c 24 80 00 00 	lea    0x80(%rsp),%rcx
  49197f:	00 
  491980:	48 89 4c 24 10       	mov    %rcx,0x10(%rsp)
  491985:	48 c7 44 24 18 01 00 	movq   $0x1,0x18(%rsp)
  49198c:	00 00 
  49198e:	48 c7 44 24 20 01 00 	movq   $0x1,0x20(%rsp)
  491995:	00 00 
  491997:	e8 94 98 ff ff       	callq  48b230 <fmt.Fprintln>
  49199c:	48 8d 05 fd 4d 01 00 	lea    0x14dfd(%rip),%rax        # 4a67a0 <type.*+0x147a0>
  4919a3:	48 89 04 24          	mov    %rax,(%rsp)
  4919a7:	48 c7 44 24 08 20 00 	movq   $0x20,0x8(%rsp)
  4919ae:	00 00 
  4919b0:	48 c7 44 24 10 00 00 	movq   $0x0,0x10(%rsp)
  4919b7:	00 00 
  4919b9:	e8 52 ae f7 ff       	callq  40c810 <runtime.makemap>
  4919be:	48 8b 44 24 18       	mov    0x18(%rsp),%rax
  4919c3:	48 89 44 24 50       	mov    %rax,0x50(%rsp)
  4919c8:	31 c9                	xor    %ecx,%ecx
  4919ca:	eb 34                	jmp    491a00 <main.main+0x1c0>
  4919cc:	48 89 4c 24 40       	mov    %rcx,0x40(%rsp)
  4919d1:	48 8d 15 c8 4d 01 00 	lea    0x14dc8(%rip),%rdx        # 4a67a0 <type.*+0x147a0>
  4919d8:	48 89 14 24          	mov    %rdx,(%rsp)
  4919dc:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  4919e1:	48 89 4c 24 10       	mov    %rcx,0x10(%rsp)
  4919e6:	e8 15 de f7 ff       	callq  40f800 <runtime.mapassign_fast64>
  4919eb:	48 8b 44 24 18       	mov    0x18(%rsp),%rax
  4919f0:	48 8b 4c 24 40       	mov    0x40(%rsp),%rcx
  4919f5:	48 89 08             	mov    %rcx,(%rax)
  4919f8:	48 ff c1             	inc    %rcx
  4919fb:	48 8b 44 24 50       	mov    0x50(%rsp),%rax
  491a00:	48 83 f9 20          	cmp    $0x20,%rcx
  491a04:	7c c6                	jl     4919cc <main.main+0x18c>
  491a06:	0f 57 c0             	xorps  %xmm0,%xmm0
  491a09:	0f 11 44 24 70       	movups %xmm0,0x70(%rsp)
  491a0e:	48 8d 0d 8b 4d 01 00 	lea    0x14d8b(%rip),%rcx        # 4a67a0 <type.*+0x147a0>
  491a15:	48 89 4c 24 70       	mov    %rcx,0x70(%rsp)
  491a1a:	48 89 44 24 78       	mov    %rax,0x78(%rsp)
  491a1f:	48 8b 05 ca 17 0d 00 	mov    0xd17ca(%rip),%rax        # 5631f0 <os.Stdout>
  491a26:	48 8d 0d 53 9c 04 00 	lea    0x49c53(%rip),%rcx        # 4db680 <go.itab.*os.File,io.Writer>
  491a2d:	48 89 0c 24          	mov    %rcx,(%rsp)
  491a31:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  491a36:	48 8d 44 24 70       	lea    0x70(%rsp),%rax
  491a3b:	48 89 44 24 10       	mov    %rax,0x10(%rsp)
  491a40:	48 c7 44 24 18 01 00 	movq   $0x1,0x18(%rsp)
  491a47:	00 00 
  491a49:	48 c7 44 24 20 01 00 	movq   $0x1,0x20(%rsp)
  491a50:	00 00 
  491a52:	e8 d9 97 ff ff       	callq  48b230 <fmt.Fprintln>
  491a57:	48 8d 05 a2 4d 01 00 	lea    0x14da2(%rip),%rax        # 4a6800 <type.*+0x14800>
  491a5e:	48 89 04 24          	mov    %rax,(%rsp)
  491a62:	48 c7 44 24 08 20 00 	movq   $0x20,0x8(%rsp)
  491a69:	00 00 
  491a6b:	48 c7 44 24 10 00 00 	movq   $0x0,0x10(%rsp)
  491a72:	00 00 
  491a74:	e8 97 ad f7 ff       	callq  40c810 <runtime.makemap>
  491a79:	48 8b 44 24 18       	mov    0x18(%rsp),%rax
  491a7e:	48 89 44 24 48       	mov    %rax,0x48(%rsp)
  491a83:	48 8d 0d 76 4d 01 00 	lea    0x14d76(%rip),%rcx        # 4a6800 <type.*+0x14800>
  491a8a:	48 89 0c 24          	mov    %rcx,(%rsp)
  491a8e:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  491a93:	48 c7 44 24 10 00 00 	movq   $0x0,0x10(%rsp)
  491a9a:	00 00 
  491a9c:	e8 5f dd f7 ff       	callq  40f800 <runtime.mapassign_fast64>
  491aa1:	48 8b 7c 24 18       	mov    0x18(%rsp),%rdi
  491aa6:	48 c7 47 08 04 00 00 	movq   $0x4,0x8(%rdi)
  491aad:	00 
  491aae:	83 3d 6b b2 0f 00 00 	cmpl   $0x0,0xfb26b(%rip)        # 58cd20 <runtime.writeBarrier>
  491ab5:	0f 85 c1 00 00 00    	jne    491b7c <main.main+0x33c>
  491abb:	48 8d 05 c3 e2 02 00 	lea    0x2e2c3(%rip),%rax        # 4bfd85 <go.string.*+0x12d>
  491ac2:	48 89 07             	mov    %rax,(%rdi)
  491ac5:	48 8d 05 34 4d 01 00 	lea    0x14d34(%rip),%rax        # 4a6800 <type.*+0x14800>
  491acc:	48 89 04 24          	mov    %rax,(%rsp)
  491ad0:	48 8b 4c 24 48       	mov    0x48(%rsp),%rcx
  491ad5:	48 89 4c 24 08       	mov    %rcx,0x8(%rsp)
  491ada:	48 c7 44 24 10 01 00 	movq   $0x1,0x10(%rsp)
  491ae1:	00 00 
  491ae3:	e8 18 dd f7 ff       	callq  40f800 <runtime.mapassign_fast64>
  491ae8:	48 8b 7c 24 18       	mov    0x18(%rsp),%rdi
  491aed:	48 c7 47 08 04 00 00 	movq   $0x4,0x8(%rdi)
  491af4:	00 
  491af5:	83 3d 24 b2 0f 00 00 	cmpl   $0x0,0xfb224(%rip)        # 58cd20 <runtime.writeBarrier>
  491afc:	75 70                	jne    491b6e <main.main+0x32e>
  491afe:	48 8d 05 84 e2 02 00 	lea    0x2e284(%rip),%rax        # 4bfd89 <go.string.*+0x131>
  491b05:	48 89 07             	mov    %rax,(%rdi)
  491b08:	0f 57 c0             	xorps  %xmm0,%xmm0
  491b0b:	0f 11 44 24 60       	movups %xmm0,0x60(%rsp)
  491b10:	48 8d 05 e9 4c 01 00 	lea    0x14ce9(%rip),%rax        # 4a6800 <type.*+0x14800>
  491b17:	48 89 44 24 60       	mov    %rax,0x60(%rsp)
  491b1c:	48 8b 44 24 48       	mov    0x48(%rsp),%rax
  491b21:	48 89 44 24 68       	mov    %rax,0x68(%rsp)
  491b26:	48 8b 05 c3 16 0d 00 	mov    0xd16c3(%rip),%rax        # 5631f0 <os.Stdout>
  491b2d:	48 8d 0d 4c 9b 04 00 	lea    0x49b4c(%rip),%rcx        # 4db680 <go.itab.*os.File,io.Writer>
  491b34:	48 89 0c 24          	mov    %rcx,(%rsp)
  491b38:	48 89 44 24 08       	mov    %rax,0x8(%rsp)
  491b3d:	48 8d 44 24 60       	lea    0x60(%rsp),%rax
  491b42:	48 89 44 24 10       	mov    %rax,0x10(%rsp)
  491b47:	48 c7 44 24 18 01 00 	movq   $0x1,0x18(%rsp)
  491b4e:	00 00 
  491b50:	48 c7 44 24 20 01 00 	movq   $0x1,0x20(%rsp)
  491b57:	00 00 
  491b59:	e8 d2 96 ff ff       	callq  48b230 <fmt.Fprintln>
  491b5e:	48 8b ac 24 a0 00 00 	mov    0xa0(%rsp),%rbp
  491b65:	00 
  491b66:	48 81 c4 a8 00 00 00 	add    $0xa8,%rsp
  491b6d:	c3                   	retq   
  491b6e:	48 8d 05 14 e2 02 00 	lea    0x2e214(%rip),%rax        # 4bfd89 <go.string.*+0x131>
  491b75:	e8 d6 a1 fc ff       	callq  45bd50 <runtime.gcWriteBarrier>
  491b7a:	eb 8c                	jmp    491b08 <main.main+0x2c8>
  491b7c:	48 8d 05 02 e2 02 00 	lea    0x2e202(%rip),%rax        # 4bfd85 <go.string.*+0x12d>
  491b83:	e8 c8 a1 fc ff       	callq  45bd50 <runtime.gcWriteBarrier>
  491b88:	e9 38 ff ff ff       	jmpq   491ac5 <main.main+0x285>
  491b8d:	e8 7e 83 fc ff       	callq  459f10 <runtime.morestack_noctxt>
  491b92:	e9 a9 fc ff ff       	jmpq   491840 <main.main>

(gdb) info locals
Python Exception <class 'gdb.error'> Attempt to take contents of a non-pointer value.:
a = map[int]int
b = map[int]int = {[1] = 1, [4] = 4, [10] = 10, [18] = 18, [23] = 23, [0] = 0, [8] = 8, [21] = 21, [28] = 28, [29] = 29, [5] = 5,
  [9] = 9, [17] = 17, [19] = 19, [24] = 24, [30] = 30, [7] = 7, [13] = 13, [16] = 16, [31] = 31, [6] = 6, [11] = 11, [12] = 12,
  [20] = 20, [14] = 14, [15] = 15, [22] = 22, [26] = 26, [27] = 27, [2] = 2, [3] = 3, [25] = 25}
c = <optimized out>
(gdb) ptype b
type = struct hash<int, int> {
    int count;
    uint8 flags;
    uint8 B;
    uint16 noverflow;
    uint32 hash0;
    struct bucket<int, int> *buckets;
    struct bucket<int, int> *oldbuckets;
    uintptr nevacuate;
    runtime.mapextra *extra;
} *
(gdb) print b.extra
$1 = (runtime.mapextra *) 0x0
(gdb) print b.buckets
$2 = (struct bucket<int, int> *) 0xc000078000
(gdb) print b.buckets[0] #查看key/value，key/value是分开存储的
$3 = {tophash = "\356C\335\332\342\000\000", keys = {1, 4, 10, 18, 23, 0, 0, 0}, values = {1, 4, 10, 18, 23, 0, 0, 0}, overflow = 0x0}
(gdb) print b.buckets[1]
$4 = {tophash = "\264\275\245\017R\000\000", keys = {0, 8, 21, 28, 29, 0, 0, 0}, values = {0, 8, 21, 28, 29, 0, 0, 0}, overflow = 0x0}
(gdb) print b.buckets[2]
$5 = {tophash = "\000\000\000\000\000\000\000", keys = {0, 0, 0, 0, 0, 0, 0, 0}, values = {0, 0, 0, 0, 0, 0, 0, 0}, overflow = 0x0}
(gdb) print b.buckets[3]
$6 = {tophash = "\376A[\332Ť\000", keys = {5, 9, 17, 19, 24, 30, 0, 0}, values = {5, 9, 17, 19, 24, 30, 0, 0}, overflow = 0x0}
(gdb) print /x b.buckets[0]
$7 = {tophash = {0xee, 0x43, 0xdd, 0xda, 0xe2, 0x0, 0x0, 0x0}, keys = {0x1, 0x4, 0xa, 0x12, 0x17, 0x0, 0x0, 0x0}, values = {0x1, 0x4,
    0xa, 0x12, 0x17, 0x0, 0x0, 0x0}, overflow = 0x0}

(gdb) list
12		fmt.Println(a)
13		b := make(map[int]int, 32)
14		for i := 0; i < 32; i++ {
15			b[i]=i
16		}
17		fmt.Println(b)
18		c := make(map[int]string, 32)
19		c[0]="2021"
20		c[1]="2022"
21		fmt.Println(c)
(gdb) break 21
Breakpoint 2 at 0x491b08: file /home/baohu.rbh/hs_engine/src/map.go, line 21.
(gdb) c
Continuing.
map[0:0 1:1 2:2 3:3 4:4 5:5 6:6 7:7 8:8 9:9 10:10 11:11 12:12 13:13 14:14 15:15 16:16 17:17 18:18 19:19 20:20 21:21 22:22 23:23 24:24 25:25 26:26 27:27 28:28 29:29 30:30 31:31]

Breakpoint 2, main.main () at /home/baohu.rbh/hs_engine/src/map.go:21
21		fmt.Println(c)
(gdb) info locals
a = map[int]int
b = map[int]int
c = map[int]string = {[1] = "2022", [0] = "2021"}
(gdb) print c[0]
$8 = {count = 2, flags = 0 '\000', B = 3 '\003', noverflow = 0, hash0 = 3299835963, buckets = 0xc00006b500, oldbuckets = 0x0,
  nevacuate = 0, extra = 0x0}
(gdb) print c.buckets[0]
$9 = {tophash = "\r\000\000\000\000\000\000", keys = {1, 0, 0, 0, 0, 0, 0, 0}, values = {"2022", "", "", "", "", "", "", ""},
  overflow = 0x0}
(gdb) print c.buckets[1]
$10 = {tophash = "\000\000\000\000\000\000\000", keys = {0, 0, 0, 0, 0, 0, 0, 0}, values = {"", "", "", "", "", "", "", ""},
  overflow = 0x0}
(gdb) print c.buckets[2]
$11 = {tophash = "\000\000\000\000\000\000\000", keys = {0, 0, 0, 0, 0, 0, 0, 0}, values = {"", "", "", "", "", "", "", ""},
  overflow = 0x0}
(gdb) print c.buckets[7]
$12 = {tophash = "\000\000\000\000\000\000\000", keys = {0, 0, 0, 0, 0, 0, 0, 0}, values = {"", "", "", "", "", "", "", ""},
  overflow = 0x0}
(gdb) print b[0]
$13 = {count = 32, flags = 3 '\003', B = 3 '\003', noverflow = 0, hash0 = 3397920808, buckets = 0xc000078000, oldbuckets = 0x0,
  nevacuate = 0, extra = 0xc00000c100}
(gdb) print b[3]
$14 = {count = 0, flags = 0 '\000', B = 0 '\000', noverflow = 0, hash0 = 0, buckets = 0x0, oldbuckets = 0x0, nevacuate = 0, extra = 0x0}
(gdb) print c[0]
$15 = {count = 2, flags = 0 '\000', B = 3 '\003', noverflow = 0, hash0 = 3299835963, buckets = 0xc00006b500, oldbuckets = 0x0,
  nevacuate = 0, extra = 0x0}
(gdb) ptype c
type = struct hash<int, string> {
    int count;
    uint8 flags;
    uint8 B;
    uint16 noverflow;
    uint32 hash0;
    struct bucket<int, string> *buckets;
    struct bucket<int, string> *oldbuckets;
    uintptr nevacuate;
    runtime.mapextra *extra;
} *
(gdb) ptype c.buckets   /*go源代码中该结构体的定义只有[8]uint8 tophash，但是实际编译时则为如下结构体类型，需要编译时计算*/
type = struct bucket<int, string> {
    [8]uint8 tophash;
    []key<int> keys;
    []val<string> values;
    struct bucket<int, string> *overflow;
} *

(gdb) stepi
runtime.makemap_small (~r0=<optimized out>) at /usr/local/go/src/runtime/map.go:292
292	func makemap_small() *hmap {
(gdb) bt
#0  runtime.makemap_small (~r0=<optimized out>) at /usr/local/go/src/runtime/map.go:292
#1  0x00000000004918d8 in main.main () at /home/baohu.rbh/hs_engine/src/map.go:9
(gdb) finish
Run till exit from #0  runtime.makemap_small (~r0=<optimized out>) at /usr/local/go/src/runtime/map.go:292
0x00000000004918d8 in main.main () at /home/baohu.rbh/hs_engine/src/map.go:9
9		a := make(map[int]int)
(gdb) print $rsp
$17 = (void *) 0xc00010eed8
(gdb) x /1xg $17     /*makemap_small创建的hmap结构体，即map*/
0xc00010eed8:	0x000000c000104150
(gdb) print a
$18 = <optimized out>
(gdb) stepi
0x00000000004918dc	9		a := make(map[int]int)
(gdb)
10		a[3]=4
(gdb) print a
$19 = map[int]int
(gdb) print &a    /*map存放在寄存器中，不在堆栈中*/
Address requested for identifier "a" which is in register $rax
(gdb) print /x $rax   /*可以看到map变量a实际就是makemap_small创建的hmap结构体*/
$21 = 0xc000104150
(gdb) stepi
0x00000000004918e8	10		a[3]=4
(gdb)
0x00000000004918ec	10		a[3]=4
(gdb)
0x00000000004918f1	10		a[3]=4
(gdb)
10		a[3]=4
(gdb)        /*t=0x4a67a0，通过汇编可以看到该值是从.rodata中获取的全局变量，所有map都会使用该变量，类型为maptype*/
runtime.mapassign_fast64 (t=0x4a67a0, h=0xc000104150, key=<optimized out>, ~r3=<optimized out>)
    at /usr/local/go/src/runtime/map_fast64.go:92
92	func mapassign_fast64(t *maptype, h *hmap, key uint64) unsafe.Pointer {
(gdb) ptype t
type = struct runtime.maptype {
    runtime._type typ;
    runtime._type *key;
    runtime._type *elem;
    runtime._type *bucket;
    func(unsafe.Pointer, uintptr) uintptr hasher;
    uint8 keysize;
    uint8 elemsize;
    uint16 bucketsize;
    uint32 flags;
} *
(gdb) print t[0]
$22 = {typ = {size = 8, ptrdata = 8, hash = 592976720, tflag = 2 '\002', align = 8 '\b', fieldAlign = 8 '\b', kind = 53 '5',
    equal = {void (void *, void *, bool *)} 0x4a67a0,
    gcdata = 0x4d7d2c "\001\002\003\004\005\006\a\b\t\n\v\f\r\016\017\020\022\024\025\026\030\031\033\034\036\037\"%&(,-256<BUXx\216\231\330\335\345\377", str = 14537, ptrToThis = 0}, key = 0x49f460, elem = 0x49f460, bucket = 0x4b0c60, hasher = {void (void *, uintptr,
    uintptr *)} 0x4a67a0, keysize = 8 '\b', elemsize = 8 '\b', bucketsize = 144, flags = 4}
(gdb) finish
Run till exit from #0  runtime.mapassign_fast64 (t=0x4a67a0, h=0xc000104150, key=<optimized out>, ~r3=<optimized out>)
    at /usr/local/go/src/runtime/map_fast64.go:92
0x00000000004918ff in main.main () at /home/baohu.rbh/hs_engine/src/map.go:10
10		a[3]=4
(gdb)


