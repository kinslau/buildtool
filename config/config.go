package config

type Log struct {
	Level        string
	Formatter    string
	ReportCaller bool
}

// LogSetting = &Log{}

// func loadLog() {
// 	sec, err := Cfg.GetSection("log")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'log': %v", err)
// 	}

// 	err = sec.MapTo(LogSetting)
// 	if err != nil {
// 		log.Fatalf("Cfg.MapTo LogSetting err: %v", err)
// 	}

// }
