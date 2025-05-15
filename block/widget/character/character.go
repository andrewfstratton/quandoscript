package character

const (
	FIXED_SPACE         = "&nbsp;"
	DOUBLE_QUOTE        = "&quot;"
	VALUE               = "⚡"
	TEXT                = "❓"
	ON_OFF              = "❗"
	NO_tEXT             = "❔"
	PLUS_MINUS          = "±"
	INVENTOR_TEXT       = "📄"
	IMAGE               = "📷"
	VIDEO               = "📺"
	LEAP_MOTION         = "👋"
	HAND_WAVE           = "🖐"
	HAND_POINT_LEFT     = "👈"
	HAND_POINT_RIGHT    = "👉"
	ROBOT_HEAD          = "🤖"
	COLOUR              = "🌈"
	OBJECT              = "🎁"
	µBIT                = "📟"
	A                   = "🅰"
	B                   = "🅱"
	MESSAGE             = "📡"
	TIME                = "⏰"
	AUDIO               = "🔊"
	BRAIN               = "🧠"
	ZOOM                = "🔍"
	MOUSE               = "🖱️"
	KEYBOARD            = "⌨️"
	JOYSTICK            = "🕹️️️️️"
	STICK               = "📍️️"
	GAMEPAD             = "🎮️️"
	SETTINGS_COG        = "️⚙️"
	FOOTPRINTS          = "👣️️"
	FLOPPY_DISK         = "💾"
	MUSIC_KEYBOARD      = "🎹"
	NOTE                = "🎵"
	DRUM                = "🥁"
	ARROW_LEFT_RIGHT    = "⇔"
	ARROW_UP_DOWN       = "⇕"
	ARROW_IN_OUT        = "⤢"
	ARROW_ROTATE        = "⤹⤸"
	ARROW_UP            = "⇧"
	ARROW_DOWN          = "⇩"
	ARROW_LEFT          = "⇦"
	ARROW_RIGHT         = "⇨"
	ARROW_IN            = "⬀"
	ARROW_OUT           = "⬃"
	ARROW_UP_THEN_LEFT  = "↰"
	ARROW_UP_THEN_RIGHT = "↱"
)

type Character struct {
	txt string
}

func New(t string) *Character {
	return &Character{txt: t}
}

func (ch *Character) Html() string {
	return ch.txt
}
