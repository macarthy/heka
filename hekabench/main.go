/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/
package main

import (
	"heka/client"
	"log"
	"os"
	"time"
)

func main() {
	var err error
	addrStr := "127.0.0.1:5565"
	sender, err := hekaclient.NewUdpSender(&addrStr)
	if err != nil {
		log.Fatalf("Error creating sender: %s\n", err.Error())
	}
	encoder := hekaclient.NewGobEncoder()
	timestamp := time.Now()
	hostname, _ := os.Hostname()
	message := hekaclient.Message{
		Type: "hekabench", Timestamp: timestamp,
		Logger: "hekabench", Severity: 6,
		Payload: "Test Payload", Env_version: "0.8",
		Pid: os.Getpid(), Hostname: hostname,
	}
	msgBytes, err := encoder.EncodeMessage(&message)

	for {
		err = sender.SendMessage(msgBytes)
		if err != nil {
			log.Printf("Error sending message: %s\n", err.Error())
		}
	}
}