package main

import (
	"html/template"
	"path"
)

type ViewType string

const (
	LoginView                ViewType = "login"
	CompletePersonalDataView ViewType = "complete-personal-data"
	InternDashboardView      ViewType = "intern-dashboard"
	AdminDashboardView       ViewType = "admin-dashboard"
	ErrorView                ViewType = "error"
	SuccessView              ViewType = "success"
)

func (v ViewType) String() string {
	return string(v)
}

var views = map[ViewType]*template.Template{}

func init() {
	views = map[ViewType]*template.Template{
		LoginView:                template.Must(template.ParseFS(Resources, path.Join("views", "login.html"))),
		CompletePersonalDataView: template.Must(template.ParseFS(Resources, path.Join("views", "complete-personal-data.html"))),
		InternDashboardView:      template.Must(template.ParseFS(Resources, path.Join("views", "intern-dashboard.html"))),
		AdminDashboardView:       template.Must(template.ParseFS(Resources, path.Join("views", "admin-dashboard.html"))),
		ErrorView:                template.Must(template.ParseFS(Resources, path.Join("views", "error.html"))),
		SuccessView:              template.Must(template.ParseFS(Resources, path.Join("views", "success.html"))),
	}
}
