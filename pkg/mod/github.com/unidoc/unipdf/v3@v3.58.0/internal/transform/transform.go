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

package transform ;import (_g "fmt";_d "github.com/unidoc/unipdf/v3/common";_cc "math";);func (_dda Point )Distance (b Point )float64 {return _cc .Hypot (_dda .X -b .X ,_dda .Y -b .Y )};func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );
};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};func (_bcb Matrix )Scale (xScale ,yScale float64 )Matrix {return _bcb .Mult (ScaleMatrix (xScale ,yScale ));};func (_dc Matrix )Identity ()bool {return _dc [0]==1&&_dc [1]==0&&_dc [2]==0&&_dc [3]==0&&_dc [4]==1&&_dc [5]==0&&_dc [6]==0&&_dc [7]==0&&_dc [8]==1;
};func (_gab *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_bad :=NewMatrix (a ,b ,c ,d ,tx ,ty );_gab .transformByMatrix (_bad );};const _bg =1e-6;func (_ecd Point )String ()string {return _g .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_ecd .X ,_ecd .Y );
};func (_efa Matrix )Inverse ()(Matrix ,bool ){_acd ,_bae :=_efa [0],_efa [1];_db ,_cf :=_efa [3],_efa [4];_gg ,_ab :=_efa [6],_efa [7];_adg :=_acd *_cf -_bae *_db ;if _cc .Abs (_adg )< _ea {return Matrix {},false ;};_fdd ,_ge :=_cf /_adg ,-_bae /_adg ;
_af ,_caa :=-_db /_adg ,_acd /_adg ;_cgg :=-(_fdd *_gg +_af *_ab );_cbf :=-(_ge *_gg +_caa *_ab );return NewMatrix (_fdd ,_ge ,_af ,_caa ,_cgg ,_cbf ),true ;};func (_cd Matrix )Angle ()float64 {_gd :=_cc .Atan2 (-_cd [1],_cd [0]);if _gd < 0.0{_gd +=2*_cc .Pi ;
};return _gd /_cc .Pi *180.0;};func (_bcbd Matrix )Transform (x ,y float64 )(float64 ,float64 ){_ddd :=x *_bcbd [0]+y *_bcbd [3]+_bcbd [6];_cb :=x *_bcbd [1]+y *_bcbd [4]+_bcbd [7];return _ddd ,_cb ;};func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_fc :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};
_fc .clampRange ();return _fc ;};func (_bbc Matrix )Rotate (theta float64 )Matrix {return _bbc .Mult (RotationMatrix (theta ))};func (_fd Matrix )String ()string {_e ,_cg ,_ba ,_bb ,_bc ,_a :=_fd [0],_fd [1],_fd [3],_fd [4],_fd [6],_fd [7];return _g .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_e ,_cg ,_ba ,_bb ,_bc ,_a );
};func (_ebc Matrix )ScalingFactorX ()float64 {return _cc .Hypot (_ebc [0],_ebc [1])};func (_ef *Matrix )Concat (b Matrix ){*_ef =Matrix {b [0]*_ef [0]+b [1]*_ef [3],b [0]*_ef [1]+b [1]*_ef [4],0,b [3]*_ef [0]+b [4]*_ef [3],b [3]*_ef [1]+b [4]*_ef [4],0,b [6]*_ef [0]+b [7]*_ef [3]+_ef [6],b [6]*_ef [1]+b [7]*_ef [4]+_ef [7],1};
_ef .clampRange ();};func (_ce Matrix )Mult (b Matrix )Matrix {_ce .Concat (b );return _ce };func (_bd *Point )transformByMatrix (_eag Matrix ){_bd .X ,_bd .Y =_eag .Transform (_bd .X ,_bd .Y )};func (_eb Matrix )Singular ()bool {return _cc .Abs (_eb [0]*_eb [4]-_eb [1]*_eb [3])< _ed };
func (_dd *Matrix )Clone ()Matrix {return NewMatrix (_dd [0],_dd [1],_dd [3],_dd [4],_dd [6],_dd [7])};func (_dcc Matrix )Translation ()(float64 ,float64 ){return _dcc [6],_dcc [7]};func (_dad Matrix )ScalingFactorY ()float64 {return _cc .Hypot (_dad [3],_dad [4])};
func (_ad *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_ad [0],_ad [1]=a ,b ;_ad [3],_ad [4]=c ,d ;_ad [6],_ad [7]=tx ,ty ;_ad .clampRange ();};func (_efd Point )Displace (delta Point )Point {return Point {_efd .X +delta .X ,_efd .Y +delta .Y }};const _acg =1e9;
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};type Point struct{X float64 ;Y float64 ;};const _ea =1.0e-6;func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};func (_fdg Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_fdg .X +t *b .X ,Y :(1-t )*_fdg .Y +t *b .Y };
};func (_ac *Matrix )Shear (x ,y float64 ){_ac .Concat (ShearMatrix (x ,y ))};func (_fbd *Point )Set (x ,y float64 ){_fbd .X ,_fbd .Y =x ,y };func RotationMatrix (angle float64 )Matrix {_ga :=_cc .Cos (angle );_b :=_cc .Sin (angle );return NewMatrix (_ga ,_b ,-_b ,_ga ,0,0);
};func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};func (_ddb *Matrix )clampRange (){for _cfg ,_gee :=range _ddb {if _gee > _acg {_d .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gee ,_acg );
_ddb [_cfg ]=_acg ;}else if _gee < -_acg {_d .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gee ,-_acg );_ddb [_cfg ]=-_acg ;};};};func (_acgb Point )Rotate (theta float64 )Point {_ec :=_cc .Hypot (_acgb .X ,_acgb .Y );
_dbd :=_cc .Atan2 (_acgb .Y ,_acgb .X );_ddbd ,_dbde :=_cc .Sincos (_dbd +theta /180.0*_cc .Pi );return Point {_ec *_dbde ,_ec *_ddbd };};type Matrix [9]float64 ;const _ed =1e-10;func (_gb Matrix )Unrealistic ()bool {_ccc ,_fcg ,_fb ,_cab :=_cc .Abs (_gb [0]),_cc .Abs (_gb [1]),_cc .Abs (_gb [3]),_cc .Abs (_gb [4]);
_gbe :=_ccc > _bg &&_cab > _bg ;_df :=_fcg > _bg &&_fb > _bg ;return !(_gbe ||_df );};func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};func (_da Matrix )Translate (tx ,ty float64 )Matrix {return _da .Mult (TranslationMatrix (tx ,ty ))};
func (_f Matrix )Round (precision float64 )Matrix {for _ca :=range _f {_f [_ca ]=_cc .Round (_f [_ca ]/precision )*precision ;};return _f ;};