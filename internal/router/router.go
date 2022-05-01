package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/flamego/flamego"
	"github.com/flamego/template"
	"github.com/qiniu/qmgo/operator"

	"github.com/midoks/vez-en/internal/mgdb"
	"github.com/midoks/vez-en/internal/robot"
)

const PAGE_NUM = 15

func Hello() string {
	return fmt.Sprintf("%s", "Hello world")
}

func Home(t template.Template, data template.Data) {

	d, _ := mgdb.ContentOriginFindId("", "-", PAGE_NUM)
	data["Articles"] = d

	dLen := len(d)
	if dLen > 1 {
		data["NextPos"] = d[dLen-1].MgID
	}

	t.HTML(http.StatusOK, "home")
}

func So(c flamego.Context, t template.Template, data template.Data) {
	kw := c.Param("kw")
	prevNext := c.Param("prevNext")
	id := c.Param("pos")

	op := operator.Lt
	if strings.EqualFold(prevNext, "prev") {
		op = operator.Gt
	}

	d, _ := mgdb.ContentOriginFindSoso(id, "-", op, kw, PAGE_NUM)

	data["Articles"] = d
	data["Keyword"] = kw

	dLen := len(d)

	if strings.EqualFold(prevNext, "prev") && dLen != PAGE_NUM {
		c.Redirect("/so/" + kw + ".html")
	}

	if strings.EqualFold(prevNext, "next") && dLen == 0 {
		c.Redirect("/so/" + kw + ".html")
	}

	if dLen > 1 {
		data["PrePos"] = d[0].MgID
		if dLen == PAGE_NUM {
			data["NextPos"] = d[dLen-1].MgID
		}
	}

	t.HTML(http.StatusOK, "soso")
}

func Prev(c flamego.Context, t template.Template, data template.Data) {
	id := c.Param("pos")
	d, _ := mgdb.ContentOriginFindIdGt(id, "-", PAGE_NUM)

	data["Articles"] = d

	dLen := len(d)
	if dLen != PAGE_NUM {
		c.Redirect("/")
		return
	}

	if dLen > 1 {
		data["PrePos"] = d[0].MgID
		data["NextPos"] = d[dLen-1].MgID
	}

	t.HTML(http.StatusOK, "home")
}

func Next(c flamego.Context, t template.Template, data template.Data) {
	id := c.Param("pos")

	d, _ := mgdb.ContentOriginFindId(id, "-", PAGE_NUM)
	data["Articles"] = d

	dLen := len(d)
	if dLen > 1 {
		data["PrePos"] = d[0].MgID
		if dLen == PAGE_NUM {
			data["NextPos"] = d[dLen-1].MgID
		}
	}

	t.HTML(http.StatusOK, "home")
}

func Rand(c flamego.Context, t template.Template, data template.Data) {
	d, err := mgdb.ContentRand()

	if err != nil {
		t.HTML(http.StatusOK, "home")
		return
	}

	url := "/html/" + d.MgID + ".html"
	c.Redirect(url)
}

func About(c flamego.Context, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "about")
}

func ContentHtml(c flamego.Context, t template.Template, data template.Data) {
	d, _ := mgdb.ContentOriginFindOne(robot.DZONE_NAME, c.Param("pos"))
	data["Article"] = d
	t.HTML(http.StatusOK, "page/content")
}
