package models

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Outline struct {
	Id      int
	Outline string `orm:"type(text)"`
}

func GetOutline(id int) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Id: id}
	err := o.Read(&outline)
	return outline, err
}

func CreateOutline(outline_str string) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Outline: outline_str}
	_, err := o.Insert(&outline)
	return outline, err
}

func UpdateOutline(id int, outline string) (Outline, error) {
	o := orm.NewOrm()
	_outline := Outline{Id: id}
	err := o.Read(&_outline)
	if err != nil {
		return Outline{}, err
	}
	_outline.Outline = outline
	_, err = o.Update(&_outline)
	return _outline, err
}

// 将outline中的"\n"字符串删去
func RefactOutline(outline Outline) Outline {
	outline.Outline = DeleteLineBreak(outline.Outline)
	return outline
}

func DeleteLineBreak(outline string) string {
	outline = strings.ReplaceAll(outline, "\n", "")
	return outline
}

func RefactXML(xmlStr string) string {
	xmlStr = strings.ReplaceAll(xmlStr, "\n", "")
	xmlStr = strings.ReplaceAll(xmlStr, "\t", "")
	regex := "(<.*>)"
	re := regexp.MustCompile(regex)
	matches := re.FindAllString(xmlStr, -1)
	return matches[0]
}

type P struct {
	XMLName xml.Name `xml:"p"`
	Content string   `xml:",innerxml"`
}

type Slide struct {
	XMLName xml.Name `xml:"section"`
	Class   string   `xml:"class,attr"`
	P_arr   []P      `xml:"p"`
	Content string   `xml:",innerxml"`
}

type Slides struct {
	XMLName  xml.Name `xml:"slides"`
	Sections []Slide  `xml:"section"`
}

func GetContentSections(xmlStr string) ([]string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides)
	if err != nil {
		return []string{}, err
	}

	// 查询<section class='content'>的项
	var contents []string
	for _, section := range slides.Sections {
		if section.Class == "content" {
			contents = append(contents, "<section class='content'>"+section.Content+"</section>")
		}
	}
	return contents, nil
}

func RefactContentSections(xmlStr string, guide_slides []string) (string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides)
	if err != nil {
		return "", err
	}

	// 查询<section class='content'>的项
	var resultxml string
	//初始化count=0
	count := 0
	for _, section := range slides.Sections {
		if section.Class == "content" {
			resultxml += guide_slides[count]
			count++
		} else {
			resultxml += "<section class='" + section.Class + "'>" + section.Content + "</section>"
		}
	}

	resultxml = "<slides>" + resultxml + "</slides>"
	return resultxml, nil
}

func coverReplace(section Slide, cover string) string {
	var title = section.P_arr[0].Content
	var description = section.P_arr[1].Content

	cover = strings.ReplaceAll(cover, "{{title}}", title)
	cover = strings.ReplaceAll(cover, "{{description}}", description)

	return cover
}

func catalogReplace(section Slide, template Template) string {
	var p_num = len(section.P_arr)
	var p_list = section.P_arr
	var catalog = p_list[0].Content
	var ret = ""

	if p_num-1 == 3 {
		ret = template.Catalog_3
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
	} else if p_num-1 == 4 {
		ret = template.Catalog_4
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
		ret = strings.ReplaceAll(ret, "{{item4}}", p_list[4].Content)
	} else if p_num-1 == 5 {
		ret = template.Catalog_5
		ret = strings.ReplaceAll(ret, "{{item1}}", p_list[1].Content)
		ret = strings.ReplaceAll(ret, "{{item2}}", p_list[2].Content)
		ret = strings.ReplaceAll(ret, "{{item3}}", p_list[3].Content)
		ret = strings.ReplaceAll(ret, "{{item4}}", p_list[4].Content)
		ret = strings.ReplaceAll(ret, "{{item5}}", p_list[5].Content)
	}
	ret = strings.ReplaceAll(ret, "{{catalog}}", catalog)
	return ret

}

func contentReplace(section Slide, template Template) []string {
	var p_num = len(section.P_arr)
	var p_list = section.P_arr
	var sub_title = p_list[0].Content
	var retList []string
	// 如果p_num-1<=4，只用生成一个页面
	var index = 1
	var ret string
	ret = template.Transition
	ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
	retList = append(retList, ret)

	for p_num-1 > 0 {
		if p_num-1 <= 4 {
			var ret string
			if p_num-1 == 1 {
				ret = template.Content_1
			} else if p_num-1 == 2 {
				ret = template.Content_2
			} else if p_num-1 == 3 {
				ret = template.Content_3
			} else if p_num-1 == 4 {
				ret = template.Content_4
			}
			ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
			for i := 1; i < p_num; i++ {
				ret = strings.ReplaceAll(ret, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret)
			index += p_num - 1
			p_num = 1
		} else if p_num-1 == 5 {
			var ret1, ret2 string
			ret1 = template.Content_3
			ret2 = template.Content_2
			ret1 = strings.ReplaceAll(ret1, "{{sub_title}}", sub_title)
			ret2 = strings.ReplaceAll(ret2, "{{sub_title}}", sub_title)
			for i := 1; i < 4; i++ {
				ret1 = strings.ReplaceAll(ret1, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			for i := 4; i < 6; i++ {
				ret2 = strings.ReplaceAll(ret2, "{{sub_title_content"+strconv.Itoa(i-3)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret1)
			retList = append(retList, ret2)
			p_num = 1
		} else {
			// 生成一个页面 4
			var ret string
			ret = template.Content_4
			ret = strings.ReplaceAll(ret, "{{sub_title}}", sub_title)
			for i := 1; i < 5; i++ {
				ret = strings.ReplaceAll(ret, "{{sub_title_content"+strconv.Itoa(i)+"}}", p_list[index].Content)
				index++
			}
			retList = append(retList, ret)

			p_num -= 4
		}
	}
	return retList
}

func GenPPT(xmlStr string, template Template) ([]string, error) {
	var slides Slides
	err := xml.Unmarshal([]byte(xmlStr), &slides)
	if err != nil {
		return []string{}, err
	}
	var ppt []string
	res := ""

	for i, section := range slides.Sections {
		if i == 0 {
			res = coverReplace(section, template.Cover)
			ppt = append(ppt, res)
		} else if i == 1 {
			res = catalogReplace(section, template)
			ppt = append(ppt, res)
		} else {
			res_list := contentReplace(section, template)
			for _, res_item := range res_list {
				ppt = append(ppt, res_item)
			}
		}

	}
	ppt = append(ppt, template.Thank)

	return ppt, nil

}
