package factory

type mes struct {
	c string
	pwd string
}

func NewMes() *mes{
	return &mes{}
}

func (m *mes) SetPwd(p string){
	m.pwd = p
}