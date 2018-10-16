# SCTE35 Decoder
scte35_decoder is a raw bytes parser for SCTE35 signal based on SCTE35 2013/2017 schema:

* 2013: [http://www.scte.org/documents/pdf/standards/ANSI_SCTE%2035%202013.pdf](http://www.scte.org/documents/pdf/standards/ANSI_SCTE%2035%202013.pdf)

* 2017: [http://www.scte.org/SCTEDocs/Standards/PublicReview/SCTE%2035%202017.pdf](http://www.scte.org/SCTEDocs/Standards/PublicReview/SCTE%2035%202017.pdf) 


## Usage

### command line tool

```
go build

./scte_decoder BASE64_ENCODED_DATA
```

### use as library

```go
func parseSCTE2017(data []byte) {	
	var o SCTE35_2017.SCTE35
	if n, err = o.DecodeFromRawBytes(data); err == nil {
	  fmt.Println("Schema Version: ", o.SchemaVersion())
	  fmt.Println("Table ID: ", o.TableID)
	  fmt.Println("splice_command_type: ", o.SpliceCommandType)
	  fmt.Println("CRC32 In Hex: ", o.CRC32InHex)
	  fmt.Println("Entire SCTE35 Structure: \n", o.JSON(" "))
	} else {
	  log.Fatal(err)
   }
}

func parseSCTE2013(data []byte) {	
	var o SCTE35_2013.SCTE35
	if n, err = o.DecodeFromRawBytes(data); err == nil {
	  fmt.Println("Schema Version: ", o.SchemaVersion())
	  fmt.Println("Table ID: ", o.TableID)
	  fmt.Println("splice_command_type: ", o.SpliceCommandType)
	  fmt.Println("CRC32 In Hex: ", o.CRC32InHex)
	  fmt.Println("Entire SCTE35 Structure: \n", o.JSON(" "))
	} else {
	  log.Fatal(err)
	}
}
```