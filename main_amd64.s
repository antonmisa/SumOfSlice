// It's all just for amd64 machine.

#include "textflag.h"

// Let's summarize elements of int64 array
// in a classical way in a circle element by element
// func _sumASM(nums []int64) int64
TEXT ·_sumASM(SB), NOSPLIT, $0
    MOVQ nums_ptr+0(FP), DI      // take a pointer to the first element
    MOVQ nums_len+8(FP), AX      // take a pointer to a len of a slice
                                 // why is it so? Let's remember what a slice is?
                                 // struct { //as described in src/runtime/slice.go
                                 //     array unsafe.Pointer      // pointer starts at 0 bit and has 8 bit long
                                 //     len int                   // starts as 8 bit and has 8 bit long on x64 machine
                                 //     cap int
                                 // }                                                                                               

    XORQ R8, R8                  // bit wise xor - empting register accumulator (here will be sum)
loop:
    CMPQ AX, $0                  // comparing AX (where len is saving) with 0 in a cycle, if eq then set a flag
    JE done                      // jump to done if flag in previos step is set
    MOVQ (DI), R9                // get each 64 bit number to R9 register
    ADDQ R9, R8                  // summarize R9 with R8 accumulator
    ADDQ $8, DI                  // move to next number, by adding 8 bit, because we have 64 bit int
    DECQ AX                      // decrease counter
    JMP loop                     // and again 

done:
    MOVQ R8, ret+24(FP)          // set ret value by setting with accumulator. +24 bit because incoming slice has 24 bit size (8(pointer)+8(len)+8(cap))
    RET


// func _sumAVX(x []int64) int64
// Requires: AVX, FMA3, SSE
TEXT ·_sumAVX(SB), NOSPLIT, $0-28
        MOVQ   x_base+0(FP), AX
        MOVQ   x_len+8(FP), CX
        VXORPS Y0, Y0, Y0

blockloop:
        CMPQ        CX, $0x00000004
        JL          tail
        VADDPD      (AX), Y0, Y0
        ADDQ        $0x00000020, AX
        SUBQ        $0x00000004, CX
        JMP         blockloop

tail:
        VXORPS X1, X1, X1

tailloop:
        CMPQ        CX, $0x00000000
        JE          reduce
        VADDSD      (AX), X1, X1
        ADDQ        $0x00000008, AX
        DECQ        CX
        JMP         tailloop

reduce:
        VEXTRACTF128 $0x01, Y0, X2
        VADDPD       X0, X2, X0
        VHADDPD      X0, X0, X0
        MOVSD        X0, ret+24(FP)
        RET
