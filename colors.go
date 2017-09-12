package GoMicroColor

type TerminalColorType int

const (
	// 30-37 fg colors, 40-47 bg colors
	Black               TerminalColorType = 30
	Red                                   = 31
	Green                                 = 32
	Yellow                                = 33
	Blue                                  = 34
	Purple                                = 35
	Cyan                                  = 36
	White                                 = 37
	Defult                                = 39
	BackgroundBlack                       = 40
	BackgroundRed                         = 41
	BackgroundGreen                       = 42
	BackgroundYellow                      = 43
	BackgroundBlue                        = 44
	BackgroundPurple                      = 45
	BackgroundCyan                        = 46
	BackgroundLightGray                   = 47
	BackgroundDefault                     = 49
)
