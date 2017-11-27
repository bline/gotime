package main

import (
	"log"
	"io/ioutil"
	"regexp"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/Joker/jade"
	"path/filepath"
)

func (c *gin.Context) Pug(code int, name string, obj interface{}) {

}

func (engine *gin.Engine) LoadPugFiles(files []string) error {
	var fullHtml string = ""
	for _, f := range files {
		htmlSrc, err = jade.ParseFile(f)
		if err != nil {
			log.Print("parsing file ", f, "; error: ", err)
		} else {
			fullHtml += htmlSrc
		}
	}
	templ, err := template.Must(template.New("").Delims(engine.delims.Left, engine.delims.Right).Funcs(engine.FuncMap).Parse(fullHtml)
	if err != nil {
		return err
	} else {
		engine.SetHTMLTemplate(templ)
	}
	return nil
}

func (engine *gin.Engine) LoadPugGlob(glob string) error {
	var (
		matches []string
		err error
	)
	matches, err = filepath.Glob(glob)
	if err != nil {
		return err
	}
	return engine.LoadPugFiles(matches)
}

func setupRoutes(c *gin.Context, r *gin.Router) {

	r.GET("/", func (c *gin.Context) {
		
	})
}
