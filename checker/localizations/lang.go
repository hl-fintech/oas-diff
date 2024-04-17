package localizations

const (
	LangDefault = LangEn
	LangEn      = "en"
	LangRu      = "ru"
	LangKr      = "kr"
)

func GetSupportedLanguages() []string {
	return []string{LangEn, LangRu, LangKr}
}
