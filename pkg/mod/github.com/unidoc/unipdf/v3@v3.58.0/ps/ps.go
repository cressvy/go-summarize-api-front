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

// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
//
// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
package ps ;import (_fb "bufio";_d "bytes";_a "errors";_ce "fmt";_e "github.com/unidoc/unipdf/v3/common";_g "github.com/unidoc/unipdf/v3/core";_f "io";_c "math";);func (_bdb *PSParser )parseFunction ()(*PSProgram ,error ){_fab ,_ :=_bdb ._gce .ReadByte ();
if _fab !='{'{return nil ,_a .New ("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_faba :=NewPSProgram ();for {_bdb .skipSpaces ();_bdb .skipComments ();_ege ,_fafed :=_bdb ._gce .Peek (2);if _fafed !=nil {if _fafed ==_f .EOF {break ;
};return nil ,_fafed ;};_e .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_ege ));if _ege [0]=='}'{_e .Log .Trace ("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");_bdb ._gce .ReadByte ();
break ;}else if _ege [0]=='{'{_e .Log .Trace ("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!");_ebbge ,_dcbe :=_bdb .parseFunction ();if _dcbe !=nil {return nil ,_dcbe ;};_faba .Append (_ebbge );}else if _g .IsDecimalDigit (_ege [0])||(_ege [0]=='-'&&_g .IsDecimalDigit (_ege [1])){_e .Log .Trace ("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!");
_eadca ,_dadf :=_bdb .parseNumber ();if _dadf !=nil {return nil ,_dadf ;};_faba .Append (_eadca );}else {_e .Log .Trace ("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f");_ege ,_ =_bdb ._gce .Peek (5);_gagg :=string (_ege );
_e .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_gagg );if (len (_gagg )> 4)&&(_gagg [:5]=="\u0066\u0061\u006cs\u0065"){_gaga ,_cbaf :=_bdb .parseBool ();if _cbaf !=nil {return nil ,_cbaf ;};_faba .Append (_gaga );
}else if (len (_gagg )> 3)&&(_gagg [:4]=="\u0074\u0072\u0075\u0065"){_eafe ,_cbac :=_bdb .parseBool ();if _cbac !=nil {return nil ,_cbac ;};_faba .Append (_eafe );}else {_aacc ,_agbb :=_bdb .parseOperand ();if _agbb !=nil {return nil ,_agbb ;};_faba .Append (_aacc );
};};};return _faba ,nil ;};func (_fdce *PSOperand )ln (_bbd *PSStack )error {_egf ,_eba :=_bbd .PopNumberAsFloat64 ();if _eba !=nil {return _eba ;};_cbce :=_c .Log (_egf );_eba =_bbd .Push (MakeReal (_cbce ));return _eba ;};var ErrRangeCheck =_a .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");


// Empty empties the stack.
func (_cfd *PSStack )Empty (){*_cfd =[]PSObject {}};

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor (program *PSProgram )*PSExecutor {_df :=&PSExecutor {};_df .Stack =NewPSStack ();_df ._cd =program ;return _df ;};var ErrUndefinedResult =_a .New ("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072");
func (_eadc *PSOperand )exch (_abd *PSStack )error {_gdd ,_fagc :=_abd .Pop ();if _fagc !=nil {return _fagc ;};_eddaf ,_fagc :=_abd .Pop ();if _fagc !=nil {return _fagc ;};_fagc =_abd .Push (_gdd );if _fagc !=nil {return _fagc ;};_fagc =_abd .Push (_eddaf );
return _fagc ;};

// PSOperand represents a Postscript operand (text string).
type PSOperand string ;func (_dedg *PSOperand )ceiling (_gf *PSStack )error {_ebf ,_afd :=_gf .Pop ();if _afd !=nil {return _afd ;};if _fec ,_def :=_ebf .(*PSReal );_def {_afd =_gf .Push (MakeReal (_c .Ceil (_fec .Val )));}else if _bebc ,_gee :=_ebf .(*PSInteger );
_gee {_afd =_gf .Push (MakeInteger (_bebc .Val ));}else {_afd =ErrTypeCheck ;};return _afd ;};func (_egc *PSOperand )div (_ddgd *PSStack )error {_baa ,_fcf :=_ddgd .Pop ();if _fcf !=nil {return _fcf ;};_cdg ,_fcf :=_ddgd .Pop ();if _fcf !=nil {return _fcf ;
};_bad ,_ddff :=_baa .(*PSReal );_aada ,_cfa :=_baa .(*PSInteger );if !_ddff &&!_cfa {return ErrTypeCheck ;};if _ddff &&_bad .Val ==0{return ErrUndefinedResult ;};if _cfa &&_aada .Val ==0{return ErrUndefinedResult ;};_fdbe ,_aab :=_cdg .(*PSReal );_dbb ,_bda :=_cdg .(*PSInteger );
if !_aab &&!_bda {return ErrTypeCheck ;};var _dac float64 ;if _aab {_dac =_fdbe .Val ;}else {_dac =float64 (_dbb .Val );};if _ddff {_dac /=_bad .Val ;}else {_dac /=float64 (_aada .Val );};_fcf =_ddgd .Push (MakeReal (_dac ));return _fcf ;};func (_cab *PSOperand )cvi (_gac *PSStack )error {_age ,_cccf :=_gac .Pop ();
if _cccf !=nil {return _cccf ;};if _ddf ,_ba :=_age .(*PSReal );_ba {_fdc :=int (_ddf .Val );_cccf =_gac .Push (MakeInteger (_fdc ));}else if _ceca ,_edfe :=_age .(*PSInteger );_edfe {_ab :=_ceca .Val ;_cccf =_gac .Push (MakeInteger (_ab ));}else {return ErrTypeCheck ;
};return _cccf ;};

