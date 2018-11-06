# asm
Assembly language for a custom, home-built processor


## Operations
I'm not sure about the endianness yet.


| Oper | Binary    | Hex  | Description  | Note |
| ---- | ---------:| ----:| -------------| -----|
| NOP  | 0000 0000 | 0x00 | No Operation | Defualt/Zero codes do nothing |
| TCK  | 0000 0001 | 0x01 | Tick         | Tick, 0x01 can be added to PC to advance |
| HLT  | 0000 0010 | 0x02 | Halt         | Stop processing |
| RST  | 0000 0011 | 0x03 | Reset        | Reset all modules |
 
### ALU Ops
By using bits 0-4 to target ALU-bound instructions, we can have some nice properties where bit 4 tells us if the op is ALU bound, and bits 0-3 are the inputs to a multiplexer to configure the logic. 

| Oper | Binary    | Hex  | Description  | Note |
| ---- | ---------:| ----:| -------------| -----|
| ZER  | 0001 0000 | 0x10 | Zero         | outputs zero always |
| ???  | 0001 0001 | 0x11 |              |      |
| ???  | 0001 0010 | 0x12 |              |      |
| ???  | 0001 0011 | 0x13 |              |      |
| ???  | 0001 0100 | 0x14 |              |      |
| ???  | 0001 0101 | 0x15 |              |      |
| ???  | 0001 0110 | 0x16 |              |      |
| ???  | 0001 0111 | 0x17 |              |      |
| ???  | 0001 1000 | 0x18 |              |      |
| ???  | 0001 1001 | 0x19 |              |      |
| ???  | 0001 1010 | 0x1A |              |      |
| ???  | 0001 1011 | 0x1B |              |      |
| ???  | 0001 1100 | 0x1C |              |      |
| ???  | 0001 1101 | 0x1D |              |      |
| ???  | 0001 1110 | 0x1E |              |      |
| ???  | 0001 1111 | 0x1F | One          | outputs one always |

### Memory Ops
Hopefully I can do a similar thing with another bit for memory access.

 
## ALU

### Logic

