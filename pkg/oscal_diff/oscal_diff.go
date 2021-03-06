package oscal_diff

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocomply/oscalkit/pkg/oscal/constants"
	"github.com/gocomply/oscalkit/types/oscal"
	"github.com/pmezard/go-difflib/difflib"
)

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
}

func Diff(a *oscal.OSCAL, b *oscal.OSCAL) (string, error) {
	if a.DocumentType() != b.DocumentType() {
		return "", fmt.Errorf("Could not compare OSCAL resources, type mismatch '%s' vs '%s'", a.DocumentType(), b.DocumentType())
	}

	hideFormatRelatedDifferences(a, b)
	as := spewConfig.Sdump(a)
	bs := spewConfig.Sdump(b)

	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(as),
		B:        difflib.SplitLines(bs),
		FromFile: "a1",
		ToFile:   "a2",
		Context:  3,
	}
	return difflib.GetUnifiedDiffString(diff)
}

// Hide xml vs json related differences
func hideFormatRelatedDifferences(a *oscal.OSCAL, b *oscal.OSCAL) {
	switch a.DocumentType() {
	case constants.CatalogDocument:
		if a.Catalog.XMLName.Space != b.Catalog.XMLName.Space {
			if a.Catalog.XMLName.Space == "" {
				a.Catalog.XMLName = b.Catalog.XMLName
			}
			if b.Catalog.XMLName.Space == "" {
				b.Catalog.XMLName = a.Catalog.XMLName
			}
		}
	case constants.ProfileDocument:
		if a.Profile.XMLName.Space != b.Profile.XMLName.Space {
			if a.Profile.XMLName.Space == "" {
				a.Profile.XMLName = b.Profile.XMLName
			}
			if b.Profile.XMLName.Space == "" {
				b.Profile.XMLName = a.Profile.XMLName
			}
		}
	case constants.SSPDocument:
		if a.SystemSecurityPlan.XMLName.Space != b.SystemSecurityPlan.XMLName.Space {
			if a.SystemSecurityPlan.XMLName.Space == "" {
				a.SystemSecurityPlan.XMLName = b.SystemSecurityPlan.XMLName
			}
			if b.SystemSecurityPlan.XMLName.Space == "" {
				b.SystemSecurityPlan.XMLName = a.SystemSecurityPlan.XMLName
			}
		}
	}
}
