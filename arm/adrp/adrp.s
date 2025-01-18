.global _main
.section __DATA,__data
data: .space 1024

.text

_main:
  adrp x0, data@PAGE
  add x0,x0, data@PAGEOFF

  mov x0, #0
  mov x16, #1
  svc #0x80
