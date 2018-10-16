package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	SCTE35_2013 "github.com/chanyk-joseph/scte35_decoder/2013"
	SCTE35_2017 "github.com/chanyk-joseph/scte35_decoder/2017"
)

func parseSCTE2013(data []byte) error {
	var o SCTE35_2013.SCTE35
	n, err := o.DecodeFromRawBytes(data)
	if err != nil {
		return err
	}

	fmt.Println("Schema Version: ", o.SchemaVersion())
	fmt.Println("Parsed Bits: ", n)
	fmt.Println("Table ID: ", o.TableID)
	fmt.Println("splice_command_type: ", o.SpliceCommandType)
	fmt.Println("CRC32 In Hex: ", o.CRC32InHex)
	fmt.Println("\nEntire SCTE35 Structure: \n", o.JSON("	"))
	return nil
}

func parseSCTE2017(data []byte) error {
	var o SCTE35_2017.SCTE35
	n, err := o.DecodeFromRawBytes(data)
	if err != nil {
		return err
	}

	fmt.Println("Schema Version: ", o.SchemaVersion())
	fmt.Println("Parsed Bits: ", n)
	fmt.Println("Table ID: ", o.TableID)
	fmt.Println("splice_command_type: ", o.SpliceCommandType)
	fmt.Println("CRC32 In Hex: ", o.CRC32InHex)
	fmt.Println("\nEntire SCTE35 Structure: \n", o.JSON("	"))
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: scte35_decoder BASE64_ENCODED_DATA\n")
		os.Exit(1)
	}

	data, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal("invalid base64 string: " + err.Error())
	}

	if err := parseSCTE2017(data); err == nil {
		os.Exit(0)
	} else {
		log.Print("could not parse as SCTE35 2017: ", err)
	}

	if err := parseSCTE2013(data); err == nil {
		os.Exit(0)
	} else {
		log.Print("could not parse as SCTE35 2013: ", err)
		os.Exit(1)
	}
}
