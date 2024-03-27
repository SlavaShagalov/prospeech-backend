package convert

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

func MP4ToWAV(mp4Data []byte) (wavData []byte, err error) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "wav", "pipe:1")

	cmd.Stdin = bytes.NewReader(mp4Data)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	if out.Len() == 0 {
		return nil, errors.New("ffmpeg output is empty")
	}

	return out.Bytes(), nil
}
