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

package pdfutil ;import (_c "github.com/unidoc/unipdf/v3/common";_g "github.com/unidoc/unipdf/v3/contentstream";_cf "github.com/unidoc/unipdf/v3/contentstream/draw";_f "github.com/unidoc/unipdf/v3/core";_d "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
//   - Normalize the page rotation.
//     Rotates the contents of the page according to the Rotate entry, thus
//     flattening the rotation. The Rotate entry of the page is set to nil.
//   - Normalize the media box.
//     If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//     the contents of the page are translated to (-Llx, -Lly). After
//     normalization, the media box is updated (Llx and Lly are set to 0 and
//     Urx and Ury are updated accordingly).
//   - Normalize the crop box.
//     The crop box of the page is updated based on the previous operations.
//
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_d .PdfPage )error {_fd ,_fda :=page .GetMediaBox ();if _fda !=nil {return _fda ;};_be ,_fda :=page .GetRotate ();if _fda !=nil {_c .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_fda .Error ());
};_a :=_be %360!=0&&_be %90==0;_fd .Normalize ();_ad ,_ac ,_cd ,_ba :=_fd .Llx ,_fd .Lly ,_fd .Width (),_fd .Height ();_baf :=_ad !=0||_ac !=0;if !_a &&!_baf {return nil ;};_ce :=func (_gb ,_ab ,_aa float64 )_cf .BoundingBox {return _cf .Path {Points :[]_cf .Point {_cf .NewPoint (0,0).Rotate (_aa ),_cf .NewPoint (_gb ,0).Rotate (_aa ),_cf .NewPoint (0,_ab ).Rotate (_aa ),_cf .NewPoint (_gb ,_ab ).Rotate (_aa )}}.GetBoundingBox ();
};_bg :=_g .NewContentCreator ();var _cc float64 ;if _a {_cc =-float64 (_be );_cee :=_ce (_cd ,_ba ,_cc );_bg .Translate ((_cee .Width -_cd )/2+_cd /2,(_cee .Height -_ba )/2+_ba /2);_bg .RotateDeg (_cc );_bg .Translate (-_cd /2,-_ba /2);_cd ,_ba =_cee .Width ,_cee .Height ;
};if _baf {_bg .Translate (-_ad ,-_ac );};_bgf :=_bg .Operations ();_cb ,_fda :=_f .MakeStream (_bgf .Bytes (),_f .NewFlateEncoder ());if _fda !=nil {return _fda ;};_de :=_f .MakeArray (_cb );_de .Append (page .GetContentStreamObjs ()...);*_fd =_d .PdfRectangle {Urx :_cd ,Ury :_ba };
if _bab :=page .CropBox ;_bab !=nil {_bab .Normalize ();_gd ,_fde ,_dg ,_gdg :=_bab .Llx -_ad ,_bab .Lly -_ac ,_bab .Width (),_bab .Height ();if _a {_acf :=_ce (_dg ,_gdg ,_cc );_dg ,_gdg =_acf .Width ,_acf .Height ;};*_bab =_d .PdfRectangle {Llx :_gd ,Lly :_fde ,Urx :_gd +_dg ,Ury :_fde +_gdg };
};_c .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_cc ,_bgf ,_fd );page .Contents =_de ;page .Rotate =nil ;
return nil ;};