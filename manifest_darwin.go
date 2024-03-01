// manifest_darwin.go - Install and Uninstall manifest file for OS X.
// Copyright (c) 2018 - 2024  Sasha Hilton <sashahilton00@users.noreply.github.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package host

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// getTargetNames returns a slice of absolute paths to native messaging host manifest
// locations for OS X.
//
// See https://developer.chrome.com/extensions/nativeMessaging#native-messaging-host-location-nix
func (h *Host) getTargetNames() []string {
	var targets []string

	// chrome
	target := "/Library/Google/Chrome/NativeMessagingHosts"

	if os.Getuid() != 0 {
		homeDir, _ := os.UserHomeDir()
		target = homeDir + "/Library/Application Support/Google/Chrome/NativeMessagingHosts"
	}

	targets = append(targets, filepath.Join(target, h.AppName+".json"))

	// firefox
	target = "/Library/Application Support/Mozilla/NativeMessagingHosts"

	if os.Getuid() != 0 {
		homeDir, _ := os.UserHomeDir()
		target = homeDir + "/Library/Application Support/Mozilla/NativeMessagingHosts"
	}

	targets = append(targets, filepath.Join(target, h.AppName+".json"))

	return targets
}

// Install creates native-messaging manifest file on appropriate location. It
// will return error when it come across one.
//
// See https://developer.chrome.com/extensions/nativeMessaging#native-messaging-host-location-nix
func (h *Host) Install() error {
	manifest, _ := json.MarshalIndent(h, "", "  ")

	targetNames := h.getTargetNames()

	for _, targetName := range targetNames {
		if err := osMkdirAll(filepath.Dir(targetName), 0755); err != nil {
			return err
		}

		if err := ioutilWriteFile(targetName, manifest, 0644); err != nil {
			return err
		}

		log.Printf("Installed: %s", targetName)
	}

	return nil
}

// Uninstall removes native-messaging manifest file from installed location.
//
// See https://developer.chrome.com/extensions/nativeMessaging#native-messaging-host-location-nix
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#manifest_location
func (h *Host) Uninstall() {
	targetNames := h.getTargetNames()

	for _, targetName := range targetNames {
		if err := os.Remove(targetName); err != nil {
			// It might never have been installed.
			log.Print(err)
		}

		if err := os.Remove(h.ExecName); err != nil {
			// It might be locked by current process.
			log.Print(err)
		}

		if err := os.Remove(h.ExecName + ".chk"); err != nil {
			// It might not exist.
			log.Print(err)
		}

		log.Printf("Uninstalled: %s", targetName)
	}

	// Exit gracefully.
	runtimeGoexit()
}