// PSInteger represents an integer.
type PSInteger struct{Val int ;};func (_aad *PSOperand )cvr (_cge *PSStack )error {_gacb ,_aadc :=_cge .Pop ();if _aadc !=nil {return _aadc ;};if _ac ,_afde :=_gacb .(*PSReal );_afde {_aadc =_cge .Push (MakeReal (_ac .Val ));}else if _bdc ,_afb :=_gacb .(*PSInteger );
_afb {_aadc =_cge .Push (MakeReal (float64 (_bdc .Val )));}else {return ErrTypeCheck ;};return _aadc ;};func (_bbdg *PSOperand )mod (_ddfc *PSStack )error {_dfc ,_ged :=_ddfc .Pop ();if _ged !=nil {return _ged ;};_acc ,_ged :=_ddfc .Pop ();if _ged !=nil {return _ged ;
};_acd ,_cbfb :=_dfc .(*PSInteger );if !_cbfb {return ErrTypeCheck ;};if _acd .Val ==0{return ErrUndefinedResult ;};_gff ,_cbfb :=_acc .(*PSInteger );if !_cbfb {return ErrTypeCheck ;};_dced :=_gff .Val %_acd .Val ;_ged =_ddfc .Push (MakeInteger (_dced ));
return _ged ;};

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand (val string )*PSOperand {_cgb :=PSOperand (val );return &_cgb };func (_dga *PSOperand )truncate (_ddeg *PSStack )error {_egcg ,_aacgd :=_ddeg .Pop ();if _aacgd !=nil {return _aacgd ;};if _aaba ,_fbee :=_egcg .(*PSReal );_fbee {_dfdf :=int (_aaba .Val );
_aacgd =_ddeg .Push (MakeReal (float64 (_dfdf )));}else if _bfbd ,_feb :=_egcg .(*PSInteger );_feb {_aacgd =_ddeg .Push (MakeInteger (_bfbd .Val ));}else {return ErrTypeCheck ;};return _aacgd ;};

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_egcb *PSStack )DebugString ()string {_aag :="\u005b\u0020";for _ ,_cbcb :=range *_egcb {_aag +=_cbcb .DebugString ();_aag +="\u0020";};_aag +="\u005d";return _aag ;};

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_ffg *PSStack )PopInteger ()(int ,error ){_cfaf ,_gagd :=_ffg .Pop ();if _gagd !=nil {return 0,_gagd ;};if _adec ,_faad :=_cfaf .(*PSInteger );_faad {return _adec .Val ,nil ;};return 0,ErrTypeCheck ;};func (_eafa *PSOperand )ne (_fdd *PSStack )error {_ebbg :=_eafa .eq (_fdd );
if _ebbg !=nil {return _ebbg ;};_ebbg =_eafa .not (_fdd );return _ebbg ;};

