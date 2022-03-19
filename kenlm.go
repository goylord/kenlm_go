package kenlm_go

/*
#cgo LDFLAGS: -lkenlm -lkenlm_builder -lkenlm_filter -lkenlm_util -llzma -lbz2 -lz -lstdc++  -ldl -lm
#include "kenlm.h"
*/
import "C"

import (
	"unsafe"
)

type KenlmModel struct {
	model unsafe.Pointer
}

func (km *KenlmModel) Score(sentence string) float64 {
	return float64(C.score(km.model, C.CString(sentence)))
}

func (km *KenlmModel) Perplexity(sentence string) float64 {
	return float64(C.score(km.model, C.CString(sentence)))
}

func LoadModel(filepath string) *KenlmModel {
	km_model := &KenlmModel{}
	km_model.model = C.load_model(C.CString(filepath))
	return km_model
}
