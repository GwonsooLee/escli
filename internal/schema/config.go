/*
copyright 2020 the Escli authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package schema

type OldConfig Config0_0_3

type Config struct {
	Profile          string `yaml:"profile"`
	ElasticSearchURL string `yaml:"elasticsearch_url"`
	AWSRegion        string `yaml:"aws_region"`
}

type Config0_0_3 struct {
	ElasticSearchURL string `yaml:"elasticsearchurl"`
	AWSRegion        string `yaml:"awsregion"`
}
