#links

Link to see gas syntax:

http://en.wikibooks.org/wiki/X86_Assembly/GAS_Syntax

# std instructions

Assignments typically result in `mov`

```asm
# move immediate value of 1 to eax
mov $1, %eax

# move value in eax to ebx.
mov %eax, %ebx

# value in location pointed by eax to ebx.
mov (%eax), %ebx

# move value in location pointed by (eax+4) to ebx.
mov 4(%eax), %ebx

# move value in location pointed by (eax+ecx+4) to ebx.
mov 4(%eax,%ecx),%ebx

# ((%eax * 8) + %ecx). Take the value at this address and move to edx
mov    (%ecx,%eax,8),%edx

# 4 + ((%eax * 8) + %ecx). Take the value at this address and move to edx
mov   4(%ecx,%eax,8),%edx

# load one byte from %eax + %ecx + 1 and zero-extend to full register
movzbl  0x01(%eax,%ecx),%eax


```

Getting just pointers result in `lea`

```asm
# multiply eax by 8 and add to ecx. Put this address in edx
lea (%ecx,%eax,8), %edx


# lea/shl combine is sometimes abused to do constant arithmetic.
# Here is one snippet that multiplies rax by 384.
lea    (%rax,%rax,2),%rdx      ; 3*rax to rdx
mov    %rdx,%rax               ; rdx back to rax
shl    $0x7,%rax               ; 128*3*orig_rax

# Just another nop (for 4 byte nop)!
lea    0x0(%esi,%eiz,1),%esi

```

* comparision

```asm


# if ebx < eax, when performing unsigned comparision, then set carry
# if both are equal, then set zero
# if XOR of (overflow/sign) is 1, then ebx is < eax
cmp %eax, %ebx

#check if eax is 0, and if 0 sets ZF. Most likely followed with je (was 0) or jne (wasn't 0)
test %eax,%eax
```

```asm

# this is same as 'mov $0,%edi'. Usually preferred over mov for taking lesser space.
xor %edi, %edi

# shift left %esi by 0x18 bits always filling 0 in lsb. The shift bits are
# lost (except for last, which goes into carry)
shl $0x18, %esi

# will shift the bits in %edi by $0x18 bits filling the new lsb with msb bits of %edx
shld   $0x18,%edx,%edi


# Right/Left shift the 2nd register by 1st-register's value time. (Compilers are usually
#                intelligent enuf to use only one-byte of the register)
# sh* doesn't extend sign bit. sar does. sal/shl are exactly the same.
# shl/sal puts the last lost bit in carry flag
shr %cl, %eax
shl %cl, %eax
sar 0x1f, %edx
sal %cl, %edx


# conditional move. Its same as test ... and then je , then %eax,%ebx (the stuff).
# The cmov avoid the brand prediction penalty for a simple one step instruction.
cmov %eax, %ebx

```

* Jumps

* Jumps reference - http://www.unixwiz.net/techtips/x86-jumps.html
* `jns`  jump if no sign.
* Suffixes

    ```
    * b = byte (8 bit)
    * s = short (16 bit integer) or single (32-bit floating point)
    * w = word (16 bit)
    * l = long (32 bit integer or 64-bit floating point)
    * q = quad (64 bit)
    * t = ten bytes (80-bit floating point)
    ```


# Typically in a function

At start, you should see

```asm
push %ebp
mov  %esp, %ebp
push <some-more-registers that are garbled here>
sub  <some-good-value>,%esp   ; <-- this value is the scratch that this frame uses.
```

Now onwards:

* `positive-offset(%ebp)` are args to this function.
* `negative-offset(%ebp)` are local-scratch area to this function, could be local vars or just scratch area.

at exit:

```asm
pop <registers that were pushed in reverse order>
pop %ebp
ret
```
--or--

```
pop <registers that were pushed in reverse order>
leave
```

# Order of arguments

```
function1 (int arg1, int arg2, int arg3)
```

will be

```
push arg3
push arg2
push arg1
call function.
```

# Traps

`__builtin_trap()` ends up as `ud2`

This instruction is provided for software testing to explicitly generate an invalid opcode exception.

Eg: `Main process exited, code=dumped, status=4/ILL`
