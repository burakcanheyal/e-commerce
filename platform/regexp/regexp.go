package regexp

import (
	"fmt"
	"regexp"
	"strings"
)

var regexpObject *regexp.Regexp

func init() {
	badWords := []string{
		"idiot", "moron", "göt", "fuck", "amcık", "götveren", "moronu", "amınakodumun", "amına kodumun", "ak", "aq",
		"mal", "pussy", "bullshit", "fuckoff", "fuck off", "ananı", "anan", "sikeyim", "sikerim", "ampute", "sefil götveren",
		"şuursuz it", "fuck you", "fuck u", "motherfucker"}

	pattern := fmt.Sprintf(`(?i)\b(?:%s)\b`, strings.Join(badWords, "|"))

	regexpObject = regexp.MustCompile(pattern)
}
func InspectFeedback(text string) string {
	return regexpObject.ReplaceAllString(text, "*")
}
