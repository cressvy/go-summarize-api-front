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

// Package common contains common properties used by the subpackages.
package common ;import (_g "fmt";_gc "io";_f "os";_df "path/filepath";_a "runtime";_c "time";);

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};func _egd (_aga _gc .Writer ,_gff string ,_dga string ,_fg ...interface{}){_ ,_ffc ,_cb ,_bbe :=_a .Caller (3);if !_bbe {_ffc ="\u003f\u003f\u003f";_cb =0;}else {_ffc =_df .Base (_ffc );
};_ca :=_g .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_gff ,_ffc ,_cb )+_dga +"\u000a";_g .Fprintf (_aga ,_ca ,_fg ...);};

// LogLevel is the verbosity level for logging.
type LogLevel int ;func (_cdg ConsoleLogger )output (_bee _gc .Writer ,_afa string ,_edd string ,_gf ...interface{}){_egd (_bee ,_afa ,_edd ,_gf ...);};var ReleasedAt =_c .Date (_aac ,_egc ,_eaa ,_fa ,_gee ,0,0,_c .UTC );

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};const _egc =4;var Log Logger =DummyLogger {};const _fa =15;

// Trace logs trace message.
func (_fb ConsoleLogger )Trace (format string ,args ...interface{}){if _fb .LogLevel >=LogLevelTrace {_ba :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fb .output (_f .Stdout ,_ba ,format ,args ...);};};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };const _dgc ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";

// Notice logs notice message.
func (_gcg ConsoleLogger )Notice (format string ,args ...interface{}){if _gcg .LogLevel >=LogLevelNotice {_gd :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_gcg .output (_f .Stdout ,_gd ,format ,args ...);};};

// Notice logs notice message.
func (_fbf WriterLogger )Notice (format string ,args ...interface{}){if _fbf .LogLevel >=LogLevelNotice {_ab :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fbf .logToWriter (_fbf .Output ,_ab ,format ,args ...);};};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _gc .Writer )*WriterLogger {_bec :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_bec ;};

// Debug logs debug message.
func (_db WriterLogger )Debug (format string ,args ...interface{}){if _db .LogLevel >=LogLevelDebug {_ff :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_db .logToWriter (_db .Output ,_ff ,format ,args ...);};};

// Error logs error message.
func (_dc ConsoleLogger )Error (format string ,args ...interface{}){if _dc .LogLevel >=LogLevelError {_ddd :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_dc .output (_f .Stdout ,_ddd ,format ,args ...);};};

// Debug logs debug message.
func (_cdf ConsoleLogger )Debug (format string ,args ...interface{}){if _cdf .LogLevel >=LogLevelDebug {_ed :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_cdf .output (_f .Stdout ,_ed ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};const Version ="\u0033\u002e\u0035\u0038\u002e\u0030";const _aac =2024;

// DummyLogger does nothing.
type DummyLogger struct{};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_de string ,_cg ...interface{});Warning (_b string ,_gg ...interface{});Notice (_e string ,_dd ...interface{});Info (_cd string ,_cda ...interface{});Debug (_eg string ,_be ...interface{});Trace (_ge string ,_ddc ...interface{});
IsLogLevel (_cgd LogLevel )bool ;};

// Warning logs warning message.
func (_gb WriterLogger )Warning (format string ,args ...interface{}){if _gb .LogLevel >=LogLevelWarning {_deb :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_gb .logToWriter (_gb .Output ,_deb ,format ,args ...);};};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _c .Time )string {return t .Format (_dgc )+"\u0020\u0055\u0054\u0043"};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Trace logs trace message.
func (_dce WriterLogger )Trace (format string ,args ...interface{}){if _dce .LogLevel >=LogLevelTrace {_eag :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_dce .logToWriter (_dce .Output ,_eag ,format ,args ...);};};const (LogLevelTrace LogLevel =5;
LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Info logs info message.
func (_fbb WriterLogger )Info (format string ,args ...interface{}){if _fbb .LogLevel >=LogLevelInfo {_dg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_fbb .logToWriter (_fbb .Output ,_dg ,format ,args ...);};};

// Info logs info message.
func (_ef ConsoleLogger )Info (format string ,args ...interface{}){if _ef .LogLevel >=LogLevelInfo {_af :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ef .output (_f .Stdout ,_af ,format ,args ...);};};const _eaa =30;

// Warning logs warning message.
func (_ea ConsoleLogger )Warning (format string ,args ...interface{}){if _ea .LogLevel >=LogLevelWarning {_ec :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ea .output (_f .Stdout ,_ec ,format ,args ...);};};const _gee =30;

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _gc .Writer ;};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fd WriterLogger )IsLogLevel (level LogLevel )bool {return _fd .LogLevel >=level };

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};func (_eac WriterLogger )logToWriter (_abb _gc .Writer ,_fc string ,_egg string ,_bb ...interface{}){_egd (_abb ,_fc ,_egg ,_bb );};

// Error logs error message.
func (_ee WriterLogger )Error (format string ,args ...interface{}){if _ee .LogLevel >=LogLevelError {_ag :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ee .logToWriter (_ee .Output ,_ag ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gce ConsoleLogger )IsLogLevel (level LogLevel )bool {return _gce .LogLevel >=level };