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

package huffman ;import (_ge "errors";_g "fmt";_a "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_da "math";_c "strings";);type FixedSizeTable struct{_ag *InternalNode };func (_faa *EncodedTable )Decode (r *_a .Reader )(int64 ,error ){return _faa ._gf .Decode (r )};
func (_cgf *StandardTable )InitTree (codeTable []*Code )error {_fac (codeTable );for _ ,_aae :=range codeTable {if _dg :=_cgf ._edd .append (_aae );_dg !=nil {return _dg ;};};return nil ;};func _fac (_df []*Code ){var _dagc int32 ;for _ ,_caf :=range _df {_dagc =_cc (_dagc ,_caf ._gff );
};_eea :=make ([]int32 ,_dagc +1);for _ ,_ede :=range _df {_eea [_ede ._gff ]++;};var _ffb int32 ;_fdf :=make ([]int32 ,len (_eea )+1);_eea [0]=0;for _ce :=int32 (1);_ce <=int32 (len (_eea ));_ce ++{_fdf [_ce ]=(_fdf [_ce -1]+(_eea [_ce -1]))<<1;_ffb =_fdf [_ce ];
for _ ,_gaf :=range _df {if _gaf ._gff ==_ce {_gaf ._bbb =_ffb ;_ffb ++;};};};};func (_cb *ValueNode )Decode (r *_a .Reader )(int64 ,error ){_ffe ,_ed :=r .ReadBits (byte (_cb ._gef ));if _ed !=nil {return 0,_ed ;};if _cb ._ac {_ffe =-_ffe ;};return int64 (_cb ._cag )+int64 (_ffe ),nil ;
};func (_gcg *ValueNode )String ()string {return _g .Sprintf ("\u0025\u0064\u002f%\u0064",_gcg ._gef ,_gcg ._cag );};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_agg ){return nil ,_ge .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_fg :=_agg [number -1];if _fg ==nil {var _ffea error ;_fg ,_ffea =_ga (_deg [number -1]);if _ffea !=nil {return nil ,_ffea ;};_agg [number -1]=_fg ;};return _fg ,nil ;};func (_bff *FixedSizeTable )Decode (r *_a .Reader )(int64 ,error ){return _bff ._ag .Decode (r )};
func (_eg *OutOfBandNode )String ()string {return _g .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_da .MaxInt64 ));};var _ Node =&OutOfBandNode {};type Tabler interface{Decode (_fef *_a .Reader )(int64 ,error );InitTree (_ade []*Code )error ;String ()string ;
RootNode ()*InternalNode ;};func (_be *Code )String ()string {var _gfed string ;if _be ._bbb !=-1{_gfed =_gd (_be ._bbb ,_be ._gff );}else {_gfed ="\u003f";};return _g .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_gfed ,_be ._gff ,_be ._dde ,_be ._gfb );
};func (_ae *EncodedTable )parseTable ()error {var (_geb []*Code ;_cd ,_ba ,_eb int32 ;_cfd uint64 ;_bc error ;);_aee :=_ae .StreamReader ();_aaf :=_ae .HtLow ();for _aaf < _ae .HtHigh (){_cfd ,_bc =_aee .ReadBits (byte (_ae .HtPS ()));if _bc !=nil {return _bc ;
};_cd =int32 (_cfd );_cfd ,_bc =_aee .ReadBits (byte (_ae .HtRS ()));if _bc !=nil {return _bc ;};_ba =int32 (_cfd );_geb =append (_geb ,NewCode (_cd ,_ba ,_eb ,false ));_aaf +=1<<uint (_ba );};_cfd ,_bc =_aee .ReadBits (byte (_ae .HtPS ()));if _bc !=nil {return _bc ;
};_cd =int32 (_cfd );_ba =32;_eb =_ae .HtLow ()-1;_geb =append (_geb ,NewCode (_cd ,_ba ,_eb ,true ));_cfd ,_bc =_aee .ReadBits (byte (_ae .HtPS ()));if _bc !=nil {return _bc ;};_cd =int32 (_cfd );_ba =32;_eb =_ae .HtHigh ();_geb =append (_geb ,NewCode (_cd ,_ba ,_eb ,false ));
if _ae .HtOOB ()==1{_cfd ,_bc =_aee .ReadBits (byte (_ae .HtPS ()));if _bc !=nil {return _bc ;};_cd =int32 (_cfd );_geb =append (_geb ,NewCode (_cd ,-1,-1,false ));};if _bc =_ae .InitTree (_geb );_bc !=nil {return _bc ;};return nil ;};var _agg =make ([]Tabler ,len (_deg ));
func _fff (_bab int32 )*InternalNode {return &InternalNode {_dd :_bab }};func (_def *FixedSizeTable )String ()string {return _def ._ag .String ()+"\u000a"};func _ga (_bfd [][]int32 )(*StandardTable ,error ){var _cbf []*Code ;for _dgd :=0;_dgd < len (_bfd );
_dgd ++{_eef :=_bfd [_dgd ][0];_dc :=_bfd [_dgd ][1];_db :=_bfd [_dgd ][2];var _cfc bool ;if len (_bfd [_dgd ])> 3{_cfc =true ;};_cbf =append (_cbf ,NewCode (_eef ,_dc ,_db ,_cfc ));};_cfb :=&StandardTable {_edd :_fff (0)};if _gfe :=_cfb .InitTree (_cbf );
_gfe !=nil {return nil ,_gfe ;};return _cfb ,nil ;};func (_bba *OutOfBandNode )Decode (r *_a .Reader )(int64 ,error ){return 0,_b .ErrOOB };func _cc (_edf ,_ffbd int32 )int32 {if _edf > _ffbd {return _edf ;};return _ffbd ;};func (_ef *StandardTable )RootNode ()*InternalNode {return _ef ._edd };
var _deg =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};
type EncodedTable struct{BasicTabler ;_gf *InternalNode ;};func (_bdb *StandardTable )String ()string {return _bdb ._edd .String ()+"\u000a"};type StandardTable struct{_edd *InternalNode };type Node interface{Decode (_ad *_a .Reader )(int64 ,error );String ()string ;
};func (_fe *InternalNode )Decode (r *_a .Reader )(int64 ,error ){_eab ,_af :=r .ReadBit ();if _af !=nil {return 0,_af ;};if _eab ==0{return _fe ._dac .Decode (r );};return _fe ._gb .Decode (r );};func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_gff :prefixLength ,_dde :rangeLength ,_gfb :rangeLow ,_gebc :isLowerRange ,_bbb :-1};
};var _ Node =&ValueNode {};type ValueNode struct{_gef int32 ;_cag int32 ;_ac bool ;};type OutOfBandNode struct{};func (_bdf *StandardTable )Decode (r *_a .Reader )(int64 ,error ){return _bdf ._edd .Decode (r )};type Code struct{_gff int32 ;_dde int32 ;
_gfb int32 ;_gebc bool ;_bbb int32 ;};var _ Tabler =&EncodedTable {};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()*_a .Reader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func (_ea *FixedSizeTable )InitTree (codeTable []*Code )error {_fac (codeTable );
for _ ,_bb :=range codeTable {_ff :=_ea ._ag .append (_bb );if _ff !=nil {return _ff ;};};return nil ;};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_f :=&EncodedTable {_gf :&InternalNode {},BasicTabler :table };if _fa :=_f .parseTable ();
_fa !=nil {return nil ,_fa ;};return _f ,nil ;};func (_e *EncodedTable )String ()string {return _e ._gf .String ()+"\u000a"};type InternalNode struct{_dd int32 ;_dac Node ;_gb Node ;};func (_ca *EncodedTable )RootNode ()*InternalNode {return _ca ._gf };
func (_bd *EncodedTable )InitTree (codeTable []*Code )error {_fac (codeTable );for _ ,_cf :=range codeTable {if _ab :=_bd ._gf .append (_cf );_ab !=nil {return _ab ;};};return nil ;};func _dag (_cg *Code )*OutOfBandNode {return &OutOfBandNode {}};var _ Node =&InternalNode {};
func _dab (_bffe *Code )*ValueNode {return &ValueNode {_gef :_bffe ._dde ,_cag :_bffe ._gfb ,_ac :_bffe ._gebc };};func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_bf :=&FixedSizeTable {_ag :&InternalNode {}};if _de :=_bf .InitTree (codeTable );
_de !=nil {return nil ,_de ;};return _bf ,nil ;};func _gd (_afe ,_beg int32 )string {var _bea int32 ;_ged :=make ([]rune ,_beg );for _ecf :=int32 (1);_ecf <=_beg ;_ecf ++{_bea =_afe >>uint (_beg -_ecf )&1;if _bea !=0{_ged [_ecf -1]='1';}else {_ged [_ecf -1]='0';
};};return string (_ged );};func (_aafa *FixedSizeTable )RootNode ()*InternalNode {return _aafa ._ag };func (_bbae *InternalNode )pad (_ddg *_c .Builder ){for _daa :=int32 (0);_daa < _bbae ._dd ;_daa ++{_ddg .WriteString ("\u0020\u0020\u0020");};};func (_fdg *InternalNode )append (_gfd *Code )(_dad error ){if _gfd ._gff ==0{return nil ;
};_ee :=_gfd ._gff -1-_fdg ._dd ;if _ee < 0{return _ge .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_egg :=(_gfd ._bbb >>uint (_ee ))&0x1;
if _ee ==0{if _gfd ._dde ==-1{if _egg ==1{if _fdg ._gb !=nil {return _g .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_gfd );};_fdg ._gb =_dag (_gfd );
}else {if _fdg ._dac !=nil {return _g .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_gfd );};_fdg ._dac =_dag (_gfd );};}else {if _egg ==1{if _fdg ._gb !=nil {return _g .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_gfd );
};_fdg ._gb =_dab (_gfd );}else {if _fdg ._dac !=nil {return _g .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_gfd );
};_fdg ._dac =_dab (_gfd );};};}else {if _egg ==1{if _fdg ._gb ==nil {_fdg ._gb =_fff (_fdg ._dd +1);};if _dad =_fdg ._gb .(*InternalNode ).append (_gfd );_dad !=nil {return _dad ;};}else {if _fdg ._dac ==nil {_fdg ._dac =_fff (_fdg ._dd +1);};if _dad =_fdg ._dac .(*InternalNode ).append (_gfd );
_dad !=nil {return _dad ;};};};return nil ;};func (_geg *InternalNode )String ()string {_bcg :=&_c .Builder {};_bcg .WriteString ("\u000a");_geg .pad (_bcg );_bcg .WriteString ("\u0030\u003a\u0020");_bcg .WriteString (_geg ._dac .String ()+"\u000a");_geg .pad (_bcg );
_bcg .WriteString ("\u0031\u003a\u0020");_bcg .WriteString (_geg ._gb .String ()+"\u000a");return _bcg .String ();};