/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/penny-vault/sonrai/db"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list objects in Sonrai",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		cmdUrl := fmt.Sprintf("%s/api/v1/jobs", viper.GetString("api.host"))
		log.Debug().Str("URL", cmdUrl).Msg("fetching jobs from API")
		resp, err := client.R().Get(cmdUrl)
		if err != nil {
			log.Error().Err(err).Msg("received error from server")
			os.Exit(1)
		}
		if resp.StatusCode() >= 400 {
			log.Error().Int("StatusCode", resp.StatusCode()).Msg("received error from server")
			os.Exit(1)
		}

		body := resp.Body()
		jobs := []*db.Job{}
		err = json.Unmarshal(body, &jobs)
		if err != nil {
			log.Error().Err(err).Msg("failed de-serializing JSON")
			os.Exit(1)
		}

		for _, job := range jobs {
			fmt.Printf("%+v\n", job)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
