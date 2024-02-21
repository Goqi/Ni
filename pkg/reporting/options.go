package reporting

import (
	"Ni/pkg/reporting/exporters/es"
	"Ni/pkg/reporting/exporters/jsonexporter"
	"Ni/pkg/reporting/exporters/jsonl"
	"Ni/pkg/reporting/exporters/markdown"
	"Ni/pkg/reporting/exporters/sarif"
	"Ni/pkg/reporting/exporters/splunk"
	"Ni/pkg/reporting/trackers/github"
	"Ni/pkg/reporting/trackers/gitlab"
	"Ni/pkg/reporting/trackers/jira"
	"github.com/projectdiscovery/retryablehttp-go"
)

// Options is a configuration file for nuclei reporting module
type Options struct {
	// AllowList contains a list of allowed events for reporting module
	AllowList *Filter `yaml:"allow-list"`
	// DenyList contains a list of denied events for reporting module
	DenyList *Filter `yaml:"deny-list"`
	// GitHub contains configuration options for GitHub Issue Tracker
	GitHub *github.Options `yaml:"github"`
	// GitLab contains configuration options for GitLab Issue Tracker
	GitLab *gitlab.Options `yaml:"gitlab"`
	// Jira contains configuration options for Jira Issue Tracker
	Jira *jira.Options `yaml:"jira"`
	// MarkdownExporter contains configuration options for Markdown Exporter Module
	MarkdownExporter *markdown.Options `yaml:"markdown"`
	// SarifExporter contains configuration options for Sarif Exporter Module
	SarifExporter *sarif.Options `yaml:"sarif"`
	// ElasticsearchExporter contains configuration options for Elasticsearch Exporter Module
	ElasticsearchExporter *es.Options `yaml:"elasticsearch"`
	// SplunkExporter contains configuration options for splunkhec Exporter Module
	SplunkExporter *splunk.Options `yaml:"splunkhec"`
	// JSONExporter contains configuration options for JSON Exporter Module
	JSONExporter *jsonexporter.Options `yaml:"json"`
	// JSONLExporter contains configuration options for JSONL Exporter Module
	JSONLExporter *jsonl.Options `yaml:"jsonl"`

	HttpClient *retryablehttp.Client `yaml:"-"`
	OmitRaw    bool                  `yaml:"-"`
}
