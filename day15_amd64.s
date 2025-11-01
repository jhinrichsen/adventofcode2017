// func day15Part1AsmB1(startA, startB uint) uint
TEXT ·day15Part1AsmB1(SB), $0-24
    // Load arguments
    MOVQ startA+0(FP), R14   // R14 = a
    MOVQ startB+8(FP), R15   // R15 = b
    XORQ R8, R8              // R8 = matches = 0

    // Constants
    MOVQ $16807, R9          // R9 = factorA
    MOVQ $48271, R10         // R10 = factorB
    MOVQ $2147483647, R11    // R11 = modulo
    MOVQ $40000000, R12      // R12 = pairs counter
    MOVQ $0xFFFF, R13        // R13 = mask

loop:
    // a = (a * factorA) % modulo
    MOVQ R14, AX
    MULQ R9                  // DX:AX = AX * R9
    DIVQ R11                 // DX = DX:AX % R11
    MOVQ DX, R14             // a = remainder

    // b = (b * factorB) % modulo
    MOVQ R15, AX
    MULQ R10                 // DX:AX = AX * R10
    DIVQ R11                 // DX = DX:AX % R11
    MOVQ DX, R15             // b = remainder

    // Compare lowest 16 bits
    MOVQ R14, AX
    ANDQ R13, AX             // AX = a & mask
    MOVQ R15, BX
    ANDQ R13, BX             // BX = b & mask
    CMPQ AX, BX
    JNE skip
    INCQ R8                  // matches++

skip:
    DECQ R12
    JNZ loop

    MOVQ R8, ret+16(FP)
    RET

// func day15Part2AsmB1(startA, startB uint) uint
TEXT ·day15Part2AsmB1(SB), $0-24
    // Load arguments
    MOVQ startA+0(FP), R14   // R14 = a
    MOVQ startB+8(FP), R15   // R15 = b
    XORQ R8, R8              // R8 = matches = 0

    // Constants
    MOVQ $16807, R9          // R9 = factorA
    MOVQ $48271, R10         // R10 = factorB
    MOVQ $2147483647, R11    // R11 = modulo
    MOVQ $5000000, R12       // R12 = pairs counter
    MOVQ $0xFFFF, R13        // R13 = mask

pairLoop:
    // Generate next a that is multiple of 4
genA:
    MOVQ R14, AX
    MULQ R9                  // DX:AX = AX * R9
    DIVQ R11                 // DX = DX:AX % R11
    MOVQ DX, R14             // a = remainder

    TESTQ $3, R14            // Check if a % 4 == 0
    JNZ genA                 // If not, generate again

    // Generate next b that is multiple of 8
genB:
    MOVQ R15, AX
    MULQ R10                 // DX:AX = AX * R10
    DIVQ R11                 // DX = DX:AX % R11
    MOVQ DX, R15             // b = remainder

    TESTQ $7, R15            // Check if b % 8 == 0
    JNZ genB                 // If not, generate again

    // Compare lowest 16 bits
    MOVQ R14, AX
    ANDQ R13, AX             // AX = a & mask
    MOVQ R15, BX
    ANDQ R13, BX             // BX = b & mask
    CMPQ AX, BX
    JNE skipPair
    INCQ R8                  // matches++

skipPair:
    DECQ R12
    JNZ pairLoop

    MOVQ R8, ret+16(FP)
    RET
