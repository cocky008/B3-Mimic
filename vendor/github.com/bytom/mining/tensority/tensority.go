package tensority

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/ -l:cSimdTs.o -lstdc++ -lgomp
// #include "./lib/cSimdTs.h"
import "C"

import(
    "unsafe"

    "github.com/bytom/protocol/bc"
)

func Hash(blockHeader, seed *bc.Hash) *bc.Hash {
    bhBytes := blockHeader.Bytes()
    sdBytes := seed.Bytes()

    // Get thearray pointer from the corresponding slice
    bhPtr := (*C.uint8_t)(unsafe.Pointer(&bhBytes[0]))
    seedPtr := (*C.uint8_t)(unsafe.Pointer(&sdBytes[0]))

    resPtr := C.SimdTs(bhPtr, seedPtr)
    
    res := bc.NewHash(*(*[32]byte)(unsafe.Pointer(resPtr)))
    return &res
}