// Exec executes the operand `op` in the state specified by `stack`.
func (_gdgc *PSOperand )Exec (stack *PSStack )error {_ead :=ErrUnsupportedOperand ;switch *_gdgc {case "\u0061\u0062\u0073":_ead =_gdgc .abs (stack );case "\u0061\u0064\u0064":_ead =_gdgc .add (stack );case "\u0061\u006e\u0064":_ead =_gdgc .and (stack );
case "\u0061\u0074\u0061\u006e":_ead =_gdgc .atan (stack );case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":_ead =_gdgc .bitshift (stack );case "\u0063e\u0069\u006c\u0069\u006e\u0067":_ead =_gdgc .ceiling (stack );case "\u0063\u006f\u0070\u0079":_ead =_gdgc .copy (stack );
case "\u0063\u006f\u0073":_ead =_gdgc .cos (stack );case "\u0063\u0076\u0069":_ead =_gdgc .cvi (stack );case "\u0063\u0076\u0072":_ead =_gdgc .cvr (stack );case "\u0064\u0069\u0076":_ead =_gdgc .div (stack );case "\u0064\u0075\u0070":_ead =_gdgc .dup (stack );
case "\u0065\u0071":_ead =_gdgc .eq (stack );case "\u0065\u0078\u0063\u0068":_ead =_gdgc .exch (stack );case "\u0065\u0078\u0070":_ead =_gdgc .exp (stack );case "\u0066\u006c\u006fo\u0072":_ead =_gdgc .floor (stack );case "\u0067\u0065":_ead =_gdgc .ge (stack );
case "\u0067\u0074":_ead =_gdgc .gt (stack );case "\u0069\u0064\u0069\u0076":_ead =_gdgc .idiv (stack );case "\u0069\u0066":_ead =_gdgc .ifCondition (stack );case "\u0069\u0066\u0065\u006c\u0073\u0065":_ead =_gdgc .ifelse (stack );case "\u0069\u006e\u0064e\u0078":_ead =_gdgc .index (stack );
case "\u006c\u0065":_ead =_gdgc .le (stack );case "\u006c\u006f\u0067":_ead =_gdgc .log (stack );case "\u006c\u006e":_ead =_gdgc .ln (stack );case "\u006c\u0074":_ead =_gdgc .lt (stack );case "\u006d\u006f\u0064":_ead =_gdgc .mod (stack );case "\u006d\u0075\u006c":_ead =_gdgc .mul (stack );
case "\u006e\u0065":_ead =_gdgc .ne (stack );case "\u006e\u0065\u0067":_ead =_gdgc .neg (stack );case "\u006e\u006f\u0074":_ead =_gdgc .not (stack );case "\u006f\u0072":_ead =_gdgc .or (stack );case "\u0070\u006f\u0070":_ead =_gdgc .pop (stack );case "\u0072\u006f\u0075n\u0064":_ead =_gdgc .round (stack );
case "\u0072\u006f\u006c\u006c":_ead =_gdgc .roll (stack );case "\u0073\u0069\u006e":_ead =_gdgc .sin (stack );case "\u0073\u0071\u0072\u0074":_ead =_gdgc .sqrt (stack );case "\u0073\u0075\u0062":_ead =_gdgc .sub (stack );case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":_ead =_gdgc .truncate (stack );
case "\u0078\u006f\u0072":_ead =_gdgc .xor (stack );};return _ead ;};func (_dcfa *PSOperand )not (_dcec *PSStack )error {_eeeg ,_fad :=_dcec .Pop ();if _fad !=nil {return _fad ;};if _bfc ,_cfg :=_eeeg .(*PSBoolean );_cfg {_fad =_dcec .Push (MakeBool (!_bfc .Val ));
return _fad ;}else if _faag ,_ggff :=_eeeg .(*PSInteger );_ggff {_fad =_dcec .Push (MakeInteger (^_faag .Val ));return _fad ;}else {return ErrTypeCheck ;};};func (_cefg *PSOperand )floor (_edc *PSStack )error {_cecd ,_abe :=_edc .Pop ();if _abe !=nil {return _abe ;
};if _bag ,_cdbg :=_cecd .(*PSReal );_cdbg {_abe =_edc .Push (MakeReal (_c .Floor (_bag .Val )));}else if _aacb ,_bebg :=_cecd .(*PSInteger );_bebg {_abe =_edc .Push (MakeInteger (_aacb .Val ));}else {return ErrTypeCheck ;};return _abe ;};func (_cec *PSReal )DebugString ()string {return _ce .Sprintf ("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f",_cec .Val );
};func (_bfa *PSOperand )sub (_dfa *PSStack )error {_cfb ,_ade :=_dfa .Pop ();if _ade !=nil {return _ade ;};_eab ,_ade :=_dfa .Pop ();if _ade !=nil {return _ade ;};_cae ,_agb :=_cfb .(*PSReal );_aefe ,_accc :=_cfb .(*PSInteger );if !_agb &&!_accc {return ErrTypeCheck ;
};_dgfe ,_geeg :=_eab .(*PSReal );_egfe ,_bccf :=_eab .(*PSInteger );if !_geeg &&!_bccf {return ErrTypeCheck ;};if _accc &&_bccf {_dfe :=_egfe .Val -_aefe .Val ;_bed :=_dfa .Push (MakeInteger (_dfe ));return _bed ;};var _edg float64 =0;if _geeg {_edg =_dgfe .Val ;
}else {_edg =float64 (_egfe .Val );};if _agb {_edg -=_cae .Val ;}else {_edg -=float64 (_aefe .Val );};_ade =_dfa .Push (MakeReal (_edg ));return _ade ;};func (_eeg *PSOperand )bitshift (_edf *PSStack )error {_aga ,_dab :=_edf .PopInteger ();if _dab !=nil {return _dab ;
};_ebe ,_dab :=_edf .PopInteger ();if _dab !=nil {return _dab ;};var _cbd int ;if _aga >=0{_cbd =_ebe <<uint (_aga );}else {_cbd =_ebe >>uint (-_aga );};_dab =_edf .Push (MakeInteger (_cbd ));return _dab ;};func (_ccc *PSOperand )copy (_dca *PSStack )error {_cdc ,_cgc :=_dca .PopInteger ();
if _cgc !=nil {return _cgc ;};if _cdc < 0{return ErrRangeCheck ;};if _cdc > len (*_dca ){return ErrRangeCheck ;};*_dca =append (*_dca ,(*_dca )[len (*_dca )-_cdc :]...);return nil ;};func (_deb *PSOperand )sin (_aeac *PSStack )error {_dfd ,_dgcg :=_aeac .PopNumberAsFloat64 ();
if _dgcg !=nil {return _dgcg ;};_bbe :=_c .Sin (_dfd *_c .Pi /180.0);_dgcg =_aeac .Push (MakeReal (_bbe ));return _dgcg ;};func (_eddafb *PSOperand )lt (_aea *PSStack )error {_fecd ,_bbbd :=_aea .PopNumberAsFloat64 ();if _bbbd !=nil {return _bbbd ;};_gebf ,_bbbd :=_aea .PopNumberAsFloat64 ();
if _bbbd !=nil {return _bbbd ;};if _c .Abs (_gebf -_fecd )< _bc {_efga :=_aea .Push (MakeBool (false ));return _efga ;}else if _gebf < _fecd {_cecdc :=_aea .Push (MakeBool (true ));return _cecdc ;}else {_bcg :=_aea .Push (MakeBool (false ));return _bcg ;
};};var ErrStackOverflow =_a .New ("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077");func (_eff *PSOperand )dup (_eeb *PSStack )error {_cbdb ,_fag :=_eeb .Pop ();if _fag !=nil {return _fag ;};_fag =_eeb .Push (_cbdb );if _fag !=nil {return _fag ;
};_fag =_eeb .Push (_cbdb .Duplicate ());return _fag ;};

