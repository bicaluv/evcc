package model10

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block10 - Communication Interface Header - To be included first for a complete interface description

const (
	ModelID          = 10
	ModelLabel       = "Communication Interface Header"
	ModelDescription = "To be included first for a complete interface description"
)

const (
	Ctl = "Ctl"
	Pad = "Pad"
	St  = "St"
	Typ = "Typ"
)

type Block10 struct {
	St  sunspec.Enum16 `sunspec:"offset=0"`
	Ctl uint16         `sunspec:"offset=1,access=rw"`
	Typ sunspec.Enum16 `sunspec:"offset=2"`
	Pad sunspec.Pad    `sunspec:"offset=3"`
}

func (block *Block10) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 4,
		Blocks: []smdx.BlockElement{
			{
				Length: 4,
				Points: []smdx.PointElement{
					{Id: St, Offset: 0, Type: typelabel.Enum16, Mandatory: true, Label: "Interface Status", Description: "Overall interface status"},
					{Id: Ctl, Offset: 1, Type: typelabel.Uint16, Access: "rw", Label: "Interface Control", Description: "Overall interface control (TBD)"},
					{Id: Typ, Offset: 2, Type: typelabel.Enum16, Label: "Physical Access Type", Description: "Enumerated value.  Type of physical media"},
					{Id: Pad, Offset: 3, Type: typelabel.Pad},
				},
			},
		}})
}
