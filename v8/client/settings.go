package client

import (
	"encoding/json"
	"fmt"
	"log"
)

// Settings holds optional client settings.
type Settings struct {
	disablePAFXFast         bool
	assumePreAuthentication bool
	preAuthEType            int32
	socksAddr               string
	logger                  *log.Logger
}

// jsonSettings is used when marshaling the Settings details to JSON format.
type jsonSettings struct {
	DisablePAFXFast         bool
	AssumePreAuthentication bool
}

// NewSettings creates a new client settings struct.
func NewSettings(settings ...func(*Settings)) *Settings {
	s := new(Settings)
	for _, set := range settings {
		set(s)
	}
	return s
}

// DisablePAFXFAST used to configure the client to not use PA_FX_FAST.
//
// s := NewSettings(DisablePAFXFAST(true))
func DisablePAFXFAST(b bool) func(*Settings) {
	return func(s *Settings) {
		s.disablePAFXFast = b
	}
}

// DisablePAFXFAST indicates is the client should disable the use of PA_FX_FAST.
func (s *Settings) DisablePAFXFAST() bool {
	return s.disablePAFXFast
}

// AssumePreAuthentication used to configure the client to assume pre-authentication is required.
//
// s := NewSettings(AssumePreAuthentication(true))
func AssumePreAuthentication(b bool) func(*Settings) {
	return func(s *Settings) {
		s.assumePreAuthentication = b
	}
}

// AssumePreAuthentication indicates if the client should proactively assume using pre-authentication.
func (s *Settings) AssumePreAuthentication() bool {
	return s.assumePreAuthentication
}

// PreAuthEType used to configure the preauthentication encryption type.
//
// s := NewSettings(PreAuthEType(true))
func PreAuthEType(e int32) func(*Settings) {
	return func(s *Settings) {
		s.preAuthEType = e
	}
}

// PreAuthEType indicates the preautneitcation encryption type.
func (s *Settings) PreAuthEType() int32 {
	return s.preAuthEType
}

// SocksAddr used to configure the client to use SOCKS5 proxy.
//
// s := NewSettings(SocksAddr("127.0.0.1:1080"))
func SocksAddr(a string) func(*Settings) {
	return func(s *Settings) {
		s.socksAddr = a
	}
}

// SocksAddr indicates is the client should use SOCKS5 proxy.
func (s *Settings) SocksAddr() string {
	return s.socksAddr
}

// Logger used to configure client with a logger.
//
// s := NewSettings(kt, Logger(l))
func Logger(l *log.Logger) func(*Settings) {
	return func(s *Settings) {
		s.logger = l
	}
}

// Logger returns the client logger instance.
func (s *Settings) Logger() *log.Logger {
	return s.logger
}

// Log will write to the service's logger if it is configured.
func (cl *Client) Log(format string, v ...interface{}) {
	if cl.settings.Logger() != nil {
		cl.settings.Logger().Output(2, fmt.Sprintf(format, v...))
	}
}

// JSON returns a JSON representation of the settings.
func (s *Settings) JSON() (string, error) {
	js := jsonSettings{
		DisablePAFXFast:         s.disablePAFXFast,
		AssumePreAuthentication: s.assumePreAuthentication,
	}
	b, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil

}