// Exec executes the program, typically leaving output values on the stack.
func (_dde *PSProgram )Exec (stack *PSStack )error {for _ ,_gg :=range *_dde {var _add error ;switch _cg :=_gg .(type ){case *PSInteger :_edd :=_cg ;_add =stack .Push (_edd );case *PSReal :_gbb :=_cg ;_add =stack .Push (_gbb );case *PSBoolean :_aaa :=_cg ;
_add =stack .Push (_aaa );case *PSProgram :_efb :=_cg ;_add =stack .Push (_efb );case *PSOperand :_gge :=_cg ;_add =_gge .Exec (stack );default:return ErrTypeCheck ;};if _add !=nil {return _add ;};};return nil ;};const _bc =0.000001;func (_cf *PSReal )Duplicate ()PSObject {_fbb :=PSReal {};
_fbb .Val =_cf .Val ;return &_fbb };func (_gacc *PSOperand )pop (_aaag *PSStack )error {_ ,_gbaf :=_aaag .Pop ();if _gbaf !=nil {return _gbaf ;};return nil ;};

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array (objects []PSObject )([]float64 ,error ){var _bg []float64 ;for _ ,_ea :=range objects {if _de ,_fe :=_ea .(*PSInteger );_fe {_bg =append (_bg ,float64 (_de .Val ));}else if _ee ,_ed :=_ea .(*PSReal );_ed {_bg =append (_bg ,_ee .Val );
}else {return nil ,ErrTypeCheck ;};};return _bg ,nil ;};

// PSBoolean represents a boolean value.
type PSBoolean struct{Val bool ;};func (_fee *PSParser )parseOperand ()(*PSOperand ,error ){var _bfd []byte ;for {_eebd ,_ecbc :=_fee ._gce .Peek (1);if _ecbc !=nil {if _ecbc ==_f .EOF {break ;};return nil ,_ecbc ;};if _g .IsDelimiter (_eebd [0]){break ;
};if _g .IsWhiteSpace (_eebd [0]){break ;};_gga ,_ :=_fee ._gce .ReadByte ();_bfd =append (_bfd ,_gga );};if len (_bfd )==0{return nil ,_a .New ("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029");
};return MakeOperand (string (_bfd )),nil ;};func (_bee *PSOperand )and (_ddd *PSStack )error {_agd ,_dee :=_ddd .Pop ();if _dee !=nil {return _dee ;};_bd ,_dee :=_ddd .Pop ();if _dee !=nil {return _dee ;};if _eea ,_deef :=_agd .(*PSBoolean );_deef {_gbe ,_aaac :=_bd .(*PSBoolean );
if !_aaac {return ErrTypeCheck ;};_dee =_ddd .Push (MakeBool (_eea .Val &&_gbe .Val ));return _dee ;};if _gdb ,_beg :=_agd .(*PSInteger );_beg {_dgbd ,_dgbc :=_bd .(*PSInteger );if !_dgbc {return ErrTypeCheck ;};_dee =_ddd .Push (MakeInteger (_gdb .Val &_dgbd .Val ));
return _dee ;};return ErrTypeCheck ;};

// Parse parses the postscript and store as a program that can be executed.
func (_caf *PSParser )Parse ()(*PSProgram ,error ){_caf .skipSpaces ();_fefe ,_aca :=_caf ._gce .Peek (2);if _aca !=nil {return nil ,_aca ;};if _fefe [0]!='{'{return nil ,_a .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b");
};_dbf ,_aca :=_caf .parseFunction ();if _aca !=nil &&_aca !=_f .EOF {return nil ,_aca ;};return _dbf ,_aca ;};func (_gcc *PSOperand )mul (_cefa *PSStack )error {_facb ,_bbf :=_cefa .Pop ();if _bbf !=nil {return _bbf ;};_dgfb ,_bbf :=_cefa .Pop ();if _bbf !=nil {return _bbf ;
};_aabc ,_afc :=_facb .(*PSReal );_dcb ,_eag :=_facb .(*PSInteger );if !_afc &&!_eag {return ErrTypeCheck ;};_bga ,_cdcb :=_dgfb .(*PSReal );_bfbg ,_ecd :=_dgfb .(*PSInteger );if !_cdcb &&!_ecd {return ErrTypeCheck ;};if _eag &&_ecd {_cgef :=_dcb .Val *_bfbg .Val ;
_aeg :=_cefa .Push (MakeInteger (_cgef ));return _aeg ;};var _eee float64 ;if _afc {_eee =_aabc .Val ;}else {_eee =float64 (_dcb .Val );};if _cdcb {_eee *=_bga .Val ;}else {_eee *=float64 (_bfbg .Val );};_bbf =_cefa .Push (MakeReal (_eee ));return _bbf ;
};func (_faf *PSOperand )index (_eaf *PSStack )error {_daa ,_deg :=_eaf .Pop ();if _deg !=nil {return _deg ;};_ceaf ,_gef :=_daa .(*PSInteger );if !_gef {return ErrTypeCheck ;};if _ceaf .Val < 0{return ErrRangeCheck ;};if _ceaf .Val > len (*_eaf )-1{return ErrStackUnderflow ;
};_bcd :=(*_eaf )[len (*_eaf )-1-_ceaf .Val ];_deg =_eaf .Push (_bcd .Duplicate ());return _deg ;};func (_eef *PSProgram )String ()string {_cef :="\u007b\u0020";for _ ,_fd :=range *_eef {_cef +=_fd .String ();_cef +="\u0020";};_cef +="\u007d";return _cef ;
};

