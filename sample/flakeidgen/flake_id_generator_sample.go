// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/hazelcast/hazelcast-go-client"
)

func main() {

	config := hazelcast.NewConfig()
	config.NetworkConfig().AddAddress("127.0.0.1:5701")
	client, err := hazelcast.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	flakeIDGenerator, _ := client.GetFlakeIDGenerator("generator")
	id, _ := flakeIDGenerator.NewID()
	fmt.Printf("new id : %d", id)

	flakeIDGenerator.Destroy()
	client.Shutdown()

}
