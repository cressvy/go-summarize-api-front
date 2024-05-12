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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_ge "bytes";_gb "crypto";_fg "crypto/rand";_de "crypto/rsa";_fe "crypto/x509";_e "crypto/x509/pkix";_cd "encoding/asn1";_fc "encoding/hex";_b "errors";_bb "fmt";_gdf "github.com/unidoc/pkcs7";_fcc "github.com/unidoc/timestamp";
_gd "github.com/unidoc/unipdf/v3/common";_dg "github.com/unidoc/unipdf/v3/core";_dgd "github.com/unidoc/unipdf/v3/model";_fa "github.com/unidoc/unipdf/v3/model/mdp";_fag "github.com/unidoc/unipdf/v3/model/sigutil";_a "hash";_f "math/big";_g "strings";_c "time";
);func _eabg (_gga []byte ,_deb int )(_baa []byte ){_ffg :=len (_gga );if _ffg > _deb {_ffg =_deb ;};_baa =make ([]byte ,_deb );copy (_baa [len (_baa )-_ffg :],_gga );return ;};func _daed (_gda *_de .PublicKey ,_bga []byte )_gb .Hash {_ecf :=_gda .Size ();
if _ecf !=len (_bga ){return 0;};_beg :=func (_db *_f .Int ,_gecb *_de .PublicKey ,_bag *_f .Int )*_f .Int {_agfb :=_f .NewInt (int64 (_gecb .E ));_db .Exp (_bag ,_agfb ,_gecb .N );return _db ;};_gaec :=new (_f .Int ).SetBytes (_bga );_gcf :=_beg (new (_f .Int ),_gda ,_gaec );
_geff :=_eabg (_gcf .Bytes (),_ecf );if _geff [0]!=0||_geff [1]!=1{return 0;};_bge :=[]struct{Hash _gb .Hash ;Prefix []byte ;}{{Hash :_gb .SHA1 ,Prefix :[]byte {0x30,0x21,0x30,0x09,0x06,0x05,0x2b,0x0e,0x03,0x02,0x1a,0x05,0x00,0x04,0x14}},{Hash :_gb .SHA256 ,Prefix :[]byte {0x30,0x31,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x01,0x05,0x00,0x04,0x20}},{Hash :_gb .SHA384 ,Prefix :[]byte {0x30,0x41,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x02,0x05,0x00,0x04,0x30}},{Hash :_gb .SHA512 ,Prefix :[]byte {0x30,0x51,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x03,0x05,0x00,0x04,0x40}},{Hash :_gb .RIPEMD160 ,Prefix :[]byte {0x30,0x20,0x30,0x08,0x06,0x06,0x28,0xcf,0x06,0x03,0x00,0x31,0x04,0x14}}};
for _ ,_afcg :=range _bge {_cfcd :=_afcg .Hash .Size ();_fbf :=len (_afcg .Prefix )+_cfcd ;if _ge .Equal (_geff [_ecf -_fbf :_ecf -_cfcd ],_afcg .Prefix ){return _afcg .Hash ;};};return 0;};type etsiPAdES struct{_fba *_de .PrivateKey ;_cdabd *_fe .Certificate ;
_da bool ;_gg bool ;_ade *_fe .Certificate ;_eag string ;

// CertClient is the client used to retrieve certificates.
CertClient *_fag .CertClient ;

// OCSPClient is the client used to retrieve OCSP validation information.
OCSPClient *_fag .OCSPClient ;

// CRLClient is the client used to retrieve CRL validation information.
CRLClient *_fag .CRLClient ;_cfd *_dgd .PdfAppender ;_eg *_dgd .DSS ;};type adobePKCS7Detached struct{_dca *_de .PrivateKey ;_cga *_fe .Certificate ;_fabf bool ;_cea int ;};

