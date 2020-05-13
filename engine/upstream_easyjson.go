// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package engine

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

func easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine(in *jlexer.Lexer, out *Upstream) {
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
		case "ID":
			out.ID = string(in.String())
		case "middleware":
			if in.IsNull() {
				in.Skip()
				out.Middleware = nil
			} else {
				in.Delim('[')
				if out.Middleware == nil {
					if !in.IsDelim(']') {
						out.Middleware = make([]Configurator, 0, 2)
					} else {
						out.Middleware = []Configurator{}
					}
				} else {
					out.Middleware = (out.Middleware)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Configurator
					(v1).UnmarshalEasyJSON(in)
					out.Middleware = append(out.Middleware, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "explorer":
			(out.Explorer).UnmarshalEasyJSON(in)
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
func easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine(out *jwriter.Writer, in Upstream) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
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
	{
		const prefix string = ",\"explorer\":"
		out.RawString(prefix)
		(in.Explorer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Upstream) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Upstream) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Upstream) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Upstream) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine(l, v)
}
func easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine1(in *jlexer.Lexer, out *Endpoint) {
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
		case "addr":
			out.Addr = string(in.String())
		case "weight":
			out.Weight = int(in.Int())
		case "ttl":
			out.TTL = int(in.Int())
		case "updated_at":
			out.UpdatedAt = int64(in.Int64())
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
func easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine1(out *jwriter.Writer, in Endpoint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addr\":"
		out.RawString(prefix[1:])
		out.String(string(in.Addr))
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		out.Int(int(in.Weight))
	}
	{
		const prefix string = ",\"ttl\":"
		out.RawString(prefix)
		out.Int(int(in.TTL))
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Int64(int64(in.UpdatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Endpoint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Endpoint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Endpoint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Endpoint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine1(l, v)
}
func easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine2(in *jlexer.Lexer, out *Configurator) {
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
func easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine2(out *jwriter.Writer, in Configurator) {
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
func (v Configurator) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Configurator) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson370d5093EncodeGithubComDxvgefTsingGatewayEngine2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Configurator) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Configurator) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson370d5093DecodeGithubComDxvgefTsingGatewayEngine2(l, v)
}
