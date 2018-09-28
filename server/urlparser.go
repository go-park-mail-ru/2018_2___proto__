package server

import (
	"regexp"
	"strings"
)

//выпарсивает апи параметры урлы
type IApiUrlParser interface {
	//возвращает true если паттерн урлы совпадает
	Match(url string) bool

	//Возвращает апи параметры урлы 
	Parse(url string) map[string]string

	//вовращает шаблон, по которому сравниваются урлы
	Pattern() string
}

type ApiUrlParser struct {
	//не скомпилированная регулярка
	patern string

	//скомпилированная регулярка
	regex *regexp.Regexp

	//имена апи параметров урлы
	parametrs []string
}

//приводит шаблон url к регулярке, по которой будем искать совпадения
// /user/{id}/description -> \/user\/([^\/]+)\/description
func NewApiUrlParser(pattern string) (*ApiUrlParser, error) {
	p := new(ApiUrlParser)
	p.parametrs = make([]string, 0)

	//разбиваем шаблон url на куски
	splitedPattern := strings.Split(pattern, "/")

	for i, val := range splitedPattern {
		len := len(val)
		if len < 2 {
			continue
		}

		if val[0:1] == "{" && val[len-1:len] == "}" {
			p.parametrs = append(p.parametrs, val[1:len-1])
			splitedPattern[i] = `([^\/]+)`
		}
	}

	//объединяем все в регулярное выражение, которым будем определять совпадения урлы
	p.patern = strings.Join(splitedPattern, `\/`)
	//компилируем регулярку
	regex, err := regexp.Compile(p.patern)
	p.regex = regex

	return p, err
}

func (p *ApiUrlParser) Pattern() string {
	return p.patern
}

func (p *ApiUrlParser) Match(url string) bool {
	return p.regex.MatchString(url)
}

func (p *ApiUrlParser) Parse(url string) map[string]string {
	parsedParams := make(map[string]string)
	apiValues := p.regex.FindStringSubmatch(url)

	//первый элемент групп это всегда изначальная строка
	for i:=1; i<len(apiValues); i++ {
		parsedParams[p.parametrs[i-1]] = apiValues[i]
	}

	return parsedParams
}