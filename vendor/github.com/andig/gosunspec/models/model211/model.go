package model211

// Code generated by ../../generators/models.go. DO NOT EDIT.
// You can regenerate it by running 'go generate ./models' from the directory above.

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block211 - single phase (AN or AB) meter -

const (
	ModelID          = 211
	ModelLabel       = "single phase (AN or AB) meter"
	ModelDescription = ""
)

const (
	A               = "A"
	AphA            = "AphA"
	AphB            = "AphB"
	AphC            = "AphC"
	Evt             = "Evt"
	Hz              = "Hz"
	PF              = "PF"
	PFphA           = "PFphA"
	PFphB           = "PFphB"
	PFphC           = "PFphC"
	PPV             = "PPV"
	PPVphAB         = "PPVphAB"
	PPVphBC         = "PPVphBC"
	PPVphCA         = "PPVphCA"
	PhV             = "PhV"
	PhVphA          = "PhVphA"
	PhVphB          = "PhVphB"
	PhVphC          = "PhVphC"
	TotVAhExp       = "TotVAhExp"
	TotVAhExpPhA    = "TotVAhExpPhA"
	TotVAhExpPhB    = "TotVAhExpPhB"
	TotVAhExpPhC    = "TotVAhExpPhC"
	TotVAhImp       = "TotVAhImp"
	TotVAhImpPhA    = "TotVAhImpPhA"
	TotVAhImpPhB    = "TotVAhImpPhB"
	TotVAhImpPhC    = "TotVAhImpPhC"
	TotVArhExpQ3    = "TotVArhExpQ3"
	TotVArhExpQ3phA = "TotVArhExpQ3phA"
	TotVArhExpQ3phB = "TotVArhExpQ3phB"
	TotVArhExpQ3phC = "TotVArhExpQ3phC"
	TotVArhExpQ4    = "TotVArhExpQ4"
	TotVArhExpQ4phA = "TotVArhExpQ4phA"
	TotVArhExpQ4phB = "TotVArhExpQ4phB"
	TotVArhExpQ4phC = "TotVArhExpQ4phC"
	TotVArhImpQ1    = "TotVArhImpQ1"
	TotVArhImpQ1phA = "TotVArhImpQ1phA"
	TotVArhImpQ1phB = "TotVArhImpQ1phB"
	TotVArhImpQ1phC = "TotVArhImpQ1phC"
	TotVArhImpQ2    = "TotVArhImpQ2"
	TotVArhImpQ2phA = "TotVArhImpQ2phA"
	TotVArhImpQ2phB = "TotVArhImpQ2phB"
	TotVArhImpQ2phC = "TotVArhImpQ2phC"
	TotWhExp        = "TotWhExp"
	TotWhExpPhA     = "TotWhExpPhA"
	TotWhExpPhB     = "TotWhExpPhB"
	TotWhExpPhC     = "TotWhExpPhC"
	TotWhImp        = "TotWhImp"
	TotWhImpPhA     = "TotWhImpPhA"
	TotWhImpPhB     = "TotWhImpPhB"
	TotWhImpPhC     = "TotWhImpPhC"
	VA              = "VA"
	VAR             = "VAR"
	VARphA          = "VARphA"
	VARphB          = "VARphB"
	VARphC          = "VARphC"
	VAphA           = "VAphA"
	VAphB           = "VAphB"
	VAphC           = "VAphC"
	W               = "W"
	WphA            = "WphA"
	WphB            = "WphB"
	WphC            = "WphC"
)