// PSObject represents a postscript object.
type PSObject interface{

// Duplicate makes a fresh copy of the PSObject.
Duplicate ()PSObject ;

// DebugString returns a descriptive representation of the PSObject with more information than String()
// for debugging purposes.
DebugString ()string ;

// String returns a string representation of the PSObject.
String ()string ;};func (_fef *PSOperand )ifelse (_gfe *PSStack )error {_afa ,_acf :=_gfe .Pop ();if _acf !=nil {return _acf ;};_ccg ,_acf :=_gfe .Pop ();if _acf !=nil {return _acf ;};_efc ,_acf :=_gfe .Pop ();if _acf !=nil {return _acf ;};_bdag ,_fgd :=_afa .(*PSProgram );
if !_fgd {return ErrTypeCheck ;};_dfb ,_fgd :=_ccg .(*PSProgram );if !_fgd {return ErrTypeCheck ;};_badb ,_fgd :=_efc .(*PSBoolean );if !_fgd {return ErrTypeCheck ;};if _badb .Val {_cdgf :=_dfb .Exec (_gfe );return _cdgf ;};_acf =_bdag .Exec (_gfe );return _acf ;
};func (_bge *PSOperand )abs (_aaf *PSStack )error {_ge ,_bff :=_aaf .Pop ();if _bff !=nil {return _bff ;};if _gdc ,_ded :=_ge .(*PSReal );_ded {_dbe :=_gdc .Val ;if _dbe < 0{_bff =_aaf .Push (MakeReal (-_dbe ));}else {_bff =_aaf .Push (MakeReal (_dbe ));
};}else if _dg ,_ggd :=_ge .(*PSInteger );_ggd {_fdb :=_dg .Val ;if _fdb < 0{_bff =_aaf .Push (MakeInteger (-_fdb ));}else {_bff =_aaf .Push (MakeInteger (_fdb ));};}else {return ErrTypeCheck ;};return _bff ;};var ErrStackUnderflow =_a .New ("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077");
func (_fc *PSOperand )cos (_fgb *PSStack )error {_gec ,_agf :=_fgb .PopNumberAsFloat64 ();if _agf !=nil {return _agf ;};_ebfd :=_c .Cos (_gec *_c .Pi /180.0);_agf =_fgb .Push (MakeReal (_ebfd ));return _agf ;};func (_eda *PSOperand )exp (_dcg *PSStack )error {_eebe ,_gc :=_dcg .PopNumberAsFloat64 ();
if _gc !=nil {return _gc ;};_fbbc ,_gc :=_dcg .PopNumberAsFloat64 ();if _gc !=nil {return _gc ;};if _c .Abs (_eebe )< 1&&_fbbc < 0{return ErrUndefinedResult ;};_beba :=_c .Pow (_fbbc ,_eebe );_gc =_dcg .Push (MakeReal (_beba ));return _gc ;};func (_cde *PSOperand )neg (_defe *PSStack )error {_abf ,_fbeg :=_defe .Pop ();
if _fbeg !=nil {return _fbeg ;};if _geef ,_ggfa :=_abf .(*PSReal );_ggfa {_fbeg =_defe .Push (MakeReal (-_geef .Val ));return _fbeg ;}else if _dbca ,_gbg :=_abf .(*PSInteger );_gbg {_fbeg =_defe .Push (MakeInteger (-_dbca .Val ));return _fbeg ;}else {return ErrTypeCheck ;
};};