// Validate validates PdfSignature.
func (_ddd *adobePKCS7Detached )Validate (sig *_dgd .PdfSignature ,digest _dgd .Hasher )(_dgd .SignatureValidationResult ,error ){_fed :=sig .Contents .Bytes ();_cdc ,_fdc :=_gdf .Parse (_fed );if _fdc !=nil {return _dgd .SignatureValidationResult {},_fdc ;
};_efe ,_abb :=digest .(*_ge .Buffer );if !_abb {return _dgd .SignatureValidationResult {},_bb .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_cdc .Content =_efe .Bytes ();if _fdc =_cdc .Verify ();
_fdc !=nil {return _dgd .SignatureValidationResult {},_fdc ;};return _dgd .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};func (_cfdfd *adobePKCS7Detached )getCertificate (_fefa *_dgd .PdfSignature )(*_fe .Certificate ,error ){if _cfdfd ._cga !=nil {return _cfdfd ._cga ,nil ;
};_gac ,_fff :=_fefa .GetCerts ();if _fff !=nil {return nil ,_fff ;};return _gac [0],nil ;};

// Validate validates PdfSignature.
func (_cfc *etsiPAdES )Validate (sig *_dgd .PdfSignature ,digest _dgd .Hasher )(_dgd .SignatureValidationResult ,error ){_aed :=sig .Contents .Bytes ();_eaf ,_fcb :=_gdf .Parse (_aed );if _fcb !=nil {return _dgd .SignatureValidationResult {},_fcb ;};_caf ,_dea :=digest .(*_ge .Buffer );
if !_dea {return _dgd .SignatureValidationResult {},_bb .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_eaf .Content =_caf .Bytes ();if _fcb =_eaf .Verify ();_fcb !=nil {return _dgd .SignatureValidationResult {},_fcb ;
};_cac :=false ;_ddc :=false ;var _ecc _c .Time ;for _ ,_fadf :=range _eaf .Signers {_edb :=_fadf .EncryptedDigest ;var _bfd RevocationInfoArchival ;_fcb =_eaf .UnmarshalSignedAttribute (_gdf .OIDAttributeAdobeRevocation ,&_bfd );if _fcb ==nil {if len (_bfd .Crl )> 0{_ddc =true ;
};if len (_bfd .Ocsp )> 0{_cac =true ;};};for _ ,_abd :=range _fadf .UnauthenticatedAttributes {if _abd .Type .Equal (_gdf .OIDAttributeTimeStampToken ){_gdfb ,_ggc :=_fcc .Parse (_abd .Value .Bytes );if _ggc !=nil {return _dgd .SignatureValidationResult {},_ggc ;
};_ecc =_gdfb .Time ;_efc :=_gdfb .HashAlgorithm .New ();_efc .Write (_edb );if !_ge .Equal (_efc .Sum (nil ),_gdfb .HashedMessage ){return _dgd .SignatureValidationResult {},_bb .Errorf ("\u0048\u0061\u0073\u0068\u0020i\u006e\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u0069s\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u006b\u0063\u0073\u0037");
};break ;};};};_cgg :=_dgd .SignatureValidationResult {IsSigned :true ,IsVerified :true ,IsCrlFound :_ddc ,IsOcspFound :_cac ,GeneralizedTime :_ecc };return _cgg ,nil ;};

// NewEtsiPAdESLevelT creates a new Adobe.PPKLite ETSI.CAdES.detached Level T signature handler.
func NewEtsiPAdESLevelT (privateKey *_de .PrivateKey ,certificate *_fe .Certificate ,caCert *_fe .Certificate ,certificateTimestampServerURL string )(_dgd .SignatureHandler ,error ){return &etsiPAdES {_cdabd :certificate ,_fba :privateKey ,_ade :caCert ,_eag :certificateTimestampServerURL },nil ;
};func (_ccg *etsiPAdES )getCRLs (_cag []*_fe .Certificate )([][]byte ,error ){_bef :=make ([][]byte ,0,len (_cag ));for _ ,_gec :=range _cag {for _ ,_cce :=range _gec .CRLDistributionPoints {if _ccg .CertClient .IsCA (_gec ){continue ;};_dfc ,_acf :=_ccg .CRLClient .MakeRequest (_cce ,_gec );
if _acf !=nil {_gd .Log .Debug ("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076",_acf );continue ;};_bef =append (_bef ,_dfc );};};return _bef ,nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_bdd *_dgd .PdfSignature ,_gcb _dgd .Hasher )([]byte ,error );

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _gb .Hash )(_dgd .SignatureHandler ,error ){return &docTimeStamp {_dcag :timestampServerURL ,_ebe :hashAlgorithm },nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_fe .Certificate ,signFunc SignFunc )(_dgd .SignatureHandler ,error ){return &adobeX509RSASHA1 {_fae :certificate ,_gbca :signFunc },nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_ead *docTimeStamp )Sign (sig *_dgd .PdfSignature ,digest _dgd .Hasher )error {_bbbba ,_gdb :=_fag .NewTimestampRequest (digest .(*_ge .Buffer ),&_fcc .RequestOptions {Hash :_ead ._ebe ,Certificates :true });if _gdb !=nil {return _gdb ;};_ebda :=_ead ._cbg ;
if _ebda ==nil {_ebda =_fag .NewTimestampClient ();};_gcda ,_gdb :=_ebda .GetEncodedToken (_ead ._dcag ,_bbbba );if _gdb !=nil {return _gdb ;};_fca :=len (_gcda );if _ead ._bfcb > 0&&_fca > _ead ._bfcb {return _dgd .ErrSignNotEnoughSpace ;};if _fca > 0{_ead ._bfcb =_fca +128;
};if sig .Contents !=nil {_dbdd :=sig .Contents .Bytes ();copy (_dbdd ,_gcda );_gcda =_dbdd ;};sig .Contents =_dg .MakeHexString (string (_gcda ));return nil ;};

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;

// Client is the timestamp client used to make the signature request.
// If no client is provided, a default one is used.
Client *_fag .TimestampClient ;};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_de .PrivateKey ,certificate *_fe .Certificate )(_dgd .SignatureHandler ,error ){return &adobePKCS7Detached {_cga :certificate ,_dca :privateKey },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_acg *adobePKCS7Detached )IsApplicable (sig *_dgd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_dfg *DocMDPHandler )ValidateWithOpts (sig *_dgd .PdfSignature ,digest _dgd .Hasher ,params _dgd .SignatureHandlerDocMDPParams )(_dgd .SignatureValidationResult ,error ){_ea ,_ga :=_dfg ._eb .Validate (sig ,digest );if _ga !=nil {return _ea ,_ga ;
};_aa :=params .Parser ;if _aa ==nil {return _dgd .SignatureValidationResult {},_b .New ("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c");};if !_ea .IsVerified {return _ea ,nil ;};_ab :=params .DiffPolicy ;
if _ab ==nil {_ab =_fa .NewDefaultDiffPolicy ();};for _fec :=0;_fec <=_aa .GetRevisionNumber ();_fec ++{_fcd ,_ec :=_aa .GetRevision (_fec );if _ec !=nil {return _dgd .SignatureValidationResult {},_ec ;};_ebg :=_fcd .GetTrailer ();if _ebg ==nil {return _dgd .SignatureValidationResult {},_b .New ("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074");
};_gc ,_ca :=_dg .GetDict (_ebg .Get ("\u0052\u006f\u006f\u0074"));if !_ca {return _dgd .SignatureValidationResult {},_b .New ("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t");};
_cda ,_ca :=_dg .GetDict (_gc .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));if !_ca {continue ;};_ef ,_ca :=_dg .GetArray (_cda .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_ca {continue ;};for _ ,_ff :=range _ef .Elements (){_cdab ,_faga :=_dg .GetDict (_ff );
if !_faga {continue ;};_fdb ,_faga :=_dg .GetDict (_cdab .Get ("\u0056"));if !_faga {continue ;};if _dg .EqualObjects (_fdb .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"),sig .Contents ){_ea .DiffResults ,_ec =_ab .ReviewFile (_fcd ,_aa ,&_fa .MDPParameters {DocMDPLevel :_dfg .Permission });
if _ec !=nil {return _dgd .SignatureValidationResult {},_ec ;};_ea .IsVerified =_ea .DiffResults .IsPermitted ();return _ea ,nil ;};};};return _dgd .SignatureValidationResult {},_b .New ("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073");
};

