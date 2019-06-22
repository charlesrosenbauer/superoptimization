package analysis


import (
  "github.com/charlesrosenbauer/superoptimization/pkg/program"
)






// This is essentially just an 8x8 bit lattice for now.
type DepMap struct{
  uint64 deps
}


func isZero(Instruction i) bool {

  isCancel := false

  switch i.Opcode {
  case SUB_OP : isCancel = true
  case XOR_OP : isCancel = true
  case SET_OP : {
    switch i.Cndcode {
    case CC_LS : isCancel = true
    case CC_GT : isCancel = true
    case CC_NE : isCancel = true
    }}
  }
  return isCancel
}