| b | a | 3 | 2 | 1 | 0 | q | label |
| --| --| --| --| --| --| --| ------|
| 0 | 0 | 0 | 0 | 0 | 0 | ? | ???   |
| 0 | 0 | 0 | 0 | 0 | 1 | ? | ???   |
| 0 | 0 | 0 | 0 | 1 | 0 | ? | ???   |
| 0 | 0 | 0 | 0 | 1 | 1 | ? | ???   |
| 0 | 0 | 0 | 1 | 0 | 0 | ? | ???   |
| 0 | 0 | 0 | 1 | 0 | 1 | ? | ???   |
| 0 | 0 | 0 | 1 | 1 | 0 | ? | ???   |
| 0 | 0 | 0 | 1 | 1 | 1 | ? | ???   |
| 0 | 0 | 1 | 0 | 0 | 0 | ? | ???   |
| 0 | 0 | 1 | 0 | 0 | 1 | ? | ???   |
| 0 | 0 | 1 | 0 | 1 | 0 | ? | ???   |
| 0 | 0 | 1 | 0 | 1 | 1 | ? | ???   |
| 0 | 0 | 1 | 1 | 0 | 0 | ? | ???   |
| 0 | 0 | 1 | 1 | 0 | 1 | ? | ???   |
| 0 | 0 | 1 | 1 | 1 | 0 | ? | ???   |
| 0 | 0 | 1 | 1 | 1 | 1 | ? | ???   |
|   |   |   |   |   |   |   |       |
| 0 | 1 | 0 | 0 | 0 | 0 | ? | ???   |
| 0 | 1 | 0 | 0 | 0 | 1 | ? | ???   |
| 0 | 1 | 0 | 0 | 1 | 0 | ? | ???   |
| 0 | 1 | 0 | 0 | 1 | 1 | ? | ???   |
| 0 | 1 | 0 | 1 | 0 | 0 | ? | ???   |
| 0 | 1 | 0 | 1 | 0 | 1 | ? | ???   |
| 0 | 1 | 0 | 1 | 1 | 0 | ? | ???   |
| 0 | 1 | 0 | 1 | 1 | 1 | ? | ???   |
| 0 | 1 | 1 | 0 | 0 | 0 | ? | ???   |
| 0 | 1 | 1 | 0 | 0 | 1 | ? | ???   |
| 0 | 1 | 1 | 0 | 1 | 0 | ? | ???   |
| 0 | 1 | 1 | 0 | 1 | 1 | ? | ???   |
| 0 | 1 | 1 | 1 | 0 | 0 | ? | ???   |
| 0 | 1 | 1 | 1 | 0 | 1 | ? | ???   |
| 0 | 1 | 1 | 1 | 1 | 0 | ? | ???   |
| 0 | 1 | 1 | 1 | 1 | 1 | ? | ???   |
|   |   |   |   |   |   |   |       |
| 1 | 0 | 0 | 0 | 0 | 0 | ? | ???   |
| 1 | 0 | 0 | 0 | 0 | 1 | ? | ???   |
| 1 | 0 | 0 | 0 | 1 | 0 | ? | ???   |
| 1 | 0 | 0 | 0 | 1 | 1 | ? | ???   |
| 1 | 0 | 0 | 1 | 0 | 0 | ? | ???   |
| 1 | 0 | 0 | 1 | 0 | 1 | ? | ???   |
| 1 | 0 | 0 | 1 | 1 | 0 | ? | ???   |
| 1 | 0 | 0 | 1 | 1 | 1 | ? | ???   |
| 1 | 0 | 1 | 0 | 0 | 0 | ? | ???   |
| 1 | 0 | 1 | 0 | 0 | 1 | ? | ???   |
| 1 | 0 | 1 | 0 | 1 | 0 | ? | ???   |
| 1 | 0 | 1 | 0 | 1 | 1 | ? | ???   |
| 1 | 0 | 1 | 1 | 0 | 0 | ? | ???   |
| 1 | 0 | 1 | 1 | 0 | 1 | ? | ???   |
| 1 | 0 | 1 | 1 | 1 | 0 | ? | ???   |
| 1 | 0 | 1 | 1 | 1 | 1 | ? | ???   |
|   |   |   |   |   |   |   |       |
| 1 | 1 | 0 | 0 | 0 | 0 | ? | ???   |
| 1 | 1 | 0 | 0 | 0 | 1 | ? | ???   |
| 1 | 1 | 0 | 0 | 1 | 0 | ? | ???   |
| 1 | 1 | 0 | 0 | 1 | 1 | ? | ???   |
| 1 | 1 | 0 | 1 | 0 | 0 | ? | ???   |
| 1 | 1 | 0 | 1 | 0 | 1 | ? | ???   |
| 1 | 1 | 0 | 1 | 1 | 0 | ? | ???   |
| 1 | 1 | 0 | 1 | 1 | 1 | ? | ???   |
| 1 | 1 | 1 | 0 | 0 | 0 | ? | ???   |
| 1 | 1 | 1 | 0 | 0 | 1 | ? | ???   |
| 1 | 1 | 1 | 0 | 1 | 0 | ? | ???   |
| 1 | 1 | 1 | 0 | 1 | 1 | ? | ???   |
| 1 | 1 | 1 | 1 | 0 | 0 | ? | ???   |
| 1 | 1 | 1 | 1 | 0 | 1 | ? | ???   |
| 1 | 1 | 1 | 1 | 1 | 0 | ? | ???   |
| 1 | 1 | 1 | 1 | 1 | 1 | ? | ???   |


| Oper  | Description | alt |
| ------| ------------| ----|
| add   | add | |
| sub   | subtract | |
| mul   | multiply | |
| div   | divide | |
| mod   | modulo/remainder | rem |
| inc   | increment | |
| dec   | decrement | |
| or    | logical or | |
| nor   | logical not or | |
| xor   | logical xor | |
| nxor  | logical not xor | |
| and   | logical and | |
| nand  | logical not and | |
| inv   | invert | | 
| lsft  | left shift | |
| rshft | right shift | | 
| hlt   | halt | stp |
| rst   | reset | |
| tck   | tick | adv / clk |
| cll   | call | exc / exe / fun / fnc / call / func / exec |
| ret   | return | rtn |
| jz    | jump if zero | |
| jc    | jump conditional / jump carry | |


Going to need some sort of load statements...
and some fetch/decode statements...



## Control Word

### Flags

| flag | name  | desc |
| -----| ------| -----|
| zf   | zero  | zero flag |
| cf   | carry | carry flag |
