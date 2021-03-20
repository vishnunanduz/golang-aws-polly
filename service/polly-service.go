package service

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

const (
	AUDIO_FORMAT = "mp3"
	JOEY_VOICE   = "Joey"
)

type PollyService interface {
	SynthesizeText(text, fileName string) error
}

type pollyConfig struct {
	voice string
}

func NewJoeyPollyService() PollyService {
	return &pollyConfig{
		voice: JOEY_VOICE,
	}
}

func CreatePollyClient() *polly.Polly {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return polly.New(sess)
}

func (p *pollyConfig) SynthesizeText(text, fileName string) error {
	pollyClient := CreatePollyClient()

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String(AUDIO_FORMAT),
		Text:         aws.String(text),
		VoiceId:      aws.String(p.voice),
	}

	output, err := pollyClient.SynthesizeSpeech(input)
	if err != nil {
		return err
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		return err
	}
	return nil
}
