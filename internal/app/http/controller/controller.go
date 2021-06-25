package controller

type Schema struct {
	*MainCtrl
	*UserCtrl
	*MusicCtrl
	*FavoriteCtrl
}

func New(u *UserCtrl, m *MusicCtrl, f *FavoriteCtrl) *Schema {
	return &Schema{
		new(MainCtrl),
		u,
		m,
		f,
	}
}
