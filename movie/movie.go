package movie

type Movie struct {
	name  string
	year  int
	genre string
}

func NewMovie(name string, year int, genre string) Movie {
	return Movie{
		name:  name,
		year:  year,
		genre: genre,
	}
}

func (m *Movie) GetTitle() string {
	return m.name
}

func (m *Movie) SetTitle(name string) {
	m.name = name
}

func (m *Movie) GetYear() int {
	return m.year
}

func (m *Movie) SetYear(year int) {
	m.year = year
}

func (m *Movie) GetGenre() string {
	return m.genre
}

func (m *Movie) SetGenre(genre string) {
	m.genre = genre
}
