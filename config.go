package main

import (
	"errors"
	"os"
	"strconv"
)

type config struct {
	GuacdHost string
	GuacdPort uint16

	ResX uint16
	ResY uint16
	DPI  uint16

	AudioMimeType string
	VideoMimeType string
	ImageMimeType string

	SSHHost     string
	SSHPort     uint16
	SSHUser     string
	SSHPassword string
}

// newConfig reads application parameters from os environment variables
// populate and return application config struct
func newConfig() config {
	c := config{}
	c.GuacdHost = os.Getenv("GUACD_HOST")

	v, _ := strconv.ParseUint(os.Getenv("GUACD_PORT"), 10, 16)
	c.GuacdPort = uint16(v)

	v, _ = strconv.ParseUint(os.Getenv("RES_X"), 10, 16)
	c.ResX = uint16(v)
	if c.ResX < 1 {
		c.ResX = 1024 // set default value
	}

	v, _ = strconv.ParseUint(os.Getenv("RES_Y"), 10, 16)
	c.ResY = uint16(v)
	if c.ResY < 1 {
		c.ResY = 768 // set default value
	}

	v, _ = strconv.ParseUint(os.Getenv("DPI"), 10, 16)
	c.DPI = uint16(v)
	if c.DPI < 1 {
		c.DPI = 96 // set default value
	}

	c.AudioMimeType = os.Getenv("AUDIO_MIMETYPE")
	c.VideoMimeType = os.Getenv("VIDEO_MIMETYPE")
	c.ImageMimeType = os.Getenv("IMAGE_MIMETYPE")

	c.SSHHost = os.Getenv("SSH_HOST")
	v, _ = strconv.ParseUint(os.Getenv("SSH_PORT"), 10, 16)
	c.SSHPort = uint16(v)
	c.SSHUser = os.Getenv("SSH_USER")
	c.SSHPassword = os.Getenv("SSH_PASSWORD")
	return c
}

// chack function is used to validate application parameters
func (c *config) check() error {
	if c.GuacdHost == "" {
		return errors.New("GUACD_HOST is required")
	}
	if c.GuacdPort < 1 {
		return errors.New("GUACD_PORT is required")
	}
	if c.SSHHost == "" {
		return errors.New("SSH_HOST is required")
	}
	if c.SSHPort < 1 {
		return errors.New("SSH_PORT is required")
	}
	if c.SSHUser == "" {
		return errors.New("SSH_USER is required")
	}
	if c.SSHPassword == "" {
		return errors.New("SSH_PASSWORD is required")
	}
	return nil
}
