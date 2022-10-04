package model19

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block19 - PPP Link - Include this model to configure a Point-to-Point Protocol link

const (
	ModelID          = 19
	ModelLabel       = "PPP Link"
	ModelDescription = "Include this model to configure a Point-to-Point Protocol link"
)

const (
	Auth   = "Auth"
	Bits   = "Bits"
	Dup    = "Dup"
	Flw    = "Flw"
	Nam    = "Nam"
	Pad    = "Pad"
	Pty    = "Pty"
	Pw     = "Pw"
	Rte    = "Rte"
	UsrNam = "UsrNam"
)

type Block19 struct {
	Nam    string         `sunspec:"offset=0,len=4,access=rw"`
	Rte    uint32         `sunspec:"offset=4,access=rw"`
	Bits   uint16         `sunspec:"offset=6,access=rw"`
	Pty    sunspec.Enum16 `sunspec:"offset=7,access=rw"`
	Dup    sunspec.Enum16 `sunspec:"offset=8,access=rw"`
	Flw    sunspec.Enum16 `sunspec:"offset=9,access=rw"`
	Auth   sunspec.Enum16 `sunspec:"offset=10"`
	UsrNam string         `sunspec:"offset=11,len=12"`
	Pw     string         `sunspec:"offset=23,len=6"`
	Pad    sunspec.Pad    `sunspec:"offset=29"`
}

func (block *Block19) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 30,
		Blocks: []smdx.BlockElement{
			{
				Length: 30,
				Points: []smdx.PointElement{
					{Id: Nam, Offset: 0, Type: typelabel.String, Access: "rw", Length: 4, Label: "Name", Description: "Interface name"},
					{Id: Rte, Offset: 4, Type: typelabel.Uint32, Units: "bps", Access: "rw", Mandatory: true, Label: "Rate", Description: "Interface baud rate in bits per second"},
					{Id: Bits, Offset: 6, Type: typelabel.Uint16, Access: "rw", Mandatory: true, Label: "Bits", Description: "Number of data bits per character"},
					{Id: Pty, Offset: 7, Type: typelabel.Enum16, Access: "rw", Mandatory: true, Label: "Parity", Description: "Bitmask value.  Parity setting"},
					{Id: Dup, Offset: 8, Type: typelabel.Enum16, Access: "rw", Label: "Duplex", Description: "Enumerated value.  Duplex mode"},
					{Id: Flw, Offset: 9, Type: typelabel.Enum16, Access: "rw", Label: "Flow Control", Description: "Flow Control Method"},
					{Id: Auth, Offset: 10, Type: typelabel.Enum16, Label: "Authentication", Description: "Enumerated value.  Authentication method"},
					{Id: UsrNam, Offset: 11, Type: typelabel.String, Length: 12, Label: "Username", Description: "Username for authentication"},
					{Id: Pw, Offset: 23, Type: typelabel.String, Length: 6, Label: "Password", Description: "Password for authentication"},
					{Id: Pad, Offset: 29, Type: typelabel.Pad},
				},
			},
		}})
}