package constants

import (
	"fmt"
	"strings"
)

// Representing OSCAL file format. XML, JSON, YAML, ...
type DocumentFormat int

const (
	UnknownFormat DocumentFormat = iota
	XmlFormat
	JsonFormat
	YamlFormat
)

func NewDocumentFormat(format string) DocumentFormat {
	switch strings.ToLower(format) {
	case "xml":
		return XmlFormat
	case "json":
		return JsonFormat
	case "yaml":
		return YamlFormat
	}
	return UnknownFormat
}

func (f DocumentFormat) String() string {
	switch f {
	case XmlFormat:
		return "xml"
	case JsonFormat:
		return "json"
	case YamlFormat:
		return "yaml"
	case UnknownFormat:
		return "unknown"
	}
	return fmt.Sprintf("unrecognized: %d", f)
}

type DocumentType int

const (
	UnknownDocument = iota
	CatalogDocument
	ProfileDocument
	SSPDocument
	ComponentDocument
	POAMDocument
	AssessmentPlanDocument
	AssessmentResultsDocument
)

func (t DocumentType) String() string {
	switch t {
	case CatalogDocument:
		return "Catalog"
	case ProfileDocument:
		return "Profile"
	case SSPDocument:
		return "System Security Plan"
	case ComponentDocument:
		return "Component"
	case POAMDocument:
		return "Plan of Action and Milestones"
	case AssessmentPlanDocument:
		return "Assessment Plan"
	case AssessmentResultsDocument:
		return "Assessment Result"
	default:
		return "Unrecognized Document Type"
	}
}
