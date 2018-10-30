package main

import (
	"log"
	"strconv"
)

func main() {
	cfg := newConfig() // read configuration from os env vars
	err := (error)(nil)
	// validate application config
	if err = cfg.check(); err != nil {
		log.Fatal(err)
	}

	// connect to guacd
	conn, err := OpenGuacdConn(cfg.GuacdHost, cfg.GuacdPort)
	if err != nil {
		log.Fatal(err)
	}
	connClose := CloseGuacdConn(conn)
	defer connClose()

	// begin handshake
	data := EncodeInstructions([]string{"select", "ssh"})
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	if data, err = ConnReadWholeInstruction(conn); err != nil {
		log.Fatal(err)
	}

	serverInstructions := DecodeInstructions(data)
	if serverInstructions[0] != "args" {
		log.Fatalf("bad guacd response opcode: %q, whant: 'args'", serverInstructions[0])
	}

	data = EncodeInstructions([]string{
		"size", // x/y dots, dpi
		strconv.FormatUint(uint64(cfg.ResX), 10),
		strconv.FormatUint(uint64(cfg.ResY), 10),
		strconv.FormatUint(uint64(cfg.DPI), 10),
	})
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	data = EncodeInstructions([]string{"audio", cfg.AudioMimeType})
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	data = EncodeInstructions([]string{"video", cfg.VideoMimeType})
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	data = EncodeInstructions([]string{"image", cfg.ImageMimeType})
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	connectInstructions := make([]string, 0, (len(serverInstructions) - 1))
	connectInstructions = append(connectInstructions, "connect")
	for i := 1; i < len(serverInstructions); i++ {
		parameter := ""
		switch serverInstructions[i] {
		case "hostname":
			parameter = cfg.SSHHost
		case "port":
			parameter = strconv.FormatUint(uint64(cfg.SSHPort), 10)
		case "username":
			parameter = cfg.SSHUser
		case "password":
			parameter = cfg.SSHPassword
		}
		connectInstructions = append(connectInstructions, parameter)
	}

	data = EncodeInstructions(connectInstructions)
	if _, err = conn.Write(data); err != nil {
		log.Fatal(err)
	}

	if data, err = ConnReadWholeInstruction(conn); err != nil {
		log.Fatal(err)
	}

	serverInstructions = DecodeInstructions(data)
	if serverInstructions[0] != "ready" {
		log.Fatalf("bad guacd response opcode: %q, whant: 'ready'", serverInstructions[0])
	}
	log.Printf("handshake completed, connection id: %q", serverInstructions[1])

	// read some data
	if data, err = ConnReadWholeInstruction(conn); err != nil {
		log.Fatal(err)
	}

	log.Printf("read some data: %q", string(data))

	return
}
