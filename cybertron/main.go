package main

import (
	"fmt"
	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/questionanswering"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"log"
	"os"
	"time"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	LoadDotenv()

	modelsDir := HasEnvVar("CYBERTRON_MODELS_DIR")
	modelName := HasEnvVar("CYBERTRON_MODEL")
	paragraph := HasEnvVar("CYBERTRON_QA_PARAGRAPH")

	m, err := tasks.Load[questionanswering.Interface](&tasks.Config{ModelsDir: modelsDir, ModelName: modelName})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer tasks.Finalize(m)

	opts := &questionanswering.Options{}

	fn := func(text string) error {
		start := time.Now()
		result, err := m.Answer(context.Background(), text, paragraph, opts)
		if err != nil {
			return err
		}
		fmt.Println(time.Since(start).Seconds())
		fmt.Println(MarshalJSON(result))
		return nil
	}

	fmt.Println(paragraph)

	err = ForEachInput(os.Stdin, fn)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
