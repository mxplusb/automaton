// Copyright © 2017 Mike Lloyd <kevin.michael.lloyd@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package initialiser

import (
	"os"
	"log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewWorkspace(cloneURL string, update bool) {
	ws := new(WorkSpace)
	ws.logger = logrus.New()
	ws.ManifestURL = cloneURL
	location, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}
	ws.Parent = &Parent{
		Location: location,
		Update:   update,
	}
	localViper := viper.GetViper()
}

func NewUser(user, email string) {
	localViper := viper.GetViper()
	localViper.Set("user", user)
	localViper.Set("email", email)
}
