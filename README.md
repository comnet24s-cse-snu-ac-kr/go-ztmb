# ZTMB (Zero Tunneling MiddleBox) w/o zkp

- Logic implementation without ZKP written in go

## Quickstart

### Demo

- Install prerequisites:

```bash
make deps
```

- Build and run an example input:

```bash
make
./ztmb ./example/input.json
```

### Output example (AES-256-GCM)

- Generated JSON file (name: `result.json`):

```json
{"key":["0","1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26","27","28","29","30","31"],"nonce":["0","1","2","3","4","5","6","7","8","9","10","11"],"packet":["36","215","1","0","0","1","0","0","0","0","0","1","63","98","87","70","106","76","84","89","48","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","72","86","116","89","87","77","116","77","84","73","52","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","50","104","63","104","77","105","48","121","78","84","89","115","97","71","49","104","89","121","49","122","97","71","69","121","76","84","85","120","77","105","120","111","98","87","70","106","76","88","78","111","89","84","69","115","97","71","49","104","89","121","49","116","90","68","85","116","90","88","82","116","81","71","57","119","90","87","63","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","109","108","119","90","87","49","107","77","84","89","119","76","87","86","48","98","85","66","118","99","71","86","117","99","51","78","111","76","109","78","118","98","83","120","111","98","87","70","106","76","88","78","111","89","20","84","69","116","79","84","89","116","90","88","82","116","81","71","57","119","90","87","53","122","99","1","56","1","102","1","49","5","49","51","57","52","48","6","116","117","110","110","101","108","7","101","120","97","109","112","108","101","3","111","114","103","0","0","5","0","1","0","12","0","247","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0"],"ciphertext":["99","213","215","27","197","228","194","27","141","65","151","138","142","171","47","43","201","154","243","77","192","42","56","69","111","61","146","176","103","10","50","213","84","105","156","197","219","173","122","206","32","253","40","160","220","234","92","81","218","40","39","180","45","140","212","239","69","244","24","94","77","154","199","215","164","80","161","18","135","168","81","12","104","23","152","38","208","231","67","161","8","211","167","54","15","16","18","47","204","181","231","149","73","170","216","234","174","161","53","178","241","25","165","115","230","83","124","229","46","63","77","39","244","8","27","255","134","182","100","253","124","183","46","148","179","178","52","74","206","6","90","158","230","131","186","20","49","182","50","246","2","239","22","161","4","147","150","76","7","40","231","38","235","23","188","14","173","159","157","173","222","219","201","79","190","206","167","19","243","26","170","150","168","214","182","249","200","140","188","65","236","253","6","255","86","177","175","139","60","171","193","145","96","52","197","184","90","187","188","165","79","56","188","191","174","53","156","240","174","10","161","34","71","120","232","165","0","214","163","176","27","31","186","160","0","34","236","174","128","33","50","32","99","116","129","60","133","23","40","4","198","221","220","57","154","117","241","196","90","59","11","22","136","243","48","230","20","209","247","144","156","218","68","62","128","223","114","7","183","45","58","10","218","180","202","205","128","139","129","175","113","47","109","81","10","244","186","75","1","3","222","72","186","197","206","23","146","60","58","90","10","149","160","211","152","202","91","94","73","148","207","23","152","49","29","22","151","61","69","233","201","149","127","42","139","96","120","233","163","224","123","32","93","211","29","98","147","26","215","132","156","20","133","8","73","163","120","102","141","177","72","255","89","47","249","217","133","84","114","203","32","175","164","35","151","181","98","40","132","253","31","121","238","241","0","193","159","229","52","220","220","87","89","50","165","131","192","217","135","7","211","220","161","112","254","127","62","64","18","135","226","11","108","187","159","32","82","212","61","16","222","70","77","230","187","86","25","122","87","107","195","209","81","114","160","184","220","82","255","248","176","165","93","107","68","18","28","131","32","229","231","185","227","237","134","171","31","120","122","122","174","221","174","87","191","125","235","26","224","87","241","148","213","81","105","138","107","82","195","154","30","60","67","197","186","176","17","255","134","180","25","239","41","1","151","147","223","206","81","170","4","141","232","172","237","128","144","172","209","154","61","174","225","32","183","2","145","93","199","232","83","28"],"counter":["0","0","0","2"]}
```

