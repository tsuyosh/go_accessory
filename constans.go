// Copyright 2014, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accessory

import (
	"errors"

	"github.com/kylelemons/gousb/usb"
)

// #include "f_accessory.h"
import "C"

const (
	USB_VENDOR_ID      usb.ID = C.USB_ACCESSORY_VENDOR_ID
	USB_PRODUCT_ID            = C.USB_ACCESSORY_PRODUCT_ID
	USB_ADB_PRODUCT_ID        = C.USB_ACCESSORY_ADB_PRODUCT_ID
)

const (
	STRING_MANUFACTURER uint16 = C.ACCESSORY_STRING_MANUFACTURER
	STRING_MODEL               = C.ACCESSORY_STRING_MODEL
	STRING_DESCRIPTION         = C.ACCESSORY_STRING_DESCRIPTION
	STRING_VERSION             = C.ACCESSORY_STRING_VERSION
	STRING_URI                 = C.ACCESSORY_STRING_URI
	STRING_SERIAL              = C.ACCESSORY_STRING_SERIAL
)

// requestType
const (
	RTYPE_IN  = (uint8(usb.ENDPOINT_DIR_IN) | usb.REQUEST_TYPE_VENDOR)
	RTYPE_OUT = (uint8(usb.ENDPOINT_DIR_OUT) | usb.REQUEST_TYPE_VENDOR)
)

// requests
const (
	GET_PROTOCOL        uint8 = C.ACCESSORY_GET_PROTOCOL
	SEND_STRING               = C.ACCESSORY_SEND_STRING
	START                     = C.ACCESSORY_START
	REGISTER_HID              = C.ACCESSORY_REGISTER_HID
	UNREGISTER_HID            = C.ACCESSORY_UNREGISTER_HID
	SET_HID_REPORT_DESC       = C.ACCESSORY_SET_HID_REPORT_DESC
	SEND_HID_EVENT            = C.ACCESSORY_SEND_HID_EVENT
	SET_AUDIO_MODE            = C.ACCESSORY_SET_AUDIO_MODE
)

// errors
var (
	ErrorNoAccessoryDevice   error = errors.New("No accessory device")
	ErrorFailedToGetProtocol       = errors.New("Failed to get protocol")
)
