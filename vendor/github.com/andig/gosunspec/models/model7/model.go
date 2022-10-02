package model7

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block7 - Secure Write Response Model (DRAFT 1) - Include a digital signature over the response

const (
	ModelID          = 7
	ModelLabel       = "Secure Write Response Model (DRAFT 1)"
	ModelDescription = "Include a digital signature over the response"
)

const (
	Alg   = "Alg"
	Alm   = "Alm"
	DS    = "DS"
	Ms    = "Ms"
	N     = "N"
	RqSeq = "RqSeq"
	Rsrvd = "Rsrvd"
	Seq   = "Seq"
	Sts   = "Sts"
	Ts    = "Ts"
)

type Block7Repeat struct {
	DS uint16 `sunspec:"offset=0,access=rw"`
}

type Block7 struct {
	RqSeq uint16         `sunspec:"offset=0,access=r"`
	Sts   sunspec.Enum16 `sunspec:"offset=1,access=r"`
	Ts    uint32         `sunspec:"offset=2,access=r"`
	Ms    uint16         `sunspec:"offset=4,access=r"`
	Seq   uint16         `sunspec:"offset=5,access=r"`
	Alm   sunspec.Enum16 `sunspec:"offset=6"`
	Rsrvd sunspec.Pad    `sunspec:"offset=7,access=r"`
	Alg   sunspec.Enum16 `sunspec:"offset=8,access=r"`
	N     uint16         `sunspec:"offset=9,access=rw"`

	Repeats []Block7Repeat
}

func (block *Block7) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 11,
		Blocks: []smdx.BlockElement{
			{
				Length: 10,
				Points: []smdx.PointElement{
					{Id: RqSeq, Offset: 0, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "Request Sequence", Description: "Sequence number from the request"},
					{Id: Sts, Offset: 1, Type: typelabel.Enum16, Access: "r", Mandatory: true, Label: "Status", Description: "Status of last write operation"},
					{Id: Ts, Offset: 2, Type: typelabel.Uint32, Access: "r", Mandatory: true, Label: "Timestamp", Description: "Timestamp value is the number of seconds since January 1, 2000"},
					{Id: Ms, Offset: 4, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "Milliseconds", Description: "Millisecond counter 0-999"},
					{Id: Seq, Offset: 5, Type: typelabel.Uint16, Access: "r", Mandatory: true, Label: "Sequence", Description: "Sequence number of response"},
					{Id: Alm, Offset: 6, Type: typelabel.Enum16, Mandatory: true, Label: "Alarm", Description: "Bitmask alarm code"},
					{Id: Rsrvd, Offset: 7, Type: typelabel.Pad, Access: "r", Mandatory: true},
					{Id: Alg, Offset: 8, Type: typelabel.Enum16, Access: "r", Mandatory: true, Label: "Algorithm", Description: "Algorithm used to compute the digital signature"},
					{Id: N, Offset: 9, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "N", Description: "Number of registers comprising the digital signature."},
				},
			},
			{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: DS, Offset: 0, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "DS", Description: "Digital Signature"},
				},
			},
		}})
}
