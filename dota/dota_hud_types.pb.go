// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_hud_types.proto

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EHeroSelectionText int32

const (
	EHeroSelectionText_k_EHeroSelectionText_Invalid                              EHeroSelectionText = -1
	EHeroSelectionText_k_EHeroSelectionText_None                                 EHeroSelectionText = 0
	EHeroSelectionText_k_EHeroSelectionText_ChooseHero                           EHeroSelectionText = 1
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_YouFirst           EHeroSelectionText = 2
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_TheyFirst          EHeroSelectionText = 3
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_BanSelected_YouFirst        EHeroSelectionText = 4
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst       EHeroSelectionText = 5
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_YouPicking                  EHeroSelectionText = 6
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TheyPicking                 EHeroSelectionText = 7
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TeammateRandomed            EHeroSelectionText = 8
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_YouPicking_LosingGold       EHeroSelectionText = 9
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold      EHeroSelectionText = 10
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_ChooseCaptain           EHeroSelectionText = 11
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain EHeroSelectionText = 12
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouSelect               EHeroSelectionText = 13
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheySelect              EHeroSelectionText = 14
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouBan                  EHeroSelectionText = 15
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheyBan                 EHeroSelectionText = 16
)

var EHeroSelectionText_name = map[int32]string{
	-1: "k_EHeroSelectionText_Invalid",
	0:  "k_EHeroSelectionText_None",
	1:  "k_EHeroSelectionText_ChooseHero",
	2:  "k_EHeroSelectionText_AllDraft_Planning_YouFirst",
	3:  "k_EHeroSelectionText_AllDraft_Planning_TheyFirst",
	4:  "k_EHeroSelectionText_AllDraft_BanSelected_YouFirst",
	5:  "k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst",
	6:  "k_EHeroSelectionText_AllDraft_YouPicking",
	7:  "k_EHeroSelectionText_AllDraft_TheyPicking",
	8:  "k_EHeroSelectionText_AllDraft_TeammateRandomed",
	9:  "k_EHeroSelectionText_AllDraft_YouPicking_LosingGold",
	10: "k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold",
	11: "k_EHeroSelectionText_CaptainsMode_ChooseCaptain",
	12: "k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain",
	13: "k_EHeroSelectionText_CaptainsMode_YouSelect",
	14: "k_EHeroSelectionText_CaptainsMode_TheySelect",
	15: "k_EHeroSelectionText_CaptainsMode_YouBan",
	16: "k_EHeroSelectionText_CaptainsMode_TheyBan",
}
var EHeroSelectionText_value = map[string]int32{
	"k_EHeroSelectionText_Invalid":                              -1,
	"k_EHeroSelectionText_None":                                 0,
	"k_EHeroSelectionText_ChooseHero":                           1,
	"k_EHeroSelectionText_AllDraft_Planning_YouFirst":           2,
	"k_EHeroSelectionText_AllDraft_Planning_TheyFirst":          3,
	"k_EHeroSelectionText_AllDraft_BanSelected_YouFirst":        4,
	"k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst":       5,
	"k_EHeroSelectionText_AllDraft_YouPicking":                  6,
	"k_EHeroSelectionText_AllDraft_TheyPicking":                 7,
	"k_EHeroSelectionText_AllDraft_TeammateRandomed":            8,
	"k_EHeroSelectionText_AllDraft_YouPicking_LosingGold":       9,
	"k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold":      10,
	"k_EHeroSelectionText_CaptainsMode_ChooseCaptain":           11,
	"k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain": 12,
	"k_EHeroSelectionText_CaptainsMode_YouSelect":               13,
	"k_EHeroSelectionText_CaptainsMode_TheySelect":              14,
	"k_EHeroSelectionText_CaptainsMode_YouBan":                  15,
	"k_EHeroSelectionText_CaptainsMode_TheyBan":                 16,
}

func (x EHeroSelectionText) Enum() *EHeroSelectionText {
	p := new(EHeroSelectionText)
	*p = x
	return p
}
func (x EHeroSelectionText) String() string {
	return proto.EnumName(EHeroSelectionText_name, int32(x))
}
func (x *EHeroSelectionText) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EHeroSelectionText_value, data, "EHeroSelectionText")
	if err != nil {
		return err
	}
	*x = EHeroSelectionText(value)
	return nil
}
func (EHeroSelectionText) EnumDescriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