type Block211 struct {
	A               float32            `sunspec:"offset=0"`
	AphA            float32            `sunspec:"offset=2"`
	AphB            float32            `sunspec:"offset=4"`
	AphC            float32            `sunspec:"offset=6"`
	PhV             float32            `sunspec:"offset=8"`
	PhVphA          float32            `sunspec:"offset=10"`
	PhVphB          float32            `sunspec:"offset=12"`
	PhVphC          float32            `sunspec:"offset=14"`
	PPV             float32            `sunspec:"offset=16"`
	PPVphAB         float32            `sunspec:"offset=18"`
	PPVphBC         float32            `sunspec:"offset=20"`
	PPVphCA         float32            `sunspec:"offset=22"`
	Hz              float32            `sunspec:"offset=24"`
	W               float32            `sunspec:"offset=26"`
	WphA            float32            `sunspec:"offset=28"`
	WphB            float32            `sunspec:"offset=30"`
	WphC            float32            `sunspec:"offset=32"`
	VA              float32            `sunspec:"offset=34"`
	VAphA           float32            `sunspec:"offset=36"`
	VAphB           float32            `sunspec:"offset=38"`
	VAphC           float32            `sunspec:"offset=40"`
	VAR             float32            `sunspec:"offset=42"`
	VARphA          float32            `sunspec:"offset=44"`
	VARphB          float32            `sunspec:"offset=46"`
	VARphC          float32            `sunspec:"offset=48"`
	PF              float32            `sunspec:"offset=50"`
	PFphA           float32            `sunspec:"offset=52"`
	PFphB           float32            `sunspec:"offset=54"`
	PFphC           float32            `sunspec:"offset=56"`
	TotWhExp        float32            `sunspec:"offset=58"`
	TotWhExpPhA     float32            `sunspec:"offset=60"`
	TotWhExpPhB     float32            `sunspec:"offset=62"`
	TotWhExpPhC     float32            `sunspec:"offset=64"`
	TotWhImp        float32            `sunspec:"offset=66"`
	TotWhImpPhA     float32            `sunspec:"offset=68"`
	TotWhImpPhB     float32            `sunspec:"offset=70"`
	TotWhImpPhC     float32            `sunspec:"offset=72"`
	TotVAhExp       float32            `sunspec:"offset=74"`
	TotVAhExpPhA    float32            `sunspec:"offset=76"`
	TotVAhExpPhB    float32            `sunspec:"offset=78"`
	TotVAhExpPhC    float32            `sunspec:"offset=80"`
	TotVAhImp       float32            `sunspec:"offset=82"`
	TotVAhImpPhA    float32            `sunspec:"offset=84"`
	TotVAhImpPhB    float32            `sunspec:"offset=86"`
	TotVAhImpPhC    float32            `sunspec:"offset=88"`
	TotVArhImpQ1    float32            `sunspec:"offset=90"`
	TotVArhImpQ1phA float32            `sunspec:"offset=92"`
	TotVArhImpQ1phB float32            `sunspec:"offset=94"`
	TotVArhImpQ1phC float32            `sunspec:"offset=96"`
	TotVArhImpQ2    float32            `sunspec:"offset=98"`
	TotVArhImpQ2phA float32            `sunspec:"offset=100"`
	TotVArhImpQ2phB float32            `sunspec:"offset=102"`
	TotVArhImpQ2phC float32            `sunspec:"offset=104"`
	TotVArhExpQ3    float32            `sunspec:"offset=106"`
	TotVArhExpQ3phA float32            `sunspec:"offset=108"`
	TotVArhExpQ3phB float32            `sunspec:"offset=110"`
	TotVArhExpQ3phC float32            `sunspec:"offset=112"`
	TotVArhExpQ4    float32            `sunspec:"offset=114"`
	TotVArhExpQ4phA float32            `sunspec:"offset=116"`
	TotVArhExpQ4phB float32            `sunspec:"offset=118"`
	TotVArhExpQ4phC float32            `sunspec:"offset=120"`
	Evt             sunspec.Bitfield32 `sunspec:"offset=122"`
}

