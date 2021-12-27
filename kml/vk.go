package kml

type Vk uint8

const (
	VkNon      Vk = iota
	VkErrRoll  Vk = iota
	VkPosFail  Vk = iota
	VkErrUndef Vk = iota
	VkA        Vk = iota
	VkB        Vk = iota
	VkC        Vk = iota
	VkD        Vk = iota
	VkE        Vk = iota
	VkF        Vk = iota
	VkG        Vk = iota
	VkH        Vk = iota
	VkI        Vk = iota
	VkJ        Vk = iota
	VkK        Vk = iota
	VkL        Vk = iota
	VkM        Vk = iota
	VkN        Vk = iota
	VkO        Vk = iota
	VkP        Vk = iota
	VkQ        Vk = iota
	VkR        Vk = iota
	VkS        Vk = iota
	VkT        Vk = iota
	VkU        Vk = iota
	VkV        Vk = iota
	VkW        Vk = iota
	VkX        Vk = iota
	VkY        Vk = iota
	VkZ        Vk = iota
	Vk1        Vk = iota
	Vk2        Vk = iota
	Vk3        Vk = iota
	Vk4        Vk = iota
	Vk5        Vk = iota
	Vk6        Vk = iota
	Vk7        Vk = iota
	Vk8        Vk = iota
	Vk9        Vk = iota
	Vk0        Vk = iota
	VkEnter
	VkEsc
	VkDel
	VkTab
	VkSpace
	VkDash      // -
	VkEqual     // =
	VkLBracket  // [
	VkRBracket  // ]
	VkBackSlash // \
	Vk50        //~
	VkSemiCol   // ;
	VkQuota     // '
	Vk53
	VkComma  // ,
	VkPeriod // .
	VkSlash  // /
	VkCap
	VkF1
	VkF2
	VkF3
	VkF4
	VkF5
	VkF6
	VkF7
	VkF8
	VkF9
	VkF10
	VkF11
	VkF12

	VkLCtrl   Vk = 224
	VkLShift  Vk = 225
	VkLAlt    Vk = 226
	VkLeftGui Vk = 227
	VkRCtrl   Vk = 228
	VkRShift  Vk = 229
	VkRAlt    Vk = 230
	VkRGui    Vk = 231
)

type MButton uint8

const (
	MbLeft   MButton = 0
	MbRight  MButton = 1
	MbMiddle MButton = 2
)
