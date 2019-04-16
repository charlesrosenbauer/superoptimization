package search


import (


  "github.com/charlesrosenbauer/superoptimization/pkg/program"
)



func (p *program.Program) ApproxCost(){

  cost := 0
  for i := 0; i < p.Size; i++ {
    switch p.Code[i].Op {
    case OP_MUL :
      cost += 3
    case OP_DIV :
      cost += 22
    case OP_MOD :
      cost += 0
    case OP_LEA1:
      cost += 2
    default:
      cost += 1
    }
  }

  return cost
}
