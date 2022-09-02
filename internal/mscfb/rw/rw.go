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

package rw ;import (_d "bytes";_f "encoding/binary";_b "errors";_bb "fmt";_g "io";_cb "io/ioutil";_cg "reflect";);func (_aba *Writer )Bytes ()[]byte {return _aba ._ada };func (_ac *Reader )ReadStringProperty (n uint32 )(string ,error ){if _af :=_ac .align (4);_af !=nil {return "",_af ;};_ee :=make ([]byte ,n );if _gb :=_f .Read (_ac ,_f .LittleEndian ,&_ee );_gb !=nil {return "",_gb ;};return string (_ee ),nil ;};func (_df *Reader )ReadPairProperty (p interface{})error {if _dfd :=_df .align (4);_dfd !=nil {return _dfd ;};_ge :=_cg .ValueOf (p );for _ge .Kind ()==_cg .Ptr {_ge =_ge .Elem ();};if !_ge .IsValid (){return _bb .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");};if _ag :=_f .Read (_df ,_f .LittleEndian ,p );_ag !=nil {return _ag ;};return nil ;};func (_eea *Writer )WriteByteAt (b byte ,off int )error {if off >=len (_eea ._ada ){return _b .New ("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};_eea ._ada [off ]=b ;return nil ;};func (_da *Writer )curPos ()int {return int (_da .Cap ())-_da .Len ()};func (_cab *Writer )AlignLength (alignTo int )error {_ce :=_cab .Len ()%alignTo ;if _ce > 0{_ ,_ba :=_cab .Write (make ([]byte ,alignTo -_ce ));if _ba !=nil {return _ba ;};};return nil ;};var _bf =_b .New ("r\u0077.\u0057\u0072\u0069\u0074\u0065\u0072\u003a\u0020t\u006f\u006f\u0020\u006car\u0067\u0065");func PushLeftUI32 (v uint32 ,flag bool )uint32 {v >>=1;if flag {v |=1<<31;};return v ;};func PopRightUI64 (v uint64 )(bool ,uint64 ){return (v &uint64 (1))==1,v >>1};type Writer struct{_ada []byte ;_geg int ;};func (_eef *Writer )reset (){_eef ._ada =_eef ._ada [:0];_eef ._geg =0};func PushLeftUI64 (v uint64 ,flag bool )uint64 {v >>=1;if flag {v |=1<<63;};return v ;};func (_ae *Writer )WritePropertyNoAlign (a interface{})error {if _eg :=_f .Write (_ae ,_f .LittleEndian ,a );_eg !=nil {return _eg ;};return nil ;};const _bcg =int (^uint (0)>>1);func (_gc *Reader )ReadProperty (a interface{})error {_ca :=_cg .ValueOf (a );for _ca .Kind ()==_cg .Ptr {_ca =_ca .Elem ();};if !_ca .IsValid (){return _bb .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");};if _ad :=_gc .align (int (_ca .Type ().Size ()));_ad !=nil {return _ad ;};if _ed :=_f .Read (_gc ,_f .LittleEndian ,a );_ed !=nil {return _ed ;};return nil ;};func (_ab *Writer )WriteProperty (a interface{})error {if _ea :=_ab .align (int (_cg .TypeOf (a ).Size ()));_ea !=nil {return _ea ;};return _ab .WritePropertyNoAlign (a );};func (_dff *Writer )Len ()int {return len (_dff ._ada )-_dff ._geg };func NewWriter ()*Writer {return &Writer {_ada :[]byte {}}};func (_cfe *Writer )grow (_ddf int )(int ,error ){_bfe :=_cfe .Len ();if _bfe ==0&&_cfe ._geg !=0{_cfe .reset ();};if _cgf ,_bgd :=_cfe .tryGrowByReslice (_ddf );_bgd {return _cgf ,nil ;};if _cfe ._ada ==nil &&_ddf <=_gce {_cfe ._ada =make ([]byte ,_ddf ,_gce );return 0,nil ;};_eb :=cap (_cfe ._ada );if _ddf <=_eb /2-_bfe {copy (_cfe ._ada ,_cfe ._ada [_cfe ._geg :]);}else if _eb > _bcg -_eb -_ddf {return 0,_bf ;}else {_dda :=_db (2*_eb +_ddf );copy (_dda ,_cfe ._ada [_cfe ._geg :]);_cfe ._ada =_dda ;};_cfe ._geg =0;_cfe ._ada =_cfe ._ada [:_bfe +_ddf ];return _bfe ,nil ;};func (_a *Reader )skip (_bc int )error {_ ,_fg :=_g .CopyN (_cb .Discard ,_a ,int64 (_bc ));if _fg !=nil {return _fg ;};return nil ;};func (_aae *Writer )tryGrowByReslice (_dab int )(int ,bool ){if _aad :=len (_aae ._ada );_dab <=cap (_aae ._ada )-_aad {_aae ._ada =_aae ._ada [:_aad +_dab ];return _aad ,true ;};return 0,false ;};func (_fae *Writer )FillWithByte (fillSize int ,b byte )error {for _abc :=0;_abc < fillSize ;_abc ++{if _ff :=_fae .WritePropertyNoAlign (b );_ff !=nil {return _ff ;};};return nil ;};func (_e *Reader )curPos ()int {return int (_e .Size ())-_e .Len ()};func (_ga *Writer )Write (p []byte )(_ddc int ,_aa error ){_ega ,_dcc :=_ga .tryGrowByReslice (len (p ));if !_dcc {var _ceb error ;_ega ,_ceb =_ga .grow (len (p ));if _ceb !=nil {return 0,_ceb ;};};return copy (_ga ._ada [_ega :],p ),nil ;};type Reader struct{*_d .Reader };func (_dc *Reader )align (_fb int )error {return _dc .skip ((_fb -_dc .curPos ()%_fb )%_fb )};func (_eaf *Writer )WriteStringProperty (s string )error {_eaf .align (4);_bga :=[]byte (s );if _fbb :=_f .Write (_eaf ,_f .LittleEndian ,&_bga );_fbb !=nil {return _fbb ;};return nil ;};func PopRightUI32 (v uint32 )(bool ,uint32 ){return (v &uint32 (1))==1,v >>1};func (_aec *Writer )align (_afc int )error {return _aec .Skip ((_afc -(_aec .Len ())%_afc )%_afc )};func _db (_dba int )[]byte {defer func (){if recover ()!=nil {panic (_bf );};}();return make ([]byte ,_dba );};func (_gf *Writer )Cap ()int {return cap (_gf ._ada )};const _gce =64;func (_cad *Writer )Skip (n int )error {if n ==0{return nil ;};_ ,_dd :=_cad .Write (make ([]byte ,n ));return _dd ;};func NewReader (b []byte )(*Reader ,error ){return &Reader {_d .NewReader (b )},nil };func (_aag *Writer )WriteTo (wTo _g .Writer )(_cf int64 ,_ege error ){if _gg :=_aag .Len ();_gg > 0{_cd ,_daa :=wTo .Write (_aag ._ada [_aag ._geg :]);if _cd > _gg {return 0,_b .New ("\u0072\u0077\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f\u003a\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0057\u0072\u0069t\u0065\u0020c\u006f\u0075\u006e\u0074");};_aag ._geg +=_cd ;_cf =int64 (_cd );if _daa !=nil {return _cf ,_daa ;};if _cd !=_gg {return _cf ,_g .ErrShortWrite ;};};_aag .reset ();return _cf ,nil ;};