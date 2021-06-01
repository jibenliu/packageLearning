//go:generate stringer -type=Pill
package stringKiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen
	COVID19
	HIV
	OIV
	SSAS = HIV
)