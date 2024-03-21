package common

import (
	"strings"

	"github.com/russross/blackfriday"
)

/**
 *
 * @author  镜湖老杨
 * @date  2021/1/14 11:00 上午
 * @version 1.0
 */

type fileUtil struct{}

func (fileUtil) GetFileExtendsByType(fileType int) []string {
	var fileExtends []string
	switch fileType {
	case 1:
		fileExtends = append(fileExtends, "bmp", "jpg", "png", "tif", "gif", "jpeg", "webp")
		break
	case 2:
		fileExtends = append(fileExtends, "doc", "docx", "txt", "hlp", "wps", "rtf", "html", "pdf", "md", "sql", "css", "js", "vue", "java")
		break
	case 3:
		fileExtends = append(fileExtends, "avi", "mp4", "mpg", "mov", "swf")
		break
	case 4:
		fileExtends = append(fileExtends, "wav", "aif", "au", "mp3", "ram", "wma", "mmf", "amr", "aac", "flac")
		break
	case 5:
		fileExtends = append(fileExtends, "bmp", "jpg", "png", "tif", "gif", "jpeg", "webp",
			"doc", "docx", "txt", "hlp", "wps", "rtf", "html", "pdf", "md", "sql", "css", "js", "vue", "java",
			"avi", "mp4", "mpg", "mov", "swf",
			"wav", "aif", "au", "mp3", "ram", "wma", "mmf", "amr", "aac", "flac")
		break
	default:
		fileExtends = []string{}
		break
	}
	return fileExtends
}

func (fileUtil) MarkdownToHTML(md string) string {
	myHTMLFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	myExtensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_HARD_LINE_BREAK |
		blackfriday.EXTENSION_FOOTNOTES |
		blackfriday.EXTENSION_HEADER_IDS

	renderer := blackfriday.HtmlRenderer(myHTMLFlags, "", "")
	bytes := blackfriday.MarkdownOptions([]byte(md), renderer, blackfriday.Options{
		Extensions: myExtensions,
	})
	theHTML := string(bytes)

	// p := bluemonday.UGCPolicy()
	// p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")

	//裁剪后数学公式无法渲染， <br /> 会裁剪为<br/>
	// return p.Sanitize(theHTML)
	return theHTML
}

func (fileUtil) GetFileName(oraginalFilename string) string {
	var fileName string
	if oraginalFilename != "" && strings.Contains(oraginalFilename, ".") {
		s := strings.Split(oraginalFilename, ".")
		fileName = s[0]
	}
	return fileName
}

var FileUtil = &fileUtil{}
