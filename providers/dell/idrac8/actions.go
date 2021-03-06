package idrac8

import (
	"fmt"
	"strings"
)

// PowerCycle reboots the machine via bmc
func (i *IDrac8) PowerCycle() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm serveraction hardreset")
	if err != nil {
		return false, fmt.Errorf(output)
	}
	if strings.Contains(output, "successful") {
		return true, err
	}

	return status, fmt.Errorf(output)
}

// PowerCycleBmc reboots the bmc we are connected to
func (i *IDrac8) PowerCycleBmc() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm racreset hard")
	if err != nil {
		return false, fmt.Errorf(output)
	}
	if strings.Contains(output, "initiated successfully") {
		return true, err
	}

	return status, fmt.Errorf(output)
}

// PowerOn power on the machine via bmc
func (i *IDrac8) PowerOn() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm serveraction powerup")
	if err != nil {
		return false, fmt.Errorf(output)
	}

	if strings.Contains(output, "successful") {
		return true, err
	}

	return status, fmt.Errorf(output)
}

// PowerOff power off the machine via bmc
func (i *IDrac8) PowerOff() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm serveraction powerdown")
	if err != nil {
		return false, fmt.Errorf(output)
	}

	if strings.Contains(output, "successful") {
		return true, err
	}

	return status, fmt.Errorf(output)
}

// PxeOnce makes the machine to boot via pxe once
func (i *IDrac8) PxeOnce() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm config -g cfgServerInfo -o cfgServerBootOnce 1")
	if err != nil {
		return false, fmt.Errorf(output)
	}
	if strings.Contains(output, "successful") {
		output, err = i.sshClient.Run("racadm config -g cfgServerInfo -o cfgServerFirstBootDevice PXE")
		if err != nil {
			return false, fmt.Errorf(output)
		}
		if strings.Contains(output, "successful") {
			return i.PowerCycle()
		}
	}

	return status, fmt.Errorf(output)
}

// IsOn tells if a machine is currently powered on
func (i *IDrac8) IsOn() (status bool, err error) {
	err = i.sshLogin()
	if err != nil {
		return status, err
	}

	output, err := i.sshClient.Run("racadm serveraction powerstatus")
	if err != nil {
		return false, fmt.Errorf("%v: %v", err, output)
	}
	if strings.Contains(output, "Server power status: ON") {
		return true, err
	}

	return status, fmt.Errorf(output)
}
