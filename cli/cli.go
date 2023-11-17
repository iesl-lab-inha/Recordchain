package cli

import (
	"flag"
	"fmt"
	"log"
	"main/core"
	"main/network"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct {
	bc  *core.Recordchain
	bcd *core.DatNode
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  startnode -port PORT -type DATA_OR_HASH                   => starts a new node at give port")
	fmt.Println("  addrecord -hash RECORD_HASH -expire EXPIRATION_IN_SECONDS => add a record to the Recordchain")
	fmt.Println("  addrecord -hash RECORD_HASH -expire EXPIRATION_IN_SECONDS => add a record to the Recordchain")
	fmt.Println("  adddata -hash DATA_LOCATION -expire EXPIRATION_IN_SECONDS => add a record to the Recordchain")
	fmt.Println("  mineblock                                                 => add a block to the Recordchain")
	fmt.Println("  minepow                                                   => add a block to the Recordchain using Proof-of-work")
	fmt.Println("  printchain                                                => print all the blocks of the Recordchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) startNode(port string, type_node string) {
	fmt.Printf("Starting Node at: %s\n", port)
	datanode := false
	if type_node == "data" {
		datanode = true
	}
	network.StartServer(port, datanode)
}

func (cli *CLI) addRecord(hash string, expire int) {
	fmt.Println("Adding a new record...")
	network.SendRx("", []byte(hash), expire)
}

func (cli *CLI) addData(hash string, expire int) {

	fmt.Println("Add Data Successfull!")
}

func (cli *CLI) mineBlock() {

	fmt.Println("Mine Block Successfull!")
	// cli.bc.AddBlock(data)
	// fmt.Println("Success!")
}

func (cli *CLI) mineBlockPOW() {

	fmt.Println("Mine Block POW Successfull!")
	// cli.bc.AddBlock(data)
	// fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	// fmt.Println("printChain Successfull!")
	bci := cli.bcd.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("\n\nPrev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Records size: %d\n", len(block.Records))
		fmt.Printf("Hash: %x\n", block.Hash)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)
	addRecordCmd := flag.NewFlagSet("addrecord", flag.ExitOnError)
	addDataCmd := flag.NewFlagSet("adddata", flag.ExitOnError)
	mineBlockCmd := flag.NewFlagSet("mineblock", flag.ExitOnError)
	minePOWCmd := flag.NewFlagSet("minepow", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	startNodePort := startNodeCmd.String("port", "", "Port number")
	startNodeType := startNodeCmd.String("type", "", "Node type: data or hash")

	switch os.Args[1] {
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addrecord":
		err := addRecordCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "adddata":
		err := addDataCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "mineblock":
		err := mineBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "minepow":
		err := minePOWCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	// if addBlockCmd.Parsed() {
	// 	if *addBlockData == "" {
	// 		addBlockCmd.Usage()
	// 		os.Exit(1)
	// 	}
	// 	cli.addBlock(*addBlockData)
	// }

	if startNodeCmd.Parsed() {
		if *startNodePort == "" || *startNodeType == "" {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(*startNodePort, *startNodeType)
	}

	if addRecordCmd.Parsed() {
		cli.addRecord("", 0)
	}

	if addDataCmd.Parsed() {
		cli.addData("", 0)
	}

	if mineBlockCmd.Parsed() {
		cli.mineBlock()
	}

	if minePOWCmd.Parsed() {
		cli.mineBlockPOW()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
