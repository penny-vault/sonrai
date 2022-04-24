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
package nats

import (
	"time"

	natsApi "github.com/nats-io/nats-server/v2/server"
	"github.com/rs/zerolog/log"
)

var Options *natsApi.Options
var Server *natsApi.Server

func Serve() {
	Options = &natsApi.Options{}
	Server, err := natsApi.NewServer(Options)

	if err != nil {
		panic(err)
	}

	go Server.Start()

	if !Server.ReadyForConnections(4 * time.Second) {
		log.Error().Msg("nats not starting in a reasonable amount of time - not ready for nats connections after 4 seconds")
		panic("nats failed")
	}
}

func Healthy() bool {
	return Server.Running()
}
