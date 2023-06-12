package cups

import (
	"log"
	"os"
	"strconv"
)

type Backend struct {
	device_uri string
	job        *Job
}

func (b *Backend) InitParameters(args []string) (err error) {
	log.Println("initializing parameters for backend..")
	job := Job{}
	job.id, err = strconv.Atoi(args[1])
	job.user = args[2]
	job.title = args[3]
	job.copies, err = strconv.Atoi(args[4])
	job.options = args[5]
	if len(args) == 7 {
		job.file = args[6]
	} else {
		job.file = ""
	}

	b.device_uri = os.Getenv("DEVICE_URI")

	log.Printf("DeviceURI: %s\n", b.device_uri)
	log.Printf("job: %s\n", job.String())

	return err
}