- STDOUT:

```
Header
  ID:        0x24d7
  Flags:     00000001 00000000
  QDCOUNT:   0x0001
  ANCOUNT:   0x0000
  NSCOUNT:   0x0000
  ARCOUNT:   0x0001
Question #0
Question
  QNMAE:     BWFJLty0Qg9WZw5zc2gUy29tlhVTYWMTmti4qG9wZw5zc2GUy29TlghTywmtC2H.hMI0YnTysAG1HyY1ZageYlTUXmixOBWfJlXNoyTesAg1HyY1TZDuTZXRTqG9WzW.5Zc2guy29tlgHTyWMtCMLWZW1kmTywlwV0BUBvcgVuC3noLMnvbsxOBWfjlxnoY.tEtOtYTzxRTqG9wZW5ZC.8.f.1.13940.TUnNeL.ExaMPle.org.
  QTYPE:     0x0005
  QCLASS:    0x0001
Remaining not-marshalled bytes

Additional Rerouces Record #0
RR OPT
  OPTCODE:  0x000c
  OPTLEN:   0x00f7
  PADDING:
    000 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    016 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    032 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    048 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    064 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    080 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    096 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    112 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    128 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    144 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    160 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    176 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    192 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    208 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    224 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00
    240 0x00 0x00 0x00 0x00 0x00 0x00 0x00

Cipher (AES-256-GCM)
  Key:                    000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f
  Nonce:                  000102030405060708090a0b
  PreCounterBlockSuffix:  00000002
  Length:                 512
  Tag:                    b70c0f8b6c0e4cc9db2805415e9ed31b
  Hex:
    000 0x63 0xd5 0xd7 0x1b 0xc5 0xe4 0xc2 0x1b 0x8d 0x41 0x97 0x8a 0x8e 0xab 0x2f 0x2b
    016 0xc9 0x9a 0xf3 0x4d 0xc0 0x2a 0x38 0x45 0x6f 0x3d 0x92 0xb0 0x67 0x0a 0x32 0xd5
    032 0x54 0x69 0x9c 0xc5 0xdb 0xad 0x7a 0xce 0x20 0xfd 0x28 0xa0 0xdc 0xea 0x5c 0x51
    048 0xda 0x28 0x27 0xb4 0x2d 0x8c 0xd4 0xef 0x45 0xf4 0x18 0x5e 0x4d 0x9a 0xc7 0xd7
    064 0xa4 0x50 0xa1 0x12 0x87 0xa8 0x51 0x0c 0x68 0x17 0x98 0x26 0xd0 0xe7 0x43 0xa1
    080 0x08 0xd3 0xa7 0x36 0x0f 0x10 0x12 0x2f 0xcc 0xb5 0xe7 0x95 0x49 0xaa 0xd8 0xea
    096 0xae 0xa1 0x35 0xb2 0xf1 0x19 0xa5 0x73 0xe6 0x53 0x7c 0xe5 0x2e 0x3f 0x4d 0x27
    112 0xf4 0x08 0x1b 0xff 0x86 0xb6 0x64 0xfd 0x7c 0xb7 0x2e 0x94 0xb3 0xb2 0x34 0x4a
    128 0xce 0x06 0x5a 0x9e 0xe6 0x83 0xba 0x14 0x31 0xb6 0x32 0xf6 0x02 0xef 0x16 0xa1
    144 0x04 0x93 0x96 0x4c 0x07 0x28 0xe7 0x26 0xeb 0x17 0xbc 0x0e 0xad 0x9f 0x9d 0xad
    160 0xde 0xdb 0xc9 0x4f 0xbe 0xce 0xa7 0x13 0xf3 0x1a 0xaa 0x96 0xa8 0xd6 0xb6 0xf9
    176 0xc8 0x8c 0xbc 0x41 0xec 0xfd 0x06 0xff 0x56 0xb1 0xaf 0x8b 0x3c 0xab 0xc1 0x91
    192 0x60 0x34 0xc5 0xb8 0x5a 0xbb 0xbc 0xa5 0x4f 0x38 0xbc 0xbf 0xae 0x35 0x9c 0xf0
    208 0xae 0x0a 0xa1 0x22 0x47 0x78 0xe8 0xa5 0x00 0xd6 0xa3 0xb0 0x1b 0x1f 0xba 0xa0
    224 0x00 0x22 0xec 0xae 0x80 0x21 0x32 0x20 0x63 0x74 0x81 0x3c 0x85 0x17 0x28 0x04
    240 0xc6 0xdd 0xdc 0x39 0x9a 0x75 0xf1 0xc4 0x5a 0x3b 0x0b 0x16 0x88 0xf3 0x30 0xe6
    256 0x14 0xd1 0xf7 0x90 0x9c 0xda 0x44 0x3e 0x80 0xdf 0x72 0x07 0xb7 0x2d 0x3a 0x0a
    272 0xda 0xb4 0xca 0xcd 0x80 0x8b 0x81 0xaf 0x71 0x2f 0x6d 0x51 0x0a 0xf4 0xba 0x4b
    288 0x01 0x03 0xde 0x48 0xba 0xc5 0xce 0x17 0x92 0x3c 0x3a 0x5a 0x0a 0x95 0xa0 0xd3
    304 0x98 0xca 0x5b 0x5e 0x49 0x94 0xcf 0x17 0x98 0x31 0x1d 0x16 0x97 0x3d 0x45 0xe9
    320 0xc9 0x95 0x7f 0x2a 0x8b 0x60 0x78 0xe9 0xa3 0xe0 0x7b 0x20 0x5d 0xd3 0x1d 0x62
    336 0x93 0x1a 0xd7 0x84 0x9c 0x14 0x85 0x08 0x49 0xa3 0x78 0x66 0x8d 0xb1 0x48 0xff
    352 0x59 0x2f 0xf9 0xd9 0x85 0x54 0x72 0xcb 0x20 0xaf 0xa4 0x23 0x97 0xb5 0x62 0x28
    368 0x84 0xfd 0x1f 0x79 0xee 0xf1 0x00 0xc1 0x9f 0xe5 0x34 0xdc 0xdc 0x57 0x59 0x32
    384 0xa5 0x83 0xc0 0xd9 0x87 0x07 0xd3 0xdc 0xa1 0x70 0xfe 0x7f 0x3e 0x40 0x12 0x87
    400 0xe2 0x0b 0x6c 0xbb 0x9f 0x20 0x52 0xd4 0x3d 0x10 0xde 0x46 0x4d 0xe6 0xbb 0x56
    416 0x19 0x7a 0x57 0x6b 0xc3 0xd1 0x51 0x72 0xa0 0xb8 0xdc 0x52 0xff 0xf8 0xb0 0xa5
    432 0x5d 0x6b 0x44 0x12 0x1c 0x83 0x20 0xe5 0xe7 0xb9 0xe3 0xed 0x86 0xab 0x1f 0x78
    448 0x7a 0x7a 0xae 0xdd 0xae 0x57 0xbf 0x7d 0xeb 0x1a 0xe0 0x57 0xf1 0x94 0xd5 0x51
    464 0x69 0x8a 0x6b 0x52 0xc3 0x9a 0x1e 0x3c 0x43 0xc5 0xba 0xb0 0x11 0xff 0x86 0xb4
    480 0x19 0xef 0x29 0x01 0x97 0x93 0xdf 0xce 0x51 0xaa 0x04 0x8d 0xe8 0xac 0xed 0x80
    496 0x90 0xac 0xd1 0x9a 0x3d 0xae 0xe1 0x20 0xb7 0x02 0x91 0x5d 0xc7 0xe8 0x53 0x1c
```

- Input/Output JSON file examples are available in [/example](./example) dir

### Additional options

- Testing

```bash
make test
```

- Formatting

```bash
make fmt
```

- Install binary

```bash
make install
```

### Pre-built binaries

- Pre-built binaries are available in [/build](./build) dir
    - `build/ztmb-darwin-amd64`: x86_64 binary
    - `build/ztmb-darwin-arm64`: ARM64 binary
