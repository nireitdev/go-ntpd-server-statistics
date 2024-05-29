package traffic

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/net/context"
	"log"
	"time"
)

type ScanInfo struct {
	//IPorigin string
	Count int
}
type Traffic struct {
	Ctx       context.Context
	Device    string
	Ip        string
	Timerange int
}

var (
	ChanInfo     chan ScanInfo
	IPcapturadas map[string]int
)

func NewCapture(t Traffic) chan ScanInfo {
	ChanInfo = make(chan ScanInfo)
	go t.capture()
	return ChanInfo
}

func (t *Traffic) capture() {

	// Inicio captura de trafico:
	handle, err := pcap.OpenLive(t.Device, 1024, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		handle.Close()
		close(ChanInfo)
	}()

	//Creo el filtro de trafico para detectar entradas al puerto 123 UDP
	filter := fmt.Sprintf(" dst host %s and not src host %s and udp port 123", t.Ip, t.Ip)

	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	IPcapturadas = make(map[string]int)


	//Capturo paquetes y analizo si es ip, tcp /udp.
	var ethLayer layers.Ethernet
	var ipLayer layers.IPv4
	var udpLayer layers.UDP

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	timelast := time.Now()
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&ipLayer,
			&udpLayer,
		)
		foundLayerTypes := []gopacket.LayerType{}

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			//algun protocolo que no me interesa analizar
			//log.Println("Ignoring traffic:", err)
		}

		var ipsrc string
		for _, layerType := range foundLayerTypes {
			if layerType == layers.LayerTypeIPv4 {
				//ipLayer.SrcIP, ipLayer.DstIP
				ipsrc = ipLayer.SrcIP.String()
			}
		}

		IPcapturadas[ipsrc]++

		if time.Since(timelast) > time.Duration(t.Timerange)*time.Second {

			ChanInfo <- ScanInfo{len(IPcapturadas)}
			clear(IPcapturadas) //clear map[]  Go version > 1.21
			timelast = time.Now()

		}

	}
}
