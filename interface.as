/*
476 func assertE2I(inter *interfacetype, e eface) (r iface) {
477     t := e._type
478     if t == nil {
479         // explicit conversions require non-nil interface value.
480         panic(&TypeAssertionError{nil, nil, &inter.typ, ""})
481     }
482     r.tab = getitab(inter, t, false)
483     r.data = e.data
484     return
485 }

 205 type eface struct { 
 206     _type *_type
 207     data  unsafe.Pointer
 208 } 

 200 type iface struct { 
 201     tab  *itab
 202     data unsafe.Pointer
 203 } 
*/

=> 0x0000000000e0b038 <+520>:	lea    0x6247c1(%rip),%rax        # 0x142f800
   0x0000000000e0b03f <+527>:	mov    %rax,(%rsp)   /*第一个参数inter入栈*/
   0x0000000000e0b043 <+531>:	mov    0x198(%rsp),%rax /*第二个参数e入栈，因为是值拷贝，所以此处相当于e._type入栈*/
   0x0000000000e0b04b <+539>:	mov    %rax,0x8(%rsp)
   0x0000000000e0b050 <+544>:	mov    0x1a0(%rsp),%rcx /*e.data入栈*/
   0x0000000000e0b058 <+552>:	mov    %rcx,0x10(%rsp)
   0x0000000000e0b05d <+557>:	callq  0x40ce00 <runtime.assertE2I2>
   0x0000000000e0b062 <+562>:	mov    0x20(%rsp),%rax  /*返回值r.data*/
   0x0000000000e0b067 <+567>:	mov    0x18(%rsp),%rcx  /*返回值r.tab*/
   0x0000000000e0b06c <+572>:	cmpb   $0x0,0x28(%rsp) /*返回的ok*/
   0x0000000000e0b071 <+577>:	je     0xe0b0e3 <gitlab.alibaba-inc.com/ngwf/warden/plugins.register+691>
   0x0000000000e0b073 <+579>:	mov    0x50(%rsp),%rdx
   0x0000000000e0b078 <+584>:	mov    %rcx,0x48(%rdx)
   0x0000000000e0b07c <+588>:	cmpl   $0x0,0x167ac8d(%rip)        # 0x2485d10 <runtime.writeBarrier>
   0x0000000000e0b083 <+595>:	jne    0xe0b462 <gitlab.alibaba-inc.com/ngwf/warden/plugins.register+1586>
   0x0000000000e0b089 <+601>:	mov    %rax,0x50(%rdx)
   0x0000000000e0b08d <+605>:	xorps  %xmm0,%xmm0
   0x0000000000e0b090 <+608>:	movups %xmm0,0x88(%rsp)
   0x0000000000e0b098 <+616>:	lea    0x79dd05(%rip),%rax        # 0x15a8da4
   0x0000000000e0b09f <+623>:	mov    %rax,0x88(%rsp)
   0x0000000000e0b0a7 <+631>:	movq   $0x7,0x90(%rsp)


(gdb) bt
#0  runtime.assertE2I2 (inter=0x142f800, e=..., r=..., b=<optimized out>)
    at /root/go/src/runtime/iface.go:487
#1  0x0000000000e0b062 in gitlab.alibaba-inc.com/ngwf/warden/plugins.register (name=..., plugin=...)
    at /home/admin/work/src/gitlab.alibaba-inc.com/ngwf/warden/plugins/register.go:75
#2  0x0000000000e0bd15 in gitlab.alibaba-inc.com/ngwf/warden/plugins.InitPlugins ()
    at /home/admin/work/src/gitlab.alibaba-inc.com/ngwf/warden/plugins/register.go:180
#3  0x0000000000e0dd5b in main.start ()
    at /home/admin/work/src/gitlab.alibaba-inc.com/ngwf/warden/main.go:223
#4  0x0000000000e0e3a9 in main.main ()
    at /home/admin/work/src/gitlab.alibaba-inc.com/ngwf/warden/main.go:288
(gdb) select-frame 1
(gdb) print name
$53 = 0x15a4a7a "waf"
(gdb) select-frame 0
(gdb) x /4xg $rsp  /*返回地址+inter+e._type+e.data*/
0xc000117ab8:	0x0000000000e0b062	0x000000000142f800
0xc000117ac8:	0x000000000158cc80	0x000000c0010da0b0
(gdb) print inter
$54 = (runtime.interfacetype *) 0x142f800
(gdb) print inter[0]
$55 = {typ = {size = 16, ptrdata = 16, hash = 2373248247, tflag = 7 '\a', align = 8 '\b',
    fieldAlign = 8 '\b', kind = 20 '\024', equal = {void (void *, void *, bool *)} 0x142f800,
    gcdata = 0x1720aff "\002\003\004\005\006\a\b\t\n\v\f\r\016\017\020\021\022\023\024\025\026\027\030\031\032\033\034\035\036\037 !\"#$%&'()*+,-.01234569:;<=>?@ABCDEHIJLOPQRTUXYZ[\\]_`abcefghijkmqtxyz{|}~\177\204\205\206\211\212\216\221\222\223\225\226\230\231\232\236\241\244\245\250\251\252\255\261\271\302\305\312\313\315\322\324\325\330\333\334\335\340\342\345\347\361\362\372\376\377", str = 215536,
    ptrToThis = 522624}, pkgpath = {bytes = 0x13707c0 ""}, mhdr = {array = 0x142f860, len = 1,
    cap = 1}}
(gdb) print e /*可以看到栈中数据跟e的值一致*/
$56 = {_type = 0x158cc80, data = 0xc0010da0b0}

