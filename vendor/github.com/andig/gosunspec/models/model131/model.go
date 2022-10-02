package model131

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block131 - Watt-PF - Watt-Power Factor

const (
	ModelID          = 131
	ModelLabel       = "Watt-PF"
	ModelDescription = "Watt-Power Factor "
)

const (
	ActCrv       = "ActCrv"
	ActPt        = "ActPt"
	CrvNam       = "CrvNam"
	ModEna       = "ModEna"
	NCrv         = "NCrv"
	NPt          = "NPt"
	PF1          = "PF1"
	PF10         = "PF10"
	PF11         = "PF11"
	PF12         = "PF12"
	PF13         = "PF13"
	PF14         = "PF14"
	PF15         = "PF15"
	PF16         = "PF16"
	PF17         = "PF17"
	PF18         = "PF18"
	PF19         = "PF19"
	PF2          = "PF2"
	PF20         = "PF20"
	PF3          = "PF3"
	PF4          = "PF4"
	PF5          = "PF5"
	PF6          = "PF6"
	PF7          = "PF7"
	PF8          = "PF8"
	PF9          = "PF9"
	PF_SF        = "PF_SF"
	Pad          = "Pad"
	ReadOnly     = "ReadOnly"
	RmpDecTmm    = "RmpDecTmm"
	RmpIncDec_SF = "RmpIncDec_SF"
	RmpIncTmm    = "RmpIncTmm"
	RmpPT1Tms    = "RmpPT1Tms"
	RmpTms       = "RmpTms"
	RvrtTms      = "RvrtTms"
	W1           = "W1"
	W10          = "W10"
	W11          = "W11"
	W12          = "W12"
	W13          = "W13"
	W14          = "W14"
	W15          = "W15"
	W16          = "W16"
	W17          = "W17"
	W18          = "W18"
	W19          = "W19"
	W2           = "W2"
	W20          = "W20"
	W3           = "W3"
	W4           = "W4"
	W5           = "W5"
	W6           = "W6"
	W7           = "W7"
	W8           = "W8"
	W9           = "W9"
	W_SF         = "W_SF"
	WinTms       = "WinTms"
)

type Block131Repeat struct {
	ActPt     uint16         `sunspec:"offset=0,len=1,access=rw"`
	W1        int16          `sunspec:"offset=1,len=1,sf=W_SF,access=rw"`
	PF1       int16          `sunspec:"offset=2,len=1,sf=PF_SF,access=rw"`
	W2        int16          `sunspec:"offset=3,len=1,sf=W_SF,access=rw"`
	PF2       int16          `sunspec:"offset=4,len=1,sf=PF_SF,access=rw"`
	W3        int16          `sunspec:"offset=5,len=1,sf=W_SF,access=rw"`
	PF3       int16          `sunspec:"offset=6,len=1,sf=PF_SF,access=rw"`
	W4        int16          `sunspec:"offset=7,len=1,sf=W_SF,access=rw"`
	PF4       int16          `sunspec:"offset=8,len=1,sf=PF_SF,access=rw"`
	W5        int16          `sunspec:"offset=9,len=1,sf=W_SF,access=rw"`
	PF5       int16          `sunspec:"offset=10,len=1,sf=PF_SF,access=rw"`
	W6        int16          `sunspec:"offset=11,len=1,sf=W_SF,access=rw"`
	PF6       int16          `sunspec:"offset=12,len=1,sf=PF_SF,access=rw"`
	W7        int16          `sunspec:"offset=13,len=1,sf=W_SF,access=rw"`
	PF7       int16          `sunspec:"offset=14,len=1,sf=PF_SF,access=rw"`
	W8        int16          `sunspec:"offset=15,len=1,sf=W_SF,access=rw"`
	PF8       int16          `sunspec:"offset=16,len=1,sf=PF_SF,access=rw"`
	W9        int16          `sunspec:"offset=17,len=1,sf=W_SF,access=rw"`
	PF9       int16          `sunspec:"offset=18,len=1,sf=PF_SF,access=rw"`
	W10       int16          `sunspec:"offset=19,len=1,sf=W_SF,access=rw"`
	PF10      int16          `sunspec:"offset=20,len=1,sf=PF_SF,access=rw"`
	W11       int16          `sunspec:"offset=21,len=1,sf=W_SF,access=rw"`
	PF11      int16          `sunspec:"offset=22,len=1,sf=PF_SF,access=rw"`
	W12       int16          `sunspec:"offset=23,len=1,sf=W_SF,access=rw"`
	PF12      int16          `sunspec:"offset=24,len=1,sf=PF_SF,access=rw"`
	W13       int16          `sunspec:"offset=25,len=1,sf=W_SF,access=rw"`
	PF13      int16          `sunspec:"offset=26,len=1,sf=PF_SF,access=rw"`
	W14       int16          `sunspec:"offset=27,len=1,sf=W_SF,access=rw"`
	PF14      int16          `sunspec:"offset=28,len=1,sf=PF_SF,access=rw"`
	W15       int16          `sunspec:"offset=29,len=1,sf=W_SF,access=rw"`
	PF15      int16          `sunspec:"offset=30,len=1,sf=PF_SF,access=rw"`
	W16       int16          `sunspec:"offset=31,len=1,sf=W_SF,access=rw"`
	PF16      int16          `sunspec:"offset=32,len=1,sf=PF_SF,access=rw"`
	W17       int16          `sunspec:"offset=33,len=1,sf=W_SF,access=rw"`
	PF17      int16          `sunspec:"offset=34,len=1,sf=PF_SF,access=rw"`
	W18       int16          `sunspec:"offset=35,len=1,sf=W_SF,access=rw"`
	PF18      int16          `sunspec:"offset=36,len=1,sf=PF_SF,access=rw"`
	W19       int16          `sunspec:"offset=37,len=1,sf=W_SF,access=rw"`
	PF19      int16          `sunspec:"offset=38,len=1,sf=PF_SF,access=rw"`
	W20       int16          `sunspec:"offset=39,len=1,sf=W_SF,access=rw"`
	PF20      int16          `sunspec:"offset=40,len=1,sf=PF_SF,access=rw"`
	CrvNam    string         `sunspec:"offset=41,len=8,access=rw"`
	RmpPT1Tms uint16         `sunspec:"offset=49,len=1,access=rw"`
	RmpDecTmm uint16         `sunspec:"offset=50,len=1,sf=RmpIncDec_SF,access=rw"`
	RmpIncTmm uint16         `sunspec:"offset=51,len=1,sf=RmpIncDec_SF,access=rw"`
	ReadOnly  sunspec.Enum16 `sunspec:"offset=52,len=1,access=r"`
	Pad       sunspec.Pad    `sunspec:"offset=53,len=1,access=r"`
}

