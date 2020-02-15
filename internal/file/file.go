package file

import (
	"bufio"
	"fmt"
	"github-commit-reput/internal/config"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func WriteRepo(message, id string) error {
	currentDate := time.Now().Format("2006-01-02")
	path := fmt.Sprintf("%v/%v/%v", config.RepoPath, config.GitUsername, currentDate)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Error().Err(err).Msgf("Error creating folder %v", path)
		return err
	}

	filePath := fmt.Sprintf("%v/%v.log", path, id)
	f, err := os.Create(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("Error creating file %v", filePath)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	size, err := w.WriteString(message)
	if err != nil {
		log.Error().Err(err).Msgf("Error writing content %v in %v", message, filePath)
		return err
	}
	err = w.Flush()
	if err != nil {
		log.Error().Err(err).Msgf("Error writing content %v in %v", message, filePath)
		return err
	}
	log.Debug().Msgf("Wrote %d bytes in %v", size, filePath)

	return nil
}
