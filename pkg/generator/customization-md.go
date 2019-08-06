// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

const (
	formDocumentation = `
# Installation Customization

This file documents the various environment variables you can set to change the functionality of the service broker.
If you are using the PCF Tile deployment, then you can manage all of these options through the operator forms.
If you are running your own, then you can set them in the application manifest of a PCF deployment, or in your pod configuration for Kubernetes.

{{ range $i, $f := .Forms }}{{ template "normalform" $f }}{{ end }}

## Custom Plans

You can specify custom plans for the following services.
The plans MUST be an array of flat JSON objects stored in their associated environment variable e.g. <code>[{...}, {...},...]</code>.
Each plan MUST have a unique UUID, if you modify the plan the UUID should stay the same to ensure previously provisioned services continue to work.
If you are using the PCF tile, it will generate the UUIDs for you.
DO NOT delete plans, instead you should change their labels to mark them as deprecated.

{{ range $i, $f := .ServicePlanForms -}}
	{{ template "customplanform" $f }}
{{- end }}

---------------------------------------

_Note: **Do not edit this file**, it was auto-generated by running <code>gcp-service-broker generate customization</code>. If you find an error, change the source code in <tt>customization-md.go</tt> or file a bug._

{{/*=======================================================================*/}}
{{ define "normalform" }}
## {{ .Label }}

{{ .Description }}

You can configure the following environment variables:

| Environment Variable | Type | Description |
|----------------------|------|-------------|
{{ range .Properties -}}
| <tt>{{upper .Name}}</tt>{{ if not .Optional }} <b>*</b>{{end}} | {{ .Type }} | <p>{{ .Label }} {{ .Description }}{{if .Default }} Default: <code>{{ js .Default }}</code>{{- end }}</p>|
{{ end }}
{{ end }}

{{/*=======================================================================*/}}
{{ define "customplanform" -}}
### {{ .Label }}

{{ .Description }}
To specify a custom plan manually, create the plan as JSON in a JSON array and store it in the environment variable: <tt>{{ upper .Name }}</tt>.

For example:
<code>
[{"id":"00000000-0000-0000-0000-000000000000", "name": "custom-plan-1"{{ range .Properties }}, "{{.Name}}": setme{{ end }}},...]
</code>

<table>
<tr>
  <th>JSON Property</th>
  <th>Type</th>
  <th>Label</th>
  <th>Details</th>
</tr>
<tr>
  <td><tt>id</tt></td>
  <td><i>string</i></td>
  <td>Plan UUID</td>
  <td>
    The UUID of the custom plan, use the <tt>uuidgen</tt> CLI command or [uuidgenerator.net](https://www.uuidgenerator.net/) to create one.
    <ul><li><b>Required</b></li></ul>
  </td>
</tr>
<tr>
  <td><tt>name</tt></td>
  <td><i>string</i></td>
  <td>Plan CLI Name</td>
  <td>
    The name of the custom plan used to provision it, must be lower-case, start with a letter a-z and contain only letters, numbers and dashes (-).
    <ul><li><b>Required</b></li></ul>
  </td>
</tr>

{{ range .Properties }}
<tr>
  <td><tt>{{ .Name }}</tt></td>
  <td><i>{{ .Type }}</i></td>
  <td>{{ .Label }}</td>
  <td>
  {{ .Description }}
  {{ template "variable-details-list" . }}
  </td>
</tr>
{{ end }}
</table>

{{ end }}

{{/*=======================================================================*/}}
{{ define "variable-details-list"}}

<ul>
  <li>{{ if .Optional }}<i>Optional</i>{{ else }}<b>Required</b>{{ end }}</li>

{{- if .Default }}
  <li>Default: <code>{{ js .Default }}</code></li>
{{- end }}

{{- if not .Configurable }}
  <li>This option _is not_ user configurable. It must be set to the default.</li>
{{- end }}

{{- if .Options }}
  <li>Valid Values:
  <ul>
    {{ range .Options }}<li><tt>{{ .Name }}</tt> - {{ .Label }}</li>{{ end }}
  </ul>
  </li>
{{- end }}
</ul>

{{ end }}
`
)

var (
	customizationTemplateFuncs = template.FuncMap{
		"upper": strings.ToUpper,
	}
	formDocumentationTemplate = template.Must(template.New("name").Funcs(customizationTemplateFuncs).Parse(formDocumentation))
)

func GenerateCustomizationMd() string {
	tileForms := GenerateForms()

	var buf bytes.Buffer
	if err := formDocumentationTemplate.Execute(&buf, tileForms); err != nil {
		log.Fatalf("Error rendering template: %s", err)
	}

	return cleanMdOutput(buf.String())
}

// Remove trailing whitespace from the document and every line
func cleanMdOutput(text string) string {
	text = strings.TrimSpace(text)

	lines := strings.Split(text, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t")
	}

	return strings.Join(lines, "\n")
}
