# go-mullvad

Simple Am I Mullvad Wrapper written in Go.

Example

```

import ( "github.com/ktpx/go-mullvad" ; "fmt" )

func main () {

	m, _  := mullvad.GetMullvad()
	fmt.Printf("My ip is %s\n", m.Ipaddr)

	// A simple method returning m.Exitip
        fmt.Printf("I am Mullvad is %v!\n", m.Amimullvad() )	
}

```
