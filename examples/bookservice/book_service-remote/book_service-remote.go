// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/wangxingge/thrift_clientpool/examples/bookservice"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Book GetBookById(string bookId)")
	fmt.Fprintln(os.Stderr, "  Book GetBookByName(string bookName)")
	fmt.Fprintln(os.Stderr, "   GetAllBooks()")
	fmt.Fprintln(os.Stderr, "  bool AddBook(Book bookInfo)")
	fmt.Fprintln(os.Stderr, "  bool RemoveBook(string bookId)")
	fmt.Fprintln(os.Stderr, "  bool DefaultKeepAlive(string clientId)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := bookservice.NewBookServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "GetBookById":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetBookById requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetBookById(value0))
		fmt.Print("\n")
		break
	case "GetBookByName":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetBookByName requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetBookByName(value0))
		fmt.Print("\n")
		break
	case "GetAllBooks":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAllBooks requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAllBooks())
		fmt.Print("\n")
		break
	case "AddBook":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddBook requires 1 args")
			flag.Usage()
		}
		arg17 := flag.Arg(1)
		mbTrans18 := thrift.NewTMemoryBufferLen(len(arg17))
		defer mbTrans18.Close()
		_, err19 := mbTrans18.WriteString(arg17)
		if err19 != nil {
			Usage()
			return
		}
		factory20 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt21 := factory20.GetProtocol(mbTrans18)
		argvalue0 := bookservice.NewBook()
		err22 := argvalue0.Read(jsProt21)
		if err22 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AddBook(value0))
		fmt.Print("\n")
		break
	case "RemoveBook":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveBook requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.RemoveBook(value0))
		fmt.Print("\n")
		break
	case "DefaultKeepAlive":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DefaultKeepAlive requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.DefaultKeepAlive(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
