package interpreter

import "math/bits"



/*
  Opcodes
*/
const (
  ADD_OP  = iota //commutative
  SUB_OP  = iota //eq=0
  MUL_OP  = iota //commutative
  DIV_OP  = iota //eq=1/err, divmod
  MOD_OP  = iota //takes 1 parameter, must be output of div, zero latency
  SHL_OP  = iota
  SHR_OP  = iota
  RTL_OP  = iota
  RTR_OP  = iota
  XOR_OP  = iota //commutative, eq=0
  AND_OP  = iota //commutative, eq=x
  OR_OP   = iota //commutative, eq=x
  NOT_OP  = iota
  SET_OP  = iota
  CMP_OP  = iota
  PCT_OP  = iota
  CTZ_OP  = iota
  CLZ_OP  = iota
  CMOV_OP = iota
  LEA0_OP = iota //a + (b * [1, 2, 4, 8]) + k
  LEA1_OP = iota //a + (b * [1, 2, 4, 8]) + c
  CNST_OP = iota
  PAR_OP  = iota
  RET_OP  = iota
)

/*
  Condition Codes
*/
const (
  CC_LS  = iota //eq=f
  CC_GT  = iota //eq=f
  CC_EQ  = iota //commutative, eq=t
  CC_NE  = iota //commutative, eq=f
  CC_LSE = iota //eq=t
  CC_GTE = iota //eq=t
  CC_E0  = iota
  CC_E1  = iota
  //CC_OVF = iota //Overflow? I have no idea how to handle this one
)

type Opcode  int8

type Regcode int8

type Cndcode int8

type Instruction struct {
  Op  Opcode
  Cnd Cndcode
  R0  Regcode
  R1  Regcode
  R2  Regcode
  Im0 int64
  Im1 int64
}


type InterpreterState struct {
  InCode   [32]Instruction
  TestCode [32]Instruction
  State    [32]int64
  Inputs   [8] int64
  Outputs  [8] int64
  InSize   int
  TestSize int
  TestData [2][8][64]int64  // 32 input-output sets, up to 8 inputs, 8 outputs each
  TestFail [64]bool         // Some tests are expected to fail, e.g, divide-by-zero errors
}


func (i *InterpreterState) Step(s int) bool {

  instruction := i.InCode[s]
  op := instruction.Op
  r0 := instruction.R0
  r1 := instruction.R1
  //r2 := instruction.R2
  //cc := instruction.Cnd
  //im := instruction.Imm

  v0 := i.State[r0]
  v1 := i.State[r1]

  switch op {
  case ADD_OP :
    i.State[s] = v0 + v1
  case SUB_OP :
    i.State[s] = v0 - v1
  case MUL_OP :
    i.State[s] = v0 * v1
  case DIV_OP :
    if i.State[r1] != 0 {
      i.State[s] = v0 / v1
    }else{
      return false
    }
  case MOD_OP :
    dr0 := i.State[i.InCode[r0].R0]
    dr1 := i.State[i.InCode[r0].R1]
    if dr1 != 0 {
      i.State[s] = dr0 % dr1
    }else{
      return false
    }
  case SHL_OP :
    i.State[s] = int64(uint64(v0) << uint64(v1))
  case SHR_OP :
    i.State[s] = int64(uint64(v0) >> uint64(v1))
  case RTL_OP :
    i.State[s] = int64(bits.RotateLeft64(uint64(v0), int( v1) ))
  case RTR_OP :
    i.State[s] = int64(bits.RotateLeft64(uint64(v0), int(-v1) ))
  case XOR_OP :
    i.State[s] = v0 ^ v1
  case AND_OP :
    i.State[s] = v0 & v1
  case OR_OP  :
    i.State[s] = v0 | v1
  case NOT_OP :
    i.State[s] = ^v0
  case SET_OP :
    // This one is more complex
  case PCT_OP :
    i.State[s] = int64(bits.OnesCount64(uint64(v0)))
  case CTZ_OP :
    i.State[s] = int64(bits.TrailingZeros64(uint64(v0)))
  case CLZ_OP :
    i.State[s] = int64(bits.LeadingZeros64(uint64(v0)))
  case CMOV_OP :
    v2 := i.State[instruction.R2]
    if v0 == 0 {
      i.State[s] = v1
    }else{
      i.State[s] = v2
    }
  case LEA0_OP :
    i.State[s] = v0 + (instruction.Im0 * v1) + instruction.Im1
  case LEA1_OP :
    v2 := i.State[instruction.R2]
    i.State[s] = v0 + (instruction.Im0 * v1) + v2
  case CNST_OP :
    i.State[s] = instruction.Im0
  case PAR_OP :
    i.State[s] = i.Inputs[r0]
  case RET_OP :
    i.Outputs[r0] = v0
  }

  return true
}
