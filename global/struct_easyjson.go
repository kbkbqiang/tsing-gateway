// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package global

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal(in *jlexer.Lexer, out *UpstreamType) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "middleware":
			if in.IsNull() {
				in.Skip()
				out.Middleware = nil
			} else {
				in.Delim('[')
				if out.Middleware == nil {
					if !in.IsDelim(']') {
						out.Middleware = make([]ModuleConfig, 0, 2)
					} else {
						out.Middleware = []ModuleConfig{}
					}
				} else {
					out.Middleware = (out.Middleware)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ModuleConfig
					(v1).UnmarshalEasyJSON(in)
					out.Middleware = append(out.Middleware, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "static_endpoint":
			out.StaticEndpoint = string(in.String())
		case "discover":
			(out.Discover).UnmarshalEasyJSON(in)
		case "load_balance":
			out.LoadBalance = string(in.String())
		case "max_cache_fault":
			out.MaxCacheFault = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(out *jwriter.Writer, in UpstreamType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if len(in.Middleware) != 0 {
		const prefix string = ",\"middleware\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Middleware {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.StaticEndpoint != "" {
		const prefix string = ",\"static_endpoint\":"
		out.RawString(prefix)
		out.String(string(in.StaticEndpoint))
	}
	if true {
		const prefix string = ",\"discover\":"
		out.RawString(prefix)
		(in.Discover).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"load_balance\":"
		out.RawString(prefix)
		out.String(string(in.LoadBalance))
	}
	{
		const prefix string = ",\"max_cache_fault\":"
		out.RawString(prefix)
		out.Int(int(in.MaxCacheFault))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpstreamType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpstreamType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpstreamType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpstreamType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal(l, v)
}
func easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal1(in *jlexer.Lexer, out *ModuleConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "config":
			out.Config = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal1(out *jwriter.Writer, in ModuleConfig) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"config\":"
		out.RawString(prefix)
		out.String(string(in.Config))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ModuleConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ModuleConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ModuleConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ModuleConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal1(l, v)
}
func easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(in *jlexer.Lexer, out *EndpointType) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "upstream_id":
			out.UpstreamID = string(in.String())
		case "addr":
			out.Addr = string(in.String())
		case "weight":
			out.Weight = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(out *jwriter.Writer, in EndpointType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"upstream_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UpstreamID))
	}
	{
		const prefix string = ",\"addr\":"
		out.RawString(prefix)
		out.String(string(in.Addr))
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		out.Int(int(in.Weight))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EndpointType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EndpointType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EndpointType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EndpointType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(l, v)
}