// PSParser is a basic Postscript parser.
type PSParser struct{_gce *_fb .Reader };func (_fbga *PSOperand )Duplicate ()PSObject {_ag :=*_fbga ;return &_ag };func (_cbc *PSOperand )eq (_cdb *PSStack )error {_bca ,_cdda :=_cdb .Pop ();if _cdda !=nil {return _cdda ;};_cecb ,_cdda :=_cdb .Pop ();if _cdda !=nil {return _cdda ;
};_ggf ,_ffdd :=_bca .(*PSBoolean );_fbbe ,_fbe :=_cecb .(*PSBoolean );if _ffdd ||_fbe {var _fbc error ;if _ffdd &&_fbe {_fbc =_cdb .Push (MakeBool (_ggf .Val ==_fbbe .Val ));}else {_fbc =_cdb .Push (MakeBool (false ));};return _fbc ;};var _dbg float64 ;
var _ecf float64 ;if _fac ,_eeaf :=_bca .(*PSInteger );_eeaf {_dbg =float64 (_fac .Val );}else if _geb ,_cad :=_bca .(*PSReal );_cad {_dbg =_geb .Val ;}else {return ErrTypeCheck ;};if _gacd ,_fdga :=_cecb .(*PSInteger );_fdga {_ecf =float64 (_gacd .Val );
}else if _eaa ,_gdgf :=_cecb .(*PSReal );_gdgf {_ecf =_eaa .Val ;}else {return ErrTypeCheck ;};if _c .Abs (_ecf -_dbg )< _bc {_cdda =_cdb .Push (MakeBool (true ));}else {_cdda =_cdb .Push (MakeBool (false ));};return _cdda ;};func (_ccgb *PSOperand )or (_dcgf *PSStack )error {_ccb ,_defef :=_dcgf .Pop ();
if _defef !=nil {return _defef ;};_cbcc ,_defef :=_dcgf .Pop ();if _defef !=nil {return _defef ;};if _gdce ,_fga :=_ccb .(*PSBoolean );_fga {_cca ,_dad :=_cbcc .(*PSBoolean );if !_dad {return ErrTypeCheck ;};_defef =_dcgf .Push (MakeBool (_gdce .Val ||_cca .Val ));
return _defef ;};if _cabg ,_efcb :=_ccb .(*PSInteger );_efcb {_dcgb ,_caba :=_cbcc .(*PSInteger );if !_caba {return ErrTypeCheck ;};_defef =_dcgf .Push (MakeInteger (_cabg .Val |_dcgb .Val ));return _defef ;};return ErrTypeCheck ;};func _fafa (_daf int )int {if _daf < 0{return -_daf ;
};return _daf ;};

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser (content []byte )*PSParser {_ebfc :=PSParser {};_ecc :=_d .NewBuffer (content );_ebfc ._gce =_fb .NewReader (_ecc );return &_ebfc ;};func (_fbd *PSReal )String ()string {return _ce .Sprintf ("\u0025\u002e\u0035\u0066",_fbd .Val )};

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger (val int )*PSInteger {_aagg :=PSInteger {};_aagg .Val =val ;return &_aagg };func (_ad *PSInteger )String ()string {return _ce .Sprintf ("\u0025\u0064",_ad .Val )};func (_ace *PSOperand )roll (_eggg *PSStack )error {_egd ,_gfc :=_eggg .Pop ();
if _gfc !=nil {return _gfc ;};_ggfg ,_gfc :=_eggg .Pop ();if _gfc !=nil {return _gfc ;};_aacg ,_egcf :=_egd .(*PSInteger );if !_egcf {return ErrTypeCheck ;};_fdaf ,_egcf :=_ggfg .(*PSInteger );if !_egcf {return ErrTypeCheck ;};if _fdaf .Val < 0{return ErrRangeCheck ;
};if _fdaf .Val ==0||_fdaf .Val ==1{return nil ;};if _fdaf .Val > len (*_eggg ){return ErrStackUnderflow ;};for _dgd :=0;_dgd < _fafa (_aacg .Val );_dgd ++{var _gcce []PSObject ;_gcce =(*_eggg )[len (*_eggg )-(_fdaf .Val ):len (*_eggg )];if _aacg .Val > 0{_gda :=_gcce [len (_gcce )-1];
_gcce =append ([]PSObject {_gda },_gcce [0:len (_gcce )-1]...);}else {_aef :=_gcce [len (_gcce )-_fdaf .Val ];_gcce =append (_gcce [1:],_aef );};_ffc :=append ((*_eggg )[0:len (*_eggg )-_fdaf .Val ],_gcce ...);_eggg =&_ffc ;};return nil ;};func (_fg *PSProgram )DebugString ()string {_ff :="\u007b\u0020";
for _ ,_cce :=range *_fg {_ff +=_cce .DebugString ();_ff +="\u0020";};_ff +="\u007d";return _ff ;};

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject ;

// String returns a string representation of the stack.
func (_egff *PSStack )String ()string {_cgee :="\u005b\u0020";for _ ,_gefc :=range *_egff {_cgee +=_gefc .String ();_cgee +="\u0020";};_cgee +="\u005d";return _cgee ;};

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_baf *PSStack )PopNumberAsFloat64 ()(float64 ,error ){_gbeg ,_cdba :=_baf .Pop ();if _cdba !=nil {return 0,_cdba ;};if _bgd ,_ecgg :=_gbeg .(*PSReal );_ecgg {return _bgd .Val ,nil ;}else if _eegb ,_cfdf :=_gbeg .(*PSInteger );_cfdf {return float64 (_eegb .Val ),nil ;
}else {return 0,ErrTypeCheck ;};};

