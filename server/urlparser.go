package server

import (
	"regexp"
	"strings"
)

//проверяет урлу на соответствие шаблону
type IApiUrlMatcher interface {
	//возвращает true если паттерн урлы совпадает
	Match(url string) bool

	//вовращает шаблон, по которому сравниваются урлы
	Pattern() string
}

//парсит апи урлы на подобие /usr/id/1/var/val
//и достает из них значения
type IApiUrlParser interface {
	//Возвращает апи параметры урлы
	Parse(url string) map[string]string

	//вовращает шаблон, по которому сравниваются урлы
	Pattern() string
}

type ApiUrlParser struct {
	//не скомпилированная регулярка
	pattern string

	//скомпилированная регулярка
	regex *regexp.Regexp

	//имена апи параметров урлы
	parameters []string
}

//приводит шаблон url к регулярке, по которой будем искать совпадения
// /user/{id}/description -> \/user\/([^\/]+)\/description
func NewApiUrlParser(pattern string) (*ApiUrlParser, error) {
	p := new(ApiUrlParser)
	p.parameters = make([]string, 0)

	//разбиваем шаблон url на куски
	splittedPattern := strings.Split(pattern, "/")

	for i, val := range splittedPattern {
		len := len(val)
		if len < 2 {
			continue
		}

		if val[0:1] == "{" && val[len-1:len] == "}" {
			p.parameters = append(p.parameters, val[1:len-1])
			splittedPattern[i] = `([^\/]+)`
		}
	}

	//объединяем все в регулярное выражение, которым будем определять совпадения урлы
	p.pattern = strings.Join(splittedPattern, `\/`)
	//компилируем регулярку
	regex, err := regexp.Compile(p.pattern)
	p.regex = regex

	return p, err
}

func (p *ApiUrlParser) Pattern() string {
	return p.pattern
}

func (p *ApiUrlParser) Match(url string) bool {
	return p.regex.MatchString(url)
}

func (p *ApiUrlParser) Parse(url string) map[string]string {
	parsedParams := make(map[string]string)
	apiValues := p.regex.FindStringSubmatch(url)

	//первый элемент групп это всегда изначальная строка
	for i := 1; i < len(apiValues); i++ {
		parsedParams[p.parameters[i-1]] = apiValues[i]
	}

	return parsedParams
}