func (block *Block211) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "ac_meter",
		Length: 124,
		Blocks: []smdx.BlockElement{
			{
				Length: 124,
				Points: []smdx.PointElement{
					{Id: A, Offset: 0, Type: typelabel.Float32, Units: "A", Mandatory: true, Label: "Amps", Description: "Total AC Current"},
					{Id: AphA, Offset: 2, Type: typelabel.Float32, Units: "A", Mandatory: true, Label: "Amps PhaseA", Description: "Phase A Current"},
					{Id: AphB, Offset: 4, Type: typelabel.Float32, Units: "A", Label: "Amps PhaseB", Description: "Phase B Current"},
					{Id: AphC, Offset: 6, Type: typelabel.Float32, Units: "A", Label: "Amps PhaseC", Description: "Phase C Current"},
					{Id: PhV, Offset: 8, Type: typelabel.Float32, Units: "V", Label: "Voltage LN", Description: "Line to Neutral AC Voltage (average of active phases)"},
					{Id: PhVphA, Offset: 10, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage AN", Description: "Phase Voltage AN"},
					{Id: PhVphB, Offset: 12, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage BN", Description: "Phase Voltage BN"},
					{Id: PhVphC, Offset: 14, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage CN", Description: "Phase Voltage CN"},
					{Id: PPV, Offset: 16, Type: typelabel.Float32, Units: "V", Label: "Voltage LL", Description: "Line to Line AC Voltage (average of active phases)"},
					{Id: PPVphAB, Offset: 18, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage AB", Description: "Phase Voltage AB"},
					{Id: PPVphBC, Offset: 20, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage BC", Description: "Phase Voltage BC"},
					{Id: PPVphCA, Offset: 22, Type: typelabel.Float32, Units: "V", Label: "Phase Voltage CA", Description: "Phase Voltage CA"},
					{Id: Hz, Offset: 24, Type: typelabel.Float32, Units: "Hz", Mandatory: true, Label: "Hz", Description: "Frequency"},
					{Id: W, Offset: 26, Type: typelabel.Float32, Units: "W", Mandatory: true, Label: "Watts", Description: "Total Real Power"},
					{Id: WphA, Offset: 28, Type: typelabel.Float32, Units: "W", Label: "Watts phase A", Description: ""},
					{Id: WphB, Offset: 30, Type: typelabel.Float32, Units: "W", Label: "Watts phase B", Description: ""},
					{Id: WphC, Offset: 32, Type: typelabel.Float32, Units: "W", Label: "Watts phase C", Description: ""},
					{Id: VA, Offset: 34, Type: typelabel.Float32, Units: "VA", Label: "VA", Description: "AC Apparent Power"},
					{Id: VAphA, Offset: 36, Type: typelabel.Float32, Units: "VA", Label: "VA phase A", Description: ""},
					{Id: VAphB, Offset: 38, Type: typelabel.Float32, Units: "VA", Label: "VA phase B", Description: ""},
					{Id: VAphC, Offset: 40, Type: typelabel.Float32, Units: "VA", Label: "VA phase C", Description: ""},
					{Id: VAR, Offset: 42, Type: typelabel.Float32, Units: "var", Label: "VAR", Description: "Reactive Power"},
					{Id: VARphA, Offset: 44, Type: typelabel.Float32, Units: "var", Label: "VAR phase A", Description: ""},
					{Id: VARphB, Offset: 46, Type: typelabel.Float32, Units: "var", Label: "VAR phase B", Description: ""},
					{Id: VARphC, Offset: 48, Type: typelabel.Float32, Units: "var", Label: "VAR phase C", Description: ""},
					{Id: PF, Offset: 50, Type: typelabel.Float32, Units: "PF", Label: "PF", Description: "Power Factor"},
					{Id: PFphA, Offset: 52, Type: typelabel.Float32, Units: "PF", Label: "PF phase A", Description: ""},
					{Id: PFphB, Offset: 54, Type: typelabel.Float32, Units: "PF", Label: "PF phase B", Description: ""},
					{Id: PFphC, Offset: 56, Type: typelabel.Float32, Units: "PF", Label: "PF phase C", Description: ""},
					{Id: TotWhExp, Offset: 58, Type: typelabel.Float32, Units: "Wh", Mandatory: true, Label: "Total Watt-hours Exported", Description: "Total Real Energy Exported"},
					{Id: TotWhExpPhA, Offset: 60, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Exported phase A", Description: ""},
					{Id: TotWhExpPhB, Offset: 62, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Exported phase B", Description: ""},
					{Id: TotWhExpPhC, Offset: 64, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Exported phase C", Description: ""},
					{Id: TotWhImp, Offset: 66, Type: typelabel.Float32, Units: "Wh", Mandatory: true, Label: "Total Watt-hours Imported", Description: "Total Real Energy Imported"},
					{Id: TotWhImpPhA, Offset: 68, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Imported phase A", Description: ""},
					{Id: TotWhImpPhB, Offset: 70, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Imported phase B", Description: ""},
					{Id: TotWhImpPhC, Offset: 72, Type: typelabel.Float32, Units: "Wh", Label: "Total Watt-hours Imported phase C", Description: ""},
					{Id: TotVAhExp, Offset: 74, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Exported", Description: "Total Apparent Energy Exported"},
					{Id: TotVAhExpPhA, Offset: 76, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Exported phase A", Description: ""},
					{Id: TotVAhExpPhB, Offset: 78, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Exported phase B", Description: ""},
					{Id: TotVAhExpPhC, Offset: 80, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Exported phase C", Description: ""},
					{Id: TotVAhImp, Offset: 82, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Imported", Description: "Total Apparent Energy Imported"},
					{Id: TotVAhImpPhA, Offset: 84, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Imported phase A", Description: ""},
					{Id: TotVAhImpPhB, Offset: 86, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Imported phase B", Description: ""},
					{Id: TotVAhImpPhC, Offset: 88, Type: typelabel.Float32, Units: "VAh", Label: "Total VA-hours Imported phase C", Description: ""},
					{Id: TotVArhImpQ1, Offset: 90, Type: typelabel.Float32, Units: "varh", Label: "Total VAR-hours Imported Q1", Description: "Total Reactive Energy Imported Quadrant 1"},
					{Id: TotVArhImpQ1phA, Offset: 92, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q1 phase A", Description: ""},
					{Id: TotVArhImpQ1phB, Offset: 94, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q1 phase B", Description: ""},
					{Id: TotVArhImpQ1phC, Offset: 96, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q1 phase C", Description: ""},
					{Id: TotVArhImpQ2, Offset: 98, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q2", Description: "Total Reactive Power Imported Quadrant 2"},
					{Id: TotVArhImpQ2phA, Offset: 100, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q2 phase A", Description: ""},
					{Id: TotVArhImpQ2phB, Offset: 102, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q2 phase B", Description: ""},
					{Id: TotVArhImpQ2phC, Offset: 104, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Imported Q2 phase C", Description: ""},
					{Id: TotVArhExpQ3, Offset: 106, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q3", Description: "Total Reactive Power Exported Quadrant 3"},
					{Id: TotVArhExpQ3phA, Offset: 108, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q3 phase A", Description: ""},
					{Id: TotVArhExpQ3phB, Offset: 110, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q3 phase B", Description: ""},
					{Id: TotVArhExpQ3phC, Offset: 112, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q3 phase C", Description: ""},
					{Id: TotVArhExpQ4, Offset: 114, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q4", Description: "Total Reactive Power Exported Quadrant 4"},
					{Id: TotVArhExpQ4phA, Offset: 116, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q4 Imported phase A", Description: ""},
					{Id: TotVArhExpQ4phB, Offset: 118, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q4 Imported phase B", Description: ""},
					{Id: TotVArhExpQ4phC, Offset: 120, Type: typelabel.Float32, Units: "varh", Label: "Total VAr-hours Exported Q4 Imported phase C", Description: ""},
					{Id: Evt, Offset: 122, Type: typelabel.Bitfield32, Mandatory: true, Label: "Events", Description: "Meter Event Flags"},
				},
			},
		}})
}
