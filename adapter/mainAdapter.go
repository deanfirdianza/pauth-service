package adapter

import "fmt"

func Adapter() {

	client := &client{}
	mac := &mac{}

	fmt.Printf("client isinya alamat sebuah struct: %+v\n", client)
	fmt.Printf("&client isinya alamat: %+v\n", &client)
	fmt.Printf("*client isinya struct asli : %+v\n", *client)
	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}

func Test() {
	fmt.Printf("test")
}
