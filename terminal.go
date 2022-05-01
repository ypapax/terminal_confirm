package terminal_confirm

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

func Ask4confirm(msg string) (bool, error) {
	var s string

	fmt.Printf(msg + " (y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		return false, errors.WithStack(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		logrus.Infof("approved")
		return true, nil
	}
	logrus.Infof("declined")
	return false, nil
}
