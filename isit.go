package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
	Talking directly to whois can be interesting,
	but why not just "proxy" the requests through
	a (hopefully) trust worthy source?

	One of the possible benefits of doing this could be
	to help circumvent whois rate-limiting.

	Note: in the future I would like to support different
	whois backends -- via directly making a traditional whois
	query, or through other sources like the whois.com with command-line
	flags to be able to change up or just compare your findings with
	other sources.
*/
const WhoisDotComLink = "https://www.whois.com/whois/"

/*
	To help avoid some common problems when working directly with
	the native golang net/http package, we can customize the Trasport
	a bit to provide a sane-ish timeout.

	Note: in the future, custom transport options could be supported
	at the command-line to fine-tune the requests happening under the hood.
*/
var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

// The client which will make the http requests.
var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

// Check is a given domain has been registered.
func IsRegistered(domain string) bool {
	resp, err := netClient.Get(WhoisDotComLink + domain)
	if err != nil {
		return false // possibly bad/false assumption?
		//fmt.Printf("%s", err)
		//os.Exit(1)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(bodyBytes), "Registrar") {
		return true
	}
	return false
}

// Check is a given domain is available ( to buy ).
func IsAvailable(domain string) bool {
	if IsRegistered(domain) {
		return false
	}
	return true
}

// Check is a given domain is resolvable ( to an IP address ).
func IsResolvable(domain string) bool {
	_, err := net.LookupHost(domain)
	if err != nil {
		return false
	}
	return true
}

// Failure function to fire off when there are now command-line arguments.
func noArgumentGiven() {
	fmt.Println("no domain given!")
	os.Exit(1)
}

// A struct to hold the results for a query of any type.
type Result struct {
	domain string
	value  bool
}

func main() {
	app := cli.NewApp()

	app.Name = "isit"
	app.Version = "1.0.0"
	app.Usage = "domain availability command-line utility"

	app.Commands = []cli.Command{
		{
			Name:    "available",
			Aliases: []string{"a"},
			Usage:   "check if the given domain(s) are available",
			Action: func(c *cli.Context) error {
				results := make(chan Result)
				argumentCount := len(c.Args())
				if argumentCount > 0 {
					for i := 0; i < argumentCount; i++ {
						go func(c *cli.Context, index int) {
							domain := c.Args().Get(index)
							results <- Result{domain: domain, value: IsAvailable(domain)}
						}(c, i)
					}
					for i := 0; i < argumentCount; i++ {
						result := <-results
						fmt.Println(result.value, "\t", result.domain)
					}
				} else {
					noArgumentGiven()
				}
				return nil
			},
		},
		{
			Name:    "registered",
			Aliases: []string{"r"},
			Usage:   "check if the given domain(s) are registered",
			Action: func(c *cli.Context) error {
				results := make(chan Result)
				argumentCount := len(c.Args())
				if argumentCount > 0 {
					for i := 0; i < argumentCount; i++ {
						go func(c *cli.Context, index int) {
							domain := c.Args().Get(index)
							results <- Result{domain: domain, value: IsRegistered(domain)}
						}(c, i)
					}
					for i := 0; i < argumentCount; i++ {
						result := <-results
						fmt.Println(result.value, "\t", result.domain)
					}
				} else {
					noArgumentGiven()
				}
				return nil
			},
		},
		{
			Name:    "resolvable",
			Aliases: []string{"R"},
			Usage:   "check if the given domain(s) are resolvable",
			Action: func(c *cli.Context) error {
				results := make(chan Result)
				argumentCount := len(c.Args())
				if argumentCount > 0 {
					for i := 0; i < argumentCount; i++ {
						go func(c *cli.Context, index int) {
							domain := c.Args().Get(index)
							results <- Result{domain: domain, value: IsResolvable(domain)}
						}(c, i)
					}
					for i := 0; i < argumentCount; i++ {
						result := <-results
						fmt.Println(result.value, "\t", result.domain)
					}
				} else {
					noArgumentGiven()
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
