// Package localizer is used to localize for different languages
package localizer

import (
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/message"

	// import the package translations so that it's init() function is run.
	// it ensures default message catalog is updated to use our translations
	// before we initialize the message.Printer instances below.
	_ "github.com/fireyun/i18n-go/pkg/i18n/translations"
)

// supportedLangs are languages the server supports.
var supportedLangs = []language.Tag{
	language.English, // en fallback
	language.Chinese, // zh
}

// matcher is used to get the most matched language.
var matcher = language.NewMatcher(supportedLangs)

// Localizer is message.Printer instance for the locale.
type Localizer struct {
	printer *message.Printer
}

// localizerMap holds the initialized Localizer for supported language.
var localizerMap = map[string]*Localizer{
	// English
	display.English.Tags().Name(language.English): {
		printer: message.NewPrinter(language.English),
	},
	// Chinese
	display.English.Tags().Name(language.Chinese): {
		printer: message.NewPrinter(language.Chinese),
	},
}

// Get gets the matched Localizer for the language.
// it will get the most matched language, see more: https://go.dev/blog/matchlang
// if the language is not supported by the server, then fall back to English.
// eg: language tags like zh-CN, zh-TW, zh-HK, cmn will use Chinese,
// other language tags like en-US, en-GB, nl or unknown language tag will use English.
func Get(lang string) *Localizer {
	tag, _, _ := matcher.Match(language.Make(lang))
	return localizerMap[display.English.Tags().Name(tag)]
}

// Translate acts as a wrapper to call message.Printer's Sprintf method.
// it returns the appropriate translation for the given message and language.
// it is concurrency safe.
func (l *Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
