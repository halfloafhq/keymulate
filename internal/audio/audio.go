package audio

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

func LoadAudioCtx() *oto.Context {
	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	context, readyChan, err := oto.NewContext(op)
	if err != nil {
		log.Fatalf("failed to create Oto context: %v", err)
	}
	<-readyChan
	return context
}

func PlayPress(otoCtx *oto.Context, key int, audio io.Reader) {
	player := otoCtx.NewPlayer(audio)
	player.Play()
	err := player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}

func PlayRelease(otoCtx *oto.Context, key int, audio io.Reader) {
	player := otoCtx.NewPlayer(audio)
	player.Play()
	err := player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}

func loadSound(switchType, soundName, action string) []byte {
	soundFile, err := os.Open(fmt.Sprintf("/home/shbhtngpl/personal/halfloafhq/keymulate/audio/%s/%s/%s.mp3", switchType, action, soundName))

	if err != nil {
		log.Fatalf("failed to open sound file: %v", err)
	}
	defer soundFile.Close()

	sound, err := io.ReadAll(soundFile)
	if err != nil {
		log.Fatalf("failed to read sound file: %v", err)
	}

	return sound
}

func LoadSoundsForKeyboard(switchType string) (map[string][]byte, map[string][]byte) {
	pressSounds := make(map[string][]byte)
	releaseSounds := make(map[string][]byte)

	pressKeys := []string{"GENERIC_R0", "GENERIC_R1", "GENERIC_R2", "GENERIC_R3", "GENERIC_R4", "ENTER", "SPACE", "BACKSPACE"}
	releaseKeys := []string{"GENERIC", "ENTER", "BACKSPACE", "SPACE"}

	for _, key := range pressKeys {
		sound := loadSound(switchType, key , "press")
		pressSounds[key] = sound
	}

	for _, key := range releaseKeys {
		sound := loadSound(switchType, key, "release")
		releaseSounds[key] = sound
	}

	return pressSounds, releaseSounds
}

func GetSoundKey(code uint16, isPress bool) string {
	// Map key codes to sound keys. Adjust this mapping as needed.
	switch {
	case code >= 2 && code <= 11: // 1-0 keys
		return fmt.Sprintf("GENERIC_R%d", (code-2)%4)
	case code == 28: // Enter key
		return "ENTER"
	case code == 57: // Space key
		return "SPACE"
	case code == 14: // Backspace key
		return "BACKSPACE"
	default:
		if isPress {
			return "GENERIC_R0" // Default press sound
		}
		return "GENERIC" // Default release sound
	}
}

func PlaySound(otoCtx *oto.Context, sound []byte) {
	reader := bytes.NewReader(sound)
	decodedMp3, err := mp3.NewDecoder(reader)
	if err != nil {
		fmt.Printf("Error decoding MP3: %v\n", err)
		return
	}

	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()

	// Wait for the sound to finish playing
	for player.IsPlaying() {
		// You might want to add a small sleep here to prevent busy-waiting
		// time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		fmt.Printf("Error closing player: %v\n", err)
	}
}
