package taskrunner

import (
	"log"
	"os"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}

	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbpos.Re
}