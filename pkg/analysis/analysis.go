package analysis


import (
  "github.com/charlesrosenbauer/superoptimization/pkg/program"
)






// This is essentially just an 8x8 bit lattice for now.
type DepMap struct{
  Deps [8]
}


func isZero(Instruction i) bool {

  /*
    We're going to need speed wherever we can get it in this superoptimizer.

    This is probably not the most performance-critical part, but it likely won't be insignificant,
    and some optimizations here to remove branches aren't too crazy.

    This is just doing some stuff with bitsets to avoid switch statements.
  */
  opPass := 0 != (((1 << SUB_OP) | (1 << XOR_OP) | (1 << SET_OP)) & (1 << uint(i.Op)))

  ccPass := 0 != (((1 << CC_VD) | (1 << CC_LS) | (1 << CC_GT) | (1 << CC_NE)) & (1 << uint(i.Cnd)))

  return opPass && ccPass && (i.R0 == i.R1)
}


func isConst(Instruction i) bool {

  opPass := 0 != (((1 << SUB_OP) | (1 << XOR_OP) | (1 << SET_OP)) & (1 << uint(i.Op)))

  ccPass := 0 != (((1 << CC_VD) | (1 << CC_LS) | (1 << CC_GT) | (1 << CC_NE) |
                   (1 << CC_EQ) | (1 << CC_LSE)| (1 << CC_GTE))              & (1 << uint(i.Cnd)))

  return opPass && ccPass && (i.R0 == i.R1)
}




func getDepMap(Program* p) DepMap {
  var ret DepMap

  var deps [32]uint8

  parset := uint64((1 << PAR_OP))
  unpset := uint64((1 << NOT_OP) | (1 << PCT_OP) | (1 << CTZ_OP) | (1 << CLZ_OP))
  bnpset := uint64((1 << ADD_OP) | (1 << SUB_OP) | (1 << MUL_OP) | (1 << DIV_OP) |
                   (1 << SHL_OP) | (1 << SHR_OP) | (1 << RTR_OP) | (1 << RTL_OP) |
                   (1 << XOR_OP) | (1 << AND_OP) | (1 <<  OR_OP) | (1 << SET_OP) |
                   (1 << LEA0_OP))
  tnpset := uint64((1 << CMOV_OP)| (1 << LEA1_OP))
  retset := uint64((1 << RET_OP))

  for i := 0; i < p.Size; i++ {
    switch{
      0 != ((1 << p.Code[i].Op) & parset) :
        deps[i] |= (1 << p.Code[i].R0)

      0 != ((1 << p.Code[i].Op) & unpset) :
        deps[i] |= deps[p.Code[i].R0]

      0 != ((1 << p.Code[i].Op) & bnpset) :
        deps[i] |= deps[p.Code[i].R0] | deps[p.Code[i].R1]

      0 != ((1 << p.Code[i].Op) & tnpset) :
        deps[i] |= deps[p.Code[i].R0] | deps[p.Code[i].R1] | deps[p.Code[i].R2]

      0 != ((1 << p.Code[i].Op) & retset) :
        ret.Deps[p.Code[i].Op] = deps[p.Code[i].R0]

      p.Code[i].Op == MOD_OP :
        deps[i] |= deps[p.Code[p.Code[i].R0].R0] | deps[p.Code[p.Code[i].R0].R1]
    }
  }
  return ret
}
