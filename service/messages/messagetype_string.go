// Code generated by "stringer -type=MessageType"; DO NOT EDIT

package messages

import "fmt"

const _MessageType_name = "HELLOHELLOOKDOHAVEHAVECONNECTREQUESTRESULTLISTENSTOPLISTENENDMOCK"

var _MessageType_index = [...]uint8{0, 5, 12, 18, 22, 29, 36, 42, 48, 58, 61, 65}

func (i MessageType) String() string {
	if i < 0 || i >= MessageType(len(_MessageType_index)-1) {
		return fmt.Sprintf("MessageType(%d)", i)
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
