package workflowPattern

import (
	"fmt"
)

// PGraphic is Subject
type PGraphic interface {
	Draw()
}

// PImage is RealSubjec
type PImage struct {
	FileName string
}

// Draw is Request()
func (im PImage) Draw() {
	fmt.Println("draw " + im.FileName)
}

// PImageProxy is Proxy
type PImageProxy struct {
	FileName string
	_PImage  *PImage
}

// GetPImage creates an Subject
func (ip PImageProxy) GetPImage() *PImage {
	if ip._PImage == nil {
		ip._PImage = &PImage{ip.FileName}
	}
	return ip._PImage
}

// Draw is Request()
func (ip PImageProxy) Draw() {
	ip.GetPImage().Draw()
}

func ProxyDemo() {
	//Client
	proxy := PImageProxy{FileName: "1.png"}
	//operation without creating a RealSubject
	fileName := proxy.FileName
	//forwarded to the RealSubject
	proxy.Draw()

	fmt.Println(fileName)
}
