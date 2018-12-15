package comm_builder

import (
	"github.com/dimiro1/banner"
	"strings"
	"fmt"
	"flag"
	"os"
	"github.com/mattn/go-colorable"
)

func Banner_show(program_version, compiler_version, build_time, author string) {
	banner_logo :=
		`***********************************************************
***********************************************************
 ____  ____  __ _   ___  ____  _  _   __   _  _   __    __          ___   __   _  _ 
(__  )(  __)(  ( \ / __)(__  )/ )( \ (  ) / )( \ / _\  (  )        / __) /  \ ( \/ )
 / _/  ) _) /    /( (_ \ / _/ ) __ (  )(  ) __ (/    \  )(    _   ( (__ (  O )/ \/ \
(____)(____)\_)__) \___/(____)\_)(_/ (__) \_)(_/\_/\_/ (__)  (_)   \___) \__/ \_)(_/

***********************************************************
****************** Compile Environment ********************
Program version : %s
Compiler version : %s
Build time : %s
Author : %s
***********************************************************
****************** Running Environment ********************
Go running version : {{ .GoVersion }}
Go running OS : {{ .GOOS }}
Startup time : {{ .Now "2006-01-02 15:04:05" }}
***********************************************************
`
	var version bool
	flag.BoolVar(&version, "v", false, "print the version info")
	flag.Parse()

	new_banner := fmt.Sprintf(banner_logo, program_version, compiler_version, build_time, author)

	banner.Init(colorable.NewColorableStdout(), true, true, strings.NewReader(new_banner))

	if version {
		os.Exit(0)
	}
}
