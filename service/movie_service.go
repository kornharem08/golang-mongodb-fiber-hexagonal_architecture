package service

import (
	"basic/errs"
	"basic/logs"
	"basic/repository"
)

type movieService struct {
	movRepo repository.MovieRepository
}

func NewMovieService(movRepo repository.MovieRepository) movieService {
	return movieService{movRepo: movRepo}
}

func (s movieService) GetMovies(page, limit int) ([]MovieResponse, error) {
	movies, err := s.movRepo.GetAll(page, limit)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpected error")
	}

	responses := []MovieResponse{}
	for _, movie := range movies {
		response := MovieResponse{
			ID:    movie.ID,
			Title: movie.Title,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
