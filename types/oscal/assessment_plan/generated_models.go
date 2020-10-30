// Code generated by https://github.com/GoComply/metaschema; DO NOT EDIT.
package assessment_plan

import (
	"github.com/gocomply/oscalkit/types/oscal/validation_root"

	"github.com/gocomply/oscalkit/types/oscal/system_security_plan"

	"github.com/gocomply/oscalkit/types/oscal/assessment_common"
)

// An assessment plan, such as those provided by a FedRAMP assessor.
type AssessmentPlan struct {

	// A RFC 4122 version 4 Universally Unique Identifier (UUID) for the containing object.
	Uuid string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`

	// Provides information about the publication and availability of the containing document.
	Metadata *Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// Used by the SAP to import information about the system being assessed.
	ImportSsp *ImportSsp `xml:"import-ssp,omitempty" json:"importSsp,omitempty"`
	// Identifies the controls and control being assessed and their control objectives. In the assessment plans, these are the planned controls and objectives. In the assessment results, these are the actual controls and objectives, and reflects any changes from the plan.
	Objectives *Objectives `xml:"objectives,omitempty" json:"objectives,omitempty"`
	// Identifies system elements being assessed, such as components, inventory items, and locations. In the assessment plan, this identifies the planned assessment subject. In the assessment results this is the actual assessment subject, and reflects any changes from the plan.
	AssessmentSubjects *AssessmentSubjects `xml:"assessment-subjects,omitempty" json:"assessmentSubjects,omitempty"`
	// Identifies the assets used to perform this assessment, such as the assessment team, scanning tools, and assumptions.
	Assets *Assets `xml:"assets,omitempty" json:"assets,omitempty"`
	// Identifies the assessment activities and schedule. In the assessment plan, these are planned activities. In the assessment results, these are the actual activities performed.
	AssessmentActivities *AssessmentActivities `xml:"assessment-activities,omitempty" json:"assessmentActivities,omitempty"`
	// A collection of citations and resource references.
	BackMatter *BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

// Identifies system elements being assessed, such as components, inventory items, and locations. In the assessment plan, this identifies the planned assessment subject. In the assessment results this is the actual assessment subject, and reflects any changes from the plan.
type AssessmentSubjects struct {

	// Identifies exactly what will be the focus of this assessment. Anything not explicitly defined is out-of-scope.
	Includes []IncludeSubject `xml:"include-subject,omitempty" json:"includes,omitempty"`
	// Identifies what is explicitly excluded from this assessment. Used to remove a subset of items from groups of explicitly included items. Also used to explicitly clarify off-limit items, such as hosts to avoid scanning.
	Excludes []ExcludeSubject `xml:"exclude-subject,omitempty" json:"excludes,omitempty"`
	// Allows control objectives, users, components, and inventory-items to be defined within the assessment plan or assessment results for circumstances where they are not appropriately defined in the SSP. NOTE: Use the assessment plan or assessment results metadata to define additional locations if needed.
	LocalDefinitions *LocalDefinitions `xml:"local-definitions,omitempty" json:"localDefinitions,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
}

// Allows control objectives, users, components, and inventory-items to be defined within the assessment plan or assessment results for circumstances where they are not appropriately defined in the SSP. NOTE: Use the assessment plan or assessment results metadata to define additional locations if needed.
type LocalDefinitions struct {

	// Used to add any components, not defined via the System Security Plan (AR->AP->SSP)
	Components ComponentMultiplexer `xml:"component,omitempty" json:"components,omitempty"`
	// Used to add any inventory-items, not defined via the System Security Plan (AR->AP->SSP)
	InventoryItems InventoryItemMultiplexer `xml:"inventory-item,omitempty" json:"inventory-items,omitempty"`
	// Used to add any users, not defined via the System Security Plan (AR->AP->SSP)
	Users UserMultiplexer `xml:"user,omitempty" json:"users,omitempty"`
	// Additional commentary on the parent item.
	Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
}

type ComponentMultiplexer = system_security_plan.ComponentMultiplexer

type InventoryItemMultiplexer = system_security_plan.InventoryItemMultiplexer

type UserMultiplexer = system_security_plan.UserMultiplexer

type AssessmentActivities = assessment_common.AssessmentActivities

type Assets = assessment_common.Assets

type BackMatter = validation_root.BackMatter

type Component = system_security_plan.Component

type ExcludeSubject = assessment_common.ExcludeSubject

type ImportSsp = assessment_common.ImportSsp

type IncludeSubject = assessment_common.IncludeSubject

type InventoryItem = system_security_plan.InventoryItem

type Metadata = validation_root.Metadata

type Objectives = assessment_common.Objectives

type Remarks = validation_root.Remarks

type User = system_security_plan.User