// Push pushes an object on top of the stack.
func (_agga *PSStack )Push (obj PSObject )error {if len (*_agga )> 100{return ErrStackOverflow ;};*_agga =append (*_agga ,obj );return nil ;};func (_fba *PSBoolean )String ()string {return _ce .Sprintf ("\u0025\u0076",_fba .Val )};func (_bf *PSBoolean )Duplicate ()PSObject {_gdg :=PSBoolean {};
_gdg .Val =_bf .Val ;return &_gdg };func (_bdd *PSOperand )sqrt (_bbg *PSStack )error {_acff ,_ddb :=_bbg .PopNumberAsFloat64 ();if _ddb !=nil {return _ddb ;};if _acff < 0{return ErrRangeCheck ;};_adb :=_c .Sqrt (_acff );_ddb =_bbg .Push (MakeReal (_adb ));
return _ddb ;};func (_bbb *PSOperand )le (_aaff *PSStack )error {_gcf ,_aaaca :=_aaff .PopNumberAsFloat64 ();if _aaaca !=nil {return _aaaca ;};_abc ,_aaaca :=_aaff .PopNumberAsFloat64 ();if _aaaca !=nil {return _aaaca ;};if _c .Abs (_abc -_gcf )< _bc {_fafe :=_aaff .Push (MakeBool (true ));
return _fafe ;}else if _abc < _gcf {_faa :=_aaff .Push (MakeBool (true ));return _faa ;}else {_fda :=_aaff .Push (MakeBool (false ));return _fda ;};};func (_bcc *PSInteger )Duplicate ()PSObject {_ef :=PSInteger {};_ef .Val =_bcc .Val ;return &_ef };func (_eb *PSProgram )Duplicate ()PSObject {_dba :=&PSProgram {};
for _ ,_efg :=range *_eb {_dba .Append (_efg .Duplicate ());};return _dba ;};func (_ga *PSInteger )DebugString ()string {return _ce .Sprintf ("\u0069\u006e\u0074\u003a\u0025\u0064",_ga .Val );};func (_gcb *PSOperand )log (_fff *PSStack )error {_ebb ,_gacf :=_fff .PopNumberAsFloat64 ();
if _gacf !=nil {return _gacf ;};_cbf :=_c .Log10 (_ebb );_gacf =_fff .Push (MakeReal (_cbf ));return _gacf ;};func (_be *PSOperand )add (_agg *PSStack )error {_ca ,_ffd :=_agg .Pop ();if _ffd !=nil {return _ffd ;};_ae ,_ffd :=_agg .Pop ();if _ffd !=nil {return _ffd ;
};_ebc ,_beb :=_ca .(*PSReal );_ec ,_cdd :=_ca .(*PSInteger );if !_beb &&!_cdd {return ErrTypeCheck ;};_eg ,_aec :=_ae .(*PSReal );_af ,_dbc :=_ae .(*PSInteger );if !_aec &&!_dbc {return ErrTypeCheck ;};if _cdd &&_dbc {_agc :=_ec .Val +_af .Val ;_dgb :=_agg .Push (MakeInteger (_agc ));
return _dgb ;};var _ddg float64 ;if _beb {_ddg =_ebc .Val ;}else {_ddg =float64 (_ec .Val );};if _aec {_ddg +=_eg .Val ;}else {_ddg +=float64 (_af .Val );};_ffd =_agg .Push (MakeReal (_ddg ));return _ffd ;};

// Append appends an object to the PSProgram.
func (_gb *PSProgram )Append (obj PSObject ){*_gb =append (*_gb ,obj )};var ErrTypeCheck =_a .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");func (_fbg *PSOperand )DebugString ()string {return _ce .Sprintf ("\u006fp\u003a\u0027\u0025\u0073\u0027",*_fbg );
};

