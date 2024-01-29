package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,  hasMX,  hasSPF,Spf records  ,hasDMARC ,DMARC records")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read from the scanner %v\n", err)
	}

}
func checkdomain(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var Spfrecords, DMARCrecords string

	mxrecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error %v", err)
	}
	if len(mxrecords) > 0 {
		hasMX = true
	}

	txtrecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error %v", err)
	}
	for _, records := range txtrecords {
		if strings.HasPrefix(records, "v=spf1") {
			hasSPF = true
			Spfrecords = records
			break
		}
	}
	dmarcrecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v", err)
	}
	for _, dmr := range dmarcrecords {
		if strings.HasPrefix(dmr, "v=DMARC1") {
			hasDMARC = true
			DMARCrecords = dmr
			break
		}
	}
	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, Spfrecords, hasDMARC, DMARCrecords)

}
