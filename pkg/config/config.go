// Copyright © 2017 Syndesis Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

func NewWithDefaults() Config {
	return Config{
		HTTPConfig{
			Address: "",
			Port:    8080,
		},
		WebhookConfig{},
		GitHubAppConfig{},
		RepoConfig{
			Labels: LabelConfig{
				Approved: "approved",
			},
		},
		nil,
	}
}

type Config struct {
	HTTP        HTTPConfig            `mapstructure:"http"`
	Webhook     WebhookConfig         `mapstructure:"webhook"`
	GitHubApp   GitHubAppConfig       `mapstructure:"github"`
	DefaultRepo RepoConfig            `mapstructure:"defaults"`
	Repos       map[string]RepoConfig `mapstructure:"repos"`
}

type HTTPConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	TLSCert string `mapstructure:"tlsCert"`
	TLSKey  string `mapstructure:"tlsKey"`
}

type WebhookConfig struct {
	Secret string `mapstructure:"secret"`
}

type GitHubAppConfig struct {
	AppID          int64  `mapstructure:"appId"`
	PrivateKeyFile string `mapstructure:"privateKey"`
}

type RepoConfig struct {
	Disabled    bool        `mapstructure:"disabled"`
	Labels      LabelConfig `mapstructure:"labels"`
	WipPatterns []string    `mapstructure:"wipPatterns"`
}

type LabelConfig struct {
	NewIssues       []string `mapstructure:"newIssues"`
	Wip             []string `mapstructure:"wip"`
	ReviewRequested string   `mapstructure:"reviewRequested"`
	Approved        string   `mapstructure:"approved"`
}