// NewPSStack returns an initialized PSStack.
func NewPSStack ()*PSStack {return &PSStack {}};

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct{Stack *PSStack ;_cd *PSProgram ;};func (_bbfd *PSParser )parseNumber ()(PSObject ,error ){_gdgdc ,_febe :=_g .ParseNumber (_bbfd ._gce );if _febe !=nil {return nil ,_febe ;};switch _deff :=_gdgdc .(type ){case *_g .PdfObjectFloat :return MakeReal (float64 (*_deff )),nil ;
case *_g .PdfObjectInteger :return MakeInteger (int (*_deff )),nil ;};return nil ,_ce .Errorf ("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054",_gdgdc );};func (_cfc *PSParser )skipSpaces ()(int ,error ){_bffc :=0;
for {_bcf ,_agcg :=_cfc ._gce .Peek (1);if _agcg !=nil {return 0,_agcg ;};if _g .IsWhiteSpace (_bcf [0]){_cfc ._gce .ReadByte ();_bffc ++;}else {break ;};};return _bffc ,nil ;};

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject ;

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool (val bool )*PSBoolean {_dcd :=PSBoolean {};_dcd .Val =val ;return &_dcd };func (_edda *PSOperand )String ()string {return string (*_edda )};func (_cea *PSBoolean )DebugString ()string {return _ce .Sprintf ("\u0062o\u006f\u006c\u003a\u0025\u0076",_cea .Val );
};

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_aa *PSExecutor )Execute (objects []PSObject )([]PSObject ,error ){for _ ,_dc :=range objects {_db :=_aa .Stack .Push (_dc );if _db !=nil {return nil ,_db ;};};_cc :=_aa ._cd .Exec (_aa .Stack );if _cc !=nil {_e .Log .Debug ("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_cc );
return nil ,_cc ;};_fa :=[]PSObject (*_aa .Stack );_aa .Stack .Empty ();return _fa ,nil ;};func (_cdcbd *PSParser )parseBool ()(*PSBoolean ,error ){_dfdg ,_cee :=_cdcbd ._gce .Peek (4);if _cee !=nil {return MakeBool (false ),_cee ;};if (len (_dfdg )>=4)&&(string (_dfdg [:4])=="\u0074\u0072\u0075\u0065"){_cdcbd ._gce .Discard (4);
return MakeBool (true ),nil ;};_dfdg ,_cee =_cdcbd ._gce .Peek (5);if _cee !=nil {return MakeBool (false ),_cee ;};if (len (_dfdg )>=5)&&(string (_dfdg [:5])=="\u0066\u0061\u006cs\u0065"){_cdcbd ._gce .Discard (5);return MakeBool (false ),nil ;};return MakeBool (false ),_a .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");
};func (_eegg *PSOperand )round (_fcg *PSStack )error {_gbba ,_faga :=_fcg .Pop ();if _faga !=nil {return _faga ;};if _bbfb ,_gbc :=_gbba .(*PSReal );_gbc {_faga =_fcg .Push (MakeReal (_c .Floor (_bbfb .Val +0.5)));}else if _fbad ,_agdc :=_gbba .(*PSInteger );
_agdc {_faga =_fcg .Push (MakeInteger (_fbad .Val ));}else {return ErrTypeCheck ;};return _faga ;};func (_ega *PSOperand )ifCondition (_gdgd *PSStack )error {_ebee ,_dcf :=_gdgd .Pop ();if _dcf !=nil {return _dcf ;};_bae ,_dcf :=_gdgd .Pop ();if _dcf !=nil {return _dcf ;
};_egg ,_fgc :=_ebee .(*PSProgram );if !_fgc {return ErrTypeCheck ;};_efbc ,_fgc :=_bae .(*PSBoolean );if !_fgc {return ErrTypeCheck ;};if _efbc .Val {_cgg :=_egg .Exec (_gdgd );return _cgg ;};return nil ;};func (_dgf *PSOperand )ge (_ffb *PSStack )error {_gcg ,_ceg :=_ffb .PopNumberAsFloat64 ();
if _ceg !=nil {return _ceg ;};_ecg ,_ceg :=_ffb .PopNumberAsFloat64 ();if _ceg !=nil {return _ceg ;};if _c .Abs (_ecg -_gcg )< _bc {_bfb :=_ffb .Push (MakeBool (true ));return _bfb ;}else if _ecg > _gcg {_dbbg :=_ffb .Push (MakeBool (true ));return _dbbg ;
}else {_eec :=_ffb .Push (MakeBool (false ));return _eec ;};};func (_dfeg *PSParser )skipComments ()error {if _ ,_ecce :=_dfeg .skipSpaces ();_ecce !=nil {return _ecce ;};_cggb :=true ;for {_abee ,_fffd :=_dfeg ._gce .Peek (1);if _fffd !=nil {_e .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_fffd .Error ());
return _fffd ;};if _cggb &&_abee [0]!='%'{return nil ;};_cggb =false ;if (_abee [0]!='\r')&&(_abee [0]!='\n'){_dfeg ._gce .ReadByte ();}else {break ;};};return _dfeg .skipComments ();};func (_fbgae *PSOperand )xor (_bba *PSStack )error {_bdfc ,_ggdb :=_bba .Pop ();
if _ggdb !=nil {return _ggdb ;};_dcc ,_ggdb :=_bba .Pop ();if _ggdb !=nil {return _ggdb ;};if _dege ,_eac :=_bdfc .(*PSBoolean );_eac {_caa ,_cfba :=_dcc .(*PSBoolean );if !_cfba {return ErrTypeCheck ;};_ggdb =_bba .Push (MakeBool (_dege .Val !=_caa .Val ));
return _ggdb ;};if _eeag ,_fafg :=_bdfc .(*PSInteger );_fafg {_deba ,_fca :=_dcc .(*PSInteger );if !_fca {return ErrTypeCheck ;};_ggdb =_bba .Push (MakeInteger (_eeag .Val ^_deba .Val ));return _ggdb ;};return ErrTypeCheck ;};var ErrUnsupportedOperand =_a .New ("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064");
func (_aed *PSOperand )idiv (_gab *PSStack )error {_afdg ,_bb :=_gab .Pop ();if _bb !=nil {return _bb ;};_aabg ,_bb :=_gab .Pop ();if _bb !=nil {return _bb ;};_bgea ,_bdf :=_afdg .(*PSInteger );if !_bdf {return ErrTypeCheck ;};if _bgea .Val ==0{return ErrUndefinedResult ;
};_bdg ,_bdf :=_aabg .(*PSInteger );if !_bdf {return ErrTypeCheck ;};_abb :=_bdg .Val /_bgea .Val ;_bb =_gab .Push (MakeInteger (_abb ));return _bb ;};

// Pop pops an object from the top of the stack.
func (_bde *PSStack )Pop ()(PSObject ,error ){if len (*_bde )< 1{return nil ,ErrStackUnderflow ;};_fbce :=(*_bde )[len (*_bde )-1];*_bde =(*_bde )[0:len (*_bde )-1];return _fbce ,nil ;};func (_gba *PSOperand )gt (_cfe *PSStack )error {_gdf ,_gaf :=_cfe .PopNumberAsFloat64 ();
if _gaf !=nil {return _gaf ;};_ada ,_gaf :=_cfe .PopNumberAsFloat64 ();if _gaf !=nil {return _gaf ;};if _c .Abs (_ada -_gdf )< _bc {_ggc :=_cfe .Push (MakeBool (false ));return _ggc ;}else if _ada > _gdf {_ecb :=_cfe .Push (MakeBool (true ));return _ecb ;
}else {_ebfb :=_cfe .Push (MakeBool (false ));return _ebfb ;};};

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal (val float64 )*PSReal {_dabb :=PSReal {};_dabb .Val =val ;return &_dabb };

// PSReal represents a real number.
type PSReal struct{Val float64 ;};func (_gag *PSOperand )atan (_bgf *PSStack )error {_ccd ,_dddb :=_bgf .PopNumberAsFloat64 ();if _dddb !=nil {return _dddb ;};_dgc ,_dddb :=_bgf .PopNumberAsFloat64 ();if _dddb !=nil {return _dddb ;};if _ccd ==0{var _fdg error ;
if _dgc < 0{_fdg =_bgf .Push (MakeReal (270));}else {_fdg =_bgf .Push (MakeReal (90));};return _fdg ;};_dce :=_dgc /_ccd ;_aac :=_c .Atan (_dce )*180/_c .Pi ;_dddb =_bgf .Push (MakeReal (_aac ));return _dddb ;};

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram ()*PSProgram {return &PSProgram {}};