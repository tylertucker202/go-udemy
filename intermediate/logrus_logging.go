package intermediate

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message using logrus.")
	log.Warn("This is a warning message using logrus.")
	log.Error("This is an error message using logrus.")
	log.Debug("This is a debug message using logrus.")

	log.WithFields(logrus.Fields{
		"username": "johndoe",
		"action":   "login",
	}).Info("User action logged.")

}