var E_HudLocalizeToken = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50501,
	Name:          "dota.hud_localize_token",
	Tag:           "bytes,50501,opt,name=hud_localize_token,json=hudLocalizeToken",
	Filename:      "dota_hud_types.proto",
}

func init() {
	proto.RegisterEnum("dota.EHeroSelectionText", EHeroSelectionText_name, EHeroSelectionText_value)
	proto.RegisterExtension(E_HudLocalizeToken)
}

func init() { proto.RegisterFile("dota_hud_types.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xcb, 0x6e, 0x13, 0x3f,
	0x14, 0xc6, 0x9b, 0x7f, 0xfb, 0x07, 0x6a, 0x28, 0x58, 0x16, 0x0b, 0x63, 0x81, 0x5a, 0xd4, 0x4d,
	0x42, 0x43, 0x12, 0x35, 0x11, 0xa5, 0x20, 0x16, 0x4d, 0x2f, 0x14, 0xa9, 0xd0, 0x50, 0x22, 0x50,
	0x10, 0xd2, 0xe0, 0x66, 0xdc, 0x89, 0xc9, 0xc4, 0x1e, 0xcd, 0x78, 0x10, 0x45, 0x2c, 0x58, 0xf3,
	0x28, 0x7e, 0x0f, 0xde, 0xc0, 0x1b, 0x9e, 0x06, 0x34, 0xe4, 0x32, 0xa1, 0x9d, 0xeb, 0xf2, 0xf8,
	0x7c, 0xdf, 0xcf, 0x67, 0xac, 0x4f, 0x07, 0xdc, 0xb6, 0xa5, 0xa2, 0xd6, 0x20, 0xb4, 0x2d, 0x75,
	0xee, 0xb1, 0xa0, 0xe6, 0xf9, 0x52, 0x49, 0xb4, 0x14, 0x55, 0xc9, 0x9a, 0x23, 0xa5, 0xe3, 0xb2,
	0xfa, 0xdf, 0xda, 0x69, 0x78, 0x56, 0xb7, 0x59, 0xd0, 0xf7, 0xb9, 0xa7, 0xa4, 0x3f, 0xee, 0x7b,
	0xf0, 0x6b, 0x05, 0xa0, 0xfd, 0x43, 0xe6, 0xcb, 0x37, 0xcc, 0x65, 0x7d, 0xc5, 0xa5, 0xe8, 0xb2,
	0x2f, 0x0a, 0x55, 0xc0, 0xdd, 0xa1, 0x75, 0xb9, 0x6e, 0xbd, 0x10, 0x9f, 0xa9, 0xcb, 0x6d, 0xf8,
	0x7b, 0xfa, 0x95, 0xd0, 0x3d, 0x70, 0x27, 0xb1, 0xf5, 0x95, 0x14, 0x0c, 0x2e, 0xa0, 0x43, 0xb0,
	0x9a, 0x78, 0xbc, 0x3b, 0x90, 0x32, 0x60, 0x51, 0x1d, 0x96, 0xc8, 0xba, 0x36, 0x78, 0x75, 0x7d,
	0xef, 0xb8, 0xbb, 0x63, 0x45, 0x15, 0x6b, 0xd6, 0x3a, 0xd7, 0x86, 0x06, 0xa0, 0x9e, 0xe8, 0xb4,
	0xe3, 0xba, 0x7b, 0x3e, 0x3d, 0x53, 0x56, 0xc7, 0xa5, 0x42, 0x70, 0xe1, 0x58, 0x3d, 0x19, 0x1e,
	0x70, 0x3f, 0x50, 0xf0, 0x3f, 0xd2, 0xd4, 0x06, 0xd7, 0x13, 0x9d, 0xd3, 0x65, 0xe8, 0x13, 0x68,
	0x14, 0x24, 0x75, 0x07, 0xec, 0x7c, 0x8c, 0x5a, 0x24, 0x2d, 0x6d, 0x70, 0xa3, 0x20, 0x6a, 0xa6,
	0x43, 0x02, 0x6c, 0x66, 0xb3, 0xda, 0x54, 0x8c, 0x4f, 0x98, 0x1d, 0x0f, 0xb6, 0x44, 0x1e, 0x69,
	0x83, 0x37, 0xb3, 0x69, 0x49, 0x4a, 0xe4, 0x81, 0x66, 0x71, 0x5e, 0x3c, 0xde, 0xff, 0x64, 0x4b,
	0x1b, 0xdc, 0x2c, 0x0e, 0x8c, 0x27, 0x7c, 0x0f, 0xca, 0xd9, 0xc4, 0x9e, 0x0c, 0x3b, 0xbc, 0x3f,
	0xe4, 0xc2, 0x81, 0x57, 0x48, 0x55, 0x1b, 0x5c, 0xce, 0xc6, 0xc4, 0xfd, 0xe8, 0x03, 0xa8, 0x64,
	0x7b, 0x47, 0xd7, 0x98, 0x9a, 0x5f, 0x25, 0x0f, 0xb5, 0xc1, 0x95, 0x6c, 0xf3, 0x39, 0x01, 0x92,
	0xa0, 0x96, 0xe3, 0xce, 0xe8, 0x68, 0x44, 0x15, 0x3b, 0xa1, 0xc2, 0x96, 0x23, 0x66, 0xc3, 0x6b,
	0xe4, 0xa9, 0x36, 0x78, 0x2b, 0x07, 0x71, 0x41, 0x65, 0x75, 0xa8, 0x90, 0x3e, 0x1d, 0xd1, 0xfc,
	0xc7, 0x89, 0x47, 0xb7, 0x8e, 0x64, 0xc0, 0x85, 0xf3, 0x5c, 0xba, 0x36, 0x5c, 0x2e, 0xf2, 0x38,
	0x89, 0x52, 0xe4, 0x83, 0x56, 0xe1, 0x1f, 0x38, 0x8f, 0x04, 0xe4, 0xb1, 0x36, 0xb8, 0x55, 0xf8,
	0x5f, 0xce, 0x33, 0xd3, 0x82, 0xbc, 0x4b, 0x3d, 0x45, 0xb9, 0x08, 0x5e, 0x4a, 0x9b, 0x4d, 0x82,
	0x3f, 0x29, 0xc1, 0xeb, 0x19, 0x41, 0x4e, 0x97, 0xa1, 0x6f, 0x60, 0x3b, 0x9f, 0xf4, 0x8e, 0x72,
	0xc5, 0x85, 0x73, 0x20, 0xfd, 0x7f, 0x99, 0x37, 0xc8, 0x33, 0x6d, 0xf0, 0x76, 0x3e, 0x33, 0xc5,
	0x00, 0x7d, 0x04, 0x1b, 0xf9, 0xf4, 0x9e, 0x0c, 0xc7, 0xa7, 0x70, 0x85, 0xd4, 0xb5, 0xc1, 0x1b,
	0xf9, 0xbc, 0x99, 0x04, 0x9d, 0x82, 0x6a, 0x3e, 0x21, 0x7a, 0x85, 0x09, 0xe2, 0x26, 0x69, 0x68,
	0x83, 0xab, 0xf9, 0x88, 0x58, 0x93, 0x1a, 0xdf, 0x8b, 0x57, 0x6a, 0x53, 0x01, 0x6f, 0x65, 0xc4,
	0x37, 0xa1, 0x3f, 0x35, 0xbe, 0x97, 0xee, 0x12, 0x99, 0xc3, 0x8c, 0xf8, 0x26, 0x09, 0x9e, 0xbc,
	0x06, 0x28, 0x5a, 0x8b, 0xae, 0xec, 0x53, 0x97, 0x7f, 0x65, 0x96, 0x92, 0x43, 0x26, 0xd0, 0xfd,
	0xda, 0x78, 0x29, 0xd6, 0xa6, 0x4b, 0xb1, 0xb6, 0x2f, 0xc2, 0xd1, 0x5b, 0xea, 0x86, 0xec, 0xd8,
	0x8b, 0xcc, 0x02, 0xfc, 0xf3, 0xc7, 0xe2, 0x5a, 0xa9, 0xbc, 0x7c, 0x02, 0x07, 0xa1, 0x7d, 0x34,
	0x51, 0x77, 0x23, 0x71, 0x7b, 0xf1, 0x7b, 0x69, 0xe1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x45, 0x4e, 0x43, 0x72, 0x07, 0x00, 0x00,
}
