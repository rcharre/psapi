package ps

import (
	"encoding/json"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/pkg/utils/importer"
	"iter"
	"log/slog"
	"path"
)

const (
	TypeFolder                     = "types"
	PokemonTypeTranslationFileName = "100003.csv"
)

var _ TypeImporter = &TypeImporterImpl{}

type TypeImporter interface {
	Import(studioFolder string, translationFolder string) (iter.Seq[*PokemonType], error)
}

type TypeImporterImpl struct {
}

func NewTypeImporter() *TypeImporterImpl {
	return &TypeImporterImpl{}
}

// Import import a types folder as an iterator
// studioFolder The studio folder path
// translationFolder The translation folder path
func (i TypeImporterImpl) Import(studioFolder string, translationFolder string) (iter.Seq[*PokemonType], error) {
	slog.Info("Import translation file for types")
	typeTranslationFilePath := path.Join(translationFolder, PokemonTypeTranslationFileName)
	typeTranslations, err := i18n.ImportTranslations(typeTranslationFilePath)
	if err != nil {
		return nil, err
	}

	typeFolderPath := path.Join(studioFolder, TypeFolder)
	slog.Info("Importing pokemon types folder", "path", typeFolderPath)

	typeContentIterator, err := importer.ImportFolder(typeFolderPath)
	if err != nil {
		return nil, err
	}

	return func(yield func(*PokemonType) bool) {
		for typeContent := range typeContentIterator {
			pokemonType := &PokemonType{}
			if err := json.Unmarshal(typeContent, pokemonType); err != nil {
				slog.Warn("Failed to unmarshal pokemon type", "error", err)
				continue
			}

			if pokemonType.TextId < len(typeTranslations) {
				pokemonType.Name = typeTranslations[pokemonType.TextId]
			}
			if !yield(pokemonType) {
				break
			}
		}
	}, nil

}