type Block131 struct {
	ActCrv       uint16              `sunspec:"offset=0,len=1,access=rw"`
	ModEna       sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms       uint16              `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms      uint16              `sunspec:"offset=3,len=1,access=rw"`
	RmpTms       uint16              `sunspec:"offset=4,len=1,access=rw"`
	NCrv         uint16              `sunspec:"offset=5,len=1,access=r"`
	NPt          uint16              `sunspec:"offset=6,len=1,access=r"`
	W_SF         sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	PF_SF        sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	RmpIncDec_SF sunspec.ScaleFactor `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block131Repeat
}

func (block *Block131) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "watt_pf",
		Length: 64,
		Blocks: []smdx.BlockElement{
			{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					{Id: ActCrv, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActCrv", Description: "Index of active curve. 0=no active curve."},
					{Id: ModEna, Offset: 1, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Is watt-PF mode active."},
					{Id: WinTms, Offset: 2, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for watt-PF change."},
					{Id: RvrtTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RvrtTms", Description: "Timeout period for watt-PF curve selection."},
					{Id: RmpTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current mode to new mode."},
					{Id: NCrv, Offset: 5, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NCrv", Description: "Number of curves supported (recommend 4)."},
					{Id: NPt, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NPt", Description: "Max number of points in array."},
					{Id: W_SF, Offset: 7, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "W_SF", Description: "Scale factor for percent WMax."},
					{Id: PF_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "PF_SF", Description: "Scale factor for PF."},
					{Id: RmpIncDec_SF, Offset: 9, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Label: "RmpIncDec_SF", Description: "Scale factor for increment and decrement ramps."},
				},
			},
			{Name: "curve",
				Length: 54,
				Type:   "repeating",
				Points: []smdx.PointElement{
					{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActPt", Description: "Number of active points in array."},
					{Id: W1, Offset: 1, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Mandatory: true, Label: "W1", Description: "Point 1 Watts."},
					{Id: PF1, Offset: 2, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Mandatory: true, Label: "PF1", Description: "Point 1 PF in EEI notation."},
					{Id: W2, Offset: 3, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W2", Description: "Point 2 Watts."},
					{Id: PF2, Offset: 4, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF2", Description: "Point 2 PF in EEI notation."},
					{Id: W3, Offset: 5, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W3", Description: "Point 3 Watts."},
					{Id: PF3, Offset: 6, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF3", Description: "Point 3 PF in EEI notation."},
					{Id: W4, Offset: 7, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W4", Description: "Point 4 Watts."},
					{Id: PF4, Offset: 8, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF4", Description: "Point 4 PF in EEI notation."},
					{Id: W5, Offset: 9, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W5", Description: "Point 5 Watts."},
					{Id: PF5, Offset: 10, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF5", Description: "Point 5 PF in EEI notation."},
					{Id: W6, Offset: 11, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W6", Description: "Point 6 Watts."},
					{Id: PF6, Offset: 12, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF6", Description: "Point 6 PF in EEI notation."},
					{Id: W7, Offset: 13, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W7", Description: "Point 7 Watts."},
					{Id: PF7, Offset: 14, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF7", Description: "Point 7 PF in EEI notation."},
					{Id: W8, Offset: 15, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W8", Description: "Point 8 Watts."},
					{Id: PF8, Offset: 16, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF8", Description: "Point 8 PF in EEI notation."},
					{Id: W9, Offset: 17, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W9", Description: "Point 9 Watts."},
					{Id: PF9, Offset: 18, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF9", Description: "Point 9 PF in EEI notation."},
					{Id: W10, Offset: 19, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W10", Description: "Point 10 Watts."},
					{Id: PF10, Offset: 20, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF10", Description: "Point 10 PF in EEI notation."},
					{Id: W11, Offset: 21, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W11", Description: "Point 11 Watts."},
					{Id: PF11, Offset: 22, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF11", Description: "Point 11 PF in EEI notation."},
					{Id: W12, Offset: 23, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W12", Description: "Point 12 Watts."},
					{Id: PF12, Offset: 24, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF12", Description: "Point 12 PF in EEI notation."},
					{Id: W13, Offset: 25, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W13", Description: "Point 13 Watts."},
					{Id: PF13, Offset: 26, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF13", Description: "Point 13 PF in EEI notation."},
					{Id: W14, Offset: 27, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W14", Description: "Point 14 Watts."},
					{Id: PF14, Offset: 28, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF14", Description: "Point 14 PF in EEI notation."},
					{Id: W15, Offset: 29, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W15", Description: "Point 15 Watts."},
					{Id: PF15, Offset: 30, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF15", Description: "Point 15 PF in EEI notation."},
					{Id: W16, Offset: 31, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W16", Description: "Point 16 Watts."},
					{Id: PF16, Offset: 32, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF16", Description: "Point 16 PF in EEI notation."},
					{Id: W17, Offset: 33, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W17", Description: "Point 17 Watts."},
					{Id: PF17, Offset: 34, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF17", Description: "Point 17 PF in EEI notation."},
					{Id: W18, Offset: 35, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W18", Description: "Point 18 Watts."},
					{Id: PF18, Offset: 36, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF18", Description: "Point 18 PF in EEI notation."},
					{Id: W19, Offset: 37, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W19", Description: "Point 19 Watts."},
					{Id: PF19, Offset: 38, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF19", Description: "Point 19 PF in EEI notation."},
					{Id: W20, Offset: 39, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "% WMax", Access: "rw", Length: 1, Label: "W20", Description: "Point 20 Watts."},
					{Id: PF20, Offset: 40, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "cos()", Access: "rw", Length: 1, Label: "PF20", Description: "Point 20 PF in EEI notation."},
					{Id: CrvNam, Offset: 41, Type: typelabel.String, Access: "rw", Length: 8, Label: "CrvNam", Description: "Optional description for curve."},
					{Id: RmpPT1Tms, Offset: 49, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpPT1Tms", Description: "The time of the PT1 in seconds (time to accomplish a change of 95%)."},
					{Id: RmpDecTmm, Offset: 50, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% PF/min", Access: "rw", Length: 1, Label: "RmpDecTmm", Description: "The maximum rate at which the power factor may be reduced in response to changes in the power value."},
					{Id: RmpIncTmm, Offset: 51, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% PF/min", Access: "rw", Length: 1, Label: "RmpIncTmm", Description: "The maximum rate at which the power factor may be increased in response to changes in the power value."},
					{Id: ReadOnly, Offset: 52, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "ReadOnly", Description: "Enumerated value indicates if curve is read-only or can be modified."},
					{Id: Pad, Offset: 53, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
