package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ngaut/log"
	"golang.org/x/text/language"
)

var i18nMap = make(map[string](map[string]string))
var defaultLocaleFolder = "locales"
var i18nMatcher language.Matcher

var (
	DefaultLang = []language.Tag{language.Chinese}
)

func String(tags []language.Tag, key string) (text string) {
	tag, _, _ := i18nMatcher.Match(tags...)
	text, _ = i18nMap[tag.String()][key]
	return
}

func Stringf(tags []language.Tag, key string, args ...interface{}) (text string) {
	f := String(tags, key)
	text = fmt.Sprintf(f, args...)
	return
}

func queryLocale(lang string, key string) (text string, found bool) {
	// full matching
	var localeTexts map[string]string
	if localeTexts, found = i18nMap[lang]; found {
		text, found = localeTexts[key]
		return
	}

	//cn mathcing
	if split := strings.Index(lang, "-"); split > 0 {
		lang = lang[0:split]
		if localeTexts, found = i18nMap[lang]; found {
			text, found = localeTexts[key]
		}
	}
	return
}

func LoadLocales(basePath string, folder string, defaultLocale string) (err error) {
	//load default tag
	var tags []language.Tag
	defaultTag, err := language.Parse(defaultLocale)
	if err != nil {
		return err
	} else {
		tags = append(tags, defaultTag)
	}

	fullPath := filepath.Join(basePath, folder)
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(fullPath); err != nil {
		return
	}
	for _, f := range files {
		localeFile := filepath.Join(fullPath, f.Name())
		locale := strings.Split(f.Name(), ".")[0]

		//parse tag
		tag, _err := language.Parse(locale)
		if _err != nil {
			log.Warn("Ingore %s, err: %s", locale, _err)
			continue
		}

		//load locale strings
		var reader *os.File
		reader, err = os.Open(localeFile)
		if err != nil {
			return
		}
		decoder := json.NewDecoder(reader)
		var localeTexts map[string]string
		if err = decoder.Decode(&localeTexts); err != nil {
			return
		}

		i18nMap[locale] = localeTexts

		//default tag has added
		if tag != defaultTag {
			tags = append(tags, tag)
		}
	}

	i18nMatcher = language.NewMatcher(tags)
	return
}
