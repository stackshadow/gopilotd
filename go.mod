module gitlab.com/gopilot/gopilotd

go 1.12

require (
	github.com/chappjc/logrus-prefix v0.0.0-20180227015900-3a1d64819adb
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.4.2
	gitlab.com/gopilot/lib v0.1.4
)

// for local develop
replace gitlab.com/gopilot/lib => ../lib
