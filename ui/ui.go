package ui

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"bytes"
	"html/template"

	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
)

// PageData - data used on the HTML pages
type PageData struct {
	AppName    string
	AppVersion string
	// URLHome - URL for the service home page
	URLHome         string
	URLCollections  string
	URLCollection   string
	URLItems        string
	URLJSON         string
	CollectionTitle string
	Layer           *data.Layer
	FeatureID       string
	UseMap          bool
}

var htmlTemp struct {
	home        *template.Template
	conformance *template.Template
	collections *template.Template
	collection  *template.Template
	items       *template.Template
	item        *template.Template
}

var HTMLDynamicLoad bool

func init() {
	HTMLDynamicLoad = false
}

func loadTemplate(curr *template.Template, filename ...string) *template.Template {
	if curr == nil || HTMLDynamicLoad {
		return template.Must(template.ParseFiles(filename...))
	}
	return curr
}

func PageHome() *template.Template {
	htmlTemp.home = loadTemplate(htmlTemp.home, "html/page.gohtml", "html/home.gohtml")
	return htmlTemp.home
}
func PageConformance() *template.Template {
	htmlTemp.conformance = loadTemplate(htmlTemp.conformance, "html/page.gohtml", "html/conformance.gohtml")
	return htmlTemp.conformance
}
func PageCollections() *template.Template {
	htmlTemp.collections = loadTemplate(htmlTemp.collections, "html/page.gohtml", "html/collections.gohtml")
	return htmlTemp.collections
}
func PageCollection() *template.Template {
	htmlTemp.collection = loadTemplate(htmlTemp.collection, "html/page.gohtml", "html/collection.gohtml")
	return htmlTemp.collection
}
func PageItems() *template.Template {
	htmlTemp.items = loadTemplate(htmlTemp.items, "html/page.gohtml", "html/map_script.gohtml", "html/items.gohtml")
	return htmlTemp.items
}
func PageItem() *template.Template {
	htmlTemp.item = loadTemplate(htmlTemp.item, "html/page.gohtml", "html/map_script.gohtml", "html/item.gohtml")
	return htmlTemp.item
}

// RenderHTML tbd
func RenderHTML(temp *template.Template, content interface{}, context interface{}) ([]byte, error) {
	bodyData := map[string]interface{}{
		"config":  config.Configuration,
		"context": context,
		"data":    content}
	contentBytes, err := renderTemplate(temp, bodyData)
	if err != nil {
		return contentBytes, err
	}
	return contentBytes, err
}

func renderTemplate(temp *template.Template, data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer

	if err := temp.ExecuteTemplate(&buf, "page", data); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
