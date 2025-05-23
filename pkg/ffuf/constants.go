package ffuf

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

var (
	//VERSION holds the current version number
	VERSION = "2.1.0"
	//VERSION_APPENDIX holds additional version definition
	VERSION_APPENDIX = "-uff-dev"
	CONFIGDIR        = filepath.Join(xdg.ConfigHome, "ffuf")
	HISTORYDIR       = filepath.Join(CONFIGDIR, "history")
	SCRAPERDIR       = filepath.Join(CONFIGDIR, "scraper")
	AUTOCALIBDIR     = filepath.Join(CONFIGDIR, "autocalibration")
)
