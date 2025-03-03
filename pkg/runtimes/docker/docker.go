/*
Copyright © 2020-2021 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package docker

import (
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

type Docker struct{}

// ID returns the identity of the runtime
func (d Docker) ID() string {
	return "docker"
}

// GetHost returns the docker daemon host
func (d Docker) GetHost() string {
	dockerHost := os.Getenv("DOCKER_HOST")
	url, err := url.Parse(dockerHost)
	if err != nil {
		return ""
	}
	log.Debugf("DockerHost: %s", url.Host)
	return url.Host
}

// GetRuntimePath returns the path of the docker socket
func (d Docker) GetRuntimePath() string {
	return "/var/run/docker.sock"
}
