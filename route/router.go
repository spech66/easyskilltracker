package route

import (
	"errors"
	"io"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spech66/easyskilltracker/api"
	"github.com/spech66/easyskilltracker/view"
)

func configExtender(data string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("data", data)
			return next(c)
		}
	}
}

// TemplateRegistry defines the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render templates by implementing e.Renderer interface => https://blog.boatswain.io/post/setup-nested-html-template-in-go-echo-web-framework/
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "_base.htm", data)
}

// Init initializes echo and routes => https://github.com/eurie-inc/echo-sample
func Init(data string) *echo.Echo {
	// Echo instance
	e := echo.New()
	//e.Debug()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(configExtender(data))

	// View routes => handler
	e.GET("/", view.HomeHandler)

	// Parse all templates
	templates := make(map[string]*template.Template)
	templates["index.htm"] = template.Must(template.ParseFiles("templates/index.htm", "templates/_base.htm"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Api routes => handler
	apiGroup := e.Group("/api")
	{
		apiGroup.GET("/skill", api.GetSkill())
	}

	return e
}
