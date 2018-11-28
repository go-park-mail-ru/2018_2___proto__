// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson4e7a3d2cDecodeProtoGameServerModels(in *jlexer.Lexer, out *ScoreRecords) {
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
		case "count":
			out.Count = int(in.Int())
		case "offset":
			out.Offset = int(in.Int())
		case "records":
			if in.IsNull() {
				in.Skip()
				out.Records = nil
			} else {
				in.Delim('[')
				if out.Records == nil {
					if !in.IsDelim(']') {
						out.Records = make([]*ScoreRecord, 0, 8)
					} else {
						out.Records = []*ScoreRecord{}
					}
				} else {
					out.Records = (out.Records)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ScoreRecord
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ScoreRecord)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Records = append(out.Records, v1)
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
func easyjson4e7a3d2cEncodeProtoGameServerModels(out *jwriter.Writer, in ScoreRecords) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	{
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Offset))
	}
	{
		const prefix string = ",\"records\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Records == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Records {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScoreRecords) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4e7a3d2cEncodeProtoGameServerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScoreRecords) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4e7a3d2cEncodeProtoGameServerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScoreRecords) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4e7a3d2cDecodeProtoGameServerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScoreRecords) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4e7a3d2cDecodeProtoGameServerModels(l, v)
}
