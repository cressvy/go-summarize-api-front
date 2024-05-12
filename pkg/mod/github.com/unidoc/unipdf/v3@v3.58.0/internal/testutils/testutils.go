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

package testutils ;import (_cg "crypto/md5";_bg "encoding/hex";_d "errors";_e "fmt";_cge "github.com/unidoc/unipdf/v3/common";_eb "github.com/unidoc/unipdf/v3/core";_gg "image";_cdc "image/png";_gf "io";_cdg "os";_a "os/exec";_g "path/filepath";_cd "strings";
_b "testing";);func CopyFile (src ,dst string )error {_de ,_cgb :=_cdg .Open (src );if _cgb !=nil {return _cgb ;};defer _de .Close ();_db ,_cgb :=_cdg .Create (dst );if _cgb !=nil {return _cgb ;};defer _db .Close ();_ ,_cgb =_gf .Copy (_db ,_de );return _cgb ;
};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_fbg :=_a .LookPath ("\u0067\u0073");_fbg !=nil {return ErrRenderNotSupported ;};return _a .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_e .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func ParseIndirectObjects (rawpdf string )(map[int64 ]_eb .PdfObject ,error ){_fge :=_eb .NewParserFromString (rawpdf );_fba :=map[int64 ]_eb .PdfObject {};for {_eac ,_fga :=_fge .ParseIndirectObject ();if _fga !=nil {if _fga ==_gf .EOF {break ;};return nil ,_fga ;
};switch _gge :=_eac .(type ){case *_eb .PdfIndirectObject :_fba [_gge .ObjectNumber ]=_eac ;case *_eb .PdfObjectStream :_fba [_gge .ObjectNumber ]=_eac ;};};for _ ,_ffb :=range _fba {_efe (_ffb ,_fba );};return _fba ,nil ;};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_bge ,_fb :=HashFile (file1 );
if _fb !=nil {return false ,_fb ;};_ee ,_fb :=HashFile (file2 );if _fb !=nil {return false ,_fb ;};if _bge ==_ee {return true ,nil ;};_fbf ,_fb :=ReadPNG (file1 );if _fb !=nil {return false ,_fb ;};_fg ,_fb :=ReadPNG (file2 );if _fb !=nil {return false ,_fb ;
};if _fbf .Bounds ()!=_fg .Bounds (){return false ,nil ;};return CompareImages (_fbf ,_fg );};var (ErrRenderNotSupported =_d .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
);func _efe (_bd _eb .PdfObject ,_gef map[int64 ]_eb .PdfObject )error {switch _bga :=_bd .(type ){case *_eb .PdfIndirectObject :_ddb :=_bga ;_efe (_ddb .PdfObject ,_gef );case *_eb .PdfObjectDictionary :_aab :=_bga ;for _ ,_ca :=range _aab .Keys (){_geb :=_aab .Get (_ca );
if _dda ,_geg :=_geb .(*_eb .PdfObjectReference );_geg {_ag ,_gcd :=_gef [_dda .ObjectNumber ];if !_gcd {return _d .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_aab .Set (_ca ,_ag );}else {_efe (_geb ,_gef );};};case *_eb .PdfObjectArray :_gedd :=_bga ;for _cb ,_dfa :=range _gedd .Elements (){if _ecf ,_cbb :=_dfa .(*_eb .PdfObjectReference );_cbb {_dfd ,_ecd :=_gef [_ecf .ObjectNumber ];if !_ecd {return _d .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_gedd .Set (_cb ,_dfd );}else {_efe (_dfa ,_gef );};};};return nil ;};func CompareImages (img1 ,img2 _gg .Image )(bool ,error ){_f :=img1 .Bounds ();_dd :=0;for _gc :=0;_gc < _f .Size ().X ;_gc ++{for _aa :=0;_aa < _f .Size ().Y ;_aa ++{_ec ,_ff ,_cdge ,_ :=img1 .At (_gc ,_aa ).RGBA ();
_fe ,_eca ,_gca ,_ :=img2 .At (_gc ,_aa ).RGBA ();if _ec !=_fe ||_ff !=_eca ||_cdge !=_gca {_dd ++;};};};_df :=float64 (_dd )/float64 (_f .Dx ()*_f .Dy ());if _df > 0.0001{_e .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_df ,_dd );
return false ,nil ;};return true ,nil ;};func RunRenderTest (t *_b .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_ede :=_cd .TrimSuffix (_g .Base (pdfPath ),_g .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_adg *_b .T ){_eg :=_g .Join (outputDir ,_ede );
_cdgb :=_eg +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _ge :=RenderPDFToPNGs (pdfPath ,0,_cdgb );_ge !=nil {_adg .Skip (_ge );};for _fd :=1;true ;_fd ++{_ea :=_e .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_eg ,_fd );_ef :=_g .Join (baselineRenderPath ,_e .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_ede ,_fd ));
if _ ,_eff :=_cdg .Stat (_ea );_eff !=nil {break ;};_adg .Logf ("\u0025\u0073",_ef );if saveBaseline {_adg .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_ea ,_ef );_edd :=CopyFile (_ea ,_ef );if _edd !=nil {_adg .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_ef ,_edd );
};continue ;};_adg .Run (_e .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_fd ),func (_af *_b .T ){_af .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_ea ,_ef );_ged ,_baf :=ComparePNGFiles (_ea ,_ef );
if _cdg .IsNotExist (_baf ){_af .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_ged {_af .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func CompareDictionariesDeep (d1 ,d2 *_eb .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_cge .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));
_cge .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_ecaa :=range d1 .Keys (){if _ecaa =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;
};_cga :=_eb .TraceToDirectObject (d1 .Get (_ecaa ));_bgf :=_eb .TraceToDirectObject (d2 .Get (_ecaa ));if _cga ==nil {_cge .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _bgf ==nil {_cge .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");
return false ;};switch _bff :=_cga .(type ){case *_eb .PdfObjectDictionary :_bgaa ,_bfa :=_bgf .(*_eb .PdfObjectDictionary );if !_bfa {_cge .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_cga ,_bgf );
return false ;};if !CompareDictionariesDeep (_bff ,_bgaa ){return false ;};continue ;case *_eb .PdfObjectArray :_fbb ,_gd :=_bgf .(*_eb .PdfObjectArray );if !_gd {_cge .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return false ;};if _bff .Len ()!=_fbb .Len (){_cge .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_bff .Len (),_fbb .Len ());
return false ;};for _egb :=0;_egb < _bff .Len ();_egb ++{_bde :=_eb .TraceToDirectObject (_bff .Get (_egb ));_fdd :=_eb .TraceToDirectObject (_fbb .Get (_egb ));if _gfd ,_da :=_bde .(*_eb .PdfObjectDictionary );_da {_ggec ,_ggb :=_fdd .(*_eb .PdfObjectDictionary );
if !_ggb {return false ;};if !CompareDictionariesDeep (_gfd ,_ggec ){return false ;};}else {if _bde .WriteString ()!=_fdd .WriteString (){_cge .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_bde .WriteString (),_fdd .WriteString ());
return false ;};};};continue ;};if _cga .String ()!=_bgf .String (){_cge .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_ecaa ,_cga .String (),_bgf .String ());
_cge .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_cga ,_bgf );_cge .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_cga ,_bgf );return false ;
};};return true ;};func ReadPNG (file string )(_gg .Image ,error ){_ed ,_bgg :=_cdg .Open (file );if _bgg !=nil {return nil ,_bgg ;};defer _ed .Close ();return _cdc .Decode (_ed );};func HashFile (file string )(string ,error ){_ga ,_ad :=_cdg .Open (file );
if _ad !=nil {return "",_ad ;};defer _ga .Close ();_ba :=_cg .New ();if _ ,_ad =_gf .Copy (_ba ,_ga );_ad !=nil {return "",_ad ;};return _bg .EncodeToString (_ba .Sum (nil )),nil ;};