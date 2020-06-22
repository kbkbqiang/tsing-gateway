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

func easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal(in *jlexer.Lexer, out *ServiceType) {
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
func easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(out *jwriter.Writer, in ServiceType) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceType) UnmarshalEasyJSON(l *jlexer.Lexer) {
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
func easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(in *jlexer.Lexer, out *HostType) {
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
		case "route_group_id":
			out.RouteGroupID = string(in.String())
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
					var v4 ModuleConfig
					(v4).UnmarshalEasyJSON(in)
					out.Middleware = append(out.Middleware, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(out *jwriter.Writer, in HostType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"route_group_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.RouteGroupID))
	}
	if len(in.Middleware) != 0 {
		const prefix string = ",\"middleware\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Middleware {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HostType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HostType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComDxvgefTsingGatewayGlobal2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HostType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HostType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComDxvgefTsingGatewayGlobal2(l, v)
}