// InitSignature initialization of the DocMDP signature.
func (_ac *DocMDPHandler )InitSignature (sig *_dgd .PdfSignature )error {_ece :=_ac ._eb .InitSignature (sig );if _ece !=nil {return _ece ;};sig .Handler =_ac ;if sig .Reference ==nil {sig .Reference =_dg .MakeArray ();};sig .Reference .Append (_dgd .NewPdfSignatureReferenceDocMDP (_dgd .NewPdfTransformParamsDocMDP (_ac .Permission )).ToPdfObject ());
return nil ;};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_fe .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_dgd .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_fae :certificate ,_gbca :signFunc ,_aff :opts .EstimateSize ,_ddcb :opts .Algorithm },nil ;
};func (_cc *etsiPAdES )makeTimestampRequest (_dcb string ,_fde []byte )(_cd .RawValue ,error ){_af :=_gb .SHA512 .New ();_af .Write (_fde );_eaa :=_af .Sum (nil );_afb :=_fcc .Request {HashAlgorithm :_gb .SHA512 ,HashedMessage :_eaa ,Certificates :true ,Extensions :nil ,ExtraExtensions :nil };
_gba :=_fag .NewTimestampClient ();_ecg ,_ed :=_gba .GetEncodedToken (_dcb ,&_afb );if _ed !=nil {return _cd .NullRawValue ,_ed ;};return _cd .RawValue {FullBytes :_ecg },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_fegb *adobeX509RSASHA1 )IsApplicable (sig *_dgd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};type docTimeStamp struct{_dcag string ;_ebe _gb .Hash ;_bfcb int ;_cbg *_fag .TimestampClient ;};

// NewDigest creates a new digest.
func (_bdfc *docTimeStamp )NewDigest (sig *_dgd .PdfSignature )(_dgd .Hasher ,error ){return _ge .NewBuffer (nil ),nil ;};

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level B signature handler.
func NewEtsiPAdESLevelB (privateKey *_de .PrivateKey ,certificate *_fe .Certificate ,caCert *_fe .Certificate )(_dgd .SignatureHandler ,error ){return &etsiPAdES {_cdabd :certificate ,_fba :privateKey ,_ade :caCert },nil ;};

