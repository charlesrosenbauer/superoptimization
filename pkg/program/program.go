package program


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
  //CMP_OP  = iota
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


type Program struct{
  Code [32]Instruction
  Size int
}
