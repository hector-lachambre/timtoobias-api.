module gitlab.com/timtoobias-projects/timtoobias-api

go 1.15

replace gitlab.com/timtoobias-projects/timtoobias-core => ../timtoobias-core
replace gitlab.com/timtoobias-projects/timtoobias-datas => ../timtoobias-datas


require (
	gitlab.com/timtoobias-projects/timtoobias-core v0.0.0
	gitlab.com/timtoobias-projects/timtoobias-datas v0.0.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.62.0
)
