//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package uuid ;import (_c "crypto/rand";_f "encoding/hex";_g "io";);func MustUUID ()UUID {uuid ,_ac :=NewUUID ();if _ac !=nil {panic (_ac );};return uuid ;};type UUID [16]byte ;var Nil =_b ;func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_cc :=_g .ReadFull (_ff ,uuid [:]);
if _cc !=nil {return _b ,_cc ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};func (_ab UUID )String ()string {var _ef [36]byte ;_ged (_ef [:],_ab );return string (_ef [:])};func _ged (_ee []byte ,_gd UUID ){_f .Encode (_ee ,_gd [:4]);
_ee [8]='-';_f .Encode (_ee [9:13],_gd [4:6]);_ee [13]='-';_f .Encode (_ee [14:18],_gd [6:8]);_ee [18]='-';_f .Encode (_ee [19:23],_gd [8:10]);_ee [23]='-';_f .Encode (_ee [24:],_gd [10:]);};var _b UUID ;var _ff =_c .Reader ;