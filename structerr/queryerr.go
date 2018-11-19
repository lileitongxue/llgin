package structerr

import "fmt"

type NsAndIp struct {
	Ns string
	IP string
}

func (err NsAndIp) Error() string {
	return fmt.Sprintf("param %s is invalid, value is %s", err.Ns, err.IP)
}