// NewDigest creates a new digest.
func (_bbb *DocMDPHandler )NewDigest (sig *_dgd .PdfSignature )(_dgd .Hasher ,error ){return _bbb ._eb .NewDigest (sig );};

// Sign sets the Contents fields for the PdfSignature.
func (_gcc *etsiPAdES )Sign (sig *_dgd .PdfSignature ,digest _dgd .Hasher )error {_dac ,_egg :=digest .(*_ge .Buffer );if !_egg {return _bb .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_cee ,_bgc :=_gdf .NewSignedData (_dac .Bytes ());
if _bgc !=nil {return _bgc ;};_cee .SetDigestAlgorithm (_gdf .OIDDigestAlgorithmSHA256 );_bfb :=_gdf .SignerInfoConfig {};_agb :=_gb .SHA256 .New ();_agb .Write (_gcc ._cdabd .Raw );var _aaa struct{Seq struct{Seq struct{Value []byte ;};};};_aaa .Seq .Seq .Value =_agb .Sum (nil );
var _dab []*_fe .Certificate ;var _dacd []*_fe .Certificate ;if _gcc ._ade !=nil {_dacd =[]*_fe .Certificate {_gcc ._ade };};_ceeb :=RevocationInfoArchival {Crl :[]_cd .RawValue {},Ocsp :[]_cd .RawValue {},OtherRevInfo :[]_cd .RawValue {}};_eff :=0;if _gcc ._cfd !=nil &&len (_gcc ._eag )> 0{_afd ,_gbae :=_gcc .makeTimestampRequest (_gcc ._eag ,([]byte )(""));
if _gbae !=nil {return _gbae ;};_fcf ,_gbae :=_fcc .Parse (_afd .FullBytes );if _gbae !=nil {return _gbae ;};_dab =append (_dab ,_fcf .Certificates ...);};if _gcc ._cfd !=nil {_dge ,_egc :=_gcc .addDss ([]*_fe .Certificate {_gcc ._cdabd },_dacd ,&_ceeb );
if _egc !=nil {return _egc ;};_eff +=_dge ;if len (_dab )> 0{_dge ,_egc =_gcc .addDss (_dab ,nil ,&_ceeb );if _egc !=nil {return _egc ;};_eff +=_dge ;};if !_gcc ._gg {_gcc ._cfd .SetDSS (_gcc ._eg );};};_bfb .ExtraSignedAttributes =append (_bfb .ExtraSignedAttributes ,_gdf .Attribute {Type :_gdf .OIDAttributeSigningCertificateV2 ,Value :_aaa },_gdf .Attribute {Type :_gdf .OIDAttributeAdobeRevocation ,Value :_ceeb });
if _caee :=_cee .AddSignerChainPAdES (_gcc ._cdabd ,_gcc ._fba ,_dacd ,_bfb );_caee !=nil {return _caee ;};_cee .Detach ();if len (_gcc ._eag )> 0{_badf :=_cee .GetSignedData ().SignerInfos [0].EncryptedDigest ;_fab ,_fabb :=_gcc .makeTimestampRequest (_gcc ._eag ,_badf );
if _fabb !=nil {return _fabb ;};_fabb =_cee .AddTimestampTokenToSigner (0,_fab .FullBytes );if _fabb !=nil {return _fabb ;};};_fdbf ,_bgc :=_cee .Finish ();if _bgc !=nil {return _bgc ;};_feb :=make ([]byte ,len (_fdbf )+1024*2+_eff );copy (_feb ,_fdbf );
sig .Contents =_dg .MakeHexString (string (_feb ));if !_gcc ._gg &&_gcc ._eg !=nil {_agb =_gb .SHA1 .New ();_agb .Write (_feb );_edc :=_g .ToUpper (_fc .EncodeToString (_agb .Sum (nil )));if _edc !=""{_gcc ._eg .VRI [_edc ]=&_dgd .VRI {Cert :_gcc ._eg .Certs ,OCSP :_gcc ._eg .OCSPs ,CRL :_gcc ._eg .CRLs };
};_gcc ._cfd .SetDSS (_gcc ._eg );};return nil ;};type adobeX509RSASHA1 struct{_adg *_de .PrivateKey ;_fae *_fe .Certificate ;_gbca SignFunc ;_aff bool ;_ddcb _gb .Hash ;};

// NewDigest creates a new digest.
func (_fdbb *adobeX509RSASHA1 )NewDigest (sig *_dgd .PdfSignature )(_dgd .Hasher ,error ){if _bdf ,_fee :=_fdbb .getHashAlgorithm (sig );_bdf !=0&&_fee ==nil {return _bdf .New (),nil ;};return _daf .New (),nil ;};type timestampInfo struct{Version int ;
Policy _cd .RawValue ;MessageImprint struct{HashAlgorithm _e .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _cd .RawValue ;GeneralizedTime _c .Time ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_dgd .SignatureHandler ,error ){return &adobePKCS7Detached {_fabf :true ,_cea :signatureLen },nil ;};

// Validate validates PdfSignature.
func (_fdf *adobeX509RSASHA1 )Validate (sig *_dgd .PdfSignature ,digest _dgd .Hasher )(_dgd .SignatureValidationResult ,error ){_gab ,_eagg :=_fdf .getCertificate (sig );if _eagg !=nil {return _dgd .SignatureValidationResult {},_eagg ;};_ffe :=sig .Contents .Bytes ();
var _aedg []byte ;if _ ,_bfcg :=_cd .Unmarshal (_ffe ,&_aedg );_bfcg !=nil {return _dgd .SignatureValidationResult {},_bfcg ;};_cca ,_gef :=digest .(_a .Hash );if !_gef {return _dgd .SignatureValidationResult {},_b .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_ceee ,_ :=_fdf .getHashAlgorithm (sig );if _ceee ==0{_ceee =_daf ;};if _ddg :=_de .VerifyPKCS1v15 (_gab .PublicKey .(*_de .PublicKey ),_ceee ,_cca .Sum (nil ),_aedg );_ddg !=nil {return _dgd .SignatureValidationResult {},_ddg ;};return _dgd .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};

// NewDigest creates a new digest.
func (_cdb *adobePKCS7Detached )NewDigest (sig *_dgd .PdfSignature )(_dgd .Hasher ,error ){return _ge .NewBuffer (nil ),nil ;};

// RevocationInfoArchival is OIDAttributeAdobeRevocation attribute.
type RevocationInfoArchival struct{Crl []_cd .RawValue `asn1:"explicit,tag:0,optional"`;Ocsp []_cd .RawValue `asn1:"explicit,tag:1,optional"`;OtherRevInfo []_cd .RawValue `asn1:"explicit,tag:2,optional"`;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gcfb *docTimeStamp )IsApplicable (sig *_dgd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};const _daf =_gb .SHA1 ;

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _gb .Hash ,opts *DocTimeStampOpts )(_dgd .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_dcag :timestampServerURL ,_ebe :hashAlgorithm ,_bfcb :opts .SignatureSize ,_cbg :opts .Client },nil ;
};

// InitSignature initialises the PdfSignature.
func (_dbd *docTimeStamp )InitSignature (sig *_dgd .PdfSignature )error {_fcfd :=*_dbd ;sig .Type =_dg .MakeName ("\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070");sig .Handler =&_fcfd ;sig .Filter =_dg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_dg .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");sig .Reference =nil ;if _dbd ._bfcb > 0{sig .Contents =_dg .MakeHexString (string (make ([]byte ,_dbd ._bfcb )));}else {_ccec ,_fadd :=_dbd .NewDigest (sig );
if _fadd !=nil {return _fadd ;};_ccec .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));if _fadd =_fcfd .Sign (sig ,_ccec );
_fadd !=nil {return _fadd ;};_dbd ._bfcb =_fcfd ._bfcb ;};return nil ;};func (_cbe *etsiPAdES )getCerts (_afe []*_fe .Certificate )([][]byte ,error ){_bg :=make ([][]byte ,0,len (_afe ));for _ ,_fbag :=range _afe {_bg =append (_bg ,_fbag .Raw );};return _bg ,nil ;
};

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct{_eb _dgd .SignatureHandler ;Permission _fa .DocMDPPermission ;};

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_cf *DocMDPHandler )Validate (sig *_dgd .PdfSignature ,digest _dgd .Hasher )(_dgd .SignatureValidationResult ,error ){return _dgd .SignatureValidationResult {},_b .New ("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065");
};

