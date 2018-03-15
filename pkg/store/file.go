package store

type FileStore struct {
	tmpl *template.Template
	mode os.FileMode
}

func File(pathTmpl string, mode os.FileMode) (s FileStore) {
	s.tmpl = template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(pathTmpl))
	s.mode = mode

	return
}

func (s FileStore) StoreBytes(ctx types.Ctx, b []byte) (c types.Ctx, err error) {
	c = ctx

	var tplOut bytes.Buffer
	s.tmpl.Execute(&tplOut, ctx)

	err = ioutil.WriteFile(tplOut.String(), ctx.Val.Bytes, 0644)
	if err != nil {
		return
	}

	return
}