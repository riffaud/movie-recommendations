package movie

import "time"

// create movie filter func type
type movieFilter func(movie Movie) (isIn bool)

func isShowingAfter(st time.Time) movieFilter {
	st = st.Add(-time.Minute * 30)
	return func(m Movie) (isIn bool) {
		for _, ms := range m.Showings {
			if ms.After(st) {
				return true
			}
		}
		return false
	}
}

func isGenre(sg string) movieFilter {
	return func(movie Movie) (isIn bool) {
		for _, g := range movie.Genres {
			if g == sg {
				return true
			}
		}
		return false
	}
}

func (movies Movies) filter(filters []movieFilter) (ret Movies) {
	for _, m := range movies {
		isIn := true
		for i := 0; isIn && i < len(filters); i++ {
			isIn = isIn && filters[i](m)
		}
		if isIn {
			ret = append(ret, m)
		}
	}
	return ret
}
