package template

type Option func(tpl *Template)

func Files(files ...string) Option {
	return func(tpl *Template) {
		tpl.Files = files
	}
}
