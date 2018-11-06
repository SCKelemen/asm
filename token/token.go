package token

type Token uint8

const (
  NOP Token = iota // 0
  TCK // 1
  HLT // 2
  RST // 3
  
  // alu
  ZER Token = 15 // 15
  AL1 // 16
  AL2 // 17
  AL3 // 18
  AL4 // 19
  AL5 // 20 
  AL6 // 21 
  AL7 // 22 
  AL8 // 23
  AL9 // 24
  ALA // 25
  ALB // 26
  ALC // 27
  ALD // 28
  ALE // 29
  ONE // 30
  
  ADD
  SUB
  MUL
  DIV
  MOD = REM
  INC
  DEC
  
  
  
)


/* 
# b  Description
0 0
1 0
2 0
3 0   ALU bit

4 0
5 0
6 0
7 0

*/
