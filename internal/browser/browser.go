package browser

import (
	"os"
	"strings"

	"hack-browser-data/internal/browingdata"
	"hack-browser-data/internal/browser/chromium"
)

type Browser interface {
	GetName() string

	GetMasterKey() ([]byte, error)

	GetBrowsingData() []browingdata.Source

	CopyItemFileToLocal() error
}

var (
	// home dir path for all platforms
	homeDir, _ = os.UserHomeDir()
)

func PickBrowser(name string) []Browser {
	var browsers []Browser
	clist := pickChromium(name)
	for _, b := range clist {
		if b != nil {
			browsers = append(browsers, b)
		}
	}
	flist := pickFirefox(name)
	for _, b := range flist {
		if b != nil {
			browsers = append(browsers, b)
		}
	}
	return browsers
}

func pickChromium(name string) []Browser {
	var browsers []Browser
	name = strings.ToLower(name)
	if name == "all" {
		for _, c := range chromiumList {
			if b, err := chromium.New(c.name, c.profilePath, c.storage, c.items); err == nil {
				browsers = append(browsers, b)
			} else {
				if strings.Contains(err.Error(), "profile path is not exist") {
					continue
				}
				panic(err)
			}
		}
		return browsers
	}
	if choice, ok := chromiumList[name]; ok {
		b, err := newChromium(choice.browserInfo, choice.items)
		if err != nil {
			panic(err)
		}
		browsers = append(browsers, b)
		return browsers
	}
	return nil
}

func pickFirefox(name string) []Browser {
	var browsers []Browser
	name = strings.ToLower(name)
	if name == "all" || name == "firefox" {
		for _, f := range firefoxList {
			multiFirefox, err := newMultiFirefox(f.browserInfo, f.items)
			if err != nil {
				panic(err)
			}
			for _, browser := range multiFirefox {
				browsers = append(browsers, browser)
			}
		}
		return browsers
	}
	return nil
}

func ListBrowser() []string {
	var l []string
	for c := range chromiumList {
		l = append(l, c)
	}
	for f := range firefoxList {
		l = append(l, f)
	}
	return l
}

type browserInfo struct {
	masterKey []byte
}

const (
	chromeName         = "Chrome"
	chromeBetaName     = "Chrome Beta"
	chromiumName       = "ChromiumBookmark"
	edgeName           = "Microsoft Edge"
	firefoxName        = "FirefoxBookmark"
	firefoxBetaName    = "FirefoxBookmark Beta"
	firefoxDevName     = "FirefoxBookmark Dev"
	firefoxNightlyName = "FirefoxBookmark Nightly"
	firefoxESRName     = "FirefoxBookmark ESR"
	speed360Name       = "360speed"
	qqBrowserName      = "QQ"
	braveName          = "Brave"
	operaName          = "Opera"
	operaGXName        = "OperaGX"
	vivaldiName        = "Vivaldi"
	coccocName         = "CocCoc"
	yandexName         = "Yandex"
)