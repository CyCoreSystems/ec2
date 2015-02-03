// Amazon EC2 utility package
package ec2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const timeout time.Duration = 1 * time.Second

// get is a utility wrapper for handling http get calls
// and reading their response bodies to a string
func get(url string) (string, error) {
	// Create a new http.Client with a timeout
	client := &http.Client{Timeout: timeout}

	// Make the GET request
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid response: %v %v", resp.StatusCode, resp.Status)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to parse response: %v", resp.Body)
	}
	if len(bodyBytes) < 1 {
		return "", fmt.Errorf("No response received")
	}
	return bytes.NewBuffer(bodyBytes).String(), nil
}

// GetPublicIPv4 returns the IPv4 public IP of the current instance
func GetPublicIPv4() (string, error) {
	return get("http://169.254.169.254/latest/meta-data/public-ipv4")
}

// GetPublicName returns the IPv4 public hostname of the current instance
func GetPublicName() (string, error) {
	return get("http://169.254.169.254/latest/meta-data/public-hostname")
}

// GetPrivateIPv4 returns the IPv4 private IP of the current instance
func GetPrivateIPv4() (string, error) {
	return get("http://169.254.169.254/latest/meta-data/local-ipv4")
}

// GetPrivateName returns the IPv4 private hostname of the current instance
func GetPrivateName() (string, error) {
	return get("http://169.254.169.254/latest/meta-data/local-hostname")
}
