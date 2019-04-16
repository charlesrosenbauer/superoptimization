package interpreter

import (
  "math/bits"

  "github.com/charlesrosenbauer/superoptimization/pkg/program"
)






type InterpreterState struct {
  InitCode program.Program
  OptmCode program.Program
  State    [32]int64
  Inputs   [8] int64
  Outputs  [8] int64
  TestData [2][8][64]int64  // 32 input-output sets, up to 8 inputs, 8 outputs each
  TestFail [64]bool         // Some tests are expected to fail, e.g, divide-by-zero errors
}

func (i *InterpreterState) StepOptm(s int) bool {

  return i.Step(&i.OptmCode, s)
}


func (i *InterpreterState) StepInit(s int) bool {

  return i.Step(&i.InitCode, s)
}


func (i *InterpreterState) Step(p *program.Program, s int) bool {

  instruction := p.Code[s]
  op := instruction.Op
  r0 := instruction.R0
  r1 := instruction.R1
  //r2 := instruction.R2
  //cc := instruction.Cnd
  //im := instruction.Imm

  v0 := i.State[r0]
  v1 := i.State[r1]

  switch op {
  case program.ADD_OP :
    i.State[s] = v0 + v1
  case program.SUB_OP :
    i.State[s] = v0 - v1
  case program.MUL_OP :
    i.State[s] = v0 * v1
  case program.DIV_OP :
    if i.State[r1] != 0 {
      i.State[s] = v0 / v1
    }else{
      return false
    }
  case program.MOD_OP :
    dr0 := i.State[p.Code[r0].R0]
    dr1 := i.State[p.Code[r0].R1]
    if dr1 != 0 {
      i.State[s] = dr0 % dr1
    }else{
      return false
    }
  case program.SHL_OP :
    i.State[s] = int64(uint64(v0) << uint64(v1))
  case program.SHR_OP :
    i.State[s] = int64(uint64(v0) >> uint64(v1))
  case program.RTL_OP :
    i.State[s] = int64(bits.RotateLeft64(uint64(v0), int( v1) ))
  case program.RTR_OP :
    i.State[s] = int64(bits.RotateLeft64(uint64(v0), int(-v1) ))
  case program.XOR_OP :
    i.State[s] = v0 ^ v1
  case program.AND_OP :
    i.State[s] = v0 & v1
  case program.OR_OP  :
    i.State[s] = v0 | v1
  case program.NOT_OP :
    i.State[s] = ^v0
  case program.SET_OP :
    cc := instruction.Cnd
    i.State[s] = 0
    switch{
    case (cc == program.CC_LS ) && (v0 <  v1):
      i.State[s] = 1
    case (cc == program.CC_GT ) && (v0 >  v1):
      i.State[s] = 1
    case (cc == program.CC_LSE) && (v0 <= v1):
      i.State[s] = 1
    case (cc == program.CC_GTE) && (v0 >= v1):
      i.State[s] = 1
    case (cc == program.CC_EQ)  && (v0 == v1):
      i.State[s] = 1
    case (cc == program.CC_NE)  && (v0 == v1):
      i.State[s] = 1
    case (cc == program.CC_E0)  && (v0 == 0):
      i.State[s] = 1
    case (cc == program.CC_E1)  && (v0 == 1):
      i.State[s] = 1
    }
  case program.PCT_OP :
    i.State[s] = int64(bits.OnesCount64(uint64(v0)))
  case program.CTZ_OP :
    i.State[s] = int64(bits.TrailingZeros64(uint64(v0)))
  case program.CLZ_OP :
    i.State[s] = int64(bits.LeadingZeros64(uint64(v0)))
  case program.CMOV_OP :
    v2 := i.State[instruction.R2]
    if v0 == 0 {
      i.State[s] = v1
    }else{
      i.State[s] = v2
    }
  case program.LEA0_OP :
    i.State[s] = v0 + (instruction.Im0 * v1) + instruction.Im1
  case program.LEA1_OP :
    v2 := i.State[instruction.R2]
    i.State[s] = v0 + (instruction.Im0 * v1) + v2
  case program.CNST_OP :
    i.State[s] = instruction.Im0
  case program.PAR_OP :
    i.State[s] = i.Inputs[r0]
  case program.RET_OP :
    i.Outputs[r0] = v0
  }

  return true
}
