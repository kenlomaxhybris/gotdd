package gotdd

import "errors"

type Review struct {
	Comment string
}

type Register struct {
	lts []LunchTalk
}

type LunchTalk struct {
	Title   string
	Speaker string
	Reviews []Review
}

func (r *Register) AddLunchTalk(lt LunchTalk) error {
	if len(lt.Title) == 0 || len(lt.Speaker) == 0 {
		return errors.New("Missing Data")
	}
	r.lts = append(r.lts, lt)
	return nil
}

func (r *Register) GetLunchTalks() []LunchTalk {
	return r.lts
}

func (r *Register) AddReview(i int, rev Review) error {
	if len(rev.Comment) == 0 {
		return errors.New("Missing Data")
	}
	if i > len(r.lts)-1 {
		return errors.New("Out of bounds")
	}
	r.lts[i].Reviews = append(r.lts[i].Reviews, rev)
	return nil
}

func (r *Register) AdjustReview(i int, j int, rev Review) error {
	if len(rev.Comment) == 0 {
		return errors.New("Missing Data")
	}
	if i > len(r.lts)-1 || j > len(r.lts[i].Reviews)-1 {
		return errors.New("Out of bounds")
	}

	r.lts[i].Reviews[j] = rev
	return nil
}