// InitSignature initialises the PdfSignature.
func (_aac *etsiPAdES )InitSignature (sig *_dgd .PdfSignature )error {if !_aac ._da {if _aac ._cdabd ==nil {return _b .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _aac ._fba ==nil {return _b .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_dce :=*_aac ;sig .Handler =&_dce ;sig .Filter =_dg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_dg .MakeName ("\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_eab ,_ebd :=_dce .NewDigest (sig );if _ebd !=nil {return _ebd ;};_ ,_ebd =_eab .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _ebd !=nil {return _ebd ;};_dce ._gg =true ;_ebd =_dce .Sign (sig ,_eab );_dce ._gg =false ;return _ebd ;};

// Sign adds a new reference to signature's references array.
func (_fb *DocMDPHandler )Sign (sig *_dgd .PdfSignature ,digest _dgd .Hasher )error {return _fb ._eb .Sign (sig ,digest );};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_cg *DocMDPHandler )IsApplicable (sig *_dgd .PdfSignature )bool {_fad :=false ;for _ ,_fef :=range sig .Reference .Elements (){if _ae ,_gbe :=_dg .GetDict (_fef );_gbe {if _dc ,_fd :=_dg .GetNameVal (_ae .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"));
_fd {if _dc !="\u0044\u006f\u0063\u004d\u0044\u0050"{return false ;};if _ce ,_dd :=_dg .GetDict (_ae .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"));_dd {_ ,_cge :=_dg .GetNumberAsInt64 (_ce .Get ("\u0050"));if _cge !=nil {return false ;
};_fad =true ;break ;};};};};return _fad &&_cg ._eb .IsApplicable (sig );};

// NewEtsiPAdESLevelLT creates a new Adobe.PPKLite ETSI.CAdES.detached Level LT signature handler.
func NewEtsiPAdESLevelLT (privateKey *_de .PrivateKey ,certificate *_fe .Certificate ,caCert *_fe .Certificate ,certificateTimestampServerURL string ,appender *_dgd .PdfAppender )(_dgd .SignatureHandler ,error ){_fefe :=appender .Reader .DSS ;if _fefe ==nil {_fefe =_dgd .NewDSS ();
};if _be :=_fefe .GenerateHashMaps ();_be !=nil {return nil ,_be ;};return &etsiPAdES {_cdabd :certificate ,_fba :privateKey ,_ade :caCert ,_eag :certificateTimestampServerURL ,CertClient :_fag .NewCertClient (),OCSPClient :_fag .NewOCSPClient (),CRLClient :_fag .NewCRLClient (),_cfd :appender ,_eg :_fefe },nil ;
};

// Sign sets the Contents fields.
func (_cegg *adobePKCS7Detached )Sign (sig *_dgd .PdfSignature ,digest _dgd .Hasher )error {if _cegg ._fabf {_cafe :=_cegg ._cea ;if _cafe <=0{_cafe =8192;};sig .Contents =_dg .MakeHexString (string (make ([]byte ,_cafe )));return nil ;};_fdg ,_deea :=digest .(*_ge .Buffer );
if !_deea {return _bb .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_gf ,_ebc :=_gdf .NewSignedData (_fdg .Bytes ());if _ebc !=nil {return _ebc ;};if _cdf :=_gf .AddSigner (_cegg ._cga ,_cegg ._dca ,_gdf .SignerInfoConfig {});
_cdf !=nil {return _cdf ;};_gf .Detach ();_ggf ,_ebc :=_gf .Finish ();if _ebc !=nil {return _ebc ;};_ccb :=make ([]byte ,8192);copy (_ccb ,_ggf );sig .Contents =_dg .MakeHexString (string (_ccb ));return nil ;};func (_acga *adobeX509RSASHA1 )getHashAlgorithm (_gbeg *_dgd .PdfSignature )(_gb .Hash ,error ){_bdce ,_bfe :=_acga .getCertificate (_gbeg );
if _bfe !=nil {if _acga ._ddcb !=0{return _acga ._ddcb ,nil ;};return _daf ,_bfe ;};if _gbeg .Contents !=nil {_afc :=_gbeg .Contents .Bytes ();var _aced []byte ;if _ ,_gfc :=_cd .Unmarshal (_afc ,&_aced );_gfc ==nil {_gcg :=_daed (_bdce .PublicKey .(*_de .PublicKey ),_aced );
if _gcg > 0{return _gcg ,nil ;};};};if _acga ._ddcb !=0{return _acga ._ddcb ,nil ;};return _daf ,nil ;};

// NewDigest creates a new digest.
func (_edce *etsiPAdES )NewDigest (_ *_dgd .PdfSignature )(_dgd .Hasher ,error ){return _ge .NewBuffer (nil ),nil ;};

// Validate validates PdfSignature.
func (_dbc *docTimeStamp )Validate (sig *_dgd .PdfSignature ,digest _dgd .Hasher )(_dgd .SignatureValidationResult ,error ){_afbc :=sig .Contents .Bytes ();_agg ,_fdbc :=_gdf .Parse (_afbc );if _fdbc !=nil {return _dgd .SignatureValidationResult {},_fdbc ;
};if _fdbc =_agg .Verify ();_fdbc !=nil {return _dgd .SignatureValidationResult {},_fdbc ;};var _bgg timestampInfo ;_ ,_fdbc =_cd .Unmarshal (_agg .Content ,&_bgg );if _fdbc !=nil {return _dgd .SignatureValidationResult {},_fdbc ;};_adea ,_fdbc :=_dfb (_bgg .MessageImprint .HashAlgorithm .Algorithm );
if _fdbc !=nil {return _dgd .SignatureValidationResult {},_fdbc ;};_gcd :=_adea .New ();_edab ,_aga :=digest .(*_ge .Buffer );if !_aga {return _dgd .SignatureValidationResult {},_bb .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");
};_gcd .Write (_edab .Bytes ());_fbfd :=_gcd .Sum (nil );_ggcc :=_dgd .SignatureValidationResult {IsSigned :true ,IsVerified :_ge .Equal (_fbfd ,_bgg .MessageImprint .HashedMessage ),GeneralizedTime :_bgg .GeneralizedTime };return _ggcc ,nil ;};func (_agc *etsiPAdES )addDss (_caaa ,_cfdf []*_fe .Certificate ,_agce *RevocationInfoArchival )(int ,error ){_acff ,_dda ,_ccf :=_agc .buildCertChain (_caaa ,_cfdf );
if _ccf !=nil {return 0,_ccf ;};_ceeg ,_ccf :=_agc .getCerts (_acff );if _ccf !=nil {return 0,_ccf ;};var _cfe ,_bc [][]byte ;if _agc .OCSPClient !=nil {_cfe ,_ccf =_agc .getOCSPs (_acff ,_dda );if _ccf !=nil {return 0,_ccf ;};};if _agc .CRLClient !=nil {_bc ,_ccf =_agc .getCRLs (_acff );
if _ccf !=nil {return 0,_ccf ;};};if !_agc ._gg {_ ,_ccf =_agc ._eg .AddCerts (_ceeg );if _ccf !=nil {return 0,_ccf ;};_ ,_ccf =_agc ._eg .AddOCSPs (_cfe );if _ccf !=nil {return 0,_ccf ;};_ ,_ccf =_agc ._eg .AddCRLs (_bc );if _ccf !=nil {return 0,_ccf ;
};};_ecea :=0;for _ ,_ace :=range _bc {_ecea +=len (_ace );_agce .Crl =append (_agce .Crl ,_cd .RawValue {FullBytes :_ace });};for _ ,_afbd :=range _cfe {_ecea +=len (_afbd );_agce .Ocsp =append (_agce .Ocsp ,_cd .RawValue {FullBytes :_afbd });};return _ecea ,nil ;
};func (_fccf *adobeX509RSASHA1 )sign (_egf *_dgd .PdfSignature ,_bed _dgd .Hasher ,_gefe bool )error {if !_gefe {return _fccf .Sign (_egf ,_bed );};_eda ,_ebcb :=_fccf ._fae .PublicKey .(*_de .PublicKey );if !_ebcb {return _bb .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_eda );
};_gae ,_cad :=_cd .Marshal (make ([]byte ,_eda .Size ()));if _cad !=nil {return _cad ;};_egf .Contents =_dg .MakeHexString (string (_gae ));return nil ;};

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler (handler _dgd .SignatureHandler ,permission _fa .DocMDPPermission )(_dgd .SignatureHandler ,error ){return &DocMDPHandler {_eb :handler ,Permission :permission },nil ;};func (_feg *etsiPAdES )buildCertChain (_gbc ,_dgg []*_fe .Certificate )([]*_fe .Certificate ,map[string ]*_fe .Certificate ,error ){_gbg :=map[string ]*_fe .Certificate {};
for _ ,_bd :=range _gbc {_gbg [_bd .Subject .CommonName ]=_bd ;};_eagb :=_gbc ;for _ ,_deeg :=range _dgg {_dff :=_deeg .Subject .CommonName ;if _ ,_acb :=_gbg [_dff ];_acb {continue ;};_gbg [_dff ]=_deeg ;_eagb =append (_eagb ,_deeg );};if len (_eagb )==0{return nil ,nil ,_dgd .ErrSignNoCertificates ;
};var _ccgf error ;for _bad :=_eagb [0];_bad !=nil &&!_feg .CertClient .IsCA (_bad );{var _bfc *_fe .Certificate ;_ ,_bfa :=_gbg [_bad .Issuer .CommonName ];if !_bfa {if _bfc ,_ccgf =_feg .CertClient .GetIssuer (_bad );_ccgf !=nil {_gd .Log .Debug ("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076",_ccgf );
break ;};_gbg [_bad .Issuer .CommonName ]=_bfc ;_eagb =append (_eagb ,_bfc );}else {break ;};_bad =_bfc ;};return _eagb ,_gbg ,nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_ede *adobeX509RSASHA1 )Sign (sig *_dgd .PdfSignature ,digest _dgd .Hasher )error {var _fea []byte ;var _fac error ;if _ede ._gbca !=nil {_fea ,_fac =_ede ._gbca (sig ,digest );if _fac !=nil {return _fac ;};}else {_fbg ,_ccc :=digest .(_a .Hash );
if !_ccc {return _b .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_ggce :=_daf ;if _ede ._ddcb !=0{_ggce =_ede ._ddcb ;};_fea ,_fac =_de .SignPKCS1v15 (_fg .Reader ,_ede ._adg ,_ggce ,_fbg .Sum (nil ));if _fac !=nil {return _fac ;
};};_fea ,_fac =_cd .Marshal (_fea );if _fac !=nil {return _fac ;};sig .Contents =_dg .MakeHexString (string (_fea ));return nil ;};

// InitSignature initialises the PdfSignature.
func (_bcc *adobeX509RSASHA1 )InitSignature (sig *_dgd .PdfSignature )error {if _bcc ._fae ==nil {return _b .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _bcc ._adg ==nil &&_bcc ._gbca ==nil {return _b .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_dae :=*_bcc ;sig .Handler =&_dae ;sig .Filter =_dg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_dg .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_dg .MakeString (string (_dae ._fae .Raw ));sig .Reference =nil ;_dde ,_ee :=_dae .NewDigest (sig );if _ee !=nil {return _ee ;};_dde .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _dae .sign (sig ,_dde ,_bcc ._aff );};

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct{

// EstimateSize specifies whether the size of the signature contents
// should be estimated based on the modulus size of the public key
// extracted from the signing certificate. If set to false, a mock Sign
// call is made in order to estimate the size of the signature contents.
EstimateSize bool ;

// Algorithm specifies the algorithm used for performing signing.
// If not specified, defaults to SHA1.
Algorithm _gb .Hash ;};func _dfb (_cff _cd .ObjectIdentifier )(_gb .Hash ,error ){switch {case _cff .Equal (_gdf .OIDDigestAlgorithmSHA1 ),_cff .Equal (_gdf .OIDDigestAlgorithmECDSASHA1 ),_cff .Equal (_gdf .OIDDigestAlgorithmDSA ),_cff .Equal (_gdf .OIDDigestAlgorithmDSASHA1 ),_cff .Equal (_gdf .OIDEncryptionAlgorithmRSA ):return _gb .SHA1 ,nil ;
case _cff .Equal (_gdf .OIDDigestAlgorithmSHA256 ),_cff .Equal (_gdf .OIDDigestAlgorithmECDSASHA256 ):return _gb .SHA256 ,nil ;case _cff .Equal (_gdf .OIDDigestAlgorithmSHA384 ),_cff .Equal (_gdf .OIDDigestAlgorithmECDSASHA384 ):return _gb .SHA384 ,nil ;
case _cff .Equal (_gdf .OIDDigestAlgorithmSHA512 ),_cff .Equal (_gdf .OIDDigestAlgorithmECDSASHA512 ):return _gb .SHA512 ,nil ;};return _gb .Hash (0),_gdf .ErrUnsupportedAlgorithm ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_de .PrivateKey ,certificate *_fe .Certificate )(_dgd .SignatureHandler ,error ){return &adobeX509RSASHA1 {_fae :certificate ,_adg :privateKey },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_bfag *etsiPAdES )IsApplicable (sig *_dgd .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// InitSignature initialises the PdfSignature.
func (_fdeb *adobePKCS7Detached )InitSignature (sig *_dgd .PdfSignature )error {if !_fdeb ._fabf {if _fdeb ._cga ==nil {return _b .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _fdeb ._dca ==nil {return _b .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_egd :=*_fdeb ;sig .Handler =&_egd ;sig .Filter =_dg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_dg .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_aace ,_badb :=_egd .NewDigest (sig );if _badb !=nil {return _badb ;};_aace .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _egd .Sign (sig ,_aace );};func (_bdc *adobeX509RSASHA1 )getCertificate (_gcbc *_dgd .PdfSignature )(*_fe .Certificate ,error ){if _bdc ._fae !=nil {return _bdc ._fae ,nil ;};_ffa ,_affc :=_gcbc .GetCerts ();if _affc !=nil {return nil ,_affc ;};
return _ffa [0],nil ;};func (_dba *docTimeStamp )getCertificate (_agbe *_dgd .PdfSignature )(*_fe .Certificate ,error ){_cfcdg ,_fgf :=_agbe .GetCerts ();if _fgf !=nil {return nil ,_fgf ;};return _cfcdg [0],nil ;};func (_aee *etsiPAdES )getOCSPs (_fda []*_fe .Certificate ,_aea map[string ]*_fe .Certificate )([][]byte ,error ){_ba :=make ([][]byte ,0,len (_fda ));
for _ ,_egb :=range _fda {for _ ,_caa :=range _egb .OCSPServer {if _aee .CertClient .IsCA (_egb ){continue ;};_ag ,_fbe :=_aea [_egb .Issuer .CommonName ];if !_fbe {_gd .Log .Debug ("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
continue ;};_ ,_agd ,_cb :=_aee .OCSPClient .MakeRequest (_caa ,_egb ,_ag );if _cb !=nil {_gd .Log .Debug ("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076",_cb );
continue ;};_ba =append (_ba ,_agd );};};return _ba ,nil ;};