#include "textflag.h"

TEXT ·archCeil(SB),NOSPLIT,$0-64
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ AX, CX
    IMULQ BX, CX
    MOVQ CX, ret1+16(FP)
    RET
isBig:
    MOVQ AX,ret+8(FP)
    RET
